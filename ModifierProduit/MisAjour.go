package ModifierProduit

import (
	"database/sql"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Public/Structures_Public"
)

func MisAjourProduit (produit Structures_Public.Produits,idprod int64) {

			db, err := sql.Open("mysql", "root:@/waystocap")
			Model_Public.CheckErr(err)

			defer db.Close()

			stmt, err:=db.Prepare("update produit set Code_barre=?,Libelle=?,PrixUnitaire=?,Date_dajout=?,Seuil=? where Id_Produit=?")
			Model_Public.CheckErr(err)

			res, err := stmt.Exec(produit.Code,produit.Libelle,produit.Prix,produit.Date_entree,produit.Seuil,idprod)
			Model_Public.CheckErr(err)

			affct,err:=res.RowsAffected()
			Model_Public.CheckErr(err)


			if affct==0 {
				//return 0,err
			}
}

func MisAjourCaracteristiqueProduit(valeurcarac string,nomcarac string,idprod int64){

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	var idcaract int64
	row:=db.QueryRow("SELECT id_caracteristique from caracteristiques where Nom_Caracteristique=?",nomcarac)
	Model_Public.CheckErr(err)
	row.Scan(&idcaract)

	stmt, err:=db.Prepare("UPDATE produit_caracteristique set Valeur=? where id_caracteristique=? AND Id_Produit=?")
	Model_Public.CheckErr(err)

	res, err := stmt.Exec(valeurcarac,idcaract,idprod)
	Model_Public.CheckErr(err)

	affct,err:=res.RowsAffected()
	Model_Public.CheckErr(err)

	if affct!=0 {
	}else {
		if valeurcarac!=""{
			stmt, err=db.Prepare("insert INTO produit_caracteristique(Valeur,id_caracteristique,Id_Produit) VALUES (?,?,?) on DUPLICATE KEY  UPDATE Valeur=?")
			Model_Public.CheckErr(err)

			res, err = stmt.Exec(valeurcarac,idcaract,idprod,valeurcarac)
			Model_Public.CheckErr(err)
		}
	}

}

func DetailsProduit(code string,libelle string,qte int64,prix float64,categorie string,seuil int64,date string ) (idpro int64,infopro Structures_Public.Produits){

	var InfoProduit Structures_Public.Produits

	db, err := sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	defer db.Close()

	var idprod int64
	row:=db.QueryRow("select Id_Produit from produit where Code_barre=? and Libelle=? and Date_dajout=? and PrixUnitaire=? and QuantiteProduit=? and Seuil=?",code,libelle,date,prix,qte,seuil)
	Model_Public.CheckErr(err)
	row.Scan(&idprod)

	var idcat int64
	row=db.QueryRow("select id_Categorie from categorie where Nom_Categorie=?",categorie)
	Model_Public.CheckErr(err)
	row.Scan(&idcat)

	rows,err:=db.Query("select Nom_Caracteristique FROM caracteristiques NATURAL JOIN  categ_carct WHERE id_Categorie=?",idcat)
	Model_Public.CheckErr(err)

	var nomcarac sql.NullString
	var toutecarac []sql.NullString

	for rows.Next(){
		rows.Scan(&nomcarac)
		toutecarac=append(toutecarac,nomcarac)

	}

	rows,err=db.Query("select Valeur from produit_caracteristique natural join caracteristiques where produit_caracteristique.Id_Produit=?",idprod)
	Model_Public.CheckErr(err)

	var valeurcarac sql.NullString
	var toutevaleurcarac []sql.NullString

	for rows.Next(){
		rows.Scan(&valeurcarac)
		toutevaleurcarac=append(toutevaleurcarac,valeurcarac)
	}

	InfoProduit=Structures_Public.Produits{code,libelle,qte,prix,date,categorie,seuil,toutecarac,toutevaleurcarac}

	return idprod,InfoProduit

}