package main

import (
    "log"
)

func main() {
    log.Println("пошли за пивом")
    log.Println("i am awesome! no! very awesome!")
    http.HandleFunc("/", index)
    http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("за мамку и двор стреляю в упор"))
}

func usefulFunction(value string) string {
  return string + string
}
