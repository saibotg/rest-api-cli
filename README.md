# rest-api-cli

## check

### simple check without auth

```shell
rest-api-cli check --url https://ipinfo.io/json --key country --regex "^DE$"
```

This command requests the URL **https://ipinfo.io/json** and reads in the resulting key **country**. When the key is "DE" the return value will be OK. When the key has any other value it will return CRITICAL.

### simple check with basic auth

```shell
rest-api-cli check --url https://ipinfo.io/json --key country --regex "^DE$" \ --username monitor --password "123456"
```

This command is the same check, like the previous, but with username and password for basic authentification.

### simple check with auth file

```shell
rest-api-cli check --url https://ipinfo.io/json --key country --regex "^DE$" \ --auth-file /etc/rest-api-cli/auth.cfg"
```

#### **`auth.cfg.js`**
```cfg
Basic bW9uaXRvcjoxMjM0NQ==
```

This command is the same check, like the previous, the content of the basic authentication header in the auth.cfg file.


