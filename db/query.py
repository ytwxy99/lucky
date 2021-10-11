from db import base

SESSION = base.SESSION

def query_by_id(model, value):
    """get records by specfied id"""
    try:
        if model.__tablename__ == "stocks":
            return SESSION.query(model).filter_by(stock_id=value).all()
    except:
        return False