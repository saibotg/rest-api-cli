# rest-api-cli

## check

### simple check without auth

```shell
rest-api-cli check --url https://ipinfo.io/json --key country -w US -c DE
```

This command requests the URL **https://ipinfo.io/json** and reads in the resulting key **country**. When the key is "US" we will get a warning and when we get the value "DE" we will get a critical. All other values will result in OK. When the key does not exist we will get UNKNOWN as result.
