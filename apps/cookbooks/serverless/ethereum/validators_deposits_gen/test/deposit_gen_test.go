package serverless_keygen

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/pretty"
	serverless_aws_automation "github.com/zeus-fyi/zeus/builds/serverless/aws_automation"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	"github.com/zeus-fyi/zeus/test/configs"

	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ServerlessDepositsGenTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

func (s *ServerlessDepositsGenTestSuite) TestServerlessSigningFunc() {
	s.Tc = configs.InitLocalTestConfigs()
	r := resty.New()
	auth := aegis_aws_auth.AuthAWS{
		Region:    "us-west-1",
		AccessKey: s.Tc.AccessKeyAWS,
		SecretKey: s.Tc.SecretKeyAWS,
	}
	fnUrl, err := serverless_aws_automation.CreateLambdaFunctionDepositGen(ctx, auth)
	s.Require().Nil(err)
	s.Require().NotEmpty(fnUrl)
	r.SetBaseURL(fnUrl)

	validatorCount := 3
	dgReq := bls_serverless_signing.EthereumValidatorDepositsGenRequests{
		MnemonicAndHDWalletSecretName: "mnemonicAndHDWalletEphemery",
		ValidatorCount:                validatorCount,
		HdOffset:                      0,
		Network:                       "ephemery",
	}
	req, err := auth.CreateV4AuthPOSTReq(ctx, "lambda", fnUrl, dgReq)
	s.Require().Nil(err)

	//depParams := make([]signing_automation_ethereum.DepositDataParams, validatorCount)
	resp, err := r.R().
		SetHeaderMultiValues(req.Header).
		SetBody(dgReq).Post("/")

	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode())

	fmt.Println("response json")
	respJSON := pretty.Pretty(resp.Body())
	respJSON = pretty.Color(respJSON, pretty.TerminalStyle)
	fmt.Println(string(respJSON))
}

func TestServerlessDepositsGenTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessDepositsGenTestSuite))
}