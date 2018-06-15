package Modele

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
)

type category struct{
	Id_category string
	Category string
	Id_caracteristique sql.NullString
	Caracteristique sql.NullString
}

func GetCategoryFromDB() []category{

		var A []category

		db, err := sql.Open("mysql", "root:@/waystocap")
		defer db.Close()
		Model_Public.CheckErr(err)

		rows,err:= db.Query("select categorie.id_Categorie,categorie.Nom_Categorie,caracteristiques.id_caracteristique,caracteristiques.Nom_Caracteristique from ((categorie left join categ_carct on categorie.id_Categorie=categ_carct.id_Categorie) LEFT join caracteristiques on caracteristiques.id_caracteristique=categ_carct.id_caracteristique)")
		Model_Public.CheckErr(err)

		var idcateg string
		var categ string
		var idcaract sql.NullString
		var caract sql.NullString

		for rows.Next() {
		err = rows.Scan(&idcateg,&categ,&idcaract,&caract)
		Model_Public.CheckErr(err)
		A=append(A,category{idcateg,categ,idcaract,caract})
	}
		return A
}

