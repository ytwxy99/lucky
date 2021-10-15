import settings
import requests

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


def fetchDetailHistory(ts_code, trade_date):
    """fetch detail history data by specified stock code"""
    history = API.daily_basic(ts_code=ts_code, trade_date=trade_date, fields='ts_code,trade_date,turnover_rate,volume_ratio,pe,pb').to_string().split("\n")
    return history


def fetchHistoryFromPublic(source, ts_code):
    """fetch hisotry data from public source"""
    start_time = settings.TRADE_START_TIME
    end_time = settings.TRADE_END_TIME

    try:
        url = settings.SOHU + "code=cn_%s3&start=%s&end=%s" % (ts_code.split(".")[0], start_time, end_time)
        return requests.get(url)[0], True
    except:
        return dict(), False



def getTradeCal(start_date, end_date, exchange='SSE'):
    cal = API.trade_cal(exchange='SSE', start_date=start_date, end_date=end_date)
    return cal