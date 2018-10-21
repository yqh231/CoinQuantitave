from db import mgo


class TraceLog(object):
    def __init__(self, *args, **kwargs):
        self.c = mgo['trx']['TraceLog'] 

    def insert(self, doc):
        self.c.insert_one(doc)


class CoinHistory(object):
    def __init__(self, *args, **kwargs):
        self.c = mgo['trx']['coin_history']

    def insert(self, doc):
        self.c.insert_one(doc)

    def bulk_write(self, docs):
        if docs:
            self.c.bulk_write(docs)


class ExceptionLog(object):
    def __init__(self, *args, **kwargs):
        self.c = mgo['trx']['exception_log']

    def insert(self, doc):
        self.c.insert_one(doc)


exception_log = ExceptionLog()
