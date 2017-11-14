package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	d := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."

	rdr := strings.NewReader(d)
	io.Copy(os.Stdout, rdr)

	buf := bytes.NewBuffer([]byte(d))
	io.Copy(os.Stdout, buf)

	res, _ := http.Get("http://www.mcleods.com")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
