package postgresql

import (
	commonlog "github.com/polarismesh/polaris/common/log"
)

var (
	log      = commonlog.GetScopeOrDefaultByName(commonlog.StoreLoggerName)
	cacheLog = commonlog.GetScopeOrDefaultByName(commonlog.CacheLoggerName)
)
