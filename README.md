# device grant [RFC 8628](https://www.rfc-editor.org/rfc/rfc8628)

A simple oauth2-enabled device grant server which implements RFC 8628.

```
go build -o device_grant.exe cmd
```

A sample client id is generated in the startup logs - see below for usage.

# endpoints

Send device authorization requests using your generated client id.
```
Content-Type application/x-www-form-urlencoded
POST http://127.0.0.1:8081/device_authorization?client_id=falling-sunset
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
POST http://127.0.0.1:8081/access_token?grant_type=urn:ietf:params:oauth:grant-type:device_code&client_id=falling-sunset&device_code=9b60c30f-5272-4bda-8fc0-562c3c8cdf57
```

Sample response:
```
HTTP 200 OK
Content-Type application/json
{
    "access_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6ImU5YTU0NTFhLWU2MjAtNGE4OS1iMTYxLWEwMjQ5MzY3OTQ2MCIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiaW1wb3J0YW50LXJlc291cmNlLXNlcnZlciJdLCJleHAiOjE2NjUwMDY0ODgsImlhdCI6MTY2NTAwMjg4OCwiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS5pby9qd2tzIiwianRpIjoiMzdlMjFmMjMtNmFhNC00NzczLWE1MDAtZjI1MWM5MTM0NzkzIiwibmJmIjoxNjY1MDAyNjgxLCJzdWIiOiJmYWxsaW5nLXN1bnNldCM5YjYwYzMwZi01MjcyLTRiZGEtOGZjMC01NjJjM2M4Y2RmNTcifQ.NgzGQnaMC_OJHTTCrkz2lbl8UMWsdw1IAOnYQpAFIAfSkATG1WDdY9swTlSNjJZeko1q6juAETxcysOBRZyrv828b8vbWDH7st9f2D1OXZrTQilxRKbB8EQGoS4Pxg5pk4Km5W0ldyChVHaIREegXI4y5oR2DCPygKM-YMgMxtgwS3RnokXXqKBM-C2uppQier4XUqd1esxHSECSrTaK1XKvqMJ_-FiajRzW3uNG-i_JY-UvMrU1j9lh1pXYsVkqHHyNo4mFIqTWfplG5CFLkXlANMFclfksVCOzkzivbnW8tB-Tc_JYifs72ZXpCzW6K2mo0SRyYoNqHpuGvP08Bw",
    "token_type": "bearer",
    "expires_in": 3600
}
```

See tests inside for more details and examples!