# firebase-auth-example

## Setup

**1. .env を作成**

```
VITE_API_KEY=apiKey
VITE_AUTH_DOMAIN=authDomain
VITE_PROJECT_ID=projectId
VITE_STORAGE_BUCKET=storageBucket
VITE_MESSAGING_SENDER_ID=messaginSenderId
VITE_APP_ID=appId
```

**2. credentials.json をダウンロードし、`backend/credentials.json` に配置**

**3. 起動**

frontend

```shell
$ yarn install
$ yarn dev
```

backend

```shell
$ go run .
```
