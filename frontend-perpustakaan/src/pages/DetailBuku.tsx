import { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { api } from '../lib/axios';
import { ArrowLeft, BookOpen, Layers, User, Building, Calendar, Hash, CheckCircle2, XCircle } from 'lucide-react';

/* Interfaces */
interface BukuDetail {
  id_buku: string;
  isbn: string;
  judul_buku: string;
  deskripsi_buku: string;
  gambar_buku: string | null;
  kondisi_buku: string;
  tahun_terbit: string;
  stok_buku: number;
  rak_buku: string;
  penulis: { penulis_buku: string; deskripsi: string };
  penerbit: { penerbit_buku: string };
  kategori: { jenis_buku: string };
}

export default function DetailBuku() {
  /* State Management */
  const { id } = useParams<{ id: string }>();
  const [buku, setBuku] = useState<BukuDetail | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  /* Data Fetching */
  useEffect(() => {
    const fetchDetailBuku = async () => {
      try {
        setIsLoading(true);
        setError(null);
        const response = await api.get(`/buku/${id}`);
        
        if (response.data && response.data.data) {
          setBuku(response.data.data);
        } else {
          setError('Data buku tidak ditemukan.');
        }
      } catch (err) {
        setError('Gagal mengambil detail buku. Pastikan server aktif.');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    if (id) {
      fetchDetailBuku();
    }
  }, [id]);

  /* Render Status Handling */
  if (isLoading) {
    return (
      <div className="min-h-screen bg-zinc-50 flex items-center justify-center">
        <div className="text-zinc-500 font-medium animate-pulse">Memuat informasi buku...</div>
      </div>
    );
  }

  if (error || !buku) {
    return (
      <div className="min-h-screen bg-zinc-50 flex flex-col items-center justify-center p-4">
        <div className="text-red-500 mb-4">{error}</div>
        <Link to="/" className="text-blue-600 hover:underline flex items-center gap-2">
          <ArrowLeft className="w-4 h-4" /> Kembali ke Katalog
        </Link>
      </div>
    );
  }

  /* Render Detail Content */
  const isTersedia = buku.stok_buku > 0;

  return (
    <div className="min-h-screen bg-zinc-50 py-12 px-4 sm:px-6 lg:px-8 font-sans">
      <div className="max-w-5xl mx-auto">
        <Link 
          to="/" 
          className="inline-flex items-center gap-2 text-sm font-medium text-zinc-500 hover:text-zinc-900 transition-colors mb-8"
        >
          <ArrowLeft className="w-4 h-4" />
          Kembali ke Katalog
        </Link>

        <div className="bg-white rounded-2xl shadow-sm border border-zinc-200 overflow-hidden">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8 p-8 md:p-12">
            
            {/* Kolom Gambar Cover */}
            <div className="col-span-1">
              <div className="aspect-[3/4] bg-zinc-100 rounded-xl flex items-center justify-center border border-zinc-200 overflow-hidden shadow-inner">
                {buku.gambar_buku ? (
                  <img src={buku.gambar_buku} alt={buku.judul_buku} className="w-full h-full object-cover" />
                ) : (
                  <BookOpen className="w-24 h-24 text-zinc-300" />
                )}
              </div>
            </div>

            {/* Kolom Detail Informasi */}
            <div className="col-span-1 md:col-span-2 flex flex-col">
              <div className="mb-6">
                <div className="flex items-center gap-3 mb-3">
                  <span className="px-3 py-1 bg-zinc-100 text-zinc-700 rounded-full text-xs font-semibold uppercase tracking-wider">
                    {buku.kategori?.jenis_buku || 'Umum'}
                  </span>
                  <span className={`flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-semibold ${isTersedia ? 'bg-emerald-100 text-emerald-700' : 'bg-red-100 text-red-700'}`}>
                    {isTersedia ? <CheckCircle2 className="w-3.5 h-3.5" /> : <XCircle className="w-3.5 h-3.5" />}
                    {isTersedia ? `${buku.stok_buku} Tersedia` : 'Stok Habis'}
                  </span>
                </div>
                <h1 className="text-3xl sm:text-4xl font-bold text-zinc-900 leading-tight mb-2">
                  {buku.judul_buku}
                </h1>
                <p className="text-lg text-zinc-500">
                  Karya <span className="font-medium text-zinc-700">{buku.penulis?.penulis_buku || 'Tidak diketahui'}</span>
                </p>
              </div>

              <div className="prose prose-zinc max-w-none text-zinc-600 mb-8">
                <p className="whitespace-pre-line">{buku.deskripsi_buku || 'Tidak ada deskripsi yang tersedia untuk buku ini.'}</p>
              </div>

              <div className="grid grid-cols-2 gap-6 mt-auto pt-8 border-t border-zinc-100">
                <div className="space-y-4">
                  <div className="flex items-start gap-3">
                    <Hash className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-xs font-medium text-zinc-500 uppercase tracking-wider">ISBN</p>
                      <p className="text-sm font-medium text-zinc-900">{buku.isbn || '-'}</p>
                    </div>
                  </div>
                  <div className="flex items-start gap-3">
                    <Building className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-xs font-medium text-zinc-500 uppercase tracking-wider">Penerbit</p>
                      <p className="text-sm font-medium text-zinc-900">{buku.penerbit?.penerbit_buku || '-'}</p>
                    </div>
                  </div>
                </div>

                <div className="space-y-4">
                  <div className="flex items-start gap-3">
                    <Calendar className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-xs font-medium text-zinc-500 uppercase tracking-wider">Tahun Terbit</p>
                      <p className="text-sm font-medium text-zinc-900">{buku.tahun_terbit || '-'}</p>
                    </div>
                  </div>
                  <div className="flex items-start gap-3">
                    <Layers className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-xs font-medium text-zinc-500 uppercase tracking-wider">Lokasi Rak</p>
                      <p className="text-sm font-medium text-zinc-900">{buku.rak_buku || '-'}</p>
                    </div>
                  </div>
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  );
}