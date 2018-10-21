from client.client import Client


class HuoBiClient(Client):
    def __init__(self, *args, **kwargs):
        super(HuoBiClient, self).__init__()
        self.caller = 'huobi'

    def get_depth(self, params):
        url = "https://api.huobipro.com/market/depth"
        return self.get(url, params, self.caller)

    def get_trade(self, params):
        url = "https://api.huobipro.com/market/trade"
        return self.get(url, params, self.caller)

    def get_kline(self, params):
        params['AccessKeyId'] = 'fff-xxx-ssss-kkk'
        url = "https://api.huobipro.com/market/history/kline"
        return self.get(url, params, self.caller)

