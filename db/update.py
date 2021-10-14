from db import base
from utils import log

LOG = log.LOG
SESSION = base.SESSION


def update(model, record):
    """update data by specified table"""
    try:
        if model.__tablename__ == "history":
            SESSION.query(model).filter(model.ts_code == record["ts_code"],
                                        model.trade_date == record["trade_date"]).update(record)
            SESSION.commit()
            return True
    except:
        return False
