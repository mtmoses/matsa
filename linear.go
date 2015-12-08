// Matsa Discovery Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package matsa

import (
    "math"
)

// Abs returns the List_f64 with nonnegative components.
func (vec List_f64) Abs() List_f64 {
    
    var vec2 List_f64
    
    if vec.Length() == 0 {
		return vec2
	}
    
    for _, val := range vec {
        vec2 = append(vec2, math.Abs(val))
    }
    
    return vec2
}

// Add returns the standard List_f64 sum of v and ov.
func (vec List_f64) Add(vec2 List_f64) List_f64 {
    
    var vec3 List_f64
    
    l1 := vec.Length()
	l2 := vec2.Length()

	if l1 == 0 || l2 == 0 {
		return vec3
	}

	if l1 != l2 {
		return vec3
	}
    
    for i, val := range vec { 
        sum := val + vec2[i]
        vec3 = append(vec3, sum)
    }
    
    return vec3
}

// Sub returns the standard List_f64 difference of v and ov.
func (vec List_f64) Sub(vec2 List_f64) List_f64 {
    
    var vec3 List_f64
    
    l1 := vec.Length()
	l2 := vec2.Length()

	if l1 == 0 || l2 == 0 {
		return vec3
	}

	if l1 != l2 {
		return vec3
	}
    
    for i, val := range vec { 
        sum := val - vec2[i]
        vec3 = append(vec3, sum)
    }
    
    return vec3
}

// Mul returns the standard scalar product of v and m.
func (vec List_f64) Mul(m float64) List_f64 {
    
    var vec2 List_f64
    
    l := vec.Length()

	if l == 0 {
		return vec2
	}
    
    for _, val := range vec { 
        sum := val * m
        vec2 = append(vec2, sum)
    }
    
    return vec2
}

// Dot returns the standard dot product of v and ov.
func (vec List_f64) Dot(vec2 List_f64) float64 {
    
    var big_sum float64
    
    l1 := vec.Length()
	l2 := vec2.Length()

	if l1 == 0 || l2 == 0 {
		return big_sum
	}

	if l1 != l2 {
		return big_sum
	}
    
    for i, val := range vec { 
        sum := val * vec2[i]
        big_sum = big_sum + sum
    }
    
    return big_sum
}

// Distance returns the Euclidean distance between v and ov.
func (vec List_f64) Distance(vec2 List_f64) float64 {
    vec3 := vec.Sub(vec2)    
    return vec3.Norm()
}

// Norm returns the List_f64's norm.
func (vec List_f64) Norm() float64 {
    if vec.Length() == 0 {
		return 0
	}
    
    dot := vec.Dot(vec)
    
    return math.Sqrt(dot)
}