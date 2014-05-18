package algo

import "testing"

func TestBinarySearch(t *testing.T) {
    arr := [] int {1, 2, 3, 4, 5, 6}

    VerifySearch(arr, 1, 0, t)
    VerifySearch(arr, 2, 1, t)
    VerifySearch(arr, 3, 2, t)
    VerifySearch(arr, 4, 3, t)
    VerifySearch(arr, 5, 4, t)
    VerifySearch(arr, 6, 5, t)
    VerifySearch(arr, 7, -7, t)
    VerifySearch(arr, 8, -7, t)

}

func VerifySearch(arr []int, key int, expected int, t *testing.T) {
    var idx int = BinarySearch(arr, key)
    if idx != expected {
        t.Error("For", key, "Expected", expected, "got", idx)
    }

}
