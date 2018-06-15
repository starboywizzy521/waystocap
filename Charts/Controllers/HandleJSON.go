package Controllers

import (
	"net/http"
	"waystocap/code/Charts/Models"
)

func HandleJSON(w http.ResponseWriter, r *http.Request){

	w.Write(Models.SendJsonData())
}
