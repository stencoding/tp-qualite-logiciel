package bowling

import (
	"fmt"
)

type Frame struct {
	firstThrow  int
	secondThrow int
}


// Vérife si le ième lancer est un strike
func isStrike(game []Frame, i int) bool {
	if game[i].firstThrow == 10 && game[i].secondThrow == 0 {	
		return true
	}
	return false
}

// Vérife si le ième lancer est un spare 
func isSpare(game []Frame, i int) bool {
	if isStrike(game, i){
		return false
	}
	if (game[i].firstThrow + game[i].secondThrow) == 10 {
		return true
	}
	return false
}

func GetScore(game []Frame) (int, error) {
	score := 0

	if len(game) < 10 {
		return 0, fmt.Errorf("Il y a moins de 10 tuples")
	}

	for i := 0; i < len(game); i++ {
		if game[i].firstThrow < 0 || game[i].secondThrow < 0 {
			return 0, fmt.Errorf("Valeur négative dans un tuple")
		}

		if game[i].firstThrow + game[i].secondThrow > 10 {
			return 0, fmt.Errorf("Somme d'un tuple est supérieur à 10")
		}

		// il existe un 11ème tuple
		if i==10 {
			// le 10ème tuple n'est pas un spare ou un strike
			if !isStrike(game, i-1) && !isSpare(game, i-1) {
			//	fmt.Printf("BOUH")	
				return 0, fmt.Errorf("Pas exactement 10 tuples")
			}
	
			// le 10ème tuple est un spare
			if isSpare(game, i-1) {
				//TODO : peut-être mettre nil
				if game[i].secondThrow != 0 {
					return 0, fmt.Errorf("Spare en fin de partie, il y a un lancer en trop")
				}
			}
		}
		// Vérification présence lancers bonus
		if i == 9 && len(game) == 10 {
			if isSpare(game, i) {				
				return 0, fmt.Errorf("Spare en fin de partie, il manque le lancer bonus")
			}
			if isStrike(game, i) {
				return 0, fmt.Errorf("Strike en fin de partie, il manque les lancers bonus")
			}
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
			// calcul d'un tuple sans spare et sans strike
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
	Mettre un lancer bonus quand la partie finie par un spare
	Mettre deux lancer bonus quand la partie finie par un strike
	Recalculer le bon score pour les tests concernés
*/