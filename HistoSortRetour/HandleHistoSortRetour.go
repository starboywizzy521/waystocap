package HistoSortRetour

import (
	"net/http"
	"html/template"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Login"
	"waystocap/code/HistoSortRetour/Model"
)

func HandleSortRetour(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	PageData:=Model.Extras{Model.RecupProduitsSortRet()}
	t:=template.Must(template.ParseFiles("HistoSortRetour/Vue/HistoSortRetour.html","HistoSortRetour/Vue/gg.html","webPages/include-header.html","webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,PageData)
}

