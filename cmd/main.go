package main

import (
	zjwt "ClothMall/jwt"
	sw "ClothMall/switcher"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	// ue "ginbase/ueditor"
	upload "ClothMall/upload"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/olahol/melody"
	"github.com/streadway/amqp"
	"github.com/zxnhl702/zConfig"
	// "gopkg.in/yaml.v2"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-sqlite3"
)

// 常量
const (
	// MainFileFolderName main.go文件夹名称
	MainFileFolderName = "cmd"
	// 禁用JWT
	JWTDeactived = "0"
)

// Rtn 返回参数
type Rtn struct {
	Success bool        `json:"success"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

// 是否开启测试函数的开关 开发时使用 正式环境应该用不到
var test bool

// 是否开启定时任务的开关
var cron bool

// 配置文件的文件路径+文件名
var cfgfile string

// 指定服务运行的端口号
var port int

// 配置信息
var cfg *zConfig.ServiceConfig

// 常量定义
var consts map[string]string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	// 读取命令行参数
	flag.BoolVar(&test, "t", false, "test flag")
	flag.BoolVar(&cron, "c", false, "定时任务开关")
	flag.StringVar(&cfgfile, "config", "./config/config.yaml", "指定配置文件")
	flag.IntVar(&port, "p", 0, "指定服务运行的端口号,只生效在http服务中,优先级高于配置文件中的端口号")
	flag.Parse()

	// 程序运行的当前路径
	currentPath, _ := os.Getwd()
	// 根据程序运行的位置调整配置文件的相对路径 保证在cmd文件夹中go run的时候可以访问到配置文件
	if strings.Index(currentPath, MainFileFolderName) >= 0 {
		cfgfile = ".." + cfgfile[strings.Index(cfgfile, "/"):]
	}
	log.Printf("定位到配置文件路径:%s", cfgfile)

	// 读取服务配置
	cfg, err = zConfig.NewServiceConfig(cfgfile)
	dealError(err, "配置文件出错")
	// 读取配置文件中的常量
	consts = cfg.Configs
	sw.ConstVar = cfg.Configs
	// 单独设置了端口号
	if port != 0 {
		// 判断单独设置的端口号是否与配置文件中的端口号重复
		for _, httpPort := range cfg.HTTP.Port {
			// 是否与配置文件中的http服务端口号重复
			if httpPort == port {
				log.Printf("-p %d 与配置文件中http服务的端口号%d重复 不做任何操作", port, httpPort)
				break
			}
		}
		for _, httpsPort := range cfg.HTTPS.Port {
			// 是否与配置文件中的https服务端口号重复
			if httpsPort == port {
				log.Fatalf("-p %d 与配置文件中https的端口号%d冲突 请重新指定端口号再运行程序", port, httpsPort)
			}
		}
		// 改写需要生效的端口号为单独设置的端口号
		cfg.HTTP.Port = []int{port}
	}
	if !test && !cron {
		// 开启服务
		server(cfg)
	} else {
		// 连接数据库
		db := ConnectDB(cfg.Database, "default")
		// 异常情况处理
		defer panicHanlderII(nil, db)
		if cron {
			// 定时任务
			sw.TimerEntry(db)
		} else if test {
			// 测试代码
			sw.TestEntry(db)
		}
	}
}

// 根据配置文件开启服务
func server(cfg *zConfig.ServiceConfig) {
	s := make(chan int)
	// http服务
	for _, p := range cfg.HTTP.Port {
		// 开启http服务
		go serverHTTP(cfg, strconv.Itoa(p))
	}
	// https服务需要的证书文件是否配置
	if "" != cfg.HTTPS.CertFile && "" != cfg.HTTPS.KeyFile {
		for _, ps := range cfg.HTTPS.Port {
			// 开启https服务
			go serverHTTPS(cfg, strconv.Itoa(ps))
		}
	}
	<-s
}

// 根据配置文件开启http服务
func serverHTTP(cfg *zConfig.ServiceConfig, port string) {
	router := gin.Default()
	//    // 开启正式模式(默认是调试模式)
	//   gin.SetMode(gin.ReleaseMode)

	// 开启gzip中间件
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// 开启默认跨域
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// ueditor 需要单独引入一个header
	config.AddAllowHeaders("x_requested_with")
	router.Use(cors.New(config))

	// // melody websocket中间件
	// m := melody.New()

	// 路由定义
	// 登陆/获取jwt token的路由 需要根据每个项目单独写
	// router.GET("/login", zjwt.LoginHandler)
	// router.POST("/login", zjwt.LoginHandler)
	// 路由组
	routerGroup := router.Group("/")
	// 根据配置文件中的设置 确定是否开启jwt token验证
	if cfg.Configs["activeJWT"] != JWTDeactived {
		routerGroup.Use(zjwt.JwtValidate())
	}
	// 路由组中的路由定义
	// routerGroup.GET("/ping/:param", pingHandler)
	// routerGroup.GET("/api", secondHanlder)
	// routerGroup.GET("/api2", thirdHandler)
	// 前端显示接口
	routerGroup.POST("/api", mallHanlder)
	// 后端管理接口
	routerGroup.POST("/api2", manageHanlder)
	// routerGroup.GET("/ueditor/:module", ue.UeditorHandler)
	// routerGroup.POST("/upload", upload.SingleFileUploadHandler)
	routerGroup.POST("/multiupload", upload.MutiUploadHanlder)
	// routerGroup.POST("/ueditor/:module", ue.UeditorUploadHandler)
	// 静态文件的路由
	router.Use(static.Serve("/", static.LocalFile("public", true)))

	// // melody weboscket路由
	// router.GET("ws2", func(c *gin.Context) {
	//     m.HandleRequest(c.Writer, c.Request)
	// })
	// // melody websocket接受消息
	// m.HandleMessage(func(s *melody.Session, msg []byte) {
	//     // log.Println(msg)
	//     // m.Broadcast(msg)
	//     mqpublisherHandler(m, "logs_direct", "direct", msg, "chat")
	// })

	// go mqsubscriberHandler(m, "logs_direct", "direct", "chat", "exam")

	router.Run(":" + port)
}

// 根据配置文件开启https服务
func serverHTTPS(cfg *zConfig.ServiceConfig, port string) {
	router := gin.Default()

	// 开启gzip中间件
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// 开启默认跨域
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// ueditor 需要单独引入一个header
	config.AddAllowHeaders("x_requested_with")
	router.Use(cors.New(config))

	// 路由定义
	// 登陆/获取jwt token的路由 需要根据每个项目单独写
	router.GET("/token", zjwt.LoginHandler)
	router.POST("/token", zjwt.LoginHandler)
	// 路由组
	routerGroup := router.Group("/")
	// 根据配置文件中的设置 确定是否开启jwt token验证
	if cfg.Configs["activeJWT"] != JWTDeactived {
		routerGroup.Use(zjwt.JwtValidate())
	}
	// 路由组中的路由定义
	routerGroup.GET("/ping/:param", pingHandler)
	routerGroup.GET("/api", secondHanlder)
	routerGroup.GET("/api2", thirdHandler)
	// 静态文件的路由
	router.Use(static.Serve("/", static.LocalFile("public", true)))

	router.RunTLS(":"+port, cfg.HTTPS.CertFile, cfg.HTTPS.KeyFile)
}

func pingHandler(c *gin.Context) {
	name := c.Param("param")
	c.String(http.StatusOK, "pong %s", name)
}

// 连接指定数据库的句柄
func mallHanlder(c *gin.Context) {
	// 调用接口的cmd
	cmd := sw.GetParamWithPanic(c, "cmd")
	log.Printf("mallHanlder cmd:%s, method:%s, url:%s", cmd, c.Request.Method, c.Request.URL)
	// 连接数据库
	db := ConnectDB(cfg.Database, "default")
	// 异常情况处理
	defer panicHanlderII(c, db)
	// 调用cmd所指定的函数并获取返回值
	switcher := sw.MallDispatch(db)
	var data interface{}
	var msg string
	if Authorize(c) {
		msg, data = switcher[cmd](c)
	} else {
		log.Println("Not authorized!")
		panic("Not authorized!")
	}

	// 拼返回数据
	rtn := &Rtn{true, msg, data}
	// 返回请求的数据
	dataResponse(c, rtn)
}

// 连接指定数据库的句柄
func manageHanlder(c *gin.Context) {
	// 调用接口的cmd
	cmd := sw.GetParamWithPanic(c, "cmd")
	log.Printf("mallHanlder cmd:%s, method:%s, url:%s", cmd, c.Request.Method, c.Request.URL)
	// 连接数据库
	db := ConnectDB(cfg.Database, "default")
	// 异常情况处理
	defer panicHanlderII(c, db)
	// 调用cmd所指定的函数并获取返回值
	switcher := sw.ManageDispatch(db)
	var data interface{}
	var msg string
	if Authorize(c) {
		msg, data = switcher[cmd](c)
	} else {
		log.Println("Not authorized!")
		panic("Not authorized!")
	}

	// 拼返回数据
	rtn := &Rtn{true, msg, data}
	// 返回请求的数据
	dataResponse(c, rtn)
}

// 连接指定数据库的句柄
func secondHanlder(c *gin.Context) {
	// 调用接口的cmd
	cmd := sw.GetParamWithPanic(c, "cmd")
	log.Printf("secondHanlder cmd:%s, method:%s, url:%s", cmd, c.Request.Method, c.Request.URL)
	// 连接数据库
	db := ConnectDB(cfg.Database, "default")
	// 异常情况处理
	defer panicHanlderII(c, db)
	// 调用cmd所指定的函数并获取返回值
	switcher := sw.MallDispatch(db)
	var data interface{}
	var msg string
	if Authorize(c) {
		msg, data = switcher[cmd](c)
	} else {
		log.Println("Not authorized!")
		panic("Not authorized!")
	}

	// 拼返回数据
	rtn := &Rtn{true, msg, data}
	// 返回请求的数据
	dataResponse(c, rtn)
}

// 连接全部数据库的句柄
func thirdHandler(c *gin.Context) {
	// 调用接口的cmd
	cmd := sw.GetParamWithPanic(c, "cmd")
	log.Printf("thirdHandler cmd:%s, method:%s, url:%s", cmd, c.Request.Method, c.Request.URL)
	// 连接数据库
	dbs := ConnectDBs(cfg.Database)
	// 异常情况处理
	defer panicHanlder(c, dbs)

	// 调用cmd所指定的函数并获取返回值
	switcher := sw.MallDispatch(dbs[0])
	var data interface{}
	var msg string
	if Authorize(c) {
		msg, data = switcher[cmd](c)
	} else {
		panic("Not authorized!")
	}

	// 拼返回数据
	rtn := &Rtn{true, msg, data}
	// 返回请求的数据
	dataResponse(c, rtn)
}

// websocket句柄 TODO
func wsHandler(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
}

// rabbitmq 消息发布句柄
func mqpublisherHandler(m *melody.Melody, exchangename, exchangekind string, msg []byte, keys ...string) error {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if nil != err {
		log.Println(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if nil != err {
		log.Println(err)
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(exchangename, exchangekind, true, false, false, false, nil)
	if nil != err {
		log.Println(err)
		return err
	}

	for _, k := range keys {
		err = ch.Publish("logs_direct", k, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
		if nil != err {
			log.Println(err)
			return err
		}
	}

	log.Printf("[x] Send %s", msg)
	return nil
}

// rabbitmq 消息订阅句柄
func mqsubscriberHandler(m *melody.Melody, exchangename, exchangekind string, keys ...string) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	dealError(err, "fail to open rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	dealError(err, "fial to open channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchangename, exchangekind, true, false, false, false, nil)
	dealError(err, "fail to declare an exchange")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	dealError(err, "fail to declare a queue")

	for _, s := range keys {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", s)
		err = ch.QueueBind(q.Name, s, "logs_direct", false, nil)
		dealError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	dealError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			m.Broadcast(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

// 返回请求的数据
func dataResponse(c *gin.Context, rtn *Rtn) {
	_, exists := c.GetQuery("callback")
	// 有callback参数时 采用jsonp返回数据
	if exists {
		c.JSONP(http.StatusOK, rtn)
	} else {
		// 没有callback参数时 采用json返回数据
		c.JSON(http.StatusOK, rtn)
	}
}

// GetJsonpResult 返回数据处理成jsonp(老的写法 新版本直接调用gin的jsonp 废弃预定)
func GetJsonpResult(callback string, rtn *Rtn) string {
	respstr, err := json.Marshal(rtn)
	if nil != err {
		panic(err)
	}
	return callback + "(" + string(respstr) + ")"
}

// Authorize 鉴权参数
func Authorize(c *gin.Context) bool {
	token := sw.GetParamWithPanic(c, "token")
	return token == "Jh2044695"
}

// 异常情况的处理函数
func panicHanlder(c *gin.Context, dbs []*sql.DB) {
	// 关闭数据库
	for _, db := range dbs {
		db.Close()
	}
	// 获取异常信息
	err := recover()
	if nil != err {
		errMsg, ok := err.(string)
		if !ok {
			errMsg = "请求失败"
		}
		if nil != c {
			// 返回异常数据 结束请求
			dataResponse(c, &Rtn{false, errMsg, nil})
		}
	}
}

// 异常情况的处理函数
func panicHanlderII(c *gin.Context, db *sql.DB) {
	// 关闭数据库
	db.Close()
	// 获取异常信息
	err := recover()
	if nil != err {
		errMsg, ok := err.(string)
		if !ok {
			errMsg = "请求失败"
		}
		if nil != c {
			// 返回异常数据 结束请求
			dataResponse(c, &Rtn{false, errMsg, nil})
		}
	}
}

// main函数中出错  打印错误信息并退出
func dealError(err error, msg string) {
	if nil != err {
		log.Println(msg)
		log.Fatal(err)
	}
}

// ConnectSqlite3DB 连接sqlite3数据库
func ConnectSqlite3DB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	db.Exec("pragma journal_mode=wal")
	return db
}

// ConnectDB 连接数据库
func ConnectDB(dbconfigs []*zConfig.ServiceDB, key string) *sql.DB {
	// 数据库连接配置
	sqlite3SourcePattern := `file:%s?cache=shared&_journal_mode=wal`
	// mysqlSourcePattern := `%s:&s@tcp(%s)/%s`
	oracleSourcePattern := `%s:&s@%s/%s`
	postgresqlSourcePattern := `host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`
	sqlserverSourcePattern := `server=%s;port%d;database=%s;user id=%s;password=%s`
	for _, c := range dbconfigs {
		// 跳过不是需要连接的数据库
		if c.Key != key {
			continue
		}
		// 连接参数
		var dataSourceName string
		// 根据不同的数据库
		switch c.Type {
		case "sqlite3":
			dataSourceName = fmt.Sprintf(sqlite3SourcePattern, c.Host)
		case "mysql":
			dataSourceName = c.User + ":" + c.Pwd + "@tcp(" + c.Host + ")/" + c.Instance
			// dataSourceName = fmt.Sprintf(mysqlSourcePattern, c.User, c.Pwd, c.Host, c.Instance)
		case "oracle":
			dataSourceName = fmt.Sprintf(oracleSourcePattern, c.User, c.Pwd, c.Host, c.Instance)
		case "postgresql":
			dataSourceName = fmt.Sprintf(postgresqlSourcePattern, c.Host, strconv.Itoa(c.Port), c.User, c.Pwd, c.Instance)
		case "sqlserver":
			dataSourceName = fmt.Sprintf(sqlserverSourcePattern, c.Host, c.Port, c.Instance, c.User, c.Pwd)
		default:
			dataSourceName = ""
		}
		// log.Println(dataSourceName)
		// 连接数据库
		db, err := sql.Open(c.Driver, dataSourceName)
		dealError(err, "连接数据库失败")
		// sqlite3读写分离
		if "sqlite3" == c.Type {
			if "wal" == c.Mode {
				db.Exec("pragma journal_mode=wal")
			}
			db.SetMaxOpenConns(1)
		}
		// 返回数据库连接
		return db
	}
	// 全都没有匹配到 返回空
	return nil
}

// ConnectDBs 连接数据库
func ConnectDBs(dbconfigs []*zConfig.ServiceDB) []*sql.DB {
	var dbs []*sql.DB
	// 数据库连接配置
	sqlite3SourcePattern := `file:%s?cache=shared&_journal_mode=wal`
	mysqlSourcePattern := `%s:&s@tcp(%s)/%s`
	oracleSourcePattern := `%s:&s@%s/%s`
	postgresqlSourcePattern := `host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`
	sqlserverSourcePattern := `server=%s;port%d;database=%s;user id=%s;password=%s`
	for _, c := range dbconfigs {
		// 连接参数
		var dataSourceName string
		// 根据不同的数据库
		switch c.Type {
		case "sqlite3":
			dataSourceName = fmt.Sprintf(sqlite3SourcePattern, c.Host)
		case "mysql":
			dataSourceName = fmt.Sprintf(mysqlSourcePattern, c.User, c.Pwd, c.Host, c.Instance)
		case "oracle":
			dataSourceName = fmt.Sprintf(oracleSourcePattern, c.User, c.Pwd, c.Host, c.Instance)
		case "postgresql":
			dataSourceName = fmt.Sprintf(postgresqlSourcePattern, c.Host, c.Port, c.User, c.Pwd, c.Instance)
		case "sqlserver":
			dataSourceName = fmt.Sprintf(sqlserverSourcePattern, c.Host, c.Port, c.Instance, c.User, c.Pwd)
		default:
			dataSourceName = ""
		}
		log.Println(dataSourceName)
		// 连接数据库
		db, err := sql.Open(c.Driver, dataSourceName)
		dealError(err, "连接数据库失败")
		// sqlite3读写分离
		if "sqlite3" == c.Type {
			if "wal" == c.Mode {
				db.Exec("pragma journal_mode=wal")
			}
			db.SetMaxOpenConns(1)
		}
		dbs = append(dbs, db)
	}
	return dbs
}
