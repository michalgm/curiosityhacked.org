package _go

import "net/http"

func init() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  })
}
