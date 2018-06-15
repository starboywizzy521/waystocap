package Login

import (
	"strings"
)

var nomutilisateur string="admin"
var motdepase string="waystocap"

func Login (username string,password string) (string, bool) {

	var Errors string

		if len(username)==0 || len(password)==0 {
			if username == "" {
				Errors= "Username is required";
			}
			if password == "" {
				Errors= "Password is required";
			}
		} else {
			if strings.Compare(nomutilisateur,username)==0 && strings.Compare(motdepase,password)==0{
				return "",true
			}else {
				Errors= "Incorrect username/password combination"
			}
		}
		return Errors,false
	}

