package config

const (
	BOT_TOKEN = "xoxb-5845993184340-5829047534279-NX1pKTMepMiKDLhU5ZIBtCVT" // required for events api callback events
	APP_TOKEN = "xapp-1-A05QLSZFB5L-5836953999030-c783e22c2dc795b48e9fa07f06d2f16556c3548101f892f43c30f0f83d32ae82" // required for socket mode

	// if below line is uncommented, dialogflowcx,json must be present at the location relative to root
	//CREDENTIALS_PATH     = "/functions/config/dialogflowcx.json"
	CREDENTIALS_PATH = ""

	SLACK_SIGNING_SECRET = "4f57242cc30e4afe92c5b69389c2448d"
	// set this to false when url verification is not yet done
	//VERIFY_SECRET = true
	PROJECT_ID = "uber-poc-397822"
	AGENT_NAME = "projects/uber-poc-397822/locations/global/agents/95c6a58e-3d5f-4fd1-8ab6-09cf8b2b7828"
)
