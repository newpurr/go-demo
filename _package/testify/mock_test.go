package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// mock的作用是赋予【调用者】动态生成期望响应的能力
// 调用者自己声明期望的入参及对应的返回
// 这样在测试调用者本身（调用者依赖于Mock实现的接口）时，
// 该调用者对Mock 接口的依赖就固化了，
// 这样就可以实现测试该类时，只有无状态逻辑，隔离http、db、redis等

// MessageService 通知客户被收取的费用
type MessageService interface {
	SendChargeNotification(int) error
}

// SMSService 是 MessageService 的实现
type SMSService struct{}

// MyService 使用 MessageService 来通知客户
type MyService struct {
	messageService MessageService
}

// SendChargeNotification 通过 SMS 来告知客户他们被收取费用
// 这就是我们将要模拟的方法
func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending Production Charge Notification")
	return nil
}

// ChargeCustomer 向客户收取费用
// 在真实系统中，我们会模拟这个
// 但是在这里，我想在每次运行测试时都赚点钱
func (a MyService) ChargeCustomer(value int) error {
	_ = a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}

// smsServiceMock
type smsServiceMock struct {
	// mocks 在测试无状态函数 (对应 FP 中的 pure function) 中意义不大，其应用场景主要在于处理不可控的第三方服务、数据库、磁盘读写等。如果这些服务的调用细节已经被封装到 interface 内部，调用方只看到了 interface 定义的一组方法，那么在测试中 mocks 就能控制第三方服务返回任意期望的结果，进而实现对调用方逻辑的全方位测试。
	mock.Mock
	SendChargeNotificationContent string
}

// 我们模拟的 smsService 方法
func (m *smsServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	// Called告诉mock对象方法已经被调用，并获取一个要返回的参数数组
	m.SendChargeNotificationContent = "SendChargeNotification content:" + strconv.Itoa(value)

	// 标记mock方法已调用
	returnArguments := m.Called(value)

	return returnArguments.Error(0)
}

// 我们将实现 MessageService 接口
// 这就意味着我们不得不改写在接口中定义的所有方法
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer 是个奇迹发生的地方
// 在这里我们将创建 SMSService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	// 声明当 100 传递给 SendChargeNotification 时，需要返回errors.New("")
	expected := 100
	err := errors.New("")
	call := smsService.
		On("SendChargeNotification", expected).
		Return(err)

	// 接下来，我们定义要测试的调用者服务
	// 然后调用方法
	// 由于这里我们注入的是Mock实现，
	// 所以MyService对于smsService的依赖就几乎没有热
	myService := MyService{smsService}
	_ = myService.ChargeCustomer(expected)

	// 断言我们声明的输入、输出同最终调用后的输入输出一致
	require.Equal(t, []*mock.Call{call}, smsService.ExpectedCalls)

	// 断言入参&出参，实际上没有任何作用
	// mock的作用是赋予【调用者】动态生成期望响应的能力
	// 调用者自己声明期望的入参及对应的返回
	// 这样在测试调用者本身（调用者依赖于Mock实现的接口）时，
	// 该调用者对Mock 接口的依赖就固化了，
	// 这样就可以实现测试该类时，只有无状态逻辑，隔离http、db、redis等
	call2 := smsService.ExpectedCalls[0]
	assert.Equal(t, "SendChargeNotification", call2.Method)
	assert.Equal(t, expected, call2.Arguments[0])

	assert.Equal(t, err, call2.ReturnArguments[0])
	assert.Equal(t, err, call.ReturnArguments.Error(0))
	assert.Equal(t, 0, call2.Repeatability)
	assert.Nil(t, call.WaitFor)

	// 断言经过myService.ChargeCustomer处理后，smsServiceMock内部获得的内容是正确的
	//  这一步断言实际上没有任何作用
	//  调用者规定入参和返回:
	//  	call := smsService.
	//		On("SendChargeNotification", expected).
	//		Return(err)
	//  mock实现内部直接获取调用者期望的返回进行返回:
	// func (m *smsServiceMock) SendChargeNotification(value int) error {
	// 		// 标记mock方法已调用
	// 		returnArguments := m.Called(value)
	//
	// 		return returnArguments.Error(0)
	// }
	assert.Equal(t, "SendChargeNotification content:"+strconv.Itoa(expected), smsService.SendChargeNotificationContent)
	fmt.Printf("value: %+v\n", smsService.SendChargeNotificationContent)

	// 最后，我们验证 myService.ChargeCustomer 调用了我们模拟的 SendChargeNotification 方法
	smsService.AssertExpectations(t)
}
