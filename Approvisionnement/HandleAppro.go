package Approvisionnement

import (
	"net/http"
	"html/template"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Login"
)

func HandleAppro(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	t:=template.Must(template.ParseFiles("Approvisionnement/approvisionnement.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,nil)
}
