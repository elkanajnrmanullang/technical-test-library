import { Link, useLocation, Outlet, useNavigate } from 'react-router-dom';
import { LayoutDashboard, BookText, ArrowRightLeft, Receipt, LogOut } from 'lucide-react';

export default function Layout() {
  const location = useLocation();
  const navigate = useNavigate();

  const menuItems = [
    { path: '/dashboard', icon: LayoutDashboard, label: 'Dashboard' },
    { path: '/buku', icon: BookText, label: 'Katalog Buku' },
    { path: '/peminjaman', icon: ArrowRightLeft, label: 'Transaksi' },
    { path: '/denda', icon: Receipt, label: 'Denda' },
  ];

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  return (
    <div className="flex h-screen bg-zinc-50">
      <div className="w-64 bg-zinc-950 text-white flex flex-col shrink-0">
        <div className="p-6 flex items-center gap-3">
          <div className="w-8 h-8 bg-white rounded-md flex items-center justify-center">
            <BookText className="w-5 h-5 text-zinc-950" />
          </div>
          <h2 className="text-xl font-bold tracking-tight">PERPUSTAKAAN</h2>
        </div>
        <div className="flex-1 py-4 flex flex-col">
          <p className="px-6 text-xs font-semibold text-zinc-500 uppercase tracking-wider mb-4">Menu Utama</p>
          <nav className="space-y-1 px-3">
            {menuItems.map((item) => {
              const isActive = location.pathname === item.path;
              return (
                <Link
                  key={item.path}
                  to={item.path}
                  className={`flex items-center gap-3 px-3 py-2 rounded-md transition-colors ${
                    isActive ? 'bg-white/10 text-white font-medium' : 'text-zinc-400 hover:text-white hover:bg-white/5'
                  }`}
                >
                  <item.icon className="w-5 h-5" />
                  {item.label}
                </Link>
              );
            })}
          </nav>
        </div>
        <div className="p-4 mt-auto">
          <button onClick={handleLogout} className="flex items-center gap-3 px-3 py-2 w-full text-zinc-400 hover:text-white hover:bg-white/5 rounded-md transition-colors">
            <LogOut className="w-5 h-5" />
            Keluar
          </button>
        </div>
      </div>

      <div className="flex-1 overflow-auto">
        <main className="p-8 max-w-7xl mx-auto">
          <Outlet />
        </main>
      </div>
    </div>
  );
}