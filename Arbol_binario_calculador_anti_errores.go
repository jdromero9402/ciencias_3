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

func CheckNum(t *Arbol) (bool, string){
    if t.Izquierda == nil && t.Derecha== nil {
      _,err:= strconv.Atoi(t.Valor)
        if err == nil{
          return true, " "
        }
        return false, "ERROR DE SINTAXIS: se encontro el String "+t.Valor+" donde se esperaba un número"
    }
    return CheckNum(t.Izquierda)
    return CheckNum(t.Derecha)
  }

  func CheckSign(t *Arbol) (bool, string){
    if t.Izquierda != nil || t.Derecha != nil {
        _,err:= strconv.Atoi(t.Valor)
        val := false
        if err == nil{
          return false, "ERROR DE SINTAXIS: se encontro el simbolo "+t.Valor+" donde se esperaba un operador"
        } else{
          val=true
        }
        return val, " "
    }
    return CheckSign(t.Izquierda)
    return CheckSign(t.Derecha)
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
  t1:= &Arbol{ &Arbol{&Arbol{nil,"x",nil}, "x", &Arbol{nil,"2",nil} }, "+", &Arbol{&Arbol{nil,"10",nil}, "*",&Arbol{nil,"3",nil} } }
  fmt.Print("\n")
 err, errtxt := CheckNum(t1)
  //err2, errtxt2 := CheckSign(t1)
 // fmt.Println(CheckSign(t1))
  if err==true{
    RecorrerInorden(t1)
    fmt.Print(" = ",Calcular(t1))
    fmt.Print("\n")
  }else{
    fmt.Println(errtxt)
  }
  /*
  t2:= &Arbol{ &Arbol{&Arbol{nil,"5",nil}, "-", &Arbol{nil,"2",nil} }, "+", &Arbol{&Arbol{nil,"10",nil}, "*",&Arbol{&Arbol{nil,"2",nil},"+",&Arbol{nil,"2",nil}}}}
  fmt.Print("\n")
  RecorrerInorden(t2)
  fmt.Print(" = ",Calcular(t2))
  fmt.Print("\n")*/
}