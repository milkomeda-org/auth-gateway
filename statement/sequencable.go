// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/23.

package statement

type Sequence interface {
	GetID() uint
	GetParentID() uint
	AppendChildren([]Sequence)
}
