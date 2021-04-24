# vezdecode-webcrawler
Webcrawler (information security 40)

Поисковик информации по веб странице

Реализован поиск телефона, емейла, социальных сетей(вконтакте и фейсбук) и IP адреса.

Написан на языка Go , требуется Go не ниже 1.14

Для запуска необходимо перейти в корневую папку проекта и выполнить 
`go run . [-args]`

Возможны следующие сценарии использования:

#### Передача HTML непосредственно в программу:
```
go run . "<body><div>tr@ytr.com</div></body>"
```
Результат:
```
RESULT:
Visited links (0):
Founded entities:
entity name: Email , value: tr@ytr.com
```

#### Передача URL адреса:

```
go run . -url=https://trukach000.ru/info/summary
```
Результат:
```
RESULT:
Visited links (1):
ulr: https://trukach000.ru/info/summary
Founded entities:
entity name: phone , value: +79131538398
entity name: Email , value: trukach000@gmail.com
```

#### Передача имени файла содержащего список URL'ов (каждый в отдельной строке)
```
go run . -file=files/urls.txt
```
Результат:
```
RESULT:
Visited links (3):
ulr: https://google.com
ulr: https://habr.com/ru/post/440400/
ulr: https://company.habr.com/?_ga=2.173711777.322719268.1619255259-610910979.1548570463
Founded entities:
entity name: Email , value: 830576edd4b7478086093f693a5a0df5@s.tmtm.ru
entity name: IP address , value: 189.057.125.126
entity name: IP address , value: 147.082.188.197
entity name: Social network , value: https://vk.com/share.php
entity name: Social network , value: https://vk.com/habr
entity name: Social network , value: https://vk.com/rtrg
entity name: Social network , value: https://www.facebook.com/sharer/sharer
entity name: Social network , value: https://www.facebook.com/habrahabr
entity name: Social network , value: https://www.facebook.com/tr
entity name: phone , value: +7 499 653-59-61
entity name: Email , value: info@habr.com
entity name: Email , value: corp@habr.team
entity name: Email , value: content@habr.team
entity name: Email , value: hr@habr.team
entity name: Email , value: adv@habr.team
entity name: Email , value: support@freelance.habr.com
```

#### Дополнительные аргументы
```
-r - включает режим рекурсии (работает только при получении контента по URL)
    Рекурсия ограничивается доменом базового URL
-d - включает дополнительные логгирование
```