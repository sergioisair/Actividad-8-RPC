package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
}

type Server struct {
	materias map[string]map[string]float64
	alumnos  map[string]map[string]float64
}

func (this *Server) AgregarCalificacion(data map[string]map[string]float64, r *string) error {
	for materia, alumno := range data {
		for nombreAlumno, calif := range alumno {
			if this.materias[materia] == nil {
				this.materias[materia] = alumno
			} else {
				this.materias[materia][nombreAlumno] = calif
			}
			if this.alumnos[nombreAlumno] == nil {
				mat := make(map[string]float64)
				mat[materia] = calif
				this.alumnos[nombreAlumno] = mat
			} else {
				this.alumnos[nombreAlumno][materia] = calif
			}
		}
	}
	fmt.Print("Llamada a función AgregarCalificación")
	return nil
}

func (this *Server) PromAlumno(nombreAlumno string, r *string) error {
	var total float64 = 0
	_, existe := this.alumnos[nombreAlumno]
	if existe {
		for _, calif := range this.alumnos[nombreAlumno] {
			total += calif
		}
		var numMat = float64(len(this.alumnos[nombreAlumno]))
		total = total / numMat
		*r = strconv.FormatFloat(total, 'f', 2, 64)
	} else {
		error := errors.New("El alumno no existe")
		fmt.Println(error)
		return error
	}
	fmt.Println("\nLlamada a funcion PromAlumno")
	return nil

}

func (this *Server) PromGeneral(noData string, r *string) error {
	var total, num float64
	num = 0
	for _, alumno := range this.materias {
		for _, calif := range alumno {
			total += calif
			num += 1
		}
	}
	if num <= 0 {
		error := errors.New("Está vacío")
		fmt.Println(error)
		return error
	}
	var promTotal = total / num
	*r = strconv.FormatFloat(promTotal, 'f', 2, 64)
	fmt.Println("\nLlamada a funcion PromGeneral")
	return nil
}

func (this *Server) PromMateria(materia string, r *string) error {
	_, existe := this.materias[materia]
	if existe {
		var total float64
		for _, calif := range this.materias[materia] {
			total += calif
		}
		var numMat = float64(len(this.materias[materia]))
		total = total / numMat
		*r = strconv.FormatFloat(total, 'f', 2, 64)
	} else {
		error := errors.New("La materia no existe")
		fmt.Println(error)
		return error
	}
	fmt.Println("\nLlamada a funcion PromMateria")
	return nil
}

func server() {
	var server = new(Server)
	server.materias = make(map[string]map[string]float64)
	server.alumnos = make(map[string]map[string]float64)
	rpc.Register(server)
	x, error := net.Listen("tcp", ":9999")
	fmt.Println("Servidor activado")
	if error != nil {
		fmt.Println(error)
	}
	for {
		con, error := x.Accept()
		if error != nil {
			fmt.Println(error)
			continue
		}
		go rpc.ServeConn(con)
	}

}
