///ARBOL BINARIO CALCULADOR
// Ana Maria Nates
// Jesus David Romero

package main

import (
  "fmt"
  "strconv"
)

type Arbol struct{
  Izquierda *Arbol
  Valor string
  Derecha *Arbol
}

func RecorrerInorden(t *Arbol){
  if t==nil{
    return
  }
  RecorrerInorden(t.Izquierda)
  fmt.Print(t.Valor)
  RecorrerInorden(t.Derecha)
}

func Check(t *Arbol) bool,string{
  //si es hoja retorna numero
    if t.Izquierda == nil && t.Derecha== nil {
      _,err:= strconv.Atoi(t.Valor)
        if err == nil{
          return true, " "
        }
        return false, "ERROR DE SINTAXIS: se encontro el String "+t.Valor+" y se esperaba un número"
    }

  }

func Calcular(t *Arbol) int {
//si es hoja retorna numero
  if t.Izquierda == nil && t.Derecha== nil {
    num, _ := strconv.Atoi(t.Valor)
    return num
  }
  switch t.Valor{
    case "+": return Calcular(t.Izquierda)+Calcular(t.Derecha)
    case "-": return Calcular(t.Izquierda)-Calcular(t.Derecha)
    case "*": return Calcular(t.Izquierda)*Calcular(t.Derecha)
    case "/": return Calcular(t.Izquierda)/Calcular(t.Derecha)
    default: return 1000
  }
}

func main(){
  t1:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "-", &Arbol{nil,"2",nil} }, "+", &Arbol{&Arbol{nil,"10",nil}, "*",&Arbol{nil,"3",nil} } }
  fmt.Print("\n")
  RecorrerInorden(t1)
  fmt.Print(" = ",Calcular(t1))
  fmt.Print("\n")
  t2:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "-", &Arbol{nil,"2",nil} }, "+", &Arbol{&Arbol{nil,"10",nil}, "*",&Arbol{&Arbol{nil,"2",nil},"+",&Arbol{nil,"2",nil}}}}
  fmt.Print("\n")
  RecorrerInorden(t2)
  fmt.Print(" = ",Calcular(t2))
  fmt.Print("\n")
}
