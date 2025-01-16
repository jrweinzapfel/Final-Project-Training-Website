import React from 'react';
import Link from 'react-router-dom';

export const Navbar = () => {
    return (
        <nav className="mt-4">
            <button
                        className="text-lg font-medium text-gray-700 hover:text-gray-900"
                        type="button"
                    >
                        <Link to="/">Home</Link>
                    </button>
            <ul className="flex justify-center space-x-4">
                <li>
                    <button
                        className="text-lg font-medium text-gray-700 hover:text-gray-900"
                        type="button"
                    >
                        <Link to="/training">Training</Link>
                    </button>
                </li>
                <li>
                    <button
                        className="text-lg font-medium text-gray-700 hover:text-gray-900"
                        type="button"
                    >
                        <Link to="/programs">Programs</Link>
                    </button>
                </li>
                <li>
                    <button
                        className="text-lg font-medium text-gray-700 hover:text-gray-900"
                        type="button"
                    >
                        <Link to="/blog">Blog</Link>
                    </button>
                </li>
                <li>
                    <button
                        className="text-lg font-medium text-gray-700 hover:text-gray-900"
                        type="button"
                    >
                        <Link to="/about">About</Link>
                    </button>
                </li>
                <li>
                    <button
                        className="text-lg font-medium text-gray-700 hover:text-gray-900"
                        type="button"
                    >
                        <Link to="/contact">Contact</Link>
                    </button>
                </li>
            </ul>
        </nav>
    );
};