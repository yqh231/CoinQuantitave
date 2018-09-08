from decimal import Decimal
from datetime import datetime

from client.coinex_client import CoinExClient
from client.huobi_client import HuoBiClient

from db.mongo import TraceLog


class TradeSvr(CoinExClient, HuoBiClient):

    def __init__(self, *args, **kwargs):
        self.coinex = CoinExClient()
        self.huobi = HuoBiClient()
        self.market_type = ["EOSETH"]
        self.trace_log = TraceLog()

        super(CoinExClient, self).__init__()
        super(HuoBiClient, self).__init__()

    def run(self):
        for market_type in self.market_type:
            res_huobi = self.huobi.get_trade({"symbol": market_type.lower()})
            res_coinex = self.coinex.get_ticker({"market": market_type.upper()})
            if 'data' not in res_coinex or 'tick' not in res_huobi:
                return
            huobi_price = Decimal(res_huobi['tick']['data'][0]['price'])
            coinex_price = Decimal(res_coinex['data']['ticker']['last'])
            if coinex_price - huobi_price > huobi_price * Decimal(0.002 + 0.003):
                coinex_depth = self.coinex.get_depth({"market": market_type, "merge": '0.00001'})
                if 'data' not in coinex_depth:
                    return
                coinex_depth_1_price = Decimal(coinex_depth['data']['bids'][0][0])
                coinex_depth_1_amount = Decimal(coinex_depth['data']['bids'][1][1])
                if coinex_depth_1_price < coinex_price:
                    return
                
                huobi_depth = self.huobi.get_depth({"symbol": market_type.lower(), "type": "step4"})
                if 'tick' not in huobi_depth:
                    return
                
                huobi_depth_1_price = Decimal(huobi_depth['tick']['asks'][0])
                huobi_depth_1_amount = Decimal(huobi_depth['tick']['asks'][1])
                if huobi_depth_1_price > huobi_price:
                    return
                trade_amount = min(coinex_depth_1_amount, huobi_depth_1_amount)
                now = datetime.now()
                self.trace_log.insert({
                    'trade_amount': trade_amount,
                    'huobi_sell_price': huobi_depth_1_price,
                    'coinex_buy_price': coinex_depth_1_price,
                    'update_time': now,
                    'create_time': now
                })
                
               

        

            

