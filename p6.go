package main


import "fmt"

func convert(s string, numRows int) string {    
    if numRows==1 {
        return s
    }else if(numRows==0) {
        return ""
    }

    rs := make(map[int] []byte, numRows)
    for i:=0;i<numRows;i++ {
        rs[i] = make([]byte, len(s))
    }
    
    const UP = 1
    const DOWN = 2
    flag := 2

    for j,row,line:=0,0,0;j<len(s);j++ {
       if flag==DOWN {
           rs[row][line] = s[j]
           if row==(numRows-1) {
               row--
               line++
               flag = UP 
           }else{
               row++
           }
       }else if flag==UP {
           rs[row][line] = s[j]
           if row==0 {
               line++
               row++
               flag = DOWN
           }else {
               row--
               line++
           }
       }
    }
    result := make([]byte, len(s))
    index := 0

    for i:=0;i<numRows;i++ {
        for _,c := range rs[i] {
            if c!=0 {
                result[index] = c
                index++
            }
        }
    }

    return string(result)
}

func main(){
    s := "PAYPALISHIRING"
    fmt.Println(convert(s,3))
    fmt.Println(convert(s,0))
    fmt.Println(convert(s,1))
    fmt.Println(convert(s,2))
}
