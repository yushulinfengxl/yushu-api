package pic

import (
	"errors"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ImageResult struct {
	Text   string      `json:"text"`
	ImgUrl string      `json:"img_url"`
	Img    *image.RGBA `json:"img"`
	Type   string      `json:"type"`
}

func ImageText(imgUrl, text, fontFilePath string, fontSize int) (res ImageResult, err error) {

	if imgUrl == "" || text == "" {
		err = errors.New("image URL and text are required")
		return
	}

	// 下载远程图片
	resp, err := http.Get(imgUrl)
	if err != nil {
		err = errors.New("failed to download image: " + err.Error())
		return
	}
	defer resp.Body.Close()

	// 检查 Content-Type 确定是否为图片
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		err = errors.New("invalid content type, expected image: " + contentType)
		return
	}

	// 解析图片
	img, format, err := image.Decode(resp.Body)
	if err != nil {
		err = errors.New("failed to decode image: " + err.Error())
		return
	}

	// 创建新的 RGBA 图像
	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{}, draw.Src)

	// 解码 URL 参数
	deText, err := url.QueryUnescape(text)
	if err != nil {
		err = errors.New("text url code 解码" + err.Error())
		return
	}

	addLabel(newImg, deText, fontFilePath, fontSize)

	// 返回结果
	res.Text = deText
	res.ImgUrl = imgUrl
	res.Img = newImg
	res.Type = format
	return
}

// 在图片上添加文本
func addLabel(img *image.RGBA, label, fontFilePath string, fontSize int) {
	// 设置绘制位置和字体
	col := color.RGBA{R: 255, G: 255, B: 255, A: 255} // 白色

	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y

	point := fixed.Point26_6{
		X: fixed.I(maxX - 10), // 初始 X 坐标
		Y: fixed.I(maxY - 10), // 文本的 Y 坐标
	}

	// 加载字体
	fontBytes, err := os.ReadFile(fontFilePath) // 替换为你的字体文件路径
	if err != nil {
		log.Printf("failed to load font: %v", err)
		return
	}

	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Printf("failed to parse font: %v", err)
		return
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("failed to create font face: %v", err)
		return
	}

	// 创建字体绘制器
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
	}

	// 获取文本宽度并更新绘制点
	textWidth := d.MeasureString(label).Round()
	point.X = fixed.I(maxX - textWidth - 10) // 更新 X 坐标

	// 设置绘制点
	d.Dot = point

	// 绘制文本到图片上
	d.DrawString(label)
}
