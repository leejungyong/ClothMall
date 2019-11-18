package upload

import (
    "encoding/json"
    "image"
    "image/jpeg"
    "image/png"
    "log"
    "os"
    "os/exec"
    "path"
    "strconv"
    "strings"
    sw "ClothMall/switcher"

    "github.com/disintegration/gift"
    "github.com/gin-gonic/gin"
)

// MutiUploadHanlder 上传多个文件的句柄函数
func MutiUploadHanlder(c *gin.Context) {
    // 异常情况下返回出错信息
    defer func() {
        err := recover()
        if err != nil {
            dataResponse(c, &Rtn{false, err.(string), nil})
        }
    }()
    
    // 模块名称
    module := sw.GetParam(c, "module")
    var pList []Pictures
    // 遍历各个不同的key(既前端多个input file标签的dom)
    form, _ := c.MultipartForm()
    for key, handle := range form.File {
        // 遍历每个文件
        for _, h := range handle {
            log.Printf("key: %s || module: %s || file: %s", key, module, h.Filename)
            var p Pictures
            // 文件重命名
            filename := GetGUID() + path.Ext(h.Filename)
            // 文件存放路径
            // filepath := sw.IMG_ROOT + "/" + module + "/" + filename
            filepath := sw.ConstVar["imgRoot"] + "/" + module + "/" + filename
            // 文件扩展名
            ext := strings.ToLower(path.Ext(filename))[1:]
            // 保存文件
            err := c.SaveUploadedFile(h, filepath)
            sw.Perror(err, "保存文件失败")
            // 生成缩略图
            var s Size
            resize := sw.GetParam(c, key + "resize")
            // 前端设置需要压缩文件的参数时
            if "" != resize {
                // 解码json形式的参数
                err = json.Unmarshal([]byte(resize), &s)
                sw.Perror(err, "参数转换失败")
                // 根据上传的文件类型 选择不同的处理方式
                if "jpg" == ext || "jpeg" == ext || "png" == ext {
                    // 上传的文件是图片(仅限jpg, jpeg, png)
                    if 0 != s.Height || 0 != s.Width {
                        // 压缩后的高度或者宽度设置过的情况下
                        // 按照给定的高度 宽度压缩图片
                        p.Thumb = GenThumbnail(filepath, ext, s.Width, s.Height)
                        if 1 == s.Thumbonly {
                            // 只需要缩略图的情况下 删除原始图片
                            err = os.Remove(filepath)
                            sw.Perror(err, "删除原始图片失败")
                            filename = ""
                        }
                    } else {
                        // 没有设置压缩后的高度和宽度的情况下 不压缩图片 设置空
                        p.Thumb = ""
                    }
                } else if "mp4" == ext {
                    // 上传的文件是视频(仅限mp4)
                    if -1 != s.Height && -1 != s.Width {
                        // 压缩后的高度与宽度都不是-1的情况下 按照尺寸截图
                        p.Thumb = GenVideoThumb(filepath, ext, s.Width, s.Height)
                    } else {
                        p.Thumb = ""
                    } 
                } else {
                    // 上传的文件除了图片和视频的情况下  不压缩 设置空
                    p.Thumb = ""
                }
            } else {
                // 前端没有设置需要压缩文件的参数时 压缩后的文件名设置空
                p.Thumb = ""
            }
            // 设置上传文件名 key等参数 绝对地址
            p.Image = "http://121.199.54.130:8079/" + module + "/" + filename
            p.Thumb = "http://121.199.54.130:8079/" + module + "/" + p.Thumb
            p.Key = key
            pList = append(pList, p)
        }
    }
    // 数据返回前端
    dataResponse(c, &Rtn{true, "上传成功", pList})
}

// GenVideoThumb 视频文件中截取缩略图
// 参数
//  filepath 带路径的文件名
//  ext 文件扩展名
//  width 缩放的宽度
//  height 缩放的高度
//   高度或者宽度为0的情况下 根据视频原始尺寸截取
// 返回值
//  string 缩放后的图片文件名 与参数中的filepath路径相同
 func GenVideoThumb(filepath, ext string, width, height int) string {
    dir := path.Dir(filepath)
    // 生成缩略图文件名
    filename := GetGUID() + ".jpg"
    // 宽度高度全都不为0
    var cmd *exec.Cmd
    if width > 0 && height > 0 {
        // 按照高宽截缩略图
        w := strconv.Itoa(width)
        h := strconv.Itoa(height)
        cmd = exec.Command("ffmpeg", "-i", filepath, "-y", "-f", "image2", "-ss", "3", "-vframes", "1", "-s", w+"*"+h, dir+"/"+filename)
        // 其他情况
    } else {
        // 按照视频原始尺寸截缩略图
        cmd = exec.Command("ffmpeg", "-i", filepath, "-y", "-f", "image2", "-ss", "3", "-vframes", "1", dir+"/"+filename)
    }
    err := cmd.Run()
    sw.Perror(err, "视频截取缩略图失败")
    return filename
}

// GenThumbnail 生成缩略图
// 参数
//  filepath 带路径的文件名
//  ext 文件扩展名
//  width 缩放的宽度
//  height 缩放的高度
//   高度或者宽度为0的情况下 根据宽度或者高度等比例缩放
// 返回值
//  string 缩放后的图片文件名 与参数中的filepath路径相同
func GenThumbnail(filepath, ext string, width, height int) string {
    // 缩放高度 宽度检查
    if height < 0 || width < 0 || (0 == height && 0 == width) {
        panic("高度或者宽度设置错误")
    }
    dir := path.Dir(filepath)
    // 读取图片
    src, ext2 := loadImage(filepath)
    // 缩放图片
    dst := resizeImage(src, width, height)
    // 生成文件名
    filename := GetGUID() + "." + ext2
    // 编码缩放后的图片
    encodeImage(dst, ext2, dir+"/"+filename)
    return filename
}

// 编码缩放后的图片
// 参数
//  img 缩放后的图片对象
//  ext 图片扩展名
//  filepath 编码之后的图片文件名
// 返回值
//  无
func encodeImage(img image.Image, ext, filename string) {
    f, err := os.Create(filename)
    defer f.Close()
    sw.Perror(err, "文件创建错误")
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


// 缩放图片
// 参数
//  src 需要缩放的图片对象
//  width 缩放的宽度
//  height 缩放的高度
//   高度或者宽度为0的情况下 根据宽度或者高度等比例缩放
// 返回值
//  image,Image 缩放后的图片对象
//
func resizeImage(src image.Image, width, height int) image.Image {
    // 图片缩放的过滤器
    g := gift.New(
        gift.Resize(width, height, gift.LanczosResampling),
    )
    // 缩放后的图片对象
    dst := image.NewRGBA(g.Bounds(src.Bounds()))
    // 缩放图片
    g.Draw(dst, src)
    return dst
}

// 
//  * 加载图片
//  * 参数
//  *  filepath 带路径的文件名
//  *  ext 文件扩展名
//  * 返回值
//  *  image.Image 图片对象
//  */
// func loadImage(filepath, ext string) image.Image {
//     // 打开文件
//     f, err := os.Open(filepath)
//     defer f.Close()
//     sw.Perror(err, "文件打开错误")
//     var img image.Image
//     // 根据文件扩展名选择相应图片解码函数
//     switch ext {
//     case "jpg":
//         img, err = jpeg.Decode(f)
//         break
//     case "jpeg":
//         img, err = jpeg.Decode(f)
//         break
//     case "png":
//         img, err = png.Decode(f)
//         log.Println(1)
//         break
//     default:
//         panic("不支持" + ext + "文件格式")
//     }
//     sw.Perror(err, "文件打开错误")
//     return img
// }


// 加载图片
// 参数
//  filepath 带路径的文件名
// 返回值
//  image.Image 图片对象
//  string 图片的扩展名
//
 func loadImage(filepath string) (image.Image, string) {
    // 打开文件
    f, err := os.Open(filepath)
    defer f.Close()
    sw.Perror(err, "文件打开错误")
    var img image.Image
    img, format, err := image.Decode(f)
    sw.Perror(err, "文件打开错误")
    return img, format
}