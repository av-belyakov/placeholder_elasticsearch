package mongodbinteractions

// SettingsInputChan
// Section - секция обработки данных
// Command - команда
// Data - данные
type SettingsInputChan struct {
	Section string
	Command string
	Data    interface{}
}

// ModuleDataBaseInteractionChannel описание типов данных циркулирующих между модулем взаимодействия с БД и Ядром приложения
// Section - секция обработки данных
// Command - команда
// AppTaskID - внутренний идентификатор задачи
type ModuleDataBaseInteractionChannel struct {
	Section   string
	Command   string
	AppTaskID string
	Data      *[]byte
}
