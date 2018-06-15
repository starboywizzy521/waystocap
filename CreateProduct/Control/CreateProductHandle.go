package Control

import (
	"net/http"
	"html/template"
	"waystocap/code/CreateProduct/Modele"
	"fmt"
	"strings"
)

func CreateProductControl(w http.ResponseWriter,r *http.Request){

	r.ParseForm()

	if r.Method=="GET"{
		t, _ := template.ParseFiles("CreateProduct/Vue/createproduct.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html")
		t.Execute(w,Modele.GetCategoryFromDB())
	} else {


		ddf:=Modele.CreateProductInDB1(r.FormValue("codebarre"),r.FormValue("libelle"),r.FormValue("prixunitaire"),r.FormValue("seuil"),r.FormValue("category"))

		bjhiu:=r.FormValue("libellecarac")
		result:=strings.Split(strings.Trim(bjhiu,"[]"),",")
			for n:= range result{
				carac:=strings.Trim(result[n],"\"")
				fmt.Println(r.FormValue("valeurcaracteristique["+carac+"]"))
				Modele.Create2(ddf,carac,r.FormValue("valeurcaracteristique["+carac+"]"))

			}
	}
}