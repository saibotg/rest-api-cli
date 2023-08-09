package nagios

type NagiosResultCode int

const (
	NagiosResultOk       NagiosResultCode = 0
	NagiosResultWarning  NagiosResultCode = 1
	NagiosResultCritical NagiosResultCode = 2
	NagiosResultUnknown  NagiosResultCode = 3
)

var resultText = map[NagiosResultCode]string{
	NagiosResultOk:       "OK",
	NagiosResultWarning:  "WARNING",
	NagiosResultCritical: "CRITICAL",
	NagiosResultUnknown:  "UNKNOWN",
}

type NagiosResult struct {
	ResultCode NagiosResultCode
	InfoText   string
}

func (n NagiosResult) ResultText() string {
	return resultText[n.ResultCode]
}
