package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	zpl := []byte("^xa^cfa,50^fo100,100^fdHello World^fs^xz")
	// adjust print density (8dpmm), label width (4 inches), label height (6 inches), and label index (0) as necessary
	req, err := http.NewRequest("POST", "http://api.labelary.com/v1/printers/8dpmm/labels/4x6/0/", bytes.NewBuffer(zpl))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Accept", "image/png") // omit this line to get PNG images back

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		file, err := os.Create("label.png") // change file name for PNG images
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		io.Copy(file, response.Body)
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Fatalln(string(body))
	}
}
