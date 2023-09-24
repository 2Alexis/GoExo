package Util

import (
	"fmt"
	"os"
)

func LireChoixUser() int {
	var choix int
	fmt.Println("Entrez votre choix :")
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de lecture de l'entrée de l'utilisateur :", err)
		os.Exit(1)
	}
	return choix
}
func Menu() int {
	fmt.Println("Menu")
	fmt.Println("Que souhaitez vous faire ")
	fmt.Println("1. Ajouter un contact ")
	fmt.Println("2. Afficher la liste des contacts ")
	fmt.Println("3. Supprimer un contact")
	fmt.Println("4. Quitter")
	return LireChoixUser()
}
