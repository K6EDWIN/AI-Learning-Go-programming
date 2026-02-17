<h1>Go Sentinel: Concurrent Performance &amp; Connectivity Monitor</h1>

<p>
  <strong>Go Sentinel</strong> is a high-performance utility designed to monitor network services.
  Developed as part of a deep dive into the Go ecosystem, it focuses on building non-blocking,
  memory-efficient systems suitable for large-scale informatics projects.
</p>

<hr />

<h2> Project Overview</h2>

<p>
  This project implements a <strong>Concurrent Web Scraper Engine</strong>.
  By moving away from traditional class-based inheritance and adopting Go’s composition model,
  the engine achieves high throughput with minimal overhead.
</p>

<p>
  The system leverages Go’s <strong>CSP (Communicating Sequential Processes)</strong> model to execute
  concurrent network checks while maintaining clean error handling and modular architecture.
</p>

<hr />

<h2> Key Features</h2>

<ul>
  <li>
    <strong>Concurrency-First Architecture:</strong>
    Utilizes Goroutines and Channels to scale network requests without blocking the main execution thread.
  </li>
  <li>
    <strong>Defensive Design:</strong>
    Implements Go's “Verbose Guardrails,” ensuring every <code>nil</code> error is audited to prevent
    memory leaks and security vulnerabilities.
  </li>
  <li>
    <strong>High-Speed Compilation:</strong>
    Leverages Go’s fast feedback loop for rapid iterative design, bridging scripting ease and compiled performance.
  </li>
  <li>
    <strong>Standardized Layout:</strong>
    Adheres to a minimalist directory structure for improved package management and reusability.
  </li>
</ul>

<hr />

<h2> Setup Instructions</h2>

<h3>System Requirements</h3>

<ul>
  <li><strong>Operating System:</strong> Windows, macOS, or Linux</li>
  <li><strong>Go Version:</strong> 1.23+</li>
</ul>

<p>Verify your Go installation:</p>

<pre><code>go version</code></pre>

<h3>Installation &amp; Execution</h3>

<ol>
  <li>
    <strong>Clone the repository</strong>
    <pre><code>git clone [https://github.com/K6EDWIN/AI-Learning-Go-programming.git]
cd AI-Learning-Go-programming</code></pre>
  </li>

  <li>
    <strong>Initialize the Go module (if required)</strong>
    <pre><code>go mod init learning-go
go mod tidy</code></pre>
  </li>

  <li>
    <strong>Run the application</strong>
    <pre><code>go run cmd/app.go</code></pre>
  </li>

  <li>
    <strong>Build the executable</strong>
    <p>Windows:</p>
    <pre><code>go build -o bin/sentinel.exe cmd/app.go</code></pre>

    <p>macOS / Linux:</p>
    <pre><code>go build -o bin/sentinel cmd/app.go</code></pre>
  </li>
</ol>

<hr />

<h2> Minimal Working Example</h2>

<p>
  The following snippet from <code>cmd/app.go</code> demonstrates the core concurrent engine:
</p>

<pre><code>// Launching Goroutines to scale network requests
for _, link := range links {
    go check(link, c)
}

// Receiving results from the channel
for i := 0; i &lt; len(links); i++ {
    fmt.Println(&lt;-c)
}</code></pre>

<p>This pattern demonstrates:</p>

<ul>
  <li>Parallel execution using Goroutines</li>
  <li>Channel-based communication</li>
  <li>Explicit error auditing</li>
  <li>Non-blocking program flow</li>
</ul>

<hr />

<h2> AI Prompt Journal &amp; Reflections</h2>

<h3>Project Structuring</h3>
<p>
  AI was used to refine the standard Go project layout (<code>cmd/</code>, <code>package/</code>, <code>bin/</code>),
  improving modularity and alignment with Go community conventions.
</p>

<h3>Concurrency Optimization</h3>
<p>
  AI assisted in transitioning from sequential HTTP polling to a Goroutine and Channel-based concurrency model,
  significantly improving scalability and responsiveness.
</p>

<h3>Error Handling Philosophy</h3>
<p>
  Go’s explicit error-handling pattern:
</p>

<pre><code>if err != nil {
    return err
}</code></pre>

<p>
  This approach forces defensive programming and promotes system reliability,
  especially in network-dependent applications.
</p>

<h3>Prompt Engineering Insight</h3>
<p>
  Providing clear architectural goals (e.g., concurrency-first, memory-efficient design)
  resulted in more accurate and optimized AI-assisted solutions.
</p>

<hr />

<h2> Tech Stack</h2>

<ul>
  <li><strong>Language:</strong> Go (Golang) 1.23+</li>
  <li>
    <strong>Standard Library:</strong>
    <ul>
      <li><code>net/http</code> – Networking</li>
      <li><code>sync</code> – Concurrency</li>
      <li><code>time</code> – Benchmarking</li>
    </ul>
  </li>
  <li><strong>Architecture Model:</strong> CSP (Channels + Goroutines)</li>
  <li><strong>Environment:</strong> Strathmore University Informatics &amp; Computer Science</li>
</ul>

<hr />

<h2> Final Directory Map</h2>

<pre><code>AI-Learning-Go-programming/
├── bin/                 # Compiled standalone executables
├── cmd/
│   └── app.go           # Entry point: Concurrent Site Checker
├── documentation/       # Research logs &amp; Notion backups
├── package/             # Reusable internal logic
├── go.mod               # Project identity &amp; dependency management
└── README.md            # Project documentation</code></pre>
