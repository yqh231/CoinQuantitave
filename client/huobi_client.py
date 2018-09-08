from client.client import Client


class HuoBiClient(Client):
    def __init__(self, *args, **kwargs):
        super(HuoBiClient, self).__init__()

    def get_depth(self, params):
        url = "https://api.huobi.pro/market/depth"
        return self.get(url, params)
    
    def get_trade(self, params):
        url = "https://api.huobi.pro/market/trade"
        return self.get(url, params)
