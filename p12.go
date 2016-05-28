package main

import "fmt"
import "strings"

func intToRoman(num int) string {    
    romanNums := [4][]string{
        {"","I","II","III","IV","V","VI","VII","VIII","IX"},
        {"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"},
        {"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"},
        {"","M","MM","MMM"}}
    r := make([]string,4)
    r[0] = romanNums[3][num/1000]
    r[1] = romanNums[2][(num%1000)/100]
    r[2] = romanNums[1][(num%100)/10]
    r[3] = romanNums[0][(num%10)]
    return strings.Join(r,"")
}

func main() {
    fmt.Println(intToRoman(3))
    fmt.Println(intToRoman(11))
    fmt.Println(intToRoman(1111))
    fmt.Println(intToRoman(111))
}
