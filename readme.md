## Pengenalan HttpRouter
- Salah satu library Open Source untuk http handler di Go Lang
- Minimalis dan cepat
- Hanya untuk routing saja
**Instalasi**
```bash
go get github.com/julienschmidt/httprouter
```
- Mirip dengan `ServeMux()`, bedanya `ServeMux()` menghandle semua HTTP Method, sedangkan HttpRouter explisit terhadap HTTP Method tertentu.

## Params
- Digunakan untuk mengambil url dinamis
- Cth: `/product/:id`, dll