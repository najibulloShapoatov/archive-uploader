package uploader

import (
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

var err error

//UploadFileServer ...
func UploadFileServer(file multipart.File, handle *multipart.FileHeader, path string) string {

	var data, err = ioutil.ReadAll(file)
	if err != nil {
		return "not readble data !!!"
	}

	//var name = strings.Split(handle.Filename, ".")
	//var ext = name[len(name)-1]
	//var fileName = uniqidString() + "." + e.Get()
	var Path = "./public/uploads/" + path
	_, err = os.Stat(Path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(Path, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
	fileName := uniqidString() + handle.Filename
	//fileName := handle.Filename

	err = ioutil.WriteFile(Path+fileName, data, 0666)

	if err != nil {
		return "not saved from folder!!!"
	}

	return fileName
}

func uniqidString() string {
	var t1 = time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	var s = b.String()
	return t1 + "_" + s + "__"
}
