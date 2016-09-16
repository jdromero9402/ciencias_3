package main

import (
	"bufio"
  "fmt"
  "os"
	//"reflect"
)

type Node struct {
	Value int
	Protagonista string
  Nombre string
	Genero string
	Año string
	}

func (n *Node) String() string {
	return fmt.Sprint(n.Value, "->" ,n.Nombre)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) buscar() *Stack {

	s2 := NewStack()
	saux := NewStack()
	var nodoaux *Node

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nSeleccione la opción de busqueda \n 1. protagonista \n 2. titulo \n 3. genero \n 4. año")
	busq, _ := reader.ReadString('\n')


	switch busq {
	case "1\n":
			reader2 := bufio.NewReader(os.Stdin)
			fmt.Print("Ingrese el nombre del protagonista: ")
			opcion, _ := reader2.ReadString('\n')
			for s.count>0 {
				nodoaux = s.Pop()
				if nodoaux.Protagonista == opcion {
					s2.Push(nodoaux)
				}
				saux.Push(nodoaux)
			}

		case "2\r\n":
			reader2 := bufio.NewReader(os.Stdin)
			fmt.Print("Ingrese el nombre del titulo: ")
		  opcion, _ := reader2.ReadString('\n')
			for s.count>0 {
				nodoaux = s.Pop()
				if nodoaux.Nombre == opcion {
				s2.Push(nodoaux)
				}
				saux.Push(nodoaux)
			}

		case "3\r\n":
			reader2 := bufio.NewReader(os.Stdin)
			fmt.Print("Ingrese el genero: ")
		  opcion, _ := reader2.ReadString('\n')
			for s.count>0 {
				nodoaux = s.Pop()
				if nodoaux.Genero == opcion {
					s2.Push(nodoaux)
				}
				saux.Push(nodoaux)
			}

		case "4\r\n":
			reader2 := bufio.NewReader(os.Stdin)
			fmt.Print("Ingrese el año: ")
		  opcion, _ := reader2.ReadString('\n')
			for s.count>0 {
				nodoaux = s.Pop()
				fmt.Println(nodoaux.Año)
				if nodoaux.Año == opcion {
					s2.Push(nodoaux)
				}
				saux.Push(nodoaux)
			}
		default:
        fmt.Println("Intente de nuevo :( !")
		}

		for saux.count>0{
			s.Push(saux.Pop())
		}

	return s2
}

func (s *Stack) añadir(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el titulo de la pelicula: ")
	titulo, _ := reader.ReadString('\n')

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el nombre del protegonista de la pelicula: ")
	protagonista, _ := reader2.ReadString('\n')

	reader3 := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el año de la pelicula: ")
	año, _ := reader3.ReadString('\n')

	reader4 := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el genero de la pelicula: ")
	genero, _ := reader4.ReadString('\n')

	s.Push(&Node{s.count+1,protagonista,titulo,genero,año})
}

func (s *Stack) mostrar() {
	for i:=s.count-1;i>=0;i--{
			fmt.Println(s.nodes[i].Nombre+"\t"+s.nodes[i].Genero+"\t"+s.nodes[i].Protagonista+"\t"+s.nodes[i].Año)
	}
}

func main() {
	s := NewStack()
	/*s.Push(&Node{1,"Sylvester Stallone"+"\r\n","Rocky I"+"\r\n","Drama"+"\r\n","1987"+"\r\n"})
	s.Push(&Node{2,"Sylvester Stallone"+"\r\n","Rocky II"+"\r\n","Drama"+"\r\n","1997"+"\r\n"})
	s.Push(&Node{3,"Sylvester Stallone"+"\r\n","Rocky III"+"\r\n","Drama"+"\r\n","1985"+"\r\n"})
	s.Push(&Node{4,"Leonardo Dicaprio"+"\r\n","Titanic"+"\r\n","Romance"+"\r\n","1987"+"\r\n"})
	s.Push(&Node{5,"Leonardo Dicaprio"+"\r\n","Inception"+"\r\n","Drama"+"\r\n","1987"+"\r\n"})
*/

	//INGRESAR PELICULA
	//for opc == "1 \n"{
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Si desea agregar pelicula presione 1. para buscar presione 2.")
//	opc, _ =reader.ReadString('\n')
	s.añadir()
 //}
	//BUSCAR
	res := NewStack()
	res = s.buscar()
	if res.count == 0  {
			fmt.Println("No existen resultados para su busqueda D:")
	} else {
		res.mostrar()
	}
}
