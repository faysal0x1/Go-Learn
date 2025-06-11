'use client'
import { useState } from 'react';

export default function SectionLayout({
    title,
    introduction,
    examples,
    exercises,
    playground
}) {
    const [activeTab, setActiveTab] = useState('introduction');

    const tabs = [
        { id: 'introduction', label: 'Introduction' },
        { id: 'examples', label: 'Examples' },
        { id: 'exercises', label: 'Exercises' },
        { id: 'playground', label: 'Playground' },
    ];

    return (
        <div className="max-w-4xl mx-auto">
            <h1 className="text-4xl font-bold mb-8">{title}</h1>

            <div className="mb-8">
                <div className="border-b border-gray-200">
                    <nav className="-mb-px flex space-x-8">
                        {tabs.map((tab) => (
                            <button
                                key={tab.id}
                                onClick={() => setActiveTab(tab.id)}
                                className={`
                  py-4 px-1 border-b-2 font-medium text-sm
                  ${activeTab === tab.id
                                        ? 'border-blue-500 text-blue-600'
                                        : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                                    }
                `}
                            >
                                {tab.label}
                            </button>
                        ))}
                    </nav>
                </div>
            </div>

            <div className="prose max-w-none">
                {activeTab === 'introduction' && (
                    <div className="space-y-6">
                        {introduction}
                    </div>
                )}

                {activeTab === 'examples' && (
                    <div className="space-y-8">
                        {examples}
                    </div>
                )}

                {activeTab === 'exercises' && (
                    <div className="space-y-8">
                        {exercises}
                    </div>
                )}

                {activeTab === 'playground' && (
                    <div className="space-y-8">
                        {playground}
                    </div>
                )}
            </div>
        </div>
    );
} 