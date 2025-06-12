'use client';

import { getAllQuizzes } from '../../lib/quiz';
import Link from 'next/link';

export default function QuizListPage() {
    const quizzes = getAllQuizzes();

    return (
        <div className="min-h-screen bg-gray-100 py-8">
            <div className="max-w-4xl mx-auto px-4">
                <h1 className="text-3xl font-bold mb-8">Available Quizzes</h1>
                <div className="grid gap-6 md:grid-cols-2">
                    {quizzes.map((quiz) => (
                        <Link
                            key={quiz.id}
                            href={`/quiz/${quiz.id}`}
                            className="block p-6 bg-white rounded-lg shadow-lg hover:shadow-xl transition-shadow"
                        >
                            <h2 className="text-xl font-semibold mb-2">{quiz.title}</h2>
                            <p className="text-gray-600">{quiz.description}</p>
                            <div className="mt-4 text-sm text-gray-500">
                                {quiz.questions.length} questions
                            </div>
                        </Link>
                    ))}
                </div>
            </div>
        </div>
    );
} 