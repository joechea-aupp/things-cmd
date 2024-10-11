# Thing3 linux command client

This project is simply providing a way for you to interact with Things3 application from commandline.
The only thing that it can do is to add a task to your Things3 inbox. It added by sending an email to your Things3 email address.

## Prerequisites
- Thing3 email address enabled. See here: https://culturedcode.com/things/support/articles/2908262/
- An email address to send the task from. If you are using gmail, please configure App Password. See here: https://support.google.com/accounts/answer/185833?hl=en
- Sqlite3 installed. If not, you can install it by running `sudo apt-get install sqlite3`

## Build
You need to have go installed. If not, you can install it by running `sudo apt-get install golang-go`
```
go build -o ./build/add
```

If you don't want to build it, you can download the binary from the releases page.

## Usage
For the first time, you need to configure email address (credentials and server) and Things3 email address.

```
add -host smtp.gmail.com -port 587 -u username@gmail.com -p AppPassword
```

**Note:** To update the configuration, you can run the same command with different values.


Then you can add a task by running the following command:
```
add -t "Task Title" -d "Task Description"
```

`-d` is optional.


Default database path is `~/.local/share/things-cmd/things.db`
**Note:** credential store in sqlite3 database is not encrypted. Please be aware of that.
