'use client';

import { useState } from 'react';
import { getQuizById, checkAnswer, getQuizProgress } from '../lib/quiz';

export default function Quiz({ quizId }) {
    const [currentQuestion, setCurrentQuestion] = useState(1);
    const [answers, setAnswers] = useState({});
    const [showResults, setShowResults] = useState(false);
    const [selectedOption, setSelectedOption] = useState(null);

    const quiz = getQuizById(quizId);
    if (!quiz) return <div>Quiz not found</div>;

    const question = quiz.questions.find(q => q.id === currentQuestion);
    const progress = getQuizProgress(quizId, answers);

    const handleAnswer = (optionIndex) => {
        setSelectedOption(optionIndex);
        setAnswers(prev => ({
            ...prev,
            [currentQuestion]: optionIndex
        }));
    };

    const handleNext = () => {
        if (currentQuestion < quiz.questions.length) {
            setCurrentQuestion(prev => prev + 1);
            setSelectedOption(null);
        } else {
            setShowResults(true);
        }
    };

    const handlePrevious = () => {
        if (currentQuestion > 1) {
            setCurrentQuestion(prev => prev - 1);
            setSelectedOption(answers[currentQuestion - 1] || null);
        }
    };

    if (showResults) {
        return (
            <div className="max-w-2xl mx-auto p-6 bg-white rounded-lg shadow-lg">
                <h2 className="text-2xl font-bold mb-4">Quiz Results</h2>
                <div className="text-lg mb-4">
                    <p>Total Questions: {progress.totalQuestions}</p>
                    <p>Correct Answers: {progress.correctAnswers}</p>
                    <p>Score: {progress.score.toFixed(1)}%</p>
                </div>
                <button
                    onClick={() => {
                        setShowResults(false);
                        setCurrentQuestion(1);
                        setAnswers({});
                        setSelectedOption(null);
                    }}
                    className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                >
                    Retry Quiz
                </button>
            </div>
        );
    }

    return (
        <div className="max-w-2xl mx-auto p-6 bg-white rounded-lg shadow-lg">
            <h2 className="text-2xl font-bold mb-4">{quiz.title}</h2>
            <p className="text-gray-600 mb-6">{quiz.description}</p>

            <div className="mb-6">
                <h3 className="text-xl font-semibold mb-4">
                    Question {currentQuestion} of {quiz.questions.length}
                </h3>
                <p className="text-lg mb-4">{question.question}</p>

                <div className="space-y-3">
                    {question.options.map((option, index) => (
                        <button
                            key={index}
                            onClick={() => handleAnswer(index)}
                            className={`w-full text-left p-4 rounded-lg border ${selectedOption === index
                                    ? 'bg-blue-100 border-blue-500'
                                    : 'hover:bg-gray-50 border-gray-200'
                                }`}
                        >
                            {option}
                        </button>
                    ))}
                </div>
            </div>

            <div className="flex justify-between">
                <button
                    onClick={handlePrevious}
                    disabled={currentQuestion === 1}
                    className={`px-4 py-2 rounded ${currentQuestion === 1
                            ? 'bg-gray-300 cursor-not-allowed'
                            : 'bg-gray-500 text-white hover:bg-gray-600'
                        }`}
                >
                    Previous
                </button>

                <button
                    onClick={handleNext}
                    disabled={selectedOption === null}
                    className={`px-4 py-2 rounded ${selectedOption === null
                            ? 'bg-gray-300 cursor-not-allowed'
                            : 'bg-blue-500 text-white hover:bg-blue-600'
                        }`}
                >
                    {currentQuestion === quiz.questions.length ? 'Finish' : 'Next'}
                </button>
            </div>
        </div>
    );
} 