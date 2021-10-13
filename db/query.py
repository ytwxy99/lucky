from db import base
from utils import log

LOG = log.LOG
SESSION = base.SESSION


def query_by_id(model, value):
    """get records by specfied id"""
    try:
        if model.__tablename__ == "stocks":
            return SESSION.query(model).filter_by(stock_id=value).all()
        if model.__tablename__ == "history":
            return SESSION.query(model).filter(model.ts_code == value.ts_code,
                                               model.trade_date == value.trade_date).all()
    except:
        LOG.error("func: def query_by_id(model, value) --- get records failed by specified id: %s" % value)
        return False


def get_all(model):
    """get all records by specified model"""
    try:
        return SESSION.query(model).all()
    except:
        LOG.error("func: def get_all(model, value) --- get all records failed by specified model: %s" % model)
        return False