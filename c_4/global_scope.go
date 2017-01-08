package main

var a string

func main(){
    a = "G"
    print(a)
    f1()
}

func f1(){
    a := "O"
    print(a)
    a = "W"
    f2()
}

func f2(){
    print(a)
}
