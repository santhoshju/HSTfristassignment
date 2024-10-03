package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter the Array Size: ")
    fmt.Scan(&n)

    if n < 2 {
        fmt.Print("Please Enter Atleast Two Elements")
        return
    }

    var array = make([]int, n)

    fmt.Println("Enter the elements of the array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&array[i])
    }

    var largest, secondlargest int
    largest = array[0]
    secondlargest = -1 

    for j := 1; j < n; j++ {
        if array[j] > largest {
            secondlargest = largest
            largest = array[j]
        } else if array[j] > secondlargest && array[j] != largest {
            secondlargest = array[j]
        }
    }

    if secondlargest == -1 {
        fmt.Println("There is no second largest element.")
    } else {
        fmt.Println("Largest Element:", largest)
        fmt.Println("Second Largest Element:", secondlargest)
    }
}
