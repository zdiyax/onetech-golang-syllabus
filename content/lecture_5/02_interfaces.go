package lecture_5

type Strings []string

//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}

//todo: сделать структуру Strings сортируемой
//todo: написать тесты для сортировки
func (s *Strings) Len() int {
	panic("implement me later!")
}

func (s *Strings) Swap(i, j int) {
	panic("implement me later!")
}

func (s *Strings) Less(i, j int) {
	panic("implement me later!")
}

func main() {
	//sort.Sort(Strings) // todo: должно работать!
}
