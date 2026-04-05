# Identity API Examples

## Create Guest Identity

### Request

```json
{
  "display_name": "henrique"
}
```

### Response

```json
{
  "player_id": "player-001",
  "display_name": "henrique",
  "claim_status": "guest",
  "device_token": "device-001"
}
```

## Resume Identity

### Request

```json
{
  "device_token": "device-001"
}
```

### Response

```json
{
  "player_id": "player-001",
  "display_name": "henrique",
  "claim_status": "guest"
}
```
