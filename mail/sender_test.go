package mail

import (
	"testing"

	"github.com/stepanleas/backend-master-class/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress)

	subject := "A test email"
	content := `
		<h1>Hello world</h1>
		<p>Some text info for the email</p>
	`

	to := []string{"noreply@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
