# device grant [RFC 8628](https://www.rfc-editor.org/rfc/rfc8628)

A simple oauth2-enabled device grant server which implements RFC 8628.
```
go build -o device_grant.exe cmd
```

A sample client id is generated in the startup logs - see below for usage.
```
created default public client_id: lingering-sound
```

# endpoints

Send device authorization requests using your generated client id.
```
Content-Type application/x-www-form-urlencoded
POST http://127.0.0.1:8081/device_authorization?client_id=lingering-sound
```

Sample response:
```
HTTP 200 OK
Content-Type application/json
{
    "device_code": "9b60c30f-5272-4bda-8fc0-562c3c8cdf57",
    "user_code": "FDUX1VQL",
    "verification_uri": "http://127.0.0.1:8081/device",
    "verification_uri_complete": "http://127.0.0.1:8081/device?user_code=FDUX1VQL",
    "expires_in": 600
}
```

Allow users to authorize devices.
```
Content-Type application/x-www-form-urlencoded
POST http://127.0.0.1:8081/device?user_code=FDUX1VQL
```

Generate access JWTs for clients post user authorization.
```
Content-Type application/x-www-form-urlencoded
POST http://127.0.0.1:8081/access_token?grant_type=urn:ietf:params:oauth:grant-type:device_code&client_id=lingering-sound&device_code=9b60c30f-5272-4bda-8fc0-562c3c8cdf57
```

Sample response:
```
HTTP 200 OK
Content-Type application/json
{
    "access_token": "eyJhbGc ... qHpuGvP",
    "token_type": "bearer",
    "expires_in": 3600
}
```

See tests inside for more details and examples!