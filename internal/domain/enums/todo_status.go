package enums

type TodoStatus uint

const (
	Finish TodoStatus = iota + 1
	NotFinish
	Fail
)

func (bt TodoStatus) String() string {
	switch bt {
	case Finish:
		return "Finish"
	case NotFinish:
		return "NotFinish"
	case Fail:
		return "Fail"
	}
	return ""
}
