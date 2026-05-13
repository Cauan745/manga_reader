package main

import (
	"fmt"
	"net/http"
)

func (a *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
}
