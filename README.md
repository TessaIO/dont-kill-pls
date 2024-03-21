## dont-kill-pls
This package monitors a Golang application and if its memory exceeds a specific threshold it would be stoped.
The reason behind this is to avoid OOM kill for the application since such the SIGKILL signal cannot be handled
gracefully.

# Examples
You can run examples by:
```
cd examples
go run ./...
```
