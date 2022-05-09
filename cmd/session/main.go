package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hieutran21198/session/internal/model"
	"github.com/hieutran21198/session/internal/session"
	"github.com/hieutran21198/session/util"
)

var (
	sessionType  = flag.String("type", "aws2fa", "get session by type")
	profile      = flag.String("profile", "$HOME/.fish_profile_$USER", "profile name")
	shell        = flag.String("shell", string(model.Fish), "profile type")
	arnMFADevice = flag.String("amd", "", "set ARN MFA Device to env variable")
)

func main() {
	flag.Parse()
	*profile = os.ExpandEnv(*profile)
	pshell := model.ProfileShellType(*shell)

	sessSrv := session.New()

	*arnMFADevice = strings.TrimSpace(*arnMFADevice)
	if *arnMFADevice != "" {
		os.Setenv("ARN_MFA_DEVICE", *arnMFADevice)
		fileContents := []string{}
		if pshell == model.Fish {
			fileContents = append(fileContents, fmt.Sprintf("set -gx ARN_MFA_DEVICE %s", *arnMFADevice))
		}

		if pshell == model.Zsh || pshell == model.Bash {
			fileContents = append(fileContents, fmt.Sprintf("export ARN_MFA_DEVICE=%s", *arnMFADevice))
		}
		util.SaveToProfile(*profile, fileContents, "#BEGIN_AWS_ARN_MFA_DEVICE", "#END_AWS_ARN_MFA_DEVICE", true)
	}

	switch *sessionType {
	case "aws2fa":
		sess := sessSrv.GetAWSSession2FA()
		fileContents := []string{}
		if pshell == model.Fish {
			fileContents = append(fileContents, fmt.Sprintf("set -gx AWS_ACCESS_KEY_ID %s", sess.Credentials.AccessKeyId))
			fileContents = append(fileContents, fmt.Sprintf("set -gx AWS_SECRET_ACCESS_KEY %s", sess.Credentials.SecretAccessKey))
			fileContents = append(fileContents, fmt.Sprintf("set -gx AWS_SESSION_TOKEN %s", sess.Credentials.SessionToken))
		}
		if pshell == model.Zsh || pshell == model.Bash {
			fileContents = append(fileContents, fmt.Sprintf("export AWS_ACCESS_KEY_ID=%s", sess.Credentials.AccessKeyId))
			fileContents = append(fileContents, fmt.Sprintf("export AWS_SECRET_ACCESS_KEY=%s", sess.Credentials.SecretAccessKey))
			fileContents = append(fileContents, fmt.Sprintf("export AWS_SESSION_TOKEN=%s", sess.Credentials.SessionToken))
		}

		util.SaveToProfile(*profile, fileContents, "#BEGIN_AWS_2FA_SESSION", "#END_AWS_2FA_SESSION", true)
	default:
		panic("unknown session type")
	}
}
