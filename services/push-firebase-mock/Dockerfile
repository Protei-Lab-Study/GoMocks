FROM golang:1.15.15-alpine3.14 as builder

#WORK_PATH - рабочкая директория
#TERM - вывод в консоль
ENV WORK_PATH=/home/uc
ENV TERM=xterm

#Создаем том для логов
VOLUME [ "/tmp/push-firebase-mock" ]

#Ставим требуемые пакеты
RUN apk update &&\
    apk upgrade &&\
    apk add --no-cache make &&\
    apk add --no-cache git &&\
    apk add --no-cache openssh

#Объявляем рабочую директорию
WORKDIR ${WORK_PATH}

#Копируем содержимое локальной папки в docker образ
COPY /services/push-firebase-mock ${WORK_PATH}/services/push-firebase-mock
COPY /services/golang/ ${WORK_PATH}/services/golang
COPY /services/.ssh/ /root/.ssh/

RUN chmod 0600 /root/.ssh/* &&\
    mkdir ~/tmp &&\
    git config --global url."ssh://git@git.protei.ru/".insteadOf "https://git.protei.ru/"
COPY /services/configs/uc-config.yaml /home/protei/config/Protei-UC/uc-config.yaml


#Объявляем директорию для сервиса push-firebase-mock
WORKDIR ${WORK_PATH}/services/push-firebase-mock

#Запускаем сборку bin для push-firebase-mock
RUN go env -w GO111MODULE=on GONOPROXY=git.protei.ru GONOSUMDB=git.protei.ru GOPROXY="https://repo.protei.ru/repository/go-modules" &&\
    cd ./src || true &&\
    go get -v -d . || true &&\
    cd ../ || true &&\
    make compile

FROM alpine:latest
WORKDIR /home/uc/services/push-firebase-mock/

#Ставим требуемые пакеты
RUN apk update &&\
    apk upgrade &&\
    apk add --no-cache bash mc

COPY --from=builder /home/uc/services/push-firebase-mock/push-firebase-mock /home/uc/services/push-firebase-mock/
COPY /services/configs/uc-config.yaml /home/protei/config/Protei-UC/uc-config.yaml
COPY /services/push-firebase-mock/config/logger.yaml /home/uc/services/push-firebase-mock/config/logger.yaml

#Команда запуска сервиса push-firebase-mock
ENTRYPOINT [ "./push-firebase-mock"]
