Placeholder_Elasticsearch v0.7.9

Конфигурационные параметры для сервиса могут быть заданы как через конфигурационный файл так и методом установки переменных окружения.

Типы конфигурационных файлов:

- config.yaml общий конфигурационный файл
- config_dev.yaml конфигурационный файл используемый для тестов при разработке
- config_prod.yaml конфигурационный файл применяемый в продуктовом режиме

Основная переменная окружения для данного приложения - GO_PHELASTIC_MAIN. На основании значения этой переменной принимается решение какой из конфигурационных файлов config_dev.yaml или config_prod.yaml использовать. При GO_PHELASTIC_MAIN=development
будет использоваться config_dev.yaml, во всех остальных случаях, в том числе и при отсутствии переменной окружения GO_PHELASTIC_MAIN будет использоваться конфигурационный файл config_prod.yaml. Перечень переменных окружения которые можно использовать для настройки приложения:

//Переменная окружения отвечающая за тип запуска приложения "development" или "production"
GO_PHELASTIC_MAIN

//Подключение к NATS
GO_PHELASTIC_NHOST
GO_PHELASTIC_NPORT
GO_PHELASTIC_SUBJECT_CASE
GO_PHELASTIC_SUBJECT_ALERT

//Подключение к СУБД Elasticsearch
GO_PHELASTIC_ESHOST
GO_PHELASTIC_ESPORT
GO_PHELASTIC_ESUSER
GO_PHELASTIC_ESPASSWD
GO_PHELASTIC_ESPREFIXCASE
GO_PHELASTIC_ESINDEXCASE
GO_PHELASTIC_ESPREFIXALERT
GO_PHELASTIC_ESINDEXALERT

//Подключение к СУБД MongoDB
GO_PHELASTIC_MONGOHOST
GO_PHELASTIC_MONGOPORT
GO_PHELASTIC_MONGOUSER
GO_PHELASTIC_MONGOPASSWD
GO_PHELASTIC_MONGONAMEDB

//Место нахождение правил
GO_PHELASTIC_RULES_DIR
GO_PHELASTIC_RULES_FILECASE
GO_PHELASTIC_RULES_FILEALERT

Приоритет значений заданных через переменные окружения выше чем значений полученных из конфигурационных файлов. Таким образом можно осуществлять гибкую временную настройку приложения.

Сервис выполняет сделующие действия:

1. Соединение с NATS.
2. Прием кейсов и алертов от TheHive в формате JSON.
3. Валидация кейсов и алертов на основе специальных правил.
4. Валидация кейсов и алертов под определенные объекты.
5. Сохранение кейсов и алертов в локальной СУБД MongoDB используемой для буферизации данных, при отсутствии подключения к СУБД Elasticsearch.
6. Запись кейсов и алертов TheHive в СУБД Elasticsearch.
