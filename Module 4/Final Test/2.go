package main

import "fmt"

func main() {
	var o1, o2, o3 string
	var x, f, g, h int

	var res int

	fmt.Scanln(&x, &o1, &o2, &o3)

	// fix
	_ = f
	_ = g
	_ = h

	operationArray := [3]string{"f", "g", "h"}

	for _, v := range operationArray {
		if o3 == v {
			for _, v := range operationArray {
				if o2 == v {
					for _, v := range operationArray {
						if o1 == v {
							if o1 == "f" {
								fa(x, &f)
								res += f
							} else if o1 == "g" {
								ga(x, &g)
								res += g
							} else if o1 == "h" {
								ha(x, &h)
								res += h
							}
						}
					}	
					if o2 == "f" {
						fa(res, &f)
						res += f
					} else if o2 == "g" {
						ga(res, &g)
						res += g
					} else if o2 == "h" {
						ha(res, &h)
						res += h
					}
				}
			}
			if o3 == "f" {
				fa(res, &f)
				res += f
			} else if o3 == "g" {
				ga(res, &g)
				res += g
			} else if o3 == "h" {
				ha(res, &h)
				res += h
			}
		}
	}

	fmt.Println(res)
}

func fa(x int, f *int) {
	fmt.Println("x is", x)
	*f = (2 * x) + 5
} 

func ga(x int, g *int) {
	fmt.Println("x is", x)
	*g = x*x + (2 * x)
}

func ha(x int, h *int) {
	fmt.Println("x is", x)
	*h = x - 3 
}