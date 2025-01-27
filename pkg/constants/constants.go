// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package constants

import (
	"time"
)

const (
	DefaultPerms755 = 0o755

	BaseDirName = ".avalanche-cli"
	LogDir      = "logs"

	ServerRunFile      = "gRPCserver.run"
	AvalancheCliBinDir = "bin"
	RunDir             = "runs"

	SuffixSeparator             = "_"
	SidecarFileName             = "sidecar.json"
	GenesisFileName             = "genesis.json"
	GenesisMainnetFileName      = "genesis_mainnet.json"
	ElasticSubnetConfigFileName = "elastic_subnet_config.json"
	SidecarSuffix               = SuffixSeparator + SidecarFileName
	GenesisSuffix               = SuffixSeparator + GenesisFileName
	NodeFileName                = "node.json"
	NodeCloudConfigFileName     = "node_cloud_config.json"
	TerraformDir                = "terraform"
	AnsibleDir                  = "ansible"
	ClusterConfigFileName       = "cluster_config.json"
	SidecarVersion              = "1.4.0"

	MaxLogFileSize   = 4
	MaxNumOfLogFiles = 5
	RetainOldFiles   = 0 // retain all old log files

	RequestTimeout    = 3 * time.Minute
	E2ERequestTimeout = 30 * time.Second

	SimulatePublicNetwork = "SIMULATE_PUBLIC_NETWORK"
	FujiAPIEndpoint       = "https://api.avax-test.network"
	MainnetAPIEndpoint    = "https://api.avax.network"

	// this depends on bootstrap snapshot
	LocalAPIEndpoint = "http://127.0.0.1:9650"
	LocalNetworkID   = 1337

	DefaultTokenName = "TEST"

	HealthCheckInterval = 100 * time.Millisecond

	// it's unlikely anyone would want to name a snapshot `default`
	// but let's add some more entropy
	SnapshotsDirName             = "snapshots"
	DefaultSnapshotName          = "default-1654102509"
	BootstrapSnapshotArchiveName = "bootstrapSnapshot.tar.gz"
	BootstrapSnapshotLocalPath   = "assets/" + BootstrapSnapshotArchiveName
	BootstrapSnapshotURL         = "https://github.com/ava-labs/avalanche-cli/raw/main/" + BootstrapSnapshotLocalPath
	BootstrapSnapshotSHA256URL   = "https://github.com/ava-labs/avalanche-cli/raw/main/assets/sha256sum.txt"

	CliInstallationURL    = "https://raw.githubusercontent.com/ava-labs/avalanche-cli/main/scripts/install.sh"
	ExpectedCliInstallErr = "resource temporarily unavailable"
	EIPLimitErr           = "AddressLimitExceeded"
	KeyDir                = "key"
	KeySuffix             = ".pk"
	YAMLSuffix            = ".yml"
	ConfigDir             = "config"

	Enable = "enable"

	Disable = "disable"

	TimeParseLayout             = "2006-01-02 15:04:05"
	MinStakeWeight              = 1
	DefaultStakeWeight          = 20
	AVAXSymbol                  = "AVAX"
	DefaultFujiStakeDuration    = "48h"
	DefaultMainnetStakeDuration = "336h"
	// The absolute minimum is 25 seconds, but set to 1 minute to allow for
	// time to go through the command
	StakingStartLeadTime                  = 5 * time.Minute
	StakingMinimumLeadTime                = 25 * time.Second
	PrimaryNetworkValidatingStartLeadTime = 20 * time.Second
	AWSCloudServerRunningState            = "running"
	TerraformNodeConfigFile               = "node_config.tf"
	AvalancheCLISuffix                    = "-avalanche-cli"
	AWSDefaultCredential                  = "default"
	CertSuffix                            = "-kp.pem"
	AWSSecurityGroupSuffix                = "-sg"
	ExportSubnetSuffix                    = "-export.dat"
	SSHTCPPort                            = 22
	AvalanchegoAPIPort                    = 9650
	AvalanchegoP2PPort                    = 9651
	CloudServerStorageSize                = 1000
	OutboundPort                          = 0
	Terraform                             = "terraform"
	AnsiblePlaybook                       = "ansible-playbook"
	SetupNodePlaybook                     = "playbook/setupNode.yml"
	CopyStakingFilesPlaybook              = "playbook/copyStakingFiles.yml"
	ExportSubnetPlaybook                  = "playbook/exportSubnet.yml"
	GetNodeIDPlaybook                     = "playbook/getNodeID.yml"
	IsBootstrappedPlaybook                = "playbook/isBootstrapped.yml"
	IsSubnetSyncedPlaybook                = "playbook/isSubnetSynced.yml"
	TrackSubnetPlaybook                   = "playbook/trackSubnet.yml"
	AvalancheGoVersionPlaybook            = "playbook/avalancheGoVersion.yml"
	IsBootstrappedJSONFile                = "isBootstrapped.json"
	AvalancheGoVersionJSONFile            = "avalancheGoVersion.json"
	NodeIDJSONFile                        = "nodeID.json"
	SubnetSyncJSONFile                    = "isSubnetSynced.json"
	AnsibleInventoryDir                   = "inventories"
	AnsiblePlaybookDir                    = "playbook"
	AnsibleStatusDir                      = "status"
	AnsibleInventoryFlag                  = "-i"
	AnsibleExtraArgsIdentitiesOnlyFlag    = "--ssh-extra-args='-o IdentitiesOnly=yes'"
	AnsibleExtraVarsFlag                  = "--extra-vars"
	DefaultConfigFileName                 = ".avalanche-cli"
	DefaultConfigFileType                 = "json"
	WriteReadReadPerms                    = 0o644

	CustomVMDir = "vms"

	AvaLabsOrg          = "ava-labs"
	AvalancheGoRepoName = "avalanchego"
	SubnetEVMRepoName   = "subnet-evm"
	CliRepoName         = "avalanche-cli"

	AvalancheGoInstallDir = "avalanchego"
	SubnetEVMInstallDir   = "subnet-evm"

	SubnetEVMBin = "subnet-evm"

	DefaultNodeRunURL = "http://127.0.0.1:9650"

	APMDir                = ".apm"
	APMLogName            = "apm.log"
	DefaultAvaLabsPackage = "ava-labs/avalanche-plugins-core"
	APMPluginDir          = "apm_plugins"

	// #nosec G101
	GithubAPITokenEnvVarName = "AVALANCHE_CLI_GITHUB_TOKEN"

	ReposDir       = "repos"
	SubnetDir      = "subnets"
	NodesDir       = "nodes"
	VMDir          = "vms"
	ChainConfigDir = "chains"

	SubnetType                 = "subnet type"
	PrecompileType             = "precompile type"
	CustomAirdrop              = "custom-airdrop"
	NumberOfAirdrops           = "airdrop-addresses"
	SubnetConfigFileName       = "subnet.json"
	ChainConfigFileName        = "chain.json"
	PerNodeChainConfigFileName = "per-node-chain.json"

	GitRepoCommitName  = "Avalanche-CLI"
	GitRepoCommitEmail = "info@avax.network"

	AvaLabsMaintainers = "ava-labs"

	UpgradeBytesFileName      = "upgrade.json"
	UpgradeBytesLockExtension = ".lock"
	NotAvailableLabel         = "Not available"
	BackendCmd                = "avalanche-cli-backend"

	AvalancheGoCompatibilityVersionAdded = "v1.9.2"
	AvalancheGoCompatibilityURL          = "https://raw.githubusercontent.com/ava-labs/avalanchego/master/version/compatibility.json"
	SubnetEVMRPCCompatibilityURL         = "https://raw.githubusercontent.com/ava-labs/subnet-evm/master/compatibility.json"

	YesLabel = "Yes"
	NoLabel  = "No"

	SubnetIDLabel     = "SubnetID: "
	BlockchainIDLabel = "BlockchainID: "

	PluginDir = "plugins"

	Network        = "network"
	MultiSig       = "multi-sig"
	SkipUpdateFlag = "skip-update-check"
	LastFileName   = ".last_actions.json"

	DefaultWalletCreationTimeout = 5 * time.Second

	DefaultConfirmTxTimeout = 20 * time.Second
)
