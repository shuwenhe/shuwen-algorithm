package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	dic map[int]int
}

func Constructor() RandomizedSet {
	dic := make(map[int]int)
	res := RandomizedSet{
		dic: dic,
	}
	return res
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dic[val]
	if !ok {
		this.dic[val]++
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.dic[val]
	if ok {
		delete(this.dic, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	num := rand.Intn(len(this.dic))
	i := 0
	var a int
	for k, _ := range this.dic {
		a = k
		if i == num {
			return k
		}
		i++
	}
	return a
}
