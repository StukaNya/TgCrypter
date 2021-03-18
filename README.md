# TgCrypter
TgCrypter is a secure data encryption storer with Telegram bot and HTTP REST API interaction.

## Installation
With a correctly configured Go toolchain:
```bash
go get -u github.com/StukaNya/TgCrypter
```

## Structure
* AES and SHA256 algorhitms for secure encryption user data.
* Using telegram-bot-api library and REST API for comfortable user interaction.
* MVC architecture style service with interface injection.
* PostgreSQL as user data (pin code and files) database.
* TODO: Awesome js-front and Amazon S3 with Minio client as file storer...
