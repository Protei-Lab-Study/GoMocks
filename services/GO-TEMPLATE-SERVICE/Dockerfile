FROM golang:1.15.15-alpine3.14 as builder

#WORK_PATH - рабочкая директория
#TERM - вывод в консоль
ENV WORK_PATH=/home/uc
ENV TERM=xterm

#Создаем том для логов
VOLUME [ "/tmp/<GO-TEMPLATE-SERVICE>" ]

#Ставим требуемые пакеты
RUN apk update &&\
    apk upgrade &&\
    apk add --no-cache make &&\
    apk add --no-cache git &&\
    apk add --no-cache openssh

#Объявляем рабочую директорию
WORKDIR ${WORK_PATH}

#Копируем содержимое локальной папки в docker образ
COPY /services/<GO-TEMPLATE-SERVICE> ${WORK_PATH}/services/<GO-TEMPLATE-SERVICE>
COPY /services/golang/ ${WORK_PATH}/services/golang
COPY /services/.ssh/ /root/.ssh/

RUN chmod 0600 /root/.ssh/* &&\
    mkdir ~/tmp &&\
    git config --global url."ssh://git@git.protei.ru/".insteadOf "https://git.protei.ru/"
COPY /services/configs/uc-config.yaml /home/protei/config/Protei-UC/uc-config.yaml


#Объявляем директорию для сервиса <GO-TEMPLATE-SERVICE>
WORKDIR ${WORK_PATH}/services/<GO-TEMPLATE-SERVICE>

#Запускаем сборку bin для <GO-TEMPLATE-SERVICE>
RUN go env -w GO111MODULE=on GONOPROXY=git.protei.ru GONOSUMDB=git.protei.ru GOPROXY="https://repo.protei.ru/repository/go-modules" &&\
    cd ./src || true &&\
    go get -v -d . || true &&\
    cd ../ || true &&\
    make compile

FROM alpine:latest
WORKDIR /home/uc/services/<GO-TEMPLATE-SERVICE>/

#Ставим требуемые пакеты
RUN apk update &&\
    apk upgrade &&\
    apk add --no-cache bash mc

COPY --from=builder /home/uc/services/<GO-TEMPLATE-SERVICE>/<GO-TEMPLATE-SERVICE> /home/uc/services/<GO-TEMPLATE-SERVICE>/
COPY /services/configs/uc-config.yaml /home/protei/config/Protei-UC/uc-config.yaml
COPY /services/<GO-TEMPLATE-SERVICE>/config/logger.yaml /home/uc/services/<GO-TEMPLATE-SERVICE>/config/logger.yaml

#Команда запуска сервиса <GO-TEMPLATE-SERVICE>
ENTRYPOINT [ "./<GO-TEMPLATE-SERVICE>"]
