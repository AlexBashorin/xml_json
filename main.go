package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseXML(xml string) string {
	thisxml := strings.NewReader(xml)
	json, err := xj.Convert(thisxml)
	if err != nil {
		panic("Ошибка парсинга xml...")
	}

	return json.String()
}

func main() {
	http.HandleFunc("/parse-xml", func(w http.ResponseWriter, r *http.Request) {
		// FORM DATA
		err := r.ParseMultipartForm(32 << 20)
		check(err)

		errPF := r.ParseForm()
		check(errPF)

		_, headerF, errF := r.FormFile("xml")
		check(errF)

		fBuff := make([]byte, headerF.Size)

		res := parseXML(string(fBuff))

		h, err := json.Marshal(res)
		check(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(h)
	})
	log.Fatal(http.ListenAndServe(":6060", nil))
}
