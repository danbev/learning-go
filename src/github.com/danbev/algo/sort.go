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

/*
var unsorted = [] int {2, 6, 4, 3, 5, 1}
var len int = 6
1. i = 1, j = 1, 6 < 2 == false, j-- == 0 break 
{2, 6, 4, 3, 5, 1}
2. i = 2, j = 2, 4 < 6 == true, swap 6 and 4 = {2, 4, 6, 3, 5, 1} , j-- == 1 continue 
    j = 1, 4 < 2 == false, j-- == 0 break
{2, 4, 6, 3, 5, 1}
3. i = 3, j = 3, 3 < 6 == true, swap 6 and 3 = {2, 4, 3, 6, 5, 1} , j-- == 2 continue 
    j = 2, 3 < 4 == true, swap 4 and 3 = {2, 3, 4, 6, 5, 1}, j-- == 1 continue
    j = 1, 3 < 2 == false, j-- == 0 break
{2, 3, 4, 6, 5, 1}
4. i = 4, j = 4, 5 < 6 == true, swap 6 and 5 = {2, 3, 4, 5, 6, 1} , j-- == 3 continue 
    j = 3, 5 < 4 == false, j-- == 2 continue
    j = 2, 4 < 3 == false, j-- == 1 continue
    j = 1, 3 < 2 == false, j-- == 0 break
{2, 3, 4, 5, 6, 1}
5. i = 5, j = 5, 1 < 6 == true, swap 6 and 1 = {2, 3, 4, 5, 1, 6} , j-- == 4 continue 
    j = 4, 1 < 5 == true, sway 6 and 1 = {2, 3, 4, 1, 5, 6}, j-- == 3 continue
    j = 3, 1 < 4 == true, sway 4 and 1 = {2, 3, 1, 4, 5, 6}, j-- == 2 continue
    j = 2, 1 < 3 == true, sway 3 and 1 = {2, 1, 3, 4, 5, 6}, j-- == 1 continue
    j = 1, 1 < 2 == true, sway 2 and 1 = {1, 2, 3, 4, 5, 6}, j-- == 0 break
{1, 2, 3, 4, 5, 6}
*/
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
