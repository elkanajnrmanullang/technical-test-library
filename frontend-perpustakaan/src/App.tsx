import { Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import Buku from './pages/Buku';
import Peminjaman from './pages/Peminjaman';
import Denda from './pages/Denda';
import Layout from './components/Layout';

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />
      <Route path="/login" element={<Login />} />
      <Route element={<Layout />}>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/buku" element={<Buku />} />
        <Route path="/peminjaman" element={<Peminjaman />} />
        <Route path="/denda" element={<Denda />} />
      </Route>
    </Routes>
  );
}