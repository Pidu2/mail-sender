## Env vars sender
* `APP_MAIL`: Gmail Address to send mails with
* `APP_PASSWORD`: The application specific password for this address
* `APP_PORT`: Port to listen to for gRPC

## Docker
```bash
cd sender/
docker build . -t pidu2/mail-sender:<TAG>
```
