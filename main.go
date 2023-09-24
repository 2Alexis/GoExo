package main

import (
	"fmt"

	"github.com/alexi/GoEx/Test/Util"
	"github.com/alexi/GoEx/Test/contacts"
)

func main() {
	for {
		choix := Util.Menu()
		switch choix {
		case 1:
			contacts.AddContact()
			fmt.Println("Vous avez choisi d'ajouter un contact :")
		case 2:
			contacts.ShowContacts()
			fmt.Println("Vous avez choisi d'afficher vos contacts :")
		case 3:
			contacts.SuprContact()
			fmt.Println("Vous avez choisi de supprimer un contact :")

		case 4:
			fmt.Println("Au revoir !")
			return
		}
	}
}
