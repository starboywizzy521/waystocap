package Modele

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
)

type caracteristique struct{
	Id_caracteristique int
	Caracteristique string
}

func GetCaracteristiqueFromDB() []caracteristique{

	var A []caracteristique

	db, err := sql.Open("mysql", "root:@/waystocap")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows,err:= db.Query("select * from caracteristiques")
	Model_Public.CheckErr(err)

	var idcaract int
	var caract string

	for rows.Next() {
		err = rows.Scan(&idcaract,&caract)
		Model_Public.CheckErr(err)
		A=append(A,caracteristique{idcaract,caract})
	}
	return A
}