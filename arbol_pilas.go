package main

import (
  "fmt"
  "strings"
  "strconv"
)


//-----------------------------
type Arbol struct{
  Izquierda *Arbol
  Valor string
  Derecha *Arbol
}

func NewArbolVacio() *Arbol{
  return &Arbol{nil,"",nil}
}

func Calcular(t *Arbol) int {
  if t.Izquierda == nil && t.Derecha== nil{
    num,err := strconv.Atoi(t.Valor)
    if err == nil{
      return num
    }
  }
  switch t.Valor{
    case "+": return Calcular(t.Izquierda)+Calcular(t.Derecha)
    case "-": return Calcular(t.Izquierda)-Calcular(t.Derecha)
    case "*": return Calcular(t.Izquierda)*Calcular(t.Derecha)
    case "/": return Calcular(t.Izquierda)/Calcular(t.Derecha)
    default: return 1000
  }
}

func RecorrerInorden(t *Arbol){
  if t==nil{
    return
  }
  RecorrerInorden(t.Izquierda)
  fmt.Print(t.Valor)
  fmt.Print(" ")
  RecorrerInorden(t.Derecha)
}
//-----------------------------
/*type Node struct {
  Value *Arbol
}*/

// Pilas
type Stack struct {
  nodes []*Arbol
  count int
}

// NewStack returns a new pila.
func NewStack() *Stack {
  return &Stack{}
}

// Push adds a node to the pila.
func (s *Stack) Push(n *Arbol) {
  s.nodes = append(s.nodes[:s.count], n)
  s.count++
}

// Pop removes and returns a node from the pila in last to first order.
func (s *Stack) Pop() *Arbol {
  if s.count == 0 {
    return nil
  }
  s.count--
  return s.nodes[s.count]
}
//-----------------------------


func Armar_arbol(arreglo []string) *Arbol {
  pila := NewStack()
  Arbolf:= NewArbolVacio()
  cont := 0
  for i := 0;i < len(arreglo); i++{
    elemento:=arreglo[i]
    _,err := strconv.Atoi(elemento)

    if err == nil{
      //if
      if cont <= 1{
        pila.Push(&Arbol{nil,elemento,nil})
        cont++
      }else{
        fmt.Println("** ERROR ** Es un arbol binario no puede ingresar mas de dos numeros seguidos")
        return Arbolf
      }
      /*else
        fmt.Println("** ERROR ** Posición",i,"Ingresó",elemento,"sin operador")
      */
    }else{
      //if pila.count > 0 && pila.count%2 ==0
      fmt.Println(cont)

        if pila.count > 0 && cont == 2 || cont==0{
        t1:=pila.Pop()
        t2:=pila.Pop()
          pila.Push(&Arbol{t2,elemento,t1})
          cont=0
        //Arbolf= &Arbol{t1,elemento,t2}
      }else{
        fmt.Println("** ERROR ** Posición",i,": Ingresó",elemento,"sin elementos suficiente para operar")
        return Arbolf
      }
    }
  }
  if pila.count == 1{
    Arbolf=pila.Pop()
    return Arbolf
  }else{
    fmt.Println("** ERROR ** No ingresó operación final")
    return Arbolf
  }


}

func main() {
  entrada:= "52+1-+"
  cadena:=strings.Split(entrada, "")
  fmt.Println("Cadena inicial","\t",cadena)
  arbolito := Armar_arbol(cadena)
  if arbolito.Valor !="" {
    fmt.Print("Resultado ","\t")
    RecorrerInorden(arbolito)
    fmt.Println("=",Calcular(arbolito))
  }

}
