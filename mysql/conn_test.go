package mysql

import (
	"bufio"
	"github.com/XiaoMi/Gaea/util/mocks/pipeTest"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestMariadbConnWithoutDB 为用来测试数据库一开始连线的详细流程，以下测试不使用 MariaDB 的服务器，只是单纯的单元测试
func TestMariadbConnWithoutDB(t *testing.T) {
	// 函数测试开始
	t.Run("MariaDB连接 的抽换缓存测试", func(t *testing.T) {
		// 开始模拟
		mockClient, mockServer := pipeTest.NewDcServerClient(t, pipeTest.TestReplyFunc) // 产生 Gaea 和 MariaDB 模拟物件

		// 针对这次测试进行临时修改
		err := mockClient.OverwriteConnBufWrite(nil, writersPool.Get().(*bufio.Writer))
		mockClient.GetBufWriter().Reset(mockClient.GetConnWrite())
		require.Equal(t, err, nil)

		// 产生一开始的讯息和预期讯息
		msg0 := []uint8{0}  // 起始传送讯息
		correct := uint8(0) // 预期之后的正确讯息

		// 开始进行讯息操作

		// 写入部份
		mockClient.SendOrReceive(msg0) // 模拟客户端传送讯息
		require.Equal(t, msg0[0], correct)

		// 读取部份
		msg1 := mockClient.Reply(mockServer) // 模拟服务端接收讯息
		correct++
		require.Equal(t, msg1[0], correct)
	})
}
