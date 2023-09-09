

def linear_search(items: [int], needle: int) -> (bool,int):
    for index,value in enumerate(items):
        if value == needle:
            return (True,index)
        else:
            continue
    return (False,0)

if __name__=='__main__':
    test_array = [12,1,232,5434,56,69]
    tests = [1,43,772,56]
    for i in tests:
        result, index = linear_search(test_array, i)
        if result == True:
            print(f"The number: {i} can be found at index: {index}")
        else:
            print(f"The number: {i} could not be found in the array")
