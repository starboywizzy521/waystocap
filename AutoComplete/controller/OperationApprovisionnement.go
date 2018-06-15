package controller

import (
	"net/http"
	"encoding/json"
	"waystocap/code/AutoComplete/Model"
)

func OperationApprovisionnement(w http.ResponseWriter,r *http.Request){

	decoder := json.NewDecoder(r.Body)
	var t string
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	Model.ModelOperationApprovisionnement(t)
}
