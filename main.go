package main

import (
	"bytes"
	"io"
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
		errPF := r.ParseForm()
		check(errPF)

		file, _, errF := r.FormFile("xml")
		if errF != nil {
			log.Fatal(errF)
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, file)
		check(err)
		bbuf := buf.Bytes()

		res := parseXML(string(bbuf))

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(res))
	})
	log.Fatal(http.ListenAndServe(":4040", nil))
}
