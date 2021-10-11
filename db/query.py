from db import base
from utils import log

LOG = log.LOG
SESSION = base.SESSION

def query_by_id(model, value):
    """get records by specfied id"""
    try:
        if model.__tablename__ == "stocks":
            return SESSION.query(model).filter_by(stock_id=value).all()
    except:
        LOG.error("func: def query_by_id(model, value) --- get records failed by specified id: %s" % value)
        return False