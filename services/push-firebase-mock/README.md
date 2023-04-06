## push-firebase-mock

Сервис для обеспечения тестирования нагрузки. Является моком firebase нотификаций. 
Сервис обеспечивает прием нотификаций по https и проксирует их на websocket для верификации.

Сборка микросервиса
-----
```
 make compile
```

Run service in docker

-----

Чтобы запустить приложение в docker нужно выполнить:

1. Скачать docker образ с Gitlab

        $ sudo docker pull localhost:8443/uc/core/push-firebase-mock

2. Запустить контейнер

        $ sudo docker run -v /log/push-firebase-mock/:/log --rm --name push-firebase-mock --network host -d localhost:8443/uc/core/push-firebase-mock

   Сервис запустится на локальном IP машины и будет использовать порты сервиса.

   Также можно запустить контейнер и в своей подсети Docker, для этого нужно будет использовать проброс портов.

        $ sudo docker run -v /log/push-firebase-mock/:/log --rm --name push-firebase-mock -p <GO-TEMPLATE-SERVICE-PORT-NO>:<GO-TEMPLATE-SERVICE-PORT-NO> -d localhost:8443/uc/core/push-firebase-mock

   Для запуска docker образа с shell оболочкой нужно использовать команду

        $ sudo docker run -it -v /log/push-firebase-mock/:/log --rm --name push-firebase-mock --network host --entrypoint /bin/bash localhost:8443/uc/core/push-firebase-mock

Генерация документации
-----
```
 make doc
```

Удаление артефактов сборки микросервиса, документации
-----
```
 make clean
```
