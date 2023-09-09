package main 

import(
    "fmt"
)


func linearSearch(array []int,needle int) (bool,int){
    for i:=0; i<len(array);i++{
        if array[i] == needle {
            return true,i 
        }else{
            continue
        }
    }
    return false, 0 
}

func main(){
    var testSlice = []int{12,3,2345,6,435,8};
    tests := []int{6,3,45023};
    for i:=0; i<len(tests);i++{
        result, index := linearSearch(testSlice,tests[i])
        if result == true{
            fmt.Println("The number:",tests[i]," can be found at index:",index)
        }else{
            fmt.Println("The number:",tests[i]," could not be found in the array")
        }
    }
}

