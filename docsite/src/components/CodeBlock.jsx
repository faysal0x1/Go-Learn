'use client'

import { useState } from 'react'
import { motion } from 'framer-motion'
import { FiCopy, FiCheck } from 'react-icons/fi'
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import { tomorrow } from 'react-syntax-highlighter/dist/cjs/styles/prism'

export default function CodeBlock({ code, language = 'go', showOutput = false, output = '' }) {
    const [copied, setCopied] = useState(false)

    const copyToClipboard = async () => {
        await navigator.clipboard.writeText(code)
        setCopied(true)
        setTimeout(() => setCopied(false), 2000)
    }

    return (
        <div className="my-4 rounded-lg overflow-hidden">
            <div className="flex items-center justify-between px-4 py-2 bg-gray-800 text-white">
                <span className="text-sm font-medium">{language}</span>
                <button
                    onClick={copyToClipboard}
                    className="p-1 rounded hover:bg-gray-700 transition-colors"
                >
                    {copied ? (
                        <FiCheck className="w-4 h-4 text-green-400" />
                    ) : (
                        <FiCopy className="w-4 h-4" />
                    )}
                </button>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="relative">
                    <SyntaxHighlighter
                        language={language}
                        style={tomorrow}
                        customStyle={{
                            margin: 0,
                            borderRadius: 0,
                            padding: '1rem',
                        }}
                    >
                        {code}
                    </SyntaxHighlighter>
                </div>
                {showOutput && (
                    <div className="p-4 bg-gray-100 dark:bg-gray-800 rounded-r-lg">
                        <h4 className="text-sm font-medium mb-2 text-gray-700 dark:text-gray-300">
                            Output
                        </h4>
                        <pre className="text-sm text-gray-800 dark:text-gray-200 whitespace-pre-wrap">
                            {output}
                        </pre>
                    </div>
                )}
            </div>
        </div>
    )
} 