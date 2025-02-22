# timetrack
A library to track execution time with ease. Don't be fooled by the long readme xD.

The API is designed to be plug and play. 

# Quick Start
Add the dependency:
```
go get github.com/rafael1mc/timetrack
```

Initialize a timer:
```
package main
func main(){
    timer := timetrack.NewNode("My New Node")

    // run some code, like:
    time.Sleep(200 * time.Millisecond)

    dur := timer.Stop()
    fmt.Println("Took", dur, "to run")

    // Took 201.063875ms to run
}
```

# Examples
I've created 3 examples.
 - simple: `go run ./examples/simple`
 - endpoint: `go run ./examples/api` (access [localhost:4000](http://localhost:4000))
 - custom report: `go run ./examples/custom_report`

This is a sample output for the endpoint example:
```
middleware: 755.155875ms (100%)
  |-- controller: 755.123ms (100.0%)
    |-- serviceFunc1: 201.836834ms (26.6%)
      |-- repositoryFunc: 100.943417ms (13.2%)
    |-- serviceFunc2: 252.188541ms (33.4%)
      |-- serviceFunc2 sub measurement: 51.0455ms (6.8%)
    |-- serviceFunc3: 301.079ms (39.9%)
```

# API
As mentioned, the API is desgined to be simple, plug and play. You can easily start as many timers as you want, or add children to an existing timer.

## Node
 - `timetrack.NewNode("my timer")`: creates a root node with that name
 - `timer.Branch(name)`: create a child node with that name inside `timer`
 - `timer.Stop()`: stop the timer and all of its children. Returns the duration of `timer`
 - `timer.CurrentDuration()`: a way to access duration without stopping the timer
 - `timer.String()`: will stop the timer and use the default reporter for generating the output, which prints all chidren durations, indentend

## Context
There are helper functions to enrich a context:
 - `BranchFrom(context, name)`: pass a context and a name of the new child. If it finds a timer in the context, it will use that, otherwise it will spawn a new root timer. The returned context already has this timer in it
 - `WithTimeNode(context, node)`: returns a new context with the given node

## Reporter
The default reporter will indent the children, printing their names, the duration and the % it represents in the total of the node being reported.

To provide your own reporter, just implement `Report(node *TimeNode) string` and pass it to `timer.SetReporter`.  

# FAQ
1. **Why don't the examples show exact milliseconds?**  
Well, tracking logic adds extra calls, which takes their own execution time, albeit very small.
2. **Why sometimes percentage doesn't add up to 100%?**  
Mainly for 2 reasons:  
    1. The default report percentage is related to the TOTAL execution time;
    2. The parent can be performing tasks  
So technically, the parent's execution time is AT LEAST its children execution time, not their sum.  
A better sum would be `(parent - children) + children`, recursively.

# TODO
 - [ ] Test concurrency safety (specially root variables)
 - [ ] Add secondary percentage to compare just chidren
 - [ ] Provide abstraction for generating [Server Timming](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server-Timing)
 - [ ] Maybe add thresholds to trigger warnings (eg: >300ms warn)

# License
Check [LICENSE](https://github.com/rafael1mc/timetrack/blob/main/LICENSE) file.