package main

import "testing"

func TestValidar(t *testing.T) {
	tarjetaPruebaMastercard := 5031755734530604
	tarjetaPruebaVisa := 4509953566233704
	tarjetaPruebaMaestro := 5000551000000407
	tarjetaPruebaVisaElectron := 4917481000000006
	tarjetaPruebaVisaCabal := 6035227716427021
	tarjetaPruebaMastercardInvalida := 5031755734530601
	tarjetaPruebaVisaInvalida := 4509953536233704

	result := validar(tarjetaPruebaMastercard)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaVisa)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaMaestro)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaVisaElectron)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaVisaCabal)
	if result != true {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, true)
	}

	result = validar(tarjetaPruebaMastercardInvalida)
	if result != false {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, false)
	}

	result = validar(tarjetaPruebaVisaInvalida)
	if result != false {
		t.Errorf("Fallo test, se obtuvo: %t, se esperaba: %t.", result, false)
	}
}
