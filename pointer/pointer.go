package main

import (
    "fmt"
    "unsafe"
)

type mystruct struct {
    a, b int
}

func main() {
    a := uint8(5)
    p := &a
    fmt.Printf("Adress var: %p Value: %d\n", &a, a )
    fmt.Printf("Sizeof var of type %T: %d\n",a, unsafe.Sizeof(a))
    fmt.Printf("Adress Pointer: %p Value: %d\n", p, *p )
    fmt.Printf("Sizeof pointer to a of type %T : %d\n",p, unsafe.Sizeof(p))
    pp := &p
    fmt.Printf("Adress Pointer to Pointer: %p Value: %p\n", pp, *pp )
    fmt.Printf("Sizeof Pointer to Pointer of type %T: %d\n",pp, unsafe.Sizeof(pp))    
    b := mystruct{1, 2}
    fmt.Printf("Adress Struct %p Value: %d\n", &b, b.a)
    fmt.Printf("Adress Struct %p Value: %d\n", &b, b)
    fmt.Printf("Sizeof struct of type %T: %d\n",b, unsafe.Sizeof(b))
    p1 := &b
    fmt.Printf("Adress Pointer to Struct: %p Value: %d\n", p1, *p1 )
    fmt.Printf("Sizeof Pointer to Struct of type %T: %d\n",p1, unsafe.Sizeof(p1))
    modifyStruct(b)
    fmt.Printf("Value b: %d\n", b )
    modifyStructPointer(&b)
    fmt.Printf("Value b: %d\n", b ) 
    c := [3]int{1, 2, 3}
    fmt.Printf("Sizeof array of type %T: %d\n",c, unsafe.Sizeof(c))
    fmt.Printf("Adress Pointer to Array: %p Value: %d\n", &c, c)
    fmt.Printf("Length array: %d \n", len(c) )
    d := []int{1, 2, 3 ,4}
    fmt.Printf("Sizeof slice of type %T: %d\n",d, unsafe.Sizeof(d))
    fmt.Printf("Adress Pointer to slice: %p Value: %d\n", &d, d)
    modifySlice(d)
    fmt.Printf("Adress Pointer to slice: %p Value: %d\n", &d, d)
    e := "String"
    fmt.Printf("Adress string: %p Value: %s\n", &e, e )
    fmt.Printf("Sizeof string of type %T: %d\n",e, unsafe.Sizeof(e))
    fmt.Printf("Length string: %d \n", len(e) )
}

func modifyStruct(s mystruct) {
    s.a = 10
}


func modifySlice(s []int) {
    s[0] = 10
}


func modifyStructPointer(s *mystruct) {
    (*s).a = 10
}


    // var t [5]float64
    // fmt.Println(t)
    // m := new(bet)
    // fmt.Printf("%p\n", m)
    // fmt.Println( m)
    // *m = bet{"Aces", []int{3, 3, 3} , 61}
    // m.name = "spades"
    // fmt.Println( m)
    // fmt.Printf("%p\n", m)
    // aces := bet{"Aces", []int{3, 3, 3} , 61}
    // fmt.Printf("%p\n", &aces)
    // peq := &aces
    // *peq = bet{"Pequeno", []int{5,6,7,8}, 1}
    // fmt.Println(peq)
    // mega := bet{"mega",[]int{5,6,7,8}, 1 }
    // fmt.Println(mega)


// func pointer() {
    
// }
  /*  pequeno := bet{"Pequeno", []int{5,6,7,8}, 1}   */
    
/*    a := &aces    
    fmt.Println(a)
    b := test(&aces)
    fmt.Println(b)
    fmt.Println(aces)*/
/*    
    sl := []bet{aces, pequeno}
/*    s2 := &sl*/
/*    fmt.Println(sl)

    test2(sl)
    fmt.Println(sl) 
    
    test3(&sl[0])
    fmt.Println(sl)*/





/*func test(tmp *bet) *bet {
    tmp.name = "hallo"
    fmt.Println(tmp)
    return tmp
}*/

/*func test2(tmp []bet)  {
    tmp[0].name = "hallo"
   /* fmt.Println(tmp)

}*/

/*func test3(tmp *bet)  {
    tmp.name = "hallo2"
    fmt.Println(tmp)

}*/