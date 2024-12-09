// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import BooksList from './components/BooksList';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<BooksList />} />
        {/* Add routes for users, transactions, etc. */}
      </Routes>
    </Router>
  );
};

export default App;
