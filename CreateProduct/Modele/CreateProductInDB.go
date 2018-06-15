package Modele

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"fmt"
	"time"
)

func CreateProductInDB1(code string,libelle string,prix string,seuil string,category string) int64{

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	var idcategory int64
	row:=db.QueryRow("SELECT id_Categorie from categorie where Nom_Categorie='"+category+"'")
	Model_Public.CheckErr(err)
	row.Scan(&idcategory)

	stmt, err:=db.Prepare("insert into produit(Id_Produit,Code_barre,Libelle,Date_dajout,PrixUnitaire,QuantiteProduit,Seuil,id_Categorie) VALUES ('',?,?,?,?,?,?,?)")
	Model_Public.CheckErr(err)

	res, err := stmt.Exec(code,libelle,time.Now(),prix,0,seuil,idcategory)
	Model_Public.CheckErr(err)

	idprod,err:=res.LastInsertId()
	Model_Public.CheckErr(err)

	fmt.Println(idprod)

	return idprod
}

func Create2(idprod int64,nomcarac string,valeurcarac string){

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	var idcaract int64
	row:=db.QueryRow("SELECT id_caracteristique from caracteristiques where Nom_Caracteristique='"+nomcarac+"'")
	Model_Public.CheckErr(err)
	row.Scan(&idcaract)

	stmt, err:=db.Prepare("insert into produit_caracteristique(Valeur,id_caracteristique,Id_Produit) VALUES (?,?,?)")
	Model_Public.CheckErr(err)

	res, err := stmt.Exec(valeurcarac,idcaract,idprod)
	Model_Public.CheckErr(err)

	affct,err:=res.RowsAffected()
	Model_Public.CheckErr(err)

	if affct!=0 {
		fmt.Println("mise à jour réussie")
	}
}