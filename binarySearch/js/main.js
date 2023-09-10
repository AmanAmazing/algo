function binarySearch(sortedArray,needle){
    let low = 0 
    let high = sortedArray.length 
    
    do{
        let midpoint = Math.floor(low+ (high-low)/2) //low is basically the offset 
        let value = sortedArray[midpoint]
        if (value === needle){
            return [true, midpoint]
        }else if (needle>value){
           low = midpoint +1 
        }else {
           high = midpoint
       }
    }while (low<high)
    return [false, -1] 
}


const array = [0,1,2,6,7,10,13,18,19,25,56,78]
const tests = [78,0,14,13,6]

for (i=0;i<tests.length;i++){
    console.log("Checking for:",tests[i])
    const [found, value] = binarySearch(array,tests[i])
    if (found ===true){
        console.log("Found value:",tests[i],"at index:",value)
    }else{
        console.log("value not found")
    }
}


