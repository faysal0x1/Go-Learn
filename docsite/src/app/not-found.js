import Link from 'next/link';

export default function NotFound() {
    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="text-center p-8 bg-white rounded-lg shadow-lg max-w-md">
                <div className="mb-6">
                    <div className="text-6xl font-bold text-blue-600 mb-2">404</div>
                    <div className="text-2xl font-semibold text-gray-800 mb-4">Page Not Found</div>
                    <div className="text-gray-600 mb-8">
                        The page you're looking for doesn't exist or has been moved.
                    </div>
                </div>

                <div className="space-y-4">
                    <Link
                        href="/"
                        className="inline-block px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                    >
                        Return Home
                    </Link>

                    <div className="mt-4 text-sm text-gray-500">
                        <p>Looking for Go programming tutorials?</p>
                        <p className="mt-2">Check out our learning resources:</p>
                        <div className="mt-4 space-y-2">
                            <Link href="/variables" className="text-blue-600 hover:underline block">Variables</Link>
                            <Link href="/functions" className="text-blue-600 hover:underline block">Functions</Link>
                            <Link href="/structs" className="text-blue-600 hover:underline block">Structs</Link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
} 