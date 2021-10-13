from utils import log
from utils import init
from db import base

LOG = log.LOG
Base = base.Base


def main():
    LOG.info("start lucky process!")
    init.init_data()


if __name__ == "__main__":
    main()
