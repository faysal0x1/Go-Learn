'use client'

import { useState } from 'react'
import { motion } from 'framer-motion'
import { FiPlay, FiLoader } from 'react-icons/fi'
import CodeBlock from './CodeBlock'

export default function Playground({ initialCode = '' }) {
    const [code, setCode] = useState(initialCode)
    const [output, setOutput] = useState('')
    const [isRunning, setIsRunning] = useState(false)
    const [error, setError] = useState('')

    const runCode = async () => {
        setIsRunning(true)
        setError('')
        setOutput('')

        try {
            // In a real implementation, you would send the code to a backend service
            // that compiles and runs the Go code. For now, we'll simulate it.
            await new Promise(resolve => setTimeout(resolve, 1000))

            // Simulate output
            if (code.includes('fmt.Println')) {
                const match = code.match(/fmt\.Println\("([^"]+)"\)/)
                if (match) {
                    setOutput(match[1])
                }
            } else {
                setOutput('Program executed successfully.')
            }
        } catch (err) {
            setError(err.message)
        } finally {
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

            <div className="p-4">
                <textarea
                    value={code}
                    onChange={(e) => setCode(e.target.value)}
                    className="w-full h-48 p-4 font-mono text-sm bg-gray-900 text-gray-100 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
                    placeholder="Write your Go code here..."
                />
            </div>

            <div className="p-4 bg-gray-50 dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700">
                <button
                    onClick={runCode}
                    disabled={isRunning}
                    className="flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
                >
                    {isRunning ? (
                        <FiLoader className="w-4 h-4 mr-2 animate-spin" />
                    ) : (
                        <FiPlay className="w-4 h-4 mr-2" />
                    )}
                    {isRunning ? 'Running...' : 'Run Code'}
                </button>
            </div>

            {(output || error) && (
                <div className="p-4 border-t border-gray-200 dark:border-gray-700">
                    <h4 className="text-sm font-medium mb-2 text-gray-700 dark:text-gray-300">
                        Output
                    </h4>
                    <pre className="p-4 bg-gray-100 dark:bg-gray-800 rounded-lg overflow-x-auto">
                        <code className={`text-sm ${error ? 'text-red-600 dark:text-red-400' : 'text-gray-800 dark:text-gray-200'}`}>
                            {error || output}
                        </code>
                    </pre>
                </div>
            )}
        </div>
    )
} 