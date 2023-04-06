## <GO-TEMPLATE-SERVICE>

Описание того, зачем необходим данный сервис

Сборка микросервиса
-----
```
 make compile
```

Run service in docker

-----

Чтобы запустить приложение в docker нужно выполнить:

1. Скачать docker образ с Gitlab

        $ sudo docker pull localhost:8443/uc/core/<GO-TEMPLATE-SERVICE-IMAGE-ID>

2. Запустить контейнер

        $ sudo docker run -v /log/<GO-TEMPLATE-SERVICE>/:/log --rm --name <GO-TEMPLATE-SERVICE> --network host -d localhost:8443/uc/core/<GO-TEMPLATE-SERVICE-IMAGE-ID>

   Сервис запустится на локальном IP машины и будет использовать порты сервиса.

   Также можно запустить контейнер и в своей подсети Docker, для этого нужно будет использовать проброс портов.

        $ sudo docker run -v /log/<GO-TEMPLATE-SERVICE>/:/log --rm --name <GO-TEMPLATE-SERVICE> -p <GO-TEMPLATE-SERVICE-PORT-NO>:<GO-TEMPLATE-SERVICE-PORT-NO> -d localhost:8443/uc/core/<GO-TEMPLATE-SERVICE-IMAGE-ID>

   Для запуска docker образа с shell оболочкой нужно использовать команду

        $ sudo docker run -it -v /log/<GO-TEMPLATE-SERVICE>/:/log --rm --name <GO-TEMPLATE-SERVICE> --network host --entrypoint /bin/bash localhost:8443/uc/core/<GO-TEMPLATE-SERVICE-IMAGE-ID>

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
