package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rrune/AoCZoom/util"
)

type Api struct {
	ImgPath string
}

func New(path string) Api {
	return Api{
		ImgPath: path,
	}
}

func (a Api) UpdateImage(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}

	err = ioutil.WriteFile(a.ImgPath, fileBytes, 0660)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}

// caching
func (a Api) GetImage(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile(a.ImgPath)
	if util.Check(err, true) {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := w.Write(f); err != nil {
		log.Println("Unable to write image")
	}
}
