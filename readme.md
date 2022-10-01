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

## Router Parameter
### Named Parameter
- /profiles/:username
  - /profiles/rasyidev
  - /profiles/taeriyaki
  - Not Match: /profiles/, profiles/rasyidev/path

### Catch All Parameter
- /profiles/*whatever
  - /profiles/rasyidev
  - /profiles/taeriyaki/path
  - Not Match: /profiles

## Serve File
- `(r httprouter) ServeHTTP()`

## Panic Handler
- Menangkap pesan error dan mengirimnya di HTTP Response.

## Not Found Handler
- Menangkap error not found (404). 