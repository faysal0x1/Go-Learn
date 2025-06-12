'use client'

import { useState, useEffect, useRef } from 'react'
import { FiPlay, FiLoader, FiAlertTriangle } from 'react-icons/fi'
import CodeMirror from '@uiw/react-codemirror'
import { go } from '@codemirror/lang-go'

export default function Playground({ initialCode = '' }) {
    const [code, setCode] = useState(initialCode || `package main

import "fmt"

func main() {
  fmt.Println("Hello, Go Playground!")
}`)
    const [output, setOutput] = useState('')
    const [isRunning, setIsRunning] = useState(false)
    const [error, setError] = useState('')
    const [gopherJsLoaded, setGopherJsLoaded] = useState(false)
    const workerRef = useRef(null)

    // Load GopherJS WebAssembly
    useEffect(() => {
        if (typeof window !== 'undefined') {
            const script = document.createElement('script')
            script.src = 'https://cdn.jsdelivr.net/npm/gopherjs-wasm@latest/dist/gopherjs.js'
            script.onload = () => {
                window.GopherJS().then((gopherjs) => {
                    window.gopherjs = gopherjs
                    setGopherJsLoaded(true)
                })
            }
            script.onerror = () => {
                setError('Failed to load Go compiler. Please try again later.')
            }
            document.body.appendChild(script)

            return () => {
                document.body.removeChild(script)
            }
        }
    }, [])

    const runCode = async () => {
        if (!gopherJsLoaded) {
            setError('Go compiler is still loading. Please wait...')
            return
        }

        setIsRunning(true)
        setError('')
        setOutput('')

        try {
            // Create a worker for safer execution
            if (!workerRef.current) {
                workerRef.current = new Worker(
                    URL.createObjectURL(
                        new Blob(
                            [
                                `
                self.onmessage = async (e) => {
                  const { code } = e.data;
                  try {
                    // Capture console.log output
                    let output = '';
                    const originalConsoleLog = console.log;
                    console.log = (...args) => {
                      output += args.join(' ') + '\\n';
                      originalConsoleLog(...args);
                    };
                    
                    // Execute the Go code
                    await self.gopherjs.run(code);
                    
                    // Restore console.log
                    console.log = originalConsoleLog;
                    
                    postMessage({ output });
                  } catch (err) {
                    postMessage({ error: err.message });
                  }
                };
                `,
                            ],
                            { type: 'application/javascript' }
                        )
                    )
                )

                workerRef.current.onmessage = (e) => {
                    if (e.data.error) {
                        setError(e.data.error)
                    } else {
                        setOutput(e.data.output)
                    }
                    setIsRunning(false)
                }
            }

            // Post the code to the worker
            workerRef.current.postMessage({ code })
        } catch (err) {
            setError(err.message)
            setIsRunning(false)
        }
    }

    return (
        <div className="my-8 rounded-lg overflow-hidden border border-gray-200 dark:border-gray-700">
            <div className="p-4 bg-gray-50 dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white">
                    Go Playground
                </h3>
                <p className="text-sm text-gray-600 dark:text-gray-300">
                    Write and run Go code directly in your browser
                </p>
            </div>

            <div className="p-0">
                <CodeMirror
                    value={code}
                    height="300px"
                    extensions={[go()]}
                    onChange={(value) => setCode(value)}
                    //   theme={githubDark}
                    basicSetup={{
                        lineNumbers: true,
                        highlightActiveLineGutter: true,
                        highlightActiveLine: true,
                        foldGutter: true,
                        autocompletion: true,
                    }}
                />
            </div>

            <div className="p-4 bg-gray-50 dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 flex justify-between items-center">
                <button
                    onClick={runCode}
                    disabled={isRunning || !gopherJsLoaded}
                    className="flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
                >
                    {isRunning ? (
                        <FiLoader className="w-4 h-4 mr-2 animate-spin" />
                    ) : (
                        <FiPlay className="w-4 h-4 mr-2" />
                    )}
                    {isRunning ? 'Running...' : 'Run Code'}
                </button>
                {!gopherJsLoaded && (
                    <div className="text-sm text-gray-500 dark:text-gray-400">
                        Loading Go compiler...
                    </div>
                )}
            </div>

            {(output || error) && (
                <div className="p-4 border-t border-gray-200 dark:border-gray-700">
                    <h4 className="text-sm font-medium mb-2 text-gray-700 dark:text-gray-300 flex items-center">
                        {error ? (
                            <>
                                <FiAlertTriangle className="mr-1 text-red-500" />
                                Error
                            </>
                        ) : (
                            'Output'
                        )}
                    </h4>
                    <pre
                        className={`p-4 rounded-lg overflow-x-auto ${error
                            ? 'bg-red-50 dark:bg-red-900/20'
                            : 'bg-gray-100 dark:bg-gray-800'
                            }`}
                    >
                        <code
                            className={`text-sm font-mono ${error
                                ? 'text-red-600 dark:text-red-400'
                                : 'text-gray-800 dark:text-gray-200'
                                }`}
                        >
                            {error || output}
                        </code>
                    </pre>
                </div>
            )}
        </div>
    )
}