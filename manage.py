import time

from trade.trade import TradeSvr


def main():
    trade_srv = TradeSvr()
    while True:
        trade_srv.run()
        time.sleep(1)


if __name__ == "__main__":
    main()