package main

type Family struct {
	No   int
	Name *Family
	Left *Family
}

func main() {
	// 构建一个二叉树
	root := &Family{
		No:   1,
		Name: "Father-shuwen",
	}

}
