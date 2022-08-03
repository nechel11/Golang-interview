// https://leetcode.com/problems/valid-parentheses/

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.

package main

import "fmt"

func main(){
	s := "()[]{}"
	s1 := "([)]"
	fmt.Println(isValid(s), isValid(s1))
}

func isValid(s string) bool {
	var collector []int32

	for _, v := range(s){
		if v == '(' || v == '{' || v == '['{
			collector = append(collector, v)
		} else{
			if len(collector) != 0{
				prev_char := collector[len(collector) - 1]
				collector = collector[:len(collector) - 1]
				if v == ']' && prev_char == '['{
					continue
				} else if v == '}' && prev_char == '{'{
					continue
				} else if v == ')' && prev_char == '(' {
					continue
				} else {
					return false
				}
            } else {
                return false
            }
		}
	}
	if len(collector) != 0{
		return false
	}
	return true
}