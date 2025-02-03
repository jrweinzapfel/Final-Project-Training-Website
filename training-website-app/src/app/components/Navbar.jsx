"use client"
import Link from 'next/link';
import { usePathname } from 'next/navigation';

export const Navbar = () => {
    const pathname = usePathname();
    return (
        <nav className="mt-4">
            <Link href="/" className={pathname === "/" ? "font-bold mr-4" : "mr-4 text-blue-500"}>Home</Link>
            <Link href="/blog" className={pathname === "/blog" ? "font-bold mr-4" : "mr-4 text-blue-500"}>Blog</Link>
            <Link href="/contact" className={pathname === "/contact" ? "font-bold mr-4" : "mr-4 text-blue-500"}>Contact</Link>
            <Link href="/about" className={pathname === "/about" ? "font-bold mr-4" : "mr-4 text-blue-500"}>About</Link>
            <Link href="/training" className={pathname === "/training" ? "font-bold mr-4" : "mr-4 text-blue-500"}>Training</Link>
            <Link href="/programs" className={pathname === "/programs" ? "font-bold mr-4" : "mr-4 text-blue-500"}>Programs</Link>
        </nav>
    );
};