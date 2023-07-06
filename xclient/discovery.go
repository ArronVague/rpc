package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

type SelectMode int

const (
	RandonSelect     SelectMode = iota //随机策略
	RoundRobinSelect                   //轮询策略
)

type Discovery interface {
	Refresh() error                      //更新服务列表
	Update(servers []string) error       //手动更新服务列表
	Get(mode SelectMode) (string, error) //根据负载均衡策略选择服务实例
	GetAll() ([]string, error)           //返回所有服务实例
}

// 服务列表
type MultiServersDiscovery struct {
	r       *rand.Rand
	mu      sync.RWMutex
	servers []string
	index   int //记录轮询算法中，被选择的服务的索引
}

func NewMultiServerDiscovery(servers []string) *MultiServersDiscovery {
	d := &MultiServersDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())), //避免产生相同的随机数记录
	}
	d.index = d.r.Intn(math.MaxInt32 - 1) //随机设定index值
	return d
}

var _ Discovery = (*MultiServersDiscovery)(nil)

func (d *MultiServersDiscovery) Refresh() error {
	return nil
}

func (d *MultiServersDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

func (d *MultiServersDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	switch mode {
	case RandonSelect:
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		s := d.servers[d.index%n]
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

func (d *MultiServersDiscovery) GetAll() ([]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	servers := make([]string, len(d.servers), len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}
