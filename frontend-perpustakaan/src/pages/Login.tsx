import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../lib/axios';
import { BookText } from 'lucide-react';

export default function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      const response = await api.post('/login', { username, password });
      
      const token = response.data?.token || response.data?.data?.token;

      if (token) {
        localStorage.setItem('token', token);
        navigate('/dashboard');
      } else {
        setError('Token tidak ditemukan pada respons server.');
      }
    } catch (err: any) {
      setError(err.response?.data?.msg || err.response?.data?.message || 'Login gagal. Periksa kembali kredensial Anda.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-zinc-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <div className="flex justify-center">
          <div className="w-12 h-12 bg-zinc-900 rounded-lg flex items-center justify-center">
            <BookText className="w-7 h-7 text-white" />
          </div>
        </div>
        <h2 className="mt-6 text-center text-3xl font-bold tracking-tight text-zinc-900">
          Sistem Perpustakaan
        </h2>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10 border border-zinc-200">
          <form className="space-y-6" onSubmit={handleLogin}>
            {error && (
              <div className="p-3 bg-red-50 border border-red-200 text-red-600 rounded-md text-sm">
                {error}
              </div>
            )}
            
            <div>
              <label className="block text-sm font-medium text-zinc-700">Username</label>
              <div className="mt-1">
                <input
                  type="text"
                  required
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  className="block w-full appearance-none rounded-md border border-zinc-300 px-3 py-2 placeholder-zinc-400 shadow-sm focus:border-zinc-900 focus:outline-none focus:ring-zinc-900 sm:text-sm"
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-zinc-700">Password</label>
              <div className="mt-1">
                <input
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="block w-full appearance-none rounded-md border border-zinc-300 px-3 py-2 placeholder-zinc-400 shadow-sm focus:border-zinc-900 focus:outline-none focus:ring-zinc-900 sm:text-sm"
                />
              </div>
            </div>

            <div>
              <button
                type="submit"
                disabled={isLoading}
                className="flex w-full justify-center rounded-md border border-transparent bg-zinc-900 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-zinc-800 focus:outline-none focus:ring-2 focus:ring-zinc-900 focus:ring-offset-2 disabled:opacity-70"
              >
                {isLoading ? 'Memverifikasi...' : 'Masuk Sistem'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}