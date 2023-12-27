package event

type IDispatcher interface {
	AddListener(ref any, id Id, call func(e IEvent))
	RemoveListener(ref any, id Id)
	Dispatch(e IEvent)
}
