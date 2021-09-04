# GoRemind
A full stack CLI reminder app with a notification feature. The Go application uses command line arguments to provide full features including create, edit, update and deletion of reminders

## Components
- Go Client
The command line client that accepts arguments from the command and parses them into appropriate commands

- Go Server
The main application server that provides endpoints for the client application to use

- Nodejs/Express notifier
A notifier service using the node-notifier package that informs the user when a reminder is due