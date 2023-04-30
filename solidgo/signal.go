package solidgo

type Signal[T any] struct {
	storage  T
	notifies []func()
}

func CreateSignal[T any](defaultValue T) (func() T, func(T)) {
	signal := Signal[T]{storage: defaultValue}
	return signal.Get, signal.Set
}

// TODO panic when called outside effect
func (signal *Signal[T]) Get() T {
	if currentEffect != nil {
		signal.notifies = append(signal.notifies, currentEffect)
	}
	return signal.storage
}
func (signal *Signal[T]) Set(value T) {
	signal.storage = value
	for _, function := range signal.notifies {
		function()
	}
}

var currentEffect func()

func CreateEffect(function func()) {
	currentEffect = function
	function()
	currentEffect = nil
}
