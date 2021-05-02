package GoCalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define nuestro tipo
type calculo struct{}

// Agrega funcion de operación
func (calculo) operacion(entrada string, operador string) int {
	// Realiza separación de datos
	valoresString := strings.Split(entrada, operador)
	operador1 := parsear(valoresString[0])
	operador2 := parsear(valoresString[1])

	// Realiza revisión de operaciones
	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		return 0
	}
}

// Obtiene entrada en int
func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)

	return operador
}

// Lee desde consola
func LeerEntrada(mensaje string) string {
	// Muestra mensaje
	fmt.Print(mensaje)

	// Define escáner de pantalla
	scanner := bufio.NewScanner(os.Stdin)

	// Inicia el escaner de datos
	scanner.Scan()

	// Asigna el escaner a un string
	operacion := scanner.Text()

	// Retorna el texto
	return operacion
}

func main() {
	// Obtiene información
	entrada := leerEntrada("Ingrese operación a realizar: ")
	operador := leerEntrada("Ingrese operador matematico: ")

	// Imprime operación a realizar
	fmt.Printf("Entrada: %s \n", entrada)
	// Imprime operación a realizar
	fmt.Printf("Operación: %s \n", operador)

	// Instancia struct
	calc := calculo{}

	// Retorna resultado
	fmt.Printf("Resultado: %d \n", calc.operacion(entrada, operador))
}
