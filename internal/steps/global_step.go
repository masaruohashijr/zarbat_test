package steps

import (
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/numbers"
	"zarbat_test/pkg/ports/sms"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var NumberSecondaryPort numbers.SecondaryPort
var NumberPrimaryPort numbers.PrimaryPort
var SmsSecondaryPort sms.SecondaryPort
var SmsPrimaryPort sms.PrimaryPort
var ResponsePlayLastRecording domains.ResponsePlayLastRecording
var ResponseGather domains.ResponseGather
var ResponseSay domains.ResponseSay
var ResponsePlay domains.ResponsePlay
var ResponseRecord domains.ResponseRecord
var ResponsePing domains.ResponsePing
var ResponsePause domains.ResponsePause
var ResponseRedirect domains.ResponseRedirect
var ResponseReject domains.ResponseReject
var ResponseHangup domains.ResponseHangup
var ResponseDial domains.ResponseDial
var ResponseDialNumber domains.ResponseDialNumber
var ResponseSMS domains.ResponseSMS
var ResponseMMS domains.ResponseMMS
var ResponseConference domains.ResponseConference
var AvailableNumbers []string
var IncomingNumbers []string
var Ch = make(chan string)
var CallSid = ""
var TestHash uint32

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	Ch = make(chan string)
	go services.RunServer(Ch, false)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	Configuration.ActionUrl = "http://zang.io/ivr/welcome/call"
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Append To "([^"]*)" config hang up$`, AppendToConfigHangUp)
	ctx.Step(`^"([^"]*)" configured as conference "([^"]*)" with size (\d+)$`, ConfiguredAsConferenceWithSize)
	ctx.Step(`^"([^"]*)" configured to dial "([^"]*)"$`, ConfiguredToDial)
	ctx.Step(`^"([^"]*)" configured to dial and send digits "([^"]*)" to "([^"]*)"$`, ConfiguredToDialAndSendDigitsTo)
	ctx.Step(`^"([^"]*)" configured to gather speech$`, ConfiguredToGatherSpeech)
	ctx.Step(`^"([^"]*)" configured to gather digits until "([^"]*)"$`, ConfiguredToGatherDigitsUntil)
	ctx.Step(`^"([^"]*)" configured to pause (\d+) seconds$`, ConfiguredToPauseSeconds)
	ctx.Step(`^"([^"]*)" configured to ping URL$`, ConfiguredToPingURL)
	ctx.Step(`^"([^"]*)" configured to play last recording$`, ConfiguredToPlayLastRecording)
	ctx.Step(`^"([^"]*)" configured to play tone "([^"]*)"$`, ConfiguredToPlayTone)
	ctx.Step(`^"([^"]*)" configured to record calls$`, ConfiguredToRecordCalls)
	ctx.Step(`^"([^"]*)" configured to record calls for download$`, ConfiguredToRecordCallsForDownload)
	ctx.Step(`^"([^"]*)" configured to say "([^"]*)"$`, ConfiguredToSay)
	ctx.Step(`^"([^"]*)" configured to send SMS "([^"]*)" to "([^"]*)"$`, ConfiguredToSendSMSTo)
	ctx.Step(`^"([^"]*)" configured to redirect to ping URL$`, ConfiguredToRedirectToPingURL)
	ctx.Step(`^"([^"]*)" configured to reject call$`, ConfiguredToRejectCall)
	ctx.Step(`^"([^"]*)" configured to hang up after (\d+) seconds$`, ConfiguredToHangUpAfterSeconds)
	ctx.Step(`^I list all available numbers$`, IListAllAvailableNumbers)
	ctx.Step(`^I should get to buy (\d+) from list$`, IShouldGetToBuyFromList)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" should had entered conference "([^"]*)"$`, ShouldHadEnteredConference)
	ctx.Step(`^"([^"]*)" should be able to listen to frequencies "([^"]*)"$`, ShouldBeAlaybleToListenToFrequencies)
	ctx.Step(`^"([^"]*)" should be able to view the SMS "([^"]*)"$`, ShouldBeAbleToViewTheSMS)
	ctx.Step(`^"([^"]*)" configured to send MMS "([^"]*)" and media "([^"]*)" to "([^"]*)"$`, ConfiguredToSendMMSAndMediaTo)
	ctx.Step(`^"([^"]*)" should be able to view the MMS "([^"]*)" with media "([^"]*)"$`, ShouldBeAbleToViewTheMMSWithMedia)
	ctx.Step(`^"([^"]*)" should be reset$`, ShouldBeReset)
	ctx.Step(`^"([^"]*)" should get a ping request on the URL$`, ShouldGetAPingRequestOnTheURL)
	ctx.Step(`^"([^"]*)" should get call cancel status$`, ShouldGetCallCancelStatus)
	ctx.Step(`^"([^"]*)" should get digits "([^"]*)" from "([^"]*)"$`, ShouldGetDigitsFrom)
	ctx.Step(`^"([^"]*)" should get last call duration more than or equals to (\d+)$`, ShouldGetLastCallDurationMoreThanOrEqualsTo)
	ctx.Step(`^"([^"]*)" should get speech "([^"]*)"$`, ShouldGetSpeech)
	ctx.Step(`^"([^"]*)" should get the incoming call from "([^"]*)"$`, ShouldGetTheIncomingCallFrom)
	ctx.Step(`^"([^"]*)" should get transcription "([^"]*)"$`, ShouldGetTranscription)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}
