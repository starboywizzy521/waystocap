package Models

import (
	"database/sql"
	"time"
	"fmt"
	"waystocap/code/Public/Structures_Public"
	"waystocap/code/Public/Model_Public"
)

func Ajout (produit Structures_Public.Produits){

		if len(produit.Libelle)!=0 && len(produit.Categorie)!=0 {

			db, err := sql.Open("mysql", "root:@/test")
			Model_Public.CheckErr(err)
			defer db.Close()

			stmt, err:=db.Prepare("insert produits set libelle=?,quantite=?,prix=?,date_entree=?,categorie=?,seuil=?")
			Model_Public.CheckErr(err)

			res, err := stmt.Exec(produit.Libelle,produit.Quantite,produit.Prix,time.Now(),produit.Categorie,produit.Seuil)
			Model_Public.CheckErr(err)

			affct,err:=res.RowsAffected()
			Model_Public.CheckErr(err)

			if affct!=0 {
				fmt.Println("insertion réussie")
			}
		}
}
