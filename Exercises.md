# Exercises


## [Exercise 1.1 (P8)](ch01/ex1.01)

Modify the `echo` program to also print `os.Args[0]`, the name of the command that invoked it.

## [Exercise 1.2 (P8)](ch01/ex1.02)

Modify the `echo` program to print the index and value of each of its arguments, one per line.

## [Exercise 1.3 (P8)](ch01/ex1.03)

Experiment to measure the difference in running time between our potentially
inefficient versions and the one last uses `strings.Join`.
(Section 1.6 illustrates part of the `time` package, and Section 11.4 shows
how to write benchmark tests for systematic performance evaluation.)

## [Exercise 1.4 (P13)](ch01/ex1.04)

Modify `dup2` to print the names of all files in which each duplicated line occurs.

## [Exercise 1.5 (P15)](ch01/ex1.05)

Change the `Lissajous` program's color palette to green on black, for added authenticity.
To create the web color `#RRGGBB`, use `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`,
where each pair of hexadecimal digits represents the intensity of the red, gree, or blue
component of the pixel.

## [Exercise 1.6 (P15)](ch01/ex1.06)

Modify the `Lissajous` program to produce images in multiple colors by adding more values to `palette`
and then displaying them by changing the third argument of `SetColorIndex` in some interesting way.

## [Exercise 1.7 (P17)](ch01/ex1.07)

The function call `io.Copy(dst, src)` reads from `src` and writes to `dst`.
Ust it instead of `ioutil.ReadAll` to copy the response body to `os.Stdout`
without requiring a buffer large enough to hold the entire stream.
Be sure to check the error result of `io.Copy`.

## [Exercise 1.8 (P17)](ch01/ex1.08)

Modify `fetch` to add prefix `http://` to each argument `URL` if it is missing.
You might want to use `strings.HasPrefix`.

## [Exercise 1.9 (P17)](ch01/ex1.09)

Modify `fetch` to also print the `HTTP` status code, found in `resp.Status`.

## [Exercise 1.11 (P19)](ch01/ex1.11)

Try `fetchall` with longer argument lists, such as samples from the top million web sites available at `alexa.com`.
How does the program behave if a web site just doesn't respond?
(Section 8.9 describes mechanisms for coping in such cases.)

## [Exercise 1.12 (P22)](ch01/ex1.12)

Modify the `Lissajous` server to read parameter values from the URL.
For example, you might arrange it so that a URL like `http://localhost:8000/?cycles=20`
sets the number of cycles to 20 instead of the default 5.
Use the `strconv.Atoi` function to convert the string parameter into an integer.
You can see its documentation with `go doc strconv.Atoi`.

## [Exercise 2.1 (P42)](ch02/ex2.01)

Add types, constants, and functions to `tempconv` for processing temperatures in
the Kelvin scale, where `zero` Kelvin is -273.15℃ and a difference of 1K has the
same magnitude as 1℃.

## [Exercise 2.2 (P44)](ch02/ex2.02)

Write a general-purpose unit-conversion program analogous to `cf` that reads
numbers from its command-line arguments or from the standard input if there
are no arguments, and converts each number into units like temperature in `Celsius`
and `Fahrenheit`, length in `feet` and `meters`, weight in `pounds` and `kilograms`, and the like.

## [Exercise 2.3 (P45)](ch02/ex2.03)

Rewrite `PopCount` to use a loop instead of a single expression. Compare the performance of two versions.
(Section 11.4 shows how to compare the performance of different implementations systematically.)

## [Exercise 2.4 (P45)](ch02/ex2.04)

Write a version of `PopCount` that counts bits by shifting its argument through 64-bit positions,
testing the rightmost bit each time. Compare its performance to the table-lookup version.

## [Exercise 2.5 (P45)](ch02/ex2.05)

The expression`x&(x-1)` clears the rightmost no-zero bit of `x`.
Write a version of `PopCount` that counts bits by using this fact, and assess its performance.

## [Exercise 3.1 (P60)](ch03/ex3.01)

If the function `f` returns a non-finite float64 value,
the SVG file will contain invalid `<polygon> ` elements
(although many SVG renderers handle this gracefully).
Modify the program to skip invalid polygons.

## [Exercise 3.2 (P60)](ch03/ex3.02)

Experiment with visualizations of other functions from the math package.
Can you produce an `egg box`, `moguls`, or a `saddle`?

## [Exercise 3.3 (P60)](ch03/ex3.03)

Color each polygon bsed on its height,
so that the peaks are colored red (#ff0000)
and the valleys blue (#0000ff)

## [Exercise 3.4 (P60)](ch03/ex3.04)

Following the approach of the `Lissajous` example in Section 1.7,
construct a web server that computes surfaces and writes SVG data to the client.
The server must set the `Content-Type` header like this:

```golang
    w.Header().Set("Content-Type", "image/svg+xml")
```

## [Exercise 3.5 (P62)](ch03/ex3.05)

Implement a full-color `Mandelbrot set` using the function `image.NewRGBA` and the type `color.RGBA` or `color.YCbCr`.

## [Exercise 3.6 (P62)](ch03/ex3.06)

Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average.
The simplestmethod is to divide each pixel into four "subpixels". Implement it.

## [Exercise 3.7 (P62)](ch03/ex3.07)

Another simple fractal uses `Newton's method` to find complex solutions to a function such as `z^4-1=0`.
Shade each starting point by the number of iterations required to get close to one of the four roots.
Color each point by the root it approaches.

## [Exercise 3.8 (P63)](ch03/ex3.08)

Rendering fractals at high zoom levels demands great arithmetic precision.
Implement the same fractal using four different representations of numbers:
`complex64`, `complex128`, `big.Float`, and `big.Rat`.
(The latter two types are found in the `math/big` package.
`Float` uses arbitrary but bounded-precision floating-point;
`Rat` uses unbounded-precision rational numbers)
How do theycompare in performance and memory usage?
At what zoom levels do rendering artifacts become visible?

## [Exercise 3.9 (P63)](ch03/ex3.09)

Write a web server that renders fractals and writes tye image data to the client.
Allow the client to specify the `x,y`, and `zoom` values as parameters to the HTTP request.

## [Exercise 3.10 (P74)](ch03/ex3.10)

Write a non-recursive version of `comma`, using `bytes.Buffer` instead of string concatenation.

## [Exercise 3.11 (P74)](ch03/ex3.11)

Enhance `comma` so that it deals correctly with floating-point numbers and an optional sign.

## [Exercise 3.12 (P74)](ch03/ex3.12)

Write a function that reports whether two strings are anagrams of each other,
that is, they contain the same letter in different order.

## [Exercise 3.13 (P78)](ch03/ex3.13)

Write const declarations for KB, MB, up through YB as compactly as you can.

## [Exercise 4.1 (P84)](ch04/ex4.01)

Write a function that counts the number of bits that are different in two SHA256 hashes.
(See PopCount from Section 2.6.2.)

## [Exercise 4.2 (P84)](ch04/ex4.02)

Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

## [Exercise 4.3 (P93)](ch04/ex4.03)

Rewrite `reverse` to use an array pointer instead of a slice.

## [Exercise 4.4 (P93)](ch04/ex4.04)

Write a version of `rotate` that operates in a single pass.

## [Exercise 4.5 (P93)](ch04/ex4.05)

Write an in-place function to eliminate adjacent duplicates in a `[]string` slice

## [Exercise 4.6 (P93)](ch04/ex4.06)

Write an in-place function that squashes each run of adjacent Unicode spaces
(see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.

## [Exercise 4.7 (P93)](ch04/ex4.07)

Modify `reverse` to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place.
Can you do it without allocating new memory?

## [Exercise 4.8 (P99)](ch04/ex4.08)

Modify `charcount` to count letters, digits, and so on in their Unicode categories,
using functions like `unicode.IsLetter`.

## [Exercise 4.9 (P99)](ch04/ex4.09)

Write a program `wordfreq` to report the frequency of each word in an input text file.
Call `input.Split(bufio.ScanWords)` before the first call to `Scan` to break the input into words instead of lines.

## [Exercise 4.10 (P112)](ch04/ex4.10)

Modify `issues` to report the results in age categories,
say less than a month old, less than a year old, and more than a year old.

## [Exercise 4.11 (P112)](ch04/ex4.11)

Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
invoking their preferred text editor when substantial text input is required.

## [Exercise 4.12 (P113)](ch04/ex4.12)

The popular web comic `xkcd` has a `JSON` interface.
For example, a request to `http://xkcd.com/571/info.0.json` produces a detailed description of comic 571, one of many favorites.
Download each URL (once!) and build an offline index.
Write a tool `xkcd` that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.

## [Exercise 4.13 (P113)](ch04/ex4.13)

The JSON-based web service of the Open Movie Database lets you search `https://omdbapi.com/`
for a movie by name and download its poster image.
Write a tool `poster` that downloads the poster image for the movie named on the command line.

## [Exercise 4.14 (P117)](ch04/ex4.14)

Create a web server that queries GitHub once and then allows navigation of the list of bug reports, milestones, and users.

## [Exercise 5.1 (P124)](ch05/ex5.01)

Change the `findlinks` program to traverse the `n.FirstChild` linked list using recursive calls to `visit` instead of a loop.


## [Exercise 5.2 (P124)](ch05/ex5.02)

Write a function to populate a mapping from element names--`p, div, span, and so on`--to the number of elements with that name in an HTML document tree.


## [Exercise 5.3 (P124)](ch05/ex5.03)

Write a function to print the contents of all text nodes. in an HTML document tree.
Do not descend into `<script>` or `<stylr>` elements, since their contents are not visible in a web browser.


## [Exercise 5.4 (P124)](ch05/ex5.04)

Extend the `visit` function so that it extracts other kinds of links from the document, such as image, scripts, and style sheets.


## [Exercise 5.5 (P127)](ch05/ex5.05)

Implement `countWordAndImages`. (See Exercise 4.9 for word-splitting.)


## [Exercise 5.6 (P127)](ch05/ex5.06)

Modify the `corner` function in `gopl.io/ch3/surface` (§3.2) to use named results and a bare return statement.


## [Exercise 5.7 (P134)](ch05/ex5.07)

Develop `startElement` and `endElement` into a general HTML pretty-printer.
Print comment nodes, text nodes, and the attributes of each element (`<a href='...'>`).
Use short forms like `<img/>` instead of `<img></img>` when an element has no children.
Write a test to ensure that the output can be parsed successfully.
(See Chapter 11.)

## [Exercise 5.8 (P134)](ch05/ex5.08)

Modify `forEachNode` so that the `pre` and `post` functions
return a boolean result indicating whether to continue the traversal.
Use it to write a function `ElementByID` with the following signature
that finds the first HTML element with the specified `id` attribute.
The function should stop the traversal as soon as a match is found.

```golang
    func ElementByID(doc *html.Node, id string) *html.Node
```

## [Exercise 5.9 (P135)](ch05/ex5.09)

Write a function `expand(s string, f func(string) string) string` that replaces each substring `"$foo"` within `s` by the text returned by `f("foo")`.

## [Exercise 5.10 (P140)](ch05/ex5.10)

Rewrite `topoSort` to use `maps` instead of `slices` and eliminate the initial sort.
Verify that the results, though nondeterministic, are valid topological orderings.

## [Exercise 5.11 (P140)](ch05/ex5.11)

The instructor of the linear algebra course decides that calculus is now a prerequisite.
Extend the `topoSort` function to report cycles.

## [Exercise 5.12 (P140)](ch05/ex5.12)

The `startElement` and `endElement` function in `gopl.io/ch5/outline2` (§5.5) share a global variable, `depth`.
Turn them into anonymous functions that share a variable local to the `outline` function.

## [Exercise 5.13 (P140)](ch05/ex5.13)

Modify `crawl` to make local copies of the pages it finds, creating directories as necessary.
Don't make copies of pages that come from a different domain.
For example, if the original page comes from `golang.org`, save all files from there, but exclude ones from `vimeo.com`.

## [Exercise 5.14 (P140)](ch05/ex5.14)

Use the `breadthFirst` function to explore a different structure.
For example, you could use the course dependencies form the `topoSort` example (a directed graph),
the file system hierarchy on your computer (a tree),
or a list of bus or subway routes downloaded from your city government's web site (an undirected graph).

## [Exercise 5.15 (P143)](ch05/ex5.15)

Write variadic functions `max` and `min`, analogous to `sum`.
What should these functions do when called with no arguments?
Write variants that require at least one argument.

## [Exercise 5.16 (P143)](ch05/ex5.16)

Write a variadic version of `strings.Join`.

## [Exercise 5.17 (P143)](ch05/ex5.17)

Write a variadic function `ElementsByTagName` that,
given an HTML node tree and zero or more names,
returns all the elements that match one of those names.
Here are two example calls:

``` Go
    func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
    images := ElementsByTagName(doc, "img")
    headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
```

## [Exercise 5.18 (P148)](ch05/ex5.18)

Without changing its behavior, rewrite the `fetch` functin to use `defer` to close the writable file.

## [Exercise 5.19 (P153)](ch05/ex5.19)

Use `panic` and `recover` to write a function that contains no return statement yet returns a non-zero value.

## [Exercise 6.1 (P167)](ch06/ex6.01)

Implement these additional methods:

``` Go
    func (*IntSet) Len() int      // return the number of elements
    func (*IntSet) Remove(x int)  // remove x from the set
    func (*IntSet) Clear()        // remove all elements from the set
    func (*IntSet) Copy() *IntSet // return a copy of the set
```

## [Exercise 6.2 (P167)](ch06/ex6.02)

Define a variadic `(*IntSet).AddAll(...int)` method that allows a list of values to be added, such as `s.AddAll(1, 2, 3)`.

## [Exercise 6.3 (P167)](ch06/ex6.03)

`(*IntSet).UnionWith` computes the union of two sets using `|`, the word-parallel bitwise OR operator.
Implement methods for `IntersectWith`, `DifferenceWith` and `SymmetricDifference` for the corresponding set operations.
(The symmetric difference of two sets contains the elements present in one set or the other but not both.)

## [Exercise 6.4 (P168)](ch06/ex6.04)

Add a method `Elems` that returns a slice containing the elements of the set, suitable for iterating over with a `range` loop.

## [Exercise 6.5 (P168)](ch06/ex6.05)

The type of each word used by `IntSet` is `uint64`,
but 64-bit arithmetic may be inefficient on a 32-bit platform.
Modify the program to use the `uint` type,
which is the most efficient unsigned integer type for the platform.
Instead of dividing by 64, define a constant holding the effective `size of uint` in bits, 32 or 64.
You can use the perhaps too-clever expression `32 << (^uint(0) >> 63) for this purpose.

## [Exercise 7.1 (P173)](ch07/ex7.01)

Using the ideas from `ByteCounter`, implement counters for words and for lines.
You will find `bufio.ScanWords` useful.

## [Exercise 7.2 (P174)](ch07/ex7.02)

Write a function `CountingWriter` with the signature below that,
given an `io.Writer`, returns a new Writer that wraps the original,
and a pointer to an `int64` variable that at any moment contains the number of bytes written to the new Writer.

``` Go
    func CountingWriter(w io.Writer) (io.Writer, *int64)
```

## [Exercise 7.3 (P174)](ch07/ex7.03)

Write a `String` method for the `*tree` type in *gopl.io/ch4/treesort (§4.4)* that reveals the sequence of values in the tree.

## [Exercise 7.4 (P175)](ch07/ex7.04)

The `strings.NewReader` function retruns a value that satisfies
the `io.Reader` interface (and others) by reading from its argument, a string.
Implement a simple version of `NewReader` yourself, and use it to make
the HTML parser (§5.2) take input from a string.

## [Exercise 7.5 (P175)](ch07/ex7.05)

The `LimitReader` function in the io package accepts an `io.Reader r` and a number of bytes `n`,
and returns another `Reader` that reads from `r` but reports an end-of-file condition after n bytes.
Implement it.

``` Go
    func LimitReader(r io.Reader, n int64) io.Reader
```

## [Exercise 7.6 (P181)](ch07/ex7.06)

Add support for Kelvin temperatures to `tempflag`.

## [Exercise 7.7 (P181)](ch07/ex7.07)

Explain why the help message contains `°C` when the default value of `20.0` does not.

**Answer**: The flag's value is a `Stringer`, and is used in command-line help messages.
`Celsius` has a `func (c Celsius) String() string` method, so `Celsius` is a `Stringer`.
When `flag` shows the help messages, it will call `Celsius' String` method to format the value.

## [Exercise 7.8 (P191)](ch07/ex7.08)

Many GUIs provide a table widget with a stateful multi-tier sort:
the primary sort key is the most recently clicked column head,
the secondary sort key is the second-most recently clicked column head, and so on.
Define an implementation of `sort.Interface` for use by such a table.
Compare that approach with repeated sorting using `sort.Stable`.

## [Exercise 7.9 (P191)](ch07/ex7.09)

Use the `html/template` package (§4.6) to replace `printTracks` with a function that displays the tracks as an HTML table.
Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.

## [Exercise 7.10 (P191)](ch07/ex7.10)

The `sort.Interface` type can be adapted to other use.
Write a function `IsPalidrome(s sort.Interface) bool` that reports whether the sequence `s` is a palindrome,
in other words, reversig the sequence would not change it.
Assume that the elements at indices `i` and `j` are equal if `!s.Less(i, j) && !s.Less(j, i)`

## [Exercise 7.11 (P195)](ch07/ex7.11)

Add additional handlers so that clients can create, read, update, and delete database entries.
For example, a request of the form `/update?item=socks&price=6` will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid.
(Warning: this change introduces concurrent variable update.)

## [Exercise 7.12 (P195)](ch07/ex7.12)

Change the handler for `/list` to print its output as an HTML table, not text.
You may find the `html/template` package (§4.6) useful.

## [Exercise 7.13 (P205)](ch07/ex7.13)

Add a `String` method to `Expr` to pretty-print the syntax tree.
Check that the results, when parsed again, yield an equivalent tree.

## [Exercise 7.14 (P205)](ch07/ex7.14)

Define a new concrete type that satisfies the `Expr` interface and provides a new operation such as computing the minimum value of its operands.
Since the `Parse` function does not create instances of this new type, to use it you will need to construct a syntax tree directly (or extend the parser).

## [Exercise 7.15 (P205)](ch07/ex7.15)

Write a program that reads a single expression from the standard input, prompts the user to provide values fro any variables, then evaluates the expression in the resulting environment. Handle all errors gracefully.

## [Exercise 7.16 (P205)](ch07/ex7.16)

Write a web-based calculator program.

## [Exercise 7.17 (P215)](ch07/ex7.17)

Extend `xmlselect` so that elements may be selected not just by name, but by their attributes too,
in the manner of CSS, so that, for instance, an element like `<div id="page" class="wide">` could be selected by a matching `id` or `class` as well as its name.

## [Exercise 7.18 (P215)](ch07/ex7.18)

Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it.
Nodes are of two kinds: `CharData` nodes represent text strings, and `Element` nodes represent named elements and their attributes.
Each element node has a slice of child nodes.

You may find the following declarations helpful.

``` Go
    import "encoding/xml"
    type Node interface{} // CharData or *Element
    type CharData string
    type Element struct {
        Type     xml.Name
        Attr     []xml.Attr
        Children []node
    }
```

## [Exercise 8.1 (P222)](ch08/ex8.01)

Modify `clock2` to accept a port number, and write a program, `clockwall`, that acts as a client of several clock servers at once,
reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices.
If you have access to `geographically` distributed computers, run instances remotely;
otherwise run local instances on different ports with fake time zones.

``` text
    $ TZ=US/Eastern    ./clock2 -port 8010 &
    $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
    $ TZ=Europe/London ./clock2 -port 8030 &
    $ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
```

## [Exercise 8.2 (P222)](ch08/ex8.02)

Implement a concurrent `File Transfer Protocol (FTP)` server.
The server should interpret commands from each client such as `cd` to chagne directore,
`ls` to list a directory,
`get` to send the contents of a file,
and `close` to close the connection.
You can use the standard ftp command as the client, or write your own.

## [Exercise 8.3 (P227)](ch08/ex8.03)

In `netcat3`, the interface value conn has the concrete type `*net.TCPConn`, which represents a TCP connection.
A TCP connection consists of two halves that may be closed independently using its `CloseRead` and `CloseWrite` methods.
Modify the main goroutine of `netcat3` to close only the write half of the connection so that the program will continue to print the final echoes from the `reverb1` server even after the standard input has been closed.
(Doing this for the `reverb2` server is harder; see Exercise 8.4)

## [Exercise 8.4 (P239)](ch08/ex8.04)

Modify the `reverb2` server to use a `sync.WaitGroup` per connection to count the number of active `echo` goroutines.
When it falls to zero, close the write half of the TCP connection as described in Exercise 8.3.
Verify that your modified `netcat3` client from that exercise waits fro the final echos of multiple concurrent shouts,
even after the standard input has been closed.

## [Exercise 8.5 (P239)](ch08/ex8.05)

Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2,
and execute its main loop in parallel using channels for communication.
How much faster does it run on a multiprocessor machine?
What is the optimal number of goroutines to use?

## Results

Render 4096*4096 Mandelbrot

``` text
    NumCPU: 4
    Done no concurrency. Used: 5.144318122s
    Done. Worker Number: 1 Used: 5.060132177s
    Done. Worker Number: 2 Used: 2.820846794s
    Done. Worker Number: 3 Used: 2.781336072s
    Done. Worker Number: 4 Used: 2.579062955s
```

## [Exercise 8.6 (P243)](ch08/ex8.06)

Add depth-limiting to the concurrent `crawler`.
That is, if the user sets `-depth=3`, then only URLs reachable by at most three links will be fetched.

## [Exercise 8.7 (P243)](ch08/ex8.07)

Write a concurrent program that creates a local mirror of a web site,
fetching each reachable page and writing it to a directory on the local disk.
Only pages within theoriginal domain (for instance, `golang.org`) should be fetched.
URLs within mirrored pages should be altered as needed so that they refer to the mirrored page, not the original.

## [Exercise 8.8 (P247)](ch08/ex8.08)

Using a select statement, add a timeout to the echo server form Section 8.3 so that it disconnects any client that shouts nothing within 10 seconds.

## [Exercise 8.9 (P251)](ch08/ex8.09)

Write a version of `du` that computes and periodically displays separate totals for each of the root directories.

## [Exercise 8.10 (P253)](ch08/ex8.10)

HTTP requests may be cancelled by closing the optional `Cancel` channel in the `http.Request` struct.
Modify the web crawler of Selection 8.6 to support cancellation.

*Hint:* the `http.Get` convenience function does not give you an opportunity to customize a `Request`.
Instead, create the request using `http.NewRequest`, set its `Cancel` field, then perform the request by calling `http.DefaultClient.Do(req)`.

## [Exercise 8.11 (P153)](ch08/ex8.11)

Following the approach of `mirroredQuery` in Section 8.4.4, implement a variant of `fetch` that requests several URLs concurrently.
As soon as the first response arrives, cancel the other requests.

## [Exercise 8.12 (P256)](ch08/ex8.12)

Make the broadcaster announce the current set of clients to each new arrival.
This requires that the `clients` set and the `entering` and `leaving` channels record the client name too.

## [Exercise 8.13 (P156)](ch08/ex8.13)

Make the chat server disconnect idle clients, such as those that have sent no messages in the last fine minutes.

*Hint:* calling `conn.Close()` in another goroutine unblocks active `Read` calls such as the one done by `input.Scan()`.

## [Exercise 8.14 (P256)](ch08/ex8.14)

Change the chat server's network protocol so that each client provides its name on entering.
Use that name instead of the netowrk address when prefixing each message with its sender's identity.

## [Exercise 8.15 (P256)](ch08/ex8.15)

Failure of any client program to read data in timely manner ultimately causes all clients to get stuck.
Modify the broadcaster to skip a message rather than wait if a client writer is not ready to acceptit.
Alternatively, add buffering to each client's outgoing message channel so that most messages are not dropped; the boradcaster should use a non-blocking send to this channel.

## [Exercise 9.1 (P262)](ch09/ex9.01)

Add a function `Withdraw(amount int) bool` to the `gopl.io/ch9/bank1` program.
The result should indicate whether the transaction succeeded or failed due to insufficient funds.
The message send to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to `Withdraw`.

## [Exercise 9.2 (P271)](ch09/ex9.02)

Rewrite the `PopCount` example from Section 2.6.2 so that it initializes the lookup table using `sync.Once` the first time it is needed.
(Realistically, the cost of synchronization whould be prohibitive for a small and highly optimized function like `PopCount`)

## [Exercise 9.3 (P279)](ch09/ex9.03)

Extend the Func type and the `(*Memo).Get` method so that callers may provide an optional done channel through which they can cancel the operation (§8.9).
The results of a cancelled Func call should not be cached.

## [Exercise 9.4 (P280)](ch09/ex9.04)

Construct a pipeline that connects an arbitrary number of goroutines with channels.
What is the maximum number of pipeline stages you can create with out running out of memory?
How long does a value take to transit the entire pipeline?

## Result

### With ResetTimer after initialize all channels

```text
BenchmarkPipeline10-4       500000  2573       ns/op  0   B/op  0  allocs/op
BenchmarkPipeline100-4      50000   24956      ns/op  0   B/op  0  allocs/op
BenchmarkPipeline1000-4     5000    286223     ns/op  0   B/op  0  allocs/op
BenchmarkPipeline10000-4    300     3980524    ns/op  0   B/op  0  allocs/op
BenchmarkPipeline100000-4   30      39138676   ns/op  26  B/op  0  allocs/op
BenchmarkPipeline1000000-4  3       379899748  ns/op  0   B/op  0  allocs/op
```

### Without ResetTimer after initialize all channels

```text
BenchmarkPipeline10-4       500000  2497        ns/op  0          B/op  0        allocs/op
BenchmarkPipeline100-4      50000   24588       ns/op  0          B/op  0        allocs/op
BenchmarkPipeline1000-4     5000    287139      ns/op  26         B/op  0        allocs/op
BenchmarkPipeline10000-4    300     3697650     ns/op  4825       B/op  41       allocs/op
BenchmarkPipeline100000-4   30      42878492    ns/op  381184     B/op  3657     allocs/op
BenchmarkPipeline1000000-4  1       4149927880  ns/op  621803856  B/op  2799323  allocs/op
```

## With a 16GB OS, ~8GB available mem

When trying to create 5000000, it running out of memory.

## [Exercise 9.5 (P281)](ch09/ex9.05)

Write a program with two gotoutines that send messages back and forth over two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?

## Result

```text
2.1975286807808625e+06 rounds per second
```

## [Exercise 9.6 (P282)](ch09/ex9.06)

Measure how the performance of a compute-bond parallel program (see Exercise 8.5) varies with GOMAXPROCS.
What is the optimal value of your computer?
How many CPUs does your computer have?

## Result

With Intel I5 6200, 4 Cores, 2 worker threads is economical.

```text
NumCPU: 4
Done no concurrency. Used: 5.101064434s
Done. Worker Number: 1 Used: 5.076399681s
Done. Worker Number: 2 Used: 2.735404646s
Done. Worker Number: 3 Used: 2.56978686s
Done. Worker Number: 4 Used: 2.522902384s
```

## [Exercise 10.1 (P288)](ch10/ex10.01)

Extend the `jpeg` program so that it converts any supported input format to any output format, using `image.Decode` to detect the input format and a flag to select the output format.

## [Exercise 10.2 (P288)](ch10/ex10.02)

Define a generic archive file-reading function capable of reading ZIP files (archive/zip) and POSIX tar filex (archive/tar).
Use a registration mechanism sililar to the one described above so that support for each file format can be plugged in using blank imports.

## [Exercise 10.3 (P293)](ch10/ex10.03)

Using `fetch http://gopl.io/ch1/helloworld?go-get=1`, find out which service hosts the code samples for this book.
(HTTP requests from `go get` include the `go-get` parameter so that servers can disginguish them from ordinary browser requests.)

## Result

```shell
go run fetch.go http://gopl.io/ch1/helloworld\?go-get\=1 | grep go-import
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
```

## [Exercise 10.4 (P300)](ch10/ex10.04)

Construct a tool that reports the set of all packages in the worksapce that transitively depend on the packages specified by the arguments.
Hint: you will need to run `to list` twice, once for the initial packages and once for all packages.
You may want to parse its JSON output using the `encodin g/json` package (§4.5)

## [Exercise 11.1 (P307)](ch11/ex11.01)

Write tests for the `charcount` program in Section 4.3.

## [Exercise 11.2 (P307)](ch11/ex11.02)

Write a set of tests for `IntSet (§6.5)` that checks that its behavior after each operation is equivalent to a set bsed on built-in maps.
Save your implementation for benchmarking in Exercise 11.7.

## [Exercise 11.3 (P308)](ch11/ex11.03)

`TestRandomPalindromes` only tests palindromes.
Write a randomized test that generates and verifies *non-palindromes.*

## [Exercise 11.4 (P308)](ch11/ex11.04)

Modify `randomPalindrome` to exercise `IsPalindrome's` handling of punctuation and spaces.

## [Exercise 11.5 (P317)](ch11/ex11.05)

Extend `TestSplit` to use a table of inputs and expected outputs.

## [Exercise 11.6 (P323)](ch11/ex11.06)

Write benchmarks to compare the `PopCount` implementation in Section 2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5.
At what point does the table-based approach break even?

```text
BenchmarkPopCountTable1-4        20000       121925    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable10-4       10000       567430    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable100-4      10000       5678211   ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable1000-4     3000        16778123  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable10000-4    300         17531809  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable100000-4   100         54971520  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift1-4        20000000    59.1      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift10-4       3000000     620       ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift100-4      200000      8679      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift1000-4     10000       113931    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift10000-4    2000        1130659   ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift100000-4   100         11198696  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears1-4       1000000000  3.12      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears10-4      30000000    39.0      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears100-4     3000000     528       ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears1000-4    200000      7529      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears10000-4   10000       100799    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears100000-4  1000        1257457   ns/op  0  B/op  0  allocs/op
```

## [Exercise 11.7 (P323)](ch11/ex11.07)

Write benchmarks for `Add`, `UnionWith`, and other methods of `*IntSet` (6.5) using large pseudo-random inputs.
How fast can you make these methods run?
How does the choice of word size affect performance?
How fast is `IntSet` compared to a set implementation based on the built-in map type?

```text
BenchmarkMapIntSetAdd10-4       1000000   1411    ns/op  323    B/op  3     allocs/op
BenchmarkMapIntSetAdd100-4      100000    16217   ns/op  3474   B/op  20    allocs/op
BenchmarkMapIntSetAdd1000-4     10000     194432  ns/op  55611  B/op  98    allocs/op
BenchmarkMapIntSetHas10-4       20000000  59.8    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetHas100-4      20000000  68.4    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetHas1000-4     20000000  63.7    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetAddAll10-4    2000000   985     ns/op  323    B/op  3     allocs/op
BenchmarkMapIntSetAddAll100-4   100000    13628   ns/op  3475   B/op  20    allocs/op
BenchmarkMapIntSetAddAll1000-4  5000      206743  ns/op  55627  B/op  98    allocs/op
BenchmarkMapIntSetString10-4    500000    2916    ns/op  368    B/op  14    allocs/op
BenchmarkMapIntSetString100-4   50000     33491   ns/op  4577   B/op  108   allocs/op
BenchmarkMapIntSetString1000-4  5000      336538  ns/op  41164  B/op  994   allocs/op
BenchmarkBitIntSetAdd10-4       3000000   536     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAdd100-4      300000    4428    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAdd1000-4     30000     45583   ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas10-4       30000000  42.8    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas100-4      30000000  41.3    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas1000-4     30000000  41.9    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll10-4    20000000  106     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll100-4   2000000   609     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll1000-4  300000    5464    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetString10-4    500000    3065    ns/op  256    B/op  12    allocs/op
BenchmarkBitIntSetString100-4   50000     27027   ns/op  3649   B/op  106   allocs/op
BenchmarkBitIntSetString1000-4  5000      209314  ns/op  33009  B/op  1002  allocs/op
```

## [Exercise 12.1 (P338)](ch12/ex12.01)

Extend `Display` so that it can display maps whose `keys` are `structs` or `arrays`.

## [Exercise 12.2 (P338)](ch12/ex12.02)

Make display safe to use on cyclic data structure by bounding the number of steps it takes before abandoning the recursion.
(In Section 13.3, we'll see another way to detect cycles.)

## [Exercise 12.3 (P341)](ch12/ex12.03)

Implement the missing cases of the `encode` function.
Encode booleans at `t` and `nil`, floating-point numbers using Go's notation, and complex numbers like `1+2i` as `#C(1.0 2.0)`.
Interfaces can be encoded as a pair of a type name and a value, for instance `("[]int" (1 2 3))`, but beware that this notation is ambiguous:
the `reflect.Type.String` method may return the same string for different types.

## [Exercise 12.4 (P341)](ch12/ex12.04)

Modify `encode` to pretty-print the S-expression in the style shown above.

## [Exercise 12.5 (P341)](ch12/ex12.05)

Adapt `encode` to emit JSON instead of S-expressions.
Test your encoder using the standard docoder, `json.Unmarshal`.

## [Exercise 12.6 (P341)](ch12/ex12.06)

Adapt `encode` so that, as an optimization, it does not encode a field whose value is the `zero` value of its type.

## [Exercise 12.7 (P341)](ch12/ex12.07)

Create a streaming API for the S-expression decoder, following the style of `json.Decoder`((§4.5).

## [Exercise 12.8 (P347)](ch12/ex12.08)

The `sexpr.Unmarshal` function, like `json.Marshal`, requires the complete input in a byte slice before it can begin decoding.
Define a `sexpr.Decoder` type that, like `json.Decoder`, allows a sequence of values to be decoded from an `io.reader`.
Change `sexpr.Unmarshal` to use this new type.

## [Exercise 12.9 (P347)](ch12/ex12.09)

Write a token-based API for decoding S-expressions, following the style of `xml.Decoder` (§7.14).
You will need five types of tokens: `Symbol`, `String`, `Int`, `StartList`, and `EndList`.

## [Exercise 12.10 (P347)](ch12/ex12.10)

Extend `sexpr.Unmarshal` to handle the booleans, floating-point numbers, and interfaces encoded by your solution to Exercise 12.3.
(Hint: to decode interfaces, you will need a mapping from the name fo each supported type to its `reflect.Type`.)

## [Exercise 12.11 (P350)](ch12/ex12.11)

Write the corresponding `Pack` function.
Given a struct value, Pack should return a URL incorporating the parameter values form the struct.

## [Exercise 12.12 (P350)](ch12/ex12.12)

Extend the field tag notation to express parameter validity requirements.
For example, a string might need to be a valid email address or credit-card number, and an integer might need to be a valid US ZIP code.
Modify `Unpack` to check these requirements.

## [Exercise 12.13 (P351)](ch12/ex12.13)

Modify the S-expression `encoder` (§12.4) and `decoder` (§12.6) so that they honor the `sexpr:"..."` field tag in a similar manner to `encoding/json (§4.5)`.

## [Exercise 13.1 (P361)](ch13/ex13.01)

Define a deep comparison function that considers numbers (of any type) equal if they differ by less than one part in a billion.

## [Exercise 13.2 (P361)](ch13/ex13.02)

Write a function that reports whether its argument is a cyclic data structure.

## [Exercise 13.3 (P366)](ch13/ex13.03)

Use `sync.Mutex` to make `bzip2.writer` safe for concurrent use by multiple goroutines.

## [Exercise 13.4 (P366)](ch13/ex13.04)

Depending on C libraries has its drawbacks.
Provide an alternative pure-Go implementation of `bzip.NewWriter` that uses the `os/exec` package to run `/bin/bzip2` as a subprocess.
