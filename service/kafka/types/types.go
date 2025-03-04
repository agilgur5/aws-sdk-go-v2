// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Details of an Amazon MSK Cluster.
type AmazonMskCluster struct {

	// The Amazon Resource Name (ARN) of an Amazon MSK cluster.
	//
	// This member is required.
	MskClusterArn *string

	noSmithyDocumentSerde
}

// Specifies the EBS volume upgrade information. The broker identifier must be set
// to the keyword ALL. This means the changes apply to all the brokers in the
// cluster.
type BrokerEBSVolumeInfo struct {

	// The ID of the broker to update.
	//
	// This member is required.
	KafkaBrokerNodeId *string

	// EBS volume provisioned throughput information.
	ProvisionedThroughput *ProvisionedThroughput

	// Size of the EBS volume to update.
	VolumeSizeGB int32

	noSmithyDocumentSerde
}

type BrokerLogs struct {
	CloudWatchLogs *CloudWatchLogs

	Firehose *Firehose

	S3 *S3

	noSmithyDocumentSerde
}

// Describes the setup to be used for Apache Kafka broker nodes in the cluster.
type BrokerNodeGroupInfo struct {

	// The list of subnets to connect to in the client virtual private cloud (VPC).
	// AWS creates elastic network interfaces inside these subnets. Client applications
	// use elastic network interfaces to produce and consume data. Client subnets can't
	// occupy the Availability Zone with ID use use1-az3.
	//
	// This member is required.
	ClientSubnets []string

	// The type of Amazon EC2 instances to use for Apache Kafka brokers. The following
	// instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge,
	// kafka.m5.4xlarge, kafka.m5.12xlarge, and kafka.m5.24xlarge.
	//
	// This member is required.
	InstanceType *string

	// The distribution of broker nodes across Availability Zones. This is an optional
	// parameter. If you don't specify it, Amazon MSK gives it the value DEFAULT. You
	// can also explicitly set this parameter to the value DEFAULT. No other values are
	// currently allowed. Amazon MSK distributes the broker nodes evenly across the
	// Availability Zones that correspond to the subnets you provide when you create
	// the cluster.
	BrokerAZDistribution BrokerAZDistribution

	// Information about the broker access configuration.
	ConnectivityInfo *ConnectivityInfo

	// The AWS security groups to associate with the elastic network interfaces in
	// order to specify who can connect to and communicate with the Amazon MSK cluster.
	// If you don't specify a security group, Amazon MSK uses the default security
	// group associated with the VPC.
	SecurityGroups []string

	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo *StorageInfo

	// The list of zoneIds for the cluster in the virtual private cloud (VPC).
	ZoneIds []string

	noSmithyDocumentSerde
}

// BrokerNodeInfo
type BrokerNodeInfo struct {

	// The attached elastic network interface of the broker.
	AttachedENIId *string

	// The ID of the broker.
	BrokerId float64

	// The client subnet to which this broker node belongs.
	ClientSubnet *string

	// The virtual private cloud (VPC) of the client.
	ClientVpcIpAddress *string

	// Information about the version of software currently deployed on the Apache
	// Kafka brokers in the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo

	// Endpoints for accessing the broker.
	Endpoints []string

	noSmithyDocumentSerde
}

// Information about the current software installed on the cluster.
type BrokerSoftwareInfo struct {

	// The Amazon Resource Name (ARN) of the configuration used for the cluster. This
	// field isn't visible in this preview release.
	ConfigurationArn *string

	// The revision of the configuration to use. This field isn't visible in this
	// preview release.
	ConfigurationRevision int64

	// The version of Apache Kafka.
	KafkaVersion *string

	noSmithyDocumentSerde
}

// Includes all client authentication information.
type ClientAuthentication struct {

	// Details for ClientAuthentication using SASL.
	Sasl *Sasl

	// Details for ClientAuthentication using TLS.
	Tls *Tls

	// Contains information about unauthenticated traffic to the cluster.
	Unauthenticated *Unauthenticated

	noSmithyDocumentSerde
}

// The client VPC connection object.
type ClientVpcConnection struct {

	// The ARN that identifies the Vpc Connection.
	//
	// This member is required.
	VpcConnectionArn *string

	// Information about the auth scheme of Vpc Connection.
	Authentication *string

	// Creation time of the Vpc Connection.
	CreationTime *time.Time

	// The Owner of the Vpc Connection.
	Owner *string

	// State of the Vpc Connection.
	State VpcConnectionState

	noSmithyDocumentSerde
}

type CloudWatchLogs struct {

	// This member is required.
	Enabled bool

	LogGroup *string

	noSmithyDocumentSerde
}

// Returns information about a cluster.
type Cluster struct {

	// The Amazon Resource Name (ARN) that uniquely identifies a cluster operation.
	ActiveOperationArn *string

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string

	// The name of the cluster.
	ClusterName *string

	// Cluster Type.
	ClusterType ClusterType

	// The time when the cluster was created.
	CreationTime *time.Time

	// The current version of the MSK cluster.
	CurrentVersion *string

	// Information about the provisioned cluster.
	Provisioned *Provisioned

	// Information about the serverless cluster.
	Serverless *Serverless

	// The state of the cluster. The possible states are ACTIVE, CREATING, DELETING,
	// FAILED, HEALING, MAINTENANCE, REBOOTING_BROKER, and UPDATING.
	State ClusterState

	// State Info for the Amazon MSK cluster.
	StateInfo *StateInfo

	// Tags attached to the cluster.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Returns information about a cluster.
type ClusterInfo struct {

	// Arn of active cluster operation.
	ActiveOperationArn *string

	// Information about the broker nodes.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo

	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string

	// The name of the cluster.
	ClusterName *string

	// The time when the cluster was created.
	CreationTime *time.Time

	// Information about the version of software currently deployed on the Apache
	// Kafka brokers in the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo

	// The current version of the MSK cluster.
	CurrentVersion *string

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo

	// Specifies which metrics are gathered for the MSK cluster. This property has the
	// following possible values: DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and
	// PER_TOPIC_PER_PARTITION. For a list of the metrics associated with each of these
	// levels of monitoring, see Monitoring (https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	// .
	EnhancedMonitoring EnhancedMonitoring

	LoggingInfo *LoggingInfo

	// The number of broker nodes in the cluster.
	NumberOfBrokerNodes int32

	// Settings for open monitoring using Prometheus.
	OpenMonitoring *OpenMonitoring

	// The state of the cluster. The possible states are ACTIVE, CREATING, DELETING,
	// FAILED, HEALING, MAINTENANCE, REBOOTING_BROKER, and UPDATING.
	State ClusterState

	StateInfo *StateInfo

	// This controls storage mode for supported storage tiers.
	StorageMode StorageMode

	// Tags attached to the cluster.
	Tags map[string]string

	// The connection string to use to connect to the Apache ZooKeeper cluster.
	ZookeeperConnectString *string

	// The connection string to use to connect to zookeeper cluster on Tls port.
	ZookeeperConnectStringTls *string

	noSmithyDocumentSerde
}

// Returns information about a cluster operation.
type ClusterOperationInfo struct {

	// The ID of the API request that triggered this operation.
	ClientRequestId *string

	// ARN of the cluster.
	ClusterArn *string

	// The time that the operation was created.
	CreationTime *time.Time

	// The time at which the operation finished.
	EndTime *time.Time

	// Describes the error if the operation fails.
	ErrorInfo *ErrorInfo

	// ARN of the cluster operation.
	OperationArn *string

	// State of the cluster operation.
	OperationState *string

	// Steps completed during the operation.
	OperationSteps []ClusterOperationStep

	// Type of the cluster operation.
	OperationType *string

	// Information about cluster attributes before a cluster is updated.
	SourceClusterInfo *MutableClusterInfo

	// Information about cluster attributes after a cluster is updated.
	TargetClusterInfo *MutableClusterInfo

	// Description of the VPC connection for CreateVpcConnection and
	// DeleteVpcConnection operations.
	VpcConnectionInfo *VpcConnectionInfo

	noSmithyDocumentSerde
}

// Step taken during a cluster operation.
type ClusterOperationStep struct {

	// Information about the step and its status.
	StepInfo *ClusterOperationStepInfo

	// The name of the step.
	StepName *string

	noSmithyDocumentSerde
}

// State information about the operation step.
type ClusterOperationStepInfo struct {

	// The steps current status.
	StepStatus *string

	noSmithyDocumentSerde
}

// Returns information about a cluster operation.
type ClusterOperationV2 struct {

	// ARN of the cluster.
	ClusterArn *string

	// Type of the backend cluster.
	ClusterType ClusterType

	// The time at which the operation finished.
	EndTime *time.Time

	// If cluster operation failed from an error, it describes the error.
	ErrorInfo *ErrorInfo

	// ARN of the cluster operation.
	OperationArn *string

	// State of the cluster operation.
	OperationState *string

	// Type of the cluster operation.
	OperationType *string

	// Properties of a provisioned cluster.
	Provisioned *ClusterOperationV2Provisioned

	// Properties of a serverless cluster.
	Serverless *ClusterOperationV2Serverless

	// The time at which operation was started.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Returns information about a provisioned cluster operation.
type ClusterOperationV2Provisioned struct {

	// Steps completed during the operation.
	OperationSteps []ClusterOperationStep

	// Information about cluster attributes before a cluster is updated.
	SourceClusterInfo *MutableClusterInfo

	// Information about cluster attributes after a cluster is updated.
	TargetClusterInfo *MutableClusterInfo

	// Description of the VPC connection for CreateVpcConnection and
	// DeleteVpcConnection operations.
	VpcConnectionInfo *VpcConnectionInfo

	noSmithyDocumentSerde
}

// Returns information about a serverless cluster operation.
type ClusterOperationV2Serverless struct {

	// Description of the VPC connection for CreateVpcConnection and
	// DeleteVpcConnection operations.
	VpcConnectionInfo *VpcConnectionInfoServerless

	noSmithyDocumentSerde
}

// Returns information about a cluster operation.
type ClusterOperationV2Summary struct {

	// ARN of the cluster.
	ClusterArn *string

	// Type of the backend cluster.
	ClusterType ClusterType

	// The time at which the operation finished.
	EndTime *time.Time

	// ARN of the cluster operation.
	OperationArn *string

	// State of the cluster operation.
	OperationState *string

	// Type of the cluster operation.
	OperationType *string

	// The time at which operation was started.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Contains source Apache Kafka versions and compatible target Apache Kafka
// versions.
type CompatibleKafkaVersion struct {

	// An Apache Kafka version.
	SourceVersion *string

	// A list of Apache Kafka versions.
	TargetVersions []string

	noSmithyDocumentSerde
}

// Represents an MSK Configuration.
type Configuration struct {

	// The Amazon Resource Name (ARN) of the configuration.
	//
	// This member is required.
	Arn *string

	// The time when the configuration was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The description of the configuration.
	//
	// This member is required.
	Description *string

	// An array of the versions of Apache Kafka with which you can use this MSK
	// configuration. You can use this configuration for an MSK cluster only if the
	// Apache Kafka version specified for the cluster appears in this array.
	//
	// This member is required.
	KafkaVersions []string

	// Latest revision of the configuration.
	//
	// This member is required.
	LatestRevision *ConfigurationRevision

	// The name of the configuration.
	//
	// This member is required.
	Name *string

	// The state of the configuration. The possible states are ACTIVE, DELETING, and
	// DELETE_FAILED.
	//
	// This member is required.
	State ConfigurationState

	noSmithyDocumentSerde
}

// Specifies the configuration to use for the brokers.
type ConfigurationInfo struct {

	// ARN of the configuration to use.
	//
	// This member is required.
	Arn *string

	// The revision of the configuration to use.
	//
	// This member is required.
	Revision int64

	noSmithyDocumentSerde
}

// Describes a configuration revision.
type ConfigurationRevision struct {

	// The time when the configuration revision was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The revision number.
	//
	// This member is required.
	Revision int64

	// The description of the configuration revision.
	Description *string

	noSmithyDocumentSerde
}

// Information about the broker access configuration.
type ConnectivityInfo struct {

	// Public access control for brokers.
	PublicAccess *PublicAccess

	// VPC connectivity access control for brokers.
	VpcConnectivity *VpcConnectivity

	noSmithyDocumentSerde
}

// Details about consumer group replication.
type ConsumerGroupReplication struct {

	// List of regular expression patterns indicating the consumer groups to copy.
	//
	// This member is required.
	ConsumerGroupsToReplicate []string

	// List of regular expression patterns indicating the consumer groups that should
	// not be replicated.
	ConsumerGroupsToExclude []string

	// Enables synchronization of consumer groups to target cluster.
	DetectAndCopyNewConsumerGroups bool

	// Enables synchronization of consumer group offsets to target cluster. The
	// translated offsets will be written to topic __consumer_offsets.
	SynchroniseConsumerGroupOffsets bool

	noSmithyDocumentSerde
}

// Details about consumer group replication.
type ConsumerGroupReplicationUpdate struct {

	// List of regular expression patterns indicating the consumer groups that should
	// not be replicated.
	//
	// This member is required.
	ConsumerGroupsToExclude []string

	// List of regular expression patterns indicating the consumer groups to copy.
	//
	// This member is required.
	ConsumerGroupsToReplicate []string

	// Enables synchronization of consumer groups to target cluster.
	//
	// This member is required.
	DetectAndCopyNewConsumerGroups bool

	// Enables synchronization of consumer group offsets to target cluster. The
	// translated offsets will be written to topic __consumer_offsets.
	//
	// This member is required.
	SynchroniseConsumerGroupOffsets bool

	noSmithyDocumentSerde
}

// Contains information about the EBS storage volumes attached to Apache Kafka
// broker nodes.
type EBSStorageInfo struct {

	// EBS volume provisioned throughput information.
	ProvisionedThroughput *ProvisionedThroughput

	// The size in GiB of the EBS volume for the data drive on each broker node.
	VolumeSize int32

	noSmithyDocumentSerde
}

// The data-volume encryption details.
type EncryptionAtRest struct {

	// The ARN of the AWS KMS key for encrypting data at rest. If you don't specify a
	// KMS key, MSK creates one for you and uses it.
	//
	// This member is required.
	DataVolumeKMSKeyId *string

	noSmithyDocumentSerde
}

// Includes encryption-related information, such as the AWS KMS key used for
// encrypting data at rest and whether you want MSK to encrypt your data in
// transit.
type EncryptionInfo struct {

	// The data-volume encryption details.
	EncryptionAtRest *EncryptionAtRest

	// The details for encryption in transit.
	EncryptionInTransit *EncryptionInTransit

	noSmithyDocumentSerde
}

// The settings for encrypting data in transit.
type EncryptionInTransit struct {

	// Indicates the encryption setting for data in transit between clients and
	// brokers. The following are the possible values. TLS means that client-broker
	// communication is enabled with TLS only. TLS_PLAINTEXT means that client-broker
	// communication is enabled for both TLS-encrypted, as well as plaintext data.
	// PLAINTEXT means that client-broker communication is enabled in plaintext only.
	// The default value is TLS_PLAINTEXT.
	ClientBroker ClientBroker

	// When set to true, it indicates that data communication among the broker nodes
	// of the cluster is encrypted. When set to false, the communication happens in
	// plaintext. The default value is true.
	InCluster bool

	noSmithyDocumentSerde
}

// Returns information about an error state of the cluster.
type ErrorInfo struct {

	// A number describing the error programmatically.
	ErrorCode *string

	// An optional field to provide more details about the error.
	ErrorString *string

	noSmithyDocumentSerde
}

type Firehose struct {

	// This member is required.
	Enabled bool

	DeliveryStream *string

	noSmithyDocumentSerde
}

// Details for IAM access control.
type Iam struct {

	// Indicates whether IAM access control is enabled.
	Enabled bool

	noSmithyDocumentSerde
}

// Indicates whether you want to turn on or turn off the JMX Exporter.
type JmxExporter struct {

	// Indicates whether you want to turn on or turn off the JMX Exporter.
	//
	// This member is required.
	EnabledInBroker bool

	noSmithyDocumentSerde
}

// Indicates whether you want to turn on or turn off the JMX Exporter.
type JmxExporterInfo struct {

	// Indicates whether you want to turn on or turn off the JMX Exporter.
	//
	// This member is required.
	EnabledInBroker bool

	noSmithyDocumentSerde
}

// Information about Kafka Cluster to be used as source / target for replication.
type KafkaCluster struct {

	// Details of an Amazon MSK Cluster.
	//
	// This member is required.
	AmazonMskCluster *AmazonMskCluster

	// Details of an Amazon VPC which has network connectivity to the Apache Kafka
	// cluster.
	//
	// This member is required.
	VpcConfig *KafkaClusterClientVpcConfig

	noSmithyDocumentSerde
}

// Details of an Amazon VPC which has network connectivity to the Apache Kafka
// cluster.
type KafkaClusterClientVpcConfig struct {

	// The list of subnets in the client VPC to connect to.
	//
	// This member is required.
	SubnetIds []string

	// The security groups to attach to the ENIs for the broker nodes.
	SecurityGroupIds []string

	noSmithyDocumentSerde
}

// Information about Kafka Cluster used as source / target for replication.
type KafkaClusterDescription struct {

	// Details of an Amazon MSK Cluster.
	AmazonMskCluster *AmazonMskCluster

	// The alias of the Kafka cluster. Used to prefix names of replicated topics.
	KafkaClusterAlias *string

	// Details of an Amazon VPC which has network connectivity to the Apache Kafka
	// cluster.
	VpcConfig *KafkaClusterClientVpcConfig

	noSmithyDocumentSerde
}

// Summarized information about Kafka Cluster used as source / target for
// replication.
type KafkaClusterSummary struct {

	// Details of an Amazon MSK Cluster.
	AmazonMskCluster *AmazonMskCluster

	// The alias of the Kafka cluster. Used to prefix names of replicated topics.
	KafkaClusterAlias *string

	noSmithyDocumentSerde
}

type KafkaVersion struct {
	Status KafkaVersionStatus

	Version *string

	noSmithyDocumentSerde
}

type LoggingInfo struct {

	// This member is required.
	BrokerLogs *BrokerLogs

	noSmithyDocumentSerde
}

// Information about cluster attributes that can be updated via update APIs.
type MutableClusterInfo struct {

	// Specifies the size of the EBS volume and the ID of the associated broker.
	BrokerEBSVolumeInfo []BrokerEBSVolumeInfo

	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication

	// Information about the changes in the configuration of the brokers.
	ConfigurationInfo *ConfigurationInfo

	// Information about the broker access configuration.
	ConnectivityInfo *ConnectivityInfo

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo

	// Specifies which Apache Kafka metrics Amazon MSK gathers and sends to Amazon
	// CloudWatch for this cluster.
	EnhancedMonitoring EnhancedMonitoring

	// Information about the Amazon MSK broker type.
	InstanceType *string

	// The Apache Kafka version.
	KafkaVersion *string

	// You can configure your MSK cluster to send broker logs to different destination
	// types. This is a container for the configuration details related to broker logs.
	LoggingInfo *LoggingInfo

	// The number of broker nodes in the cluster.
	NumberOfBrokerNodes int32

	// The settings for open monitoring.
	OpenMonitoring *OpenMonitoring

	// This controls storage mode for supported storage tiers.
	StorageMode StorageMode

	noSmithyDocumentSerde
}

// Indicates whether you want to turn on or turn off the Node Exporter.
type NodeExporter struct {

	// Indicates whether you want to turn on or turn off the Node Exporter.
	//
	// This member is required.
	EnabledInBroker bool

	noSmithyDocumentSerde
}

// Indicates whether you want to turn on or turn off the Node Exporter.
type NodeExporterInfo struct {

	// Indicates whether you want to turn on or turn off the Node Exporter.
	//
	// This member is required.
	EnabledInBroker bool

	noSmithyDocumentSerde
}

// The node information object.
type NodeInfo struct {

	// The start time.
	AddedToClusterTime *string

	// The broker node info.
	BrokerNodeInfo *BrokerNodeInfo

	// The instance type.
	InstanceType *string

	// The Amazon Resource Name (ARN) of the node.
	NodeARN *string

	// The node type.
	NodeType NodeType

	// The ZookeeperNodeInfo.
	ZookeeperNodeInfo *ZookeeperNodeInfo

	noSmithyDocumentSerde
}

// JMX and Node monitoring for the MSK cluster.
type OpenMonitoring struct {

	// Prometheus settings.
	//
	// This member is required.
	Prometheus *Prometheus

	noSmithyDocumentSerde
}

// JMX and Node monitoring for the MSK cluster.
type OpenMonitoringInfo struct {

	// Prometheus settings.
	//
	// This member is required.
	Prometheus *PrometheusInfo

	noSmithyDocumentSerde
}

// Prometheus settings.
type Prometheus struct {

	// Indicates whether you want to turn on or turn off the JMX Exporter.
	JmxExporter *JmxExporter

	// Indicates whether you want to turn on or turn off the Node Exporter.
	NodeExporter *NodeExporter

	noSmithyDocumentSerde
}

// Prometheus settings.
type PrometheusInfo struct {

	// Indicates whether you want to turn on or turn off the JMX Exporter.
	JmxExporter *JmxExporterInfo

	// Indicates whether you want to turn on or turn off the Node Exporter.
	NodeExporter *NodeExporterInfo

	noSmithyDocumentSerde
}

// Provisioned cluster.
type Provisioned struct {

	// Information about the brokers.
	//
	// This member is required.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo

	// The number of broker nodes in the cluster.
	//
	// This member is required.
	NumberOfBrokerNodes int32

	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication

	// Information about the Apache Kafka version deployed on the brokers.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo

	// Specifies the level of monitoring for the MSK cluster. The possible values are
	// DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and PER_TOPIC_PER_PARTITION.
	EnhancedMonitoring EnhancedMonitoring

	// Log delivery information for the cluster.
	LoggingInfo *LoggingInfo

	// The settings for open monitoring.
	OpenMonitoring *OpenMonitoringInfo

	// This controls storage mode for supported storage tiers.
	StorageMode StorageMode

	// The connection string to use to connect to the Apache ZooKeeper cluster.
	ZookeeperConnectString *string

	// The connection string to use to connect to the Apache ZooKeeper cluster on a
	// TLS port.
	ZookeeperConnectStringTls *string

	noSmithyDocumentSerde
}

// Provisioned cluster request.
type ProvisionedRequest struct {

	// Information about the brokers.
	//
	// This member is required.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo

	// The Apache Kafka version that you want for the cluster.
	//
	// This member is required.
	KafkaVersion *string

	// The number of broker nodes in the cluster.
	//
	// This member is required.
	NumberOfBrokerNodes int32

	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication

	// Represents the configuration that you want Amazon MSK to use for the brokers in
	// a cluster.
	ConfigurationInfo *ConfigurationInfo

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo

	// Specifies the level of monitoring for the MSK cluster. The possible values are
	// DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and PER_TOPIC_PER_PARTITION.
	EnhancedMonitoring EnhancedMonitoring

	// Log delivery information for the cluster.
	LoggingInfo *LoggingInfo

	// The settings for open monitoring.
	OpenMonitoring *OpenMonitoringInfo

	// This controls storage mode for supported storage tiers.
	StorageMode StorageMode

	noSmithyDocumentSerde
}

// Contains information about provisioned throughput for EBS storage volumes
// attached to kafka broker nodes.
type ProvisionedThroughput struct {

	// Provisioned throughput is enabled or not.
	Enabled bool

	// Throughput value of the EBS volumes for the data drive on each kafka broker
	// node in MiB per second.
	VolumeThroughput int32

	noSmithyDocumentSerde
}

// Public access control for brokers.
type PublicAccess struct {

	// The value DISABLED indicates that public access is turned off.
	// SERVICE_PROVIDED_EIPS indicates that public access is turned on.
	Type *string

	noSmithyDocumentSerde
}

// Specifies configuration for replication between a source and target Kafka
// cluster.
type ReplicationInfo struct {

	// Configuration relating to consumer group replication.
	//
	// This member is required.
	ConsumerGroupReplication *ConsumerGroupReplication

	// The ARN of the source Kafka cluster.
	//
	// This member is required.
	SourceKafkaClusterArn *string

	// The compression type to use when producing records to target cluster.
	//
	// This member is required.
	TargetCompressionType TargetCompressionType

	// The ARN of the target Kafka cluster.
	//
	// This member is required.
	TargetKafkaClusterArn *string

	// Configuration relating to topic replication.
	//
	// This member is required.
	TopicReplication *TopicReplication

	noSmithyDocumentSerde
}

// Specifies configuration for replication between a source and target Kafka
// cluster (sourceKafkaClusterAlias -> targetKafkaClusterAlias)
type ReplicationInfoDescription struct {

	// Configuration relating to consumer group replication.
	ConsumerGroupReplication *ConsumerGroupReplication

	// The alias of the source Kafka cluster.
	SourceKafkaClusterAlias *string

	// The compression type to use when producing records to target cluster.
	TargetCompressionType TargetCompressionType

	// The alias of the target Kafka cluster.
	TargetKafkaClusterAlias *string

	// Configuration relating to topic replication.
	TopicReplication *TopicReplication

	noSmithyDocumentSerde
}

// Summarized information of replication between clusters.
type ReplicationInfoSummary struct {

	// The alias of the source Kafka cluster.
	SourceKafkaClusterAlias *string

	// The alias of the target Kafka cluster.
	TargetKafkaClusterAlias *string

	noSmithyDocumentSerde
}

// Details about the state of a replicator
type ReplicationStateInfo struct {

	// Code that describes the current state of the replicator.
	Code *string

	// Message that describes the state of the replicator.
	Message *string

	noSmithyDocumentSerde
}

// Information about a replicator.
type ReplicatorSummary struct {

	// The time the replicator was created.
	CreationTime *time.Time

	// The current version of the replicator.
	CurrentVersion *string

	// Whether this resource is a replicator reference.
	IsReplicatorReference bool

	// Kafka Clusters used in setting up sources / targets for replication.
	KafkaClustersSummary []KafkaClusterSummary

	// A list of summarized information of replications between clusters.
	ReplicationInfoSummaryList []ReplicationInfoSummary

	// The Amazon Resource Name (ARN) of the replicator.
	ReplicatorArn *string

	// The name of the replicator.
	ReplicatorName *string

	// The Amazon Resource Name (ARN) of the replicator resource in the region where
	// the replicator was created.
	ReplicatorResourceArn *string

	// State of the replicator.
	ReplicatorState ReplicatorState

	noSmithyDocumentSerde
}

type S3 struct {

	// This member is required.
	Enabled bool

	Bucket *string

	Prefix *string

	noSmithyDocumentSerde
}

// Details for client authentication using SASL.
type Sasl struct {

	// Indicates whether IAM access control is enabled.
	Iam *Iam

	// Details for SASL/SCRAM client authentication.
	Scram *Scram

	noSmithyDocumentSerde
}

// Details for SASL/SCRAM client authentication.
type Scram struct {

	// SASL/SCRAM authentication is enabled or not.
	Enabled bool

	noSmithyDocumentSerde
}

// Serverless cluster.
type Serverless struct {

	// The configuration of the Amazon VPCs for the cluster.
	//
	// This member is required.
	VpcConfigs []VpcConfig

	// Includes all client authentication information.
	ClientAuthentication *ServerlessClientAuthentication

	noSmithyDocumentSerde
}

// Includes all client authentication information.
type ServerlessClientAuthentication struct {

	// Details for ClientAuthentication using SASL.
	Sasl *ServerlessSasl

	noSmithyDocumentSerde
}

// Serverless cluster request.
type ServerlessRequest struct {

	// The configuration of the Amazon VPCs for the cluster.
	//
	// This member is required.
	VpcConfigs []VpcConfig

	// Includes all client authentication information.
	ClientAuthentication *ServerlessClientAuthentication

	noSmithyDocumentSerde
}

// Details for client authentication using SASL.
type ServerlessSasl struct {

	// Indicates whether IAM access control is enabled.
	Iam *Iam

	noSmithyDocumentSerde
}

type StateInfo struct {
	Code *string

	Message *string

	noSmithyDocumentSerde
}

// Contains information about storage volumes attached to MSK broker nodes.
type StorageInfo struct {

	// EBS volume information.
	EbsStorageInfo *EBSStorageInfo

	noSmithyDocumentSerde
}

// Details for client authentication using TLS.
type Tls struct {

	// List of ACM Certificate Authority ARNs.
	CertificateAuthorityArnList []string

	// Specifies whether you want to turn on or turn off TLS authentication.
	Enabled bool

	noSmithyDocumentSerde
}

// Details about topic replication.
type TopicReplication struct {

	// List of regular expression patterns indicating the topics to copy.
	//
	// This member is required.
	TopicsToReplicate []string

	// Whether to periodically configure remote topic ACLs to match their
	// corresponding upstream topics.
	CopyAccessControlListsForTopics bool

	// Whether to periodically configure remote topics to match their corresponding
	// upstream topics.
	CopyTopicConfigurations bool

	// Whether to periodically check for new topics and partitions.
	DetectAndCopyNewTopics bool

	// List of regular expression patterns indicating the topics that should not be
	// replicated.
	TopicsToExclude []string

	noSmithyDocumentSerde
}

// Details for updating the topic replication of a replicator.
type TopicReplicationUpdate struct {

	// Whether to periodically configure remote topic ACLs to match their
	// corresponding upstream topics.
	//
	// This member is required.
	CopyAccessControlListsForTopics bool

	// Whether to periodically configure remote topics to match their corresponding
	// upstream topics.
	//
	// This member is required.
	CopyTopicConfigurations bool

	// Whether to periodically check for new topics and partitions.
	//
	// This member is required.
	DetectAndCopyNewTopics bool

	// List of regular expression patterns indicating the topics that should not be
	// replicated.
	//
	// This member is required.
	TopicsToExclude []string

	// List of regular expression patterns indicating the topics to copy.
	//
	// This member is required.
	TopicsToReplicate []string

	noSmithyDocumentSerde
}

type Unauthenticated struct {

	// Specifies whether you want to turn on or turn off unauthenticated traffic to
	// your cluster.
	Enabled bool

	noSmithyDocumentSerde
}

// Error info for scram secret associate/disassociate failure.
type UnprocessedScramSecret struct {

	// Error code for associate/disassociate failure.
	ErrorCode *string

	// Error message for associate/disassociate failure.
	ErrorMessage *string

	// AWS Secrets Manager secret ARN.
	SecretArn *string

	noSmithyDocumentSerde
}

// Description of the requester that calls the API operation.
type UserIdentity struct {

	// A unique identifier for the requester that calls the API operation.
	PrincipalId *string

	// The identity type of the requester that calls the API operation.
	Type UserIdentityType

	noSmithyDocumentSerde
}

// The configuration of the Amazon VPCs for the cluster.
type VpcConfig struct {

	// The IDs of the subnets associated with the cluster.
	//
	// This member is required.
	SubnetIds []string

	// The IDs of the security groups associated with the cluster.
	SecurityGroupIds []string

	noSmithyDocumentSerde
}

// The VPC connection object.
type VpcConnection struct {

	// The ARN that identifies the Cluster which the Vpc Connection belongs to.
	//
	// This member is required.
	TargetClusterArn *string

	// The ARN that identifies the Vpc Connection.
	//
	// This member is required.
	VpcConnectionArn *string

	// Information about the auth scheme of Vpc Connection.
	Authentication *string

	// Creation time of the Vpc Connection.
	CreationTime *time.Time

	// State of the Vpc Connection.
	State VpcConnectionState

	// The vpcId that belongs to the Vpc Connection.
	VpcId *string

	noSmithyDocumentSerde
}

// Description of the VPC connection.
type VpcConnectionInfo struct {

	// The time when Amazon MSK creates the VPC Connnection.
	CreationTime *time.Time

	// The owner of the VPC Connection.
	Owner *string

	// Description of the requester that calls the API operation.
	UserIdentity *UserIdentity

	// The Amazon Resource Name (ARN) of the VPC connection.
	VpcConnectionArn *string

	noSmithyDocumentSerde
}

// Description of the VPC connection.
type VpcConnectionInfoServerless struct {

	// The time when Amazon MSK creates the VPC Connnection.
	CreationTime *time.Time

	// The owner of the VPC Connection.
	Owner *string

	// Description of the requester that calls the API operation.
	UserIdentity *UserIdentity

	// The Amazon Resource Name (ARN) of the VPC connection.
	VpcConnectionArn *string

	noSmithyDocumentSerde
}

// VPC connectivity access control for brokers.
type VpcConnectivity struct {

	// Includes all client authentication information for VPC connectivity.
	ClientAuthentication *VpcConnectivityClientAuthentication

	noSmithyDocumentSerde
}

// Includes all client authentication information for VPC connectivity.
type VpcConnectivityClientAuthentication struct {

	// SASL authentication type details for VPC connectivity.
	Sasl *VpcConnectivitySasl

	// TLS authentication type details for VPC connectivity.
	Tls *VpcConnectivityTls

	noSmithyDocumentSerde
}

// Details for IAM access control for VPC connectivity.
type VpcConnectivityIam struct {

	// SASL/IAM authentication is on or off for VPC connectivity.
	Enabled bool

	noSmithyDocumentSerde
}

// Details for SASL client authentication for VPC connectivity.
type VpcConnectivitySasl struct {

	// Details for SASL/IAM client authentication for VPC connectivity.
	Iam *VpcConnectivityIam

	// Details for SASL/SCRAM client authentication for VPC connectivity.
	Scram *VpcConnectivityScram

	noSmithyDocumentSerde
}

// Details for SASL/SCRAM client authentication for VPC connectivity.
type VpcConnectivityScram struct {

	// SASL/SCRAM authentication is on or off for VPC connectivity.
	Enabled bool

	noSmithyDocumentSerde
}

// Details for TLS client authentication for VPC connectivity.
type VpcConnectivityTls struct {

	// TLS authentication is on or off for VPC connectivity.
	Enabled bool

	noSmithyDocumentSerde
}

// Zookeeper node information.
type ZookeeperNodeInfo struct {

	// The attached elastic network interface of the broker.
	AttachedENIId *string

	// The virtual private cloud (VPC) IP address of the client.
	ClientVpcIpAddress *string

	// Endpoints for accessing the ZooKeeper.
	Endpoints []string

	// The role-specific ID for Zookeeper.
	ZookeeperId float64

	// The version of Zookeeper.
	ZookeeperVersion *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
