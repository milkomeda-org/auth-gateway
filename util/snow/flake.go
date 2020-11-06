// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package snow

import (
	"github.com/sony/sonyflake"
)

var (
	SFlake *Flake
)

// Flake SnowFlake算法结构体
type Flake struct {
	sFlake *sonyflake.Sonyflake
}

func init() {
	SFlake = NewSnowFlake()
}

// 模拟获取本机的机器ID
func getMachineID() (mID uint16, err error) {
	mID = 10
	return
}

func NewSnowFlake() *Flake {
	st := sonyflake.Settings{}
	// machineID是个回调函数
	st.MachineID = getMachineID
	return &Flake{
		sFlake: sonyflake.NewSonyflake(st),
	}
}

func (s *Flake) GetID() (uint64, error) {
	return s.sFlake.NextID()
}
