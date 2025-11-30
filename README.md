D20 Kockadobó Alkalmazás

Az alkalmazás egy véletlenszerű D20 kockadobást végez el, és ennek eredményét webes felületen jeleníti meg.
A felhasználói felület elérhető a http://localhost:8080 url-en azon a gépen, ami az alkalmazást futtatja.

Build:
1. Telepítsük az operációs rendszerünknek megfelelő Go-t: https://go.dev/doc/install
2. A rendszerünk parancssoros felületével navigáljunk a forráskód könyvtárába.
3. Adjuk ki a következő build parancsok valamelyikét: 
    3.1 Linux rendszereken: go build -o d20app
    3.2 Windowson: go build -o d20app.exe
4. Ellenőrizzük, hogy létrejött-e a bináris futtatható állomány.