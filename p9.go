package main


import "fmt"

func isPalindrome(x int) bool {   
    
    if x<10 && x>-10 {
        return true
    }

    n := x
    if x<0 {
        return false
    }

    nums := make([]byte,0)
    for {
        if n/10 == 0 {
            nums = append(nums, byte(n))
            break
        }else{
            nums = append(nums, byte(n%10))
            n /= 10
        }
    }
    i,j := 0,len(nums)-1 
    for i,j=0,len(nums)-1;nums[i]==nums[j] && i<=j;i,j=i+1,j-1 {
    }
    
    if i>j {
        return true
    }
    return false 

}

func main() {
    fmt.Println(isPalindrome(12321))
    fmt.Println(isPalindrome(1))
    fmt.Println(isPalindrome(11))
    fmt.Println(isPalindrome(112))
    fmt.Println(isPalindrome(-2147447412))
}
