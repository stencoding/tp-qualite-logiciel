package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)
	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

// à quoi sert cette function ??
func TestNullScore(t *testing.T) {
	input := []Frame{{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0}}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// test le score total d'une partie sans spare ni strike
func TestScore(t *testing.T) {
	input := []Frame{{1, 0},{1, 0},{1, 0},{1, 0},{1, 0},{1, 0},{1, 0},{1, 0},{1, 0},{1, 0}}
	expected := 10
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	input = []Frame{{1, 2},{1, 2},{1, 2},{1, 2},{1, 2},{1, 2},{1, 2},{1, 2},{1, 2},{1, 2}}
	expected = 30
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// test de moins de 10 tuples
func TestNbTuples(t *testing.T) {
	input := []Frame{{7, 3},{1, 2}}

	expectedError:=fmt.Errorf("Il y a moins de 10 tuples")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// test une valeur négative dans un au moins un tuple
func TestValeurTuples(t *testing.T) {
	input := []Frame{{-5, 8},{13, -2},{7, 5},{13, -2},{7, 5},{13, -2},{7, 5},{13, -2},{7, 5},{13, -2}}

	expectedError:=fmt.Errorf("Valeur négative dans un tuple")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

}

// test la somme des valeurs d'au moins un tuple supérieur à 10
func TestSommeValeurTuples(t *testing.T) {
	input := []Frame{{5, 6},{8, 2},{7, 5},{13, 2},{7, 5},{13, 2},{7, 5},{13, 4},{7, 5},{13, 2}}

	expectedError:=fmt.Errorf("Somme d'un tuple est supérieur à 10")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

}

// test le score total d'un partie avec un spare
func TestScoreSpare(t *testing.T) {
	// Spare en début de partie
	input := []Frame{{8, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{0, 2}}
	expected := 77
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	// Spare au dernier tour
	input = []Frame{{6, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{8, 2}, {3,0}}
	expected = 84
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

// test le score total d'un partie avec un strike
func TestScoreStrike(t *testing.T) {
	// Strike en début de partie
	input := []Frame{{10,0},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{0, 2}}
	expected := 79
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	// deux strikes à la suite
	input = []Frame{{6, 2},{10,0},{10,0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{8, 2}, {6,0}}
	expected = 116
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}

	// Strike au dernier tour
	input = []Frame{{3, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{10,0},{2,3}}
	expected = 85
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}

	// deux strikes aux derniers tours
	input = []Frame{{3, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{10,0},{10,0}, {2,3}}
	expected = 99
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}


// Test si plus de 10 tuples et dernier lancer n'est pas un strike ou spare
func TestLanceFinStrikeSpare(t *testing.T) {
	// le 10 lancer n'est pas un strike ou un spare
	input := []Frame{{6, 2},{10, 0},{10, 0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{9, 0},{5, 2}}

	expectedError:=fmt.Errorf("Pas exactement 10 tuples")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

	// le 10ème tuple est un spare mais ne respecte pas la règle d'un seul lancer bonus
	input = []Frame{{6, 2},{10, 0},{10, 0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{9, 1},{5, 2}}

	expectedError=fmt.Errorf("Spare en fin de partie, il y a un lancer en trop")
	expected = 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
	// le 10ème tuple est un spare donc il manque un lancer bonus
	input = []Frame{{6, 2},{10, 0},{10, 0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{9, 1}}

	expectedError=fmt.Errorf("Spare en fin de partie, il manque le lancer bonus")
	expected = 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

	// le 10ème tuple est un strike donc il manque les lancers bonus
	input = []Frame{{6, 2},{10, 0},{10, 0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{10, 0}}

	expectedError=fmt.Errorf("Strike en fin de partie, il manque les lancers bonus")
	expected = 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

}