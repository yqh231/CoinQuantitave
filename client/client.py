import requests

from db.mongo import exception_log 


class Client(object):

    def __init__(self, *args, **kwargs):
        self.session = requests.Session()

        self.session.headers.update({
            "Content-Type": "application/json; charset=utf-8",
            "Accept": "application/json",
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"
        })

    def get(self, url, params, caller):
            r = self.session.get(url, params=params, timeout=5).json()
            if caller == 'huobi' and r['status'] != 'ok':
                exception_log.insert({
                    'caller': 'huobi',
                    'position': 'Client',
                    'msg': r['err-msg'] if 'err_msg' in r else 'null'
                })
                raise Exception('error')

            return r

    def post(self, url, data):
        r = self.session.post(url, data=data)
        return r.json()
