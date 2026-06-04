import { useState, useEffect } from 'react';
import { api } from '../lib/axios';
import { Plus, Search, Edit, Trash2, X } from 'lucide-react';
import Toast from '../components/Toast';

interface Penerbit {
  id: string;
  penerbit_buku: string;
  email_penerbit: string;
  telp_penerbit: string;
  alamat_penerbit: string;
  deskripsi: string;
}

interface FormData {
  id: string;
  penerbit_buku: string;
  email_penerbit: string;
  telp_penerbit: string;
  alamat_penerbit: string;
  deskripsi: string;
}

const initialForm: FormData = {
  id: '',
  penerbit_buku: '',
  email_penerbit: '',
  telp_penerbit: '',
  alamat_penerbit: '',
  deskripsi: ''
};

export default function Penerbit() {
  const [data, setData] = useState<Penerbit[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [modalMode, setModalMode] = useState<'add' | 'edit'>('add');
  const [formData, setFormData] = useState<FormData>(initialForm);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [notification, setNotification] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  const showNotification = (message: string, type: 'success' | 'error') => setNotification({ message, type });

  const fetchData = async () => {
    try {
      setIsLoading(true);
      const res = await api.get('/admin/buku/penbuk');
      if (res.data?.data) setData(res.data.data);
    } catch {
      showNotification('Gagal memuat data penerbit.', 'error');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => { fetchData(); }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    
    // Payload disesuaikan dengan struct backend
    const payload = {
        id: formData.id,
        penerbit_buku: formData.penerbit_buku,
        email_penerbit: formData.email_penerbit,
        telp_penerbit: formData.telp_penerbit,
        alamat_penerbit: formData.alamat_penerbit,
        deskripsi: formData.deskripsi
    };

    try {
      if (modalMode === 'add') {
        await api.post('/admin/buku/penbuk/create', payload);
        showNotification('Penerbit berhasil ditambahkan.', 'success');
      } else {
        await api.put('/admin/buku/penbuk/update', payload);
        showNotification('Penerbit berhasil diperbarui.', 'success');
      }
      setIsModalOpen(false);
      fetchData();
    } catch (err: any) {
      showNotification(err.response?.data?.msg || 'Gagal menyimpan data.', 'error');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = async (id: string, nama: string) => {
    if (window.confirm(`Hapus penerbit "${nama}"?`)) {
      try {
        await api.delete('/admin/buku/penbuk/delete', { data: { id } });
        showNotification('Penerbit berhasil dihapus.', 'success');
        fetchData();
      } catch (err: any) {
        if (err.response?.status === 502 || err.response?.status === 500) {
           showNotification('Data sedang digunakan, tidak bisa dihapus.', 'error');
        } else {
           showNotification('Gagal menghapus data.', 'error');
        }
      }
    }
  };

  return (
    <div className="space-y-6">
      {notification && <Toast message={notification.message} type={notification.type} onClose={() => setNotification(null)} />}
      
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Manajemen Penerbit</h1>
        <button onClick={() => {setModalMode('add'); setFormData(initialForm); setIsModalOpen(true)}} className="bg-zinc-900 text-white px-4 py-2 rounded-md flex items-center gap-2"><Plus className="w-4 h-4"/> Tambah</button>
      </div>
      
      <div className="bg-white rounded-lg border border-zinc-200 overflow-hidden">
        <div className="p-4 border-b">
           <input type="text" placeholder="Cari penerbit..." onChange={(e) => setSearchQuery(e.target.value)} className="w-full max-w-md p-2 border rounded text-sm"/>
        </div>
        <table className="w-full text-left text-sm">
          <thead className="bg-zinc-50 border-b">
            <tr><th className="px-6 py-3">Nama</th><th className="px-6 py-3">Email</th><th className="px-6 py-3">Telepon</th><th className="px-6 py-3 text-right">Aksi</th></tr>
          </thead>
          <tbody className="divide-y">
            {data.filter(i => i.penerbit_buku.toLowerCase().includes(searchQuery.toLowerCase())).map((item) => (
              <tr key={item.id} className="hover:bg-zinc-50">
                <td className="px-6 py-4 font-medium">{item.penerbit_buku}</td>
                <td className="px-6 py-4">{item.email_penerbit}</td>
                <td className="px-6 py-4">{item.telp_penerbit}</td>
                <td className="px-6 py-4 text-right flex justify-end gap-2">
                  <button onClick={() => {setModalMode('edit'); setFormData(item); setIsModalOpen(true)}} className="text-zinc-400 hover:text-blue-600"><Edit className="w-4 h-4"/></button>
                  <button onClick={() => handleDelete(item.id, item.penerbit_buku)} className="text-zinc-400 hover:text-red-600"><Trash2 className="w-4 h-4"/></button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {isModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-zinc-900/50 backdrop-blur-sm">
          <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg w-full max-w-md space-y-4">
            <div className="flex justify-between items-center"><h3 className="font-bold">Form Penerbit</h3><X className="cursor-pointer" onClick={() => setIsModalOpen(false)}/></div>
            <input required type="text" placeholder="Nama" value={formData.penerbit_buku} onChange={e => setFormData({...formData, penerbit_buku: e.target.value})} className="w-full p-2 border rounded"/>
            <input required type="email" placeholder="Email" value={formData.email_penerbit} onChange={e => setFormData({...formData, email_penerbit: e.target.value})} className="w-full p-2 border rounded"/>
            <input required type="text" placeholder="Telepon" value={formData.telp_penerbit} onChange={e => setFormData({...formData, telp_penerbit: e.target.value})} className="w-full p-2 border rounded"/>
            <textarea required placeholder="Alamat" value={formData.alamat_penerbit} onChange={e => setFormData({...formData, alamat_penerbit: e.target.value})} className="w-full p-2 border rounded" rows={2}/>
            <textarea required placeholder="Deskripsi" value={formData.deskripsi} onChange={e => setFormData({...formData, deskripsi: e.target.value})} className="w-full p-2 border rounded" rows={2}/>
            <button type="submit" disabled={isSubmitting} className="w-full bg-zinc-900 text-white p-2 rounded">{isSubmitting ? '...' : 'Simpan'}</button>
          </form>
        </div>
      )}
    </div>
  );
}