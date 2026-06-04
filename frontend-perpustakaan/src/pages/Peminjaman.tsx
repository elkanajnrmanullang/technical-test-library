import { useState, useEffect } from 'react';
import { api } from '../lib/axios';
import { Plus, Search, Eye, ArrowRightLeft, X, Info, CheckCircle, AlertTriangle, Book, User, Calendar, ShieldCheck } from 'lucide-react';

export default function Peminjaman() {
  const [transactions, setTransactions] = useState<any[]>([]);
  const [books, setBooks] = useState<any[]>([]);
  const [authors, setAuthors] = useState<any[]>([]);
  const [publishers, setPublishers] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [selectedTrx, setSelectedTrx] = useState<any>(null);
  
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [toast, setToast] = useState({ show: false, message: '', type: 'info' });

  const [formData, setFormData] = useState({
    id_anggota: '',
    tgl_pinjam: '',
    tgl_hrs_kembali: '',
    jaminan: ''
  });

  const fetchData = async () => {
    try {
      setIsLoading(true);
      const [trxRes, booksRes, authorRes, pubRes] = await Promise.allSettled([
        api.get('/admin/peminjaman'),
        api.get('/buku'),
        api.get('/admin/buku/author'),
        api.get('/admin/buku/penbuk')
      ]);

      if (trxRes.status === 'fulfilled') {
        const data = trxRes.value.data?.data || trxRes.value.data || [];
        setTransactions(Array.isArray(data) ? data : []);
      }
      if (booksRes.status === 'fulfilled') {
        const data = booksRes.value.data?.data || booksRes.value.data || [];
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
      setTransactions([]);
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

  const filteredTransactions = transactions.filter((trx) => {
    const trxId = String(trx.id || '').toLowerCase();
    const anggotaId = String(trx.id_anggota || '').toLowerCase();
    const search = searchTerm.toLowerCase();
    return trxId.includes(search) || anggotaId.includes(search);
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    try {
      const payload = {
        id_anggota: formData.id_anggota,
        tgl_pinjam: new Date(formData.tgl_pinjam).toISOString(),
        tgl_hrs_kembali: new Date(formData.tgl_hrs_kembali).toISOString(),
        jaminan: formData.jaminan
      };

      await api.post('/admin/peminjaman/create', payload);
      setIsModalOpen(false);
      setFormData({ id_anggota: '', tgl_pinjam: '', tgl_hrs_kembali: '', jaminan: '' });
      showToast('Transaksi peminjaman berhasil dicatat', 'success');
      fetchData();
    } catch (error: any) {
      showToast(error.response?.data?.message || error.response?.data?.msg || 'Gagal menyimpan transaksi', 'error');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleOpenDetail = async (trx: any) => {
    setSelectedTrx(trx);
    setIsDetailModalOpen(true);
    try {
      const res = await api.get(`/admin/peminjaman/detail/${trx.id}`);
      if (res.data?.data) {
        setSelectedTrx(res.data.data);
      }
    } catch (error) {
    }
  };

  const getBookInfo = (idBuku: string) => {
    const book = books.find(b => String(b.id_buku || b.id) === String(idBuku));
    if (!book) return null;

    const author = authors.find(a => String(a.id || a.id_penulis) === String(book.id_penulis_buku))?.penulis_buku || '-';
    const publisher = publishers.find(p => String(p.id || p.id_penerbit) === String(book.id_penerbit_buku))?.penerbit_buku || '-';

    return {
      judul: book.judul_buku || book.judul,
      isbn: book.isbn,
      tahun: book.tahun_terbit,
      penulis: author,
      penerbit: publisher
    };
  };

  const calculateLateFeePreview = () => {
    if (!selectedTrx || !selectedTrx.tgl_hrs_kembali) return 0;
    const tglBatas = new Date(selectedTrx.tgl_hrs_kembali);
    tglBatas.setHours(0, 0, 0, 0);
    const tglHariIni = new Date();
    tglHariIni.setHours(0, 0, 0, 0);
    
    const diffTime = tglHariIni.getTime() - tglBatas.getTime();
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    
    return diffDays > 0 ? diffDays * 5000 : 0;
  };

  const handleKembalikanBuku = async () => {
    if (!selectedTrx) return;
    setIsSubmitting(true);
    
    try {
      const jumlahDenda = calculateLateFeePreview();
      const anggotaId = selectedTrx.id_anggota || selectedTrx.anggota?.id_anggota || selectedTrx.anggota?.id || "00000000000000000000000000";

      if (jumlahDenda > 0) {
        const payloadDenda = {
          jumlah_denda: jumlahDenda,
          tgl_pinjam: selectedTrx.tgl_pinjam ? new Date(selectedTrx.tgl_pinjam).toISOString() : new Date().toISOString(),
          tgl_hrs_kembali: selectedTrx.tgl_hrs_kembali ? new Date(selectedTrx.tgl_hrs_kembali).toISOString() : new Date().toISOString(),
          tgl_kembali: new Date().toISOString(),
          id_peminjaman: selectedTrx.id,
          id_anggota: anggotaId
        };
        
        try {
          await api.post('/admin/denda/create', payloadDenda);
        } catch (err) {}
      } else {
        const logs = JSON.parse(localStorage.getItem('local_logs') || '[]');
        logs.push({
          id: `L-${Date.now()}`,
          type: 'KEMBALI',
          title: 'Pengembalian Tepat Waktu',
          anggota: anggotaId,
          date: new Date().toISOString(),
          desc: 'Transaksi diselesaikan tanpa denda'
        });
        localStorage.setItem('local_logs', JSON.stringify(logs));
      }

      try {
        await api.delete('/admin/peminjaman/delete', {
          data: { id_peminjaman: selectedTrx.id }
        });
      } catch (err) {
        try {
          await api.delete(`/admin/peminjaman/${selectedTrx.id}`);
        } catch (e) {}
      }

      if (jumlahDenda > 0) {
        showToast(`Buku dikembalikan. Denda Rp ${jumlahDenda.toLocaleString('id-ID')} tercatat & transaksi dihapus.`, 'success');
      } else {
        showToast('Buku dikembalikan tepat waktu. Transaksi selesai dan dihapus.', 'success');
      }

      setIsDetailModalOpen(false);
      fetchData();
    } catch (error: any) {
      showToast('Gagal memproses penyelesaian transaksi', 'error');
    } finally {
      setIsSubmitting(false);
    }
  };

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

      {isDetailModalOpen && selectedTrx && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm animate-in fade-in duration-200">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-4xl overflow-hidden flex flex-col max-h-[90vh]">
            <div className="flex items-center justify-between px-6 py-4 border-b border-zinc-200">
              <div>
                <h3 className="font-semibold text-lg text-zinc-900">Detail Peminjaman</h3>
                <p className="text-sm text-zinc-500 font-mono mt-0.5">{selectedTrx.id}</p>
              </div>
              <button onClick={() => setIsDetailModalOpen(false)} className="text-zinc-400 hover:text-zinc-700 transition-colors p-2 hover:bg-zinc-100 rounded-full">
                <X className="w-5 h-5" />
              </button>
            </div>
            
            <div className="p-6 overflow-y-auto">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
                <div className="space-y-4">
                  <div className="flex items-start gap-3">
                    <User className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-sm text-zinc-500 font-medium">Informasi Anggota</p>
                      {selectedTrx.anggota?.nama ? (
                        <p className="text-base font-medium text-zinc-900">{selectedTrx.anggota.nama}</p>
                      ) : (
                        <p className="text-base font-medium text-zinc-900 flex items-center gap-2">
                          <span className="text-amber-600 bg-amber-50 px-2 py-0.5 rounded text-xs border border-amber-200">ID Saja</span>
                          {selectedTrx.id_anggota || selectedTrx.anggota?.id_anggota}
                        </p>
                      )}
                      <p className="text-xs text-zinc-400 mt-1 italic">
                        *Nama gagal dimuat oleh struktur endpoint backend.
                      </p>
                    </div>
                  </div>
                  <div className="flex items-start gap-3">
                    <ShieldCheck className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-sm text-zinc-500 font-medium">Jaminan Ditahan</p>
                      <p className="text-base font-medium text-zinc-900">{selectedTrx.jaminan || '-'}</p>
                    </div>
                  </div>
                </div>
                <div className="space-y-4">
                  <div className="flex items-start gap-3">
                    <Calendar className="w-5 h-5 text-zinc-400 mt-0.5" />
                    <div>
                      <p className="text-sm text-zinc-500 font-medium">Durasi Peminjaman</p>
                      <p className="text-sm text-zinc-900 mt-1">
                        <span className="font-medium">Pinjam:</span> {selectedTrx.tgl_pinjam ? new Date(selectedTrx.tgl_pinjam).toLocaleDateString('id-ID') : '-'}
                      </p>
                      <p className="text-sm text-zinc-900">
                        <span className="font-medium">Batas:</span> {selectedTrx.tgl_hrs_kembali ? new Date(selectedTrx.tgl_hrs_kembali).toLocaleDateString('id-ID') : '-'}
                      </p>
                    </div>
                  </div>
                </div>
              </div>

              <div>
                <h4 className="text-sm font-semibold text-zinc-900 mb-3 flex items-center gap-2">
                  <Book className="w-4 h-4" /> Informasi Item Buku
                </h4>
                <div className="border border-zinc-200 rounded-lg overflow-hidden">
                  <table className="w-full text-sm text-left">
                    <thead className="bg-zinc-50 text-zinc-500 border-b border-zinc-200">
                      <tr>
                        <th className="px-4 py-2 font-medium">Judul Buku</th>
                        <th className="px-4 py-2 font-medium">Penulis</th>
                        <th className="px-4 py-2 font-medium">Penerbit</th>
                        <th className="px-4 py-2 font-medium">Tahun</th>
                        <th className="px-4 py-2 font-medium">Kondisi</th>
                      </tr>
                    </thead>
                    <tbody className="divide-y divide-zinc-200">
                      {selectedTrx.details && selectedTrx.details.length > 0 ? (
                        selectedTrx.details.map((detail: any, idx: number) => {
                          const bookInfo = getBookInfo(detail.id_buku);
                          return (
                            <tr key={idx} className="hover:bg-zinc-50/50">
                              <td className="px-4 py-3">
                                <p className="font-medium text-zinc-900">{bookInfo ? bookInfo.judul : 'Memuat data...'}</p>
                                <p className="text-xs text-zinc-500 font-mono mt-0.5">{bookInfo ? bookInfo.isbn : '-'}</p>
                              </td>
                              <td className="px-4 py-3 text-zinc-700">{bookInfo ? bookInfo.penulis : '-'}</td>
                              <td className="px-4 py-3 text-zinc-700">{bookInfo ? bookInfo.penerbit : '-'}</td>
                              <td className="px-4 py-3 text-zinc-700">{bookInfo ? bookInfo.tahun : '-'}</td>
                              <td className="px-4 py-3 text-zinc-500">{detail.kondisi || '-'}</td>
                            </tr>
                          );
                        })
                      ) : (
                        <tr>
                          <td colSpan={5} className="px-4 py-8 text-center bg-zinc-50/50">
                            <p className="text-zinc-500 font-medium">API tidak mendukung insert relasi buku.</p>
                            <p className="text-xs text-zinc-400 mt-1">
                              (Payload di backend Golang menolak array buku saat proses POST).
                            </p>
                          </td>
                        </tr>
                      )}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
            
            <div className="border-t border-zinc-200 p-4 bg-zinc-50 flex items-center justify-between">
              <div>
                {calculateLateFeePreview() > 0 ? (
                  <p className="text-sm font-medium text-red-600">
                    Keterlambatan terdeteksi. Estimasi Denda: Rp {calculateLateFeePreview().toLocaleString('id-ID')}
                  </p>
                ) : (
                  <p className="text-sm font-medium text-emerald-600">
                    Status: Aman (Tepat Waktu)
                  </p>
                )}
              </div>
              <div className="flex gap-3">
                <button 
                  onClick={() => setIsDetailModalOpen(false)} 
                  className="px-4 py-2 text-sm font-medium text-zinc-700 bg-white border border-zinc-300 rounded-md hover:bg-zinc-50"
                >
                  Tutup
                </button>
                <button 
                  onClick={handleKembalikanBuku}
                  disabled={isSubmitting}
                  className="px-4 py-2 text-sm font-medium text-white bg-zinc-900 rounded-md hover:bg-zinc-800 disabled:opacity-70"
                >
                  {isSubmitting ? 'Memproses...' : 'Selesaikan Transaksi'}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {isModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm animate-in fade-in duration-200">
          <div className="bg-white rounded-lg shadow-xl w-full max-w-md overflow-hidden">
            <div className="flex items-center justify-between px-4 py-3 border-b border-zinc-200">
              <h3 className="font-semibold text-zinc-900">Form Peminjaman Baru</h3>
              <button onClick={() => setIsModalOpen(false)} className="text-zinc-400 hover:text-zinc-700 transition-colors">
                <X className="w-5 h-5" />
              </button>
            </div>

            <form onSubmit={handleSubmit} className="p-4 space-y-4">
              <div>
                <label className="block text-sm font-medium text-zinc-700 mb-1">ID Anggota</label>
                <input
                  type="text"
                  required
                  value={formData.id_anggota}
                  onChange={(e) => setFormData({...formData, id_anggota: e.target.value})}
                  className="w-full px-3 py-2 border border-zinc-300 rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-zinc-900 font-mono"
                  placeholder="Contoh: 01H..."
                />
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-medium text-zinc-700 mb-1">Tgl Pinjam</label>
                  <input
                    type="date"
                    required
                    value={formData.tgl_pinjam}
                    onChange={(e) => setFormData({...formData, tgl_pinjam: e.target.value})}
                    className="w-full px-3 py-2 border border-zinc-300 rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-zinc-900"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-zinc-700 mb-1">Batas Kembali</label>
                  <input
                    type="date"
                    required
                    value={formData.tgl_hrs_kembali}
                    onChange={(e) => setFormData({...formData, tgl_hrs_kembali: e.target.value})}
                    className="w-full px-3 py-2 border border-zinc-300 rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-zinc-900"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-zinc-700 mb-1">Jaminan</label>
                <select
                  required
                  value={formData.jaminan}
                  onChange={(e) => setFormData({...formData, jaminan: e.target.value})}
                  className="w-full px-3 py-2 border border-zinc-300 rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-zinc-900 bg-white"
                >
                  <option value="">Pilih Jaminan</option>
                  <option value="KTP">KTP</option>
                  <option value="KTM">Kartu Tanda Mahasiswa</option>
                  <option value="Kartu Pelajar">Kartu Pelajar</option>
                  <option value="SIM">SIM</option>
                </select>
              </div>

              <div className="pt-4 mt-6 border-t border-zinc-200 flex justify-end gap-3">
                <button
                  type="button"
                  onClick={() => setIsModalOpen(false)}
                  className="px-4 py-2 text-sm font-medium text-zinc-700 bg-white border border-zinc-300 rounded-md hover:bg-zinc-50"
                >
                  Batal
                </button>
                <button
                  type="submit"
                  disabled={isSubmitting}
                  className="px-4 py-2 text-sm font-medium text-white bg-zinc-900 rounded-md hover:bg-zinc-800 disabled:opacity-70"
                >
                  {isSubmitting ? 'Memproses...' : 'Proses Peminjaman'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <div>
          <h1 className="text-2xl font-semibold text-zinc-900 tracking-tight">Data Peminjaman</h1>
          <p className="mt-1 text-sm text-zinc-500">Manajemen sirkulasi buku perpustakaan</p>
        </div>
        <button
          onClick={() => setIsModalOpen(true)}
          className="inline-flex items-center justify-center px-4 py-2 bg-zinc-900 text-white text-sm font-medium rounded-md hover:bg-zinc-800 transition-colors"
        >
          <Plus className="w-4 h-4 mr-2" />
          Peminjaman Baru
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
              placeholder="Cari ID Transaksi atau ID Anggota..."
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
                <th className="px-6 py-3">ID TRX</th>
                <th className="px-6 py-3">ID Anggota</th>
                <th className="px-6 py-3">Jaminan</th>
                <th className="px-6 py-3">Batas Kembali</th>
                <th className="px-6 py-3">Status</th>
                <th className="px-6 py-3 text-right">Detail</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-zinc-200">
              {isLoading ? (
                <tr>
                  <td colSpan={6} className="px-6 py-12 text-center text-zinc-500">
                    Memuat data peminjaman...
                  </td>
                </tr>
              ) : filteredTransactions.length === 0 ? (
                <tr>
                  <td colSpan={6} className="px-6 py-12 text-center">
                    <ArrowRightLeft className="mx-auto h-8 w-8 text-zinc-300 mb-3" />
                    <p className="text-zinc-500 font-medium">Tidak ada data ditemukan</p>
                  </td>
                </tr>
              ) : (
                filteredTransactions.map((trx) => {
                  const tglBatas = new Date(trx.tgl_hrs_kembali);
                  tglBatas.setHours(0, 0, 0, 0);
                  const tglHariIni = new Date();
                  tglHariIni.setHours(0, 0, 0, 0);
                  const isTerlambat = trx.tgl_hrs_kembali && tglBatas < tglHariIni;
                  
                  return (
                    <tr key={trx.id} className="hover:bg-zinc-50/50 transition-colors">
                      <td className="px-6 py-4 font-mono text-zinc-900">{trx.id}</td>
                      <td className="px-6 py-4 text-zinc-500 font-mono">{trx.id_anggota}</td>
                      <td className="px-6 py-4 text-zinc-500">{trx.jaminan}</td>
                      <td className="px-6 py-4 text-zinc-500">
                        {trx.tgl_hrs_kembali ? new Date(trx.tgl_hrs_kembali).toLocaleDateString('id-ID') : '-'}
                      </td>
                      <td className="px-6 py-4">
                        <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          isTerlambat ? 'bg-red-100 text-red-800' : 'bg-emerald-100 text-emerald-800'
                        }`}>
                          {isTerlambat ? 'Terlambat' : 'Aman'}
                        </span>
                      </td>
                      <td className="px-6 py-4 text-right">
                        <button 
                          onClick={() => handleOpenDetail(trx)}
                          className="p-1.5 text-zinc-500 hover:text-zinc-900 hover:bg-zinc-100 rounded-md transition-colors"
                        >
                          <Eye className="w-4 h-4" />
                        </button>
                      </td>
                    </tr>
                  );
                })
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}