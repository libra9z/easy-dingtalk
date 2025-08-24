package utils

import (
	"fmt"
	"time"

	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RequestOption struct {
	Protocal *string
	Domain   *string
	Headers  map[string]*string
}

func DoRequest(method string, path string, query map[string]*string, body interface{}, opts ...*RequestOption) (response *tea.Response, err error) {
	req := tea.NewRequest()
	req.Protocol = tea.String("https")
	req.Pathname = tea.String(path)
	req.Domain = tea.String("oapi.dingtalk.com")
	req.Method = tea.String(method)
	req.Headers = map[string]*string{
		"accept":       tea.String("application/json"),
		"content-type": tea.String("application/json; charset=utf-8"),
		"host":         tea.String("oapi.dingtalk.com"),
	}
	for _, opt := range opts {
		if opt.Protocal != nil {
			req.Protocol = opt.Protocal
		}
		if opt.Domain != nil {
			req.Domain = opt.Domain
		}
		if opt.Headers != nil {
			req.Headers = tea.Merge(req.Headers, opt.Headers)
		}
	}
	req.Query = query
	req.Body = tea.ToReader(util.ToJSONString(body))
	runtimeOptions := &tea.RuntimeObject{
		ConnectTimeout: tea.Int(int(time.Second) * 2),
	}
	response, err = tea.DoRequest(req, tea.ToMap(runtimeOptions))
	if err != nil {
		err = fmt.Errorf("%w", err)
		return
	}
	return
}
