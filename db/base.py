import os
import sqlalchemy
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# https://www.cnblogs.com/lsdb/p/9835894.html
ENGINE = sqlalchemy.create_engine('sqlite:///%s/lucky.db' % os.getcwd(), echo=False)
Base = declarative_base()
Session = sessionmaker(bind=ENGINE)
SESSION = Session()