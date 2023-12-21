/*
Строитель(Builder) - порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово. Строитель дает возможность использовать один и тот же код строительства для получения разных представлений объектов.

Проблема: Представьте сложный объект, требующий кропотливой пошаговой инициализации множества полей и вложенных объектов. Код инициализации таких объектов обычно спрятан внутри монструозного конструктора с десятком параметров. Либо ещё хуже - распылён по всему клиентскому коду.

Решение: Паттерн Строитель предлагает вынести конструирование объектов за пределы его собственного класса, поручив это дело отдельным объектам, называемым строителями.
*/

package builder

import "fmt"

func BuilderMain() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %v\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %v\n", normalHouse.windowType)
	fmt.Printf("Normal House Number of Floors: %v\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("Igloo House Door Type: %v\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %v\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Number of Floors: %v\n", iglooHouse.floor)
}
