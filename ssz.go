package main

type SigningData struct {
	ObjectRoot [32]byte `json:"current_version" ssz-size:"32"`
	Domain     [32]byte `json:"domain" ssz-size:"32"`
}

type ForkData struct {
	CurrentVersion       [4]byte  `json:"current_version" ssz-size:"4"`
	GenesisValidatorRoot [32]byte `json:"genesis_validators_root" ssz-size:"32"`
}

func compute_deposit_domain(fork_version [4]byte) ([]byte, error) {
	domain_type := DOMAIN_DEPOSIT
	fork_data_root, err := compute_deposit_fork_data_root(fork_version)
	if err != nil {
		return nil, err
	}

	return append(domain_type[:], fork_data_root[:28]...), nil
}

func compute_deposit_fork_data_root(current_version [4]byte) ([]byte, error) {
	forkData := new(ForkData)
	forkData.CurrentVersion = current_version
	forkData.GenesisValidatorRoot = ZERO_BYTES32

	root, err := forkData.HashTreeRoot()
	if err != nil {
		return nil, err
	}

	return root[:], nil
}