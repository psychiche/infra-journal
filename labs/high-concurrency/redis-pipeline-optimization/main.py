import time
import redis

r = redis.Redis(host='localhost', port=6379)

N = 10000
batch = 100

start = time.time()
pipe = r.pipeline(transaction=False)
for i in range(N):
    pipe.set(f"k:{i}", "v")
    if (i + 1) % batch == 0:
        pipe.execute()
# flush remainder
pipe.execute()

dur = time.time() - start
print(f"ops={N} dur={dur:.3f}s ops/s={N/dur:.1f}")
