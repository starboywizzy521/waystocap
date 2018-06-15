package Models

import (
	"database/sql"
	"encoding/json"
	"waystocap/code/Public/Model_Public"
)

type Data struct{
	Mois string  `json:"mois"`
	Prix float32 `json:"prix"`
}

type EphrauJson struct{
Mois string  `json:"moi"`
TotalE float32 `json:"totalE"`
TotalS float32 `json:"totalS"`
}


func SendJsonData() []byte {
	var data []Data
	db, err:= sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	rows,err:=db.Query("SELECT Nom_Categorie,SUM(QuantiteProduit) FROM `produit` NATURAL JOIN categorie GROUP by id_Categorie")
	Model_Public.CheckErr(err)
	var mois string
	var prix float32
	for rows.Next() {
		_=rows.Scan(&mois,&prix)
		data=append(data,Data{mois,prix})
	}
	db.Close()
	data1,err:=json.Marshal(data)
	data=nil
	Model_Public.CheckErr(err)
	return data1
}

func SendJsonEphrau() []byte{
	var data []EphrauJson
	db, err:= sql.Open("mysql", "root:@/waystocap")
	Model_Public.CheckErr(err)
	rows0,err:=db.Query("SELECT distinct MONTHNAME( `Date` ) AS `Mois`FROM operation where YEAR(`Date`)=YEAR(DATE(NOW())) ORDER BY MONTH(`Date`) ASC ")
	Model_Public.CheckErr(err)
	var Mois string
	var TotalEntree float32
	var TotalSortie float32
	for rows0.Next() {
		_=rows0.Scan(&Mois)
		errE:=db.QueryRow("SELECT sum(po.Quantite * (select PrixUnitaire from produit where Id_Produit= po.Id_Produit)) " +
			"FROM produit_operation po " +
			"WHERE po.id_operation IN ( select id_operation from operation where MONTHNAME(`Date`)=? and YEAR(`Date`)=YEAR(DATE(NOW())) AND Type_operation='entry')",Mois).Scan(&TotalEntree)

		errS:=db.QueryRow("SELECT sum(po.Quantite * (select PrixUnitaire from produit where Id_Produit= po.Id_Produit)) " +
			"FROM produit_operation po " +
			"WHERE po.id_operation IN ( select id_operation from operation where MONTHNAME(`Date`)=? and YEAR(`Date`)=YEAR(DATE(NOW())) AND Type_operation='output')",Mois).Scan(&TotalSortie)


		if errE != nil {
			TotalEntree=0
		}
		if errS != nil {
			TotalSortie=0
		}
		data=append(data,EphrauJson{Mois,TotalEntree,TotalSortie})
	}
	db.Close()
	data1,err:=json.Marshal(data)
	data=nil
	Model_Public.CheckErr(err)
	return data1
}