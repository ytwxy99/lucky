import time

from utils import upgrade
from utils import log
from utils import init
from db import base
from db import add
from db import models

LOG = log.LOG
Base = base.Base

def main():
    upgrade.upgrade_tushare()
    LOG.info("start lucky process!")
    init.initAllStock()

if __name__ == "__main__":
    main()