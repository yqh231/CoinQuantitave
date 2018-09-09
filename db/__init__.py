from pymongo import MongoClient

from config import conf

mgo = MongoClient('mongodb://{}:{}/'.format(conf.host, conf.port))