from db import add
from api import stocks
from db import models
from utils import log

LOG = log.LOG

def initAllStock():
    """init all stock information"""
    sks = stocks.get_all_stocks()
    for index, sk in enumerate(sks):
        if index == 0:
            continue
        sk_items = [value for value in sk.split(" ") if value]
        record = models.Stocks(name=sk_items[3], stock_id=sk_items[2], ts_code=sk_items[1], classify=sk_items[5], region=sk_items[4])
        if not add.add_one(record, models.Stocks, sk_items[2]):
            LOG.error("func initAllStock() -- Init all stock information failed")
            return False

    return True
