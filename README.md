# go_cam

Start: go run server.go

SO: Windows

Libreria GO: Echo

Server HTTP

Registrazione senza audio

Chiamate:

    - startRecord
        - parametro pathname cartella dove salvare file
        - parametro stringa prefisso
        - (eventuale) parametro formato video (default mp4)
        - formato nome registrazione: <prefisso comunicato dal client>_<timestamp>.<formato video>  
    - stopRecord
        - risposta nome file generato

--------------
Electron porta 8080
