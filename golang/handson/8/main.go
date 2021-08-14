package main

import "strings"

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
// hint how we test slice of string  comparation:
// https://www.geeksforgeeks.org/reflect-deepequal-function-in-golang-with-examples/

func Split(s, sep string) []string {
   var result []string
   i := strings.Index(s, sep)
   for i > -1 {
       result = append(result, s[:i])
       s = s[i+len(sep):]
       i = strings.Index(s, sep)
   }
   return append(result, s)
}