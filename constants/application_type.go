package constants

type ApplicationType string

const (
	Unknown    ApplicationType = "UNKNOWN"
	CLOUD      ApplicationType = "CLOUD"
	EDGE       ApplicationType = "EDGE"
	SWG        ApplicationType = "SWG"
	HOST       ApplicationType = "HOST"
	VM         ApplicationType = "VM"
	DOCKER     ApplicationType = "DOCKER"
	IDS        ApplicationType = "IDS"
	IPS        ApplicationType = "IPS"
	FW         ApplicationType = "FW"
	NGFW       ApplicationType = "NGFW"
	DNS        ApplicationType = "DNS"
	LB         ApplicationType = "LB"
	GW         ApplicationType = "GW"
	VNF        ApplicationType = "VNF"
	SSLVPN     ApplicationType = "SSLVPN"
	OTHER      ApplicationType = "OTHER"
	ZTP_SERVER ApplicationType = "ZTP_SERVER"
	ZTP_CLIENT ApplicationType = "ZTP_CLIENT"
)
