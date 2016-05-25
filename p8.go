package main


import "fmt"


func myAtoi(str string) int {

    if len(str)==0 {
        return 0
    }

    numIndexStart := 0
    numIndexEnd := 0
    flag := 1 

    //trim space
    for i:=0;i<len(str) && str[i]==' ';i++ {
            numIndexStart ++
    }

    //deal with '-' at first
    if numIndexStart>(len(str)-1) { // no nums
        return 0
    }else{
       if str[numIndexStart] == '-' {
           flag = -1
           numIndexStart++
       }else if str[numIndexStart] == '+' {
           numIndexStart++
       }
    }

    numIndexEnd = numIndexStart
    for j:=numIndexEnd;j<len(str) && str[j]>47 && str[j]<58;j++ {
        numIndexEnd++
    }

    nums := []byte(str[numIndexStart:numIndexEnd])

    //generate numbers,positve
    //overflow
    const INT_MAX = 2<<30-1
    const INT_MIN = -2<<30
    
    var r int

    for i:=0;i<len(nums);i++ {
        n := int(nums[i]- 48) // 48 is '0'
        if flag==1 {
            if (INT_MAX-n)/10 < r {
                return INT_MAX
            }
        }else{
            if (INT_MAX-n+1)/10 < r {
                return INT_MIN
            }
        }

        r = r*10+n
    }

    return r
}


func main(){
   //fmt.Println(myAtoi("  -123a"))
   //fmt.Println(myAtoi("2147483647"))
   //fmt.Println(myAtoi("2147483648"))
   //fmt.Println(myAtoi("-2147483648"))
   //fmt.Println(myAtoi("-8147483648"))
   fmt.Println(myAtoi("+1"))
}
