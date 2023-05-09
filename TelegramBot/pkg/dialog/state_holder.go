package dialog

type StateHolder struct {
	currentStates map[int64]State
}

var stateHolder *StateHolder = nil

func (sh *StateHolder) init() {
	sh.currentStates = make(map[int64]State)
}

func initHolder() {
	stateHolder = new(StateHolder)
	stateHolder.init()
}

func GetState(chatId int64) State {
	if stateHolder == nil {
		initHolder()
	}
	return stateHolder.currentStates[chatId]
}

func SetState(chatId int64, state State) {
	if stateHolder == nil {
		initHolder()
	}
	stateHolder.currentStates[chatId] = state
}
