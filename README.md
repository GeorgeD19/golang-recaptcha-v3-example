# golang-recaptcha-v3-example

An example project where a web app using recaptcha v3 submits a token to a Golang server which is able to authenticate that token.

## Get Started

Create a .env file with the following content

```env
RECAPTCHA_SECRET_KEY=""
```

Fill in your secret, ensuring during setup you add localhost to accepted domains if doing local.

First start the go server

```bash
go run main.go
```

Then start the web app in another terminal

```bash
npm start
```

Open in your browser the URL http-server offers and add the url param reCAPTCHA_site_key passing in your site key.

For example: [http://localhost:8081/?reCAPTCHA_site_key={{site_key}}](http://localhost:8081/?reCAPTCHA_site_key={{site_key}})
