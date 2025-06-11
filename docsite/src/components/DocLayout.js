'use client'

import { useState, useEffect } from 'react'
import { usePathname } from 'next/navigation'
import { motion, AnimatePresence } from 'framer-motion'
import Sidebar from './Sidebar'
import TopBar from './TopBar'

export default function DocLayout({ children, title, description }) {
    const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false)
    const pathname = usePathname()

    // Close mobile menu when route changes
    useEffect(() => {
        setIsMobileMenuOpen(false)
    }, [pathname])

    return (
        <div className="min-h-screen bg-white dark:bg-gray-900 flex flex-col">
            <TopBar onMenuToggle={() => setIsMobileMenuOpen(!isMobileMenuOpen)} />

            <div className="flex flex-1">
                <Sidebar isOpen={isMobileMenuOpen} onClose={() => setIsMobileMenuOpen(false)} />

                <div className="flex-1 lg:ml-64">
                    <main className="py-6">
                        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                            <AnimatePresence mode="wait">
                                <motion.div
                                    key={pathname}
                                    initial={{ opacity: 0, y: 20 }}
                                    animate={{ opacity: 1, y: 0 }}
                                    exit={{ opacity: 0, y: -20 }}
                                    transition={{ duration: 0.2 }}
                                >
                                    <div className="mb-8">
                                        {title && (
                                            <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
                                                {title}
                                            </h1>
                                        )}
                                        {description && (
                                            <p className="mt-2 text-lg text-gray-600 dark:text-gray-300">
                                                {description}
                                            </p>
                                        )}
                                    </div>
                                    <div className="prose dark:prose-invert max-w-none">
                                        {children}
                                    </div>
                                </motion.div>
                            </AnimatePresence>
                        </div>
                    </main>
                </div>
            </div>
        </div>
    )
}