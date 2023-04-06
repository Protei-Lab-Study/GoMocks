#!/bin/bash
FILE_CERT_NAME=firebaseMockCert

openssl genrsa -out $FILE_CERT_NAME.key 2048
openssl req -new -x509 -sha256 -key $FILE_CERT_NAME.key -out $FILE_CERT_NAME.crt -days 365000 -subj "/C=RU/ST=SPB/L=SPB/O=Protei/OU=UC/CN=localhost"