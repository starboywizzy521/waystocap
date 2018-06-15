package main

import (
	"net/http"
	"log"
	"waystocap/code/Charts/Controllers"
	"waystocap/code/ModifierProduit/Controller"
	"waystocap/code/Login"
	"waystocap/code/Tables"
	"waystocap/code/Dashboard"
	"waystocap/code/Approvisionnement"
	"waystocap/code/HistoriqueApprovisionnement"
	"waystocap/code/AutoComplete/controller"
	"waystocap/code/CreateProduct/Control"
	"waystocap/code/Sortie"
	"waystocap/code/SortieInstant/ControllerSortieInst"
	"waystocap/code/SortieArticle/controllerSortieArticle"
	"waystocap/code/HistoSortRetour"
	"waystocap/code/Historique"
)



func main() {

	http.HandleFunc("/createcategory",Control.CreateCategoryHandle)

	http.HandleFunc("/createproduct",Control.CreateProductControl)

	http.HandleFunc("/autocomplete",controller.HandleAutoComplete)

	http.HandleFunc("/OperationApprovisionnement",controller.OperationApprovisionnement)

	http.HandleFunc("/ephraujson",Controllers.HandleEphrauJSON)

	http.HandleFunc("/json",Controllers.HandleJSON)

	http.HandleFunc("/misajour", Controller.HandleMisAjour)

	http.HandleFunc("/Détailsduproduit", Controller.HandleModif)

	http.HandleFunc("/receive", Controller.ReceiveAjax)

	http.HandleFunc("/login", Login.HandleLogin)

	http.HandleFunc("/ChangeLogin", Login.HandleChangeLogin)

	http.HandleFunc("/tables", Tables.HandleTables)

	http.HandleFunc("/charts", Controllers.HandleCharts)

	http.HandleFunc("/index", Dashboard.HandleIndex)

	http.HandleFunc("/appro", Approvisionnement.HandleAppro)

	http.HandleFunc("/HistoAppro",HistoriqueApprovisionnement.HandleHistoAppro)

	http.HandleFunc("/sortie", Sortie.HandleSortie)

	http.HandleFunc("/SortieInst", ControllerSortieInst.HandleSortieInst)

	http.HandleFunc("/SortieArticle", controllerSortieArticle.HandleSortieArticle)

	http.HandleFunc("/receiveSortInst", ControllerSortieInst.ReceiveAjax)

	http.HandleFunc("/OperationSortie", controllerSortieArticle.OperationSortie)

	http.HandleFunc("/HistoSortRetour",HistoSortRetour.HandleSortRetour)

	http.HandleFunc("/Histo", Historique.HandleHisto)

	http.Handle("/Login_v1/",http.StripPrefix("/Login_v1",http.FileServer(http.Dir("Login_v1/"))))

	http.Handle("/webPages/",http.StripPrefix("/webPages",http.FileServer(http.Dir("webPages/"))))

		err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
