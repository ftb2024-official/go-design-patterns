/*
Фабричный метод - порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Применимость:
	- когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код;
	- когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки;
	- когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых;
*/

package factory_method

import "fmt"

func FactoryMethodMain() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %v\n", g.getName())
	fmt.Printf("Power: %v\n", g.getPower())
}
