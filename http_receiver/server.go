package http_receiver

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func StartHTTPReceiver(port int, responseText string) {
	handler := HTTPReceiver{responseText: responseText}
	err := http.ListenAndServe(":"+strconv.Itoa(port), handler)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type HTTPReceiver struct {
	responseText string
}

func (hr HTTPReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r)
	w.WriteHeader(400)
	w.Write(bytes.NewBufferString(hr.responseText).Bytes())
}
