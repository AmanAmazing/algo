function linear_search(elementArray,needle){
    for (let i =0;i<elementArray.length;i++) {
        if (elementArray[i]== needle){
            return true  
        }else {
            continue
        }
    }
    return false
}


const a = [12,3,15,34,0,4]

console.log(linear_search(a,34)) // true
console.log(linear_search(a,21312)) // false
