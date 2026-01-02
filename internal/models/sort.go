package models

import "sort"

func (m *Model) ApplySort() {
	sort.SliceStable(m.Tasks, func(i, j int) bool {
		t1, t2 := m.Tasks[i], m.Tasks[j]
		switch m.SortMode {
		case SortTodoFirst:
			if t1.Done != t2.Done {
				return !t1.Done
			}
		case SortDoneFirst:
			if t1.Done != t2.Done {
				return t1.Done
			}
		}
		return t1.ID < t2.ID
	})
	if m.Cursor >= len(m.Tasks) && len(m.Tasks) > 0 {
		m.Cursor = len(m.Tasks) - 1
	}
}

