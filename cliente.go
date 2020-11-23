package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client()
}

func AgregarCalificacion() map[string]map[string]float64 {
	var materias = make(map[string]map[string]float64)
	var nombre, mat string
	var calif float64

	fmt.Print("Alumno: ")
	fmt.Scanln(&nombre)
	fmt.Print("Materia: ")
	fmt.Scanln(&mat)
	fmt.Print("Calificacion: ")
	fmt.Scanln(&calif)

	alumno := make(map[string]float64)
	alumno[nombre] = calif
	materias[mat] = alumno

	return materias
}

func client() {

	c, error := rpc.Dial("tcp", "127.0.0.1:9999")
	if error != nil {
		fmt.Println(error)
		return
	}
	var op int64
	for {
		fmt.Println("\n********************************")
		fmt.Println("1-Agregar calificación")
		fmt.Println("2-Obtener promedio de Alumno")
		fmt.Println("3-Obtener promedio General")
		fmt.Println("4-Obtener promedio de Materia")
		fmt.Println("0-Salir")
		fmt.Print("  Ingrese opción >>> ")
		fmt.Scanln(&op)
		fmt.Println("")
		switch op {
		case 1:
			var respuesta string
			error = c.Call("Server.AgregarCalificacion", AgregarCalificacion(), &respuesta)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("\nSe ha agregado correctamente")
			}
		case 2:
			var nombre, respuesta string
			fmt.Print("Ingrese el nombre del alumno: ")
			fmt.Scanln(&nombre)
			error = c.Call("Server.PromAlumno", nombre, &respuesta)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("\nEl promedio total del alumno ", nombre, " es ", respuesta)
			}
		case 3:
			var respuesta string
			error = c.Call("Server.PromGeneral", "", &respuesta)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("El promedio General es ", respuesta)
			}
		case 4:
			var mat, respuesta string
			fmt.Print("Ingrese el nombre de la materia: ")
			fmt.Scanln(&mat)
			error = c.Call("Server.PromMateria", mat, &respuesta)
			if error != nil {
				fmt.Println(error)
			} else {
				fmt.Println("\nEl promedio de la materia ", mat, " es ", respuesta)
			}
		case 0:
			return
		}
	}
}

