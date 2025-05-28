## JSON Support

Aplikasi sederhana berbasis Go untuk memproses data dari file JSON dan mengekspornya ke file Excel.

## 🔧 Cara Penggunaan

1. Download file binary myapp-mac (Mac Intel), myapp-mac-arm64 (Intel Apple Chip), myapp-win.exe (Windows 64-bit)
2. Buka terminal
3. Pastikan workdir berada di dalam binary yang sudah di download tadi
    ```shell
    cd ~/Downloads/<nama_folder>
    ```
4. Pastikan file json berada dalam 1 folder binary
    ```shell
    my_app/
    ├── main.go                # File utama aplikasi
    ├── go.mod                 # Modul Go
    ├── README.md              # Dokumentasi penggunaan
    ├── input/                 # Direktori input files
    │   └── sample.json        # Contoh input JSON (opsional)
    ├── output/                # Folder hasil output (Excel)
    │   └── result-2025-05-27T11-00-00.xlsx
    ```
5. Jalankan perintah berikut
    ```bash
    my_app <lokasi_json_file>
    ```
   
   Contoh:

    ```bash
    my_app-mac input/smaple.json
    ```
6. Hasil output ada di folder `outputs/`