package ueditor

import (
    "bytes"
    "crypto/md5"
    "crypto/rand"
    "database/sql"
    "encoding/base64"
    "encoding/hex"
    "image"
    "image/draw"
    "image/jpeg"
    "image/png"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path"
    "strings"
    sw "ginbase/switcher"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

type Ret struct {
    Success bool        `json:"success"`
    ErrMsg  string      `json:"errMsg"`
    Data    interface{} `json:"data"`
}

// 初始化读取配置文件
func UeditorHandler(c *gin.Context) {
    log.Println(sw.GetParamWithPanic(c, "action"), c.Param("module"))
    file, err := os.Open("./ueditor/config.json")
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    defer file.Close()
    buf := bytes.NewBuffer(nil)
    buf.ReadFrom(file)
    // 通过callback判断是否是jsonp
    cb, exists := c.GetQuery("callback")
    if exists {
        // 返回jsonp
        c.String(http.StatusOK, cb + "(" + buf.String() + ")")
    } else {
        // 返回json
        c.String(http.StatusOK, buf.String())
    }
}

// 上传文件
func UeditorUploadHandler(c *gin.Context) {
    // log.Println(sw.GetParameter(r, "action"), p.ByName("module"))
    log.Println(sw.GetParamWithPanic(c, "action"), c.Param("module"))
    op := sw.GetParamWithPanic(c, "action")
    switch op {
    case "uploadimage":
        uploadFile(c)
    case "uploadscrawl":
        uploadScrawl(c)
    case "uploadvideo":
        uploadFile(c)
    case "uploadfile":
        uploadFile(c)
    }
}

// 上传图片、视频、文件
func uploadFile(c *gin.Context) {
    file, err := c.FormFile("upfile")
    if err != nil {
        log.Println(err)
        panic("上传失败")
    }

    // filename
    filename := file.Filename
    log.Println(filename)
    filename = GetGuid() + path.Ext(filename)

    // save file
    // err = c.SaveUploadedFile(file, sw.IMG_ROOT + "/" + c.Param("module") + "/" + filename)
    err = c.SaveUploadedFile(file, sw.ConstVar["imgRoot"] + "/" + c.Param("module") + "/" + filename)
    if err != nil {
        log.Println(err)
        panic("保存文件失败")
    }
    // 根据管理员编号确定是否需要添加图片水印
    adminId, err := c.Cookie("adminId")
    if nil == err {
        // filename = getWaterMarkedImg(sw.IMG_ROOT + "/" + c.Param("module") + "/" + filename, filename, adminId)
        filename = getWaterMarkedImg(sw.ConstVar["imgRoot"] + "/" + c.Param("module") + "/" + filename, filename, adminId)
    }
    c.JSON(http.StatusOK, map[string]string{
        // "url":      sw.IMG_ROOT_URL + c.Param("module") + "/" + filename,  //保存后的文件路径
        "url":      sw.ConstVar["imgRootUrl"] + c.Param("module") + "/" + filename,  //保存后的文件路径
        "title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
        "original": file.Filename,                                         //原始文件名
        "state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
    })
}

// 上传涂鸦
// P.S 上传涂鸦生成的图片只能网页端打开 直接电脑中打开。jpg文件会显示错误
func uploadScrawl(c *gin.Context) {
    str := sw.GetParam(c, "upfile")
    data, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        log.Println(err)
        panic("上传失败")
    }

    // filename
    filename := GetGuid() + ".jpg"

    // save file
    // err = ioutil.WriteFile(sw.IMG_ROOT + "/" + c.Param("module") + "/" + filename, data, 0777)
    err = ioutil.WriteFile(sw.ConstVar["imgRoot"] + "/" + c.Param("module") + "/" + filename, data, 0777)
    if err != nil {
        log.Println(err)
        panic("保存文件失败")
    }
    c.JSON(http.StatusOK, map[string]string{
        // "url":      sw.IMG_ROOT_URL + c.Param("module") + "/" + filename,  //保存后的文件路径
        "url":      sw.ConstVar["imgRootUrl"] + c.Param("module") + "/" + filename,  //保存后的文件路径
        "title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
        "original": filename,                                              //原始文件名
        "state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
    })
}

/**
 * 获取水印图片 如果不是图片文件或者当前帐号没有设置水印怎么返回原始图片
 * 参数
 *  filepath 原始图片路径
 *  filename 原始图片文件名
 *  adminId 管理员帐号编号
 * 返回
 *  stirng 带水印的图片文件名或者原始图片文件名
 */
func getWaterMarkedImg(filepath, filename, adminId string) string {
    // 文件扩展名
    ext := strings.ToLower(path.Ext(filepath))[1:]
    // 图片的情况下
    if ("jpg" == ext || "jpeg" == ext || "png" == ext) {
        wfile, wext := getWaterMarkFilepath(adminId)
        
        // 如果没有水印 直接返回没有水印的图片
        if "" == wfile && "" == wext {
            return filename
        } else {
            // 添加水印
            wmFilename := addWaterMark(filepath, ext, wfile, wext)
            os.Remove(filepath)
            log.Println(wmFilename, filepath)
            return wmFilename
        }
    } else {
        return filename
    }
}

/**
 * 添加水印
 * 参数
 *  filepath 带路径的文件名
 *  ext 文件扩展名
 *  wfile 水印文件
 *  wext 水印文件扩展名
 * 返回值
 *  string 添加水印之后的图片文件名
 */
func addWaterMark(filepath, ext, wfile, wext string) string {
    dir := path.Dir(filepath)
    // 读取图片
    src := loadImage(filepath, ext)
    // 水印图片
    wm := loadImage(wfile, wext)
    //把水印写到右下角，并向0坐标各偏移10个像素
    offset := image.Pt(src.Bounds().Dx()-wm.Bounds().Dx()-10, src.Bounds().Dy()-wm.Bounds().Dy()-10)
    // 绘制带水印的图片
    b := src.Bounds()
    m := image.NewRGBA(b)
    draw.Draw(m, b, src, image.ZP, draw.Src)
    draw.Draw(m, wm.Bounds().Add(offset), wm, image.ZP, draw.Over)
    // 生成文件名
    filename := GetGuid() + "." + ext
    // 编码缩放后的图片
    encodeImage(m, ext, dir + "/" + filename)
    return filename
}

/**
 * 编码缩放后的图片
 * 参数
 *  img 缩放后的图片对象
 *  ext 图片扩展名
 *  filepath 编码之后的图片文件名
 * 返回值
 *  无
 */
func encodeImage(img image.Image, ext, filename string) {
    f, err := os.Create(filename)
    defer f.Close()
    if nil != err {
        log.Println(err)
        panic("文件创建错误")
    }
    // 根据文件扩展名选择相应图片解码函数
    switch ext {
        case "jpg":
            err = jpeg.Encode(f, img, nil)
            break
        case "jpeg":
            err = jpeg.Encode(f, img, nil)
            break
        case "png":
            err = png.Encode(f, img)
            break
        default:
            panic("不支持" + ext + "文件格式")
    }
}

/**
 * 加载图片
 * 参数
 *  filepath 带路径的文件名
 *  ext 文件扩展名
 * 返回值
 *  image.Image 图片对象
 */
func loadImage(filepath, ext string) image.Image {
    // 打开文件
    f, err := os.Open(filepath)
    defer f.Close()
    if nil != err {
        log.Println(err)
        panic("文件打开错误")
    }
    var img image.Image
    // 根据文件扩展名选择相应图片解码函数
    switch ext {
        case "jpg":
            img, err = jpeg.Decode(f)
            break
        case "jpeg":
            img, err = jpeg.Decode(f)
            break
        case "png":
            img, err = png.Decode(f)
            log.Println(1)
            break
        default:
            panic("不支持" + ext + "文件格式")
    }
    if nil != err {
        log.Println(err)
        panic("文件打开错误")
    }
    return img
}

/**
 * 获取水印图片文件
 * 参数
 *  adminId 管理员编号
 * 返回值
 *  string 带路径的水印图片文件名  如果没有水印或者查询出错返回''
 *  string 水印图片文件扩展名  如果没有水印或者查询出错返回''
 */
func getWaterMarkFilepath(adminId string) (string, string) {
    // wm := "/home/zhangxiangnan/web-location/wxt/waterMark/wifizswatermark.png"
    // return wm, strings.ToLower(path.Ext(wm))[1:]
    db := ConnectDB("./middle.db")
    defer db.Close()
    var wm string
    err := db.QueryRow("select filepath from live_waterMark where adminId = ? limit 1", adminId).Scan(&wm)
    if nil != err {
        return "", ""
    } else {
        return wm, strings.ToLower(path.Ext(wm))[1:]
    }
}

func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

func GetGuid() string {
    b := make([]byte, 48)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        panic("创建文件名出错")
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func ConnectDB(dbPath string) *sql.DB {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        panic(err)
    }
    db.Exec("pragma journal_mode=wal")
    return db
}
