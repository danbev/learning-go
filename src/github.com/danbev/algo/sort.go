package algo

func InsertionSort(arr []int) []int {
    var len int = len(arr)
    for i := 0 ; i < len; i++ {
        var min = i
        for j := i + 1; j < len; j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }
        var tmp = arr[i]
        arr[i] = arr[min]
        arr[min] = tmp
    }
    return arr
}
