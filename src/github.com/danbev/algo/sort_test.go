package algo

import "testing"

func TestSelectionSort(t *testing.T) {
    var unsorted = [] int {2, 6, 4, 3, 5, 1}
    expected := [] int {1, 2, 3, 4, 5, 6}
    var sorted = SelectionSort(unsorted);
    for i := range sorted {
        if sorted[i] != expected[i] {
            t.Error("For", unsorted, "Expected", expected, "got", sorted)
            break
        }
    }
}

