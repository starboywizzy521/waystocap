package Controllers

import (
	"net/http"
	"waystocap/code/Charts/Models"
)

func HandleEphrauJSON(w http.ResponseWriter, r *http.Request){

	w.Write(Models.SendJsonEphrau())
}