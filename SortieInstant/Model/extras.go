package Model


type ProduitSortie struct {
	Code int64
	Libelle string
	Quantite int64
	Prix float64
	Date_entree string
	Categorie string
	Benef string
	IdOP int
}


type Extras struct{
	P [] ProduitSortie
}



