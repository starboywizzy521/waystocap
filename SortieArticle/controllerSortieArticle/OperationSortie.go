package controllerSortieArticle

import (
	"net/http"
	"waystocap/code/SortieArticle/Models"
)

func OperationSortie(w http.ResponseWriter,r *http.Request){

	benef:= r.FormValue("benef")
	t:= r.FormValue("tableau")
	Models.ModelOperationSortie(t,benef)
}
