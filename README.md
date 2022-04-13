# Taskr

A simple task time tracker for terminal built in golang

## Usage
-   taskr [command]

### Available Commands
-   `add [task_name]`        add task
-   `finish [task_name]`      finish task
-   `get [task_name]`         get task
-   `list`                              list tasks
-   `pause [task_name]`    pause task
-   `resume [task_name]`  resume task

### Example
`taskr add mytask`

## Installation

### From release
Download the latest binary
```bash
  chmod +x taskr
  sudo cp taskr /usr/local/bin
```

### Custom build
```bash
  go build -o taskr cmd/main.go
  chmod +x taskr
  sudo cp taskr /usr/local/bin
```
