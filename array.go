package main

import "fmt"

// func main() {
//     var x [5]int
//     x[4] = 100
//     fmt.Println(x)
// }

// func main() {
//     var x [5]float64
//     x[0] = 98
//     x[1] = 93
//     x[2] = 77
//     x[3] = 82
//     x[4] = 83
//
//     var total float64 = 0
//     for i := 0; i < 5; i++ {
//         total += x[i]
//     }
//      fmt.Println(total / 5)
//   }

// func main() {
//     slice1 := []int{1,2,3}
//     slice2 := make([]int, 2)
//     copy(slice2, slice1)
//     fmt.Println(slice1, slice2)
// }

// func main() {
//     elements := make(map[string]string)
//     elements["H"] = "Hydrogen"
//     elements["He"] = "Helium"
//     elements["Li"] = "Lithium"
//     elements["Be"] = "Beryllium"
//     elements["B"] = "Boron"
//     elements["C"] = "Carbon"
//     elements["N"] = "Nitrogen"
//     elements["O"] = "Oxygen"
//     elements["F"] = "Fluorine"
//     elements["Ne"] = "Neon"
//
//     fmt.Println(elements["Li"])
// }
func main() {
    elements := map[string]map[string]string{
            "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }

    if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }
}
