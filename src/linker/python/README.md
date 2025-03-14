# Py deps
outputs JSON

```bash
pydeps myscript.py --json
```
example output
```json
{
    "myscript.py":["utils.py","config.py"],
    "utils.py":["helpers.py"]
}
```