package algo

func SelectionSort(arr []int) []int {
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

func InsertionSort(arr []int) []int {
    var len int = len(arr)
    for i := 1 ; i < len; i++ {
        for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
            var tmp = arr[j-1]
            arr[j-1] = arr[j]
            arr[j] = tmp
        }
    }
    return arr
}
