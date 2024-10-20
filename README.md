# Tasks

A classic Todo app, but with Go, for the command-line. It's a simple CLI app to manage tasks. You can add, list, edit, remove and clean tasks. It stores everything in the local machine only as a CSV file.

Check the Command Mapping and Data Schema [here](./assets/tasks.png).

## Installation

1. Clone the repository

```bash
git clone https://github.com/PratikDev/tasks
```

2. Go to the project directory

```bash
cd tasks
```

3. Get the dependencies

```bash
go mod tidy
```

4. Build the binary

```bash
go build .
```

## Commands

### Add

- `./tasks add "task"`: Add a new task with title "task" and status "Pending"
- `./tasks add "task" -w`: Add a new task with title "task" and status "Working"
- `./tasks add "task" -d`: Add a new task with title "task" and status "Done"
- `./tasks add "task" -c`: Add a new task with title "task" and status "Canceled"

### List

- `./tasks ls`: List all tasks with "Working" status
- `./tasks ls -a`: List all tasks
- `./tasks ls -w`: List all tasks with "Working" status
- `./tasks ls -d`: List all tasks with "Done" status
- `./tasks ls -c`: List all tasks with "Canceled" status

### Edit

- `./tasks edit <id> -t "new title"`: Edit the task with id <id> and change the title to "new title"
- `./tasks edit <id> -w`: Edit the task with id <id> and change the status to "Working"
- `./tasks edit <id> -d`: Edit the task with id <id> and change the status to "Done"
- `./tasks edit <id> -c`: Edit the task with id <id> and change the status to "Canceled"
- `./tasks edit <id> -t "new title" -d`: Edit the task with id <id> and change the title to "new title" and the status to "Done"

### Remove

- `./tasks rm <id>`: Remove the task with id <id>

### Clean

- `./tasks prune`: Remove all tasks with "Cancelled" status
- `./tasks prune -d`: Remove all tasks with "Done" status
- `./tasks prune -w`: Remove all tasks with "Working" status
- `./tasks prune -p`: Remove all tasks with "Pending" status

### Help

- `./tasks [command] -h`: Show the help for any command

## Technologies

- **Golang** - As the main language
- [Cobra](https://github.com/spf13/cobra) - To create the CLI
- [Cobra-CLI](https://github.com/spf13/cobra-cli) - To initialize the commands
- [Timediff](https://github.com/mergestat/timediff) - To show the task creation time in a human-readable format
- [Heredoc](https://github.com/MakeNowJust/heredoc) - To create the help messages with good indentation
- [text/tabwriter](https://pkg.go.dev/text/tabwriter) - To show the tasks in a tabular format
- [encoding/csv](https://pkg.go.dev/encoding/csv) - To handle the CSV file
