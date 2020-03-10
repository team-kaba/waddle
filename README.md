# waddle

`waddle` is a simple task runner based on json syntax.
You can define and describe your tasks in json file `waddle.json` and execute them with `waddle` command.

`waddle.json` is like below.

```json
{
	"tasks": [
		{
			"hello": {
				"cmd": "echo hello",
				"comment": "say hello"
			}
		},
		{
			"goodbye": {
				"cmd": "echo goodbye",
				"comment": "say good bye"
			}
		}
	]
}
``` 

