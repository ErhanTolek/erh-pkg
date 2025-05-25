package fallback

import "strconv"

type Response interface {
	GetItems(count int) []any
	TotalCount() int
	RawItems() []any
}

type RequestHandler interface {
	Call(endpoint string, params map[string]string) (Response, error)
}

type Executor struct {
	Handler RequestHandler
}

func NewExecutor(handler RequestHandler) *Executor {
	return &Executor{Handler: handler}
}

func (e *Executor) FallbackExecutor(endpoint string, params map[string]string, fallbackEndpoints map[string]int, sizeParamKey string) ([]any, error) {
	baseResp, err := e.Handler.Call(endpoint, params)
	if err != nil {
		return nil, err
	}

	baseCount := baseResp.TotalCount()

	requestedTotal, _ := strconv.Atoi(params[sizeParamKey])
	remainingSlots := requestedTotal - baseCount

	if remainingSlots <= 0 {
		return baseResp.GetItems(requestedTotal), nil
	}

	var fallbackItems []any

	for fallbackEndpoint, desiredCount := range fallbackEndpoints {
		if remainingSlots <= 0 {
			break
		}

		resp, err := e.Handler.Call(fallbackEndpoint, params)
		if err != nil {
			continue
		}

		items := resp.GetItems(desiredCount)
		fallbackItems = append(fallbackItems, items...)
		remainingSlots -= len(items)
	}

	return append(baseResp.GetItems(baseCount), fallbackItems...), nil
}
