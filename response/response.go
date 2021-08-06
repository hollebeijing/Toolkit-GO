package response

type Response struct {
	Total  int           `json:"total"`
	Result bool          `json:"result"`
	Errors []*Error      `json:"errors"`
	Rows   interface{} `json:"rows"`
}

func NewResponseDefault() *Response {
	return &Response{
		Total:  0,
		Result: false,
		Errors: make([]*Error, 0),
		Rows:   nil,
	}
}

func NewResponse(total int, result bool, rows interface{}, errs []*Error) *Response {
	return &Response{
		Total:  total,
		Result: result,
		Errors: errs,
		Rows:   rows,
	}
}

func (r *Response) WithTotal(total int) *Response {
	r.Total = total
	return r
}

func (r *Response) WithResult(result bool) *Response {
	r.Result = result
	return r
}

func (r *Response) WithRows(rows interface{}) *Response {
	r.Rows = rows
	return r
}

func (r *Response) WithErrors(errs []*Error) *Response {
	r.Errors = errs
	return r
}

func (r *Response) AppendError(err *Error) *Response {
	r.Errors = append(r.Errors, err)
	return r
}

