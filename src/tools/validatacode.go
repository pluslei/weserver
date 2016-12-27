package tools

import (
	"github.com/freetype"
	//"strconv"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

//验证码
func PicCode(w http.ResponseWriter, req *http.Request) (string, error) {
	strCode := NewCdoe(6)
	err := DrawToImg(strCode, w)
	return strCode, err
}

//生成一个新的验证码
func NewCdoe(len int) string {
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	var strCode string
	Letternumber := []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
	for i := 0; i < len; i++ {
		n := r.Intn(62)
		//strCode += strconv.Itoa(n)
		strCode += string(Letternumber[n])
	}
	return strCode
}

func DrawToImg(strCode string, w io.Writer) error {
	arrFontFile := []string{"arial.ttf", "arialbd.ttf", "segoeuiz.ttf", "calibril.ttf", "times.ttf"}
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	fIndex := r.Intn(len(arrFontFile))
	strFontFile := "./static/fonts/" + arrFontFile[fIndex]
	fontBytes, err := ioutil.ReadFile(strFontFile)
	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}
	c := freetype.NewContext()
	var fontSize float64
	fontSize = 16
	c.SetDPI(100)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	//计算字符串的宽度，对于高度，还有此问题，懂的可以改改
	width, startY := c.MeasureString(strCode)
	heigth := c.FUnitToPixelRU(font.UnitsPerEm())
	//width += 10
	//heigth += 10
	width += 26
	heigth = 33
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, width, heigth))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	disturbBitmap(rgba)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	pt := freetype.Pt(12, 10+startY)
	_, err = c.DrawString(strCode, pt)
	if err != nil {
		return err
	}
	err = png.Encode(w, rgba)
	return nil
}

//绘制干扰背景
func disturbBitmap(img *image.RGBA) {
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	for i := 0; i < img.Rect.Max.X; i++ {
		for j := 0; j < img.Rect.Max.Y; j++ {
			n := r.Intn(100)
			if n < 40 {
				c := color.NRGBA{uint8(r.Intn(150)), uint8(r.Intn(150)), uint8(r.Intn(150)), uint8(r.Intn(100))}
				img.Set(i, j, c)
			}
		}

	}
}
