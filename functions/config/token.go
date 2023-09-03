package config

const (
	BOT_TOKEN = "xoxb-5838129240597-5841951492402-PO0TDJNF9pr1qUuznQF5o6Hp" // required for events api callback events
	APP_TOKEN = "xapp-1-A05QP0V6SKX-5839371371029-5ccfe0e1b135ef4e2f752ad148007ab0fbfb84555428d8987949dd0231cb53a9" // required for socket mode

	// if below line is uncommented, dialogflowcx,json must be present at the location relative to root
	CREDENTIALS_PATH     = "/functions/config/dialogflowcx.json"
	//CREDENTIALS_PATH = ""

	SLACK_SIGNING_SECRET = "2a71a7cc3819b077538b63151b7af04b"
	// set this to false when url verification is not yet done
	//VERIFY_SECRET = true
	PROJECT_ID = "uber-poc-397822"
	AGENT_NAME = "projects/uber-poc-397822/locations/global/agents/95c6a58e-3d5f-4fd1-8ab6-09cf8b2b7828"
)
