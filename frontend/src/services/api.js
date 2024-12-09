import axios from 'axios';

const API = axios.create({
  baseURL: 'http://localhost:9010', 
});

export const getBooks = () => API.get('/book');
export const getBookById = (id) => API.get(`/book/${id}`);
export const createBook = (book) => API.post('/book', book);
export const updateBook = (id, book) => API.put(`/book/${id}`, book);
export const deleteBook = (id) => API.delete(`/book/${id}`);

export const getUsers = () => API.get('/user');
export const getUserById = (id) => API.get(`/user/${id}`);
export const createUser = (user) => API.post('/user', user);

export const getTransactions = () => API.get('/transaction');
export const getTransactionById = (id) => API.get(`/transaction/${id}`);
export const createTransaction = (transaction) => API.post('/transaction', transaction);
