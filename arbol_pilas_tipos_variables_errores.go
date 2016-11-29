/*

Ana María Nates
Jesús Romero
*/

package main

import (
  "fmt"
  "strings"
  "strconv"
  "bufio"
  "os"
  "regexp"
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

func Calcular(t *Arbol) float64 {
  if t.Izquierda == nil && t.Derecha== nil{
    num,err := strconv.ParseFloat(t.Valor,64)
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
    _,err := strconv.ParseFloat(elemento,64)

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

//Lee la cadena y guarda las variables en un map
func leerCadenas() (map[string][]string, []string, []string, []string, []string, []string) {
  reader := bufio.NewReader(os.Stdin)
  mapa := make(map[string][]string)
  var cad_final []string
  fmt.Print("Ingrese una sentencia por linea.\n Separe los terminos con un espacio: '12 3 + Y :='\n Termine con FIN:\n")
  var tiposV []string
  var ops []string
  var vars []string
  var nums []string
  for{
    e,_ := reader.ReadString('\n')
    if strings.Contains(e,"FIN") {
      break
    } else if strings.Contains(e,":=") {
      cad:=strings.Split(e," ")
    //  fmt.Println("cad:", cad)
    tiposVaux,opsAux,varAux,numsAux:= tiposVariables(cad)
    //Le mete las variables y tipos de variables al arreglo
    tiposV = append(tiposV,tiposVaux...)
    ops = append(ops,opsAux...)
    vars = append(vars,varAux...)
    nums = append(nums,numsAux...)
    v:=cad[len(cad)-2]
    fmt.Println("variable:", v)
    mapa[v] = cad[:len(cad)-2]
    //cad = cad[:len(cad)-3]
      cad_final = cad[:len(cad)] //Quitar el enter
    }
  }
  //fmt.Println(mapa)
//  fmt.Println("",cad_final)
  return mapa, cad_final, tiposV, ops, vars, nums
}

//modificar un arreglo
func modificarArreglo(arr map[string][]string, final []string) []string{
  cont := 0
  tamaño := -1
  var newArr2 []string
for cont != tamaño {
  var newArr []string
  tam := len(final)
    for i := 0;i < tam-2;i++{
      _,ver := arr[final[i]]
      if !ver{
        newArr = append(newArr, final[i])
      }else{
        arrAux := arr[final[i]]
        for i:=0;i < len(arrAux);i++{
          newArr = append(newArr, arrAux[i])
        }
      }
    }
    tam2 := len(newArr)
    cont = 0
    for j:=0;j<tam2-2;j++{
    //fmt.Println(tam2)
      /*elemento := newArr[j]
      _,err := strconv.ParseFloat(elemento,64)
      if err != nil{
          if elemento=="+" || elemento=="-" || elemento=="*" || elemento=="/"{
            cont++
          }else{
            cont++
          }
      }else{
        cont++
      }*/
      cont++
    }
    tamaño = len(newArr)-2
    final = newArr
    newArr2 = newArr
    //fmt.Println(cont, newArr)
    //fmt.Println(arr,final)
  }
  return newArr2
}

func tiposVariables(arreglo []string) ([]string,[]string,[]string,[]string){
  //m := make(map[string]string)
  // var caracteres []string
  var operadores []string
  var variables []string
  var numeros []string
  // var tipos []string
  var res []string
  for i:=0;i < len(arreglo);i++{
    elemento := arreglo[i]
    _,err := strconv.ParseFloat(elemento,64)
    if err != nil{
      if elemento=="+" || elemento=="-" || elemento=="*" || elemento=="/" || elemento==":=\n"{
        operadores = append(operadores,elemento)
        res = append(res,elemento+"-> OP")
        // tipos = append(tipos, "OP")
        // caracteres = append(caracteres, elemento)
      }else{
        variables = append(variables,elemento)
        res = append(res,elemento+"-> ID")
        // tipos = append(tipos, "ID")
        // caracteres = append(caracteres, elemento)
      }
    }else{
      //m[elemento]="VAL"
      numeros = append(numeros,elemento)
      res = append(res,elemento+"-> VAL")
      // tipos = append(tipos, "VAL")
      // caracteres = append(caracteres, elemento)
    }
  }
  /*for i:=0;i < len(tipos);i++{
    tipo:=tipos[i]
    carac:=caracteres[i]
    res = append(res,carac+"->"+tipo)
  }*/
//  fmt.Println(arr)
  return res, operadores, variables, numeros
}

func imprimirArr(arr []string){

    for i:=0; i < len(arr); i++{
      fmt.Println(arr[i])
    }
}

//Evalua caracteres y reconoce si hay decimales o caracteres invalidos
func errores(variables []string,numeros []string)([]string,bool){
  var mensajes []string
  var vars, _ = regexp.Compile("^[a-zA-Z_][a-zA-Z0-9_]*$")
  var nums,_ = regexp.Compile("^[0-9]*$")
  var flag = false
  for i:=0; i < len(variables); i++{
    elemento := variables[i]
    if !vars.MatchString(elemento) {
      mensajes = append(mensajes,"¡ERROR! El caracter "+elemento+" es invalido, las variables solo pueden ser letras.")
    }
  }
  for i:=0; i < len(numeros); i++{
    elemento := numeros[i]
    if !nums.MatchString(elemento) {
      mensajes = append(mensajes,"¡ERROR! El numero "+elemento+" es invalido, por ser decimal.")
    }
  }
  if len(mensajes)>0{
    flag=true
  }

  return mensajes,flag
}

func main() {
  mapa,cadena,tiposV,_,vars,nums := leerCadenas()
  //fmt.Println(mapa,cadena)
  newCadena := modificarArreglo(mapa,cadena)
  fmt.Println("Tabla caracteres")
  imprimirArr(tiposV)
  err,flag := errores(vars,nums)
  if !flag{
    arbolito := Armar_arbol(newCadena)
    if arbolito.Valor !="" {
      fmt.Println(newCadena)
      fmt.Print("Resultado ","\t")
      RecorrerInorden(arbolito)
      fmt.Println("=",Calcular(arbolito))
    }
  }else{
    imprimirArr(err)
  }
//  fmt.Println(tiposV)

}
