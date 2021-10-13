from db import base
from sqlalchemy import Column, Integer, String


class Stocks(base.Base):
    # set table name
    __tablename__ = 'stocks'

    stock_id = Column(String(20), primary_key=True)
    ts_code = Column(String(20))
    name = Column(String(20))
    classify = Column(String(32))
    region = Column(String(32))

    def __repr__(self):
        return "<Stocks(name='%s', stock_id='%s', classify='%s')>" % (
            self.name, self.stock_id, self.classify)


class History(base.Base):
    # set table name
    __tablename__ = 'history'

    id = Column(Integer, primary_key=True)
    ts_code = Column(String(20))
    trade_date = Column(String(20))
    open = Column(String(10))
    high = Column(String(10))
    low = Column(String(10))
    close = Column(String(10))
    pre_close = Column(String(10))
    change = Column(String(10)) # 涨跌额
    pct_chg = Column(String(10)) # 涨跌幅（未复权）
    vol = Column(String(20)) # 成交量
    amount = Column(String(20)) # 成交额
    turnover_rate = Column(String(20), nullable=True)
    turnover_rate_f = Column(String(20), nullable=True)  # 流通换手率
    volume_ratio = Column(String(20), nullable=True)  # 量比

    def __repr__(self):
        return "<History(ts_code='%s', open='%s', high='%s')>" % (
            self.ts_code, self.open, self.high)


# create all tables when the table is not existed
base.Base.metadata.create_all(base.ENGINE, checkfirst=True)
