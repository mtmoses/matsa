// Matsa Discovery Library
// Copyright (C) 2015 by Todd Moses (todd@toddmoses.com)
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package matsa

//without - returns a copy of the array with all instances of the values removed.
func (list List_str) Without(values ...string) List_str {

    var predicate = func(v string) bool {
       
        for _, val := range values {
            if v == val {
                return true
            }
        }
        
        return false
    }
    
    return list.Reject(predicate)
}

func (list List_f64) Without(values ...float64) List_f64 {
    
    var predicate = func(v float64) bool {
       
        for _, val := range values {
            if v == val {
                return true
            }
        }
        
        return false
    }
    
    return list.Reject(predicate)
}

//union - computes the union of the passed-in lists: the list of unique items,
//in order, that are present in one or more of the lists.
func (list List_str) Union(lists ...List_str) List_str {

    vsf := make(List_str, 0)
    
    //append all lists into list
    for _, lst := range lists {
        vsf = list.Append(lst)
    }
    
    //return uniq (non-duplicate)
    return vsf.Uniq()
}

func (list List_f64) Union(lists ...List_f64) List_f64 {
    
    vsf := make(List_f64, 0)
    
    //append all lists into list
    for _, lst := range lists {
        vsf = list.Append(lst)
    }
    
    //return uniq (non-duplicate)
    return vsf.Uniq()
}

//intersection - computes the list of values that are the intersection
//of all the lists. Each value in the result is present in each of the lists.
func (list List_str) Intersection(lists ...List_str) List_str {

    vsf := make(List_str, 0)
    
    //iterate over main list to get values
    for _, lval := range list {
        
        var is_inter bool
        
        if len(lval) > 0 {
            
            //iterate over all passed in lists
            for _, lst := range lists {
                is_inter = lst.Contains(lval)
                if is_inter == false {
                    break
                }
            }
            
            if is_inter == true {
                vsf = append(vsf, lval)
            }
            
        }
        
    }
    
    return vsf
}

func (list List_f64) Intersection(lists ...List_f64) List_f64 {
    
    vsf := make(List_f64, 0)
    
    //iterate over main list to get values
    for _, lval := range list {
        
        var is_inter bool
        
        //iterate over all passed in lists
        for _, lst := range lists {
            is_inter = lst.Contains(lval)
            if is_inter == false {
                break
            }
        }
        
        if is_inter == true {
            vsf = append(vsf, lval)
        }
        
    }
    
    return vsf
}

//difference - returns the values from list that are not present in the other lists
func (list List_str) Difference(lists ...List_str) List_str {
    
    vsf := make(List_str, 0)
    
    //iterate over main list to get values
    for _, lval := range list {
        
        var is_inter bool
        
        if len(lval) > 0 {
            
            //iterate over all passed in lists
            for _, lst := range lists {
                is_inter = lst.Contains(lval)
                if is_inter == true {
                    break
                }
            }
            
            if is_inter == false {
                vsf = append(vsf, lval)
            }
            
        }
        
    }
    
    return vsf
}

func (list List_f64) Difference(lists ...List_f64) List_f64 {
    
    vsf := make(List_f64, 0)
    
    //iterate over main list to get values
    for _, lval := range list {
        
        var is_inter bool
        
        //iterate over all passed in lists
        for _, lst := range lists {
            is_inter = lst.Contains(lval)
            if is_inter == true {
                break
            }
        }
        
        if is_inter == false {
            vsf = append(vsf, lval)
        }
        
    }
    
    return vsf
}

//uniq - produces a duplicate-free version of the list
func (list List_str) Uniq() List_str {
    
    vsf := make(List_str, 0)
    
    for _, lval := range list {
        if vsf.Contains(lval) == false {
            vsf = append(vsf, lval)
        }
    }
    
    return vsf
}

func (list List_f64) Uniq() List_f64 {
    
    vsf := make(List_f64, 0)
    
    for _, lval := range list {
        if vsf.Contains(lval) == false {
            vsf = append(vsf, lval)
        }
    }
    
    return vsf
}