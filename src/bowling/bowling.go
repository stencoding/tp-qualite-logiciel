package bowling

import (
	"fmt"
)

type Frame struct {
	firstThrow  int
	secondThrow int
}

var score = 0

// Vérife si le ième lancer est un strike
func IsStrike(game []Frame, i int) bool {
	if game[i].firstThrow == 10 && game[i].secondThrow == 0 {	
		return true
	}
	return false
}

// Vérife si le ième lancer est un spare 
func IsSpare(game []Frame, i int) bool {
	if IsStrike(game, i){
		return false
	}
	if (game[i].firstThrow + game[i].secondThrow) == 10 {
		return true
	}
	return false
}

func GetScore(game []Frame) (int, error) {
	score := 0
	// Gestion des erreurs
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
			if !IsStrike(game, i-1) && !IsSpare(game, i-1) {
				return 0, fmt.Errorf("Pas exactement 10 tuples")
			}
	
			// le 10ème tuple est un spare
			if IsSpare(game, i-1) {
				if game[i].secondThrow != 0 {
					return 0, fmt.Errorf("Spare en fin de partie, il y a un lancer en trop")
				}
			}
		}
		// Vérification présence lancers bonus
		if i == 9 && len(game) == 10 {
			if IsSpare(game, i) {				
				return 0, fmt.Errorf("Spare en fin de partie, il manque le lancer bonus")
			}
			if IsStrike(game, i) {
				return 0, fmt.Errorf("Strike en fin de partie, il manque les lancers bonus")
			}
		}

		// Calcul du score

		//calcul du strike
		if IsStrike(game, i) {
			// deux strikes à la suite
			if(i+1 < len(game)) && IsStrike(game, i+1) {
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
		if !IsStrike(game, i) && IsSpare(game, i) && (i+1 < len(game)) {		
			score = score + game[i+1].firstThrow
		}
	}
	return score, nil
}
