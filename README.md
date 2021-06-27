# pfn-play
a playground for the pfn programming language i made, there's no need to install anything, just [try it](https://parserfun.herokuapp.com/) online!
# about pfn
![](https://github.com/TsukiGva2/pfn/blob/17418d6471f37c9a8c25124e5d54465024955a6b/img/assoc.png)
## pfn is...
a transpiled programming language i made while trying to learn parsers, it's still a work in progress.
the language borrows some design ideas from languages like lisp, Haskell, ruby and clojure, but it isn't quite consistent when it comes to language constructs,  for example, functions are defined with .f(|x,y|) but are called with (f x y), another example, function blocks use ( and ) while "when" blocks use "do" and "end".
it transpiles to python btw
## why making yet another language?
i know, most of the stuff i do is mess around with
programming languages, but this is the last one, i
promise.
## About the language
### code sample?
here you go.
```ruby
.f (
  |x,y|
  -> (+ x y)
)
a:=2
(print (f a 5))
```
this transpiles to its equivalent python code
```py
NOTE: before your code comes a big chunk of code that defines standard functions and utilities like __pfn_call
...
def pfn_f(*args):
        if len(args) < 2:
                raise ArgcountError('too few arguments for function f')
        x = args[0]
        y = args[1]
        return (x+y)
a=2
print(__pfn_call([pfn_f], a,5))
```
The output code isn't the most pretty or optimized, i plan to work on that part after i finish the language itself.
### Your code can be compact too
```ruby
.f ( |x,y| -> (+ x y) )
a := 2
(print (f a 5))
```
but try to not end up with code like this:
```ruby
.f(|x,y|->(+x y))a:=2(print(f a 5))
```
### why is the design mixed between lisp-like and whatever that other style is
I actualy spent 2 weeks trying to figure out how to parse expressions like 1 + 2 -2 + 2 / 3 * 2,
but i gave up after a few stack overflows, infinite loops and weird results.
I then switched to a new syntax, "+(x,y)", but after a few tests, i realized that it got pretty similar to the lisp-like syntax after some nesting, so i just switched to lisp-like.
### and here is a more complex example showing most of the language's features
```ruby
# return the square root of a given number
py {import math}
.iSqrt(|0|->"undefined") # return "undefined" when called with 0
.iSqrt(|1|->1)
.iSqrt(|x|
  when (< x 0) do
    -> "no real sqrt"
  end 
  -> (/ 1 (math.sqrt x))
)
x:=(input "type a number: ")
(print (iSqrt (int x)))
```
# how to use this code?

## running

just head over to https://parserfun.herokuapp.com/ and try typing some of the samples shown above!
just head over to https://parserfun.herokuapp.com/ and try typing some of the samples shown above,
the script will not run, but instead will be transpiled to the equivalent python code, then you can just copy-paste it to
a python file and run it.
