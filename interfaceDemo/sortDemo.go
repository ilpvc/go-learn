package main

type Student struct {
	ID    int
	Name  string
	Score int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
