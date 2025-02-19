package main

func main1() {
	_ = stackIt()
 }
 //go:noinline
 func stackIt() int {
	y := 2
	return y * 2
 }


 func main2() {
	_ = stackIt2()
 }
 //go:noinline
 func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
 }


 // slices and vector
 func mainSlices() int {
	sum := 0
   a := slices()
   for _, v := range a {
		sum += v
	}
	return sum
 }
 func slices() []int{
	// heap
	v := []int{}
	v = append(v, 4, 2, 1, 2, 3, 9999, 12, 23, 12, 23)
	for i := range v {
		v[i] += 4
	}

	return v
 }

 func mainVector() int{
	sum := 0
	a := vector()
	for _, v := range a {
		sum += v
	}
	return sum
 }
 func vector() [10]int{
	// stack
	var v [10]int
	v[0] = 4
	v[1] = 2
	v[2] = 1
	v[3] = 2
	v[4] = 3
	for i := range v {
		v[i] += 4
	}

	return v
 }