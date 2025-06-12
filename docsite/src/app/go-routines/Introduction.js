export default function Introduction() {
    return (
        <>
            <h2 className="text-2xl font-bold mb-4">Goroutines in Go</h2>
            <p className="mb-4">
                Goroutines are lightweight threads managed by the Go runtime. They enable concurrent programming, allowing you to run multiple functions independently and efficiently.
            </p>
            <p className="mb-4">Key features of goroutines in Go:</p>
            <ul className="list-disc pl-6 space-y-2">
                <li>Start with the <code>go</code> keyword</li>
                <li>Efficient and lightweight (thousands per program)</li>
                <li>Communicate via channels</li>
                <li>Synchronization with WaitGroup, Mutex, and atomic operations</li>
                <li>Support for context cancellation and timeouts</li>
                <li>Patterns: worker pool, pipeline, fan-out/fan-in, producer-consumer</li>
            </ul>
            <p className="mt-4">
                Goroutines are fundamental to Go's concurrency model. They're not OS threads but are multiplexed onto a small number of OS threads by the Go runtime, making them very efficient.
            </p>
        </>
    );
}