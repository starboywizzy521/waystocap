package Login

import "strings"

func ChangeLogin(oldusername string,oldpassword string,newusername string,newpassword string) (string, bool){

	if strings.Compare(nomutilisateur,oldusername)==0 && strings.Compare(motdepase,oldpassword)==0{
		nomutilisateur=newusername
		motdepase=newpassword
		return "",true
	}
	return "combinaison incorrecte",false
}