Rob Pike 的[Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)视频学习笔记 

## What is concurrency？
composition of independently executing computations
not parallelism

1. Concurrency is not parallelism
If you have only one processor, your program can still be concurrent but it can't be parallel.

2. Concurrency is a model for software construction
Easy to understand, use

## CSP 
Hoare 1978

## Goroutines
It’s an independent executing function, launched by a go statement.
It has its own call stack, which grows and shrinks as required.
It’s very cheap. It's practical to have thousands, even hundreds of thousands of goroutines.
It’s not a thread.
There maybe only one thread in a program with thousands of goroutines.
Instead, goroutines are multiplexed dynamically onto threads as needed to keep all the goroutines running.

## Communication with channel
Channel
A channel in Go provides a connection between two goroutines, allowing them to communicate.
first-class value

### Synchronization
A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.

### Buffered
Buffering removes synchronization.
Buffering makes them more like Erlang's mailboxes.

## THe Go approach
Don't communicate by sharing memory, share memory by communicating.
You don't have blob of memory and then put locks and mutexes and condition variables around it to protect it from parallel access. Instead you actually use the channel to pass the data back and forth between the goroutines and make your concurrent program operate that way.

## Patterns
### Generator: function that returns a channel
Channels are first-class values, just like integers or strings.

[Channels as a handle on a service](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/handleOnService)

[Multiplexing](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/multiplexing)

[Restoring sequencing](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/restoringSequencing)

    type Message struct {
        str string
        wait chan bool
    }

### select
All channels are evaluated.
Selection blocks until one communication can proceed, which then does.
If multiple can proceed, select chooses pseudo-randomly （假随机）.
A default clause, if present, executes immediately if no channel is ready.
```
select {
case <- c1:
case <- c2:
default:
}
```

[Fan-in again](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/FanInWithSelect)

Only one goroutine is needed.

[Timeout with select](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/timeout)

[Quit channel](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/quit)

[Receive on quit channel](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/receiveOnQuitChannel)

How do we know it's finished? Wait for it to tell us it's done: receive on the quit channel.

[Daisy-chain](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/channel/daisy-chain)

## Systems software
Google Search
Q: What does Google search do?
A: Give a query, return a page of search result (with some ads).
Q: How do we get the search results?
A: Send the query to Web search, Image search, YouTube, Maps, News etc., then mix the results.

Google Search: A fake framework
```
var (
Web = fakeSearch("web")
Image = fakeSearch("image")
Video = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
    return func(query string) Result {
        time.Sleep(time.second)
        return Result(fmt.Sprintf("%s result for %q\n", kind, query))
         }
}
}
```

Google Search 1.0
```
func Google(query string) (results []Result){
    results = append(results, Web(query))
    results = append(results, Image(query))
    results = append(results, Video(query))
    return
}
```

[Google Search 2.0](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/fakeSearch/googleSearch2.0)

[Google Search 2.1](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/fakeSearch/googleSearch2.1)

Don't wait for slow servers. No locks. No condition variables. No callbacks.

Avoid timeout
Q: How do we avoid discarding results from slow servers?
A: Replicate the servers. Send requests to multiple replicates, and use the first response.

[Google Search 3.0](https://github.com/kyl2016/GoCocurrencyPatterns/tree/master/fakeSearch/googleSearch3.0)

Reduce tail latency using replicated search servers.

And still...
No locks. No condition variables. No callbacks.

Summary
use Go's concurrency primitives to convert a 
- slow
- sequential 
- failure-sensitive
program into one that is
- fast
- concurrent
- replicated
- robust

Chatroulette toy: tinyurl.com/gochatroulette
Load balancer: tinyurl.com/goloadbalancer
Concurrent prime sieve: tinyurl.com/gosieve
Concurrent power series (by Mcllroy): tinyurl.com/gopowerseries

Conclusions
Goroutines and channels make it easy to express complex operations dealing with
- multiple inputs
- multiple outputs
- timeouts
- failure
- independent execution
- replication
- robustness


