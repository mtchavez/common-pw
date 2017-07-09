# Common Password Checker

common-pw is a go packge implementing a JSON API endpoint to check if a password
is in the top X most commonly used. Data is used from [Probable Word List][pwl]

## Usage

### Installing

Clone repo `git clone https://github.com/mtchavez/common-pw`

#### Large Datasets

Download the data from the [probably word list][pwl] repository and place
in the data directory. Currently `Top32Million-probable.txt` is not in version
control because of size.

### Server

Run the server on port 3000 with `go run *.go`

### Checking Passwords

Use cURL or something similar to post to `/validate` with a JSON body:

```json
{
  "password": "the password to check"
}
```

**Examples:**

```
$ curl -s -XPOST -H "Content-Type: application/json" http://localhost:3000/validate --data '{"password": "iloveyou"}' | jq
{
  "status": "OK",
  "top196": "true",
  "top3575": "true",
  "top95k": "true"
}
```
---
```
$ curl -s -XPOST -H "Content-Type: application/json" http://localhost:3000/validate --data '{"password": "annabelle"}' | jq
{
  "status": "OK",
  "top196": "false",
  "top3575": "true",
  "top95k": "true"
}
```
---
```
$ curl -s -XPOST -H "Content-Type: application/json" http://localhost:3000/validate --data '{"password": "!@#$%^&*("}' | jq
{
  "status": "OK",
  "top196": "false",
  "top3575": "false",
  "top95k": "true"
}
```
---
```
$ curl -s -XPOST -H "Content-Type: application/json" http://localhost:3000/validate --data '{"password": "Portlandia"}' | jq
{
  "status": "OK",
  "top196": "false",
  "top3575": "false",
  "top95k": "false"
}
```
---
```
$ curl -s -XPOST -H "Content-Type: application/json" http://localhost:3000/validate --data '{"password": "Portlandia"}' | jq
{
  "status": "OK",
  "top196": "false",
  "top32m": "true",
  "top3575": "false",
  "top95k": "false"
}
```
---
**A password is required**
```
$ curl -s -XPOST -H "Content-Type: application/json" http://localhost:3000/validate --data '{}' | jq
{
  "error": "a password must be provided",
  "status": "failed"
}
```

[pwl]: https://github.com/berzerk0/Probable-Wordlists
