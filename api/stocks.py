import settings
from utils import getApi

API = getApi.API

def get_all_stocks():
    # get all stocks information
    stocks = API.query('stock_basic', exchange='', list_status='L', fields='ts_code,symbol,name,area,industry,list_date')
    return stocks.to_string().split("\n")