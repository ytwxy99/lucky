import sqlalchemy
from db import base
from db import query
from utils import log

LOG = log.LOG
SESSION = base.SESSION

def add_one(record, model, filter_value):
    """
        insert one record into specified model, if record is existed,
        we won't add it.
    """
    if not query.query_by_id(model, filter_value):
        try:
            SESSION.add(record)
            SESSION.commit()
            return True
        except sqlalchemy.exc.IntegrityError:
            LOG.error("func add_one(record, model, filter_value) -- add record error, may it is existed, record is : %s" % record)
            return False
    else:
        return True


def add_multiple(records):
    """
        add multiple records into specified model, if record is existed,
        we won't add it.
    """
    SESSION.add_all(records)