package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    existChar := make(map[rune] int)
    
    maxLen := 0
    l := 0
    currentIndex := 0
    for i, ch := range s {
       v, ok := existChar[ch] 
       if !ok || (ok && v==-1){
           l ++
           existChar[ch] = i
       }else{
           t := existChar[ch]
           l = i-t 
           for c := currentIndex;c < t;c++ {
                existChar[rune(s[c])] = -1
           }
           existChar[ch] = i
           currentIndex = t+1
       }

       if maxLen < l {
           maxLen = l
       }
    }
    return maxLen
}

func main(){
    s := "abcdec"
    fmt.Println(lengthOfLongestSubstring(s))
}
