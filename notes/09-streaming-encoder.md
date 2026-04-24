# Streaming encoder

- Sebelumnya kita belajar package json dengan melakukan konversi data JSON yang sudah dalam bentuk variable dan data string atau []byte
- Pada kenyataanya, kadang data JSON nya berasal dari Input berupa io.Reader (File, Network, Request Body)
- Kita bisa saja membaca semua datanya terlebih dahulu, lalu simpan di variable, baru lakukan konversi dari JSON, namun hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur untuk membaca data dari Stream

Next: []()
