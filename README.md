![logo](https://i.imgur.com/KGKYp3F.png)



# Garbage Lisp  
Well it works but isn't good.

**Go** is definitely not the best language to do this, so it came out kinda hacky with **Go** Interfaces.

I coded this to learn about Interpreters, so this is definitely not a good example nor it works like most lisps so far.

I tried following this guide http://norvig.com/lispy.html, yet I had to change a lot of stuff because **Go** isn't Dynamically
typed like **Python** and would be easier to code this using some language that has generics. 
I'm not saying that **Go** needs them, I like it's minimalism :) .

# TODO
- ~~User defined functions (lambdas)~~
- ~~Quotation ~~ 
- ~~Better error messages~~
- ~~More native/out of the box functions and constants~~
- ~~Remove unused "Methods"~~
- ~~Overall improve the interaction with the interpreter (REPL, File, ETC)~~
- ~~some examples~~
- ~~Strings~~ (Only 2 basic functions to make it possible manipulate strings from chars)
- Files 
- Comments 
- Improve Prints 
- "Debug Mode"
- concurrency (jk)
- IO

(To be honest this looks more like a Shopping list)


# How to Run
1. Clone the Project Repo `git clone [url of this repo]`
2. Run `go run main.go`


# Options

- `-f [FILENAME]` read program from file.
- `-r` REPL.
- `-i` read program from stdin.

# Examples
Check the [Examples](/tree/master/examples) folder.
 
- `test.gl` - simple define and print value.
- `listCount.l` - user defined function (lambda) that counts the number of elements equal to an atom.
- `42.gl` - big program with multiple user defined functions and input.
- `strings.gl` - some basic string functions.



