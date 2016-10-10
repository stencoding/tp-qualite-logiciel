package bowling

import (
	"fmt"
)

type Frame struct {
	firstThrow  int
	secondThrow int
}


// Vérife si le dernier lancer est un strike {10,0}
func isStrike(game []Frame, key int) bool {

	if game[key].firstThrow == 10 && game[key].firstThrow == 0 {
		return true
	}

	return false
}

// Vérife si le dernier lancer est un spare
func isSpare(game []Frame, key int) bool {
fmt.Println("titoi")
	if isStrike(game, key){
		return false
	}

	if game[key].firstThrow + game[key].secondThrow == 10 {
		return true
	}

	return false
}

/*func nbLancerBonus(game []Frame) int {
	nb := 0
	if game[10].firstThrow != 0 {
		nb = 1
	}
	if game[10].secondThrow != 0 {
		nb = nb + 1
	}
	return nb
}*/

func GetScore(game []Frame) (int, error) {
	score := 0

	/*if len(game) != 10 {
		// pour qu'il y ai 11 tuples, il faut que le 10ème tuple soit un spare ou un strike
		if len(game) == 11 {

			if game[9].firstThrow < 0 || game[9].secondThrow < 0 {
				return 0, fmt.Errorf("Valeur négative dans un tuple")
			}

			// si 10ème tuple n'est ni un spare ni un strike
			if !isSpareLastShot(game) || !isStrikeLastShot(game) {
				return 0, fmt.Errorf("Le dernier lancé n'est pas un spare ou un strike")
			}

			// si 10ème tuple est un spare
			if isSpareLastShot(game) {
				if nbLancerBonus(game) != 1 {
					return 0, fmt.Errorf("Spare en fin de partie, il y a un lancer en trop")
				}
			}
			if isStrike(game) {
				if nbLancerBonus(game) != 2 {
					return 0, fmt.Errorf("Strike en fin de partie, il y a un lancer en trop")
				}
			}

		} else {
			return 0, fmt.Errorf("Pas exactement 10 tuples")
		}
	}*/

	if len(game) < 10 {
		fmt.Println("bingo")
		fmt.Printf("%d\n", len(game))
		return 0, fmt.Errorf("Pas exactement 10 tuples")
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
			if !isStrike(game, i-1) || !isSpare(game, i-1) {
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