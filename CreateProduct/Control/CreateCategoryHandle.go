package Control

import (
	"net/http"
	"html/template"
	"waystocap/code/CreateProduct/Modele"
)

func CreateCategoryHandle(w http.ResponseWriter,r *http.Request){

	r.ParseForm()

	if r.Method=="GET"{
		t, _ := template.ParseFiles("CreateProduct/Vue/createcategory.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html")
		t.Execute(w,Modele.GetCaracteristiqueFromDB())
	} else {
		Modele.InsertCaracteristiqueIntoDB(Modele.InsertCategorieIntoDB(r.FormValue("category")),r.FormValue("caracteristique"))
	}
}
