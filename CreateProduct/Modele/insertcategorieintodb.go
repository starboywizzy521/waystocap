package Modele

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
)

func InsertCategorieIntoDB(category string) int64{

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	stmt, err:=db.Prepare("insert INTO categorie VALUES ('',?)")
	Model_Public.CheckErr(err)

	res, err := stmt.Exec(category)
	Model_Public.CheckErr(err)

	categoryID,err:=res.LastInsertId()
	Model_Public.CheckErr(err)

	return categoryID
}
