package main

import "fmt"
import "unsafe"

func Add(i, j int) int {
    fmt.Printf("%d + %d =%d\n", i, j, i+j)
    return i+j
}

func Max(i, j int) int {
// input param:
// i int, j int
    if i > j {
        return i
    }else{
        return j
    }

}

func Split(s string) (string, int){
    //return s, unsafe.Sizeof(s)
    return s, 10

}

type Books struct {
    title  string
    author string
    subject string
    book_id int
}


func main() {
    fmt.Printf("hello again!\n")
    fmt.Println("hello again!")

    /* assign new variable .. */
    var a int
    var b = 90.2
    //b := -9
    c :=  -9
    d := 10 // no use varibale 
    _ = d
    //fmt.Printf(a, b, c)
    println(a, b, c)
    // const

    const LENGTH, WIDTH = 90, 100
    println()
    println(LENGTH, WIDTH)
    // const & iota
    const ( 
        A1 = iota
        A2 = iota
        A3 = iota
    )
    println("Const with iota")
    println(A1, A2, A3)
    // unsafe 
    println("unsafe demo")
    println("unsafe.Sizeof(A1)==", unsafe.Sizeof(A1))
    // if --else
    if  b > A1{

        println("b > A1")
    }else{
        println("b < A1")
    }
     const MAX_LOOP = 10
     var index = 0
    // for loop
    // go to 

    for true{
        println(" hello loop")
        if index > MAX_LOOP{
            break
        }
        index ++
    }
    index2 := 0
    GOTO_HERE:
        println(" loop by goto")
        if index2 > MAX_LOOP{
                goto END_LOOP
            }
        index2 ++
        goto GOTO_HERE
        END_LOOP:
    // test function declare ..
    println("Max value : Max(9, 19)=", Max(9, 19))

    // test multi return value ..
    var str string
    var str_len int
    str, str_len = Split("hello world")
    println(str, str_len)
    // array for go
    var n [10]int
    the_index := 0
    for the_index < 10{
        n[the_index] = the_index * the_index 
        the_index ++
    }
    the_index = 0 //reset flag
    for the_index < 10{
        println("the i =",the_index, "value is==>", n[the_index])
        the_index ++
    }
    // point for go
    var  int_ptr *int
    var int_var = 90
    println("the address of int_var =", &int_var)
    int_ptr = &int_var
    println("the value from ptr store=", int_ptr)

    // structure for go
    var book Books;
    book.title = "Golang"
    book.author = "Frank"
    book.subject = "hello golang"
    book.book_id = 101

    // read data from structure  ..
    println("book.book_id=", book.book_id)

    // structure with point
    var book_ptr *Books;
    book_ptr = & book
    println("the size of book=", unsafe.Sizeof(book), "address of book=", &book, "ptr=", book_ptr)
    // slice
    s:=make([] int, 3, 5)
    println("Hello slice ==>",s)
    numbers := []int {0, 1, 2, 3, 4, 5, 6}
    fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
    numbers = append(numbers, 99, 88,77)
    fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
    // range for go
    for _, num  := range numbers{
        println(" number -->", num)
    }

    // iter a map
    kvs := map[string] string{"a":"hello a", "b":"hello b"}
    for k, v := range kvs{
        fmt.Printf("key=%s --> value=%s\n", k, v)
    }

    // iter a string
    for _, ss := range "go"{
        println("ss ==>", ss)
    }
    // declare a map 
    var name2nickname map[string]string

    // create a init data
    name2nickname = make(map[string]string)

    name2nickname["frank"] = "landpack"
    name2nickname["jack"] = "luss Jack"
    for k, v := range name2nickname{
        fmt.Printf("name=%s -- nickname=%s\n", k, v)
        }

    // delete a jack
    delete(name2nickname, "jack")
    for k, v := range name2nickname{
        fmt.Printf("name=%s -- nickname=%s\n", k, v)
        }
}
