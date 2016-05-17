package main

import "fmt"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var m, n int
    r := 0.0
    m = len(nums1)
    n = len(nums2)

    if m==0 && n!=0 {
        if n%2!=0 {
            r = float64(nums2[(n-1)/2])
        }else{
            r = (float64(nums2[n/2])+float64(nums2[n/2-1]))/2
        }
    }else if m!=0 && n==0 {
        if m%2!=0 {
            r = float64(nums1[(m-1)/2])
        }else{
            r = (float64(nums1[m/2])+float64(nums1[m/2-1]))/2
        }

    }else {
        ma := 0
        mb := 0
        flag := 0
        
        k := (m+n+1)/2
        cc := 0
        for {
            if k==1 {
                if (m+n)%2!=0 {
                    if mb>(n-1) {
                        r = float64(nums1[ma])
                        break
                    }
                    if ma>(m-1){
                        r = float64(nums2[mb])
                        break
                    }
                    if nums1[ma]<nums2[mb] {
                        r = float64(nums1[ma])
                        break
                    }else{
                        r = float64(nums2[mb])
                        break
                    }
                }else {
                    if mb>n-1 {
                        r = (float64(nums1[ma])+float64(nums1[ma+1]))/2
                        break
                    }
                    if ma>m-1 {
                        r = (float64(nums2[mb])+float64(nums2[mb+1]))/2
                        break
                    }
                    if nums1[ma]<nums2[mb] {
                        if (ma+1)<m && nums1[ma+1]<nums2[mb] {
                            r = (float64(nums1[ma])+float64(nums1[ma+1]))/2
                            break
                        }
                    }
                    if nums1[ma]>nums2[mb] {
                        if (mb+1)<n && nums2[mb+1]<nums1[ma] {
                            r = (float64(nums2[mb])+float64(nums2[mb+1]))/2
                            break
                        }
                    }

                    r = (float64(nums2[mb])+float64(nums1[ma]))/2
                    break
                }
            }

            t := k/2
            A := ma+t-1
            B := mb+t-1
            if A>(m-1) {
                if flag==0 {
                    cc = 1
                    A = m-1
                }else{
                    if (m+n)%2!=0 {
                        r = float64(nums2[mb+k-1])
                        break
                    }else{
                        r = (float64(nums2[mb+k-1])+float64(nums2[mb+k]))/2
                        break
                    }
                }
            }
            if B>(n-1) {
                if flag==0 {
                    B = n-1
                    cc = 2 
                }else{
                    if (m+n)%2!=0 {
                        r = float64(nums1[ma+k-1])
                        break
                    }else{
                        r = (float64(nums1[ma+k-1])+float64(nums1[ma+k]))/2
                        break
                    }
                }
            }
            if nums1[A] == nums2[B] {
                if (m+n)%2!=0 && k==2 {
                    r = float64(nums1[A])
                    break
                }else{
                    if (B+1)>(n-1) || nums1[A+1] < nums2[B+1] {
                        ma = A+1 
                        k = k-t
                        continue
                    }else{
                        mb = B+1
                        k = k-t
                        continue
                    }
                }
            }else if nums1[A] < nums2[B] {
                ma = A+1
                if cc == 1 {
                    k = k-m
                }else {
                    k = k-t
                }
            }else{
                mb = B+1
                if cc == 2 {
                    k = k-n
                }else {
                    k = k-t
                }

            }
            flag = 1
            cc = 0
        }
    }
    return r
}

func main(){
    nums1 := []int{3}
    nums2 := []int{1,2,4,5,6,7}
    fmt.Println(findMedianSortedArrays(nums1, nums2))
}
