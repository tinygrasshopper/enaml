package enaml

type CloudConfigManifest struct {
	AZs         []AZ                `yaml:"azs,omitempty"`
	VMTypes     []VMType            `yaml:"vm_types,omitempty"`
	DiskTypes   []DiskPool          `yaml:"disk_types,omitempty"`
	Networks    []DeploymentNetwork `yaml:"networks,omitempty"`
	Compilation *Compilation        `yaml:"compilation,omitempty"`
}

type VMType struct {
	Name            string      `yaml:"name:omitempty"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

type AZ struct {
	Name            string      `yaml:"name:omitempty"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

func (s *CloudConfigManifest) AddAZ(az AZ) (err error) {
	s.AZs = append(s.AZs, az)
	return
}

func (s *CloudConfigManifest) AddVMType(vmt VMType) (err error) {
	s.VMTypes = append(s.VMTypes, vmt)
	return
}

func (s *CloudConfigManifest) AddDiskType(dskt DiskPool) (err error) {
	s.DiskTypes = append(s.DiskTypes, dskt)
	return
}

func (s *CloudConfigManifest) AddNetwork(ntw DeploymentNetwork) (err error) {
	s.Networks = append(s.Networks, ntw)
	return
}