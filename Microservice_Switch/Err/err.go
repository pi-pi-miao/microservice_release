package Err

var (
	ErrorResponseReadFail = ErrorResponse{
		HttpCode: 401,
		Err: Err{
			Error:     "read request Body err",
			ErrorCode: "001",
		},
	}

	ErrorJsonFailed = ErrorResponse{
		HttpCode: 400,
		Err: Err{
			Error:     "request body can not parse",
			ErrorCode: "002",
		},
	}

	ErrorPutEtcd = ErrorResponse{
		HttpCode: 500,
		Err: Err{
			Error:     "put etcd err",
			ErrorCode: "003",
		},
	}

	ErrorInsertDb = ErrorResponse{
		HttpCode: 500,
		Err: Err{
			Error:     "put etcd err",
			ErrorCode: "004",
		},
	}

	ErrorSelectDb = ErrorResponse{
		HttpCode: 500,
		Err: Err{
			Error:     "put etcd err",
			ErrorCode: "005",
		},
	}

	ErrGetEtcd = ErrorResponse{
		HttpCode: 500,
		Err: Err{
			Error:     "get etcd err",
			ErrorCode: "006",
		},
	}

	ErrorRequest = ErrorResponse{
		HttpCode: 400,
		Err: Err{
			Error:     "get etcd err",
			ErrorCode: "007",
		},
	}

	ErrorDeleteDb = ErrorResponse{
		HttpCode: 500,
		Err: Err{
			Error:     "delete Db err",
			ErrorCode: "008",
		},
	}

	ErrorUpdateDb = ErrorResponse{
		HttpCode: 500,
		Err: Err{
			Error:     "update Db err",
			ErrorCode: "008",
		},
	}
)

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	Err      Err
	HttpCode int
}
