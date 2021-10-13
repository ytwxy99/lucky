import settings
from db import add
from api import stocks
from db import models
from db import query
from utils import log
from utils import upgrade

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
    all_history = list()
    ts_codes = [sk.ts_code for sk in query.get_all(models.Stocks)]

    if len(ts_codes)/100 > int(len(ts_codes)/100):
        pages = int(len(ts_codes)/100) + 1
    else:
        pages = int(len(ts_codes)/100)

    for page in range(pages):
        query_codes = ts_codes[page*100: page*100 + 100]
        ts_code_query = ','.join(query_codes)
        history = stocks.fetchHistory(ts_code_query, settings.TRADE_START_TIME, settings.TRADE_END_TIME)
        if not history:
            LOG.error("func initHistoryData() -- Init all history data failed")
            return False
        else:
            # NOTE(ytwxy99), remove title of items which index is 0.
            history.remove(history[0])
            all_history.extend(history)

    for history in all_history:
        h = [h for h in history.split(" ") if h]
        record = models.History(ts_code=h[1],
                                trade_date=h[2],
                                open=h[3],
                                high=h[4],
                                low=h[5],
                                close=h[6],
                                pre_close=h[7],
                                change=h[8],
                                pct_chg=h[9],
                                vol=h[10],
                                amount=h[11])
        if not add.add_one(record, models.History, record):
            LOG.error("func initAllStock() -- Init all stock information failed")
            return False

    return True

def init_data(check_upgrade=False):
    """init all data"""
    upgrade.upgrade_tushare(check_upgrade)
    # if initAllStock():
    #     LOG.info("Init all stocks information successful!")
    if initHistoryData():
        LOG.info("Init all history information successful!")
