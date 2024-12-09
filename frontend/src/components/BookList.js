// src/components/BooksList.js
import React, { useEffect, useState } from 'react';
import { getBooks } from '../services/api';

const BooksList = () => {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const { data } = await getBooks();
        setBooks(data);
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    };
    fetchBooks();
  }, []);

  return (
    <div>
      <h1>Books</h1>
      <ul>
        {books.map((book) => (
          <li key={book.ID}>
            {book.Name} by {book.Author} ({book.Publication})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BooksList;
