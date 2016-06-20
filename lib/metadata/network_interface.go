// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metadata

import "net"

// NetworkEndpoint describes a network presence in the form a vNIC in sufficient detail that it can be:
// a. created - the vNIC added to a VM
// b. identified - the guestOS can determine which interface it corresponds to
// c. configured - the guestOS can configure the interface correctly
type NetworkEndpoint struct {
	// Common.Name - the nic alias requested (only one name and one alias possible in linux)
	// Common.ID - pci slot of the vnic allowing for interface identifcation in-guest
	Common

	// IP address to assign - nil if DHCP
	Static *net.IPNet `vic:"0.1" scope:"read-only" key:"staticip"`

	// Actual IP address assigned
	Assigned net.IP `vic:"0.1" scope:"read-write" key:"ip"`

	// The PCI slot for the vNIC - this allows for interface idenitifcaton in the guest
	PCISlot int32 `vic:"0.1" scope:"read-only" key:"pcislot"`

	// The network in which this information should be interpreted. This is embedded directly rather than
	// as a pointer so that we can ensure the data is consistent
	Network ContainerNetwork `vic:"0.1" scope:"read-only" key:"network"`
}

// ContainerNetwork is the data needed on a per container basis both for vSphere to ensure it's attached
// to the correct network, and in the guest to ensure the interface is correctly configured.
type ContainerNetwork struct {
	// The symbolic name of the network
	Name string `vic:"0.1" scope:"read-only" key:"name"`

	// The network scope the IP belongs to.
	// The IP address is the default gateway
	Gateway net.IPNet `vic:"0.1" scope:"read-only" key:"gateway"`
	// Should this gateway be the default route for containers on the network
	Default bool `vic:"0.1" scope:"read-only" key:"default"`

	// The set of nameservers associated with this network - may be empty
	Nameservers []net.IP `vic:"0.1" scope:"read-only" key:"dns"`

	// The IP range for this network
	FirstIP net.IP `vic:"0.1" scope:"read-only" key:"first_ip"`
	LastIP  net.IP `vic:"0.1" scope:"read-only" key:"last_ip"`
}