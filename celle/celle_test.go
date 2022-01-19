package cell

import "testing"

var cell bool

func TestInitCell(t*Testing){
	wanted := false; //bruker semantikken i Golangs konseptuelle modell
	state := InitCell;
	if state != wanted{
		t.Errorf("Feil, fikk %t, ønsket %t, state, wanted")
	}

}
