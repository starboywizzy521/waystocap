package controllerSortieArticle

import (
"net/http"
"html/template"
"strconv"
"waystocap/code/Public/Model_Public"
"waystocap/code/SortieArticle/Models"
"waystocap/code/Public/Structures_Public"
)

func HandleSortieArticle(w http.ResponseWriter,r *http.Request){

	r.ParseForm()

	if r.Method=="GET"{

		t, _ := template.ParseFiles("SortieArticle/sortiearticle.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html")
		t.Execute(w,Models.AutocompleteSortie())
	}

	if r.Method=="POST"{
		if r.Method=="POST" {

			libelle:=r.Form.Get("Libelle")
			qte,err:=strconv.ParseInt(r.Form.Get("Quantite"),10,0)
			Model_Public.CheckErr(err)
			prix,err:=strconv.ParseFloat(r.Form.Get("Prix"),64)
			Model_Public.CheckErr(err)
			categorie:=r.Form.Get("Categorie")
			seuil,err:=strconv.ParseInt(r.Form.Get("Seuil"),10,0)
			Model_Public.CheckErr(err)

			Models.Ajout(Structures_Public.Produits{"",libelle,qte,prix,"",categorie,seuil,nil,nil})
		}
	}

}