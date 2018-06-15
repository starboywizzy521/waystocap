package Controller

import (
	"net/http"
	"strconv"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Public/Structures_Public"
	"waystocap/code/ModifierProduit"
)

var prod Structures_Public.Produits
var idprod int64

func ReceiveAjax(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.Method == "POST"{

		code:=r.FormValue("code")
		libelle:=r.FormValue("libelle")
		qte,err:=strconv.ParseInt(r.FormValue("quantite"),10,0)
		Model_Public.CheckErr(err)
		prix,err:=strconv.ParseFloat(r.FormValue("prix"),64)
		Model_Public.CheckErr(err)
		categorie:=r.FormValue("categorie")
		seuil,err:=strconv.ParseInt(r.FormValue("seuil"),10,0)
		Model_Public.CheckErr(err)
		date:=r.FormValue("date")

		idprod,prod=ModifierProduit.DetailsProduit(code,libelle,qte,prix,categorie,seuil,date)
	}
}
