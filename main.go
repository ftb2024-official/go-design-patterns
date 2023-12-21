package main

import (
	abstractFactory "design-patterns/abstract_factory"
	builder "design-patterns/builder"
	factoryMethod "design-patterns/factory_method"
	prototype "design-patterns/prototype"
	singleton "design-patterns/singleton"
	"fmt"
)

func main() {
	builder.BuilderMain()
	fmt.Println()
	abstractFactory.AbstractFactoryMain()
	fmt.Println()
	factoryMethod.FactoryMethodMain()
	fmt.Println()
	prototype.PrototypeMain()
	fmt.Println()
	singleton.SingletonMain()
}
