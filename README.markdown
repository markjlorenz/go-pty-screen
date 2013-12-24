# go/pty/screen
> Easily screen-share any terminal application

## TL;DR
![demo.gif](https://github.com/dapplebeforedawn/go-pty-screen/raw/master/demo.gif)

## Installation
 - One step! (try that with ruby)
 - Copy both files from `bin` to somewhere in your `$PATH`
 ```bash
 cp bin/* ~/dir/in/path
 ```

## Usage
### The Host
The host (running the to-be-shared applications):
```bash
go-pty-server
```

Now start some application termials with the `new` command:
```
> new our-tests bash 20 80
# new <alias> <command> <terminal rows> <terminal cols>
```

### The Clients
Any number of clients:
```bash
go-pty-client <ip of server>

# its not a bad idea to chain a call to `reset` at the end
go-pty-client <ip of server>; reset
```

## Scripting
go-pty-screen uses HTTP as it's inter-application communication protocol.  This makes it really easy to script.  You can even send it commands from your web browser (try loading `http://localhost:2000/servers` after starting some application terminals.

You can provide a startup script using the `--config-file` option, or by default `~/.go-pty-rc` will be loaded at boot.  For an example rc file, you can use `test/create-test-3.http` (e.g. `cp test/create-test-3.http ~/.go-pty-rc`)

## Project Structure
```
├── Makefile                    # `make build` to build the client and server
├── README.markdown             # you are here
├── bin
│   ├── go-pty-client           # an OSX 10.8 pre-compiled binary
│   └── go-pty-server           # an OSX 10.8 pre-compiled binary
├── clients
│   ├── go-pty-client.go        # manages a single application connection to the server
│   └── list.go                 # manages the clients list view of app choices
├── go-pty-client.go            # top level setup for the client
├── go-pty-server.go            # top level setup for the server
├── options
│   ├── client.go               # options parsing for the client
│   └── server.go               # options parsing for the server
├── pty_interface
│   └── pty_interface.go        # manages pseudo-termainal and it's contained application
├── servers
│   ├── key_server.go           # receives key strokes all connected clients to a single app
│   ├── screen_server.go        # sends screens state for a single app to all connected clients
│   ├── supervisor.go           # responds to HTTP requests about the state of the supervisor
│   └── supervisor_rc_loader.go # loads the startup config file
├── test
│   ├── create-test-2.http      # fixture data
│   ├── create-test-3.http
│   ├── create-test.http
│   ├── integration.rb          # an integration test `rspec integration.rb`
│   └── list-test.http
└── views
    ├── client
    │   └── list.go             # the clients list of available apps
    ├── supervisor
    │   ├── command.go          # the command window on the server
    │   └── list.go             # the list of app running on the server
    └── supervisor.go           # coordinates operation of `supervisor/command.go` and `supervisor/list.go`
```

## Building:
If you need to build from source (maybe to cross-compile for a 64 bit linux)
```
GOOS=linux GOARCH=amd64 make build
```

## Note:
Yes, another one of these.  The ruby versions were fun to code, but since a ruby application can not reliably be "distributed" as a stand alone application, it's nearly impossible for non-ruby devs to enjoy the fun.

More importantly, running ruby scripts rely on the version of ruby in the $PATH.  Since dev/pty/screen and dev/pty/vim both require ruby >= 2.0 this makes it impossible to use them while developing a <= 1.9 application.  Go-lang makes distributing standalone binaries stupid easy, so here we are!

After developing the Golang version of this application, wow is it a better tool for concurency than Ruby.  The code is significantly shorter, and easier to read and reason about thanks to go-routines and channels.

 - [dev/pty/screen](https://github.com/dapplebeforedawn/dev-pty-screen)
 - [dev/pty/vim](https://github.com/dapplebeforedawn/dev-pty-vim)
