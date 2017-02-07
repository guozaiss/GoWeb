package controllers

import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Deal(w, r)
	fmt.Fprintf(w, "Index")
}
