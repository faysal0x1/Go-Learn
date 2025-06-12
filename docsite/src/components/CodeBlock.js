'use client'

import { useState } from 'react'
import { motion } from 'framer-motion'
import { FiCopy, FiCheck } from 'react-icons/fi'
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import {
    tomorrow,
    tomorrowNight,
    dracula,
    materialLight,
    materialDark,
    oneLight,
    oneDark
} from 'react-syntax-highlighter/dist/cjs/styles/prism'

const availableThemes = {
    'tomorrow': tomorrow,
    'tomorrow-night': tomorrowNight,
    'dracula': dracula,
    'material-light': materialLight,
    'material-dark': materialDark,
    'one-light': oneLight,
    'one-dark': oneDark
}

export default function CodeBlock({
    code,
    language = 'go',
    theme = 'tomorrow'
}) {
    const [copied, setCopied] = useState(false)
    const selectedTheme = availableThemes[theme] || tomorrow

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
                    aria-label="Copy code"
                >
                    {copied ? (
                        <FiCheck className="w-4 h-4 text-green-400" />
                    ) : (
                        <FiCopy className="w-4 h-4" />
                    )}
                </button>
            </div>
            <SyntaxHighlighter
                language={language}
                style={selectedTheme}
                customStyle={{
                    margin: 0,
                    borderRadius: 0,
                    padding: '1rem',
                }}
            >
                {code}
            </SyntaxHighlighter>
        </div>
    )
}