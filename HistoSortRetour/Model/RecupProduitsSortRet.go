package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"fmt"
)

func RecupProduitsSortRet() [] ProduitSortie{

	var P [] ProduitSortie=nil
	db, err := sql.Open("mysql", "root:@/waystocap")
	defer db.Close()
	Model_Public.CheckErr(err)

	rows1,err:=db.Query("SELECT distinct id_operation from operation where Type_operation='Retour' ORDER BY id_operation DESC")
	Model_Public.CheckErr(err)

	for(rows1.Next()) {

		var idOP int
		rows1.Scan(&idOP)
		fmt.Println(idOP)
		rows, err := db.Query("SELECT DISTINCT p.Code_barre,p.Id_Produit, p.Libelle,po.Quantite, p.PrixUnitaire,"+
			"(select `Date` from operation where id_operation="+
			"(select substr(Description,1,instr(Description,'*|*')-1) from produit_operation where id_operation=?)),"+
			"(select `Date` from operation where id_operation=po.id_operation),"+
			"(select Nom_Categorie from Categorie where id_Categorie=p.id_Categorie),"+
			"( select concat(Nom,' ',Prenom) from employee where id_employee=(select id_employee from operation where id_operation=?)),"+
			"substr(po.Description,instr(Description,'*|*')+3,999999)"+
			"FROM produit p, produit_operation po where po.id_operation=? and p.Id_Produit in (select Id_Produit from produit_operation where id_operation=po.id_operation)",idOP,idOP,idOP)
		Model_Public.CheckErr(err)

		var codebarre string
		var code1 int64
		var libelle1 string
		var quantite1 int64
		var prix1 float64
		var date_sortie1 string
		var date_retour1 string
		var categorie1 string
		var benef string
		var description string

		for rows.Next() {
			err = rows.Scan(&codebarre,&code1, &libelle1, &quantite1, &prix1, &date_sortie1,&date_retour1, &categorie1, &benef, &description)
			Model_Public.CheckErr(err)
			P = append(P, ProduitSortie{codebarre, libelle1, quantite1, prix1, date_sortie1,date_retour1, categorie1, benef,description,idOP})
		}
	}

	return P
}