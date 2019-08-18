# When Do

Schedule *when* to *do* things.

## General Workflow

- Schedule a CloudWatch Event every `PULSE_GAP` minutes.
- Each Event will trigger a Lambda to run ([src/when](src/when))
  - finds all tasks that need to run this minute
    - ???
  - for each task, fire a lambda to execute it

## Building

```
GOOS=linux GOARCH=amd64 go build -o src/when/when.go
zip main.zip main

GOOS=linux GOARCH=amd64 go build -o src/do/main.go
```