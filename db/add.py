from db import base
from db import query

SESSION = base.SESSION

def add_one(record, model, filter_value):
    """
        insert one record into specified model, if record is existed,
        we won't add it.
    """
    if not query.query_by_id(model, filter_value):
        SESSION.add(record)
        SESSION.commit()


def add_multiple(records):
    """
        add multiple records into specified model, if record is existed,
        we won't add it.
    """
    SESSION.add_all(records)