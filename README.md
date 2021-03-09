# NEWS-AGGREGATOR

## ОПИСАНИЕ

Перед вами новостной агрегатор, который умеет парсить:

* новости и статьи сайта https://habr.com/ru/
* новости сайта https://hi-tech.news/

## ЛОКАЛЬНЫЙ ЗАПУСК

    $ cd deployments
    $ docker compose up

## REST API

```http request
url: http://localhost:8086/
```

### POST /habr

Загрузка в базу статей или новостей с хабра. Критерии парсинга задаются в теле запроса:

```json
{
    "articles": false
}
```
или
```json
{
    "articles": true,
    "all": true,
    "scope": 10
}
```
или
```json
{
    "articles": true,
    "best": true,
    "period": "day"
}
```

**Критерии**

Param         | Value     | Description
---------     |--------   | -----------
**articles**  | **bool**  | **обязательный критерий. определяет что собирать**
`             | true      | будут статьи с сайта
`             | false     | будут последние новости с сайта
**all**       | **bool**  | **новые статьи. работает если articles = true и best = false**
`             | true      | 
`             | false     |  
**rating**    | **uint**  | **оценка статьи. можно не указывать, тогда парсер будет собирать рандомные статьи**
`             | 10        | выше 10
`             | 25        | выше 25 
`             | 50        | выше 50 
`             | 100       | выше 100  
`             | empty     | без порога            
**best**      | **bool**  | **лучшие статьи. работает если articles = true и all = false** 
`             | true      |        
`             | false     |                                                      
**period**    | **string**| **период в топе**
`             | "day"     | сутки
`             | "week"    | неделя 
`             | "month"   | месяц
`             | "year"    | год 
`             | empty     | выведется за сутки   

**Return values**  
```json
{
    "data": {
        "message": "Habr's site data was parsed and saved"
    }
}
```

### GET /habr?limit=3&offset=1

Получение всех новостей из базы.

**Return values**  
```json
{
    "data": [
        {
            "id": "e0a01eb0-40e4-481a-9304-de0ed5a1cad5",
            "author": "AnnieBronson",
            "author_link": "https://habr.com/ru/users/AnnieBronson/",
            "title": "Исследование зумбомбинга утверждает, что контрмеры эффекта не дают",
            "views": "2,7k",
            "publication_date": "вчера в 10:53",
            "link": "https://habr.com/ru/news/t/541008/",
            "created_at": "2021-02-06T21:06:48.099648Z"
        },
        {
            "id": "627dbe1b-2791-4b25-a79c-ad8190be7152",
            "author": "AnnieBronson",
            "author_link": "https://habr.com/ru/users/AnnieBronson/",
            "title": "Google добавит в смартфоны Pixel функцию отслеживания пульса: палец нужно будет приложить к камере",
            "views": "2,5k",
            "publication_date": "вчера в 11:33",
            "link": "https://habr.com/ru/news/t/541022/",
            "created_at": "2021-02-06T21:06:48.094766Z"
        },
        {
            "id": "f1bc2db3-68df-410b-8101-f9f3f993e788",
            "author": "avouner",
            "author_link": "https://habr.com/ru/users/avouner/",
            "title": "Администрация Байдена пообещала не снимать санкции с Huawei и других китайских компаний",
            "views": "1,6k",
            "publication_date": "вчера в 11:57",
            "link": "https://habr.com/ru/news/t/541026/",
            "created_at": "2021-02-06T21:06:48.092152Z"
        }
    ]   
}
```

### GET /habr/search?title=Как сжать

Поиск новости по подстроке в заголовке.

**Return values**  
```json
{
    "data": [
        {
            "id": "d3d5f9e1-3523-4104-8dca-0e5c2a926360",
            "author": "BlackBox",
            "author_link": "https://habr.com/ru/users/BlackBox/",
            "title": "Паспортный контроль, или Как сжать полтора гигабайта до 42 мегабайт",
            "views": "12,2k",
            "publication_date": "вчера в 12:52",
            "link": "https://habr.com/ru/post/538358/",
            "created_at": "2021-02-06T09:01:58.483185Z"
        }
    ]
}
```

### POST /hi_tech_news

Загрузка в базу статей или новостей с hi-tech news. Критерии парсинга задаются в теле запроса:

```json
{
    "category": "Смартфоны",
    "page": 2
}
```

**Критерии**

Param         | Value         | Description
---------     | ------------- | -----------
**category**  | **string**    | **определяет какую категорию собирать**
`             | "Смартфоны"   | 
`             | "Медицина"    |        
`             | "Прочее"      | 
`             | empty         | _P.S.: Категорий на сайте больше, я взяла эти для примера собирает все_         
**page**      | **uint**      | **определяет какую страницу парсить**
`             | от 1 до 90    | 

**Return values**  
```json
{
    "data": {
        "message": "Hi-tech's site data was parsed and saved"
    }
}
```

### GET /hi_tech_news?limit=3&offset=1

Получение всех новостей из базы.

**Return values**  
```json
{
    "data": [
        {
            "id": "c9e9ca43-a273-4421-9c0c-3c1a897b5b64",
            "category": "Телевизоры",
            "title": "Philips представила телевизоры серии 9000 Mini LED. Особый интерес ожидается от геймеров",
            "preview": "Уже через несколько месяцев в магазины поступят первые телевизоры Philips, оснащенные мини-светодиодными матрицами – Mini LED. Серия 9000 также предлагает улучшения для геймеров и улучшенную...",
            "link": "https://hi-tech.news/tv/2683-philips-predstavila-televizory-serii-9000-mini-led-osobyj-interes-ozhidaetsja-ot-gejmerov.html",
            "created_at": "2021-02-06T20:06:05.007291Z"
        },
        {
            "id": "aa5d81a6-4400-4d31-b123-e5251f5490f3",
            "category": "Смартфоны",
            "title": "Глобальная премьера Xiaomi Mi 11 и MIUI 12.5 запланированы на 8 февраля",
            "preview": "После успешного запуска в Китае Xiaomi Mi 11 готов выйти на мировой рынок – это произойдет 8 февраля. Однако официального сообщения о Mi 11 Pro пока нет. Но это не значит, что премьера нового...",
            "link": "https://hi-tech.news/smartphone/2689-globalnaja-premera-xiaomi-mi-11-i-miui-125-zaplanirovany-na-8-fevralja.html",
            "created_at": "2021-02-06T20:06:05.004547Z"
        },
        {
            "id": "037ad1d2-5326-4275-a051-db8471c9a3af",
            "category": "Интернет",
            "title": "Последнее обновление Telegram позволяет импортировать историю переписки с других приложений",
            "preview": "Telegram v7.4.0 теперь предоставит возможность перенести историю чатов из других приложений обмена сообщениями, таких как WhatsApp, Line и KakaoTalk. Этот шаг последовал после спорного изменения...",
            "link": "https://hi-tech.news/internet/2690-poslednee-obnovlenie-telegram-pozvoljaet-importirovat-istoriju-perepiski-s-drugih-prilozhenij.html",
            "created_at": "2021-02-06T20:06:05.002304Z"
        }
    }
]
```

### GET /hi_tech_news/search?title=COVID

Поиск новости по подстроке в заголовке.

**Return values**  
```json
{
    "data": [
        {
            "id": "4080e8e2-150e-485c-a109-ecaa6a13e626",
            "category": "Медицина",
            "title": "COVID Hunter – бесконтактный сканер для обнаружения коронавируса (видео)",
            "preview": "Американская компания Advanced Medical Solutions International заявила, что изобрела первый в мире бесконтактный сканер для обнаружения коронавируса с 99-процентной эффективностью. Так называемый...",
            "link": "https://hi-tech.news/med/2705-covid-hunter-beskontaktnyj-skaner-dlja-obnaruzhenija-koronavirusa-video.html",
            "created_at": "2021-02-06T09:02:13.705096Z"
        }
    ]
}
```

### GET /health

Проверка соединения с бд. Отправляет 500 status code, если нет связи с бд.   

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
