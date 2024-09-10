package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randInt(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randInt(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randInt(-3, -1))
	default:
		return "неизвестный класс персонажа"
	}
}

func defence(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randInt(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randInt(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randInt(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

func special(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	default:
		return "неизвестный класс персонажа"
	}
}

func startTraining(charName, charClass string) string {
	switch charClass {
	case "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanln(&cmd)

		switch cmd {
		case "attack":
			fmt.Println(attack(charName, charClass))
		case "defence":
			fmt.Println(defence(charName, charClass))
		case "special":
			fmt.Println(special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

func chooseCharClass() string {
	var approveChoice string
	var charClass string

	for strings.ToLower(approveChoice) != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanln(&charClass)

		switch charClass {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanln(&approveChoice)
	}
	return charClass
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanln(&charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := chooseCharClass()

	fmt.Println(startTraining(charName, charClass))
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}
