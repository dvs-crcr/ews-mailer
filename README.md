[![Go Report Card](https://goreportcard.com/badge/github.com/dvs-crcr/ews-mailer)](https://goreportcard.com/report/github.com/dvs-crcr/ews-mailer)

# ews-mailer

> A tiny application for sending automated plaintext messages via EWS w/o any external dependencies.

## Configuration
### Environment variables

```dotenv
EWS_URL=https://mail.domainname.com/EWS/Exchange.asmx
EWS_DOMAIN=DOMAINNAME
EWS_USERNAME=username.io
EWS_PASSWORD=superSecretPass$
EWS_FROM=username.io@domainname.com
```

### Flags
```text
-to      - Comma-separated recipients (required)
-subject - Email subject (required)
-body    - Plaintext message body (required)
```

You can also customize the `Request timeout` and `User-Agent` header at build time by editing const in the `main.go` file.

## Usage
Install the application

```shell
go install github.com/dvs-crcr/ews-mailer
```

Fill arguments and run the application using the command below.

```shell
ews-mailer \
  -to "alice@domainname.com,bob@domainname.com" \
  -subject "System Notification" \
  -body "This message was sent automatically from the EWS-Mailer."
```
