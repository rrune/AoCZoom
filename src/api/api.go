package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rrune/AoCZoom/util"
)

type Api struct {
	ImgPath string
	Image   []byte
}

func New(path string) Api {
	a := Api{
		ImgPath: path,
	}
	a.loadImage()
	return a
}

func (a *Api) loadImage() {
	f, err := ioutil.ReadFile(a.ImgPath)
	if util.Check(err, true) {
		log.Println(err)
		return
	}
	a.Image = f
}

func (a *Api) UpdateImage(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(a.ImgPath, fileBytes, 0666)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	a.Image = fileBytes
	w.WriteHeader(http.StatusOK)
}

func (a *Api) GetImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := w.Write(a.Image); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
