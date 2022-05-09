package session

import "time"

// AWSSession2FA is a 2FA session for AWS.
/*
example:
{
    "Credentials": {
        "AccessKeyId": "ASIA3KI22IVFZJISHO64",
        "SecretAccessKey": "P/lM5OZRqpYH3mfYuOz1HZFeOEXvOnTvDaT1hZCO",
        "SessionToken": "IQoJb3JpZ2luX2VjEBwaDmFwLXNvdXRoZWFzdC0xIkgwRgIhAORnkqiALOL1YV/nzrumql5xY+n47SOi4tO+hvCCawZgAiEAzbYVwxqrpqMIebxjs2fzHy48/2125otHZXDGBpIZodYq+AE
I5f//////////ARAAGgw3Nzc5ODIxOTkxMTUiDOiySl8fslwoEUaN3irMAX7nolEejfE4WiWOVdTALIPJWzkgHNb5fsqoRqKO+VuoESkP+LVd1hJ+IOh55+UNppdwf1bPsw2Z2ubPpkFcVmzja7iZ7nZbWCdbr/Um2uujn1o
q+jnM8ynGzVh8M6l2fAQmFI/tzYr2KXJufFKngalOXOiWwmXm+2g6mtrIZrhjmfNCXa/XPvsQ1GBXHb/vxyYD6vTPL4ZoPbPhn7P4VKGoz7yzzm2wnlxbZLD0WmI7MmdBNTZ/+9nWvSDvzihO5pBlh112To21LNrATTDuk+K
TBjqXAYz38BatL+kVApePrrzwiWtiVBSoQF45VyNuW6Ivcq0jSczjfeKlj1uRhGuTP+jfp8K10tDqox0d9v9DioKyW6bT1hxeCTcb3P1hXK1PwT6qNXonSGZYiYmcq7I8F3Gh5EkquNi/3nhs9EwBu43yxFxh4hcDCtRYSoE
QxETx5vrsjHfvCHMYURc89x2HgHFTr1rgAvw/IVo=",
        "Expiration": "2022-05-09T15:26:38+00:00"
    }
}
*/
type AWSSession2FA struct {
	Credentials struct {
		AccessKeyId     string    `json:"AccessKeyId"`
		SecretAccessKey string    `json:"SecretAccessKey"`
		SessionToken    string    `json:"SessionToken"`
		Expiration      time.Time `json:"Expiration"`
	} `json:"credentials"`
}
