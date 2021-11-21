// 产考了 https://github.com/afex/hystrix-go/blob/master/hystrix/rolling/rolling.go

package rolling

import (
	"sync"
	"time"
)

/*
Numbers tracks a numberBucket over number of time bucket.
The buckets is 10 seconds long with 1 second interval
+---------------------+
|1|2|3|4|5|6|7|8|9|10|
+---------------------+
*/
type Numbers struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

func NewNumbers() *Numbers {
	return &Numbers{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}
}

// Max returns the maximum value seen in last 10 seconds
func (r *Numbers) Max(now time.Time) float64 {
	var max float64

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}
	return max
}

// Max Sliding Window returns max value in the sliding window of k
func (r *Numbers) MaxSlidingWindow(now time.Time, k int) []float64 {
	nums := []float64{}
	max := []float64{}

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			nums = append(nums, bucket.Value)
		}
	}
	max = getMaxSlidingWindow(nums, k)

	return max
}

func getMaxSlidingWindow(nums []float64, k int) []float64 {
	// Descending queue
	q := []int{}
	ans := []float64{}

	for i := 0; i < len(nums); i++ {
		// delete first element from q if window greater than k
		for len(q) != 0 && q[0] <= i-k {
			q = q[1:]
		}
		// if new number is greater, remove the smaller number
		for len(q) != 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		// take the first element from q, update answer
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}

// Sum sums the value over the last 10 seconds
func (r *Numbers) Sum(now time.Time) float64 {
	var sum float64

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			sum += bucket.Value
		}
	}
	return sum
}

// Avg returns the average bucket value over the last 10 seconds
func (r *Numbers) Avg(now time.Time) float64 {
	return r.Sum(now) / 10
}

func (r *Numbers) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool

	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}

	return bucket
}

func (r *Numbers) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

// Increment increments the number in the current timeBucket
func (r *Numbers) Increment(n float64) {
	if n == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	b.Value += n
	r.removeOldBuckets()
}

// UpdateMax updates the maximum value in the current timeBucket
func (r *Numbers) UpdateMax(n float64) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	r.removeOldBuckets()
}
