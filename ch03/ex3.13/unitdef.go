package unitdef

const (
	Kilo = 1000
	B    = 1
	KB   = Kilo * B
	MB   = Kilo * KB
	GB   = Kilo * MB
	TB   = Kilo * GB
	PB   = Kilo * TB
	EB   = Kilo * PB
	ZB   = Kilo * EB
	YB   = Kilo * ZB
)

func step() int {
	return YB / ZB
}
