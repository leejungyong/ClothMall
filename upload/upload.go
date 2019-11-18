package upload

import (
	sw "ClothMall/switcher"
	"log"
	"path"

	"github.com/gin-gonic/gin"
)

// SingleFileUploadHandler 上传单个文件的句柄函数
func SingleFileUploadHandler(c *gin.Context) {
	// 异常情况下返回出错信息
	defer func() {
		err := recover()
		if err != nil {
			dataResponse(c, &Rtn{false, err.(string), nil})
		}
	}()

	var filename string
	var imgurl string

	form, _ := c.MultipartForm()
	for key := range form.File {
		// 上传的文件
		file, err := c.FormFile(key)
		sw.Perror(err, "上传失败")
		// 原始文件名
		filename = file.Filename
		// 文件重命名
		rename, exists := sw.GetParamExist(c, "rename")
		if exists {
			// 前端明确指定重命名的情况选  使用前端指定的文件名
			filename = rename
		} else {
			// 否则随机生成文件名
			filename = GetGUID() + path.Ext(filename)
		}
		// 文件保存路径
		dst := sw.ConstVar["imgRoot"] + "/" + key + "/" + filename
		log.Println("filename:", filename, "destination:", dst)
		// 保存文件
		err = c.SaveUploadedFile(file, dst)
		sw.Perror(err, "保存文件失败")
		imgurl = sw.ConstVar["imgRootUrl"] + key + "/" + filename
	}

	// 存储路径返回前端
	dataResponse(c, &Rtn{true, "上传成功", imgurl})
}
