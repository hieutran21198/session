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
	GetAWSSession2FA() *AWSSession2FA
}

// New creates a new session service.
func New() *service {
	return &service{}
}

// GetAWSSession2FA gets the AWS session with 2FA.
func (s *service) GetAWSSession2FA() *AWSSession2FA {
	sess := &AWSSession2FA{}

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
