package main

import "fmt"

var rti = map[uint8]int{
    73 : 1, //"I"
    86 : 5, //"V"
    88: 10, //"X"
    76 : 50, //"L"
    67 : 100, //"C"
    68 : 500, //"D"
    77 : 1000, //"M"
}

func main() {
    for k,_ := range rti {
       fmt.Printf("%v = %v\n", string(k), k)
    }
}

func romeToInt2(roma string) int {
    curIndex := 0
    nextIndex := 0
    result := 0
    for curIndex < len(roma) {
        curNumber := rti[roma[curIndex]]
        nextIndex = curIndex + 1
        if nextIndex > len(roma) - 1 {
            result = result + curNumber
            return result
        }
        nextNumber :=  rti[roma[nextIndex]]
        if curNumber < nextNumber {
            result = result + (nextNumber - curNumber)
            curIndex= curIndex + 2
        } else {
            result = result + curNumber
            curIndex++
        }
    }
    return result;
}

func romeToInt(roma string) int {
    lastSymbol := roma[len(roma) - 1]
    lastValue := rti[lastSymbol]
    result := lastValue
    for i := len(roma) - 2; i >= 0; i-- {
        currenSymbol := roma[i]
        currentValue := rti[currenSymbol]
        if currentValue < lastValue {
            result = result - currentValue
        } else {
            result = result + currentValue
        }
        lastValue = currentValue
    }

    return result;
}
