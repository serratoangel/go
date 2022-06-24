package main

import "fmt"

func main() {
	
}

/*
Junio 23 2022
Código Facilito
Clase 3

- version go
go version
- nombre paquete
package main
- imprimir en consola
import "fmt"
- funcion principal, debe ser main y no recibe parametros
func main()
- imprimir con salto de linea
fmt.Println("Hola mundo!!!")
- compilar un programa, main.go se refiere al nombre del programa en este caso
go build main.go
- correr sin necesidad de compilar
go run main.go

Clase 4

- declarar variables
var nombre string => tipo string se inicia como vacio
var edad int => tipo entero se inicia como cero
- no se puede cambiar el tipo de dato de la variable
var nombre string
var edad int
var name string = "Angel Serrato"
nombre = "Eduardo Ismael"
edad = 35
fmt.Println(nombre)
fmt.Println(edad)
fmt.Println(name)

Clase 5, 6

- se puede crear la variable sin el tipo de dato
nombre := "Angel Serrato"
edad := 35
var altura = 1.83 valor float
- para imprimir el valor de varias variables en consola, solo se separa con una coma
fmt.Println(nombre, edad, altura)
fmt.Println("El nombre es:", nombre, edad, altura)

Clase 7

- declarar varias variables
	var nombre, apellido, pais string
	var nombre, apellido, pais = "Angel", "Serrato", "Colombia"
	nombre, apellido, pais := "Angel", "Serrato", "Colombia"

Clase 8

- declarar constantes, se hacen fuera de la funcion principal
	const Curso string = "Curso de Go"

Clase 9
- operadores relacionales, cuando estos se usan el resultado es un valor booleano, true or false
>
<
>=
<=
==
!=
- para comparar strings, si tiene un espacio demas los strings no daran iguales
	var nombre = "Angel"
	var resultado = nombre == "Angel"

Clase 10
- operadores logicos
&&
||
!
- ejemplo
	resultado := 20 == 20 && 30 == 30 && ( 90 < 100 || 100 == 90 )
	resultado := 20 == 20 && 30 == 30
	resultado := !true
	fmt.Println(resultado)

Clase 11
- definir n cantidad de constantes en secuencia
const Domingo int = 0
const Lunes int = 1
const Martes int = 2
const Miercoles int = 3
const Jueves int = 4
const Viernes int = 5
const Sabado int = 6
- forma correcta
const (
	Domingo int = 0
	Lunes int = 1
	Martes int = 2
	Miercoles int = 3
	Jueves int = 4
	Viernes int = 5
	Sabado int = 6
)
- palabra reservada iota, esto le indica que las demas tendran valor en secuencia ya que el valor inicial es 0
const (
	Domingo int = iota // 0
	Lunes 
	Martes 
	Miercoles 
	Jueves 
	Viernes 
	Sabado 
)
- si se quiere iniciar desde otro numero 
	Domingo int = iota + 1 // 1

Clase 12
- strings
	var nombre string // inicia vacio
	var curso string = "Go"
	var dato = "Angel"
	hola := "Didier"
- para la longitud
	longitud := len(hola) //cantidad de caracteres del string se toman los espacios en blanco
- para acceder al indice del string, en este caso es un entero positivo de 8 bits
	primerCaracter := dato[0] // Character -> uint8
- para obtener el caracter en lugar del numero raro
	fmt.Printf("%c\n", primerCaracter)
- para imprimir el ultimo caracter del string
	primerCaracter := dato[ len(dato) - 1 ]

Clase 13 
- lectura de valores de teclado
	
*/