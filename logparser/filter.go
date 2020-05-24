package logparser

type Filter interface {
	Accept(e LogEntry) bool
}

type AndFilter struct {
	subs []Filter
}

func NewAndFilter(subs []Filter) *AndFilter {
	return &AndFilter{subs: subs}
}

func (f *AndFilter) Accept(e LogEntry) bool {
	for _, sub := range f.subs {
		if !sub.Accept(e) {
			return false
		}
	}
	return true
}

type OrFilter struct {
	subs []Filter
}

func NewOrFilter(subs []Filter) *OrFilter {
	return &OrFilter{subs: subs}
}

func (f *OrFilter) Accept(e LogEntry) bool {
	for _, sub := range f.subs {
		if sub.Accept(e) {
			return true
		}
	}
	return false
}

type NotFilter struct {
	sub Filter
}

func NewNotFilter(sub Filter) *NotFilter {
	return &NotFilter{sub: sub}
}

func (f *NotFilter) Accept(e LogEntry) bool {
	return !f.sub.Accept(e)
}
