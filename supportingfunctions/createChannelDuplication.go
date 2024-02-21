package supportingfunctions

// CreateChannelDuplication выполняет дублирование одного канала
// на заданное число каналов такого же типа
func CreateChannelDuplication[T any](input chan T, count int) []chan T {
	outputs := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		outputs = append(outputs, make(chan T))
	}

	go func() {
		for data := range input {
			for _, channel := range outputs {
				channel <- data
			}
		}
	}()

	return outputs
}
