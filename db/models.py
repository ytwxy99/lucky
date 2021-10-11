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

# create all tables when the table is not existed
base.Base.metadata.create_all(base.ENGINE, checkfirst=True)