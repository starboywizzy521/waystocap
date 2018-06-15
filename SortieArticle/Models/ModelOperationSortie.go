package Models

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"strings"
	"fmt"
)

func ModelOperationSortie(t string,benef string){

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()
	stmt, err:=db.Prepare("insert INTO operation VALUES ('',NOW(),'output',NULL)")
	Model_Public.CheckErr(err)

	res, err := stmt.Exec()
	Model_Public.CheckErr(err)

	idop,err:=res.LastInsertId()
	Model_Public.CheckErr(err)

	s:=strings.Trim(t,"|")
	result:=strings.Split(s,"|")

	emp:= strings.Split(benef," ")
	for j:=range result{

		cells:=strings.Split(result[j]," ")
		stmt, err:=db.Prepare("call operation_sortie(?,?,?,?,?,?)")
		Model_Public.CheckErr(err)

		res, err := stmt.Exec(cells[1],cells[2],cells[0],emp[0],emp[1],idop)
		Model_Public.CheckErr(err)

		affct,err:=res.RowsAffected()
		Model_Public.CheckErr(err)

		if affct!=0 {
			fmt.Println("mise à jour réussie")
		}

	}
}
