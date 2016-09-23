package exo2

import (
	"fmt"
)

type Frame struct {
	firstThrow  int
	secondThrow int
}


// Vérife si le dernier lancer est un strike
func isStrike(game []Frame) bool {

	if game[9].firstThrow == 10 {
		return true
	}

	return false
}

// Vérife si le dernier lancer est un spare
func isSpare(game []Frame) bool {

	lastFirstThrow := game[9].firstThrow
	lastSecondThrow := game[9].secondThrow

	if lastFirstThrow != 10 && (lastFirstThrow + lastSecondThrow == 10) {
		return true
	}

	return false
}

func GetScore(game []Frame) (int, error) {
	score := 0

	if len(game) != 10 {
		if len(game) == 11 {

			// si le dernier n'est pas un spare on génère une erreur
			if !isSpare(game) {
				return 0, fmt.Errorf("Le dernier lancé n'est pas un spare")
			}

			// si le dernier n'est pas un strike on génère une erreur
			if !isStrike(game) {
				return 0, fmt.Errorf("Le dernier lancé n'est pas un strike")
			}

		} else {
			return 0, fmt.Errorf("Pas exactement 10 tuples")
		}
	}

	for i := 0; i < len(game); i++ {
		if game[i].firstThrow < 0 || game[i].secondThrow < 0 {
			return 0, fmt.Errorf("Valeur négative dans un tuple")
		}

		if game[i].firstThrow + game[i].secondThrow > 10 {
			return 0, fmt.Errorf("Somme d'un tuple est supérieur à 10")
		}

		//calcul du strike
		if game[i].firstThrow == 10 {
			// deux strikes à la suite
			if(i+1 < len(game)) && game[i+1].firstThrow == 10 {
				score = score + game[i].firstThrow + game[i+1].firstThrow
				// on est au début de la partie
				if i+2 < len(game) {
					score = score + game[i+2].firstThrow
				}				
			} else { // un seul strike
				score = score + game[i].firstThrow
				if i+1 < len(game) {
					score = score + game[i+1].firstThrow + game[i+1].secondThrow	
				} 
			}
		} else {
			score = score + game[i].firstThrow + game[i].secondThrow
		}


		// calcul du spare
		if game[i].firstThrow != 10 && game[i].firstThrow + game[i].secondThrow == 10 && (i+1 < len(game)) {			
			score = score + game[i+1].firstThrow
		}

	}

	return score, nil
}

/*
version 2:
	Mettre 10 tuples à chaque input
	Ne pas mettre de valeur négative
	La somme du tuple ne doit pas dépasser 10

version 3 : pas de modif des tests

version bonus : 

*/