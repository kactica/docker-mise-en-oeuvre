from flask import Flask
from datetime import datetime

app = Flask(__name__)


@app.route('/')
def home():
    return f'<h1>Bienvenue sur Flask dans Docker!</h1><p>Heure actuelle: {datetime.now()}</p>'


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
