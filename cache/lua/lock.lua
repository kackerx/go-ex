val = redis.call('get', KEYS[1])
if val == false then
    -- key不存在
    return redis.call('set', KEYS[1], ARGV[1], 'EX', ARGV[2])
elseif val == ARGV[1] then
    -- 其实上次加锁成功
    redis.call('expire', KEYS[1], ARGV[2])
    return 'OK'
else
    -- 锁被别人抢走
    return ''
end