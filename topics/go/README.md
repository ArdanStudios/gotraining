## Ultimate Go

This is material for any intermediate-level developer who has some experience with other programming languages and wants to learn Go. We believe these classes are perfect for anyone who wants a jump start in learning Go or who wants a more thorough understanding of the language and its internals.

[Ultimate Go](../courses/go/README.md)

## Design Guidelines

You must develop a design philosophy that establishes a set of guidelines. This is more important than developing a set of rules or patterns you apply blindly. Guidelines help to formulate, drive and validate decisions. You can't begin to make the best decisions without understanding the impact of your decisions. Every decision you make, every line of code you write comes with trade-offs.

* [Prepare Your Mind](https://github.com/ardanlabs/gotraining/tree/master/topics/go#prepare-your-mind)
* [Productivity vs Performance](https://github.com/ardanlabs/gotraining/tree/master/topics/go#productivity-vs-performance)
* [Correctness vs Performance](https://github.com/ardanlabs/gotraining/tree/master/topics/go#correctness-vs-performance)
* [Code Reviews](https://github.com/ardanlabs/gotraining/tree/master/topics/go#code-reviews)
* [Data Oriented Design](https://github.com/ardanlabs/gotraining/tree/master/topics/go#data-oriented-design)
* [Interface And Composition Design](https://github.com/ardanlabs/gotraining/tree/master/topics/go#interface-and-composition-design)
* [Package Oriented Design](https://github.com/ardanlabs/gotraining/tree/master/topics/go#package-oriented-design)
* [Concurrent Software Design](https://github.com/ardanlabs/gotraining/tree/master/topics/go#concurrent-software-design)
* [Channel Design](https://github.com/ardanlabs/gotraining/tree/master/topics/go#channel-design)

---

### Prepare Your Mind

**Somewhere Along The Line**  
* We became impressed with programs that contain large amounts of code.
* We strived to create large abstractions in our code base.
* We forgot that the hardware is the platform.
* We lost the understanding that every decision comes with a cost.

**These Days Are Gone**  
* We can throw more hardware at the problem.
* We can throw more developers at the problem.

**Aspire To**  
* Be a champion for quality, efficiency and simplicity.
* Have a point of view.
* Value introspection and self-review.

**Open Your Mind**  
* Technology changes quickly but people's minds change slowly.
* Easy to adopt new technology but hard to adopt new ways of thinking.

---

### Productivity vs Performance

Productivity and performance both matter, but in the past you couldn’t have both. You needed to choose one over the other. We naturally gravitated to productivity, with the idea or hope that the hardware would resolve our performance problems for free. This movement towards productivity has resulted in the design of programming languages that produce sluggish software that is out pacing the hardware’s ability to make them faster.

By following Go’s idioms and a few guidelines, we can write code that can be reasoned about by anyone who looks at it. We can write software that simplifies, minimizes and reduces the amount of code we need to solve the problems we are working on. We don’t have to choose productivity over performance or performance over productivity anymore. We can have both.

**Quotes**

_"The hope is that the progress in hardware will cure all software ills. However, a critical observer may observe that software manages to outgrow hardware in size and sluggishness. Other observers had noted this for some time before, indeed the trend was becoming obvious as early as 1987." - Niklaus Wirth_

_"The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry." - Henry Petroski (2015)_

_"The hardware folks will not put more cores into their hardware if the software isn’t going to use them, so, it is this balancing act of each other staring at each other, and we are hoping that Go is going to break through on the software side.” - Rick Hudson (2015)_

_"C is the best balance I've ever seen between power and expressiveness. You can do almost anything you want to do by programming fairly straightforwardly and you will have a very good mental model of what's going to happen on the machine; you can predict reasonably well how quickly it's going to run, you understand what's going on .... - Brian Kernighan (2000)_

_"The trend in programming language design has been to create languages that enhance software reliability and programmer productivity. What we should do is develop languages alongside sound software engineering practices so the task of developing reliable programs is distributed throughout the software lifecycle, especially into the early phases of system design." - Al Aho (2009)_

---

### Correctness vs Performance

You want to write code that is optimized for correctness. Don't make coding decisions based on what you think might perform better. You must benchmark or profile to know if code is not fast enough. Then and only then should you optimize for performance. This can't be done until you have something working.

Improvement comes from writing code and thinking about the code you write. Then refactoring the code to make it better. This requires the help of other people to also read the code you are writing. Prototype ideas first to validate them. Try different approaches or ask others to attempt a solution. Then compare what you have learned.

Too many developers are not prototyping their ideas first before writing production code. It is through prototyping that you can validate your thoughts, ideas and designs. This is the time when you can break down walls and figure out how things work. Prototype in the concrete and consider contracts after you have a working prototype.

Refactoring must become part of the development cycle. Refactoring is the process of improving the code from the things that you learn on a daily basis. Without time to refactor, code will become impossible to manage and maintain over time. This creates the legacy issues we are seeing today.

**Quotes**

_"The correctness of the implementation is the most important concern, but there is no royal road to correctness. It involves diverse tasks such as thinking of invariants, testing and code reviews. Optimization should be done, but not prematurely." - Al Aho (inventor of AWK)_

_"The basic ideas of good style, which are fundamental to write clearly and simply, are just as important now as they were 35 years ago. Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code." - Brian Kernighan_

_"Unless the developer has a really good idea of what the software is going to be used for, there's a very high probablility that the software will turn out badly. If the developers don't know and understand the application well, then it's crucial to get as much user input and experience as possible." - Brian Kernighan_

_"The hardest bugs are those where your mental model of the situation is just wrong, so you can't see the problem at all" - Brian Kernighan_

_"There are two kinds of software projects: those that fail, and those that turn into legacy horrors." - Peter Weinberger (inventor of AWK)_

_"Legacy software is an unappreciated but serious problem. Legacy code may be the downfall of our civilization." - Chuck Moore (inventor of Forth)_

**Resources:**

[Prototype your design!](https://www.youtube.com/watch?v=vLxX3yZmw5Q) - Robert Griesemer

---

### Code Reviews

You can't look at a piece of code, function or algorithm and determine if it smells good or bad without a design philosophy. These four major categories are the basis for code reviews and should be prioritized in this order: Integrity, Readability, Simplicity and then Performance. You must consciously and with great reason be able to explain the category you are choosing.

**_Note: There are exceptions to everything but when you are not sure an exception applies, follow the guidelines presented the best you can._** 

#### 1) Integrity

**_We need to become very serious about reliability._**

There are two driving forces behind integrity:  
* Integrity is about every allocation, read and write of memory being accurate, consistent and efficient. The type system is critical to making sure we have this `micro` level of integrity.
* Integrity is about every data transformation being accurate, consistent and efficient. Writing less code and error handling is critical to making sure we have this `macro` level of integrity.

**Error Handling:**

When error handling is treated as an exception and not part of the main code, you can expect the majority of your critical failures to be due to error handling.

48 critical failures were found in a study looking at a couple hundred bugs in Cassandra, HBase, HDFS, MapReduce, and Redis.  
* 92% : Failures from bad error handling
    * 35% : Incorrect handling
        * 25% : Simply ignoring an error
        * 8% : Catching the wrong exception
        * 2% : Incomplete TODOs
    * 57% System specific
        * 23% : Easily detectable
        * 34% : Complex bugs
* 8% : Failures from latent human errors

**Ignorance vs Carelessness:**

Anytime we identify an integrity issue we need to ask ourselves why it happened. 
```
                    Not Deliberate               Deliberate
              ------------------------------------------------------
              |                          |                         |
              |                          |                         |
   Ignorance  |  Learning / Prototyping  |    Hacking / Guessing   |
              |                          |                         |
              |                          |                         |
              |-----------------------------------------------------
              |                          |                         |
              |                          |                         |
Carelessness  |        Education         |   Dangerous Situation   |
              |                          |                         |
              |                          |                         |
              ------------------------------------------------------
```

**Write Less Code:**

There have been studies that have researched the number of bugs you can expect to have in your software. The industry average is around 15 to 50 bugs per 1000 lines of code. One simple way to reduce the number of bugs, and increase the integrity of your software, is to write less code.

Bjarne Stroustrup stated that writing more code than you need results in `Ugly`, `Large` and `Slow` code:

* `Ugly`: Leaves places for bugs to hide.
* `Large`: Ensures incomplete tests.
* `Slow`: Encourages the use of shortcuts and dirty tricks.

**Resources:**

[Software Development for Infrastructure](http://www.stroustrup.com/Software-for-infrastructure.pdf) - Bjarne Stroustrup  
[Normalization of Deviance in Software](http://danluu.com/wat/) - danluu.com  
[Lessons learned from reading postmortems](http://danluu.com/postmortem-lessons/) - danluu.com  
[Technical Debt Quadrant](https://martinfowler.com/bliki/TechnicalDebtQuadrant.html) - Martin Fowler  
[Design Philosophy On Integrity](https://www.goinggo.net/2017/02/design-philosophy-on-integrity.html) - William Kennedy  
[Ratio of bugs per line of code](https://www.mayerdan.com/ruby/2012/11/11/bugs-per-line-of-code-ratio) - Dan Mayer  
[Masterminds of Programming](http://dl.acm.org/citation.cfm?id=1592983) - Federico Biancuzzi and Shane Warden  

#### 2) Readability

**_We must structure our systems to be more comprehensible._**  
**_Readability means reliability._**

This is about writing simple code that is easy to read and understand without the need of mental exhaustion. Just as important, it's about not hiding the cost/impact of the code per line, function, package and the overall ecosystem it runs in.

In Go, the underlying machine is the real machine rather than a single abstract machine. The model of computation is that of the computer. Here is the key, Go gives you direct access to the machine while still providing abstraction mechanisms to allow higher-level ideas to be expressed.

**_"A well-designed language has a one-one correlation between source code and object code. It's obvious to the programmer what code will be generated from their source. This provides its own satisfaction, is efficient, and reduces the need for documentation." - Chuck Moore (inventor of Forth)_**

**_"Can you explain it to the median user (developer)? as opposed to will the smartest user (developer) figure it out?" - Peter Weinberger (inventor of AWK)_**

[Example Readability Issue](http://cpp.sh/6i7d)  

#### 3) Simplicity

**_We must understand that simplicity is hard to design and complicated to build._**  

This is about hiding complexity. A lot of care and design must go into simplicity because this can cause more problems then good. It can create issues with readability and it can cause issues with performance. Validate that abstractions are not generalized or even too concise. You might think you are helping the programmer or the code but validate things are still easy to use, understand, debug and maintain.

**_"The simple fact is that complexity will emerge somewhere, if not in the language definition, then in thousands of applications and libraries." - Bjarne Stroustrup (inventor of C++)_**

**_"Everything should be made as simple as possible, but not simpler." - Albert Einstein_**

**Resources:**

[Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM) - Rob Pike  

#### 4) Performance

**_We must compute less to get the results we need._**

This is about not wasting effort and achieving execution efficiency. Writing code that is mechanically sympathetic with the runtime, operating system and hardware. Achieving performance by writing less and more efficient code but staying within the idioms and framework of the language.

Rules of Performance:   
    * Never guess about performance.  
    * Measurements must be relevant.  
    * Profile before you decide something is performance critical.  
    * Test to know you are correct.

[Example Benchmark](https://github.com/ardanlabs/gotraining/blob/master/topics/go/testing/benchmarks/basic/basic_test.go)  

**_"When we're computer programmers we're concentrating on the intricate little fascinating details of programming and we don't take a broad engineering point of view about trying to optimize the total system. You try to optimize the bits and bytes." - Tom Kurtz (inventor of BASIC)_**

#### 5) Micro-Optimizations

Micro-Optimizations are about squeezing every ounce of performance as possible. When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic. You are writing clever code that may require the unsafe package or you may need to drop into assembly.

[Example Micro Optimization](https://play.golang.org/p/D_bImirgXL)

---

### Data-Oriented Design

**Design Philosophy:**

* If you don't understand the data, you don't understand the problem.
* All problems are unique and specific to the data you are working with.
* Data transformations are at the heart of solving problems. Each function, method and work-flow must focus on implementing the specific data transformations required to solve the problems.
* If your data is changing, your problems are changing. When your problems are changing, the data transformations needs to change with it.
* Uncertainty about the data is not a license to guess but a directive to STOP and learn more.
* Solving problems you don't have, creates more problems you now do.
* If performance matters, you must have mechanical sympathy for how the hardware and operating system work.
* Minimize, simplify and REDUCE the amount of code required to solve each problem. Do less work by not wasting effort.
* Code that can be reasoned about and does not hide execution costs can be better understood, debugged and performance tuned.
* Coupling data together and writing code that produces predictable access patterns to the data will be the most performant.
* Changing data layouts can yield more significant performance improvements than changing just the algorithms.
* Efficiency is obtained through algorithms but performance is obtained through data structures and layouts.

**Resources:**

[Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc) - Mike Acton  
[Efficiency with Algorithms, Performance with Data Structures](https://www.youtube.com/watch?v=fHNmRkzxHWs) - Chandler Carruth

---

### Interface And Composition Design

**Design Philosophy:**

* Interfaces give programs structure.
* Interfaces encourage design by composition.
* Interfaces enable and enforce clean divisions between components.
    * The standardization of interfaces can set clear and consistent expectations.
* Decoupling means reducing the dependencies between components and the types they use.
    * This leads to correctness, quality and performance.
* Interfaces allow you to group concrete types by what they do.
    * Don't group types by a common DNA but by a common behavior.
    * Everyone can work together when we focus on what we do and not who we are.
* Interfaces help your code decouple itself from change.
    * You must do your best to understand what could change and use interfaces to decouple.
    * Interfaces with more than one method have more than one reason to change.
    * Uncertainty about change is not a license to guess but a directive to STOP and learn more.
* You must distinguish between code that:
    * defends against fraud vs protects against accidents

**Validation:**

Declare an interface when:  
* users of the API need to provide an implementation detail.
* API’s have multiple implementations they need to maintain internally.
* parts of the API that can change have been identified and require decoupling.

Don't declare an interface:  
* for the sake of using an interface.
* to generalize an algorithm.
* when users can declare their own interfaces.
* if it's not clear how the ineterface makes the code better.

**Resources:**

[Methods, interfaces and Embedding](https://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy  
[Composition with Go](https://www.goinggo.net/2015/09/composition-with-go.html) - William Kennedy  
[Reducing type hierarchies](https://www.goinggo.net/2016/10/reducing-type-hierarchies.html) - William Kennedy  
[Interface pollution in Go](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275) - Burcu Dogan  
[Application Focused API Design](https://www.goinggo.net/2016/11/application-focused-api-design.html) - William Kennedy  
[Avoid interface pollution](https://www.goinggo.net/2016/10/avoid-interface-pollution.html) - William Kennedy  

---

### Package-Oriented Design

_Package Oriented Design allows a developer to identify where a package belongs inside a Go project and the design guidelines the package must respect. It defines what a Go project is and how a Go project is structured. Finally, it improves communication between team members and promotes clean package design and project architecture that is discussable._

**Language Mechanics:**

* Packaging directly conflicts with how we have been taught to organize source code in other languages.
* In other languages, packaging is a feature that you can choose to use or ignore.
* You can think of packaging as applying the idea of microservices on a source tree.
* All packages are "first class," and the only hierarchy is what you define in the source tree for your project.
* There needs to be a way to “open” parts of the package to the outside world.
* Two packages can’t cross-import each other. Imports are a one way street. 

**Design Philosophy:**

* To be purposeful, packages must provide, not contain.
    * Packages must be named with the intent to describe what it provides.
    * Packages must not become a dumping ground of disparate concerns.
* To be usable, packages must be designed with the user as their focus.
    * Packages must be intuitive and simple to use.
    * Packages must respect their impact on resources and performance.
    * Packages must protect the user’s application from cascading changes.
    * Packages must prevent the need for type assertions to the concrete.
    * Packages must reduce, minimize and simplify its code base.
* To be portable, packages must be designed with reusability in mind.
    * Packages must aspire for the highest level of portability.
    * Packages must reduce taking on opinions when it’s reasonable and practical.
    * Packages must not become a single point of dependency.

**Project Structure:**

```
Kit                     Application

├── CONTRIBUTORS        ├── cmd/
├── LICENSE             ├── internal/
├── README.md           │   └── platform/
├── cfg/                └── vendor/
├── examples/
├── log/
├── pool/
├── tcp/
├── timezone/
├── udp/
└── web/
```

* **vendor/**  
Good documentation for the `vendor/` folder can be found in this Gopher Academy [post](https://blog.gopheracademy.com/advent-2015/vendor-folder) by Daniel Theophanes. For the purpose of this post, all the source code for 3rd party packages need to be vendored (or copied) into the `vendor/` folder. This includes packages that will be used from the company `Kit` project. Consider packages from the `Kit` project as 3rd party packages.

* **cmd/**  
All the programs this project owns belongs inside the `cmd/` folder. The folders under `cmd/` are always named for each program that will be built. Use the letter `d` at the end of a program folder to denote it as a daemon. Each folder has a matching source code file that contains the `main` package.

* **internal/**  
Packages that need to be imported by multiple programs within the project belong inside the `internal/` folder. One benefit of using the name `internal/` is that the project gets an extra level of protection from the compiler. No package outside of this project can import packages from inside of `internal/`. These packages are therefore internal to this project only.

* **internal/platform/**  
Packages that are foundational but specific to the project belong in the `internal/platform/` folder. These would be packages that provide support for things like databases, authentication or even marshaling.

**Validation:**

<u>Validate the location of a package.</u>
* `Kit`
    * Packages that provide foundational support for the different `Application` projects that exist.
    * logging, configuration or web functionality.
* `cmd/`
    * Packages that provide support for a specific program that is being built.
    * startup, shutdown and configuration.
* `internal/`
    * Packages that provide support for the different programs the project owns.
    * CRUD, services or business logic.
* `internal/platform/`
    * Packages that provide internal foundational support for the project..
    * database, authentication or marshaling.
    
<u>Validate the dependency choices.</u>
* `All`
    * Validate the cost/benefit of each dependency.
    * Question imports for the sake of sharing existing types.
    * Question imports to others packages at the same level.
    * If a package wants to import another package at the same level:
        * Question the current design choices of these packages.
        * If reasonable, move the package inside the source tree for the package that wants to import it.
        * Use the source tree to show the dependency relationships.
* `internal/`
    * Packages from these locations CAN’T be imported:
        * `cmd/`
* `internal/platform/`
    * Packages from these locations CAN’T be imported:
        * `cmd/`
        * `internal/`
        
<u>Validate the opinions being imposed.</u>
* `Kit`, `internal/platform/`
    * NOT allowed to have opinions about any application concerns.
    * NOT allowed to log, but access to trace information must be decoupled.
    * Configuration and runtime changes must be decoupled.
    * Retrieving metric and telemetry values must be decoupled.
* `cmd/`, `internal/`
    * Allowed to have opinions about any application concerns.
    * Allowed to log and handle configuration natively.
    
<u>Validate how data is accepted/returned.</u>
* `All`
    * Validate the consistent use of value/pointer semantics for a given type.
    * When using an interface type to accept a value, the focus must be on the behavior that is required and not the value itself.
    * If behavior is not required, use a concrete type.
    * When reasonable, use an existing type before declaring a new one.
    * Question types from dependencies that leak into the exported API.
        * An existing type may no longer be reasonable to use.
        
<u>Validate how errors are handled.</u>
* `All`
    * Handling an error means:
        * The error is no longer a concern.
        * There is no more action that needs to be taken.
        * It has been logged if necessary.
* `Kit`
    * NOT allowed to panic an application.
    * NOT allowed to wrap errors.
    * Return only root cause error values.
* `cmd/`
    * Allowed to panic an application.
    * Wrap errors with context if not being handled.
    * Majority of handling errors happen here.
* `internal/`
    * NOT allowed to panic an application.
    * Wrap errors with context if not being handled.
    * Minority of handling errors happen here.
* `internal/platform/`
    * NOT allowed to panic an application.
    * NOT allowed to wrap errors.
    * Return only root cause error values.

**Resources:**

[Design Philosophy on Packaging](https://www.goinggo.net/2017/02/design-philosophy-on-packaging.html)  
[Package Oriented Design](https://www.goinggo.net/2017/02/package-oriented-design.html)  

---

### Concurrent Software Design

Concurrency is about managing multiple things at once. Like one person washing the dishes while they are also cooking dinner. You're making progress on both but you're only ever doing one of those things at the same time. Parallelism is about doing multiple things at once. Like one person cooking and placing dirty dishes in the sink, while another washes the dishes. They are happening at the same time.

Both you and the runtime have a responsibility in managing the concurrency of the application. You are responsible for managing these three things when writing concurrent software:

**Design Philosophy:**

* The application must startup and shutdown with integrity.
    * Know how and when every goroutine you create terminates.
    * All goroutines you create should terminate before main returns.
    * Applications should be capable of shutting down on demand, even under load, in a controlled way.
        * You want to stop accepting new requests and finish the requests you have (load shedding).
* Identify and monitor critical points of back pressure that can exist inside your application.
    * Channels, mutexes and atomic functions can create back pressure when goroutines are required to wait.
    * A little back pressure is good, it means there is a good balance of concerns.
    * A lot of back pressure is bad, it means things are imbalanced.
    * Back pressure that is imbalanced will cause:
        * Failures inside the software and across the entire platform.
        * Your application to collapse, implode or freeze.
    * Measuring back pressure is a way to measure the health of the application.
* Rate limit to prevent overwhelming back pressure inside your application.
    * Every system has a breaking point, you must know what it is for your application.
    * Applications should reject new requests as early as possible once they are overloaded.
        * Don’t take in more work than you can reasonably work on at a time.
        * Push back when you are at critical mass. Create your own external back pressure.
    * Use an external system for rate limiting when it is reasonable and practical.
* Use timeouts to release the back pressure inside your application.
    * No request or task is allowed to take forever.
    * Identify how long users are willing to wait.
    * Higher-level calls should tell lower-level calls how long they have to run.
    * At the top level, the user should decide how long they are willing to wait.
    * Use the `Context` package.
        * Functions that users wait for should take a `Context`.
            * These functions should select on <-ctx.Done() when they would otherwise block indefinitely.
        * Set a timeout on a `Context` only when you have good reason to expect that a function's execution has a real time limit.
        * Allow the upstream caller to decide when the `Context` should be canceled.
        * Cancel a `Context` whenever the user abandons or explicitly aborts a call.
* Architect applications to:
    * Identify problems when they are happening.
    * Stop the bleeding.
    * Return the system back to a normal state.

---

### Channel Design

Channels allow goroutines to communicate with each other through the use of signaling semantics. Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.

**Language Mechanics:**

* Use channels to orchestrate and coordinate goroutines.
    * Focus on the signaling semantics and not the sharing of data.
    * Signal by passing data or by closing the channel.
    * Question their use for synchronizing access to shared state.
        * _There are cases where channels can be simpler for this but initially question._
* Unbuffered channels:
    * Blocks the sending goroutine from moving forward until a different goroutine has received the data signal.
        * The sending goroutine gains immediate knowledge their data signal has been received.
    * Both goroutines involved must be at the channel at the same time.
        * More important: The Receive happens before the Send.
    * Trade-offs:
        * We take the benefit of knowing the data signal has been received for the cost of higher blocking latency.
* Buffered channels:
    * Does NOT block the sending goroutine from moving forward until a different goroutine has received the data signal.
        * The sending goroutine gains no knowledge that their data signal has been received.
    * Both goroutines involved don't need to be at the channel at the same time.
        * More important: The Send happens before the Receive.
    * Trade-offs:
        * We take the benefit of reducing blocking latency for the cost of not knowing if/when the data signal is received.
* Closing channels:
    * Signaling without the need for data passing.
    * Perfect for signaling cancellations and deadlines.
* NIL channels:
    * Turn off signaling
    * Perfect for rate limiting or short term stoppages.

**Design Philosophy:**

Depending on the problem you are solving, you may require different channel semantics. Depending on the semantics you need, different architectural choices must be taken.

* If any given Send on a channel `CAN` cause the goroutine to block:
    * You have less buffers compared to the number of goroutines that will perform a Send at any given time.
    * An example would be a resource pool of database connections.
    * This requires knowing what happens when the Send blocks because this will create a situation of back pressure inside the application in front of the channel.
    * The things discussed above about [writing concurrent software](https://github.com/ardanlabs/gotraining/blob/master/reading/design_guidelines.md#writing-concurrent-software) must be taken into account for this channel.
    * Not knowing what happens when the Send blocks on the channel is not a license to guess but a directive to STOP, understand and take appropriate measures.
* If any given Send on a channel `WON'T` cause the goroutine to block:
    * You have a buffer for every goroutine that will perform a Send.
    * You will abandon the Send immediately if it can't be performed.
    * An example would be a fan out pattern or pipelining.
    * This requires knowing if the size of the buffer is appropriate and if it is acceptable to abandon the Send.
* Less is more with buffers.
    * Don’t think about performance when thinking about buffers.
    * Buffers can help to reduce blocking latency between signaling.
        * Reducing blocking latency towards zero does not necessarily mean better throughput.
        * If a buffer of one is giving you good enough throughput then keep it.
        * Question buffers that are larger than one and measure for size.
        * Find the smallest buffer possible that provides good enough throughput.