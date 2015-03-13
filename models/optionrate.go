package models

type OptionRate struct {
	Name string
	Rate float32
}

func (opt *OptionRate) GetName() string {
	return opt.Name
}

func (opt *OptionRate) GetRate() float32 {
	return opt.Rate
}
