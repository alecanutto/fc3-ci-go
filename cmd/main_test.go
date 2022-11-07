package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido. Resultado: %d Experado: %d", total, 30)
	}
}
