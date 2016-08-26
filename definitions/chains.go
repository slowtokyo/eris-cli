package definitions

type Chain struct {
	// name of the chain
	Name string `json:"name" yaml:"name" toml:"name"`
	// chain_id of the chain
	ChainID string `mapstructure:"chain_id" json:"chain_id" yaml:"chain_id" toml:"chain_id"`
	// type of the chain
	ChainType string `mapstructure:"chain_type" json:"chain_type" yaml:"chain_type" toml:"chain_type"`
	Chain string `json:"chain,omitempty" yaml:"chain,omitempty" toml:"chain,omitempty"`

	// same fields as in the Service Struct/Service Specification
	Service      *Service      `json:"service,omitempty" yaml:"service,omitempty" toml:"service,omitempty"`
	Dependencies *Dependencies `json:"dependencies,omitempty" yaml:"dependencies,omitempty" toml:"dependencies,omitempty"`
	Maintainer   *Maintainer   `json:"maintainer,omitempty" yaml:"maintainer,omitempty" toml:"maintainer,omitempty"`
	Location     *Location     `json:"location,omitempty" yaml:"location,omitempty" toml:"location,omitempty"`
	Machine      *Machine      `json:"machine,omitempty" yaml:"machine,omitempty" toml:"machine,omitempty"`
	Operations   *Operation
}

func BlankChain() *Chain {
	return &Chain{
		Service:    BlankService(),
		Maintainer: BlankMaintainer(),
		Location:   BlankLocation(),
		Machine:    BlankMachine(),
		Operations: BlankOperation(),
	}
}
