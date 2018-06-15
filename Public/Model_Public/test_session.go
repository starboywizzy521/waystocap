package Model_Public

import (
	"github.com/gorilla/sessions"
	"fmt"
)

func Test_session (session *sessions.Session){

	if session!=nil {
		fmt.Println("authentifié")
	} else {
	fmt.Println("non authentifié")
	}
	/*if session.Values["authentification"]!=true{
	}*/
}
