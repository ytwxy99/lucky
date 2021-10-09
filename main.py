import time

from utils import upgrade
from utils import log
from api import stocks

LOG = log.LOG

def main():
    upgrade.windows()
    LOG.info("start lucky process!")


if __name__ == "__main__":
    main()