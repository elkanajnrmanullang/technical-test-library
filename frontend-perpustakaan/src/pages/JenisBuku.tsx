import { useState, useEffect } from 'react';
import { api } from '../lib/axios';
import { Plus, Search, Edit, Trash2, X } from 'lucide-react';
import Toast from '../components/Toast';

interface JenisBuku {
  id: string;
  jenis_buku: string;
  deskripsi: string;
}

interface FormData {
  id?: string;
  jenis_buku: string;
  deskripsi: string;
}

export default function JenisBuku() {
  const [jenisBuku, setJenisBuku] = useState<JenisBuku[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [modalMode, setModalMode] = useState<'add' | 'edit'>('add');
  const [formData, setFormData] = useState<FormData>({ jenis_buku: '', deskripsi: '' });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [notification, setNotification] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  const showNotification = (message: string, type: 'success' | 'error') => {
    setNotification({ message, type });
  };

  const fetchJenisBuku = async () => {
    try {
      setIsLoading(true);
      const response = await api.get('/admin/buku/jenbuk');
      if (response.data?.data) {
        setJenisBuku(response.data.data);
      }
    } catch {
      showNotification('Gagal memuat data kategori.', 'error');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchJenisBuku();
  }, []);

  const handleOpenModal = (mode: 'add' | 'edit', data?: JenisBuku) => {
    setModalMode(mode);
    setFormData(data ? { id: data.id, jenis_buku: data.jenis_buku, deskripsi: data.deskripsi } : { jenis_buku: '', deskripsi: '' });
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
    setFormData({ jenis_buku: '', deskripsi: '' });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    
    try {
      if (modalMode === 'add') {
        await api.post('/admin/buku/jenbuk/create', {
          jenis_buku: formData.jenis_buku,
          deskripsi: formData.deskripsi,
        });
        showNotification('Kategori berhasil ditambahkan.', 'success');
      } else {
        await api.put('/admin/buku/jenbuk/update', {
          id: formData.id,
          jenis_buku: formData.jenis_buku,
          deskripsi: formData.deskripsi,
        });
        showNotification('Kategori berhasil diperbarui.', 'success');
      }
      handleCloseModal();
      fetchJenisBuku();
    } catch {
      showNotification('Gagal menyimpan data.', 'error');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = async (id: string, nama: string) => {
    if (window.confirm(`Hapus kategori "${nama}"?`)) {
      try {
        await api.delete('/admin/buku/jenbuk/delete', { data: { id } });
        showNotification('Kategori berhasil dihapus.', 'success');
        fetchJenisBuku();
      } catch (err: any) {
        if (err.response?.status === 500) {
          showNotification('Kategori sedang digunakan, tidak dapat dihapus.', 'error');
        } else {
          showNotification('Gagal menghapus data.', 'error');
        }
      }
    }
  };

  const filteredJenisBuku = jenisBuku.filter((item) =>
    item.jenis_buku.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="space-y-6">
      {notification && (
        <Toast 
          message={notification.message} 
          type={notification.type} 
          onClose={() => setNotification(null)} 
        />
      )}

      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-zinc-900">Jenis Buku</h1>
        <button 
          onClick={() => handleOpenModal('add')}
          className="flex items-center gap-2 bg-zinc-900 text-white px-4 py-2 rounded-md hover:bg-zinc-800 transition-colors"
        >
          <Plus className="w-4 h-4" /> Tambah
        </button>
      </div>

      <div className="bg-white rounded-lg border border-zinc-200 shadow-sm overflow-hidden">
        <div className="p-4 border-b border-zinc-200">
          <input
            type="text"
            placeholder="Cari kategori..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="w-full max-w-md pl-4 py-2 border border-zinc-300 rounded-md focus:outline-none focus:ring-1 focus:ring-zinc-900 text-sm"
          />
        </div>

        <table className="w-full text-left text-sm text-zinc-600">
          <thead className="bg-zinc-50 border-b border-zinc-200">
            <tr>
              <th className="px-6 py-3">Jenis Buku</th>
              <th className="px-6 py-3">Deskripsi</th>
              <th className="px-6 py-3 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-zinc-200">
            {isLoading ? (
              <tr><td colSpan={3} className="px-6 py-8 text-center">Memuat...</td></tr>
            ) : filteredJenisBuku.map((item) => (
              <tr key={item.id} className="hover:bg-zinc-50">
                <td className="px-6 py-4 font-medium text-zinc-900">{item.jenis_buku}</td>
                <td className="px-6 py-4">{item.deskripsi}</td>
                <td className="px-6 py-4 text-right flex justify-end gap-2">
                  <button onClick={() => handleOpenModal('edit', item)} className="p-1.5 text-zinc-400 hover:text-blue-600"><Edit className="w-4 h-4" /></button>
                  <button onClick={() => handleDelete(item.id, item.jenis_buku)} className="p-1.5 text-zinc-400 hover:text-red-600"><Trash2 className="w-4 h-4" /></button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {isModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-zinc-900/50 backdrop-blur-sm">
          <div className="bg-white rounded-lg w-full max-w-md p-6">
            <div className="flex justify-between items-center mb-4">
              <h3 className="font-bold">{modalMode === 'add' ? 'Tambah' : 'Edit'} Kategori</h3>
              <button onClick={handleCloseModal}><X className="w-5 h-5 text-zinc-400" /></button>
            </div>
            <form onSubmit={handleSubmit} className="space-y-4">
              <input type="text" required placeholder="Nama Jenis" value={formData.jenis_buku} onChange={(e) => setFormData({...formData, jenis_buku: e.target.value})} className="w-full p-2 border rounded" />
              <textarea required placeholder="Deskripsi" value={formData.deskripsi} onChange={(e) => setFormData({...formData, deskripsi: e.target.value})} className="w-full p-2 border rounded" rows={3} />
              <button type="submit" disabled={isSubmitting} className="w-full bg-zinc-900 text-white p-2 rounded">{isSubmitting ? '...' : 'Simpan'}</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}