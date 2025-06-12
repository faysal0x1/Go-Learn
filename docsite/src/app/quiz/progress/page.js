'use client';

import { useState, useEffect } from 'react';
import { getUserProgress } from '../../../lib/db';

export default function ProgressPage() {
    const [progress, setProgress] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const userId = localStorage.getItem('userId') || 'anonymous';
        getUserProgress(userId).then((data) => {
            setProgress(data);
            setLoading(false);
        });
    }, []);

    if (loading) {
        return (
            <div className="max-w-2xl mx-auto p-6">
                <div className="bg-white rounded-lg shadow-lg p-8 text-center">
                    <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
                    <p className="mt-4 text-gray-600">Loading progress...</p>
                </div>
            </div>
        );
    }

    if (!progress || !progress.quizHistory || progress.quizHistory.length === 0) {
        return (
            <div className="max-w-2xl mx-auto p-6">
                <div className="bg-white rounded-lg shadow-lg p-8 text-center">
                    <h2 className="text-2xl font-bold mb-4">No Quiz History</h2>
                    <p className="text-gray-600 mb-6">
                        You haven't taken any quizzes yet. Start your learning journey now!
                    </p>
                    <a
                        href="/quiz"
                        className="inline-block bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700"
                    >
                        Take Quiz
                    </a>
                </div>
            </div>
        );
    }

    return (
        <div className="max-w-2xl mx-auto p-6">
            <div className="bg-white rounded-lg shadow-lg p-8">
                <h2 className="text-2xl font-bold mb-6">Quiz Progress</h2>

                <div className="space-y-6">
                    {progress.quizHistory.map((quiz, index) => (
                        <div key={index} className="border-b pb-4 last:border-b-0">
                            <div className="flex justify-between items-center mb-2">
                                <span className="text-gray-600">
                                    {new Date(quiz.date).toLocaleDateString()}
                                </span>
                                <span className="font-semibold">
                                    Score: {quiz.isCorrect ? '1' : '0'} / 1
                                </span>
                            </div>
                            <div className="text-sm text-gray-500">
                                Question ID: {quiz.questionId}
                            </div>
                        </div>
                    ))}
                </div>

                <div className="mt-8 text-center">
                    <a
                        href="/quiz"
                        className="inline-block bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700"
                    >
                        Take Another Quiz
                    </a>
                </div>
            </div>
        </div>
    );
} 