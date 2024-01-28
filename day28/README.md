1. Key / Value
   - Contoh: DynamoDB
   - DynamoDB bersifat cloud-native karena tidak bisa berjalan secara lokal atau di cloud hibrid. DynamoDB hanya bisa berjalan pada Amazon Web Services (AWS). DynamoDB digunakan pada berbagai use case, diantaranya pengembangan aplikasi perangkat lunak, sebagai penyimpanan metadata media, sampai membuat platform game. Data pada DynamoDB disimpan dalam format key-value (JSON). Setiap item pada DynamoDB wajib memiliki Partition Key. Selain Partition Key, kita juga bisa mendefinisikan Sort Key untuk mengurutkan item. Jadi, peletakan item dilakukan berdasarkan Partition Key, dan item-item dengan Partition Key yang sama diurutkan berdasarkan Sort Key. Partition Key dan Sort Key ini berperan seperti primary key pada SQL.
   - Kelebihan DynamoDB antara lain:
     1. Menyediakan penyimpanan yang hampir tidak terbatas
     2. Tidak memerlukan server fisik karena dihandle AWS
     3. Aman karena sudah diakui secara global oleh PCI DSS, HIPAA, dan NIST
     4. Waktu respon cepat
        
    -Kekurangan DynamoDB:
     1. Opsi Query yang diberikan terbatas
     2. Sulit memprediksi biaya yang dipakai
     3. Tidak dapat menggunakan join table
  
        
2. Column - family
     - Contoh: HBase
     - HBase adalah database terdistribusi yang berorientasi pada kolom. HBase memiliki struktur data yang cukup sederhana, yang hanya terdiri atas Key dan Value. Pada HBase, Key terdiri atas Row Key, Column Family, Column, dan Timestamp. Sedangkan Value adalah data yang disimpan dalam bentuk 'byte array' yang bisa berupa data teks, angka, website pages, maupun data binary. Row Key juga berupa 'byte array' dan bertindak sebagai 'Primary Key'. Dalam suatu Table HBase, Row Key disusun berurutan pada bagian baris tabel, sedangkan Column Family, Column dan Timestamp menempati bagian kolom dari tabel tersebut. Data yang disimpan disortir berdasarkan urutan Row Key. Ilustrasi struktur data pada HBase
       ![image](https://github.com/galihpra/ALTA_BE19_galih/assets/146093220/cbacdf67-e884-47a1-bee8-c4ad39664bc1)

     - Kelebihan HBase:
       1. Dapat mengelola data yang sangat besar dalam sistem terdistribusi.
       2. Dapat bekerja secara otomatis maupun manual.
       3. Menjamin keutuhan data meskipun terjadi kegagalan pada beberapa komputer yang dipekerjakannya.
     
3. Graph
     - Contoh: JanusGraph
     - JanusGraph adalah database grafik terdistribusi yang dioptimalkan untuk skalabilitas dan performa. JanusGraph dirancang untuk menangani data grafik berskala besar dan menyediakan kemampuan kueri dan analitik grafik yang kuat. Contoh kasus penggunaan JanusGraph antara lain analisa social network dan mesin pencarian berbasis grafik. 
     - Kelebihan JanusGraph
         1. ACID (atomicy, consistencym isolation, and durability) transactions.
         2. Fleksibel karena dapat terhubung dengan berbagai penyimpanan backend seperti Cassandra, HBase, dan Google Cloud Bigtable.
         3. Memungkinkan query data kompleks yang efisien di dalam grafik.
            
     - Kekurangan JanusGraph
         1. Rumit dalam konfigurasi dan penyetelan
         2. Opsi bahasa Query yang terbatas, hanya menggunakan Gremlin untuk query grafik.
            
4. Document - based
     - Contoh: Amazon SimpleDB
     - Amazon SimpleDB lebih kompatibel untuk database yang tidak terlalu rumit di mana pengguna dengan cepat mencari dan mengakses data dalam format NoSQL. Contoh penggunaan Amazon SimpleDB adalah untuk menangani blog website, dimana data artikel, komentar, kategori, dan tag dapat disimpan dan diakses dengan cepat. Jenis database ini menggunakan data yang berisi sepasang key dan value yang disimpan di dokumen dengan format JSON atau XML. Berikut adalah ilustrasi struktur datanya:
       ![image](https://github.com/galihpra/ALTA_BE19_galih/assets/146093220/f5760690-069f-4cc4-bb21-4feef232ca45)

     - Kelebihan Amazon SimpleDB
         1. Operasional yang mudah
         2. Beban administrasi berkurang
         3. Data otomatis terindeks
         
     - Kekurangan Amazon SimpleDB
         1. Konsistensi yang tidak terlalu baik.
         2. Keterbatasan penyimpanan.
