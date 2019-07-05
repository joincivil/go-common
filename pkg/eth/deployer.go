package eth

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/go-common/pkg/generated/contract"
)

// DeployerContractAddresses tracks smart contract addresses
type DeployerContractAddresses struct {
	NewsroomFactory       common.Address
	MultisigFactory       common.Address
	CivilTokenController  common.Address
	CreateNewsroomInGroup common.Address
	PLCR                  common.Address
	TCR                   common.Address
	CVLToken              common.Address
	Parameterizer         common.Address
	Government            common.Address

	// libraries
	AttributeStore   common.Address
	DLL              common.Address
	ECRecovery       common.Address
	MessagesAndCodes common.Address
}

// Deployer provides helper methods to deploy Civil smart contracts
type Deployer struct {
	eth       *Helper
	addresses *DeployerContractAddresses
}

// NewDeployer builds a new Deployer instance
func NewDeployer(helper *Helper, addresses *DeployerContractAddresses) *Deployer {
	return &Deployer{eth: helper, addresses: addresses}
}

// DeployECRecovery deploys ECRecovery library
func (d *Deployer) DeployECRecovery() error {
	addr, _, _, err := contract.DeployECRecoveryContract(d.eth.Auth, d.eth.Blockchain)
	if err != nil {
		return err
	}

	d.addresses.ECRecovery = addr

	return nil
}

// DeployMessagesAndCodes deploys the MessagesAndCodes library
func (d *Deployer) DeployMessagesAndCodes() error {
	addr, _, _, err := contract.DeployMessagesAndCodesContract(d.eth.Auth, d.eth.Blockchain)
	if err != nil {
		return err
	}

	d.addresses.MessagesAndCodes = addr

	return nil
}

// DeployNewsroomFactory deploys the NewsroomFactory contract and links required libraries
func (d *Deployer) DeployNewsroomFactory() error {
	if (d.addresses.ECRecovery == common.Address{}) {
		return errors.New("cannot deploy contract without ECRecovery library")
	}

	addr, _, _, err := DeployContractWithLinks(
		d.eth.Auth,
		d.eth.Blockchain,
		contract.NewsroomFactoryABI,
		contract.NewsroomFactoryBin,
		map[string]common.Address{"ECRecovery": d.addresses.ECRecovery},
		d.addresses.MultisigFactory,
	)
	if err != nil {
		return err
	}

	d.addresses.NewsroomFactory = addr

	return nil
}

// DeployCivilTokenController creates a new CivilTokenController contract
func (d *Deployer) DeployCivilTokenController() error {
	if (d.addresses.MessagesAndCodes == common.Address{}) {
		return errors.New("cannot deploy contract without MessagesAndCodes library")
	}

	addr, _, _, err := DeployContractWithLinks(
		d.eth.Auth,
		d.eth.Blockchain,
		contract.CivilTokenControllerContractABI,
		contract.CivilTokenControllerContractBin,
		map[string]common.Address{"MessagesAndCodes": d.addresses.MessagesAndCodes},
	)
	if err != nil {
		return err
	}

	d.addresses.CivilTokenController = addr

	return nil
}

// DeployCreateNewsroomInGroup creates a new instance of the CreateNewsroomInGroup contract
func (d *Deployer) DeployCreateNewsroomInGroup() error {
	if (d.addresses.MultisigFactory == common.Address{} || d.addresses.NewsroomFactory == common.Address{}) {
		return errors.New("cannot deploy contract without MultisigFactory and NewsroomFactory addresses set")
	}
	addr, _, _, err := contract.DeployCreateNewsroomInGroupContract(
		d.eth.Auth,
		d.eth.Blockchain,
		d.addresses.NewsroomFactory,
		d.addresses.CivilTokenController,
	)
	if err != nil {
		return err
	}
	d.addresses.CreateNewsroomInGroup = addr

	return nil
}

// DeployMultiSigWalletFactory creates a new instance of the MultiSigWalletFactory contract
func (d *Deployer) DeployMultiSigWalletFactory() error {
	addr, _, _, err := contract.DeployMultiSigWalletFactoryContract(
		d.eth.Auth,
		d.eth.Blockchain,
	)
	if err != nil {
		return err
	}
	d.addresses.MultisigFactory = addr

	return nil
}

// DeployCVLToken deploys the CVL Token contract
func (d *Deployer) DeployCVLToken() error {
	if (d.addresses.CivilTokenController == common.Address{}) {
		return errors.New("cannot deploy contract without CivilTokenController addresses set")
	}

	_initialAmount := &big.Int{}
	_initialAmount, _ = _initialAmount.SetString("1000000000000000000000000", 10)
	_tokenName := "CVL"
	_decimalUnits := uint8(18)
	_tokenSymbol := "CVL"

	addr, _, _, err := contract.DeployCVLTokenContract(
		d.eth.Auth,
		d.eth.Blockchain,
		_initialAmount,
		_tokenName,
		_decimalUnits,
		_tokenSymbol,
		d.addresses.CivilTokenController,
	)
	if err != nil {
		return err
	}

	d.addresses.CVLToken = addr

	return nil
}

// DeployDLL deploys the DLL library
func (d *Deployer) DeployDLL() error {
	addr, _, _, err := contract.DeployDLLContract(
		d.eth.Auth,
		d.eth.Blockchain,
	)
	if err != nil {
		return err
	}

	d.addresses.DLL = addr

	return nil
}

// DeployAttributeStore deploys the AttributeStore library
func (d *Deployer) DeployAttributeStore() error {
	addr, _, _, err := contract.DeployAttributeStoreContract(
		d.eth.Auth,
		d.eth.Blockchain,
	)
	if err != nil {
		return err
	}

	d.addresses.AttributeStore = addr

	return nil
}

// DeployPLCR deploys the PLCR contract
func (d *Deployer) DeployPLCR() error {
	if (d.addresses.DLL == common.Address{} || d.addresses.AttributeStore == common.Address{}) {
		return errors.New("cannot deploy contract without AttributeStore and DLL addresses set")
	}

	addr, _, _, err := DeployContractWithLinks(
		d.eth.Auth,
		d.eth.Blockchain,
		contract.CivilPLCRVotingContractABI,
		contract.CivilPLCRVotingContractBin,
		map[string]common.Address{
			"AttributeStore": d.addresses.AttributeStore,
			"DLL":            d.addresses.DLL,
		},
		d.addresses.CVLToken,
		d.addresses.CivilTokenController,
	)
	if err != nil {
		return err
	}

	d.addresses.PLCR = addr

	return nil
}

// DefaultParameterizerConfig returns default values to use when creating the Parameterizer contract
func DefaultParameterizerConfig() []*big.Int {
	// values copied from civil-event-crawler
	// https://github.com/joincivil/go-common/blob/master/pkg/token/cvl_token_service_test.go
	return []*big.Int{
		big.NewInt(10),    // minDeposit
		big.NewInt(100),   // pMinDeposit
		big.NewInt(0),     // applyStageLength
		big.NewInt(120),   // pApplyStageLength
		big.NewInt(18000), // commitStageLength
		big.NewInt(120),   // pCommitStageLength
		big.NewInt(18000), // revealStageLength
		big.NewInt(120),   // pRevealStageLength
		big.NewInt(50),    // dispensationPct
		big.NewInt(50),    // pDispensationPct
		big.NewInt(50),    // voteQuorum
		big.NewInt(50),    // pVoteQuorum
		big.NewInt(18000), // pProcessBy
		big.NewInt(18000), // challengeAppealLength
		big.NewInt(16000), // appealChallengeCommitStageLength
		big.NewInt(16000), // appealChallengeRevealStageLength
	}
}

// DeployParameterizer deploys the Parameterizer contract
func (d *Deployer) DeployParameterizer(config []*big.Int) error {

	if (d.addresses.DLL == common.Address{} || d.addresses.AttributeStore == common.Address{}) {
		return errors.New("cannot deploy contract without AttributeStore and DLL addresses set")
	}

	addr, _, _, err := DeployContractWithLinks(
		d.eth.Auth,
		d.eth.Blockchain,
		contract.ParameterizerContractABI,
		contract.ParameterizerContractBin,
		map[string]common.Address{
			"AttributeStore": d.addresses.AttributeStore,
			"DLL":            d.addresses.DLL,
		},
		d.addresses.CVLToken,
		d.addresses.PLCR,
		config,
	)
	if err != nil {
		return err
	}

	d.addresses.Parameterizer = addr

	return nil
}

// DefaultGovernmentConfig returns default values to use when creating the Parameterizer contract
func DefaultGovernmentConfig() []interface{} {
	// values copied from civil-event-crawler
	// https://github.com/joincivil/civil-events-crawler/blob/a7754bc767a7c0f09ef1ace4dd67b86b0b322326/pkg/contractutils/contractutils.go#L487
	return []interface{}{
		big.NewInt(1000),       // appealFeeAmount
		big.NewInt(36000),      // requestAppealPhaseLength
		big.NewInt(36000),      // judgeAppealPhaseLength
		big.NewInt(66),         // appealSupermajorityPercentage
		big.NewInt(66),         // appealChallengeVoteDispensationPct
		big.NewInt(150),        // pDeposit
		big.NewInt(120),        // pCommitStageLength
		big.NewInt(120),        // pRevealStageLength
		[32]byte{},             // constitutionHash
		"http://madeupURL.com", // constitutionURI
	}
}

// DeployGovernment creates a new Government contract
// mostly borrowed from https://github.com/joincivil/civil-events-crawler/blob/a7754bc767a7c0f09ef1ace4dd67b86b0b322326/pkg/contractutils/contractutils.go#L487
func (d *Deployer) DeployGovernment(config []interface{}, appellateAddr common.Address, governmentControllerAddr common.Address) error {
	addr, _, _, err := DeployContractWithLinks(
		d.eth.Auth,
		d.eth.Blockchain,
		contract.GovernmentContractABI,
		contract.GovernmentContractBin,
		map[string]common.Address{},
		append(
			[]interface{}{
				appellateAddr, // aka council multisig
				governmentControllerAddr,
				d.addresses.PLCR,
			},
			config...,
		)...,
	)
	if err != nil {
		return err
	}

	d.addresses.Government = addr

	return nil
}

// DeployTCR creates a new TCR contract
func (d *Deployer) DeployTCR() error {
	addr, _, _, err := DeployContractWithLinks(
		d.eth.Auth,
		d.eth.Blockchain,
		contract.CivilTCRContractABI,
		contract.CivilTCRContractBin,
		map[string]common.Address{
			"AttributeStore": d.addresses.AttributeStore,
			"DLL":            d.addresses.DLL,
		},
		d.addresses.CVLToken,
		d.addresses.PLCR,
		d.addresses.Parameterizer,
		d.addresses.Government,
	)
	if err != nil {
		return err
	}

	d.addresses.TCR = addr

	return nil
}
