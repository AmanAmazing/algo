import math

def binarySearch(sortedArray: [int],needle: int):
    low = 0 
    high = len(sortedArray)

    while (low<high):
       midpoint = math.floor(low + (high-low)/2) 
       value = sortedArray[midpoint]
       if value == needle:
           return (True,midpoint)
       elif value<needle:
            low = midpoint + 1
       else:
            high = midpoint
    return (False,-1)

if __name__ == '__main__':
    data = [0, 1, 2, 6, 7, 10, 13, 18, 19, 25, 56, 78]
    testData = [78,0,14,13,6]

    for i in testData:
        found, index = binarySearch(data,i)
        if found:
            print(f"The number: {i} can be found at index: {index}")
        else:
            print(f"The number: {i} could not be found in the array")
