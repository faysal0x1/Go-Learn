import Link from 'next/link';

const sections = [
    { id: 1, title: 'Hello World', path: '/hello-world' },
    // { id: 2, title: 'Simple Values', path: '/simple-values' },
    { id: 3, title: 'Variables', path: '/variables' },
    { id: 4, title: 'Constants', path: '/constants' },
    { id: 5, title: 'For Loop', path: '/for-loop' },
    { id: 6, title: 'If Else', path: '/if-else' },
    { id: 7, title: 'Switches', path: '/switches' },
    { id: 8, title: 'Arrays', path: '/arrays' },
    { id: 9, title: 'Slices', path: '/slices' },
    { id: 10, title: 'Maps', path: '/maps' },
    { id: 11, title: 'Range', path: '/range' },
    { id: 12, title: 'Functions', path: '/functions' },
    { id: 13, title: 'Variadic Functions', path: '/variadic-functions' },
    { id: 14, title: 'Closures', path: '/closures' },
    { id: 15, title: 'Pointers', path: '/pointers' },
    { id: 16, title: 'Structs', path: '/structs' },
    { id: 17, title: 'Interfaces', path: '/interfaces' },
    { id: 18, title: 'Enums', path: '/enums' },
    { id: 19, title: 'Generics', path: '/generics' },
    { id: 20, title: 'Go Routines', path: '/go-routines' },
];

export default function Navigation() {
    return (
        <nav className="w-64 h-screen bg-gray-800 text-white p-4 fixed left-0 top-0 overflow-y-auto">
            <div className="mb-8">
                <h1 className="text-2xl font-bold">Go Learning</h1>
                <p className="text-sm text-gray-400">Complete Guide</p>
            </div>
            <ul className="space-y-2">
                {sections.map((section) => (
                    <li key={section.id}>
                        <Link
                            href={section.path}
                            className="block px-4 py-2 rounded hover:bg-gray-700 transition-colors"
                        >
                            {section.title}
                        </Link>
                    </li>
                ))}
            </ul>
        </nav>
    );
} 