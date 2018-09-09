from db import mgo
from config import conf


class TraceLog(object):
    def __init__(self, *args, **kwargs):
        self.c = mgo[conf.db][conf.collection] 
    
    def insert(self, doc):
        self.c.insert_one(doc)

