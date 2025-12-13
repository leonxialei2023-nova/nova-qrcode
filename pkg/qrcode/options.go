package qrcode

type Level int

const (
	Low Level = iota
	Medium
	High
	Highest
)

type Options struct {
	Size  int   // 像素尺寸，如 256 / 384
	Level Level // 容错等级
}
