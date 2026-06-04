import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite'; 
// import { VitePWA } from 'vite-plugin-pwa';

export default defineConfig({
  plugins: [
    react(),
    tailwindcss(), 
    // VitePWA({ ... MATIKAN SEMENTARA DENGAN KOMENTAR ...
    //   registerType: 'autoUpdate',
    //   ... (biarkan baris lainnya di bawah sini tetap dikomentari) ...
    // }),
  ],
});