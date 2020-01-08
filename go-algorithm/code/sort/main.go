package main

import (
	"math/rand"
)

// quick sort
// 分治
func main()  {
	var z []int

	for i := 0; i < 3; i++{
		z = append(z, rand.Intn(3))
	}
	
	quickSort(z)	
}

func quickSort(list []int)  {
	if len(list) <= 1{
		return
	}

	i, j := 0, len(list) - 1
	index := 1 // 第一次比较索引位置
	key := list[0] // 第一次比较参考值，选择第一个

	if list[index] > key{
		list[i], list[j] = list[j], list[i]
		j--
	}else{
		list[i], list[index] = list[index],list[i]
		i++
		index++
	}

	quickSort(list[:i]) // 处理参考值前面值
	quickSort(list[i+1:])
}


