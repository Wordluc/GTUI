package Utils

import "slices"

///Return the elements that aren't in from but are in to,"the new elements"
//from =[1,2,3]   to=[1,2,3,4,5]  result=[4,5]
//from =[1,2,3,4,5]   to=[1,2,3]  result=[]
func Diff[element comparable](from []element, to []element)[]element {
	var result []element=[]element{}
	for _, v := range to {
		if !slices.Contains(from, v) {
			result = append(result, v)
		}
	}
	return result
}
