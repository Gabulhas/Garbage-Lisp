![logo](https://i.imgur.com/KGKYp3F.png)



# Garbage Lisp  
Well it works but isn't good.

**Go** is definitely not the best language to do this, so it came out kinda hacky with **Go** Interfaces.

I coded this to learn about Interpreters, so this is definitely not a good example nor it works like most lisps so far.

I tried following this guide http://norvig.com/lispy.html, yet I had to change a lot of stuff because **Go** isn't Dynamically
typed like **Python** and would be easier to code this using some language that has generics. 
I'm not saying that **Go** needs them, I like it's minimalism :) .

# TODO
- [x] User defined functions (lambdas)
- [x] Quotation  
- [x] Better error messages
- [x] More native/out of the box functions and constants
- [x] Remove unused "Methods"
- [x] Overall improve the interaction with the interpreter (REPL, File, ETC)
- [x] some examples
- [x] Strings
- [x] IO
- [ ] Fix Eval
- [ ] Files 
- [ ] Comments 
- [ ] Improve Prints 
- [ ] "Debug Mode"
- [ ] concurrency (jk)

(To be honest this looks more like a Shopping list)


# How to Run

1. Clone the Project Repo `git clone [url of this repo]`
2. Compile with Go `go build .`
3. Run with `./Garbage-Lisp`

# Options
- `-f [FILENAME]` read program from file.
- `-r` REPL.
- `-i` read program from stdin.

# Examples
Check the [Examples](/tree/master/examples) folder.
- `test.gl`     - simple define and print value.
- `listCount.gl`- user defined function (lambda) that counts the number of elements equal to an atom.
- `42.gl`       - big program with multiple user defined functions and input.
- `strings.gl`  - some basic string functions.

# Project Structure
- Env       - Environment Logic, like variables in the scope and native functions.
- Evaluator - AST walker. Includes interpretation/evaluation logic.
- LispTypes - All possible types. LispToken is the name of the Interface which all types implement.
- Parser    - Parser (obviously). Transforms code into an AST tree.
- main.go   - Starting File. Parses command line argument. Either Reads from a file (or Stdin) or starts a REPL.

# Contributions
**Don't**
Unless you want to suggest stuff.
