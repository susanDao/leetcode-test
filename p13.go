package main

import "fmt"

func romanToInt(s string) int {
    romanNums := map[string] int{
        "MMM":3000,"MM":2000,"M":1000,
        "CM":900,"DCCC":800,"DCC":700,"DC":600,"D":500,"CD":400,"CCC":300,"CC":200,"C":100,
        "XC":90,"LXXX":80,"LXX":70,"LX":60,"L":50,"XL":40,"XXX":30,"XX":20,"X":10,
        "IX":9,"VIII":8,"VII":7,"VI":6,"V":5,"IV":4,"III":3,"II":2,"I":1}
    var r,tmp int
    flag := 0
    for i:=0;i<len(s);i+=flag {
        flag = 0
        tmp = 0
        for j:=i;j<len(s);j++ {
            ss := string(s[i:(j+1)])
            v,ok := romanNums[ss]
            if !ok {
                break
            }else{
                tmp = v 
                flag++ 
            }
        }
        r += tmp
    }
    return r 
}

func main() {
    fmt.Println(romanToInt("MCM"))
    fmt.Println(romanToInt("MM"))
    fmt.Println(romanToInt("MMM"))
    fmt.Println(romanToInt("VIII"))
    fmt.Println(romanToInt("III"))
    fmt.Println(romanToInt("IX"))
}
