package channel

import (
	"fmt"
	"sync"
)

// 定义资源类型
type Resource struct {
    ID int
}

// 定义资源池结构体
type ResourcePool struct {
    pool    chan *Resource
    counter int
    maxSize int
}

// 从资源池中获取资源
func (rp *ResourcePool) Get() (*Resource, error) {
    select {
    case res := <-rp.pool:
        return res, nil
    default:
        if rp.counter < rp.maxSize {
            rp.counter++
            return &Resource{ID: rp.counter}, nil
        }
        return nil, fmt.Errorf("no resources available and pool is at max capacity")
    }
}