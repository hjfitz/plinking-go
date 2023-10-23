package main

import "fmt"

func main() {
				rules := map[int]string{
								3: "fizz",
								5: "buzz",
				}

				fizzbuzz(rules, 100)
}

func fizzbuzz(rules map[int]string, max int) {
				for i := 0; i < max; i += 1 {
								out := ""
								for n, s := range rules {
												if i % n == 0 {
																out += s
												}
								}
								if len(out) == 0 {
												fmt.Println(i)
								} else {
												fmt.Println(out)
								}
				}
}
