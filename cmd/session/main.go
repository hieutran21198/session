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

// ActionKey represents the action key. It is used to identify the action.
type ActionKey string

// all of the action keys.
var (
	AWSMFA ActionKey = "aws_mfa"
	AWSAMD ActionKey = "aws_amd"
)

type block struct {
	prefix string
	suffix string
}

var (
	openningHeaderTag = "#BEGIN"
	closingHeaderTag  = "#END"

	// used for handle profile based on shell type.
	profile = flag.String("profile", "$HOME/.fish_profile_$USER", "profile name")
	shell   = flag.String("shell", string(model.Fish), "profile type")

	// AWS MFA session, it is used to get the session token. It uses action keys (AWSMFA,...) as the value.
	sessionType = flag.String("type", string(AWSMFA), "get session by type")

	awsamd = flag.String("aws-amd", "", "set ARN MFA Device to env variable")

	// clean content, it is used to clean the content of the profile.
	cleanContent = flag.String("clean", "", "clean profile")

	pshell model.ProfileShellType

	profileBlocks = map[ActionKey]*block{
		AWSMFA: {
			prefix: "AWS_MFA_SESSION",
			suffix: "AWS_MFA_SESSION",
		},
		AWSAMD: {
			prefix: "AWS_ARN_MFA_DEVICE",
			suffix: "AWS_ARN_MFA_DEVICE",
		},
	}
)

func getBlockWithHeaderTag(key ActionKey) *block {
	b := profileBlocks[key]
	if b == nil {
		panic("unknown block")
	}

	return &block{
		prefix: fmt.Sprintf("%s_%s", openningHeaderTag, b.prefix),
		suffix: fmt.Sprintf("%s_%s", closingHeaderTag, b.suffix),
	}
}

func main() {
	flag.Parse()
	*profile = os.ExpandEnv(*profile)
	*awsamd = strings.TrimSpace(*awsamd)
	pshell = model.ProfileShellType(*shell)

	if *cleanContent != "" {
		clean(ActionKey(*cleanContent))
		return
	}

	if *awsamd != "" {
		setAWSAMD()
	}

	setSession()
}

func clean(ak ActionKey) {
	b := getBlockWithHeaderTag(ak)
	util.DeleteContent(*profile, b.prefix, b.suffix)
}

func setSession() {
	sessSrv := session.New()

	switch *sessionType {
	case string(AWSMFA):
		setAWSSessionMFA(sessSrv)
	default:
		panic("unknown session type")
	}
}

func setAWSSessionMFA(sessSrv session.Service) {
	sess := sessSrv.GetAWSSessionMFA()
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

	block := getBlockWithHeaderTag(AWSMFA)

	util.SaveToProfile(*profile, fileContents, block.prefix, block.suffix, true)
}

func setAWSAMD() {
	os.Setenv("ARN_MFA_DEVICE", *awsamd)

	fileContents := []string{}
	if pshell == model.Fish {
		fileContents = append(fileContents, fmt.Sprintf("set -gx ARN_MFA_DEVICE %s", *awsamd))
	}

	if pshell == model.Zsh || pshell == model.Bash {
		fileContents = append(fileContents, fmt.Sprintf("export ARN_MFA_DEVICE=%s", *awsamd))
	}

	block := getBlockWithHeaderTag(AWSAMD)

	util.SaveToProfile(*profile, fileContents, block.prefix, block.suffix, true)
}
