package models

type OptionRate struct {
	name string
	rate float32
}

func (opt *OptionRate) GetName() string {
	return opt.name
}

func (opt *OptionRate) GetRate() float32 {
	return opt.rate
}
