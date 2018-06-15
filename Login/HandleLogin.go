package Login

import (
	"net/http"
	"html/template"
	"github.com/gorilla/sessions"
	"fmt"
)

var Store = sessions.NewCookieStore([]byte("something-very-secret"))
var Session *sessions.Session

func HandleLogin(w http.ResponseWriter,r *http.Request){

	if r.Method == "GET" {
		t, _ := template.ParseFiles("Login/index.html")
		t.Execute(w,nil)
	}else if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username");
		password := r.Form.Get("password");
		err,boolean:=Login(username,password)

		if err=="" && boolean==true{
			Session, _= Store.Get(r, "cookie-name")
			Session.Values["username"]=username
			Session.Values["authentification"]=true
			Session.Save(r,w)
			http.Redirect(w,r,"index",301)
		}else {
			fmt.Println(err)
			t, _ := template.ParseFiles("Login/index.html")
			t.Execute(w,nil)
		}
	}
}
