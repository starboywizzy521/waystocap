package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"fmt"
	"strings"
	"strconv"
)

func RecupProduitsSortie() [] ProduitSortie{

	var P [] ProduitSortie=nil
	db, err := sql.Open("mysql", "root:@/waystocap")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows1,err:=db.Query("SELECT distinct id_operation from operation where Type_operation='output' and id_operation in (select id_operation from produit_operation where description='1') ORDER BY id_operation DESC")
	Model_Public.CheckErr(err)

	for(rows1.Next()) {

		var idOP int
		rows1.Scan(&idOP)
		fmt.Println(idOP)
		rows, err := db.Query("SELECT DISTINCT p.Id_Produit,"+
			"p.Libelle,"+
			"(select Quantite from produit_operation where Id_Produit=p.Id_Produit and id_operation=?),"+
			"p.PrixUnitaire,"+
				"(select `Date` from operation where id_operation=?),"+
				"(select Nom_Categorie from Categorie where id_Categorie=p.id_Categorie),"+
					"( select concat(Nom,' ',Prenom) from employee where id_employee=(select id_employee from operation where id_operation=?))"+
					"	FROM produit p"+
					"	where p.Id_Produit in (select Id_Produit from produit_operation where id_operation=? and description='1')",idOP,idOP,idOP,idOP)
		Model_Public.CheckErr(err)

		var code1 int64
		var libelle1 string
		var quantite1 int64
		var prix1 float64
		var date_entree1 string
		var categorie1 string
		var benef string

		for rows.Next() {
			err = rows.Scan(&code1, &libelle1, &quantite1, &prix1, &date_entree1, &categorie1, &benef)
			Model_Public.CheckErr(err)
			P = append(P, ProduitSortie{code1, libelle1, quantite1, prix1, date_entree1, categorie1, benef,idOP})
		}
	}

	return P
}

func OperatRetour(idOP int, code int64, benef string, qte int64, description string){

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	stmt0, err:=db.Prepare("update produit_operation set description='2' where id_operation=? AND Id_Produit=?")
	Model_Public.CheckErr(err)

	res0,err := stmt0.Exec(idOP,code)
	Model_Public.CheckErr(err)

	affct0,err:=res0.RowsAffected()
	Model_Public.CheckErr(err)
	if affct0!=0 {
		fmt.Println("sortie marqué comme revenue")
	}


	stmt1, err:=db.Prepare("insert INTO operation VALUES ('',NOW(),'Retour',NULL)")
	Model_Public.CheckErr(err)

	res1, err := stmt1.Exec()
	Model_Public.CheckErr(err)

	DerIdOP,err:=res1.LastInsertId()
	Model_Public.CheckErr(err)

	affct1,err:=res1.RowsAffected()
	Model_Public.CheckErr(err)
	if affct1!=0 {
		fmt.Println("Insertion d'operation retour")
	}
	var OperatSortieDescrip =""+strconv.Itoa(idOP)+"*|*"+description
	emp:= strings.Split(benef," ")

		stmt2, err:=db.Prepare("call operation_retour(?,?,?,?,?,?)")
		Model_Public.CheckErr(err)

		res2, err := stmt2.Exec(qte,code,emp[0],emp[1],DerIdOP,OperatSortieDescrip)
		Model_Public.CheckErr(err)

		affct2,err:=res2.RowsAffected()
		Model_Public.CheckErr(err)

		if affct2!=0 {
			fmt.Println("mise à jour réussie")
		}
}