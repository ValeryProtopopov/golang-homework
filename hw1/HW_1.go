package main

import "fmt"

// 1. Напишите функцию `multiply` с двумя вещественными аргументами, которая возвращает их произведение.
//    При этом, разрешается использовать только операции сложения и присваивания.
func multiply(x float64, y float64) float64 {
	result := 0.0
	for i := 0; i < int(y); i++ {
		result += x
	}
	temp := y - float64(int(y))

	return result + x*temp
}

// 2. Напишите функцию `fibonacci1` с одним целочисленным аргументом `n`,
//	  которая с помощью цикла находит и возвращает `n`-ый член последовательности Фибоначчи.
func fibonacci1(n int) int {
	if n == 1 || n == -1 || n == 0 {
		return n
	}
	result := 0
	fib0 := 0
	if n > 0 {
		fib1 := 1
		for i := 0; i < n-1; i++ {
			result = fib0 + fib1
			fib0 = fib1
			fib1 = result
		}
	} else {
		fib1 := -1
		for i := 0; i > n+1; i-- {
			result = fib0 + fib1
			fib0 = fib1
			fib1 = result
		}
	}

	return result
}

// 3. Напишите функцию `fibonacci2` с одним целочисленным аргументом `n`,
//    которая с помощью рекурсии находит возвращает `n`-ый член последовательности Фибоначчи.
func fibonacci2(n int) int {
	result := 0
	if n == 0 {
		return 0
	} else if n > 0 {
		if n == 1 {
			return 1
		}
		if n > 1 {
			result += fibonacci2(n-1) + fibonacci2(n-2)
		}
	} else if n < 0 {
		if n == -1 {
			return -1
		}
		if n < -1 {
			result += fibonacci2(n+1) + fibonacci2(n+2)
		}
	}

	return result
}

// 4. Напишите фунцию `bubble_sort` с одним аргументом типа ссылка на массив вещественных чисел.
//    Функция должна отсортировать массив методом пузырька. Функция не должна ничего возвращать.
func bubbleSort(array *[]float64) {
	for i := 0; i < len(*array); i++ {
		for j := 0; j < len(*array)-1-i; j++ {
			if (*array)[j] > (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
	}
}

// 5. Напишите фунцию `unique_count` с одним аргументом типа ссылка на массив целых чисел.
// Функция должна вернуть целое число — количество различных значений. Например для массива [1, 2, 3, 4, 1, 2, 2, 3, 2], правильный ответ — 4.
func uniqueCount(array *[]int) int {
	newArray := []int{}
	newArray = append(newArray, (*array)[0])

	count := 0
	for i := 0; i < len(*array); i++ {
		if (*array)[i] == 0 {
			count++
			break
		}
	}
	for i := 0; i < len(*array)-1; i++ {
		x := (*array)[i]
		if x == 0 {
			continue
		}
		for j := i + 1; j < len(*array); j++ {
			if (*array)[j] == x {
				(*array)[j] = 0
			}
		}
	}
	for i := 0; i < len(*array); i++ {
		if (*array)[i] != 0 {
			count++
		}
	}
	return count
}

func main() {
	//1
	fmt.Println("Task_1:")
	task1VChislo1 := 55.9
	task1VChislo2 := 1.1
	fmt.Println(multiply(task1VChislo1, task1VChislo2))
	//2
	fmt.Println("Task_2:")
	task2IChisloN := -10
	fmt.Println(fibonacci1(task2IChisloN))
	//3
	fmt.Println("Task_3:")
	task3IChisloN := 10
	fmt.Println(fibonacci2(task3IChisloN))
	//4
	fmt.Println("Task_4:")
	task4VArray := []float64{0, -5, 6.6, -5, 10}
	bubbleSort(&task4VArray)
	fmt.Println(task4VArray)
	//5
	fmt.Println("Task_5:")
	task5IArray := []int{1, 2, 3, 4, 1, 2, 2, 3, 2}
	fmt.Println(uniqueCount(&task5IArray))
}
