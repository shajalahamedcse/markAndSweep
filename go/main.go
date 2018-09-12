package main

import (
	"fmt"
)

const (
	objInt objType = iota
	objPair
)

const stackMax = 6

type objType int

type Object struct {
	marked bool
	tp     objType

	value int

	head *Object
	tail *Object
}

/*
Now we can wrap that in a little virtual machine structure.
Its role in this story is to have a stack that stores the variables that are currently in scope.
Most language VMs are either stack-based (like the JVM and CLR) or register-based (like Lua).
In both cases, there is actually still a stack.
It’s used to store local variables and temporary variables needed in the middle of an expression.

We’ll model that explicitly and simply like so:
*/

type VM struct {
	stack     [stackMax]*Object
	stackSize int
}

/*
	Now that we’ve got our basic data structures in place, let’s slap together a bit of code to create some stuff.
 	First, let’s write a function that creates and initializes a VM:
*/
func newVM() *VM {
	vm := &VM{}
	vm.stackSize = 0
	return vm
}

func newObject(vm *VM, tp objType) *Object {
	object := &Object{}
	object.tp = tp
	return object
}

func push(vm *VM, value *Object) {
	if vm.stackSize > stackMax {
		fmt.Println("Stack overflow!")
	}
	vm.stackSize++
	vm.stack[vm.stackSize] = value
}

func pushInt(vm *VM, tp objType) *Object {
	object := newObject(vm, tp)
	object.tp = tp

	return object
}

func pushPair(vm *VM) *Object {
	object := newObject(vm, objPair)
	return object
}

func pop(vm *VM) *Object {
	if vm.stackSize < 0 {
		fmt.Println("Stack Underflow")
	}
	vm.stackSize--
	return vm.stack[vm.stackSize]
}

func main() {

}
