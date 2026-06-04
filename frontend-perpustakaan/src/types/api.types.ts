export interface ApiResponse<T> {
  error: boolean;
  msg: string;
  data: T;
}

export interface AuthData {
  username: string;
  token: string;
  refresh_token: string;
}

export interface Buku {
  id_buku: string;
  isbn: string;
  id_kategori_buku: string;
  judul_buku: string;
  id_penulis_buku: string;
  id_penerbit_buku: string;
  tahun_terbit: string;
  stok_buku: number;
  rak_buku: string;
  deskripsi_buku: string;
  gambar_buku: string | null;
  kondisi_buku: string | null;
  created_at: string;
  updated_at: string;
}