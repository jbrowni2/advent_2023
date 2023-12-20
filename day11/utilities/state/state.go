package state

import (
	m "day11/utilities/mat"
	"fmt"
	"math"
)



type Corridinent struct{
    Row int
    Col int
}

type State struct{
    visited [][]bool
    distance [][]int64
    rows int
    cols int
    init Corridinent
}


func (s *State) Constructor(c Corridinent, m *m.Mat){
    rows,cols := m.Shape()
    s.rows = rows
    s.cols = cols
    for i:=0; i<s.rows; i++{
        vis_raw := make([]bool, s.cols)
        dis_row := make([]int64, s.cols)
        for j:=0; j<s.cols; j++{
            dis_row[j] = 9223372036854775807 // represents infinity
        }

        s.visited = append(s.visited, vis_raw)
        s.distance = append(s.distance, dis_row)
    }

    s.distance[c.Row][c.Col] = 0 //starting position
    s.init.Col = c.Col
    s.init.Row = c.Row
}

func (s *State) Reset(){
    for row:=0;row<s.rows;row++{
        for col:=0;col<s.cols;col++{
            s.visited[row][col] = false
            s.distance[row][col] =  900000000000 //
        }
    }

    s.distance[s.init.Row][s.init.Col]=0
}

func (s *State) NonVisited() bool{
    for _, row := range s.visited{
        for _, v := range row {
            if v == false{
                return true
            }
        }
    }

    return false
}

func (s *State) Dijkstra(){
    current := s.init
    calculating := true
    

    for calculating{
        //check states around calculating
        row := current.Row
        col := current.Col

        next := Corridinent{Row: row, Col: col}
        var next_d int64 = 2000000

        var new_d int64


        if (row >= s.rows) || (col >= s.cols){
            fmt.Printf("%v,%v\n", row, col)
        } else {
            new_d = s.distance[row][col] + 1
        }
        for i:=row-1;i<=row+1;i++{
                if (i >= 0) && (i < s.rows){
                    if s.visited[i][col] == false {
                        if s.distance[i][col] > new_d{
                            s.distance[i][col] = new_d
                        }
                    
                   
                    }
                    
                }
        }

        for j:=col-1;j<=col+1;j++{
            if(j >= 0) && (j < s.cols){
                if s.visited[row][j] == false {
                    if s.distance[row][j] > new_d{
                        s.distance[row][j] = new_d
                    }
                }
                    
            }
        }
        s.visited[row][col] = true

        for i:=0;i<s.rows;i++{
            for j:=0; j<s.cols;j++{
                if s.visited[i][j] == false{
                    if next_d >= s.distance[i][j]{
                        next_d = s.distance[i][j]
                        next.Row = i
                        next.Col = j
                    }
                }
            }
        }

        current = next

        calculating = s.NonVisited()


    }
}

func (s *State) PrintState(){
    for row:=0; row<s.rows; row++{
        for col:=0; col<s.cols; col++{
            fmt.Printf("%v ", s.distance[row][col])
        }
        fmt.Printf("\n")
    }
}

func (s *State) GetIdx(row int, col int) int64{
    return s.distance[row][col]
}


func (s *State) PrintVisited(){
    for row:=0; row<s.rows; row++{
        for col:=0; col<s.cols; col++{
            fmt.Printf("%v ", s.visited[row][col])
        }
        fmt.Printf("\n")
    }
}

func (s *State) heuristic(current Corridinent, goal Corridinent) float64{
    gn := float64(s.distance[current.Row][current.Col])
    hn := math.Pow(float64(current.Row)-float64(goal.Row), 2.0) + math.Pow(float64(current.Col)-float64(goal.Col), 2.0)
    fn := gn + hn

    return fn
}

func (s *State) AStar(goal Corridinent) int64{
    current := s.init
    calculating := true
    

    for calculating{
        //check states around calculating
        row := current.Row
        col := current.Col

        next := Corridinent{Row: row, Col: col}

        var new_d int64


        if (row >= s.rows) || (col >= s.cols){
            fmt.Printf("%v,%v\n", row, col)
        } else {
            new_d = s.distance[row][col] + 1
        }
        for i:=row-1;i<=row+1;i++{
                if (i >= 0) && (i < s.rows){
                    if s.visited[i][col] == false {
                        if s.distance[i][col] > new_d{
                            s.distance[i][col] = new_d
                        }
                    
                   
                    }
                    
                }
        }

        for j:=col-1;j<=col+1;j++{
            if(j >= 0) && (j < s.cols){
                if s.visited[row][j] == false {
                    if s.distance[row][j] > new_d{
                        s.distance[row][j] = new_d
                    }
                }
                    
            }
        }

        s.visited[row][col] = true
        next_d := math.Inf(1)

        for i:=row-1;i<=row+1;i++{
            for j:=col-1; j<=col+1;j++{
                if (i>=0) && (i <s.rows) && (j>=0) && (j<s.cols){
                    if s.visited[i][j] == false{
                        fn := s.heuristic(Corridinent{Row: i, Col: j}, goal)
                        if next_d > fn{
                            next_d = fn
                            next.Row = i
                            next.Col = j
                        }
                    }
                }
            }
        }

        current = next
        if (current.Row == goal.Row) && (current.Col == goal.Col){
            return s.distance[current.Row][current.Col]
        }

        calculating = s.NonVisited()

    }

    return -10
}

func (s *State) PrintH(goal Corridinent){
    for row:=0; row<s.rows; row++{
        for col:=0; col<s.cols; col++{
            fn := s.heuristic(Corridinent{Row: row, Col: col}, goal)
            fmt.Printf("%v: %v   ", s.distance[row][col], fn)
        }
        fmt.Printf("\n")
    }
}
