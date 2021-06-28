package gogenerate

import "testing"

func TestCode(t *testing.T) {
	cases := []struct {
		errCode ErrCode
		expect  string
	}{
		{ErrCodeOk, "OK"},
		{ErrCodeInvalidParams, "无效参数"},
		{ErrCodeTimeout, "超时"},
	}

	for _, testCase := range cases {
		if testCase.errCode.String() != testCase.expect {
			t.Errorf("error code %d description inconsistant actual:%s expect:%s", int(testCase.errCode), testCase.errCode, testCase.expect)
		}
	}
}
