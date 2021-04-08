# splatnet2statink-docker

https://github.com/frozenpandaman/splatnet2statink

## generate config.txt using docker

```sh
$ touch config.txt
$ docker run -it -v ${PWD}/config.txt:/splatnet2statink/config.txt nu50218/splatnet2statink:latest
```

## Examples

### Example 1

- mount local config.txt
- upload new salmon records

```sh
$ docker run -v ${PWD}/config.txt:/splatnet2statink/config.txt nu50218/splatnet2statink:latest --salmon -r
```

### Example 2

- generate config.txt from environment variables at startup
- upload new battle records every 15min

```sh
$ docker run \
-e USE_ENVIRONMENT_VARIABLE=1 \
-e API_KEY=key123 \
-e COOKIE=cookie456 \
-e SESSION_TOKEN=token789 \
-e USER_LANG=en-US \
nu50218/splatnet2statink:latest -M 900
```

this will generate

```json
{
    "api_key": "key123",
    "cookie": "cookie456",
    "session_token": "token789",
    "user_lang": "en-US"
}
```