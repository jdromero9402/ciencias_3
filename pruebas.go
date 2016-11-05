package main

import (
  "fmt"
)

//devolver un mapa
func devolver() map[string][]string{
  m := make(map[string][]string)
  a := []string{"a", "b", "1", "c"}
  m["x"] = a
  return m
}

//devolver un arreglo de mapas
func arreglo() []map[string][]string{
    m1 := make(map[string][]string)
    m2 := make(map[string][]string)
    m3 := make(map[string][]string)
    a := []string{"a", "b", "1", "c"}
    m1["x"] = a
    m2["y"] = a
    m3["z"] = a
    arr := []map[string][]string{m1, m2, m3}
    return arr

//modificar un arreglo
func modificarArreglo(arr []string, pos int) []string{
    var newArr = []string
    newArr[0] = "a"
    newArr[1] = "b"
    tam := len(arr) + len(newarr) - 1
    for i := pos; i < tam ; i++{
        arr[i+len(newArr)]
    }
  }
}

func main() {
  //m := devolver()
  arr := arreglo()
  for i:=0;i<len(arr);i++{
    fmt.Println(arr[i])
  }
  //fmt.Println(m)
}
