import { useState, useEffect } from 'react';
import { api } from '../lib/axios';
import { Search, Eye, X, Info, CheckCircle, AlertTriangle, AlertCircle } from 'lucide-react';

export default function Denda() {
  const [denda, setDenda] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const [toast, setToast] = useState({ show: false, message: '', type: 'info' });

  const fetchData = async () => {
    try {
      setIsLoading(true);
      const response = await api.get('/admin/denda');
      const data = response.data?.data || response.data || [];
      setDenda(Array.isArray(data) ? data : []);
    } catch (error) {
      setDenda([]);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const showToast = (message: string, type: 'info' | 'success' | 'error' = 'info') => {
    setToast({ show: true, message, type });
    setTimeout(() => setToast({ show: false, message: '', type: 'info' }), 3000);
  };

  const filteredDenda = denda.filter((item) => {
    const idDenda = String(item.id || '').toLowerCase();
    const idPem = String(item.id_peminjaman || '').toLowerCase();
    const idAnggota = String(item.id_anggota || '').toLowerCase();
    const search = searchTerm.toLowerCase();
    return idDenda.includes(search) || idPem.includes(search) || idAnggota.includes(search);
  });

  return (
    <div className="space-y-6 relative">
      {toast.show && (
        <div className={`fixed bottom-6 right-6 z-[100] animate-in slide-in-from-bottom-4 fade-in duration-300`}>
          <div className={`px-4 py-3 rounded-md shadow-lg flex items-center gap-3 border ${
            toast.type === 'success' ? 'bg-emerald-50 border-emerald-200 text-emerald-800' :
            toast.type === 'error' ? 'bg-red-50 border-red-200 text-red-800' :
            'bg-zinc-900 border-zinc-800 text-white'
          }`}>
            {toast.type === 'success' && <CheckCircle className="w-4 h-4" />}
            {toast.type === 'error' && <AlertTriangle className="w-4 h-4" />}
            {toast.type === 'info' && <Info className="w-4 h-4 text-zinc-400" />}
            <span className="text-sm font-medium">{toast.message}</span>
            <button onClick={() => setToast({ show: false, message: '', type: 'info' })} className="ml-2 opacity-70 hover:opacity-100 transition-opacity">
              <X className="w-4 h-4" />
            </button>
          </div>
        </div>
      )}

      <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <div>
          <h1 className="text-2xl font-semibold text-zinc-900 tracking-tight">Data Denda</h1>
          <p className="mt-1 text-sm text-zinc-500">Manajemen log tagihan keterlambatan pengembalian buku</p>
        </div>
      </div>

      <div className="bg-white rounded-lg border border-zinc-200 shadow-sm overflow-hidden">
        <div className="p-4 border-b border-zinc-200 flex items-center justify-between gap-4">
          <div className="relative flex-1 max-w-md">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <Search className="h-4 w-4 text-zinc-400" />
            </div>
            <input
              type="text"
              placeholder="Cari ID Denda, Transaksi, atau Anggota..."
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
                <th className="px-6 py-3">ID Denda</th>
                <th className="px-6 py-3">ID Transaksi</th>
                <th className="px-6 py-3">ID Anggota</th>
                <th className="px-6 py-3">Tgl Dikembalikan</th>
                <th className="px-6 py-3">Nominal Denda</th>
                <th className="px-6 py-3 text-right">Detail</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-zinc-200">
              {isLoading ? (
                <tr>
                  <td colSpan={6} className="px-6 py-12 text-center text-zinc-500">
                    Memuat riwayat denda...
                  </td>
                </tr>
              ) : filteredDenda.length === 0 ? (
                <tr>
                  <td colSpan={6} className="px-6 py-12 text-center">
                    <AlertCircle className="mx-auto h-8 w-8 text-zinc-300 mb-3" />
                    <p className="text-zinc-500 font-medium">Tidak ada histori denda ditemukan</p>
                  </td>
                </tr>
              ) : (
                filteredDenda.map((item) => (
                  <tr key={item.id} className="hover:bg-zinc-50/50 transition-colors">
                    <td className="px-6 py-4 font-mono text-zinc-900">{item.id}</td>
                    <td className="px-6 py-4 text-zinc-500 font-mono">{item.id_peminjaman}</td>
                    <td className="px-6 py-4 text-zinc-500 font-mono">{item.id_anggota}</td>
                    <td className="px-6 py-4 text-zinc-500">
                      {item.tgl_kembali ? new Date(item.tgl_kembali).toLocaleDateString('id-ID') : '-'}
                    </td>
                    <td className="px-6 py-4 font-medium text-red-600">
                      Rp {item.jumlah_denda ? item.jumlah_denda.toLocaleString('id-ID') : '0'}
                    </td>
                    <td className="px-6 py-4 text-right">
                      <button
                        onClick={() => showToast('Bukti denda digital sedang dikembangkan', 'info')}
                        className="p-1.5 text-zinc-500 hover:text-zinc-900 hover:bg-zinc-100 rounded-md transition-colors"
                      >
                        <Eye className="w-4 h-4" />
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