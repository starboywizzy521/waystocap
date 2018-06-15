package Sortie

import (
	"net/http"
	"html/template"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Login"
)

func HandleSortie(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	t:=template.Must(template.ParseFiles("Sortie/sortie.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,nil)
}
