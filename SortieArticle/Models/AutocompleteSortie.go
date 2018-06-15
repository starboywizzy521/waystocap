package Models

import (
"database/sql"
"waystocap/code/Public/Model_Public"
)

type libelletab struct{
	Code string
	CodeBR string
	Libelle string
	Quantite int
	Categorie string
}

type Employee struct{
	id_emp int
	Nom string
	Prenom string
}

type libelletabSortie struct{
	Produit []libelletab
	Emp []Employee
}

func AutocompleteSortie() libelletabSortie{

	var A []libelletab
	var E []Employee
	db, err := sql.Open("mysql", "root:@/waystocap")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows,err:= db.Query("SELECT p.Id_Produit,p.Code_barre,p.Libelle,p.QuantiteProduit,(SELECT Nom_Categorie From categorie where id_categorie = p.id_Categorie) FROM produit p")
	Model_Public.CheckErr(err)

	var libelle string
	var code string
	var codeBR string
	var quantite int
	var Categ string


	for rows.Next() {
		err = rows.Scan(&code,&codeBR,&libelle,&quantite,&Categ)
		Model_Public.CheckErr(err)
		A=append(A,libelletab{code,codeBR,libelle,quantite,Categ})
	}

	rows2,err:= db.Query("SELECT id_employee,Nom,Prenom FROM employee")
	Model_Public.CheckErr(err)

	var Id int
	var Nom string
	var Prenom string

	for rows2.Next() {
		err = rows2.Scan(&Id,&Nom,&Prenom)
		Model_Public.CheckErr(err)
		E=append(E,Employee{Id,Nom,Prenom})
	}

	var Data libelletabSortie

	Data.Emp=E
	Data.Produit=A

	return Data

}
