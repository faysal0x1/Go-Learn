'use client'

import { FiMenu, FiSearch, FiBell, FiUser, FiSun, FiMoon, FiPlus } from 'react-icons/fi'
import { useTheme } from 'next-themes'
import { useEffect, useState } from 'react'
import { AnimatePresence } from "framer-motion";

export default function TopBar({ onMenuToggle }) {
    const { theme, setTheme } = useTheme()
    const [mounted, setMounted] = useState(false)
    const [searchOpen, setSearchOpen] = useState(false)

    useEffect(() => setMounted(true), [])

    return (
        <header className="sticky top-0 z-20 bg-white/80 dark:bg-gray-900/80 backdrop-blur-md border-b border-gray-200 dark:border-gray-800">
            <div className="px-4 sm:px-6 lg:px-8">
                <div className="flex items-center justify-between h-16">
                    {/* Left side */}
                    <div className="flex items-center">
                        <button
                            onClick={onMenuToggle}
                            className="mr-2 p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 rounded-full transition-colors lg:hidden"
                        >
                            <FiMenu className="w-5 h-5" />
                        </button>

                        {/* Search - mobile */}
                        <button
                            onClick={() => setSearchOpen(true)}
                            className="md:hidden p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 rounded-full"
                        >
                            <FiSearch className="w-5 h-5" />
                        </button>

                        {/* Search - desktop */}
                        <div className={`relative max-w-md text-center w-full hidden md:block transition-all duration-300 ${searchOpen ? 'opacity-100' : 'opacity-100'}`}>
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <FiSearch className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                                type="text"
                                placeholder="Search documentation..."
                                className="block w-full pl-10 pr-3 py-2 border border-gray-200 dark:border-gray-700 rounded-full leading-5 bg-gray-50 dark:bg-gray-800 placeholder-gray-500 dark:placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 dark:text-white"
                            />
                        </div>
                    </div>

                    {/* Right side */}
                    <div className="flex items-center space-x-3">
                        <button className="hidden md:flex items-center space-x-1.5 px-3 py-1.5 bg-gradient-to-r from-blue-500 to-blue-600 text-white text-sm font-medium rounded-full hover:from-blue-600 hover:to-blue-700 transition-all">
                            <FiPlus className="w-4 h-4" />
                            <span>New</span>
                        </button>

                        <button className="p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 rounded-full relative">
                            <FiBell className="w-5 h-5" />
                            <span className="absolute top-1 right-1 h-2 w-2 rounded-full bg-red-500 ring-2 ring-white dark:ring-gray-900"></span>
                        </button>

                        <button
                            onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
                            className="p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 rounded-full"
                        >
                            {mounted ? (
                                theme === 'dark' ? (
                                    <FiSun className="w-5 h-5" />
                                ) : (
                                    <FiMoon className="w-5 h-5" />
                                )
                            ) : (
                                <FiSun className="w-5 h-5" />
                            )}
                        </button>

                        <div className="relative">
                            <button className="flex items-center space-x-2 focus:outline-none group">
                                <div className="h-9 w-9 rounded-full bg-gradient-to-r from-blue-500 to-purple-500 flex items-center justify-center text-white">
                                    <span className="font-medium text-sm">JD</span>
                                </div>
                                <span className="hidden lg:inline-block text-sm font-medium text-gray-700 dark:text-gray-300 group-hover:text-gray-900 dark:group-hover:text-white">
                                    John Doe
                                </span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            {/* Mobile search overlay */}
            <AnimatePresence>
                {searchOpen && (
                    <motion.div
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                        exit={{ opacity: 0 }}
                        className="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4"
                        onClick={() => setSearchOpen(false)}
                    >
                        <motion.div
                            initial={{ y: 20, opacity: 0 }}
                            animate={{ y: 0, opacity: 1 }}
                            exit={{ y: 20, opacity: 0 }}
                            className="w-full max-w-md bg-white dark:bg-gray-800 rounded-xl p-4"
                            onClick={(e) => e.stopPropagation()}
                        >
                            <div className="relative">
                                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                    <FiSearch className="h-5 w-5 text-gray-400" />
                                </div>
                                <input
                                    type="text"
                                    placeholder="Search documentation..."
                                    autoFocus
                                    className="block w-full pl-10 pr-3 py-3 border border-gray-200 dark:border-gray-700 rounded-lg leading-5 bg-white dark:bg-gray-800 placeholder-gray-500 dark:placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 dark:text-white"
                                />
                            </div>
                        </motion.div>
                    </motion.div>
                )}
            </AnimatePresence>
        </header>
    )
}