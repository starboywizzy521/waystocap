package Controller

import (
	"net/http"
	"html/template"
	"strconv"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/ModifierProduit"
	"waystocap/code/Public/Structures_Public"
	"strings"
	"fmt"
)

func HandleMisAjour(w http.ResponseWriter,r *http.Request){

	r.ParseForm()
	if r.Method=="GET" {
		tmpl:= template.Must(template.ParseFiles("ModifierProduit/ModifierProduit.html","webPages/include-header.html","webPages/include-body.html","webPages/include-footer.html"))
		tmpl.Execute(w,prod)
	} else{

		code:=r.FormValue("Code")
		libelle:=r.Form.Get("Libelle")
		qte,err:=strconv.ParseInt(r.Form.Get("Quantite"),10,0)
		Model_Public.CheckErr(err)
		prix,err:=strconv.ParseFloat(r.Form.Get("Prix"),64)
		Model_Public.CheckErr(err)
		categorie:=r.Form.Get("Categorie")
		seuil,err:=strconv.ParseInt(r.Form.Get("Seuil"),10,0)
		Model_Public.CheckErr(err)
		date:=r.Form.Get("DateEntre")

		ModifierProduit.MisAjourProduit(Structures_Public.Produits{code,libelle,qte,prix,date,categorie,seuil,nil,nil},idprod)
		fmt.Println("bonjkou")
		NomCaracteristique:=r.FormValue("NomCaracteristique")
		result:=strings.Split(strings.Trim(NomCaracteristique,"[]"),",")

		valeurcaracteristique:=r.FormValue("valeurcaracteristique")
		resultvaleurcara:=strings.Split(strings.Trim(valeurcaracteristique,"[]"),",")

		for n:= range result{
			ModifierProduit.MisAjourCaracteristiqueProduit(strings.Trim(resultvaleurcara[n],"\""),strings.Trim(result[n],"\""),idprod)
		}

		w.Write(Model_Public.Sendsuccessalert("modification réussie"))

	}
}
