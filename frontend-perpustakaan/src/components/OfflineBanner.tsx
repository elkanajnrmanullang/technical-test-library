import { WifiOff } from 'lucide-react';
import { useOfflineStatus } from '../hooks/useOfflineStatus';

export default function OfflineBanner() {
  const isOffline = useOfflineStatus();

  if (!isOffline) return null;

  return (
    <div className="fixed bottom-0 left-0 right-0 z-50 bg-red-600 text-white px-4 py-3 shadow-lg flex items-center justify-center gap-3">
      <WifiOff className="w-5 h-5" />
      <p className="text-sm font-medium">
        Koneksi terputus. Anda sedang berada dalam mode luring (offline). Katalog buku yang telah tersimpan tetap dapat dilihat.
      </p>
    </div>
  );
}