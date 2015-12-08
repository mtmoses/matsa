// Matsa Discovery Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package matsa

import (
    "math"
    "math/rand"
    "time"
    "sort"
    "errors"
)

func (list List_f64) Sum() float64 {
    
    if list.Length() == 0 {
		return 0
	}

    var sum float64
    
	// Add em up
	for _, n := range list {
		sum += n
	}

	return sum
}

func (list List_f64) Mean() float64 {
    
    l := list.Length()
    if l == 0 {
        return 0
    }
    
    sum := list.Sum()
    
    return sum / float64(l)
}

func (list List_f64) GeometricMean() float64 {
    
    l := list.Length()
    if l == 0 {
        return 0
    }
    
    // Get the product of all the numbers
    var p float64
    for _, n := range list {
        if p == 0 {
            p = n
        } else {
            p *= n
        }
    }
    
    // Calculate the geometric mean
    return math.Pow(p, 1/float64(l))
}

func (list List_f64) HarmonicMean() float64 {
    
    l := list.Length()
	if l == 0 {
		return 0
	}

	// Get the sum of all the numbers reciprocals and return an
	// error for values that cannot be included in harmonic mean
	var p float64
	for _, n := range list {
		if n < 0 {
			return 0
		} else if n == 0 {
			return 0
		}
		p += (1 / n)
	}

	return float64(l) / p
}

func (list List_f64) Median() float64 {
    // Start by sorting a copy of the slice
	c := sortedCopy(list)

    var median float64
    
	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	l := len(c)
    
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = c[l/2-1 : l/2+1].Mean()
	} else {
		median = float64(c[l/2])
	}

	return median
}

func (list List_f64) Mode() List_f64 {
    
    // Return the input if there's only one number
	l := list.Length()
	if l == 1 {
		return list
	} else if l == 0 {
		return nil
	}

	c := sortedCopyDif(list)
	// Traverse sorted array,
	// tracking the longest repeating sequence
	mode := make(List_f64, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		switch {
		case c[i] == c[i-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, c[i-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], c[i-1])
			maxCnt, cnt = cnt, 1
		}
	}
	switch {
	case cnt == maxCnt:
		mode = append(mode, c[l-1])
	case cnt > maxCnt:
		mode = append(mode[:0], c[l-1])
		maxCnt = cnt
	}

	// Since length must be greater than 1,
	// check for slices of distinct values
	if maxCnt == 1 {
		return []float64{}
	}

	return mode
}

func (list List_f64) StandardDeviation() float64 {
    
    if list.Length() == 0 {
		return 0
	}

	// Get the population variance
	vp := list.Variance()

	// Return the population standard deviation
	return math.Pow(vp, 0.5)
}

func (list List_f64) Variance() float64 {
    
    l := list.Length()
	if l == 0 {
		return 0
	}

	// Sum the square of the mean subtracted from each number
	m := list.Mean()

    var variance float64
    
	for _, n := range list {
		variance += (float64(n) - m) * (float64(n) - m)
	}

	vrnce := variance / float64(l)

	return vrnce
}

func (list List_f64) Percentile(percent float64) float64 {
    
    if list.Length() == 0 {
		return 0
	}

    var percentile float64
    
	// Start by sorting a copy of the slice
	c := sortedCopy(list)

	// Multiple percent by length of input
	index := (percent / 100) * float64(len(c))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := float64ToInt(index)

		// Find the average of the index and following values
        percentile = List_f64{c[i-1], c[i]}.Mean()

	} else {

		// Convert float to int
		i := float64ToInt(index)

		// Find the value at the index
		percentile = c[i-1]

	}

	return percentile
}

func (list List_f64) Quartile() List_f64 {
    
    il := list.Length()
	if il == 0 {
		return List_f64{0,0,0}
	}

	// Start by sorting a copy of the slice
	cpy := sortedCopy(list)

	// Find the cutoff places depeding on if
	// the input slice length is even or odd
	var c1 int
	var c2 int
	if il%2 == 0 {
		c1 = il / 2
		c2 = il / 2
	} else {
		c1 = (il - 1) / 2
		c2 = c1 + 1
	}

	// Find the Medians with the cutoff points
	Q1 := cpy[:c1].Median()
	Q2 := cpy.Median()
	Q3 := cpy[c2:].Median()

	return List_f64{Q1, Q2, Q3}
}

// InterQuartileRange finds the range between Q1 and Q3
func (list List_f64) InterQuartileRange() float64 {
	if list.Length() == 0 {
		return 0
	}
	qs := list.Quartile()
	iqr := qs[2] - qs[0]
	return iqr
}

// Trimean finds the average of the median and the midhinge
func (list List_f64) Trimean() float64 {
	if list.Length() == 0 {
		return 0
	}

	q := list.Quartile()

	return (q[0] + (q[1] * 2) + q[2]) / 4
}

func (list List_f64) Correlation(list2 List_f64) float64 {
    l1 := list.Length()
	l2 := list2.Length()

	if l1 == 0 || l2 == 0 {
		return 0
	}

	if l1 != l2 {
		return 0
	}

	sdev1 := list.StandardDeviation()
	sdev2 := list2.StandardDeviation()

	if sdev1 == 0 || sdev2 == 0 {
		return 0
	}

	covp := list.Covariance(list2)
	return covp / (sdev1 * sdev2)
}

func (list List_f64) Pearson(list2 List_f64) float64 {
    return list.Correlation(list2)
}

func (list List_f64) Covariance(list2 List_f64) float64 {
    
    l1 := list.Length()
	l2 := list2.Length()

	if l1 == 0 || l2 == 0 {
		return 0
	}

	if l1 != l2 {
		return 0
	}

	m1 := list.Mean()
	m2 := list2.Mean()

	// Calculate sum of squares
	var ss float64
	for i := 0; i < l1; i++ {
		delta1 := (list[i] - m1)
		delta2 := (list2[i] - m2)
		ss += (delta1*delta2 - ss) / float64(i+1)
	}

	return ss * float64(l1) / float64(l1-1)
}

//utility functions
func Round(input float64, places int) (rounded float64, err error) {

	// If the float is not a number
	if math.IsNaN(input) {
		return 0.0, errors.New("Not a number")
	}

	// Find out the actual sign and correct the input for later
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	// Use the places arg to get the amount of precision wanted
	precision := math.Pow(10, float64(places))

	// Find the decimal place we are looking to round
	digit := input * precision

	// Get the actual decimal number as a fraction to be compared
	_, decimal := math.Modf(digit)

	// If the decimal is less than .5 we round down otherwise up
	if decimal >= 0.5 {
		rounded = math.Ceil(digit)
	} else {
		rounded = math.Floor(digit)
	}

	// Finally we do the math to actually create a rounded number
	return rounded / precision * sign, nil
}

func GenRandomFloat() float64 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return (float64(rand.Intn(100)) + r.Float64())
}

// sortedCopy returns a sorted copy of List_f64
func sortedCopy(list List_f64) List_f64 {
	cpy := list.Copy()
	sort.Float64s(cpy)
	return cpy
}

// sortedCopyDif returns a sorted copy of List_f64
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func sortedCopyDif(list List_f64) List_f64 {
	if sort.Float64sAreSorted(list) {
		return list
	}
	cpy := list.Copy()
	sort.Float64s(cpy)
	return cpy
}

// float64ToInt rounds a float64 to an int
func float64ToInt(input float64) int {
	r, _ := Round(input, 0)
	return int(r)
}