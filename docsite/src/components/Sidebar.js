'use client'

import { useState, useEffect } from 'react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { motion, AnimatePresence } from 'framer-motion'
import {
    FiChevronDown,
    FiChevronRight,
    FiX,
    FiMenu,
    FiHome,
    FiBook,
    FiSettings,
    FiHelpCircle,
    FiLogOut,
    FiChevronLeft
} from 'react-icons/fi'
import { navigation } from '@/config/navigation'

export default function Sidebar({ isOpen, onClose }) {
    const [expandedSections, setExpandedSections] = useState({})
    const [collapsed, setCollapsed] = useState(false)
    const pathname = usePathname()

    useEffect(() => {
        const newExpandedSections = {}
        navigation.forEach(section => {
            if (section.items) {
                const isActive = section.items.some(item => pathname === item.href)
                if (isActive) newExpandedSections[section.title] = true
            }
        })
        setExpandedSections(newExpandedSections)
    }, [pathname])

    const toggleSection = (sectionTitle) => {
        setExpandedSections(prev => ({
            ...prev,
            [sectionTitle]: !prev[sectionTitle]
        }))
    }

    const isActive = (href) => pathname === href

    const isSectionActive = (section) => {
        if (section.href && isActive(section.href)) return true
        if (section.items) return section.items.some(item => isActive(item.href))
        return false
    }

    return (
        <>
            <AnimatePresence>
                {(isOpen || window.innerWidth >= 1024) && (
                    <>
                        {isOpen && (
                            <motion.div
                                initial={{ opacity: 0 }}
                                animate={{ opacity: 0.5 }}
                                exit={{ opacity: 0 }}
                                className="lg:hidden fixed inset-0 bg-black z-30"
                                onClick={() => onClose(false)}
                            />
                        )}

                        <motion.aside
                            initial={{ x: -300 }}
                            animate={{ x: 0 }}
                            exit={{ x: -300 }}
                            transition={{ type: 'spring', damping: 20 }}
                            className={`fixed inset-y-0 left-0 ${collapsed ? 'w-20' : 'w-64'} bg-gradient-to-b from-gray-900 to-gray-800 border-r border-gray-700 overflow-y-auto z-40 shadow-xl transition-all duration-300`}
                        >
                            <div className="flex flex-col h-full">
                                {/* Header */}
                                <div className="p-4 border-b border-gray-700 flex items-center justify-between">
                                    {!collapsed && (
                                        <Link href="/" className="flex items-center space-x-2">
                                            <div className="w-8 h-8 rounded-lg bg-blue-500 flex items-center justify-center">
                                                <FiBook className="text-white" />
                                            </div>
                                            <span className="text-xl font-bold text-white">GoLearn</span>
                                        </Link>
                                    )}
                                    {collapsed && (
                                        <div className="w-8 h-8 rounded-lg bg-blue-500 flex items-center justify-center mx-auto">
                                            <FiBook className="text-white" />
                                        </div>
                                    )}
                                    <button
                                        onClick={() => setCollapsed(!collapsed)}
                                        className="hidden lg:flex items-center justify-center w-8 h-8 rounded-full bg-gray-700 text-gray-300 hover:bg-gray-600 transition-colors"
                                    >
                                        {collapsed ? <FiChevronRight /> : <FiChevronLeft />}
                                    </button>
                                </div>

                                {/* Navigation */}
                                <nav className="flex-1 p-2 space-y-1 overflow-y-auto">
                                    {navigation.map((section) => (
                                        <div key={section.title}>
                                            {section.items ? (
                                                <>
                                                    <button
                                                        onClick={() => toggleSection(section.title)}
                                                        className={`w-full flex items-center ${collapsed ? 'justify-center' : 'justify-between'} px-3 py-3 text-sm rounded-lg transition-colors ${isSectionActive(section)
                                                            ? 'bg-blue-500/20 text-blue-400'
                                                            : 'text-gray-300 hover:bg-gray-700/50'
                                                            }`}
                                                    >
                                                        <div className="flex items-center">
                                                            <span className="text-lg">{section.icon || <FiHome />}</span>
                                                            {!collapsed && <span className="ml-3">{section.title}</span>}
                                                        </div>
                                                        {!collapsed && (expandedSections[section.title] ? (
                                                            <FiChevronDown className="w-4 h-4" />
                                                        ) : (
                                                            <FiChevronRight className="w-4 h-4" />
                                                        ))}
                                                    </button>

                                                    <AnimatePresence>
                                                        {!collapsed && expandedSections[section.title] && (
                                                            <motion.div
                                                                initial={{ height: 0, opacity: 0 }}
                                                                animate={{ height: 'auto', opacity: 1 }}
                                                                exit={{ height: 0, opacity: 0 }}
                                                                transition={{ duration: 0.2 }}
                                                                className="pl-11 mt-1 space-y-1"
                                                            >
                                                                {section.items.map((item) => (
                                                                    <Link
                                                                        key={item.href}
                                                                        href={item.href}
                                                                        className={`block px-3 py-2.5 text-sm rounded-lg transition-colors ${isActive(item.href)
                                                                            ? 'text-blue-400 font-medium bg-blue-500/10'
                                                                            : 'text-gray-400 hover:bg-gray-700/30'
                                                                            }`}
                                                                    >
                                                                        {item.title}
                                                                    </Link>
                                                                ))}
                                                            </motion.div>
                                                        )}
                                                    </AnimatePresence>
                                                </>
                                            ) : (
                                                <Link
                                                    href={section.href}
                                                    className={`flex items-center ${collapsed ? 'justify-center' : 'px-3'} py-3 text-sm rounded-lg transition-colors ${isActive(section.href)
                                                        ? 'bg-blue-500/20 text-blue-400'
                                                        : 'text-gray-300 hover:bg-gray-700/50'
                                                        }`}
                                                >
                                                    <span className="text-lg">{section.icon || <FiHome />}</span>
                                                    {!collapsed && <span className="ml-3">{section.title}</span>}
                                                </Link>
                                            )}
                                        </div>
                                    ))}
                                </nav>

                                {/* Footer */}
                                <div className="p-4 border-t border-gray-700">
                                    {!collapsed && (
                                        <div className="space-y-3">
                                            <Link
                                                href="/settings"
                                                className="flex items-center px-3 py-2 text-sm text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
                                            >
                                                <FiSettings className="text-lg" />
                                                <span className="ml-3">Settings</span>
                                            </Link>
                                            <Link
                                                href="/help"
                                                className="flex items-center px-3 py-2 text-sm text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
                                            >
                                                <FiHelpCircle className="text-lg" />
                                                <span className="ml-3">Help & Support</span>
                                            </Link>
                                            <button className="w-full flex items-center px-3 py-2 text-sm text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors">
                                                <FiLogOut className="text-lg" />
                                                <span className="ml-3">Logout</span>
                                            </button>
                                        </div>
                                    )}
                                    {collapsed && (
                                        <div className="flex flex-col items-center space-y-3">
                                            <Link href="/settings" className="p-2 text-gray-300 hover:bg-gray-700/50 rounded-lg">
                                                <FiSettings className="text-lg" />
                                            </Link>
                                            <Link href="/help" className="p-2 text-gray-300 hover:bg-gray-700/50 rounded-lg">
                                                <FiHelpCircle className="text-lg" />
                                            </Link>
                                            <button className="p-2 text-gray-300 hover:bg-gray-700/50 rounded-lg">
                                                <FiLogOut className="text-lg" />
                                            </button>
                                        </div>
                                    )}
                                    {!collapsed && (
                                        <div className="mt-4 pt-4 border-t border-gray-700 text-xs text-gray-500 text-center">
                                            v2.4.0 • © 2023 GoLearn
                                        </div>
                                    )}
                                </div>
                            </div>
                        </motion.aside>
                    </>
                )}
            </AnimatePresence>
        </>
    )
}