package Model

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"strings"
	"fmt"
)

func ModelOperationApprovisionnement(t string){

	db, err := sql.Open("mysql", "ql7243021:nEFCqrWjib@tcp(sql7.freemysqlhosting.net)/sql7243021?charset=utf8")
	Model_Public.CheckErr(err)
	defer db.Close()

	stmt, err:=db.Prepare("insert INTO operation VALUES ('',NOW(),'entry',NULL)")
	Model_Public.CheckErr(err)

	res, err := stmt.Exec()
	Model_Public.CheckErr(err)

	idop,err:=res.LastInsertId()
	Model_Public.CheckErr(err)

	s:=strings.Trim(t,"/")
	result:=strings.Split(s,"/")

	for j:=range result{
		cells:=strings.Split(result[j]," ")

		stmt, err:=db.Prepare("call operation_entre(?,?,?,?,?)")
		Model_Public.CheckErr(err)

		res, err := stmt.Exec(cells[1],cells[2],cells[3],cells[0],idop)
		Model_Public.CheckErr(err)

		affct,err:=res.RowsAffected()
		Model_Public.CheckErr(err)

		if affct!=0 {
			fmt.Println("mise à jour réussie")
		}

	}
}
