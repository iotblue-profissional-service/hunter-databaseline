package cervello

type CervelloConfigurations struct {
	OrganizationID string
	ApplicationID  string
}

var (
	// olympic city deployment organization
	envApplicationID  = "5a347170-cabc-4d6f-9a6e-c63324c7c9ab"
	envOrganizationID = "348d69d2-6332-4c3a-817f-8c5556842f10"
	envAPIURL         = "https://api.cervello.olympiccity.prv"
	envBrokerAPIURL   = "https://broker.cervello.olympiccity.prv:7443"
	envAuthURI        = "https://accounts.cervello.olympiccity.prv"
	envAuthREALM      = "cervello"
	envAuthClientID   = "cervello-ui"
	envAuthUsername   = "gis@olympiccity.prv"
	envAuthPassword   = "gis1234"
	envAuthGrantType  = "password"
	cervelloSdkLogs   = "false"
	NatsHost          = "nats://nats:4222"
	keycloakHost      = "http://accounts.demo.cervello.io"
	keycloakRealm     = "cervello"

	// test deployment organization
	//envApplicationID  = "310803e8-2b80-48ba-b0cd-83bb1ac3a152"
	//envOrganizationID = "a433cd79-3533-4db0-b3a7-240595f428d7"
	//envAPIURL         = "https://api.cervello.olympiccity.prv"
	//envBrokerAPIURL   = "https://broker.cervello.olympiccity.prv:7443"
	//envAuthURI        = "https://accounts.cervello.olympiccity.prv"
	//envAuthREALM      = "cervello"
	//envAuthClientID   = "cervello-ui"
	//envAuthUsername   = "gis@olympiccity.prv"
	//envAuthPassword   = "gis1234"
	//envAuthGrantType  = "password"
	//cervelloSdkLogs   = "false"
	//NatsHost          = "nats://nats:4222"
	//keycloakHost      = "http://accounts.demo.cervello.io"
	//keycloakRealm     = "cervello"
	//
	//// olympic city demo
	//envApplicationID  = "75050789-dd8f-4cfe-9b86-0136376be621"
	//envOrganizationID = "6b099687-0ae0-4df0-9273-a06bd146eb1f"
	//envAPIURL         = "https://api.demo.cervello.io/"
	//envBrokerAPIURL   = "https://broker.demo.cervello.io"
	//envAuthURI        = "https://accounts.demo.cervello.io"
	//envAuthREALM      = "cervello"
	//envAuthClientID   = "cervello-ui"
	//envAuthUsername   = "gis@olympiccity.prv"
	//envAuthPassword   = "gis1234"
	////envAuthUsername  = "super@cervello.local"
	////envAuthPassword  = "Iotblue55$"
	//envAuthGrantType = "password"
	//cervelloSdkLogs  = "false"
	//NatsHost         = "nats://nats:4222"
	//keycloakHost     = "http://accounts.demo.cervello.io"
	//keycloakRealm    = "cervello"
)
