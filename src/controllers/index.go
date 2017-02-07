package controllers

import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	DealBegin(w, r)
	fmt.Fprintf(w, "Index")
	DealEnd(w, r)
}
