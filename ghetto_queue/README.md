Ghetto Queue
============

A simple ``Queue``. Currently only works with ``string`` types.


Usage
-----

<code>
var queue = Queue{}
queue.put("Hello")
queue.put("World")
fmt.Println("Empty?", queue.empty())
fmt.Println("Length", queue.length())
first := queue.get()
fmt.Println("Saw:", first)
queue.put("foo")
queue.put("bar")
fmt.Println("Length", queue.length())
</code>


Disclaimer
----------

Since this is only the second thing I've ever written in Go, this is likely
**not** concurrency-safe. I have not tried it in conjunction with Goroutines
or threading, but am dubious "it'll be fine".

It was mostly an exercise to slices (after struggling to figure out ``Array``
then learning it was the wrong thing to be using). Don't put it in production,
please.
