package main

type StringArray []string
type ForEachFunc func(index int, value string) bool

func (l StringArray) ForEach(each ForEachFunc) {
	for index, val := range l {
		if cont := each(index, val); cont == false {
			break
		}
	}
}
