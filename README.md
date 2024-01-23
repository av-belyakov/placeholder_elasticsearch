Placeholder_Elasticsearch v0.1.2

Конфигурационные параметры для сервиса могут быть заданы как через конфигурационный файл так и методом установки переменных окружения.

Типы конфигурационных файлов:

- config.yaml общий конфигурационный файл
- config_dev.yaml конфигурационный файл используемый для тестов при разработке
- config_prod.yaml конфигурационный файл применяемый в продуктовом режиме

Основная переменная окружения для данного приложения - GO_PHELASTIC_MAIN. На основании
значения этой переменной принимается решение какой из конфигурационных файлов config_dev.yaml или config_prod.yaml использовать. При GO_PHELASTIC_MAIN=development
будет использоваться config_dev.yaml, во всех остальных случаях, в том числе и при отсутствии переменной окружения GO_PHELASTIC_MAIN будет использоваться конфигурационный файл config_prod.yaml. Перечень переменных окружения которые можно использовать для настройки приложения:

//Переменная окружения отвечающая за тип запуска приложения "development" или "production"
GO_PHELASTIC_MAIN

//Подключение к NATS
GO_PHELASTIC_NHOST
GO_PHELASTIC_NPORT

//Подключение к СУБД Elasticsearch
GO_PHELASTIC_ESSEND
GO_PHELASTIC_ESHOST
GO_PHELASTIC_ESPORT
GO_PHELASTIC_ESPREFIX
GO_PHELASTIC_ESINDEX
GO_PHELASTIC_ESUSER
GO_PHELASTIC_ESPASSWD

//Подключение к СУБД MongoDB
GO_PHELASTIC_MONGOHOST
GO_PHELASTIC_MONGOPORT
GO_PHELASTIC_MONGOUSER
GO_PHELASTIC_MONGOPASSWD
GO_PHELASTIC_MONGONAMEDB

Приоритет значений заданных через переменные окружения выше чем значений полученных из конфигурационных файлов. Таким образом можно осуществлять гибкую временную настройку приложения.

Сервис выполняет сделующие действия:

1. Соединение с NATS.
2. Прием кейсов от TheHive в формате JSON.
3. Валидация кейсов на основе специальных правил.
4. Валидация кейсов под определенные объекты.
5. Сохранение кейсов в локальной СУБД MongoDB используемой для буферизации данных, при отсутствии подключения к СУБД Elasticsearch.
6. Запись кейсов TheHive в СУБД Elasticsearch.
