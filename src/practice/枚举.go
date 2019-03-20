package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota // 开始生成枚举值, 默认为0
	Shuriken
	SniperRifle
	Rifle
	Blower
)

// 输出所有枚举值

// 使用枚举类型并

func main() {
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	var weapon Weapon = Blower
	fmt.Println(weapon)
}
