import settings
from utils import getApi

API = getApi.API


def get_all_stocks():
    # get all stocks information
    stocks = API.query('stock_basic', exchange='', list_status='L',
                       fields='ts_code,symbol,name,area,industry,list_date')
    return stocks.to_string().split("\n")


def fetchHistory(ts_code, start_date, end_date):
    """fetch stock history data by specified stock code"""
    history = API.daily(ts_code=ts_code, start_date=start_date, end_date=end_date).to_string().split("\n")
    return history
