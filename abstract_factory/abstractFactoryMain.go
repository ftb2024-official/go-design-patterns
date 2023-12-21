/*
Абстрактная фабрика(Abstract factory) - порождающий паттерн проектирования, который позволяет создавать семейства связанных объектов, не привязываясь к конкретным классам создаваемых объектов.

Проблема: Представьте, что вы пишете симулятор мебельного магазина. Ваш код содержит: 1) Семейство зависимых продуктов. Скажем: Кресло + Диван + Столик. 2) Несколько вариаций эирнр семейства. Например, продукты Кресло
*/

package abstract_factory

import "fmt"

func AbstractFactoryMain() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(shoe IShoe) {
	fmt.Printf("Logo: %v\n", shoe.getLogo())
	fmt.Printf("Size: %v\n", shoe.getSize())
}

func printShirtDetails(shirt IShirt) {
	fmt.Printf("Logo: %v\n", shirt.getLogo())
	fmt.Printf("Size: %v\n", shirt.getSize())
}
