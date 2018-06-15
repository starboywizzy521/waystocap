package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"waystocap/code/Public/Structures_Public"
	"waystocap/code/Public/Model_Public"
)

func HistoApr() ([] Structures_Public.Produits,[]float32, []int){

	var Total []float32=nil
	var P [] Structures_Public.Produits=nil
	var Ope []int

	db, err := sql.Open("mysql", "root:@/waystocap")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows1,err:=db.Query("SELECT distinct id_operation from operation where Type_operation='entry' ORDER BY id_operation DESC")
	Model_Public.CheckErr(err)

	for rows1.Next(){

		var idOP int
		rows1.Scan(&idOP)
		rows,err:= db.Query("SELECT DISTINCT p.Code_barre,p.Id_Produit,p.Libelle,(select Quantite from produit_operation where id_operation=? and Id_Produit=p.Id_Produit),p.PrixUnitaire,(select `Date` from operation where id_operation=?),(select Nom_Categorie from Categorie where id_Categorie=p.id_Categorie),p.Seuil FROM produit p where p.Id_Produit in (select Id_Produit from produit_operation where id_operation=?)",idOP,idOP,idOP)
		Model_Public.CheckErr(err)
		var codebarre string
		var code1 string
		var libelle1 string
		var quantite1 int64
		var prix1 float64
		var date_entree1 string
		var categorie1 string
		var seuil int64

		for rows.Next() {
			err = rows.Scan(&codebarre,&code1,&libelle1,&quantite1,&prix1,&date_entree1,&categorie1,&seuil)
			Model_Public.CheckErr(err)
			P=append(P,Structures_Public.Produits{codebarre,libelle1,quantite1,prix1,date_entree1,categorie1,seuil,nil,nil})
			Ope=append(Ope,idOP)
		}
		var totaux float32
		db.QueryRow("SELECT sum(po.Quantite * (select PrixUnitaire from produit where Id_Produit= po.Id_Produit)) FROM produit_operation po WHERE po.id_operation=?",idOP).Scan(&totaux)
		Total=append(Total,totaux)
	}

	return P,Total,Ope
}