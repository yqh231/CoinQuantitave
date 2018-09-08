from client.client import Client


class CoinExClient(Client):

    def get_depth(self, params):
        url = "https://api.coinex.com/v1/market/depth"
        return self.get(url, params)

    def get_ticker(self, params):
        url = " https://api.coinex.com/v1/market/ticker"
        return self.get(url, params)
