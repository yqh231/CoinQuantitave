from decimal import Decimal

from client.huobi_client import HuoBiClient 
from tools.email import send_email


class BaseStrategy(HuoBiClient):

    def __init__(self, symbol='ethusdt'):
        super().__init__()
        self.symbol = symbol

    def _get_kiline(self, symbol, size=10, period='1day'):
        return self.get_kline({'symbol': symbol,
                                'size': size,
                                'period': period})

    def _get_average_price(self):
        klines = self._get_kiline(self.symbol)['data']
        average_price = sum([Decimal(k['close']) for k in klines]) / Decimal(len(klines))
        return average_price

    def _max_price(self):
        klines = self._get_kiline(self.symbol)['data']
        max_price = max([Decimal(k['close']) for k in klines])
        return max_price

    def _min_price(self):
        klines = self._get_kiline(self.symbol)['data']
        min_price = min([Decimal(k['close']) for k in klines])
        return min_price

    def long_poiot(self):
        max_price = self._max_price()
        trade_details = self.get_trade({'symbol': self.symbol})['tick']['data'][0]
        if Decimal(trade_details['price']) > max_price:
            send_email('time to do long',
                       'in {} current pirce is {}, max pirce is {}'.format(self.symbol, trade_details['price'], str(max_price)))

    def short_point(self):
        min_price = self._min_price()
        trade_details = self.get_trade({'symbol': self.symbol})['tick']['data'][0]
        if trade_details['price'] < min_price:
            send_email('time to do short',
                       'in {} current pirce is {}, max pirce is {}'.format(self.symbol, trade_details['price'], str(min_price)))




