package Model


type ProduitSortie struct {
	Code string
	Libelle string
	Quantite int64
	Prix float64
	Date_sortie string
	Date_retour string
	Categorie string
	Benef string
	Description string
	IdOP int
}


type Extras struct{
	P [] ProduitSortie
}



