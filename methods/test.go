package main

import "fmt"

type Vertex struct {
	x int
	y int
}

// this is a method with receiver argument
func (v Vertex) sumOf() int {
	return v.x + v.y
}

// pointer receiver
func (v *Vertex) update() {
	v.x = v.x * 10
	v.y = v.y * 10

}

// this is a method a defined has normal function
func sum(v Vertex) int {
	return v.x + v.y
}

func main() {
	test := Vertex{1, 2}
	fmt.Println(test)
	fmt.Println(test.sumOf())
	fmt.Println(sum(Vertex{1, 4}))
	test.update()
	fmt.Println(test)
}

/*
so in go there is no classes but the methods can be used on types
here for give type you defining what method can be used
a function with receive argument is called method
a recevier argument means in the function your are defining which type belongs to thats what is a method
pointer receiver is another extension of it basically it modifies source value of give type or struct
if you see most of the go projects use pointer receive by default because it does it create a copy of struct in each function call 
you can use pointer recevier for just returning intstead og only modifying the base thing thats it 

interesting thing: test.update() works even without using &test because Go has automatic dereferencing
when the compiler sees test.update() and knows update requires a *Vertex receiver, it interprets it as (&test).update() behind the scenes
this only works because test is addressable (its a local variable). Vertex{1, 2}.update() would NOT compile because a struct literal is not addressable
the same works the other way too - if you have a pointer and call a value receiver method, Go auto-dereferences it
so (&test).sumOf() would also work even though sumOf takes a Vertex not *Vertex
*/
