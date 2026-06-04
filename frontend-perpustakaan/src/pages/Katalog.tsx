import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { api } from '../lib/axios';
import { Search, BookOpen, LogIn, Library } from 'lucide-react';

/* Interfaces */
interface Buku {
  id_buku: string;
  judul_buku: string;
  gambar_buku: string | null;
  penulis: { penulis_buku: string };
  penerbit: { penerbit_buku: string };
  kategori: { jenis_buku: string };
  tahun_terbit: string;
  stok_buku: number;
}

export default function Katalog() {
  /* State Management */
  const [buku, setBuku] = useState<Buku[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');

  /* Data Fetching */
  useEffect(() => {
    const fetchBuku = async () => {
      try {
        setIsLoading(true);
        const response = await api.get('/buku');
        if (response.data && response.data.data) {
          setBuku(response.data.data);
        }
      } catch (error) {
        console.error('Gagal mengambil data katalog buku', error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchBuku();
  }, []);

  /* Filter Data */
  const filteredBuku = buku.filter((item) =>
    item.judul_buku.toLowerCase().includes(searchQuery.toLowerCase()) ||
    item.penulis.penulis_buku.toLowerCase().includes(searchQuery.toLowerCase()) ||
    item.kategori.jenis_buku.toLowerCase().includes(searchQuery.toLowerCase())
  );

  /* Utility Functions */
  const getStokBadge = (stok: number) => {
    if (stok > 3) return { label: 'Tersedia', style: 'bg-emerald-100 text-emerald-700' };
    if (stok > 0) return { label: `Tersisa ${stok}`, style: 'bg-amber-100 text-amber-700' };
    return { label: 'Habis', style: 'bg-red-100 text-red-700' };
  };

  /* Render */
  return (
    <div className="min-h-screen bg-zinc-50 font-sans">
      <nav className="bg-white border-b border-zinc-200 sticky top-0 z-10">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 bg-zinc-900 rounded-lg flex items-center justify-center text-white">
              <Library className="w-6 h-6" />
            </div>
            <span className="font-bold text-xl text-zinc-900 tracking-tight">Katalog Perpustakaan</span>
          </div>
          <Link 
            to="/login" 
            className="flex items-center gap-2 text-sm font-medium text-zinc-600 hover:text-zinc-900 transition-colors"
          >
            <LogIn className="w-4 h-4" />
            Area Staf
          </Link>
        </div>
      </nav>

      <div className="bg-zinc-900 text-zinc-50 py-16 px-4">
        <div className="max-w-3xl mx-auto text-center space-y-6">
          <h1 className="text-4xl font-extrabold tracking-tight sm:text-5xl">
            Temukan Buku Pilihanmu
          </h1>
          <p className="text-lg text-zinc-400">
            Jelajahi koleksi literatur, jurnal, dan buku pengetahuan terbaru yang tersedia di perpustakaan kami.
          </p>
          <div className="relative max-w-xl mx-auto mt-8">
            <Search className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-zinc-400" />
            <input
              type="text"
              placeholder="Cari berdasarkan judul, penulis, atau jenis buku..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="w-full pl-12 pr-4 py-4 rounded-full bg-white/10 border border-white/20 text-white placeholder-zinc-400 focus:outline-none focus:bg-white focus:text-zinc-900 focus:placeholder-zinc-500 transition-all shadow-sm"
            />
          </div>
        </div>
      </div>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        {isLoading ? (
          <div className="text-center py-20 text-zinc-500">Memuat katalog...</div>
        ) : filteredBuku.length === 0 ? (
          <div className="text-center py-20 text-zinc-500">Tidak ada buku yang cocok dengan pencarian Anda.</div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
            {filteredBuku.map((item) => {
              const badge = getStokBadge(item.stok_buku);
              return (
                <Link 
                  to={`/buku/${item.id_buku}`}
                  key={item.id_buku} 
                  className="bg-white rounded-xl border border-zinc-200 overflow-hidden shadow-sm hover:shadow-md transition-shadow group flex flex-col"
                >
                  <div className="aspect-[3/4] bg-zinc-100 flex items-center justify-center relative overflow-hidden border-b border-zinc-100">
                    {item.gambar_buku ? (
                      <img src={item.gambar_buku} alt={item.judul_buku} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300" />
                    ) : (
                      <BookOpen className="w-16 h-16 text-zinc-300" />
                    )}
                    <div className="absolute top-3 right-3">
                      <span className={`px-2.5 py-1 rounded-full text-xs font-semibold shadow-sm ${badge.style}`}>
                        {badge.label}
                      </span>
                    </div>
                  </div>
                  
                  <div className="p-5 flex flex-col flex-1">
                    <div className="text-xs font-medium text-blue-600 mb-2 uppercase tracking-wider">
                      {item.kategori?.jenis_buku || 'Umum'}
                    </div>
                    <h3 className="font-bold text-zinc-900 text-lg leading-tight mb-1 line-clamp-2">
                      {item.judul_buku}
                    </h3>
                    <p className="text-sm text-zinc-500 mb-4 line-clamp-1">
                      Oleh: {item.penulis?.penulis_buku || '-'}
                    </p>
                    
                    <div className="mt-auto pt-4 border-t border-zinc-100 flex items-center justify-between text-xs text-zinc-500">
                      <span>Penerbit: {item.penerbit?.penerbit_buku || '-'}</span>
                      <span className="font-medium">{item.tahun_terbit}</span>
                    </div>
                  </div>
                </Link>
              );
            })}
          </div>
        )}
      </main>
    </div>
  );
}