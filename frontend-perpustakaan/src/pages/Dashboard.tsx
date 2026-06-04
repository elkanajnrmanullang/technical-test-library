import { useState, useEffect } from 'react';
import { BookText, ArrowRightLeft, AlertCircle, Clock, CheckCircle, BookPlus, CheckSquare } from 'lucide-react';
import { api } from '../lib/axios';

export default function Dashboard() {
  const [stats, setStats] = useState({
    buku: 0,
    peminjamanAktif: 0,
    keterlambatan: 0
  });

  const [activities, setActivities] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchDashboardData = async (isInitial: boolean) => {
      if (isInitial) {
        setIsLoading(true);
      }
      try {
        const [bukuRes, peminjamanRes, dendaRes] = await Promise.allSettled([
          api.get('/buku'),
          api.get('/admin/peminjaman'),
          api.get('/admin/denda')
        ]);

        const rawBuku = bukuRes.status === 'fulfilled' ? (bukuRes.value.data?.data || bukuRes.value.data) : [];
        const rawPeminjaman = peminjamanRes.status === 'fulfilled' ? (peminjamanRes.value.data?.data || peminjamanRes.value.data) : [];
        const rawDenda = dendaRes.status === 'fulfilled' ? (dendaRes.value.data?.data || dendaRes.value.data) : [];

        const bukuData = Array.isArray(rawBuku) ? rawBuku : [];
        const peminjamanData = Array.isArray(rawPeminjaman) ? rawPeminjaman : [];
        const dendaData = Array.isArray(rawDenda) ? rawDenda : [];

        const today = new Date();
        today.setHours(0, 0, 0, 0);

        let terlambatCount = 0;

        peminjamanData.forEach((trx: any) => {
          if (trx.tgl_hrs_kembali) {
            const batas = new Date(trx.tgl_hrs_kembali);
            batas.setHours(0, 0, 0, 0);
            if (batas < today) {
              terlambatCount++;
            }
          }
        });

        setStats({
          buku: bukuData.length,
          peminjamanAktif: peminjamanData.length,
          keterlambatan: terlambatCount
        });

        const combinedActivities: any[] = [];

        bukuData.forEach(b => {
          if (b.created_at || b.updated_at) {
            combinedActivities.push({
              id: `B-${b.id_buku}`,
              type: 'BUKU',
              title: 'Penambahan Buku Baru',
              anggota: 'Sistem',
              date: b.created_at || b.updated_at,
              desc: `Judul: ${b.judul_buku || b.judul}`
            });
          }
        });

        peminjamanData.forEach(trx => {
          combinedActivities.push({
            id: `P-${trx.id}`,
            type: 'PINJAM',
            title: 'Peminjaman Baru',
            anggota: trx.id_anggota,
            date: trx.tgl_pinjam || trx.created_at || new Date().toISOString(),
            desc: `Batas pengembalian: ${trx.tgl_hrs_kembali ? new Date(trx.tgl_hrs_kembali).toLocaleDateString('id-ID') : '-'}`
          });
        });

        dendaData.forEach(d => {
          combinedActivities.push({
            id: `D-${d.id}`,
            type: 'DENDA',
            title: 'Pengembalian & Denda',
            anggota: d.id_anggota,
            date: d.tgl_kembali || d.created_at || new Date().toISOString(),
            desc: `Tagihan denda tercatat sebesar Rp ${d.jumlah_denda ? d.jumlah_denda.toLocaleString('id-ID') : '0'}`
          });
        });

        const localLogs = JSON.parse(localStorage.getItem('local_logs') || '[]');
        localLogs.forEach((log: any) => {
          combinedActivities.push(log);
        });

        combinedActivities.sort((a, b) => {
          const dateA = new Date(a.date).getTime() || 0;
          const dateB = new Date(b.date).getTime() || 0;
          return dateB - dateA;
        });
        
        setActivities(combinedActivities.slice(0, 8));

      } catch (error) {
      } finally {
        if (isInitial) {
          setIsLoading(false);
        }
      }
    };

    fetchDashboardData(true);

    const intervalId = setInterval(() => {
      fetchDashboardData(false);
    }, 5000);

    return () => clearInterval(intervalId);
  }, []);

  const statCards = [
    { title: 'Total Inventaris Buku', value: isLoading ? '...' : stats.buku, icon: BookText, color: 'text-blue-600', bg: 'bg-blue-50' },
    { title: 'Peminjaman Berjalan', value: isLoading ? '...' : stats.peminjamanAktif, icon: ArrowRightLeft, color: 'text-emerald-600', bg: 'bg-emerald-50' },
    { title: 'Kasus Keterlambatan', value: isLoading ? '...' : stats.keterlambatan, icon: AlertCircle, color: 'text-red-600', bg: 'bg-red-50' },
  ];

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold text-zinc-900 tracking-tight">Dashboard</h1>
          <p className="mt-1 text-sm text-zinc-500">Pemantauan real-time metrik sistem perpustakaan</p>
        </div>
        <div className="flex items-center gap-2 px-3 py-1.5 bg-emerald-50 text-emerald-700 rounded-full text-xs font-medium border border-emerald-200">
          <span className="relative flex h-2 w-2">
            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
            <span className="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
          </span>
          Sinkronisasi Aktif
        </div>
      </div>

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {statCards.map((item) => (
          <div key={item.title} className="bg-white p-5 rounded-lg border border-zinc-200 shadow-sm transition-all hover:shadow-md">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-sm font-medium text-zinc-500">{item.title}</h3>
              <div className={`p-2 rounded-md ${item.bg}`}>
                <item.icon className={`h-4 w-4 ${item.color}`} />
              </div>
            </div>
            <div className="flex items-baseline">
              <span className="text-2xl font-semibold text-zinc-900 tracking-tight">{item.value}</span>
            </div>
          </div>
        ))}
      </div>

      <div className="bg-white rounded-lg border border-zinc-200 shadow-sm overflow-hidden">
        <div className="px-5 py-4 border-b border-zinc-200 flex items-center justify-between">
          <h2 className="text-base font-semibold text-zinc-900">Aktivitas Sistem Terkini</h2>
          <Clock className="w-4 h-4 text-zinc-400" />
        </div>
        
        <div className="p-0">
          {isLoading ? (
            <div className="flex justify-center py-12">
              <p className="text-sm text-zinc-500">Memuat log aktivitas...</p>
            </div>
          ) : activities.length === 0 ? (
            <div className="flex flex-col items-center justify-center py-12 text-zinc-400 bg-zinc-50/50">
              <ArrowRightLeft className="h-8 w-8 mb-3 text-zinc-300" />
              <p className="text-sm font-medium">Belum ada aktivitas terekam di sistem.</p>
            </div>
          ) : (
            <ul className="divide-y divide-zinc-100">
              {activities.map((act) => (
                <li key={act.id} className="p-5 hover:bg-zinc-50/80 transition-colors">
                  <div className="flex gap-4">
                    <div className="shrink-0 mt-1">
                      {act.type === 'PINJAM' ? (
                        <div className="w-8 h-8 rounded-full bg-blue-50 border border-blue-100 flex items-center justify-center">
                          <CheckCircle className="w-4 h-4 text-blue-600" />
                        </div>
                      ) : act.type === 'BUKU' ? (
                        <div className="w-8 h-8 rounded-full bg-indigo-50 border border-indigo-100 flex items-center justify-center">
                          <BookPlus className="w-4 h-4 text-indigo-600" />
                        </div>
                      ) : act.type === 'KEMBALI' ? (
                        <div className="w-8 h-8 rounded-full bg-emerald-50 border border-emerald-100 flex items-center justify-center">
                          <CheckSquare className="w-4 h-4 text-emerald-600" />
                        </div>
                      ) : (
                        <div className="w-8 h-8 rounded-full bg-red-50 border border-red-100 flex items-center justify-center">
                          <AlertCircle className="w-4 h-4 text-red-600" />
                        </div>
                      )}
                    </div>
                    <div className="flex-1 space-y-1">
                      <div className="flex items-center justify-between">
                        <p className="text-sm font-medium text-zinc-900">{act.title}</p>
                        <p className="text-xs text-zinc-500 whitespace-nowrap">
                          {new Date(act.date).toLocaleDateString('id-ID', { hour: '2-digit', minute: '2-digit' })}
                        </p>
                      </div>
                      <p className="text-sm text-zinc-500 font-mono text-xs mb-1">Entitas Log: {act.anggota}</p>
                      <p className="text-sm text-zinc-600">{act.desc}</p>
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}