package lifecyclemanagement

import (
	"Middleware/App/impl"
	"sync"
	"time"
)

var mutex sync.Mutex

type Servant struct {
	Impl      impl.BookService
	id        int
	available bool
	expired   bool
	time      time.Time
}

type LifecycleManager struct {
	lease_time   int
	max_servants int
	pool         []*Servant
}

func NewLifecycleManager(max_servants int, lease_time int) LifecycleManager {
	lcm := LifecycleManager{}
	lcm.max_servants = max_servants
	lcm.lease_time = lease_time
	return lcm
}

func (lcm *LifecycleManager) Pooling() {
	current_quantity := len(lcm.pool)

	for i := current_quantity; i < lcm.max_servants; i++ {
		servant := &Servant{Impl: impl.BookService{}, id: i, available: true, expired: false, time: time.Now()}

		lcm.pool = append(lcm.pool, servant)
	}
}
func (lcm *LifecycleManager) GetServant() *Servant {

	var servant *Servant
	for i := 0; i < len(lcm.pool); i++ {
		if lcm.pool[i].available {
			servant = lcm.pool[i]
			servant.available = false
			servant.expired = false
			servant.time = time.Now()
			break
		}
	}
	return servant
}

func (lcm *LifecycleManager) ReturnServant(servant *Servant) {
	for i := 0; i < len(lcm.pool); i++ {
		if lcm.pool[i] == servant {
			mutex.Lock()
			lcm.pool[i].available = true
			lcm.pool[i].expired = false
			lcm.pool[i].time = time.Now()
			mutex.Unlock()
			break
		}
	}
}

func (lcm *LifecycleManager) Leasing() {
	for {
		time.Sleep(1000 * time.Millisecond)
		var new_pool []*Servant
		for i := 0; i < len(lcm.pool); i++ {
			if time.Now().Sub(lcm.pool[i].time).Seconds() > float64(lcm.lease_time) {
				// Expirou...
				mutex.Lock()
				lcm.pool[i].expired = true
				mutex.Unlock()
			} else {
				new_pool = append(new_pool, lcm.pool[i])
			}
		}
		lcm.pool = new_pool
	}
}

func (servant *Servant) Update_life() {
	mutex.Lock()
	servant.time = time.Now()
	mutex.Unlock()
}

func (servant *Servant) IsExpired() bool {
	return servant.expired
}
