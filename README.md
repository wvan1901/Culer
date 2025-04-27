# Culer
This is a cli program that takes in input from stdin and wraps text with
ascii escape code coloring and outputs it to stdout.

# Table Of Contents
- [How to install](#how-to-install)
- [How to run](#how-to-run)
- [Color Options](#color-options)
- [Questions](#questions)

# How to install
Run this command to install Culer
```bash
go get github.com/wvan1901/Culer
go install github.com/wvan1901/Culer
```

# How to run
Culer has defaults set in place so it works without any configuration.
If you just run the program nothing will be displayed since Culer requires an input.
Culer is meant to be used with other programs. If I want to run a server and want
Culer to wrap the programs output then I can do so by piping the running server like so.
```bash
# This pipes server logs to culer, then culer prints the output into stdout
go run server.go | Culer
```
Since we can pipe a string input to culer it doesn't really matter if we pipe a server
or not but this is its main use case. If you want to use Culer to color a file feel free
to do so!
```bash
# Pipe a file to Culer
cat server.log | Culer
```

# Color Options
Culer has a multitude of flags to customize log level text string and color.
Here are the following color options supported:
```
black,red,green,yellow,blue,
magenta,cyan,light-gray,dark-gray,light-red,
light-green,light-yellow,light-blue,light-magenta,light-cyan,
white
```
# Questions
## Why did I build this
When running a API server we typically have lots of log statements. Its difficult
to tell the difference between log levels or to find a certain text within the logs.
Culer fixes this by taking in text and looking for the log level or custom text and
wraps this text with colors. This allows for easy differenciation by using colors to
highlight the information.

## Why not just use x y and z
This problem has been solved a multitude of ways. I even created a colored slog
handler that solves this exact same issue. The issues is that not all projects use
a slog handler or the team just doesn't want an unnesssary dependency. Instead of
creating a solution inside the projecy I decided to also create a solution that works
outside the projects. If this doesn't suit your needs then thats okay this project is
tailored to suit my specific set up and requirements.


