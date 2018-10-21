import sys
import requests
import os
from pymongo import InsertOne
sys.path.append(os.getcwd())

from db.mongo import CoinHistory

coins = ["BTC", "BCH", "ETH", "EOS", "XRP"]


def main():
    base_url = "http://coincap.io/history/"
    coin_history = CoinHistory()

    for coin in coins:
        resp = requests.get(base_url + coin).json()
        market_cap = resp['market_cap']
        price = resp['price']
        volume = resp['volume']
        request = []
        for item in market_cap:
            carrier = {
                "coin": coin,
                "time": item[0],
                "market_cap": item[1],
            }
            for i in price:
                if i[0] == item[0]:
                    carrier['price'] = i[1]

            for i in volume:
                if i[0] == item[0]:
                    carrier['volume'] = i[1]

            request.append(InsertOne(carrier))

        coin_history.bulk_write(request)


if __name__ == "__main__":
    main()
