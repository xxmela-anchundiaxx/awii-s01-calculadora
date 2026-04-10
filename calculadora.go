package main

import (
	"fmt"
)

func main() {

	var historial string
	contador := 0

	for { // bucle infinito

		var a float64
		var b int
		var operacion string

		fmt.Println("\n---- CALCULADORA CIENTÍFICA ----")

		// Entrada
		fmt.Print("Ingresa el primer número: ")
		fmt.Scan(&a)

		fmt.Print("Ingresa el segundo número: ")
		fmt.Scan(&b)

		fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
		fmt.Scan(&operacion)

		var resultadoStr string
		errorOperacion := false

		switch operacion {

		case "+":
			res := a + float64(b)
			resultadoStr = fmt.Sprintf("%.2f + %d = %.2f", a, b, res)

		case "-":
			res := a - float64(b)
			resultadoStr = fmt.Sprintf("%.2f - %d = %.2f", a, b, res)

		case "*":
			res := a * float64(b)
			resultadoStr = fmt.Sprintf("%.2f * %d = %.2f", a, b, res)

		case "/":
			if b == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
				errorOperacion = true
				break
			}
			res := a / float64(b)
			resultadoStr = fmt.Sprintf("%.2f / %d = %.2f", a, b, res)

		case "^":
			if b < 0 {
				fmt.Println("Error: el exponente debe ser positivo")
				errorOperacion = true
				break
			}

			resultado := 1.0
			if b == 0 {
				resultado = 1
			} else {
				for i := 0; i < b; i++ {
					resultado *= a
				}
			}

			resultadoStr = fmt.Sprintf("%.2f ^ %d = %.2f", a, b, resultado)

		case "!":
			if a < 0 {
				fmt.Println("Error: no existe factorial de números negativos")
				errorOperacion = true
				break
			}

			n := int(a)
			resultado := 1

			if n == 0 {
				resultado = 1
			} else {
				for i := 1; i <= n; i++ {
					resultado *= i
				}
			}

			resultadoStr = fmt.Sprintf("%d ! = %d", n, resultado)

		default:
			fmt.Println("Error: operación no reconocida")
			errorOperacion = true
		}

		// Guardar en historial si no hubo error
		if !errorOperacion {
			fmt.Println("Resultado:", resultadoStr)

			contador++
			linea := fmt.Sprintf("%d) %s", contador, resultadoStr)

			historial = historial + linea + "\n"
		}

		// Preguntar si continúa
		var opcion string
		fmt.Print("¿Otra operación? (s/n): ")
		fmt.Scan(&opcion)

		if opcion == "n" {
			break
		}
	}

	fmt.Println("\n==== HISTORIAL DE OPERACIONES ====")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", contador)
	fmt.Println("¡Hasta luego!")
}
