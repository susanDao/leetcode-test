package main

import "fmt"

func reverse(x int) int {
      if x<10 && x>-10 {
          return x
      }

      const POSITIVE=1
      const NEGATIVE=2
      const MAX = 2<<31-1
      const MIN = -(2<<31)
      flag := 0
      
      v := x
      if x<0 {
          flag = NEGATIVE
          v = -v
      }
      rs := make([]int,0)
      for v > 0 {
          rs = append(rs, v%10) 
          v /= 10
      }
      var r int
      for i := 0;i<len(rs);i++ {
        r = r*10+rs[i]
      }
      if flag==NEGATIVE {
          r = -r
          if r<MIN {
              r = 0
          }
      }else{
          if r>MAX {
              r=0
          }
      }
      fmt.Println(MAX)
      fmt.Println(MIN)
      return r
}

func main(){
    fmt.Println(reverse(123))
    fmt.Println(reverse(-123))
}
