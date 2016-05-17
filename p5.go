package main

import "fmt"

func longestPalindrome(s string) string {
    var i,j,k int
    
    l := len(s)

    var start,end int
    start = 0
    end = 0
    
    if l==0 {
        return ""
    }else if l==1 {
        return s
    }else{
        n := (2*l+1)
        str := make([]byte,n) 
        for a,b := 0,0;b<n;b++ {
            if (b+1)%2!=0 {
                str[b] = '|'
            }else{
                str[b] = s[a]
                a++
            }
        }

        for k=1;k<=(n-1);k++ {
            i = k-1
            j = k+1
            if j>(n-1) {
                break
            }

            if str[i] != str[j] {
                continue
            }
            
            for {
                if str[i] != str[j] {
                    i = i+1
                    j = j-1
                    break
                }
                if (i-1)<0 || (j+1)>(n-1){
                    break
                }
                i,j = i-1,j+1
            }
            
            if (j-i) > (end-start) {
                start = i
                end = j
            }
        }

        subStr := str[start:end+1]
        rs := make([]byte,0)

        for _,v := range subStr {
            if v!='|' {
                rs = append(rs, v)
            }
        }

        return string(rs) 
    }
}

func main() {
    s := "bb"
    fmt.Println(longestPalindrome(s))
}
