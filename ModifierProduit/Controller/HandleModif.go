package Controller

import (
	"net/http"
	"html/template"
)

func HandleModif(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method=="GET"{
		tmpl:= template.Must(template.ParseFiles("ModifierProduit/DétailsProduit.html","webPages/include-header.html","webPages/include-body.html","webPages/include-footer.html"))
		tmpl.Execute(w,prod)
	}
}
