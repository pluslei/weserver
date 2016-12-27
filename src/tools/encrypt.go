package tools

import (
	//"bufio"
	"bytes"
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"os"
	//"path"
	"strings"
)

const (
	base64key = "a59b6c15d69ef" //加密用的常量
	marksign  = "||"            //分隔符
)

func EncodeB64(message string) string {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	base64Text = bytes.TrimRight(base64Text, "\x00")
	return string(base64Text)
}

func DecodeB64(message string) string {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(base64Text, []byte(message))
	base64Text = bytes.TrimRight(base64Text, "\x00")
	return string(base64Text)
}

func Encodedata(data string) string {
	var dataresult string
	if len(data) < 1 {
		return dataresult
	}
	//skey := base64key
	skey := EncodeB64(base64key)
	data_string := EncodeB64(data)

	datalength := len(data_string)
	skeylength := len(skey)
	if skeylength > datalength {
		skeylength = datalength
	}
	for i := 0; i < skeylength; i++ {
		dataresult += string(data_string[i]) + string(skey[i])
	}

	if skeylength < datalength {
		dataresult = dataresult + data_string[skeylength:]
	}

	dataresult = strings.Replace(dataresult, "=", "O0O0O", -1) //改内容
	dataresult = strings.Replace(dataresult, "+", "o000o", -1) //改内容
	dataresult = strings.Replace(dataresult, "/", "oo00o", -1) //改内容
	return dataresult
}

func Decodedata(data string) string {
	var dataresult string
	if len(data) < 2 {
		return dataresult
	}
	//skey := base64key
	skey := EncodeB64(base64key)
	data = strings.Replace(data, "O0O0O", "=", -1) //改内容
	data = strings.Replace(data, "o000o", "+", -1) //改内容
	data = strings.Replace(data, "oo00o", "/", -1) //改内容

	datalength := len(data)
	skeylength := len(skey)
	if skeylength >= datalength/2 {
		skeylength = datalength/2 - 1
	}
	dataresult += data[:1]
	for i := 1; i < skeylength+1; i++ {
		if string(data[2*i-1]) == string(skey[i-1]) {
			dataresult += data[2*i : 2*i+1]
		}
	}

	if len(skey) == skeylength {
		dataresult = dataresult + data[2*skeylength+1:]
	}

	return DecodeB64(dataresult)
}

//公司代码用户名转换
func Transformname(codeid string, uname string, sel int) string {
	var Rescontent string
	switch sel {
	case 0:
		Rescontent = codeid + marksign + uname
	case 1:
		{
			index := strings.Index(uname, marksign) //在同一级找内容
			if -1 != index {
				Rescontent = uname[0:index]
			}
		}
	case 2:
		{
			index := strings.Index(codeid, "_") //在同一级找内容
			if -1 != index {
				Rescontent = codeid[index+1:]
			}
		}
	case 3:
		{
			index := strings.Index(codeid, "_") //在同一级找内容
			if -1 != index {
				Rescontent = codeid[0:index]
			}
		}
	default:
		{
			decodeid, _ := MainDecrypt(codeid)
			Rescontent = string(decodeid)
		}
	}
	return Rescontent
}

//单个图片操作
func Imagepro(oldpath string, newpath string, height uint, width uint) error {
	lastIndex := strings.LastIndex(oldpath, ".")
	ext := strings.ToLower(oldpath[lastIndex+1:])
	if ext == "png" {
		err := Imag_thumbpng(oldpath, 200, 100, newpath)
		return err
	} else if ext == "jpg" || ext == "jpeg" {
		err := Imag_thumbjpg(oldpath, 200, 100, newpath)
		return err
	}
	return nil
}

func Imag_thumbjpg(file string, width uint, height uint, to string) error {
	file_origin, err := os.Open(file)
	defer file_origin.Close()
	origin, _ := jpeg.Decode(file_origin)
	canvas := resize.Resize(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	if err != nil {
		beego.Error(err)
	}
	defer file_out.Close()
	err = jpeg.Encode(file_out, canvas, &jpeg.Options{100})
	return err
}

func Imag_thumbpng(file string, width uint, height uint, to string) error {
	file_origin, _ := os.Open(file)
	defer file_origin.Close()
	origin, _ := png.Decode(file_origin)
	canvas := resize.Resize(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	if err != nil {
		beego.Error(err)
	}
	defer file_out.Close()
	err = png.Encode(file_out, canvas)
	return err
}
