package handler

import (
	"fmt"
	"net/http"
	"time"
	mylog "uploader/pkg/log"
	"uploader/pkg/uploader"
)

var log = mylog.Log

//UploadFile ...
func UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Info(fmt.Sprintf("____%v", r))
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Error(fmt.Sprintf("%v, %v, %v, %v", w, r, "/", http.StatusSeeOther))
	}
	var file, handle, err = r.FormFile("file") //file key from request

	if err != nil {
		fmt.Fprintf(w, "%v", err.Error())
		log.Error(err.Error())
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	log.Info(mimeType)
	switch mimeType {
	case "zip", "rar", "tar.gz", "gz", "7z", "tar.bz2", "application/octet-stream":
		y, m, d := time.Now().Date()
		uploader.UploadFileServer(file, handle, fmt.Sprintf("%v/%v/%v/", y, int(m), d))
		jsonResponse(w, http.StatusCreated, "Succesfully uploaded !!!\n")
	default:
		jsonResponse(w, http.StatusBadRequest, "Not valid format file !!!")
	}
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, message)
}
