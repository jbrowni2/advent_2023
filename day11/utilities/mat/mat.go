package mat

import "fmt"

type Node struct{
    info string
    next *Node
}

type List struct{
    head *Node
    tail *Node
    Length int
}

type ListNode struct{
    info *List
    next *ListNode
}

type Mat struct{
    head *ListNode
    tail *ListNode
    Length int
}


func (l *List) Insert(d string){
    lis := &Node{info: d, next: nil}
    if l.head == nil{
        l.head = lis
        l.tail = lis
        l.Length = 1
    } else {
        p := l.tail
        p.next = lis
        l.tail = lis
        l.Length += 1
    }
}

func (l *List) PrintList(){
    p := l.head
    for p != nil {
        fmt.Printf("%s",p.info)
        p = p.next
    }
}

func (l *List) GetIdx(col int) string{
    p := l.head
    for i:=0; i<col;i++{
        p = p.next
    }

    return p.info
}

func (l *List) CheckStar() bool{
    p := l.head
    for p != nil{
        if p.info == "#"{
            return false
        }
        p = p.next
    }
    return true
}

func (l *List) InsertIdx(col int, d string){
    lis := &Node{info: d, next: nil}
    p := l.head
    for j:=0; j<col; j++{
       p = p.next 
    }

    if p.next != nil {
        lis.next = p.next
        p.next = lis
    } else {
        p.next = lis
        l.tail = lis
    }
    
    l.Length = l.Length + 1
    
}

func (l *List) Clone() List{
    p := l.head
    lis := List{}
    for p != nil{
        lis.Insert(p.info)
        p = p.next
    }

    return lis
}


//--------------------------------------------------
//       Matrix Implementation
//--------------------------------------------------

func (l *Mat) Insert(d *List){
    lis := &ListNode{info: d, next: nil}
    if l.head == nil{
        l.head = lis
        l.tail = lis
        l.Length = 1

        if lis.info.CheckStar(){
            d2 := lis.info.Clone()
            lis2 := &ListNode{info: &d2, next: nil}
            p := l.tail
            p.next = lis2
            l.tail = lis2
            l.Length += 1
        }
    } else {
        p := l.tail
        p.next = lis
        l.tail = lis
        l.Length += 1

        if lis.info.CheckStar(){
            for i:=0;i<1000000-1; i++{
                fmt.Println(i)
                p = p.next
                d2 := lis.info.Clone()
                lis2 := &ListNode{info: &d2, next: nil}
                p.next = lis2
                l.tail = lis2
                l.Length += 1
            }
        }
    }
}

func (l *Mat) PrintMat(){
    lis := l.head
    for lis != nil{
        lis.info.PrintList()
        fmt.Printf("\n")
        lis = lis.next
    }
}

func (l *Mat) GetIdx(row int, col int) string{
    lis := l.head
    for i:=0;i<row;i++{
        lis = lis.next
    }
    st := lis.info.GetIdx(col)

    return st
}

func (l *Mat) Shape() (int, int){
    lis := l.head
    return l.Length, lis.info.Length
}

func (l *Mat) InsertIdx(row int, col int, d string){
    lis := l.head
    for i:=0;i<row;i++{
        lis = lis.next
    }

    lis.info.InsertIdx(col, d)
}

func (l *Mat) CheckStarCol(){
    var cols []int
    for col:=0; col<l.head.info.Length; col++{
        stars := true
        for row:=0; row<l.Length; row++{
            if l.GetIdx(row, col) == "#"{
                stars = false
            }
        }

        if stars == true{
            cols = append(cols, col)
        }
    }

    i := 0
    for _, c := range cols{
    for j:=0; j<1000000-1; j++{
        col := c + i
        for row:=0; row<l.Length; row++{
            l.InsertIdx(row, col, ".")
        }
        i += 1
        }
    }
}


func (m *Mat) GetGals() ([]int, []int){
    var gal_rows []int
    var gal_cols []int
    lis := m.head
    for row:=0; row<m.Length; row++{
        p := lis.info.head
        for col:=0; col < lis.info.Length; col++{
            if p.info == "#"{
                gal_rows = append(gal_rows, row)
                gal_cols = append(gal_cols, col)
            }
            p = p.next
        }

        lis = lis.next
    }

    return gal_rows, gal_cols
}
