// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/9/14.

package statement

// Execute 可执行接口
type Execute interface {
	// Execute 可返回err的执行接口
	Execute() error
}
