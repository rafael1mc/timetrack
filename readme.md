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
timer := timetrack.NewNode("My New Node")

// run some code, like:
time.Sleep(200 * time.Millisecond)

dur := timer.Stop()
fmt.Println("Took", dur, "to run")

// Took 201.063875ms to run
```

# Examples
I've created two examples, a simple one with children nodes and one to simulate tracking endpoint execution times. They are in the `examples` directory, and you can run them by cloning the repo and executing:
```
go run ./examples/simple
```
The terminal will show the results:
```
root node: 378.344666ms (100%)
  |-- child node: 77.28075ms (20.4%)
    |-- grandchild1 node: 51.100084ms (13.5%)
    |-- grandchild2 node: 26.049791ms (6.9%)
```
OR
```
go run ./examples/api
```
And then open your browser at [localhost:4000](http://localhost:4000).  
The terminal will show the results:
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

# FAQ
1. **Why don't the examples show exact milliseconds?**  
Well, tracking logic adds extra calls, which takes their own execution time, albeit very small.
2. **Why sometimes percentage doesn't add up to 100%?**  
Mainly for 2 reasons:  
    1. The default report percentage is related to the TOTAL execution time;
    2. The parent can be performing tasks  
So technically, the parent's execution time is AT LEAST its children execution time, not their sum.  
A better sum would be `(parent - children) + children`, recursively.

# License
Check [LICENSE](https://github.com/rafael1mc/timetrack/blob/main/LICENSE) file.