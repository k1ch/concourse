package atccmd_test

import (
	"testing"

	"github.com/concourse/concourse/atc/atccmd"
	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/acme/autocert"
)

type CommandSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *CommandSuite) TestLetsEncryptDefaultIsUpToDate() {
	cmd := &atccmd.ATCCommand{}

	parser := flags.NewParser(cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	opt := parser.Find("run").FindOptionByLongName("lets-encrypt-acme-url")
	s.NotNil(opt)

	s.Equal(opt.Default, []string{autocert.DefaultACMEDirectory})
}

func (s *CommandSuite) TestInvalidConcurrencyLimit() {
	cmd := &atccmd.RunCommand{}
	flags.ParseArgs(cmd, []string{
		"--client-secret",
		"client-secret",
		"--concurrent-request-limit",
		"ListAllJobs=2",
		"--concurrent-request-limit",
		"ListAllJobs=2",
	})

	_, err := cmd.Runner([]string{})

	s.Errorf(err, "invalid concurrent request limits: multiple limits for 'ListAllJobs'")
}

func TestSuite(t *testing.T) {
	suite.Run(t, &CommandSuite{
		Assertions: require.New(t),
	})
}
