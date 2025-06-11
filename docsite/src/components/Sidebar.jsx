'use client'
import { useState } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { useTheme } from 'next-themes'
import { FiSun, FiMoon, FiMenu, FiX, FiChevronDown, FiChevronRight } from 'react-icons/fi'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { navigation } from '@/config/navigation'

export default function Sidebar() {
    const [isOpen, setIsOpen] = useState(false)
    const [expandedSections, setExpandedSections] = useState({})
    const { theme, setTheme } = useTheme()
    const pathname = usePathname()

    const toggleSection = (title) => {
        setExpandedSections(prev => ({
            ...prev,
            [title]: !prev[title]
        }))
    }

    const isActive = (href) => pathname === href

    return (
        <>
            {/* Mobile menu button */}
            <button
                className="fixed top-4 left-4 z-50 p-2 rounded-md lg:hidden"
                onClick={() => setIsOpen(!isOpen)}
            >
                {isOpen ? <FiX size={24} /> : <FiMenu size={24} />}
            </button>

            {/* Sidebar */}
            <motion.aside
                initial={{ x: -300 }}
                animate={{ x: isOpen ? 0 : -300 }}
                className="fixed top-0 left-0 z-40 w-72 h-screen bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 lg:translate-x-0 overflow-y-auto"
            >
                <div className="flex flex-col h-full">
                    {/* Logo */}
                    <div className="p-4 border-b border-gray-200 dark:border-gray-700">
                        <h1 className="text-xl font-bold text-gray-900 dark:text-white">
                            Go Programming
                        </h1>
                    </div>

                    {/* Navigation */}
                    <nav className="flex-1 p-4 space-y-1 overflow-y-auto">
                        {navigation.map((item) => (
                            <div key={item.title}>
                                {item.subsections ? (
                                    <>
                                        <button
                                            onClick={() => toggleSection(item.title)}
                                            className="flex items-center justify-between w-full px-4 py-2 text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md"
                                        >
                                            <span>{item.title}</span>
                                            {expandedSections[item.title] ? (
                                                <FiChevronDown className="w-4 h-4" />
                                            ) : (
                                                <FiChevronRight className="w-4 h-4" />
                                            )}
                                        </button>
                                        <AnimatePresence>
                                            {expandedSections[item.title] && (
                                                <motion.div
                                                    initial={{ height: 0, opacity: 0 }}
                                                    animate={{ height: 'auto', opacity: 1 }}
                                                    exit={{ height: 0, opacity: 0 }}
                                                    transition={{ duration: 0.2 }}
                                                    className="pl-4 space-y-1"
                                                >
                                                    {item.subsections.map((subsection) => (
                                                        <Link
                                                            key={subsection.href}
                                                            href={subsection.href}
                                                            className={`block px-4 py-2 text-sm text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md ${isActive(subsection.href)
                                                                ? 'bg-gray-100 dark:bg-gray-700 font-medium'
                                                                : ''
                                                                }`}
                                                        >
                                                            {subsection.title}
                                                        </Link>
                                                    ))}
                                                </motion.div>
                                            )}
                                        </AnimatePresence>
                                    </>
                                ) : (
                                    <Link
                                        href={item.href}
                                        className={`block px-4 py-2 text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md ${isActive(item.href)
                                            ? 'bg-gray-100 dark:bg-gray-700 font-medium'
                                            : ''
                                            }`}
                                    >
                                        {item.title}
                                    </Link>
                                )}
                            </div>
                        ))}
                    </nav>

                    {/* Theme toggle */}
                    <div className="p-4 border-t border-gray-200 dark:border-gray-700">
                        <button
                            onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
                            className="flex items-center justify-center w-full px-4 py-2 text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md"
                        >
                            {theme === 'dark' ? (
                                <FiSun className="w-5 h-5 mr-2" />
                            ) : (
                                <FiMoon className="w-5 h-5 mr-2" />
                            )}
                            {theme === 'dark' ? 'Light Mode' : 'Dark Mode'}
                        </button>
                    </div>
                </div>
            </motion.aside>

            {/* Overlay */}
            {isOpen && (
                <motion.div
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }}
                    className="fixed inset-0 z-30 bg-black bg-opacity-50 lg:hidden"
                    onClick={() => setIsOpen(false)}
                />
            )}
        </>
    )
} 