// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BrokerAZDistribution string

// Enum values for BrokerAZDistribution
const (
	BrokerAZDistributionDefault BrokerAZDistribution = "DEFAULT"
)

// Values returns all known values for BrokerAZDistribution. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BrokerAZDistribution) Values() []BrokerAZDistribution {
	return []BrokerAZDistribution{
		"DEFAULT",
	}
}

type ClientBroker string

// Enum values for ClientBroker
const (
	ClientBrokerTls          ClientBroker = "TLS"
	ClientBrokerTlsPlaintext ClientBroker = "TLS_PLAINTEXT"
	ClientBrokerPlaintext    ClientBroker = "PLAINTEXT"
)

// Values returns all known values for ClientBroker. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClientBroker) Values() []ClientBroker {
	return []ClientBroker{
		"TLS",
		"TLS_PLAINTEXT",
		"PLAINTEXT",
	}
}

type ClusterState string

// Enum values for ClusterState
const (
	ClusterStateActive          ClusterState = "ACTIVE"
	ClusterStateCreating        ClusterState = "CREATING"
	ClusterStateDeleting        ClusterState = "DELETING"
	ClusterStateFailed          ClusterState = "FAILED"
	ClusterStateHealing         ClusterState = "HEALING"
	ClusterStateMaintenance     ClusterState = "MAINTENANCE"
	ClusterStateRebootingBroker ClusterState = "REBOOTING_BROKER"
	ClusterStateUpdating        ClusterState = "UPDATING"
)

// Values returns all known values for ClusterState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClusterState) Values() []ClusterState {
	return []ClusterState{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"FAILED",
		"HEALING",
		"MAINTENANCE",
		"REBOOTING_BROKER",
		"UPDATING",
	}
}

type ClusterType string

// Enum values for ClusterType
const (
	ClusterTypeProvisioned ClusterType = "PROVISIONED"
	ClusterTypeServerless  ClusterType = "SERVERLESS"
)

// Values returns all known values for ClusterType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ClusterType) Values() []ClusterType {
	return []ClusterType{
		"PROVISIONED",
		"SERVERLESS",
	}
}

type ConfigurationState string

// Enum values for ConfigurationState
const (
	ConfigurationStateActive       ConfigurationState = "ACTIVE"
	ConfigurationStateDeleting     ConfigurationState = "DELETING"
	ConfigurationStateDeleteFailed ConfigurationState = "DELETE_FAILED"
)

// Values returns all known values for ConfigurationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationState) Values() []ConfigurationState {
	return []ConfigurationState{
		"ACTIVE",
		"DELETING",
		"DELETE_FAILED",
	}
}

type EnhancedMonitoring string

// Enum values for EnhancedMonitoring
const (
	EnhancedMonitoringDefault              EnhancedMonitoring = "DEFAULT"
	EnhancedMonitoringPerBroker            EnhancedMonitoring = "PER_BROKER"
	EnhancedMonitoringPerTopicPerBroker    EnhancedMonitoring = "PER_TOPIC_PER_BROKER"
	EnhancedMonitoringPerTopicPerPartition EnhancedMonitoring = "PER_TOPIC_PER_PARTITION"
)

// Values returns all known values for EnhancedMonitoring. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnhancedMonitoring) Values() []EnhancedMonitoring {
	return []EnhancedMonitoring{
		"DEFAULT",
		"PER_BROKER",
		"PER_TOPIC_PER_BROKER",
		"PER_TOPIC_PER_PARTITION",
	}
}

type KafkaVersionStatus string

// Enum values for KafkaVersionStatus
const (
	KafkaVersionStatusActive     KafkaVersionStatus = "ACTIVE"
	KafkaVersionStatusDeprecated KafkaVersionStatus = "DEPRECATED"
)

// Values returns all known values for KafkaVersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KafkaVersionStatus) Values() []KafkaVersionStatus {
	return []KafkaVersionStatus{
		"ACTIVE",
		"DEPRECATED",
	}
}

type NodeType string

// Enum values for NodeType
const (
	NodeTypeBroker NodeType = "BROKER"
)

// Values returns all known values for NodeType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (NodeType) Values() []NodeType {
	return []NodeType{
		"BROKER",
	}
}

type ReplicatorState string

// Enum values for ReplicatorState
const (
	ReplicatorStateRunning  ReplicatorState = "RUNNING"
	ReplicatorStateCreating ReplicatorState = "CREATING"
	ReplicatorStateUpdating ReplicatorState = "UPDATING"
	ReplicatorStateDeleting ReplicatorState = "DELETING"
	ReplicatorStateFailed   ReplicatorState = "FAILED"
)

// Values returns all known values for ReplicatorState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicatorState) Values() []ReplicatorState {
	return []ReplicatorState{
		"RUNNING",
		"CREATING",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type StorageMode string

// Enum values for StorageMode
const (
	StorageModeLocal  StorageMode = "LOCAL"
	StorageModeTiered StorageMode = "TIERED"
)

// Values returns all known values for StorageMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageMode) Values() []StorageMode {
	return []StorageMode{
		"LOCAL",
		"TIERED",
	}
}

type TargetCompressionType string

// Enum values for TargetCompressionType
const (
	TargetCompressionTypeNone   TargetCompressionType = "NONE"
	TargetCompressionTypeGzip   TargetCompressionType = "GZIP"
	TargetCompressionTypeSnappy TargetCompressionType = "SNAPPY"
	TargetCompressionTypeLz4    TargetCompressionType = "LZ4"
	TargetCompressionTypeZstd   TargetCompressionType = "ZSTD"
)

// Values returns all known values for TargetCompressionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TargetCompressionType) Values() []TargetCompressionType {
	return []TargetCompressionType{
		"NONE",
		"GZIP",
		"SNAPPY",
		"LZ4",
		"ZSTD",
	}
}

type UserIdentityType string

// Enum values for UserIdentityType
const (
	UserIdentityTypeAwsaccount UserIdentityType = "AWSACCOUNT"
	UserIdentityTypeAwsservice UserIdentityType = "AWSSERVICE"
)

// Values returns all known values for UserIdentityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UserIdentityType) Values() []UserIdentityType {
	return []UserIdentityType{
		"AWSACCOUNT",
		"AWSSERVICE",
	}
}

type VpcConnectionState string

// Enum values for VpcConnectionState
const (
	VpcConnectionStateCreating     VpcConnectionState = "CREATING"
	VpcConnectionStateAvailable    VpcConnectionState = "AVAILABLE"
	VpcConnectionStateInactive     VpcConnectionState = "INACTIVE"
	VpcConnectionStateDeactivating VpcConnectionState = "DEACTIVATING"
	VpcConnectionStateDeleting     VpcConnectionState = "DELETING"
	VpcConnectionStateFailed       VpcConnectionState = "FAILED"
	VpcConnectionStateRejected     VpcConnectionState = "REJECTED"
	VpcConnectionStateRejecting    VpcConnectionState = "REJECTING"
)

// Values returns all known values for VpcConnectionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VpcConnectionState) Values() []VpcConnectionState {
	return []VpcConnectionState{
		"CREATING",
		"AVAILABLE",
		"INACTIVE",
		"DEACTIVATING",
		"DELETING",
		"FAILED",
		"REJECTED",
		"REJECTING",
	}
}
