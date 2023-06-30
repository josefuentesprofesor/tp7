package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Ingrese un numero entero de 16 dígitos: ")
	fmt.Scanln(&numero)

	validar(numero)
}

func validar(numero int) (valido bool) {
	if numero < 1000000000000000 || numero > 9999999999999999 {
		fmt.Println("El numero ", numero, " no tiene 16 digitos")
		valido = false
		return
	}

	verificador := numero % 10
	numero = numero / 10

	fmt.Println("El digito verificador es ", verificador)

	fmt.Println("Aplicando calculo en ", numero)

	acumulador := 0
	for i := 0; i < 15; i++ {
		//Extraigo el ultimo dígito de la derecha
		digito := numero % 10

		if i%2 == 0 {
			//posicion par
			//TODO actualizar el acumulador
		}

		acumulador += digito

		//Elimino el ultimo dígito de la derecha
		numero = numero / 10
	}

	resultado := (acumulador * 9) % 10

	fmt.Println("El resultado es ", resultado)

	if resultado == verificador {
		//TODO imprimir el mensaje correspondiente
		valido = true
	} else {
		//TODO imprimir el mensaje correspondiente
		valido = false
	}

	return
}
