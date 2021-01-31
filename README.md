# NEWS-AGGREGATOR

Test task to the "Tochka Bank" company

## INSTALLATION

    $ cd deployments
    $ docker compose up

## REST API

Коллекцию для postman можно получить по этой ссылке:

https://www.getpostman.com/collections/-

### GET /health

Liveness probe для проверки соединения с бд. Отправляет 500 status code, если нет связи с бд.   

## Environments Variables
   
   Variable             | Description
   ------------------   | -----------
   APP_NAME             | Имя сервиса
   HTTP_SERVER_PORT     | Порт http сервера
   DB_USER              | Имя пользователя базы данных 
   DB_PASS              | Пароль базы данных
   DB_HOST              | Адрес базы данных
   DB_PORT              | Порт базы данных
   DB_NAME              | Имя базы данных
   DB_MAX_CONNS         | Максимальное количество соединений в пуле (по умолчанию 32)
   DB_MAX_LIFETIME      | Максимальное время жизни соединения (по умолчанию 5m)
   DB_TIMEOUT           | Ограничение времени выполнения запроса (по умолчанию 30s)
   DB_RETRIES           | Количество попыток повтора запросов (по умолчанию 5)
   LOGGER_ADDRESS       | UDP адрес логгера
   LOGGER_LEVEL         | Уровень логирования для отправки [TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC] (по умолчанию INFO)
   LOGGER_POD_NAME      | kube pod
   LOGGER_POD_NODE      | kube node
   LOGGER_POD_NAMESPACE | kube pod namespace