package prime_composite

func Flags(a []int) int {
    var numPicks int
    var sumDist int
    var distances = make([]int, len(a) / 2)
    var lastPeakI int
    
    // Count peaks and distances
    for i:=1; i < len(a) - 1; i++ {
        if a[i] > a[i-1] && a[i] > a[i+1] {
            if lastPeakI != 0 {
                distances[numPicks] = i - lastPeakI
                sumDist += i - lastPeakI
            }
            numPicks++
            lastPeakI = i
        }
    }
    
    // Search maximum reasonable number of flags
    for i := numPicks ; i > 0 ; i-- {
        if i * (i - 1) <= sumDist {
            // Test found number of flags
            flags := 1
            prevSum := 0
            for j:=1;j < len(distances);j++ {
                if prevSum + distances[j] >= i {
                    flags++
                    prevSum = 0
                } else {
                    prevSum += distances[j]
                }
            }
            
            if flags >= i {
                return i
            }
        } 
    }
    
    return 0
}
