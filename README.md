# rest-api-cli

## check

### simple check without auth

```shell
rest-api-cli check --url https://ipinfo.io/json --key country --regex "^DE$" --severity crit
```

This command requests the URL **https://ipinfo.io/json** and reads in the resulting key **country**. When the key is "DE" the return value will be OK. When the key has any other value it will return CRITICAL.

### simple check with basic auth

```shell
rest-api-cli check --url https://ipinfo.io/json --key country --regex "^DE$" --severity warn --username monitor --password "123456"
```

This command is the same check, like the previous, but with username and password for basic authetification and the result is WARNING, when the key is not "OK".
