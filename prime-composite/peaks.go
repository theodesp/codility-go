package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// count the number of peaks, then go to each last index within each block, checking whether for each block, the number of peaks has gone up by at least 1. 
// If not, that block is missing a peak. Return the last divisor for which a peak was found in each block.

func Solution(A []int) int {
    peaks := 0
    var divisors []int 
    peakArray := []int{0}
    length := len(A)
    result := 0

    for i := 1; i < length -1; i++{
        if A[i-1] < A[i] && A[i] > A[i+1]{
            peaks++
        }
        peakArray = append(peakArray, peaks)
        if length % i == 0{
            divisors = append(divisors, i)
        }
    }
    peakArray = append(peakArray, peaks)

    for _, divisor := range divisors{
        block := length / divisor 
        lastPeak := 0
        ok := true
        
        for i := 1; i <= divisor; i++{
            if peakArray[i*block-1] > lastPeak{
                lastPeak = peakArray[i*block-1]
            } else{
                ok = false
                break
            }
        }
        if ok {
            result = divisor
        }
    }

    return result

}
