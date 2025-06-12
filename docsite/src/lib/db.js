import { MongoClient } from 'mongodb';
const uri = process.env.MONGODB_URI;
const client = new MongoClient(uri);

export async function connectToDatabase() {
    try {
        await client.connect();
        return client.db('go-quiz');
    } catch (error) {
        console.error('Error connecting to database:', error);
        throw error;
    }
}

export async function saveUserProgress(userId, quizResults) {
    const db = await connectToDatabase();
    const collection = db.collection('user-progress');

    return collection.updateOne(
        { userId },
        {
            $push: {
                quizHistory: {
                    date: new Date(),
                    ...quizResults
                }
            },
            $setOnInsert: { userId }
        },
        { upsert: true }
    );
}

export async function getUserProgress(userId) {
    const db = await connectToDatabase();
    const collection = db.collection('user-progress');

    return collection.findOne({ userId });
}

export async function getAnsweredQuestions(userId) {
    const db = await connectToDatabase();
    const collection = db.collection('user-progress');

    const userProgress = await collection.findOne({ userId });
    if (!userProgress) return [];

    return userProgress.quizHistory.flatMap(quiz =>
        quiz.questions.map(q => q.questionId)
    );
} 