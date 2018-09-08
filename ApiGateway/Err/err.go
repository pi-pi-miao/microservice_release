package Err

var (
	ErrorJsonFailed = ErrorResponse{
		HttpCode: 400,
		Error: Err{
			Error:     "request body can not parse",
			ErrorCode: "001",
		},
	}

	ErrorReadBodyFailed = ErrorResponse{
		HttpCode: 401,
		Error: Err{
			Error:     "read request Body err",
			ErrorCode: "002",
		},
	}

	ErrorPassword = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "user register err",
			ErrorCode: "003",
		},
	}

	ErrorPasswordNotSame = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "the password not same",
			ErrorCode: "004",
		},
	}

	ErrorUserName = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "the English not in Username",
			ErrorCode: "005",
		},
	}

	ErrorEmail = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "User register email not right",
			ErrorCode: "006",
		},
	}

	ErrorRegisterTime = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "Register time not right",
			ErrorCode: "007",
		},
	}

	ErrorRpcConnFailed = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "cant not conn",
			ErrorCode: "008",
		},
	}

	ErrorMethodFailed = ErrorResponse{
		HttpCode: 409,
		Error: Err{
			Error:     "the method err",
			ErrorCode: "009",
		},
	}

	ErrorNotRequest = ErrorResponse{
		HttpCode: 403,
		Error: Err{
			Error:     "the method err",
			ErrorCode: "010",
		},
	}

	ErrorTimeOut = ErrorResponse{
		HttpCode: 408,
		Error: Err{
			Error:     "the method time out",
			ErrorCode: "011",
		},
	}

	ErrorRequestFaild = ErrorResponse{
		HttpCode: 400,
		Error: Err{
			Error:     "request Failed",
			ErrorCode: "012",
		},
	}

	ErrorCall = ErrorResponse{
		HttpCode: 403,
		Error: Err{
			Error:     "request Failed",
			ErrorCode: "013",
		},
	}
)

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	Error    Err
	HttpCode int
}
