import requests

class Client(object):

    def __init__(self, *args, **kwargs):
        self.session = requests.Session()

        self.session.headers.update({
            "Content-Type": "application/json; charset=utf-8",
            "Accept": "application/json",
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"
        })
    
    def get(self, url, params):
        r = self.session.get(url, params=params)
        return r.json()
    
    def post(self, url, data):
        r = self.session.post(url, data=data)
        return r.json()
        
        