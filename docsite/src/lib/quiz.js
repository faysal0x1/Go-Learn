import quizData from '../data/quizzes.json';

export function getAllQuizzes() {
    return quizData.quizzes;
}

export function getQuizById(quizId) {
    return quizData.quizzes.find(quiz => quiz.id === quizId);
}

export function getQuestionById(quizId, questionId) {
    const quiz = getQuizById(quizId);
    if (!quiz) return null;
    return quiz.questions.find(q => q.id === questionId);
}

export function checkAnswer(quizId, questionId, selectedAnswer) {
    const question = getQuestionById(quizId, questionId);
    if (!question) return false;
    return question.correctAnswer === selectedAnswer;
}

export function getQuizProgress(quizId, answers) {
    const quiz = getQuizById(quizId);
    if (!quiz) return null;

    const totalQuestions = quiz.questions.length;
    const correctAnswers = Object.entries(answers).reduce((acc, [questionId, answer]) => {
        return acc + (checkAnswer(quizId, parseInt(questionId), answer) ? 1 : 0);
    }, 0);

    return {
        totalQuestions,
        correctAnswers,
        score: (correctAnswers / totalQuestions) * 100
    };
} 