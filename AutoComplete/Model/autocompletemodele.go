package Model

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
)

type libelletab struct{
	Code string
	Libelle string
	Quantite int
	Prix float64
}

func Autocomplete() []libelletab{

	var A []libelletab

	db, err := sql.Open("mysql", "ql7243021:nEFCqrWjib@tcp(sql7.freemysqlhosting.net)/sql7243021?charset=utf8")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows,err:= db.Query("SELECT Id_Produit,Libelle,QuantiteProduit,PrixUnitaire FROM produit")
	Model_Public.CheckErr(err)

	var libelle string
	var code string
	var quantite int
	var prix float64


	for rows.Next() {
		err = rows.Scan(&code,&libelle,&quantite,&prix)
		Model_Public.CheckErr(err)
		A=append(A,libelletab{code,libelle,quantite,prix})
	}
	return A
}

func Entry(){

	db, err := sql.Open("mysql", "root:@/waystocap")
	defer db.Close()
	Model_Public.CheckErr(err)

}
