# Madge
outputs JSON

```bash
madge --json src/
```
example output
```json
{
    "src/index.ts":["src/utils.ts","src/config.ts"],
    "src/utils.ts":["src/helpers.ts"]
}
```