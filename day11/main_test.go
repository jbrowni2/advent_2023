package main

import (
	"bufio"
	l "day11/utilities/mat"
	s "day11/utilities/state"
	"log"
	"os"
    "testing"
)



func TestMain1(t *testing.T){
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
    galaxies := s.State{}

    init := s.Corridinent{Row: 6, Col: 1}
    
    galaxies.Constructor(init, &data)


    num := galaxies.AStar(s.Corridinent{Row: 11, Col: 5})


    if num != 9{
        t.Fatalf("Expected 9 got %v\n",num)
    }
    

}


func TestMain2(t *testing.T){
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

    galaxies := s.State{}

    init := s.Corridinent{Row: 2, Col: 0}
    
    galaxies.Constructor(init, &data)


    num := galaxies.AStar(s.Corridinent{Row: 7, Col: 12})

    if num != 17{
        t.Fatalf("Expected 9 got %v\n",num)
    }
    

}



func TestMain3(t *testing.T){
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

    galaxies := s.State{}

    init := s.Corridinent{Row: 0, Col: 4}
    
    galaxies.Constructor(init, &data)


    num := galaxies.AStar(s.Corridinent{Row: 10, Col: 9})

    if num != 15{
        t.Fatalf("Expected 9 got %v\n",num)
    }
    

}



func TestMain4(t *testing.T){
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

    galaxies := s.State{}

    init := s.Corridinent{Row: 11, Col: 0}
    
    galaxies.Constructor(init, &data)


    num := galaxies.AStar(s.Corridinent{Row: 11, Col: 5})

    if num != 5{
        t.Fatalf("Expected 9 got %v\n",num)
    }
    

}
