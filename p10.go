package main

import "fmt"

func isMatch(s string, p string) bool {
    fmt.Println("111",s,"222",p)
    sLen := len(s)
    pLen := len(p)

    if s==p {
        return true
    }

    if sLen==0 {
        if pLen==0 {
            return true
        }else{
            if pLen==1 {
                return false
            }else{
                if p[1]=='*'{
                    return isMatch(s,p[2:])
                }else{
                    return false
                }
            }
        }
    }else {
        if pLen==0 {
            return false
        }
    }

    sIndex := 0
    pIndex:= 0

    const STAR = '*'
    const DOT = '.'

    for sIndex<sLen && pIndex<pLen{
        if p[pIndex]==DOT || p[pIndex]==s[sIndex] {
            if pIndex+1>=pLen {
                break
            }else{
                if p[pIndex+1] != STAR {
                    sTmp := string(s[sIndex+1:])
                    pTmp := string(p[pIndex+1:])
                    return isMatch(sTmp, pTmp)
                }else{
                    sTmp := string(s[sIndex+1:])
                    pTmp1 := string(p[pIndex+2:])
                    pTmp2 := string(p[pIndex:])
                    return isMatch(sTmp,pTmp1) || isMatch(sTmp,pTmp2) || isMatch(s[sIndex:],pTmp1)
                }
            }
        }else{
            if p[pIndex] == STAR {
                pIndex++
                continue
            }

            if pIndex+1>=pLen {
                return false
            }else{
                if p[pIndex+1] == STAR {
                    //sIndex++
                    pIndex += 2
                }else{
                    return false
                }
            }
        }
    }

    if (sIndex>=sLen && pIndex>=pLen) || (sIndex==sLen-1 && pIndex==pLen-1){
        return true
    }
    return false
}

func main() {
    //fmt.Println(isMatch("aa","aaa"))
    //fmt.Println(isMatch("ab",".*"))
    //fmt.Println(isMatch("ab","ab.*c"))
    //fmt.Println(isMatch("ab",".*c"))
    //fmt.Println(isMatch("aaa","a.a"))
    //fmt.Println(isMatch("aaa","ab*a"))
    //fmt.Println(isMatch("a","ab*"))
    //fmt.Println(isMatch("bbbba",".*a*a"))
    fmt.Println(isMatch("a",".*..a"))
}
