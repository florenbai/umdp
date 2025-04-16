package license

import (
	"github.com/shirou/gopsutil/v3/host"
)

// GetHostId MachineCode、PhysicalId
func GetHostId() string {
	info, _ := host.Info()
	return info.HostID
}
