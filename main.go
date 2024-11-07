package main

import (
	"github.com/gin-gonic/gin"
	_ "golang.org/x/image/webp" // 支持 webp 图片格式
	"image/jpeg"
	"image/png"
	"log"
	"yushu/opt/file"
	"yushu/opt/pic"
)

func main() {
	r := gin.New()
	//conf := config.New()

	r.GET("/img-text", func(c *gin.Context) {
		imgUrl := c.Query("img")
		text := c.Query("text")
		fontFilePath := file.AppPath("./font-jianti.ttf")

		img, err := pic.ImageText(imgUrl, text, fontFilePath, 48)
		if err != nil {
			return
		}

		//addLabel(newImg, deText, 48)

		// 设置响应头，指定返回的内容类型
		switch img.Type {
		case "png":
			c.Header("Content-Type", "image/png")
			err = png.Encode(c.Writer, img.Img) // 将 PNG 图片写入响应
		case "jpeg", "jpg":
			c.Header("Content-Type", "image/jpeg")
			err = jpeg.Encode(c.Writer, img.Img, nil) // 将 JPEG 图片写入响应
		case "webp":
			c.Header("Content-Type", "image/webp")
			err = png.Encode(c.Writer, img.Img) // 对 webp 格式，用 png 编码返回
		default:
			c.String(400, "Unsupported image format")
			return
		}

		if err != nil {
			log.Printf("failed to encode image: %v", err)
			c.String(500, "Failed to encode image")
		}
	})

	r.GET("/acg/img", func(c *gin.Context) {
		imgUrl := "https://api.miaomc.cn/image/get"
		text := "i站:小林"
		fontFilePath := file.AppPath("./font-jianti.ttf")

		img, err := pic.ImageText(imgUrl, text, fontFilePath, 48)
		if err != nil {
			return
		}

		//addLabel(newImg, deText, 48)

		// 设置响应头，指定返回的内容类型
		switch img.Type {
		case "png":
			c.Header("Content-Type", "image/png")
			err = png.Encode(c.Writer, img.Img) // 将 PNG 图片写入响应
		case "jpeg", "jpg":
			c.Header("Content-Type", "image/jpeg")
			err = jpeg.Encode(c.Writer, img.Img, nil) // 将 JPEG 图片写入响应
		case "webp":
			c.Header("Content-Type", "image/webp")
			err = png.Encode(c.Writer, img.Img) // 对 webp 格式，用 png 编码返回
		default:
			c.String(400, "Unsupported image format")
			return
		}

		if err != nil {
			log.Printf("failed to encode image: %v", err)
			c.String(500, "Failed to encode image")
		}
	})

	r.GET("/my", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"data": "this data",
		})
	})

	r.Run(":8080")
}
