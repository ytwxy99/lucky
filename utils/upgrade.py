import os

from utils import log

LOG = log.LOG


def upgrade_tushare(upgrade_tushare):
    """upgrade tushare to latest version"""
    if not upgrade_tushare:
        return

    if os.system("pip3 install tushare --upgrade"):
        LOG.error("upgrade tushare error")
    else:
        LOG.info("upgrade tushare successful")