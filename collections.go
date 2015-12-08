// Matsa Discovery Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package matsa

import (
    "math/rand"
)

//sample - produce a random sample from the list. Pass a number to return n random elements
//from the collection.
func (list List_str) Sample(n int) List_str {
    if n > len(list) {
        n = len(list)
    } else if n == 0 {
        n = 1
    }
    vsm := make(List_str, n)
    for i:=0; i<n; i++ {
        j := rand.Intn(i + 1)
        vsm[i], vsm[j] = list[j], list[i]
    }
    return vsm
}

func (list List_f64) Sample(n int) List_f64 {
    if n > len(list) {
        n = len(list)
    } else if n == 0 {
        n = 1
    }
    vsm := make(List_f64, n)
    for i:=0; i<n; i++ {
        j := rand.Intn(i + 1)
        vsm[i], vsm[j] = list[j], list[i]
    }
    return vsm
}

//reverse
func (list List_str) Reverse() List_str {
    vsm := make(List_str, len(list))
    for i:=0; i<len(list); i++ {
        opp := len(list)-1-i
        vsm[i], vsm[opp] = list[opp], list[i]
    }
    return vsm
}

func (list List_f64) Reverse() List_f64 {
    vsm := make(List_f64, len(list))
    for i:=0; i<len(list); i++ {
        opp := len(list)-1-i
        vsm[i], vsm[opp] = list[opp], list[i]
    }
    return vsm
}

//min - returns the smallest value from the collection
func (list List_f64) Min() float64 {
    // Get the count of numbers in the slice
	l := list.Length()

	// Return an error if there are no numbers
	if l == 0 {
		return 0
	}

	// Get the first value as the starting point
	min := list[0]

	// Iterate until done checking for a lower value
	for i := 1; i < l; i++ {
		if list[i] < min {
			min = list[i]
		}
	}
    
	return min
}

//max - returns the largest value from the collection
func (list List_f64) Max() float64  {
    // Get the count of numbers in the slice
	l := list.Length()
    
    // Return an error if there are no numbers
	if l == 0 {
		return 0
	}

	// Get the first value as the starting point
	max := list[0]

	// Loop and replace higher values
	for i := 1; i < l; i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

//map - produces a new array of values by mapping each value in list through a transformation function (iteratee)
func (list List_str) Map(iteratee func(string) string) List_str {
    vsm := make(List_str, len(list))
    for i, v := range list {
        vsm[i] = iteratee(v)
    }
    return vsm
}

func (list List_f64) Map(iteratee func(float64) float64) List_f64 {
    vsm := make(List_f64, len(list))
    for i, v := range list {
        vsm[i] = iteratee(v)
    }
    return vsm
}

//reduce - boils down a list of values into a single value. Memo is the initial state of the reduction,
//and each successive step of it should be returned by iteratee.
func (list List_str) Reduce(iteratee func(string, string) string, memo string) string {
    
    var vsm string
    
    for _, v := range list {
        vsm = iteratee(memo, v)
        memo = vsm
    }
    return vsm
}

func (list List_f64) Reduce(iteratee func(float64, float64) float64, memo float64) float64 {
    
    var vsm float64
    
    for _, v := range list {
        vsm = iteratee(memo, v)
        memo = vsm
    }
    return vsm
}

//index - returns index of value
func (list List_str) Index(t string) int {
    for i, v := range list {
        if v == t {
            return i
        }
    }
    return -1
}

func (list List_f64) Index(t float64) int {
    for i, v := range list {
        if v == t {
            return i
        }
    }
    return -1
}

//filter - looks through each value in the list, returning all the values that pass a truth test (predicate).
func (list List_str) Filter(predicate func(string) bool) List_str {
    vsf := make(List_str, 0)
    for _, v := range list {
        if predicate(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func (list List_f64) Filter(predicate func(float64) bool) List_f64 {
    vsf := make(List_f64, 0)
    for _, v := range list {
        if predicate(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

//reject - returns the values in list without the elements that the truth test (predicate) passes. opposite of filter
func (list List_str) Reject(predicate func(string) bool) List_str {
    vsf := make(List_str, 0)
    for _, v := range list {
        if predicate(v) == false {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func (list List_f64) Reject(predicate func(float64) bool) List_f64 {
    vsf := make(List_f64, 0)
    for _, v := range list {
        if predicate(v) == false {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

//all (alias of every) - returns true if all of the values in the list pass the predicate truth test.
func (list List_str) All(predicate func(string) bool) bool {
    for _, v := range list {
        if !predicate(v) {
            return false
        }
    }
    return true
}

func (list List_f64) All(predicate func(float64) bool) bool {
    for _, v := range list {
        if !predicate(v) {
            return false
        }
    }
    return true
}

//any (alias of some) - returns true if any of the values in the list pass the predicate truth test.
func (list List_str) Any(predicate func(string) bool) bool {
    for _, v := range list {
        if predicate(v) {
            return true
        }
    }
    return false
}

func (list List_f64) Any(predicate func(float64) bool) bool {
    for _, v := range list {
        if predicate(v) {
            return true
        }
    }
    return false
}

//contains - returns true if the value is present in the collection
func (list List_str) Contains(t string) bool {
    return list.Index(t) >= 0
}

func (list List_f64) Contains(t float64) bool {
    return list.Index(t) >= 0
}

//shuffle - returns a shuffled copy of the collection, using a version of the Fisher-Yates shuffle.
func (list List_str) Shuffle() List_str {
    vsm := make(List_str, len(list))
    for i := range list {
        j := rand.Intn(i + 1)
        vsm[i], vsm[j] = list[j], list[i]
    }
    return vsm
}

func (list List_f64) Shuffle() List_f64 {
    vsm := make(List_f64, len(list))
    for i := range list {
        j := rand.Intn(i + 1)
        vsm[i], vsm[j] = list[j], list[i]
    }
    return vsm
}
