package main

import (
	"bufio"
	l "day11/utilities/mat"
	s "day11/utilities/state"
	"fmt"
	"log"
	"os"
)





func main(){
    f, err := os.Open("inputTest.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        for _, c := range line{
            lis.Insert(string(c))
        }
        data.Insert(&lis)
    }

    data.CheckStarCol()


    gals_rows, gals_cols := data.GetGals()

    var sum int64 = 0

    p := 0
    for i:=0; i<len(gals_rows); i++{
        prog := float64(p)/float64(len(gals_rows))
        fmt.Printf("Progress: %v\n", prog)
        galaxies := s.State{}

        init := s.Corridinent{Row: gals_rows[i], Col: gals_cols[i]}
    
        galaxies.Constructor(init, &data)



        for j:=i+1; j<len(gals_rows); j++{
            sum += galaxies.AStar(s.Corridinent{Row: gals_rows[j], Col: gals_cols[j]})
            galaxies.Reset()

            
        }

        p += 1
    }

    fmt.Printf("%v\n",sum)
}
