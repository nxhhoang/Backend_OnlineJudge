# Problem service

## API Schema
```
- POST /problem/add?problemId=...
- GET /problem/get/${problemId}/{static files}
```

### Problem directory structure
```
+-- problem.json
+-- statement.pdf
+-- checker
+-- tests/
|   +-- input/
|   |   +-- 1
|   |   +-- 2
|   |   +-- ...
|   +-- output/
|   |   +-- 1
|   |   +-- 2
|   |   +-- ...
```
### Sample of a problem.json
```json
{
  "ID": "...",
  "problem-id": 327480,
  "name": "Chơi đùa cùng bộ bài",
  "short-name": "playful-card-game",
  "tags": [
    "dp"
  ],
  "test-num": 103,
  "time-limit": 1000,
  "memory-limit": 268435456
}
```

## Current problems
- Only accept English statement (could use Vietnamese inside)