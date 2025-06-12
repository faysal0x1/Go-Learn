export default function Exercises() {
    return (
        <>
            <h2 className="text-2xl font-bold mb-4">Exercises</h2>
            <div className="space-y-6">
                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 1: Parallel Sum</h3>
                    <p className="mb-4">Write a function that splits a slice into N parts and sums each part in a separate goroutine, then combines the results.</p>
                    <p className="text-sm text-gray-600">Hint: Use sync.WaitGroup to wait for all goroutines to complete.</p>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 2: Pipeline Pattern</h3>
                    <p className="mb-4">Implement a pipeline with three stages: generate numbers, square them, and double them, using goroutines and channels.</p>
                    <p className="text-sm text-gray-600">Hint: Create separate functions for each stage connected by channels.</p>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 3: Worker Pool</h3>
                    <p className="mb-4">Create a worker pool that processes jobs concurrently with a fixed number of workers.</p>
                    <p className="text-sm text-gray-600">Hint: Use a channel for jobs and another for results.</p>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 4: Graceful Shutdown</h3>
                    <p className="mb-4">Implement a server that can be started and gracefully shut down using channels and goroutines.</p>
                    <p className="text-sm text-gray-600">Hint: Use a quit channel to signal shutdown.</p>
                </div>

                <div className="bg-gray-50 p-6 rounded-lg">
                    <h3 className="text-xl font-semibold mb-2">Exercise 5: Error Handling</h3>
                    <p className="mb-4">Create goroutines that can return errors to the main thread using channels.</p>
                    <p className="text-sm text-gray-600">Hint: Use a channel of error type.</p>
                </div>
            </div>
        </>
    );
}