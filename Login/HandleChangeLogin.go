package Login

import (
	"net/http"
	"html/template"
)

func HandleChangeLogin(w http.ResponseWriter,r *http.Request) {

	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Login/ChangeLogin.html")
		t.Execute(w,nil)
	}else{
		err,boolean:=ChangeLogin(r.FormValue("oldusername"),r.FormValue("oldpassword"),r.FormValue("newusername"),r.FormValue("newpassword"))

		if err=="" && boolean==true{
			http.Redirect(w,r,"login",301)
		}else{
			t, _ := template.ParseFiles("Login/ChangeLogin.html")
			t.Execute(w,nil)
		}
	}
}