# srv
A simple, non-production static file server.

## Simple
```
$ srv
serving . on :8080
```

## With options
```
$ srv -dir ./my/directory -port 1618
```

## Help
```
$ srv -h

srv: a simple static file server

Usage:

  -dir string
        directory to server (default ".")
  -port string
        static file server port (default "8080")
```
