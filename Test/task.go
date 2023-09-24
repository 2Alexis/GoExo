package contacts

import "fmt"

type Contact struct {
	Name    string
	Surname string
	Email   string
}

var contactsList []Contact

func AddContact(name string, surname string, email string) {
	newcontact := Contact{
		Name:    name,
		Surname: surname,
		Email:   email,
	}
	contactsList = append(contactsList, newcontact)
}

func ShowContacts(name string, surname string, email string) {
	for _, contacts := range contactsList {
		fmt.Printf("Liste de vos contacts : \n Nom : %s\n Prénom : %s\n Email : %s\n", contacts.Name, contacts.Surname, contacts.Email)
	}
}
func SuprContact(name string) {
	for _, contacts := range contactsList {
		if contacts.Name == name {
			contactsList = append(contactsList, contacts)
		}
		fmt.Printf("Vous avez supprimé %s de vos contacts.", contacts.Name)
	}
}
