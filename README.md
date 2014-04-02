atap
=====

atap is a tool that allows you to run a command
a given number of times synchronously.

### Install

You will need:

  - Git
  - Go (tested with version 1.1.2)
  - Make (tested with version 3.81)

```bash
$ git clone https://github.com/hazbo/atap.git
$ cd atap
$ make
$ sudo make install
```

### Usage

atap requires one configuration file written in JSON. An example
looks like this:

```json
{
	"command" : "ls",
	"args" : [
		"-a"
	],
	"count" : 5
}
```

Then you simply pass through that file as the param when you run atap

	$ atap config.json

You may call this file what you wish. Your given command with arguments
will then be ran the number of times you specify. In the case above, five.

### License

[MIT](https://github.com/hazbo/atap/blob/master/LICENSE)