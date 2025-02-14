package main

type ProductOfNumbers struct {
	products []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num != 0 {
		if len(this.products) == 0 {
			this.products = append(this.products, num)
		} else {
			this.products = append(this.products, num*this.products[len(this.products)-1])
		}
	} else {
		this.products = []int{}
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	l := len(this.products)
	if k > l {
		return 0
	} else if k == l {
		return this.products[l-1]
	} else {
		return this.products[l-1] / this.products[l-k-1]
	}
}
