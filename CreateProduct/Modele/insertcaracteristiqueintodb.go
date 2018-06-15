package Modele

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"strings"
	"fmt"
)

func InsertCaracteristiqueIntoDB(id_category int64,t string){


	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	s:=strings.Trim(t,"/")
	result:=strings.Split(s,"/")

	for j:=range result{

		row:=db.QueryRow("select id_caracteristique from caracteristiques WHERE Nom_Caracteristique='"+result[j]+"'")
		var idcara int64
		Model_Public.CheckErr(err)
		row.Scan(&idcara)

		fmt.Println(idcara)
		if idcara==0 {

			stmt, err:=db.Prepare("insert into caracteristiques VALUES ('',?)")
			Model_Public.CheckErr(err)

			res, err := stmt.Exec(result[j])
			Model_Public.CheckErr(err)

			idcara,err=res.LastInsertId()
			Model_Public.CheckErr(err)
		}

		fmt.Println(idcara)
		stmt, err:=db.Prepare("insert into categ_carct(id_Categorie,id_caracteristique) VALUES (?,?)")
		Model_Public.CheckErr(err)

		res, err := stmt.Exec(id_category,idcara)
		Model_Public.CheckErr(err)

		affct,err:=res.RowsAffected()
		Model_Public.CheckErr(err)

		if affct!=0 {
			fmt.Println("mise à jour réussie")
		}

	}

}
