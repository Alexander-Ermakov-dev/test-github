package main

import ("fmt")

/*func main() {
    var x string
    x = "first"
    fmt.Println(x)
    x = "second"
    fmt.Println(x)
}*/
/*func main() {
    var x string
    x = "first "
    fmt.Println(x)
    x = x + "second"
    fmt.Println(x)
}*/

/*func main() {
  var x string = "hello"
  var y string = "hello"
  fmt.Println(x == y)
} */

/*var x string = "Hello World"

func main() {
    fmt.Println(x)

}

func f() {
    fmt.Println(x)
}
*/

/*func main() {
    const x string = "Hello World"
    fmt.Println(x)
}
*/


func main() {
    fmt.Print("Введите температуру по Фаренгейту : ")
    var F float64
    fmt.Scanf("%f", &F)

    C := (F - 32) * 5/9

    fmt.Println(C)
}
