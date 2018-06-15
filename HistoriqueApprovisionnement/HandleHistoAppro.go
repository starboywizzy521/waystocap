package HistoriqueApprovisionnement

import (
	"net/http"
	"html/template"
	"waystocap/code/HistoriqueApprovisionnement/Models"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Login"
)

func HandleHistoAppro(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	ListeProduits,total,Ope:=Models.HistoApr()
	PageData:=Models.Mystruct{ListeProduits,total,Ope}
	t:=template.Must(template.ParseFiles("HistoriqueApprovisionnement/Vue/HistoAppro.html","webPages/include-header.html","HistoriqueApprovisionnement/Vue/HistoApproDataTables.html", "webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,PageData)
}