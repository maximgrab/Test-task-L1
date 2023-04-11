package main

// Настроен как внешний под XML, но у нас всё в JSON
// Применим паттерн к аналитическому сервису работающему только с xml для отправки JSON
// Для применения паттерна создадим структуру адаптера наследующую JSON
// По итогу сервис видит XML

type AnalyticalDataService interface {
	SendXmlData()
}
type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	println("Отправка XML")
}

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() {
}

type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.ConvertToXml()
	println("Отправка XML")
}

func main() {

}
