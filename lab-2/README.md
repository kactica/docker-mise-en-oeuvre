# Lab 2 - Créer un Dockerfile

Pour ce lab, il faut comprendre quelques bases du Dockerfile. Pour cela je vous redirige vers les sections suivantes:
- <a href="https://wiki.antoine-chiris.com/docker-deep-dive/dockerfile" target="_blank">Docker Deep Dive: Dockerfile</a>
- <a href="https://wiki.antoine-chiris.com/docker-deep-dive/build-et-publication-des-images" target="_blank">Docker Deep Dive: build et publication des images</a>

Pour bien comprendre le processus, nous allons pas recréer une image... Mais trois.

Exercice 1 :

- Simplement, recréer une image Nginx basique avec comme image de base Ubuntu 22.04. La page d'accueil nginx doit être remplacée par une autre page par défaut (au choix).

Exercice 2 :

- Créer une image docker qui permet de faire tourner mon code Python Flask suivant:

```python
from flask import Flask
from datetime import datetime

app = Flask(__name__)

@app.route('/')
def home():
    return f'<h1>Bienvenue sur Flask dans Docker!</h1><p>Heure actuelle: {datetime.now()}</p>'

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
```

Exercice 3 :

- Plus complexe maintenant, créer une image Docker légère pour une application Go en utilisant le multi-stage building. Voici le code de l'application:

```golang
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Bonjour, Docker!")
    })

    fmt.Println("Serveur démarré sur le port 8080")
    http.ListenAndServe(":8080", nil)
}
```

- Initialiser le dossier avec `go mod init main`

---

Maintenant, essayer de `push` une image sur votre repository hub.docker
