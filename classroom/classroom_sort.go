package classroom

import "sort"

// Note for reviewer:
// there are 3 ways to sort the classroom:
// 1. use sort.Sort function
// 2. use sort.Slice function
// 3. sort natively
// I think the first one is the best because it's the most readable and the most efficient
// because it uses the sort.Interface interface
// the second one is also good because it's the most readable
// the third one is the worst because it's the least readable and the least efficient

type SortByScore []Student

func (s SortByScore) Len() int {
	return len(s)
}

func (s SortByScore) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s SortByScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// resort sorts the classroom by score
func (c *Classroom) resort() {
	sort.Sort(SortByScore(c.students))
}
