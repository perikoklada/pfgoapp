# [GoApp REST API](#goapp)

## GET /goapp

Returns an example websocket page.

## GET /goapp/ws

The message sent by the server containing the counter value:
```json
{
  "iteration": 10,
  "output": "0x2221E3bC64"
}
```

The message sent by the client to reset the counter to 0 is currently empty object :

```json
{}
```

## [GET /goapp/health](#health)
| _health_ |

Returns an HTTP code for health status. Currently it only returns 200 status code when the endpoint is accessible.
