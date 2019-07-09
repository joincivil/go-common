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

// ParameterizerConfig contains the attributes needed to instantiate the Parameterizer contract
type ParameterizerConfig struct {
	MinDeposit                       *big.Int
	PMinDeposit                      *big.Int
	ApplyStageLength                 *big.Int
	PApplyStageLength                *big.Int
	CommitStageLength                *big.Int
	PCommitStageLength               *big.Int
	RevealStageLength                *big.Int
	PRevealStageLength               *big.Int
	DispensationPct                  *big.Int
	PDispensationPct                 *big.Int
	VoteQuorum                       *big.Int
	PVoteQuorum                      *big.Int
	PProcessBy                       *big.Int
	ChallengeAppealLength            *big.Int
	AppealChallengeCommitStageLength *big.Int
	AppealChallengeRevealStageLength *big.Int
}

// NewDefaultParameterizerConfig returns default values to use when creating the Parameterizer contract
func NewDefaultParameterizerConfig() *ParameterizerConfig {
	// values copied from civil-event-crawler
	// https://github.com/joincivil/go-common/blob/master/pkg/token/cvl_token_service_test.go
	return &ParameterizerConfig{
		MinDeposit:                       big.NewInt(10),    // minDeposit
		PMinDeposit:                      big.NewInt(100),   // pMinDeposit
		ApplyStageLength:                 big.NewInt(0),     // applyStageLength
		PApplyStageLength:                big.NewInt(120),   // pApplyStageLength
		CommitStageLength:                big.NewInt(18000), // commitStageLength
		PCommitStageLength:               big.NewInt(120),   // pCommitStageLength
		RevealStageLength:                big.NewInt(18000), // revealStageLength
		PRevealStageLength:               big.NewInt(120),   // pRevealStageLength
		DispensationPct:                  big.NewInt(50),    // dispensationPct
		PDispensationPct:                 big.NewInt(50),    // pDispensationPct
		VoteQuorum:                       big.NewInt(50),    // voteQuorum
		PVoteQuorum:                      big.NewInt(50),    // pVoteQuorum
		PProcessBy:                       big.NewInt(18000), // pProcessBy
		ChallengeAppealLength:            big.NewInt(18000), // challengeAppealLength
		AppealChallengeCommitStageLength: big.NewInt(16000), // appealChallengeCommitStageLength
		AppealChallengeRevealStageLength: big.NewInt(16000), // appealChallengeRevealStageLength
	}
}

// AsArray returns the ParameterizerConfig as an IntArray needed for the contract parameter
func (p *ParameterizerConfig) AsArray() []*big.Int {
	return []*big.Int{
		p.MinDeposit,
		p.PMinDeposit,
		p.ApplyStageLength,
		p.PApplyStageLength,
		p.CommitStageLength,
		p.PCommitStageLength,
		p.RevealStageLength,
		p.PRevealStageLength,
		p.DispensationPct,
		p.PDispensationPct,
		p.VoteQuorum,
		p.PVoteQuorum,
		p.PProcessBy,
		p.ChallengeAppealLength,
		p.AppealChallengeCommitStageLength,
		p.AppealChallengeRevealStageLength,
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

// GovernmentConfig contains the fields needed to instatiate the Government contract
type GovernmentConfig struct {
	AppealFeeAmount                    *big.Int
	RequestAppealPhaseLength           *big.Int
	JudgeAppealPhaseLength             *big.Int
	AppealSupermajorityPercentage      *big.Int
	AppealChallengeVoteDispensationPct *big.Int
	PDeposit                           *big.Int
	PCommitStageLength                 *big.Int
	PRevealStageLength                 *big.Int
	ConstitutionHash                   [32]byte
	ConstitutionURI                    string
}

// NewDefaultGovernmentConfig returns default values to use when creating the Government contract
func NewDefaultGovernmentConfig() *GovernmentConfig {
	// values copied from civil-event-crawler
	// https://github.com/joincivil/civil-events-crawler/blob/a7754bc767a7c0f09ef1ace4dd67b86b0b322326/pkg/contractutils/contractutils.go#L487
	return &GovernmentConfig{
		AppealFeeAmount:                    big.NewInt(1000),       // appealFeeAmount
		RequestAppealPhaseLength:           big.NewInt(36000),      // requestAppealPhaseLength
		JudgeAppealPhaseLength:             big.NewInt(36000),      // judgeAppealPhaseLength
		AppealSupermajorityPercentage:      big.NewInt(66),         // appealSupermajorityPercentage
		AppealChallengeVoteDispensationPct: big.NewInt(66),         // appealChallengeVoteDispensationPct
		PDeposit:                           big.NewInt(150),        // pDeposit
		PCommitStageLength:                 big.NewInt(120),        // pCommitStageLength
		PRevealStageLength:                 big.NewInt(120),        // pRevealStageLength
		ConstitutionHash:                   [32]byte{},             // constitutionHash
		ConstitutionURI:                    "http://madeupURL.com", // constitutionURI
	}
}

// AsArray returns a config fields as an array needed for the contract parameter
func (g *GovernmentConfig) AsArray() []interface{} {
	return []interface{}{
		g.AppealFeeAmount,
		g.RequestAppealPhaseLength,
		g.JudgeAppealPhaseLength,
		g.AppealSupermajorityPercentage,
		g.AppealChallengeVoteDispensationPct,
		g.PDeposit,
		g.PCommitStageLength,
		g.PRevealStageLength,
		g.ConstitutionHash,
		g.ConstitutionURI,
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
