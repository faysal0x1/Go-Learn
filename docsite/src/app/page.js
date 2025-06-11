"use client";
import { motion } from 'framer-motion'
import Link from 'next/link'
import { FiArrowRight, FiCode, FiBook, FiPlay } from 'react-icons/fi'

export default function Home() {
  return (
    <div className="max-w-4xl mx-auto">
      <div className="text-center mb-12">
        <h1 className="text-5xl font-bold mb-4">Welcome to Go Learning</h1>
        <p className="text-xl text-gray-600">
          A comprehensive guide to learning the Go programming language
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-2xl font-bold mb-4">Getting Started</h2>
          <p className="text-gray-600 mb-4">
            Begin your Go journey with our structured learning path, starting from the basics
            and progressing to advanced topics.
          </p>
          <ul className="list-disc pl-6 space-y-2">
            <li>Hello World and Basic Syntax</li>
            <li>Variables and Constants</li>
            <li>Control Structures</li>
            <li>Functions and Methods</li>
          </ul>
        </div>

        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-2xl font-bold mb-4">Advanced Topics</h2>
          <p className="text-gray-600 mb-4">
            Master advanced Go concepts and best practices for building robust applications.
          </p>
          <ul className="list-disc pl-6 space-y-2">
            <li>Concurrency with Goroutines</li>
            <li>Interfaces and Type Systems</li>
            <li>Error Handling</li>
            <li>Testing and Benchmarking</li>
          </ul>
        </div>

        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-2xl font-bold mb-4">Interactive Learning</h2>
          <p className="text-gray-600 mb-4">
            Practice what you learn with our interactive playgrounds and exercises.
          </p>
          <ul className="list-disc pl-6 space-y-2">
            <li>Code Examples</li>
            <li>Hands-on Exercises</li>
            <li>Interactive Playground</li>
            <li>Real-world Projects</li>
          </ul>
        </div>

        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-2xl font-bold mb-4">Best Practices</h2>
          <p className="text-gray-600 mb-4">
            Learn Go's idiomatic patterns and best practices for writing clean, efficient code.
          </p>
          <ul className="list-disc pl-6 space-y-2">
            <li>Code Organization</li>
            <li>Error Handling</li>
            <li>Concurrency Patterns</li>
            <li>Performance Optimization</li>
          </ul>
        </div>
      </div>
    </div>
  );
}
