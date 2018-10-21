from decimal import Decimal
from datetime import datetime

from client.coinex_client import CoinExClient
from client.huobi_client import HuoBiClient

from db.mongo import TraceLog


class TradeSvr(CoinExClient, HuoBiClient):

    def __init__(self, *args, **kwargs):
        self.coinex = CoinExClient()
        self.huobi = HuoBiClient()
        self.market_type = ["BTMETH"]
        self.trace_log = TraceLog()

        super(CoinExClient, self).__init__()
        super(HuoBiClient, self).__init__()

    def run(self):
        for market_type in self.market_type:
            coinex_depth = self.coinex.get_depth({"market": market_type, "merge": '0.00000001'})
            if 'data' not in coinex_depth:
                return

            coinex_depth_1_price = Decimal(coinex_depth['data']['bids'][0][0])
            coinex_depth_1_amount = Decimal(coinex_depth['data']['bids'][1][1])
            huobi_depth = self.huobi.get_depth({"symbol": market_type.lower(), "type": "step1"})
            if 'tick' not in huobi_depth:
                return

            huobi_depth_1_price = Decimal(huobi_depth['tick']['asks'][0][0])
            huobi_depth_1_amount = Decimal(huobi_depth['tick']['asks'][0][1])
            print(huobi_depth_1_price, coinex_depth_1_price)
            if huobi_depth_1_price - coinex_depth_1_price > huobi_depth_1_price * Decimal(0.002 + 0.003):
                trade_amount = min(coinex_depth_1_amount, huobi_depth_1_amount)
                now = datetime.now()
                self.trace_log.insert({
                    'trade_amount': str(trade_amount),
                    'huobi_sell_price': str(huobi_depth_1_price),
                    'coinex_buy_price': str(coinex_depth_1_price),
                    'update_time': now,
                    'create_time': now
                })
                print('time to trade {}'.format(str(now)))
