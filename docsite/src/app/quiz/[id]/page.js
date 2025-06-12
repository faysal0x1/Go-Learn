import Quiz from '../../../components/Quiz';

export default function QuizPage({ params }) {
    return (
        <div className="min-h-screen bg-gray-100 py-8">
            <Quiz quizId={params.id} />
        </div>
    );
} 