package main

import "fmt"

func main() {
    mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
    mm2["scala"] = 25
    mm2["erlang"] = 50
    mm2["python"] = 0
    fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])
}