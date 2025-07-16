
## API Schema
- POST /problem/add?problemId=...&?packageId=...
- GET /problem/get/${problemId}/ 

#### Problem directory structure
```
+-- problem.json // simple version with only needed properties
+-- statement.pdf // haven't implemented yet
+-- checker // haven't implemented yet
+-- tests/
|   +-- input/
|   |   +-- 01
|   |   +-- 02
|   |   +-- ...
|   +-- output/
|   |   +-- 01
|   |   +-- 02
|   |   +-- ...
```
#### Sampe of a problem.json
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
