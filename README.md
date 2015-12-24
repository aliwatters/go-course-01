# Go Course - Go Lang Coding.
Notes and examples from Udemy course.

[https://www.udemy.com/learn-how-to-code](https://www.udemy.com/learn-how-to-code)

## Section 1: Setup.

Installation - already had 1.4 - moved to 1.5.2 (64bit linux)

```
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
```

Note: go recommends a workspace folder - mine is in /home/ali/go - made a go-course-01 repo for this.

Note: hit errors with upgrade of Go

```
package runtime: C source files not allowed when not using cgo or SWIG: atomic_amd64x.c defs.c float.c heapdump.c lfstack.c malloc.c mcache.c mcentral.c mem_linux.c mfixalloc.c mgc0.c mheap.c msize.c os_linux.c panic.c parfor.c proc.c runtime.c signal.c signal_amd64x.c signal_unix.c stack.c string.c sys_x86.c vdso_linux_amd64.c
```

Seems there is old c code left over from previous installation:

```
sudo rm -rf /usr/local/go/
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
```

Resolved Problem - HELLO WORLD! :)

Additionally - if installing first time - have to set up some go env vars. GOPATH (workspace) GOROOT (/usr/local/bin/go - or whereever) etc - easy done before.

For the course then run: `go get github.com/goestoeleven/golangtraining` - adds to ~/go/src/... (note mid-way in section 4 he changes username/repo to lowercased versions - better but - seriously - edit the original video!)

## Section 2

### Packages

Exported / unexported - capitalized is visible outside of the package, lowercase internal. public/private not used.

## Section 4 - scope

Had some problems when I tried to create and use a package called foo. Switched to bar and things worked... odd. I did have foo/foo.go - changed to bar/foo.go - might have been something to do with it.

Func expressions - eg:

inside a function you can assign an anonymous function to a var very similar to JS at this point.

We can also - write functions that return functions - works very similar to javascript (first class functions)

```
func wrapper() func() int {
  x := 0
  return func() int {
    x++
    return x
  }
}

func main() {
  increment := wrapper() // IIFE!
  fmt.Println(increment())
  fmt.Println(increment())
}
```


# Section 5 - fundamentals

## Constants

iota - increments by 1 each time used (neat)
untyped constants - allows 42 to be used with any numeric value (regardless of it's type!)
but there are also typed constants which can be used as static typed

## Pointers

show the address with & so

```
a := 42
fmt.Println("a's address is", &a)
```


note: fmt.Scan - reads into a memory address, eg.

```
var note string
fmt.Scan(&note)
```

create pointers with * and dereference pointers with *

```
a := 43
var b *int = &a
fmt.Println("b has address", b, "and value", *b)
```

# Section 6 - control flow

for loops - does more than in most languages!

```
// standard
for init; cond; post { }

// like a while
for condition { }

// endless (break to get out)
for { }

// also for will iterate over an array
for key, value := range oldMap {
  newMap[key] = value
}

// value can be ignored with
for key := range m { ... }

// key can be ignored with the blank identifier _
for _, value := range m { ... }

// strings can be treated as array like!
for pos, char := range "some utf-8 string" { ... }

// and you can multiple init & post
// reverse a in place
for i, j := 0, len(a) - 1; i < j; i = i + 1, j-1 {
  a[i], a[j] = a[j], a[i]
}

```


## Rune

a rune is character from utf-8 - it is also int32 -- maps to the utf-8 character set (4 bytes - 4*8 = 32)

```
x := 'a'
// x - is now the rune value of a - a int32 value 92
```


## String Literals

```
a := `Preserve  whitespace newlines etc and "quotes" don't need
				escaped`
b := "Single line strings"
```

## Switch statements

```
switch "value" {
	case "value":
		fmt.Println("Value found")
	case "b":
		fmt.Println("b found")
		fallthrough
	case "c":
		fmt.Println("c found") // b will also print this
	case "d", "e":
		fmt.Println("d or e found")
	default:
		fmt.Println("the default case")
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}


switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

Impressive - again go is really powerful. 

## more on if

```
b := true

if food := "chocolate"; b {
	fmt.Println(food) // initialized and scoped in this block only - keep scope tight
}
```


