package Dashboard

import (
	"net/http"
	"html/template"
	"waystocap/code/Login"
	"waystocap/code/Public/Model_Public"
	"waystocap/code/Tables/Model"
)

func HandleIndex(w http.ResponseWriter,r *http.Request){

	Model_Public.Test_session(Login.Session)
	PageData:=Model.Extras{Model.RecupProduits()}
	t:=template.Must(template.ParseFiles("Dashboard/index.html","Tables/Vue/gg.html","Charts/charts.html","webPages/include-header.html","webPages/include-body.html","webPages/include-footer.html"))
	t.Execute(w,PageData)

}
