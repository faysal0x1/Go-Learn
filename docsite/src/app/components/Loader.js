export default function Loader() {
    return (
        <div className="fixed inset-0 flex items-center justify-center bg-white bg-opacity-90 z-50">
            <div className="text-center">
                <div className="relative w-24 h-24">
                    {/* Go Gopher-inspired loading animation */}
                    <div className="absolute inset-0 border-4 border-blue-600 rounded-full animate-spin border-t-transparent"></div>
                    <div className="absolute inset-0 flex items-center justify-center">
                        <div className="w-16 h-16 bg-blue-600 rounded-full animate-pulse"></div>
                    </div>
                </div>
                <div className="mt-4 text-blue-600 font-semibold animate-pulse">
                    Loading Go Tutorials...
                </div>
            </div>
        </div>
    );
} 