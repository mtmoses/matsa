// Matsa Discovery Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package matsa

import (
    "math/rand"
)

//lessthen - detemrine if one element is less than the other
func (list List_f64) LessThan(i, j int) bool {
    return list[i] < list[j]
}

func (list List_str) LessThan(i, j int) bool {
    return list[i] < list[j]
}

//swap - swap elements place with another
func (list List_f64) Swap(i, j int) List_f64 {
    list[i], list[j] = list[j], list[i]
    return list
}

func (list List_str) Swap(i, j int) List_str {
    list[i], list[j] = list[j], list[i]
    return list
}

//quicksort - sorting algorithms based on divide et impera strategy,
//resulting in an O(n log n) complexity
func (list List_f64) QuickSort() List_f64 {
    
    if len(list) < 2 {
        return list
    }

    left, right := 0, len(list) - 1
    
    // Pick a pivot
    pivotIndex := rand.Int() % len(list)
    
    // Move the pivot to the right
    list[pivotIndex], list[right] = list[right], list[pivotIndex]
    
    // Pile elements smaller than the pivot on the left
    for i := range list {
      if list[i] < list[right] {
        list[i], list[left] = list[left], list[i]
        left++
      }
    }
    
    // Place the pivot after the last smaller element
    list[left], list[right] = list[right], list[left]
    
    // Go down the rabbit hole
    list[:left].QuickSort()
    list[left + 1:].QuickSort()
    
    return list
}

func (list List_str) QuickSort() List_str {
    
    if len(list) < 2 {
        return list
    }

    left, right := 0, len(list) - 1
    
    // Pick a pivot
    pivotIndex := rand.Int() % len(list)
    
    // Move the pivot to the right
    list[pivotIndex], list[right] = list[right], list[pivotIndex]
    
    // Pile elements smaller than the pivot on the left
    for i := range list {
      if list[i] < list[right] {
        list[i], list[left] = list[left], list[i]
        left++
      }
    }
    
    // Place the pivot after the last smaller element
    list[left], list[right] = list[right], list[left]
    
    // Go down the rabbit hole
    list[:left].QuickSort()
    list[left + 1:].QuickSort()
    
    return list
}

//shellsort - a more efficient variation of insertion sort
//sort, it works by comparing items of varying distances apart resulting
//in a run time complexity of O(n log2 n).
func (list List_f64) ShellSort() List_f64 {
    h := 1
	for h < len(list) {
		h = 3 * h + 1
	}
	for h >= 1 {
		for i := h; i < len(list); i++ {
			for j := i; j >= h && list[j] < list[j - h]; j = j - h {
                list.Swap(i, j)
			}
		}
		h = h/3
	}
    return list
}

func (list List_str) ShellSort() List_str {
    h := 1
	for h < len(list) {
		h = 3 * h + 1
	}
	for h >= 1 {
		for i := h; i < len(list); i++ {
			for j := i; j >= h && list[j] < list[j - h]; j = j - h {
                list.Swap(i, j)
			}
		}
		h = h/3
	}
    return list
}