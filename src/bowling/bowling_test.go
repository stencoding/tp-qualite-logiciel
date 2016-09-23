package exo2

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

/*func TestNullScore(t *testing.T) {
	input := []Frame{{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0},{0, 0}}
	expected := 0
	// err := scoreChecker(input, expected, nil)
	// if err != nil
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}*/

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

func TestNbTuples(t *testing.T) {
	input := []Frame{{7, 5},{13, -2}}

	expectedError:=fmt.Errorf("Pas exactement 10 tuples")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestValeurTuples(t *testing.T) {
	input := []Frame{{-5, 8},{13, -2},{7, 5},{13, -2},{7, 5},{13, -2},{7, 5},{13, -2},{7, 5},{13, -2}}

	expectedError:=fmt.Errorf("Valeur négative dans un tuple")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

}

func TestSommeValeurTuples(t *testing.T) {
	input := []Frame{{5, 6},{8, 2},{7, 5},{13, 2},{7, 5},{13, 2},{7, 5},{13, 4},{7, 5},{13, 2}}

	expectedError:=fmt.Errorf("Somme d'un tuple est supérieur à 10")
	expected := 0

	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}

}

func TestScoreSpare(t *testing.T) {
	// Spare en début de partie
	input := []Frame{{8, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{0, 2}}
	expected := 77
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	// Spare au dernier tour
	input = []Frame{{6, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{8, 2}}
	expected = 78
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreStrike(t *testing.T) {
	// Strike en début de partie
	input := []Frame{{10,0},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{0, 2}}
	expected := 79
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
	// deux strikes à la suite
	input = []Frame{{6, 2},{10,0},{10,0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{8, 2}}
	expected = 104
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}

	// Strike au dernier tour
	input = []Frame{{3, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{10,0}}
	expected = 75
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}

	// deux strikes aux derniers tours
	input = []Frame{{3, 2},{5, 2},{4, 5},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{10,0},{10,0}}
	expected = 87
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

/*func TestLanceBonusSpare(t *testing.T) {
	// deux strikes à la suite
	input := []Frame{{6, 2},{10,0},{10,0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{8, 2},{5}}
	expected := 104
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}


func TestLanceBonusStrike(t *testing.T) {
	// deux strikes à la suite
	input := []Frame{{6, 2},{10,0},{10,0},{5, 2},{7, 0},{7, 2},{7, 1},{1, 4},{7, 1},{10,0},{5,2}}
	expected := 104
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}*/