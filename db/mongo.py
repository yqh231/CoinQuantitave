from db import mgo


class TraceLog(object):
    def __init__(self, *args, **kwargs):
        self.c = mgo['trx']['TraceLog'] 
    
    def insert(self, doc):
        self.c.insert_one(doc)

