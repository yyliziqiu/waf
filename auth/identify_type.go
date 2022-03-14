package auth

const AccTypeConsumer = "consumer"
const AccTypeVisitor = "visitor"

type Identify struct {
	Base   Base   `json:"base"`
	Device Device `json:"device"`
}

func (i Identify) GetAccId() string {
	return i.Base.AccId
}

func (i Identify) GetAccType() string {
	return i.Base.AccType
}

func (i Identify) GetDeviceId() string {
	return i.Device.DeviceId
}

func (i Identify) GetDeviceType() string {
	return i.Device.DeviceType
}

func (i Identify) IsConsumer() bool {
	return i.GetAccType() == AccTypeConsumer
}

func (i Identify) IsVisitor() bool {
	return i.GetAccType() == AccTypeVisitor
}

type Base struct {
	AccId     string `json:"acc_id"`
	AccType   string `json:"acc_type"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
}

type Device struct {
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	DeviceMode string `json:"device_model"`
	AppVersion int    `json:"app_version"`
	OTAVersion int    `json:"ota_version"`
	IP         string `json:"ip"`
	Market     string `json:"market"`
	UpdatedAt  uint64 `json:"updated_at"`
}
