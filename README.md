# google-access-token

Utilitary for generating access/refresh tokens.

## Installation

```bash
$ go get -u github.com/tpisani/google-access-token
```

## Usage

```bash
$ google-access-token <client_id> <client_secret>
```

## Walkthrough

Access the [Google Developers Console](https://console.developers.google.com/).
![Google Developers Console](walkthrough/developers-console.png)

Make sure Google Drive API is active.
![Google Drive API](walkthrough/apis-drive.png)
![Google Drive API active](walkthrough/drive-api-active.png)

Create a new OAuth client.
![Google Drive API active](walkthrough/oauth2-client-selection.png)
![Google Drive API active](walkthrough/oauth2-client-creation.png)

Copy client ID and secret.
![Google Drive API active](walkthrough/oauth2-client-credentials.png)

Run `google-access-token` with these credentials.
```bash
$ google-access-token <client_id> <client_secret>
```
