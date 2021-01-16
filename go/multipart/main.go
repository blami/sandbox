package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

const Index = `<!DOCTYPE html>
<html>
  <head>
      <meta charset="UTF-8" />
  </head>
  <body>
    <div>
        <form method="POST" enctype="multipart/form-data" action="/">
            <input name="img-1" type="file" />
            <input type="submit" value="submit" />
        </form>
    </div>
  </body>
</html>`

func main() {
	http.HandleFunc("/", handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, Index)

	case "POST":
		file, header, err := r.FormFile("img-1")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		b := bytes.NewBuffer(nil)
		if _, err := io.Copy(b, file); err != nil {
			log.Fatal(err)
		}
		// convert bytes to base64
		fmt.Println(header.Filename, base64.URLEncoding.EncodeToString(b.Bytes()))
	}
}
