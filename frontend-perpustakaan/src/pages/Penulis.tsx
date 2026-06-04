import { useState, useEffect } from 'react';
import { api } from '../lib/axios';
import { Plus, Search, Edit, Trash2, X } from 'lucide-react';
import Toast from '../components/Toast';

interface Penulis {
  id: string;
  penulis_buku: string;
  alamat_penulis: string;
  email_penulis: string;
  deskripsi: string;
}

interface FormData {
  id?: string;
  penulis_buku: string;
  alamat_penulis: string;
  email_penulis: string;
  deskripsi: string;
}

export default function Penulis() {
  const [data, setData] = useState<Penulis[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [modalMode, setModalMode] = useState<'add' | 'edit'>('add');
  const [formData, setFormData] = useState<FormData>({ penulis_buku: '', alamat_penulis: '', email_penulis: '', deskripsi: '' });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [notification, setNotification] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  const showNotification = (message: string, type: 'success' | 'error') => setNotification({ message, type });

  const fetchData = async () => {
    try {
      setIsLoading(true);
      const res = await api.get('/admin/buku/author');
      if (res.data?.data) setData(res.data.data);
    } catch {
      showNotification('Gagal memuat data penulis.', 'error');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => { fetchData(); }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    const payload = {
        id: formData.id,
        penulis_buku: formData.penulis_buku,
        alamat_penulis: formData.alamat_penulis,
        email_penulis: formData.email_penulis,
        deskripsi: formData.deskripsi
    };

    try {
      if (modalMode === 'add') {
        await api.post('/admin/buku/author/create', payload);
        showNotification('Penulis berhasil ditambahkan.', 'success');
      } else {
        await api.put('/admin/buku/author/update', payload);
        showNotification('Penulis berhasil diperbarui.', 'success');
      }
      setIsModalOpen(false);
      fetchData();
    } catch (err: any) {
      const msg = err.response?.data?.msg || 'Gagal menyimpan data.';
      showNotification(msg, 'error');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = async (id: string, nama: string) => {
    if (window.confirm(`Hapus penulis "${nama}"?`)) {
      try {
        await api.delete('/admin/buku/author/delete', { data: { id } });
        showNotification('Penulis berhasil dihapus.', 'success');
        fetchData();
      } catch (err: any) {
        // Menangkap error 502/500 dari backend dengan pesan yang jelas
        if (err.response?.status === 502 || err.response?.status === 500) {
           showNotification('Data tidak dapat dihapus karena masih digunakan buku lain.', 'error');
        } else {
           showNotification(err.response?.data?.msg || 'Gagal menghapus data.', 'error');
        }
      }
    }
  };

  return (
    <div className="space-y-6">
      {notification && <Toast message={notification.message} type={notification.type} onClose={() => setNotification(null)} />}
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Manajemen Penulis</h1>
        <button onClick={() => {setModalMode('add'); setFormData({penulis_buku:'', alamat_penulis:'', email_penulis:'', deskripsi:''}); setIsModalOpen(true)}} className="bg-zinc-900 text-white px-4 py-2 rounded-md flex items-center gap-2"><Plus className="w-4 h-4"/> Tambah</button>
      </div>
      
      <div className="bg-white rounded-lg border border-zinc-200 overflow-hidden">
        <div className="p-4 border-b">
           <input type="text" placeholder="Cari penulis..." onChange={(e) => setSearchQuery(e.target.value)} className="w-full max-w-md p-2 border rounded text-sm"/>
        </div>
        <table className="w-full text-left text-sm">
          <thead className="bg-zinc-50 border-b">
            <tr><th className="px-6 py-3">Nama</th><th className="px-6 py-3">Email</th><th className="px-6 py-3 text-right">Aksi</th></tr>
          </thead>
          <tbody className="divide-y">
            {data.filter(i => i.penulis_buku.toLowerCase().includes(searchQuery.toLowerCase())).map((item) => (
              <tr key={item.id} className="hover:bg-zinc-50">
                <td className="px-6 py-4 font-medium">{item.penulis_buku}</td>
                <td className="px-6 py-4">{item.email_penulis}</td>
                <td className="px-6 py-4 text-right flex justify-end gap-2">
                  <button onClick={() => {setModalMode('edit'); setFormData(item); setIsModalOpen(true)}} className="text-zinc-400 hover:text-blue-600"><Edit className="w-4 h-4"/></button>
                  <button onClick={() => handleDelete(item.id, item.penulis_buku)} className="text-zinc-400 hover:text-red-600"><Trash2 className="w-4 h-4"/></button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {isModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-zinc-900/50 backdrop-blur-sm">
          <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg w-full max-w-md space-y-4">
            <div className="flex justify-between items-center"><h3 className="font-bold">Form Penulis</h3><X className="cursor-pointer" onClick={() => setIsModalOpen(false)}/></div>
            <input required type="text" placeholder="Nama Penulis" value={formData.penulis_buku} onChange={e => setFormData({...formData, penulis_buku: e.target.value})} className="w-full p-2 border rounded"/>
            <input required type="email" placeholder="Email" value={formData.email_penulis} onChange={e => setFormData({...formData, email_penulis: e.target.value})} className="w-full p-2 border rounded"/>
            <input required type="text" placeholder="Alamat" value={formData.alamat_penulis} onChange={e => setFormData({...formData, alamat_penulis: e.target.value})} className="w-full p-2 border rounded"/>
            <textarea required placeholder="Deskripsi" value={formData.deskripsi} onChange={e => setFormData({...formData, deskripsi: e.target.value})} className="w-full p-2 border rounded" rows={3}/>
            <button type="submit" disabled={isSubmitting} className="w-full bg-zinc-900 text-white p-2 rounded">{isSubmitting ? '...' : 'Simpan'}</button>
          </form>
        </div>
      )}
    </div>
  );
}