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

### check scalar value with

```shell
rest-api-cli check --url https://ipinfo.io/json --key postal --warning "50000" --critical "70000"
```

Critical if value is over 70000, else warn of over 50000. Also critical when < 0.

## Threshold and Ranges

Source: [Nagios Plugins Development Guidelines](https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT)

The scalar value is checked against the range which is defined in critical or warning. If the value is part of the defined range, the check will result in the corresponding status. 

### Rules and Syntax

```text
[@]start:end
```

* start ≤ end
* start and ":" is not required if start=0
* if range is of format "start:" and end is not specified, assume end is infinity
* to specify negative infinity, use "~"
* alert is raised if metric is outside start and end range (inclusive of endpoints)
* if range starts with "@", then alert if inside this range (inclusive of endpoints)

### Example ranges

| Range definition | Generate an alert if x... |
| --- | -- |
| ```10``` | < 0 or > 10, (outside the range of {0 .. 10}) |
| ```10:``` | < 10, (outside {10 .. [infinity]}) |
| ```~:10```| > 10, (outside the range of {[negativ infinity] .. 10}) |
| ```10:20``` | < 10 or > 20, (outside the range of {10 .. 20}) |
| ```@10:20``` | ≥ 10 and ≤ 20, (inside the range of {10 .. 20}) |

