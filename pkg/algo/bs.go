package algo

func BinarySearch(arr []int, key int) int {
    var low int = 0
    var high int = len(arr)
    for low < high {
        var mid int = low + (high - low) /2
        if key < arr[mid] {
            high = mid
        } else if key > arr[mid] {
            low = mid + 1;
        } else {
            return mid;
        }
    }
    return -(low + 1)
}
