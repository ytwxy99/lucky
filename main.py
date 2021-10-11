import time

from utils import upgrade
from utils import log
from utils import init
from db import base

LOG = log.LOG
Base = base.Base

def main():
    upgrade.upgrade_tushare()
    LOG.info("start lucky process!")
    if init.initAllStock():
        LOG.info("Init all stocks information successful!")

if __name__ == "__main__":
    main()