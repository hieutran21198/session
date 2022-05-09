package session

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hieutran21198/session/util"
)

type service struct {
}

// Service is the interface for the session service.
type Service interface {
	GetAWSSessionMFA() *AWSSessionMFA
}

// New creates a new session service.
func New() Service {
	return &service{}
}

// GetAWSSessionMFA gets the AWS session with MFA.
func (s *service) GetAWSSessionMFA() *AWSSessionMFA {
	sess := &AWSSessionMFA{}

	var tokenStr string
	fmt.Println("Token from registered device:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokenStr = scanner.Text()

	if err := util.GetJSONOutputFromCMD("aws", sess, "sts", "get-session-token", "--serial-number", os.ExpandEnv("$ARN_MFA_DEVICE"), "--token-code", tokenStr); err != nil {
		panic(err)
	}

	return sess
}
