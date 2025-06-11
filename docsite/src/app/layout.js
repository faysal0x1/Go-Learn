import './globals.css';
import Navigation from '../components/Navigation';
import { Suspense } from 'react';
import Loader from './components/Loader';

export const metadata = {
  title: 'Go Learning - Complete Guide',
  description: 'A comprehensive guide to learning Go programming language',
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <div className="flex min-h-screen">
          <Navigation />
          <main className="flex-1 ml-64 p-8">
            <Suspense fallback={<Loader />}>
              {children}
            </Suspense>
          </main>
        </div>
      </body>
    </html>
  );
}
