import time

from trade.turtule_monitor import Turtle 

ONE_HOUR = 60 * 60


def main():
    turtle = Turtle()
    while True:
        try:
            turtle.run()
            time.sleep(ONE_HOUR)
        except Exception as e:
            print(str(e))


if __name__ == "__main__":
    main()