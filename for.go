package main

import "fmt"

/*func main() {
    i := 1              // перечисляет числа с 1 до 10
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }
}*/

/* func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
} */

/*func main() {
    for i := 1; i <= 10; i++ {
       if i % 2 == 0 {
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }
} */

/* func main() {
      for i := 1; i <= 10; i++ {
      switch i {
        case 0: fmt.Println("Zero")
  case 1: fmt.Println("One")
  case 2: fmt.Println("Two")
  case 3: fmt.Println("Three")
  case 4: fmt.Println("Four")
  case 5: fmt.Println("Five")
  default: fmt.Println("Unknown Number")
    }
  }
} */

func main() {
      for i := 1; i <= 101; i++ { // выводит числа делющиеся на 3
        if i % 3 == 0 {
          fmt.Println(i)
        }
      }
    }
