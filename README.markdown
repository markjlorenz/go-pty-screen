# go/pty/screen
> Easily screen-share any terminal application

## Installation
 - One step! (try that with ruby)
 - Copy both files from `bin` to somewhere in your `$PATH`
 ```bash
 cp bin/* ~/dir/in/path
 ```

## Usage
The host (running the to-be-shared application):
```bash
go-pty-server application-name

# e.g.
go-pty-server vim  # let's pair program
go-pty-server --key_port=2002 --screen_port=2003  "bundle exec guard" # so you can see the test runner

# Note, when flag arguments are used, the application name _must_ come last
```

Any number of clients:
```bash
go-pty-client <ip of server>

# e.g.
go-pty-client 192.168.100.1   # shows us vim
go-pty-client --key_port=2002 --screen_port=2003  192.168.100.1  # shows us the test runner

# its not a bad idea to chain a call to `reset` at the end
go-pty-client <ip of server>; reset
```

## Project Structure
```
├── README.markdown
├── go-pty-client.go      # the client application
├── go-pty-server.go      # the server application
├── options
│   ├── client.go         # options parser for the client
│   └── server.go         # options parser for the server
├── pty_interface
│   └── pty_interface.go  # used by the server to interact with the pty
└── servers
    ├── key_server.go     # network connection for the stdin of the pty
    └── screen_server.go  # network connection for the stdout of the pty
```

## Note:
Yes, another one of these.  The ruby versions were fun to code, but since a ruby application can not reliably be "distributed" as a stand alone application, it's nearly impossible for non-ruby devs to enjoy the fun.

More importantly, running ruby scripts rely on the version of ruby in the $PATH.  Since dev/pty/screen and dev/pty/vim both require ruby >= 2.0 this makes it impossible to use them while developing a <= 1.9 application.  Go-lang makes distributing standalone binaries stupid easy, so here we are!

After developing the Golang version of this application, wow is it a better tool for concurency than Ruby.  The code is significantly shorter, and easier to read and reason about thanks to go-routines and channels.

 - [dev/pty/screen](https://github.com/dapplebeforedawn/dev-pty-screen)
 - [dev/pty/vim](https://github.com/dapplebeforedawn/dev-pty-vim)
