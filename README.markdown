 go/pty/screen
> Easily screen-share any terminal application

## TL;DR
![demo.gif](https://github.com/dapplebeforedawn/go-pty-screen/raw/master/demo.gif)

## Installation
 - One step! (try that with ruby)
 - Copy server, client (and optionally tunnel) from `bin` to somewhere in your `$PATH`
 ```bash
 cp bin/* ~/dir/in/path
 ```

## Usage
### The Host
The host (running the to-be-shared applications):
```bash
go-pty-server
```

Now start some application terminals with the `new` command:
```
> new our-tests bash 20 80
# new <alias> <command> <terminal rows> <terminal cols>
```

### The Clients
Any number of clients:
```bash
go-pty-client <ip of server>  # server ip is option if on a local network
```

### The Tunnel
The tunnel is a quick network pipe that can be used to easily transfer files, or use program on the client that are no available on the server.  Some examples:
* client: `echo "come here watson" | go-pty-tunnel`
* server: `go-pty-tunnel`
* output (server): `come here watson`

* client: `go-pty-tunnel | pbcopy`
* server: `pbpaste | go-pty-tunnel`
* output: the clients clipboard now has the same contents as the servers clipboard.

* client: `cowsay "midwestern commit message" | go-pty-tunnel`
* server: `go-pty-tunnel | git commit -F -`  (the server does not have the cowsay program installed)
* output: the commit message now has a cow.

## Scripting
go-pty-screen uses HTTP as it's inter-application communication protocol.  This makes it really easy to script.  You can even send it commands from your web browser (try loading `http://localhost:2000/servers` after starting some application terminals.

You can provide a startup script using the `--config-file` option, or by default `~/.go-pty-rc` will be loaded at boot.  For an example rc file, you can use `test/create-test-3.http` (e.g. `cp test/create-test-3.http ~/.go-pty-rc`)

## Project Structure
```
├── ./Makefile                            # run `make` and the `Makefile` will do the rest
├── ./README.markdown                     # you are here
├── ./bin                                 # pre-compiled OSX binaries
│   ├── ./bin/go-pty-client
│   ├── ./bin/go-pty-server
│   └── ./bin/go-pty-tunnel
├── ./clients
│   ├── ./clients/go-pty-client.go
│   └── ./clients/list.go
├── ./go-pty-client.go                    # `main` package for the client
├── ./go-pty-server.go                    # `main` package for the server
├── ./go-pty-tunnel.go                    # `main` package for the tunnel
├── ./options                             # options files
│   ├── ./options/client.go
│   ├── ./options/server.go
│   └── ./options/tunnel.go
├── ./pty_interface                       # setup and manage the underlying ptys
│   └── ./pty_interface/pty_interface.go
├── ./servers
│   ├── ./servers/key_server.go           # recieves keystrokes from the clients
│   ├── ./servers/screen_server.go        # serves the current pty screen state to clients
│   ├── ./servers/supervisor.go           # handles HTTP requests for new ptys, and listing of current ptys
│   └── ./servers/supervisor_rc_loader.go # loads the rc file and sends it to the supervisor via HTTP
├── ./test                                # the `.http` files here can be used as bases for your `~/.go-pty-rc`
│   ├── ./test/create-test-2.http
│   ├── ./test/create-test-3.http
│   ├── ./test/create-test.http
│   ├── ./test/integration.rb
│   └── ./test/list-test.http
├── ./views
│   ├── ./views/client
│   │   └── ./views/client/list.go        # the clients menu
│   ├── ./views/supervisor
│   │   ├── ./views/supervisor/command.go # the command window
│   │   └── ./views/supervisor/list.go    # the list of running ptys
│   └── ./views/supervisor.go
└── ./zeroconf                            # client and server for interacing with dns-sd  (Bonjour)
    ├── ./zeroconf/client.go
    └── ./zeroconf/server.go
```

## Building
If you need to build from source (maybe to cross-compile for a 64 bit linux)
```
GOOS=linux GOARCH=amd64 make build
```

## Tests
The interaction between these two apps is a little hard to test.  The included integration test is an attempt, but not a great one.  To run the spec:

```bash
cd /path/to/repo/root
rspec test/integration.rb
```

but don't expect them all to pass.

## Note
Yes, another one of these.  The ruby versions were fun to code, but since a ruby application can not reliably be "distributed" as a stand alone application, it's nearly impossible for non-ruby devs to enjoy the fun.

More importantly, running ruby scripts rely on the version of ruby in the $PATH.  Since dev/pty/screen and dev/pty/vim both require ruby >= 2.0 this makes it impossible to use them while developing a <= 1.9 application.  Go-lang makes distributing standalone binaries stupid easy, so here we are!

After developing the Golang version of this application, wow is it a better tool for concurrency than Ruby.  The code is significantly shorter, and easier to read and reason about thanks to go-routines and channels.

 - [dev/pty/screen](https://github.com/dapplebeforedawn/dev-pty-screen)
 - [dev/pty/vim](https://github.com/dapplebeforedawn/dev-pty-vim)
