package redis_lock

// LuaCheckAndDeleteDistributionLock 判断是否拥有分布式锁的归属权，是则删除
const LuaCheckAndDeleteDistributionLock = `
  local lockerKey = KEYS[1]
  local targetToken = ARGV[1]
  local getToken = redis.call('get',lockerKey)
  if (not getToken or getToken ~= targetToken) then
    return 0
	else
		return redis.call('del',lockerKey)
  end
`
