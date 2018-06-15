package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"waystocap/code/Public/Structures_Public"
	"waystocap/code/Public/Model_Public"
)

func RecupProduits() [] Structures_Public.Produits{

	var P [] Structures_Public.Produits=nil
	db, err := sql.Open("mysql", "sql7243021:nEFCqrWjib@tcp(sql7.freemysqlhosting.net)/sql7243021?charset=utf8")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows,err:= db.Query("SELECT Code_barre,Libelle,QuantiteProduit,PrixUnitaire,Date_dajout,Nom_Categorie,Seuil FROM produit left join categorie on produit.id_Categorie=categorie.id_Categorie")
	Model_Public.CheckErr(err)

	var code1 string
	var libelle1 string
	var quantite1 int64
	var prix1 float64
	var date_entree1 string
	var categorie1 string
	var seuil1 int64

	for rows.Next() {
		err = rows.Scan(&code1,&libelle1,&quantite1,&prix1,&date_entree1,&categorie1,&seuil1)
		Model_Public.CheckErr(err)
		P=append(P,Structures_Public.Produits{code1,libelle1,quantite1,prix1,date_entree1,categorie1,seuil1,nil,nil})
	}
	return P
}