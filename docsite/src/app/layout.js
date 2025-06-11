import './globals.css';
import Navigation from '../components/Navigation';

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
            {children}
          </main>
        </div>
      </body>
    </html>
  );
}
