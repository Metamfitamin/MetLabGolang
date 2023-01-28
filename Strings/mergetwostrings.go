package main

import "fmt"

func merge(s1 string, s2 string) string {
  var s string
  len1, len2 := len(s1), len(s2)
  ptr1, ptr2 := 0, 0

   for ptr1 < len1 && ptr2 < len2 {
    if s1[ptr1] < s2[ptr2] {
      s += string(s1[ptr1])
      ptr1++
    } else {
      s += string(s2[ptr2])
      ptr2++
    }
  }

  if ptr1 < len1 {
    s += s1[ptr1:]
  }

  if ptr2 < len2 {
    s += s2[ptr2:]
  }

  return s
}

func main() {
    var s1 string
    var s2 string
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Println(merge(s1, s2))
}
//merge two strings like this "1236" "1234" = 11223346
