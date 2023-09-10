package main

import (
	"fmt"
) 


func binarySearch(sortedArray []int,needle int) (bool,int){
    low := 0 
    high := len(sortedArray)
    for {
        midpoint := low + (high-low)/2
        value := sortedArray[midpoint]
        if value == needle{
            return true, midpoint
        }else if value<needle{
            low = midpoint + 1
        }else {
            high = midpoint
        }
        if low >=high {
            break
        }
    }
    return false, -1
}


func main() {
	data := []int{0, 1, 2, 6, 7, 10, 13, 18, 19, 25, 56, 78}
	testData := []int{78, 0, 14, 13, 6}

	for i := 0; i < len(testData); i++ {
        found,index:= binarySearch(data, testData[i])
        if found{
            fmt.Println("Found value:",testData[i],"at index:",index)
        }else{
            fmt.Println("Could not find value:",testData[i])
        }
	}
}
