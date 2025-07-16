# Software Engineering Interview Theory Questions & Answers

## Table of Contents
1. [Data Structures & Algorithms](#data-structures--algorithms)
2. [Object-Oriented Programming](#object-oriented-programming)
3. [System Design](#system-design)
4. [Database Concepts](#database-concepts)
5. [Software Engineering Principles](#software-engineering-principles)
6. [Go-Specific Questions](#go-specific-questions)
7. [Concurrency & Parallelism](#concurrency--parallelism)
8. [Software Testing](#software-testing)
9. [Performance & Optimization](#performance--optimization)
10. [Security](#security)

---

## Data Structures & Algorithms

### Q1: What is the difference between an array and a linked list?

**Answer:**
- **Array:**
  - Elements stored in contiguous memory locations
  - Fixed size (in most languages)
  - O(1) random access by index
  - O(n) insertion/deletion at arbitrary positions
  - Better cache locality
  - Less memory overhead per element

- **Linked List:**
  - Elements stored in nodes with pointers to next element
  - Dynamic size
  - O(n) access to arbitrary element
  - O(1) insertion/deletion if you have reference to node
  - Poor cache locality
  - Extra memory overhead for pointers

### Q2: Explain the different types of time complexity notation.

**Answer:**
- **O(1) - Constant Time:** Algorithm takes same time regardless of input size
- **O(log n) - Logarithmic Time:** Time increases logarithmically with input size (binary search)
- **O(n) - Linear Time:** Time increases linearly with input size
- **O(n log n) - Log-linear Time:** Common in efficient sorting algorithms (merge sort, heap sort)
- **O(n²) - Quadratic Time:** Nested loops over input (bubble sort, insertion sort)
- **O(2ⁿ) - Exponential Time:** Time doubles with each additional input element

### Q3: What is a hash table and how does it work?

**Answer:**
A hash table is a data structure that implements an associative array abstract data type, mapping keys to values.

**How it works:**
1. Uses a hash function to compute an index into an array of buckets
2. Hash function converts key into array index
3. Handles collisions using chaining or open addressing
4. Average case: O(1) for search, insert, delete
5. Worst case: O(n) when many collisions occur

**Collision Resolution:**
- **Chaining:** Store multiple values in same bucket using linked list
- **Open Addressing:** Find next available slot (linear probing, quadratic probing)

---

## Object-Oriented Programming

### Q4: Explain the four pillars of OOP.

**Answer:**

1. **Encapsulation:**
   - Bundling data and methods that operate on that data within a single unit
   - Hiding internal implementation details
   - Provides data protection and interface abstraction

2. **Inheritance:**
   - Mechanism where new class acquires properties of existing class
   - Promotes code reusability
   - Creates "is-a" relationship between classes

3. **Polymorphism:**
   - Same interface can be used for different underlying data types
   - Method overriding and method overloading
   - Runtime polymorphism (dynamic binding) and compile-time polymorphism

4. **Abstraction:**
   - Hiding complex implementation details
   - Showing only essential features of object
   - Achieved through abstract classes and interfaces

### Q5: What is the difference between composition and inheritance?

**Answer:**

**Inheritance ("is-a" relationship):**
- Child class inherits properties from parent class
- Tight coupling between parent and child
- Can lead to fragile base class problem
- Single inheritance limitation in many languages

**Composition ("has-a" relationship):**
- Object contains instances of other objects
- Loose coupling between objects
- More flexible and easier to test
- Favors delegation over inheritance
- Can combine multiple behaviors

**Best Practice:** "Favor composition over inheritance" - more flexible and maintainable.

---

## System Design

### Q6: What are the key principles of scalable system design?

**Answer:**

1. **Horizontal Scaling:** Add more servers rather than upgrading existing ones
2. **Load Balancing:** Distribute incoming requests across multiple servers
3. **Caching:** Store frequently accessed data in fast storage (Redis, Memcached)
4. **Database Partitioning:** Shard data across multiple databases
5. **Microservices:** Break application into smaller, independent services
6. **CDN:** Content Delivery Network for static assets
7. **Asynchronous Processing:** Use message queues for time-consuming tasks
8. **Database Replication:** Master-slave setup for read scalability

### Q7: Explain CAP Theorem.

**Answer:**
CAP Theorem states that any distributed system can guarantee at most two of the following three properties:

- **Consistency:** All nodes see the same data simultaneously
- **Availability:** System remains operational and responsive
- **Partition Tolerance:** System continues despite network failures

**Examples:**
- **CP Systems:** Traditional RDBMS (MySQL, PostgreSQL)
- **AP Systems:** DNS, Web caching
- **CA Systems:** Single-node systems (not truly distributed)

In practice, partition tolerance is usually required, so choice is between consistency and availability.

---

## Database Concepts

### Q8: What is the difference between SQL and NoSQL databases?

**Answer:**

**SQL Databases:**
- Structured data with predefined schema
- ACID properties (Atomicity, Consistency, Isolation, Durability)
- Complex queries with JOINs
- Vertical scaling
- Examples: MySQL, PostgreSQL, Oracle

**NoSQL Databases:**
- Flexible schema or schema-less
- Eventually consistent (BASE properties)
- Simple queries, no complex JOINs
- Horizontal scaling
- Types: Document (MongoDB), Key-Value (Redis), Column-family (Cassandra), Graph (Neo4j)

### Q9: Explain database normalization.

**Answer:**
Database normalization is the process of organizing data to reduce redundancy and dependency.

**Normal Forms:**

1. **1NF (First Normal Form):**
   - Eliminate repeating groups
   - Each cell contains single value
   - Each row is unique

2. **2NF (Second Normal Form):**
   - Must be in 1NF
   - No partial dependencies on composite primary key

3. **3NF (Third Normal Form):**
   - Must be in 2NF
   - No transitive dependencies
   - Non-key attributes depend only on primary key

**Benefits:** Reduces data redundancy, improves data integrity
**Drawbacks:** May require more JOINs, potentially slower queries

---

## Software Engineering Principles

### Q10: Explain SOLID principles.

**Answer:**

1. **Single Responsibility Principle (SRP):**
   - Class should have only one reason to change
   - One class, one responsibility

2. **Open/Closed Principle (OCP):**
   - Software entities should be open for extension, closed for modification
   - Use interfaces and abstract classes

3. **Liskov Substitution Principle (LSP):**
   - Objects of superclass should be replaceable with objects of subclass
   - Inheritance should not break functionality

4. **Interface Segregation Principle (ISP):**
   - Clients should not depend on interfaces they don't use
   - Many specific interfaces are better than one general interface

5. **Dependency Inversion Principle (DIP):**
   - High-level modules should not depend on low-level modules
   - Both should depend on abstractions

### Q11: What is the difference between waterfall and agile methodologies?

**Answer:**

**Waterfall:**
- Sequential phases (Requirements → Design → Implementation → Testing → Deployment)
- Fixed scope, timeline, and budget
- Heavy documentation
- Changes are expensive
- Good for well-defined, stable requirements

**Agile:**
- Iterative and incremental development
- Flexible scope, fixed timeline
- Working software over documentation
- Embraces change
- Regular customer collaboration
- Popular frameworks: Scrum, Kanban

---

## Go-Specific Questions

### Q12: What are goroutines and how do they differ from threads?

**Answer:**
Goroutines are lightweight threads managed by Go runtime.

**Differences from OS threads:**
- **Memory:** Goroutines start with 2KB stack vs 1-2MB for threads
- **Creation Cost:** Cheaper to create and destroy
- **Scheduling:** Cooperatively scheduled by Go runtime, not OS
- **Communication:** Use channels for communication
- **Scalability:** Can have millions of goroutines

**Example:**
```go
go func() {
    fmt.Println("Running in goroutine")
}()
```

### Q13: Explain channels in Go and their types.

**Answer:**
Channels are Go's way of communicating between goroutines.

**Types:**
1. **Unbuffered Channels:** Synchronous communication
```go
ch := make(chan int)
```

2. **Buffered Channels:** Asynchronous up to buffer size
```go
ch := make(chan int, 5)
```

**Operations:**
- Send: `ch <- value`
- Receive: `value := <-ch`
- Close: `close(ch)`

**Patterns:**
- **Select Statement:** Non-blocking operations
- **Range over Channel:** Iterate until closed
- **Worker Pools:** Distribute work among goroutines

### Q14: What is the difference between arrays and slices in Go?

**Answer:**

**Arrays:**
- Fixed size declared at compile time
- Value types (copied when assigned)
- Size is part of the type
```go
var arr [5]int
```

**Slices:**
- Dynamic arrays with flexible size
- Reference types (point to underlying array)
- Has length and capacity
```go
var slice []int
slice = make([]int, 5, 10) // length 5, capacity 10
```

---

## Concurrency & Parallelism

### Q15: What is the difference between concurrency and parallelism?

**Answer:**

**Concurrency:**
- Dealing with multiple things at once
- About structure and composition
- Can run on single-core processor
- Goroutines taking turns on same thread

**Parallelism:**
- Doing multiple things at once
- About execution
- Requires multi-core processor
- Goroutines running simultaneously on different cores

**Go's Approach:**
- GOMAXPROCS controls parallelism
- Goroutines provide concurrency
- Runtime scheduler maps goroutines to OS threads

### Q16: What are race conditions and how to prevent them?

**Answer:**
Race condition occurs when multiple goroutines access shared data simultaneously and at least one modifies it.

**Prevention Methods:**

1. **Mutex (Mutual Exclusion):**
```go
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

2. **Channels:**
```go
// Use channels for communication instead of shared memory
```

3. **Atomic Operations:**
```go
import "sync/atomic"
atomic.AddInt64(&counter, 1)
```

4. **sync.Once:** Ensure code runs only once
```go
var once sync.Once
once.Do(func() {
    // initialization code
})
```

---

## Software Testing

### Q17: Explain different types of software testing.

**Answer:**

**By Level:**
1. **Unit Testing:** Test individual components in isolation
2. **Integration Testing:** Test component interactions
3. **System Testing:** Test complete system
4. **Acceptance Testing:** Validate against business requirements

**By Type:**
1. **Functional Testing:** Tests what system does
2. **Non-functional Testing:** Tests how system performs
   - Performance testing
   - Security testing
   - Usability testing

**By Approach:**
1. **Black Box Testing:** Test without knowing internal structure
2. **White Box Testing:** Test with knowledge of internal code
3. **Gray Box Testing:** Combination of both

### Q18: What is Test-Driven Development (TDD)?

**Answer:**
TDD is a development process where tests are written before the actual code.

**TDD Cycle (Red-Green-Refactor):**
1. **Red:** Write failing test
2. **Green:** Write minimal code to pass test
3. **Refactor:** Improve code while keeping tests passing

**Benefits:**
- Better code design
- Higher test coverage
- Confidence in code changes
- Living documentation

**Challenges:**
- Initial learning curve
- Time investment upfront
- Requires discipline

---

## Performance & Optimization

### Q19: What are common performance optimization techniques?

**Answer:**

**Algorithm Optimization:**
- Choose better algorithms and data structures
- Reduce time complexity
- Optimize space complexity

**Caching:**
- Memory caching (Redis, Memcached)
- HTTP caching (browser, CDN)
- Database query result caching

**Database Optimization:**
- Proper indexing
- Query optimization
- Connection pooling
- Database partitioning

**Code-Level Optimization:**
- Avoid premature optimization
- Profile before optimizing
- Reduce object allocations
- Use efficient data structures

**Go-Specific:**
- Use sync.Pool for object reuse
- Avoid memory leaks in goroutines
- Optimize garbage collection
- Use build constraints for different environments

### Q20: Explain memory management in Go.

**Answer:**

**Garbage Collector:**
- Concurrent, tri-color mark-and-sweep collector
- Automatic memory management
- Low-latency, designed for < 10ms pauses

**Memory Layout:**
- **Stack:** Local variables, function parameters
- **Heap:** Dynamic allocations, shared data
- **Escape Analysis:** Compiler decides stack vs heap allocation

**Best Practices:**
- Reuse objects with sync.Pool
- Avoid memory leaks (close channels, cancel contexts)
- Profile memory usage with pprof
- Use efficient data structures

---

## Security

### Q21: What are common web application security vulnerabilities?

**Answer:**

**OWASP Top 10:**
1. **Injection:** SQL injection, command injection
2. **Broken Authentication:** Session management flaws
3. **Sensitive Data Exposure:** Unencrypted data transmission
4. **XML External Entities (XXE):** XML parser vulnerabilities
5. **Broken Access Control:** Improper authorization
6. **Security Misconfiguration:** Default configurations
7. **Cross-Site Scripting (XSS):** Malicious script injection
8. **Insecure Deserialization:** Object injection attacks
9. **Using Components with Known Vulnerabilities**
10. **Insufficient Logging & Monitoring**

**Prevention:**
- Input validation and sanitization
- Parameterized queries
- HTTPS everywhere
- Regular security updates
- Security testing

### Q22: Explain different types of authentication and authorization.

**Answer:**

**Authentication (Who are you?):**
- **Basic Authentication:** Username/password over HTTP
- **Token-based:** JWT, OAuth tokens
- **Multi-factor Authentication:** Something you know + have + are
- **SSO:** Single Sign-On

**Authorization (What can you do?):**
- **Role-Based Access Control (RBAC):** Users assigned roles
- **Attribute-Based Access Control (ABAC):** Policy-based decisions
- **Access Control Lists (ACL):** Resource-specific permissions

**Best Practices:**
- Principle of least privilege
- Regular access reviews
- Secure token storage
- Session timeout

---

## Conclusion

This document covers fundamental software engineering concepts commonly asked in technical interviews. Understanding these concepts deeply, along with practical implementation experience, will help you succeed in software engineering interviews.

Remember to:
- Practice implementing these concepts in code
- Understand trade-offs and when to use each approach
- Be prepared to discuss real-world examples
- Stay updated with current technologies and best practices

For hands-on practice, refer to the coding problems in this repository's other directories.
