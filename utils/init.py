import settings
from db import add
from api import stocks
from db import models
from db import query
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


def initHistoryData():
    """init all stock history data"""
    ts_codes = [sk.ts_code for sk in query.get_all(models.Stocks)]

    if len(ts_codes)/100 > int(len(ts_codes)/100):
        pages = int(len(ts_codes)/100) + 1
    else:
        pages = int(len(ts_codes)/100)

    for page in range(pages):
        query_codes = ts_codes[page*10: page*10 + 10]
        ts_code_query = ','.join(query_codes)
        stocks.fetchHistory(ts_code_query, settings.TRADE_START_TIME, settings.TRADE_END_TIME)
