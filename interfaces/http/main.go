package main

import (
	"log"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	log.Println(string(bs))
	log.Println("Bytes written:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	log.Println(string(bs))

	Similar can be achieved using io.Copy
	io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
