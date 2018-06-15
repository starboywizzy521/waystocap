// models/produits.go

package Structures_Public

import "database/sql"

type Produits struct{
	Code string
	Libelle string
	Quantite int64
	Prix float64
	Date_entree string
	Categorie string
	Seuil int64
	NomCaracteristique []sql.NullString
	ValeurCaracteristique []sql.NullString
}