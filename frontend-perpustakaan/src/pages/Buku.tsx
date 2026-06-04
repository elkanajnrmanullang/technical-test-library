import { useState, useEffect } from 'react';
import { api } from '../lib/axios';
import { Plus, Search, Edit2, Trash2, BookText, Info, X } from 'lucide-react';

export default function Buku() {
  const [books, setBooks] = useState<any[]>([]);
  const [authors, setAuthors] = useState<any[]>([]);
  const [publishers, setPublishers] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const [toast, setToast] = useState({ show: false, message: '' });

  const fetchData = async () => {
    try {
      setIsLoading(true);
      const [bukuRes, authorRes, pubRes] = await Promise.allSettled([
        api.get('/buku'),
        api.get('/admin/buku/author'),
        api.get('/admin/buku/penbuk')
      ]);

      if (bukuRes.status === 'fulfilled') {
        const data = bukuRes.value.data?.data || bukuRes.value.data || [];
        setBooks(Array.isArray(data) ? data : []);
      }

      if (authorRes.status === 'fulfilled') {
        const data = authorRes.value.data?.data || authorRes.value.data || [];
        setAuthors(Array.isArray(data) ? data : []);
      }

      if (pubRes.status === 'fulfilled') {
        const data = pubRes.value.data?.data || pubRes.value.data || [];
        setPublishers(Array.isArray(data) ? data : []);
      }
    } catch (error) {
      setBooks([]);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const booksWithDetails = books.map(buku => {
    const authorName = authors.find(a => String(a.id || a.id_penulis) === String(buku.id_penulis_buku))?.penulis_buku || buku.id_penulis_buku || '-';
    const publisherName = publishers.find(p => String(p.id || p.id_penerbit) === String(buku.id_penerbit_buku))?.penerbit_buku || buku.id_penerbit_buku || '-';
    
    return {
      ...buku,
      authorName,
      publisherName
    };
  });

  const filteredBooks = booksWithDetails.filter((buku) => {
    const judul = (buku.judul_buku || '').toLowerCase();
    const penulis = (buku.authorName || '').toLowerCase();
    const penerbit = (buku.publisherName || '').toLowerCase();
    const search = searchTerm.toLowerCase();
    
    return judul.includes(search) || penulis.includes(search) || penerbit.includes(search);
  });

  const handleAction = (action: string) => {
    setToast({ show: true, message: `Tindakan ${action} dinonaktifkan pada sesi demonstrasi ini.` });
    setTimeout(() => setToast({ show: false, message: '' }), 3000);
  };

  return (
    <div className="space-y-6 relative">
      {toast.show && (
        <div className="fixed bottom-6 right-6 z-[100] animate-in slide-in-from-bottom-4 fade-in duration-300">
          <div className="bg-zinc-900 text-white px-4 py-3 rounded-md shadow-lg flex items-center gap-3 border border-zinc-800">
            <Info className="w-4 h-4 text-zinc-400" />
            <span className="text-sm font-medium">{toast.message}</span>
            <button onClick={() => setToast({ show: false, message: '' })} className="ml-2 text-zinc-400 hover:text-white transition-colors">
              <X className="w-4 h-4" />
            </button>
          </div>
        </div>
      )}

      <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <div>
          <h1 className="text-2xl font-semibold text-zinc-900 tracking-tight">Katalog Buku</h1>
          <p className="mt-1 text-sm text-zinc-500">Kelola daftar buku yang tersedia di perpustakaan</p>
        </div>
        <button 
          onClick={() => handleAction('Tambah Buku')}
          className="inline-flex items-center justify-center px-4 py-2 bg-zinc-900 text-white text-sm font-medium rounded-md hover:bg-zinc-800 transition-colors"
        >
          <Plus className="w-4 h-4 mr-2" />
          Tambah Buku
        </button>
      </div>

      <div className="bg-white rounded-lg border border-zinc-200 shadow-sm overflow-hidden">
        <div className="p-4 border-b border-zinc-200 flex items-center justify-between gap-4">
          <div className="relative flex-1 max-w-md">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <Search className="h-4 w-4 text-zinc-400" />
            </div>
            <input
              type="text"
              placeholder="Cari judul, penulis, atau penerbit..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="block w-full pl-9 pr-3 py-2 border border-zinc-300 rounded-md text-sm placeholder-zinc-400 focus:outline-none focus:ring-1 focus:ring-zinc-900 focus:border-zinc-900"
            />
          </div>
        </div>

        <div className="overflow-x-auto">
          <table className="w-full text-sm text-left">
            <thead className="bg-zinc-50 text-zinc-500 font-medium border-b border-zinc-200">
              <tr>
                <th className="px-6 py-3 w-1/3">Buku & Deskripsi</th>
                <th className="px-6 py-3">Penulis</th>
                <th className="px-6 py-3">Penerbit</th>
                <th className="px-6 py-3">ISBN</th>
                <th className="px-6 py-3 text-center">Tahun</th>
                <th className="px-6 py-3 text-center">Stok</th>
                <th className="px-6 py-3 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-zinc-200">
              {isLoading ? (
                <tr>
                  <td colSpan={7} className="px-6 py-8 text-center text-zinc-500">
                    Memuat sinkronisasi data relasional...
                  </td>
                </tr>
              ) : filteredBooks.length === 0 ? (
                <tr>
                  <td colSpan={7} className="px-6 py-12 text-center">
                    <BookText className="mx-auto h-8 w-8 text-zinc-300 mb-3" />
                    <p className="text-zinc-500 font-medium">
                      {searchTerm ? 'Buku yang dicari tidak ditemukan' : 'Belum ada data buku'}
                    </p>
                  </td>
                </tr>
              ) : (
                filteredBooks.map((buku, index) => (
                  <tr key={index} className="hover:bg-zinc-50/50 transition-colors">
                    <td className="px-6 py-4">
                      <p className="font-medium text-zinc-900">{buku.judul_buku}</p>
                      <p className="text-xs text-zinc-500 mt-1 line-clamp-2 pr-4 leading-relaxed">
                        {buku.deskripsi_buku || '-'}
                      </p>
                    </td>
                    <td className="px-6 py-4 text-zinc-700">
                      {buku.authorName}
                    </td>
                    <td className="px-6 py-4 text-zinc-700">
                      {buku.publisherName}
                    </td>
                    <td className="px-6 py-4 text-zinc-500 font-mono text-xs">
                      {buku.isbn}
                    </td>
                    <td className="px-6 py-4 text-zinc-500 text-center">
                      {buku.tahun_terbit}
                    </td>
                    <td className="px-6 py-4 text-zinc-900 font-medium text-center">
                      {buku.stok_buku}
                    </td>
                    <td className="px-6 py-4 text-right space-x-2 whitespace-nowrap">
                      <button 
                        onClick={() => handleAction('Edit Buku')}
                        className="p-1.5 text-zinc-500 hover:text-zinc-900 hover:bg-zinc-100 rounded-md transition-colors"
                      >
                        <Edit2 className="w-4 h-4" />
                      </button>
                      <button 
                        onClick={() => handleAction('Hapus Buku')}
                        className="p-1.5 text-red-500 hover:text-red-700 hover:bg-red-50 rounded-md transition-colors"
                      >
                        <Trash2 className="w-4 h-4" />
                      </button>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}