# Go Callvis
outputs JSON

```bash
go list -json ./... | jq -s "."  > deps.json
```
example output
```json
{
    "nodes":["main","handler","databaseQuery"],
    "edges":[{
        "from":"main","to":"databaseQuery"
    }]
}
```