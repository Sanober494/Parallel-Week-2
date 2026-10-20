# Homework \#2

This ungraded homework is intended to serve as an introduction to using atomic
operations and importance of understanding the architecture of the
parallel hardware a parallel application is running on.

### Programming Questions

For this assignment you are **only** allowed to use the following Go
concurrent constructs:

  - `go` statement
  - `sync/atomic` package. You may use any of the atomic operations.
  - `sync.WaitGroup` and its associated methods.

As the course progresses, you will be able to use other constructs
provided by Go such as channels, conditional variables, mutexes, etc. I
want you to learn the basics and build up from there. This way you'll
have a good picture of how parallel programs and their constructs are
built from each other. **If you are unsure about whether you are able to
use a language feature then please ask on Slack or during office hours
before using it**.

## Static Task Decomposition

For problems 1 and 2, you will be implementing your first parallel
programs in the course. Each goroutine will need to perform a certain
amount of work to complete. *But how should the work be distributed to
each goroutine*? For this assignment, we use the task distribution
technique we saw during class, which equally distributes the amount of
work to each goroutine.

## Problem 0: Set Up Cluster Access

I encourage you to set up your access to the Linux cluster, if you have not
done so already for other classes. Make sure you know your credentials,
know how to log in via `ssh`, and try to run a simple Go program, for example
the week 2 examples. Knowing that this all works will help you immensely with
the projects.

## Problem 1: Estimating Pi using the Leibniz formula

Inside the `pi` directory, open the file called `pi.go` and write a
concurrent Go program that uses an infinite series method to estimate pi.
One of the most widely known infinite series that can be used for this is
the Leibniz formula, as described here:

  - [How To Make
    Pi](https://en.wikipedia.org/wiki/Leibniz_formula_for_π)

Your program should use this formula, and must use the following 
required command-line arguments, and print the below usage statement
if called incorrectly:

``` go
const usage = "Usage: pi interval threads\n" +
"    interval = the number of iterations to perform\n" +
"    threads = the number of threads (i.e., goroutines to spawn)\n"
```

The main goroutine should read in the `interval` and `threads` argument
and **only** print the estimate of pi. The program **must** use the
fork-join pattern described in class and a static task distribution
technique. The threshold for running the sequential version is 1000
intervals (that means that below the threshold, you must run sequentially
even though the threads value is always set); otherwise, you
must run the parallel implementation.

**Assumptions**: You can assume that `interval` and `threads` will
always be given as integers and be provided. Thus, you do not need to
perform error checking on those arguments. You also will not need to
print out the usage statement for this problem. The usage is given here
for clarification purposes.

Sample Runs ($: is just mimicking the command line) :

    $: go run hw2/pi 100 2
    3.1315929036
    $: go run hw2/pi 10 1
    3.0418396189
    $: go run hw2/pi 10 3
    3.0418396189

The tests for this problem is inside `pi_test.go`. Please go back to
homework \#1 if you are having trouble running tests.

## Problem 2: Waitgroup Implementation

Read the following about go interfaces before starting this problem:

  - <https://gobyexample.com/interfaces>
  - <https://www.alexedwards.net/blog/interfaces-explained>

Inside the `waitgroup` directory, open the file called `waitgroup.go`.
Consider the following interface and function:

``` go
type WaitGroup interface {
  Add(amount uint)
  Done()
  Wait()
}

// NewWaitGroup returns a instance of a waitgroup
// This instance must be a pointer and should not
// be copied after creation.
func NewWaitGroup() WaitGroup {
  //IMPLEMENT ME!
}
```

Provide an implementation of a `WaitGroup` that mimics the Go
implementation:

  - <https://golang.org/pkg/sync/#WaitGroup>

The only difference is that our implementation requires uint (rather
than just int) to be passed to the `Add` method. Remember, you can only
use atomic operations in this implementation. Don't overthink this
problem. Many of these methods can be implemented in very few lines.
Note that there are various types of int (64 or 32 bit, signed or
unsigned), and there are special atomic operations for each type. The
test does not check which type you use inside your waitgroup
implementation.

The tests for this problem is inside `waitgroup_test.go`. Please go back
to homework \#1 if you are having trouble running tests.


### Program Performance

Do not worry about the performance of your programs for this week. We
will reexamine these implementations for the next homework assignment
once you a better understanding about parallel performance, which is
covered in module 3.
