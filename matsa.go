// Matsa Discovery Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

// version 0.0.1

package matsa

type List_f64 []float64
type List_str []string

// general list methods

//length - return the number of values in the collection
func (list List_str) Length() int {
    return len(list)
}

func (list List_f64) Length() int {
    return len(list)
}

//isempty - provides an clear way to determine if list has elements or not
func (list List_f64) IsEmpty() bool {
    if list.Length() == 0 {
        return true
    }
    return false
}

func (list List_str) IsEmpty() bool {
    if list.Length() == 0 {
        return true
    }
    return false
}

//insert item into list
func (list List_f64) Insert(values ...float64) List_f64 {
    for _, val := range values {
        list = append(list, val)
    }
    return list
}

func (list List_str) Insert(values ...string) List_str {
    for _, val := range values {
        list = append(list, val)
    }
    return list
}

//removebyindex - item from list by index
func (list List_f64) RemoveByIndex(i int) List_f64 {
    list, list[len(list)-1] = append(list[:i], list[i+1:]...), 0
    return list
}

func (list List_str) RemoveByIndex(i int) List_str {
    list, list[len(list)-1] = append(list[:i], list[i+1:]...), ""
    return list
}

//removebyvalue - item from list by value
func (list List_f64) RemoveByValue(v float64) List_f64 {
    return list.Without(v)
}

func (list List_str) RemoveByValue(v string) List_str {
    return list.Without(v)
}

//pop element from list
func (list List_f64) Pop() (float64, List_f64) {
    out, list := list[len(list)-1], list[:len(list)-1]
    return out, list
}

func (list List_str) Pop() (string, List_str) {
    out, list := list[len(list)-1], list[:len(list)-1]
    return out, list
}

//push element onto list
func (list List_f64) Push(val float64) List_f64 {
    list.Insert(val)
    return list
}

func (list List_str) Push(val string) List_str {
    list.Insert(val)
    return list
}

//copy - create a copy of the list
func (list List_f64) Copy() List_f64 {
    out := make(List_f64, len(list))
    copy(out, list)
    return out
}

func (list List_str) Copy() List_str {
    out := make(List_str, len(list))
    copy(out, list)
    return out
}

//append - add a list onto the end of the current list
func (list List_f64) Append(list2 List_f64) List_f64 {
    list = append(list, list2...)
    return list
}

func (list List_str) Append(list2 List_str) List_str {
    list = append(list, list2...)
    return list
}
