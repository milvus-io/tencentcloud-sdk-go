// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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

package v20190725

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AddNodeList struct {
	// 需要删除的节点角色。
	// - SECONDARY：Mongod 节点。
	// - READONLY：只读节点。
	// - MONGOS：Mongos 节点。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 节点所对应的可用区。
	// - 单可用区，所有节点在同一可用区。
	// - 多可用区：当前标准规格是三可用区分布，主从节点不在同一可用区，需注意配置新增节点对应的可用区，且新增后必须满足任意2个可用区节点数大于第3个可用区原则。
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

// Predefined struct for user
type AssignProjectRequestParams struct {
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 项目ID，用户已创建项目的唯一ID,非自定义
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type AssignProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 项目ID，用户已创建项目的唯一ID,非自定义
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *AssignProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignProjectResponseParams struct {
	// 返回的异步任务ID列表
	FlowIds []*uint64 `json:"FlowIds,omitnil" name:"FlowIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssignProjectResponse struct {
	*tchttp.BaseResponse
	Response *AssignProjectResponseParams `json:"Response"`
}

func (r *AssignProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Auth struct {
	// 当前账号具有的权限信息。<ul><li>0：无权限。</li><li>1：只读。</li><li>2：只写。</li><li>3：读写。</li></ul>
	Mask *int64 `json:"Mask,omitnil" name:"Mask"`

	// 指具有当前账号权限的数据库名。
	// <ul><li>* ：表示所有数据库。</li><li>db.name：表示特定name的数据库。</li></ul>
	NameSpace *string `json:"NameSpace,omitnil" name:"NameSpace"`
}

type BackupDownloadTask struct {
	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 备份文件名
	BackupName *string `json:"BackupName,omitnil" name:"BackupName"`

	// 分片名称
	ReplicaSetId *string `json:"ReplicaSetId,omitnil" name:"ReplicaSetId"`

	// 备份数据大小，单位为字节
	BackupSize *int64 `json:"BackupSize,omitnil" name:"BackupSize"`

	// 任务状态。0-等待执行，1-正在下载，2-下载完成，3-下载失败，4-等待重试
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 任务进度百分比
	Percent *int64 `json:"Percent,omitnil" name:"Percent"`

	// 耗时，单位为秒
	TimeSpend *int64 `json:"TimeSpend,omitnil" name:"TimeSpend"`

	// 备份数据下载链接
	Url *string `json:"Url,omitnil" name:"Url"`

	// 备份文件备份类型，0-逻辑备份，1-物理备份
	BackupMethod *int64 `json:"BackupMethod,omitnil" name:"BackupMethod"`

	// 发起备份时指定的备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupDesc *string `json:"BackupDesc,omitnil" name:"BackupDesc"`
}

type BackupDownloadTaskStatus struct {
	// 分片名
	ReplicaSetId *string `json:"ReplicaSetId,omitnil" name:"ReplicaSetId"`

	// 任务当前状态。0-等待执行，1-正在下载，2-下载完成，3-下载失败，4-等待重试
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type BackupInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 备份方式，0-自动备份，1-手动备份
	BackupType *uint64 `json:"BackupType,omitnil" name:"BackupType"`

	// 备份名称
	BackupName *string `json:"BackupName,omitnil" name:"BackupName"`

	// 备份备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupDesc *string `json:"BackupDesc,omitnil" name:"BackupDesc"`

	// 备份文件大小，单位KB
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSize *uint64 `json:"BackupSize,omitnil" name:"BackupSize"`

	// 备份开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 备份结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 备份状态，1-备份中，2-备份成功
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 备份方法，0-逻辑备份，1-物理备份
	BackupMethod *uint64 `json:"BackupMethod,omitnil" name:"BackupMethod"`
}

type ClientConnection struct {
	// 连接的客户端IP
	IP *string `json:"IP,omitnil" name:"IP"`

	// 对应客户端IP的连接数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 是否为内部ip
	InternalService *bool `json:"InternalService,omitnil" name:"InternalService"`
}

// Predefined struct for user
type CreateAccountUserRequestParams struct {
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 新账号名称。其格式要求如下：<ul><li>字符范围[1,32]。</li><li>可输入[A,Z]、[a,z]、[1,9]范围的字符以及下划线“_”与短划线“-”。</li></ul>
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 新账号密码。密码复杂度要求如下：<ul><li>字符长度范围[8,32]。</li><li>至少包含字母、数字和特殊字符（叹号“!”、at"@"、井号“#”、百分号“%”、插入符“^”、星号“*”、小括号“()”、下划线“_”）中的两种。</li></ul>
	Password *string `json:"Password,omitnil" name:"Password"`

	// mongouser 账号对应的密码。mongouser 为系统默认账号，即为创建实例时，设置的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil" name:"MongoUserPassword"`

	// 账号备注信息。
	UserDesc *string `json:"UserDesc,omitnil" name:"UserDesc"`

	// 账号的读写权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil" name:"AuthRole"`
}

type CreateAccountUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 新账号名称。其格式要求如下：<ul><li>字符范围[1,32]。</li><li>可输入[A,Z]、[a,z]、[1,9]范围的字符以及下划线“_”与短划线“-”。</li></ul>
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 新账号密码。密码复杂度要求如下：<ul><li>字符长度范围[8,32]。</li><li>至少包含字母、数字和特殊字符（叹号“!”、at"@"、井号“#”、百分号“%”、插入符“^”、星号“*”、小括号“()”、下划线“_”）中的两种。</li></ul>
	Password *string `json:"Password,omitnil" name:"Password"`

	// mongouser 账号对应的密码。mongouser 为系统默认账号，即为创建实例时，设置的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil" name:"MongoUserPassword"`

	// 账号备注信息。
	UserDesc *string `json:"UserDesc,omitnil" name:"UserDesc"`

	// 账号的读写权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil" name:"AuthRole"`
}

func (r *CreateAccountUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "MongoUserPassword")
	delete(f, "UserDesc")
	delete(f, "AuthRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountUserResponseParams struct {
	// 创建任务ID。
	FlowId *uint64 `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAccountUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountUserResponseParams `json:"Response"`
}

func (r *CreateAccountUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDBInstanceRequestParams struct {
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 设置备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil" name:"BackupMethod"`

	// 备份备注信息。
	BackupRemark *string `json:"BackupRemark,omitnil" name:"BackupRemark"`
}

type CreateBackupDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 设置备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil" name:"BackupMethod"`

	// 备份备注信息。
	BackupRemark *string `json:"BackupRemark,omitnil" name:"BackupRemark"`
}

func (r *CreateBackupDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDBInstanceResponseParams struct {
	// 查询备份流程的状态。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBackupDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupDBInstanceResponseParams `json:"Response"`
}

func (r *CreateBackupDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDownloadTaskRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 要下载的备份文件名，可通过DescribeDBBackups接口获取。
	BackupName *string `json:"BackupName,omitnil" name:"BackupName"`

	// 指定要下载的副本集的节点名称 或 分片集群的分片名称列表。
	// 如副本集cmgo-p8vnipr5，示例(固定取值)：BackupSets.0=cmgo-p8vnipr5_0，可下载全量数据。
	// 如分片集群cmgo-p8vnipr5，示例：BackupSets.0=cmgo-p8vnipr5_0&BackupSets.1=cmgo-p8vnipr5_1，即下载分片0和分片1的数据，分片集群如需全量下载，请按示例方式传入全部分片名称。
	BackupSets []*ReplicaSetInfo `json:"BackupSets,omitnil" name:"BackupSets"`
}

type CreateBackupDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 要下载的备份文件名，可通过DescribeDBBackups接口获取。
	BackupName *string `json:"BackupName,omitnil" name:"BackupName"`

	// 指定要下载的副本集的节点名称 或 分片集群的分片名称列表。
	// 如副本集cmgo-p8vnipr5，示例(固定取值)：BackupSets.0=cmgo-p8vnipr5_0，可下载全量数据。
	// 如分片集群cmgo-p8vnipr5，示例：BackupSets.0=cmgo-p8vnipr5_0&BackupSets.1=cmgo-p8vnipr5_1，即下载分片0和分片1的数据，分片集群如需全量下载，请按示例方式传入全部分片名称。
	BackupSets []*ReplicaSetInfo `json:"BackupSets,omitnil" name:"BackupSets"`
}

func (r *CreateBackupDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "BackupSets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDownloadTaskResponseParams struct {
	// 下载任务状态
	Tasks []*BackupDownloadTaskStatus `json:"Tasks,omitnil" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBackupDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupDownloadTaskResponseParams `json:"Response"`
}

func (r *CreateBackupDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourRequestParams struct {
	// 实例内存大小，单位：GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// 指副本集数量。
	// - 创建副本集实例，该参数只能为1。
	// - 创建分片实例，指分片的数量。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// 指每个副本集内节点个数。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 指版本信息。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 机器类型。
	// - HIO：高IO型。
	// - HIO10G：高IO万兆。
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 实例数量，最小值1，最大值为10。
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 私有网络ID。如果不设置该参数，则默认选择基础网络。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络下的子网 ID，如果配置参数 VpcId，则 SubnetId必须配置。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例密码。自定义密码长度为8-32个字符，至少包含字母、数字和字符（!@#%^*()_）中的两种。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 项目ID。若不设置该参数，则为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	Clone *int64 `json:"Clone,omitnil" name:"Clone"`

	// 父实例 ID。当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	Father *string `json:"Father,omitnil" name:"Father"`

	// 安全组。
	SecurityGroup []*string `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 克隆实例回档时间。
	// - 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。
	// - 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil" name:"RestoreTime"`

	// 实例名称。仅支持长度为60个字符的中文、英文、数字、下划线_、分隔符- 。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 多可用区部署的节点列表。具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)获取。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	// - 不支持4.2及以上版本。
	// - 不支持只读灾备实例。
	// - 不能选择基础网络。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil" name:"AvailabilityZoneList"`

	// Mongos CPU 核数。购买分片集群时，必须填写。
	MongosCpu *uint64 `json:"MongosCpu,omitnil" name:"MongosCpu"`

	// Mongos 内存大小。购买分片集群时，必须填写。
	// 
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// Mongos 数量。购买分片集群时，必须填写。
	// 
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil" name:"MongosNodeNum"`

	// 只读节点数量，最大不超过7个。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil" name:"HiddenZone"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// 实例内存大小，单位：GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// 指副本集数量。
	// - 创建副本集实例，该参数只能为1。
	// - 创建分片实例，指分片的数量。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// 指每个副本集内节点个数。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 指版本信息。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 机器类型。
	// - HIO：高IO型。
	// - HIO10G：高IO万兆。
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 实例数量，最小值1，最大值为10。
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 私有网络ID。如果不设置该参数，则默认选择基础网络。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络下的子网 ID，如果配置参数 VpcId，则 SubnetId必须配置。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例密码。自定义密码长度为8-32个字符，至少包含字母、数字和字符（!@#%^*()_）中的两种。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 项目ID。若不设置该参数，则为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	Clone *int64 `json:"Clone,omitnil" name:"Clone"`

	// 父实例 ID。当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	Father *string `json:"Father,omitnil" name:"Father"`

	// 安全组。
	SecurityGroup []*string `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 克隆实例回档时间。
	// - 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。
	// - 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil" name:"RestoreTime"`

	// 实例名称。仅支持长度为60个字符的中文、英文、数字、下划线_、分隔符- 。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 多可用区部署的节点列表。具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)获取。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	// - 不支持4.2及以上版本。
	// - 不支持只读灾备实例。
	// - 不能选择基础网络。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil" name:"AvailabilityZoneList"`

	// Mongos CPU 核数。购买分片集群时，必须填写。
	MongosCpu *uint64 `json:"MongosCpu,omitnil" name:"MongosCpu"`

	// Mongos 内存大小。购买分片集群时，必须填写。
	// 
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// Mongos 数量。购买分片集群时，必须填写。
	// 
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil" name:"MongosNodeNum"`

	// 只读节点数量，最大不超过7个。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil" name:"HiddenZone"`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ReplicateSetNum")
	delete(f, "NodeNum")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "ClusterType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "ProjectId")
	delete(f, "Tags")
	delete(f, "Clone")
	delete(f, "Father")
	delete(f, "SecurityGroup")
	delete(f, "RestoreTime")
	delete(f, "InstanceName")
	delete(f, "AvailabilityZoneList")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNodeNum")
	delete(f, "ReadonlyNodeNum")
	delete(f, "ReadonlyNodeAvailabilityZoneList")
	delete(f, "HiddenZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// 订单ID。
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// 创建的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceHourResponseParams `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
	// 指每个副本集内节点个数。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例内存大小，单位：GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// 指版本信息。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 实例数量, 最小值1，最大值为10。
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 实例时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// 机器类型。
	// - HIO：高IO型。
	// - HIO10G：高IO万兆。
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 指副本集数量。
	// - 创建副本集实例，该参数只能为1。
	// - 创建分片实例，指分片的数量。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// 项目ID。若不设置该参数，则为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 私有网络ID。如果不设置该参数，则默认选择基础网络。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络下的子网 ID，如果配置参数 VpcId，则 SubnetId必须配置。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例密码。自定义密码长度为8-32个字符，至少包含字母、数字和字符（!@#%^*()_）中的两种。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 自动续费标记。
	// - 0：不自动续费。默认为不自动续费。
	// - 1：自动续费。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 是否自动选择代金券。
	// - 1：是。
	// - 0：否。默认为0。
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	Clone *int64 `json:"Clone,omitnil" name:"Clone"`

	// 父实例 ID。当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	Father *string `json:"Father,omitnil" name:"Father"`

	// 安全组。
	SecurityGroup []*string `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 克隆实例回档时间。
	// - 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。
	// - 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil" name:"RestoreTime"`

	// 实例名称。仅支持长度为60个字符的中文、英文、数字、下划线_、分隔符- 。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 多可用区部署的节点列表。具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)获取。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	// - 不支持4.2及以上版本。
	// - 不支持只读灾备实例。
	// - 不能选择基础网络。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil" name:"AvailabilityZoneList"`

	// Mongos CPU 核数，购买MongoDB 4.2 及以上WiredTiger存储引擎版本的分片集群时，必须填写。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	MongosCpu *uint64 `json:"MongosCpu,omitnil" name:"MongosCpu"`

	// Mongos 内存大小。购买MongoDB 4.2 及以上WiredTiger存储引擎版本的分片集群时，必须填写。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// Mongos 数量。购买MongoDB 4.2 及以上WiredTiger存储引擎版本的分片集群时，必须填写。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。为了保障高可用，取值范围为[3,32]。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil" name:"MongosNodeNum"`

	// 只读节点数量，取值范围[0,5]。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil" name:"HiddenZone"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指每个副本集内节点个数。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例内存大小，单位：GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例硬盘大小，单位：GB。
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// 指版本信息。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 实例数量, 最小值1，最大值为10。
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 实例时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// 机器类型。
	// - HIO：高IO型。
	// - HIO10G：高IO万兆。
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 指副本集数量。
	// - 创建副本集实例，该参数只能为1。
	// - 创建分片实例，指分片的数量。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// 项目ID。若不设置该参数，则为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 私有网络ID。如果不设置该参数，则默认选择基础网络。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络下的子网 ID，如果配置参数 VpcId，则 SubnetId必须配置。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例密码。自定义密码长度为8-32个字符，至少包含字母、数字和字符（!@#%^*()_）中的两种。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 自动续费标记。
	// - 0：不自动续费。默认为不自动续费。
	// - 1：自动续费。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 是否自动选择代金券。
	// - 1：是。
	// - 0：否。默认为0。
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	Clone *int64 `json:"Clone,omitnil" name:"Clone"`

	// 父实例 ID。当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	Father *string `json:"Father,omitnil" name:"Father"`

	// 安全组。
	SecurityGroup []*string `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 克隆实例回档时间。
	// - 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。
	// - 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil" name:"RestoreTime"`

	// 实例名称。仅支持长度为60个字符的中文、英文、数字、下划线_、分隔符- 。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 多可用区部署的节点列表。具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)获取。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	// - 不支持4.2及以上版本。
	// - 不支持只读灾备实例。
	// - 不能选择基础网络。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil" name:"AvailabilityZoneList"`

	// Mongos CPU 核数，购买MongoDB 4.2 及以上WiredTiger存储引擎版本的分片集群时，必须填写。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	MongosCpu *uint64 `json:"MongosCpu,omitnil" name:"MongosCpu"`

	// Mongos 内存大小。购买MongoDB 4.2 及以上WiredTiger存储引擎版本的分片集群时，必须填写。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// Mongos 数量。购买MongoDB 4.2 及以上WiredTiger存储引擎版本的分片集群时，必须填写。具体售卖规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。为了保障高可用，取值范围为[3,32]。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil" name:"MongosNodeNum"`

	// 只读节点数量，取值范围[0,5]。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil" name:"HiddenZone"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "Period")
	delete(f, "MachineCode")
	delete(f, "ClusterType")
	delete(f, "ReplicateSetNum")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "Tags")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "Clone")
	delete(f, "Father")
	delete(f, "SecurityGroup")
	delete(f, "RestoreTime")
	delete(f, "InstanceName")
	delete(f, "AvailabilityZoneList")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNodeNum")
	delete(f, "ReadonlyNodeNum")
	delete(f, "ReadonlyNodeAvailabilityZoneList")
	delete(f, "HiddenZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// 创建的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CurrentOp struct {
	// 操作序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpId *int64 `json:"OpId,omitnil" name:"OpId"`

	// 操作所在的命名空间，形式如db.collection
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ns *string `json:"Ns,omitnil" name:"Ns"`

	// 操作执行语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Query *string `json:"Query,omitnil" name:"Query"`

	// 操作类型，可能的取值：aggregate、count、delete、distinct、find、findAndModify、getMore、insert、mapReduce、update和command
	// 注意：此字段可能返回 null，表示取不到有效值。
	Op *string `json:"Op,omitnil" name:"Op"`

	// 操作所在的分片名称
	ReplicaSetName *string `json:"ReplicaSetName,omitnil" name:"ReplicaSetName"`

	// 筛选条件，节点状态，可能的取值为：Primary、Secondary
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil" name:"State"`

	// 操作详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// 操作所在的节点名称
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 操作已执行时间（ms）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicrosecsRunning *uint64 `json:"MicrosecsRunning,omitnil" name:"MicrosecsRunning"`
}

type DBInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitnil" name:"Region"`
}

type DBInstancePrice struct {
	// 单价
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil" name:"UnitPrice"`

	// 原价
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// 折扣价
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`
}

// Predefined struct for user
type DeleteAccountUserRequestParams struct {
	// 指定待删除账号的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 配置待删除的账号名。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 配置 mongouser 对应的密码。mongouser为系统默认账号，输入其对应的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil" name:"MongoUserPassword"`
}

type DeleteAccountUserRequest struct {
	*tchttp.BaseRequest
	
	// 指定待删除账号的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 配置待删除的账号名。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 配置 mongouser 对应的密码。mongouser为系统默认账号，输入其对应的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil" name:"MongoUserPassword"`
}

func (r *DeleteAccountUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "MongoUserPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountUserResponseParams struct {
	// 账户删除任务ID。
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAccountUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountUserResponseParams `json:"Response"`
}

func (r *DeleteAccountUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountUsersRequestParams struct {
	// 指定待获取账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeAccountUsersRequest struct {
	*tchttp.BaseRequest
	
	// 指定待获取账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeAccountUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountUsersResponseParams struct {
	// 实例账号列表。
	Users []*UserInfo `json:"Users,omitnil" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountUsersResponseParams `json:"Response"`
}

func (r *DescribeAccountUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// 异步请求Id，涉及到异步流程的接口返回，如CreateBackupDBInstance
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// 异步请求Id，涉及到异步流程的接口返回，如CreateBackupDBInstance
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// 状态。返回参数有：initial-初始化、running-运行中、paused-任务执行失败，已暂停、undoed-任务执行失败，已回滚、failed-任务执行失败, 已终止、success-成功
	Status *string `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadTaskRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 备份文件名，用来过滤指定文件的下载任务
	BackupName *string `json:"BackupName,omitnil" name:"BackupName"`

	// 指定查询时间范围内的任务，StartTime指定开始时间，不填默认不限制开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 指定查询时间范围内的任务，EndTime指定截止时间，不填默认不限制截止时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 此次查询返回的条数，取值范围为1-100，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定此次查询返回的页数，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序字段，取值为createTime，finishTime两种，默认为createTime
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 排序方式，取值为asc，desc两种，默认desc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// 根据任务状态过滤。0-等待执行，1-正在下载，2-下载完成，3-下载失败，4-等待重试。不填默认返回所有类型
	Status []*int64 `json:"Status,omitnil" name:"Status"`
}

type DescribeBackupDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 备份文件名，用来过滤指定文件的下载任务
	BackupName *string `json:"BackupName,omitnil" name:"BackupName"`

	// 指定查询时间范围内的任务，StartTime指定开始时间，不填默认不限制开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 指定查询时间范围内的任务，EndTime指定截止时间，不填默认不限制截止时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 此次查询返回的条数，取值范围为1-100，默认为20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定此次查询返回的页数，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序字段，取值为createTime，finishTime两种，默认为createTime
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 排序方式，取值为asc，desc两种，默认desc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// 根据任务状态过滤。0-等待执行，1-正在下载，2-下载完成，3-下载失败，4-等待重试。不填默认返回所有类型
	Status []*int64 `json:"Status,omitnil" name:"Status"`
}

func (r *DescribeBackupDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadTaskResponseParams struct {
	// 满足查询条件的所有条数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 下载任务列表
	Tasks []*BackupDownloadTask `json:"Tasks,omitnil" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackupDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadTaskResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsRequestParams struct {
	// 指定待查询的实例ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 单次请求返回的数量。最小值为1，最大值为1000，默认值为1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认值为0。Offset=Limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeClientConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 指定待查询的实例ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 单次请求返回的数量。最小值为1，最大值为1000，默认值为1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认值为0。Offset=Limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeClientConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsResponseParams struct {
	// 客户端连接信息，包括客户端 IP 和对应 IP 的连接数量。
	Clients []*ClientConnection `json:"Clients,omitnil" name:"Clients"`

	// 满足条件的记录总条数，可用于分页查询。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClientConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientConnectionsResponseParams `json:"Response"`
}

func (r *DescribeClientConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentOpRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 筛选条件，操作所属的命名空间namespace，格式为db.collection
	Ns *string `json:"Ns,omitnil" name:"Ns"`

	// 筛选条件，操作已经执行的时间（单位：毫秒），结果将返回超过设置时间的操作，默认值为0，取值范围为[0, 3600000]
	MillisecondRunning *uint64 `json:"MillisecondRunning,omitnil" name:"MillisecondRunning"`

	// 筛选条件，操作类型，可能的取值：none，update，insert，query，command，getmore，remove和killcursors
	Op *string `json:"Op,omitnil" name:"Op"`

	// 筛选条件，分片名称
	ReplicaSetName *string `json:"ReplicaSetName,omitnil" name:"ReplicaSetName"`

	// 筛选条件，节点状态，可能的取值为：primary
	// secondary
	State *string `json:"State,omitnil" name:"State"`

	// 单次请求返回的数量，默认值为100，取值范围为[0,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认值为0，取值范围为[0,10000]
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回结果集排序的字段，目前支持："MicrosecsRunning"/"microsecsrunning"，默认为升序排序
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回结果集排序方式，可能的取值："ASC"/"asc"或"DESC"/"desc"
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`
}

type DescribeCurrentOpRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 筛选条件，操作所属的命名空间namespace，格式为db.collection
	Ns *string `json:"Ns,omitnil" name:"Ns"`

	// 筛选条件，操作已经执行的时间（单位：毫秒），结果将返回超过设置时间的操作，默认值为0，取值范围为[0, 3600000]
	MillisecondRunning *uint64 `json:"MillisecondRunning,omitnil" name:"MillisecondRunning"`

	// 筛选条件，操作类型，可能的取值：none，update，insert，query，command，getmore，remove和killcursors
	Op *string `json:"Op,omitnil" name:"Op"`

	// 筛选条件，分片名称
	ReplicaSetName *string `json:"ReplicaSetName,omitnil" name:"ReplicaSetName"`

	// 筛选条件，节点状态，可能的取值为：primary
	// secondary
	State *string `json:"State,omitnil" name:"State"`

	// 单次请求返回的数量，默认值为100，取值范围为[0,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认值为0，取值范围为[0,10000]
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回结果集排序的字段，目前支持："MicrosecsRunning"/"microsecsrunning"，默认为升序排序
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 返回结果集排序方式，可能的取值："ASC"/"asc"或"DESC"/"desc"
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`
}

func (r *DescribeCurrentOpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentOpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ns")
	delete(f, "MillisecondRunning")
	delete(f, "Op")
	delete(f, "ReplicaSetName")
	delete(f, "State")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurrentOpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentOpResponseParams struct {
	// 符合查询条件的操作总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 当前操作列表
	CurrentOps []*CurrentOp `json:"CurrentOps,omitnil" name:"CurrentOps"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCurrentOpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurrentOpResponseParams `json:"Response"`
}

func (r *DescribeCurrentOpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentOpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 备份方式，当前支持：0-逻辑备份，1-物理备份，2-所有备份。默认为逻辑备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil" name:"BackupMethod"`

	// 分页大小，最大值为100，不设置默认查询所有。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，最小值为0，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 备份方式，当前支持：0-逻辑备份，1-物理备份，2-所有备份。默认为逻辑备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil" name:"BackupMethod"`

	// 分页大小，最大值为100，不设置默认查询所有。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，最小值为0，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsResponseParams struct {
	// 备份列表
	BackupList []*BackupInfo `json:"BackupList,omitnil" name:"BackupList"`

	// 备份总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBBackupsResponseParams `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDealRequestParams struct {
	// 订单ID，通过CreateDBInstance等接口返回
	DealId *string `json:"DealId,omitnil" name:"DealId"`
}

type DescribeDBInstanceDealRequest struct {
	*tchttp.BaseRequest
	
	// 订单ID，通过CreateDBInstance等接口返回
	DealId *string `json:"DealId,omitnil" name:"DealId"`
}

func (r *DescribeDBInstanceDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceDealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDealResponseParams struct {
	// 订单状态，1：未支付，2：已支付，3：发货中，4：发货成功，5：发货失败，6：退款，7：订单关闭，8：超时未支付关闭。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 订单原价。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// 订单折扣价格。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`

	// 订单行为，purchase：新购，renew：续费，upgrade：升配，downgrade：降配，refund：退货退款。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBInstanceDealResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceDealResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceNodePropertyRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点ID。
	NodeIds []*string `json:"NodeIds,omitnil" name:"NodeIds"`

	// 节点角色。可选值包括：
	// <ul><li>PRIMARY：主节点。</li><li>SECONDARY：从节点。</li><li>READONLY：只读节点。</li><li>ARBITER：仲裁节点。</li></ul>
	Roles []*string `json:"Roles,omitnil" name:"Roles"`

	// 该参数指定节点是否为Hidden节点，默认为false。
	OnlyHidden *bool `json:"OnlyHidden,omitnil" name:"OnlyHidden"`

	// 该参数指定选举新主节点的优先级。其取值范围为[0,100]，数值越高，优先级越高。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 该参数指定节点投票权。
	// <ul><li>1：具有投票权。</li><li>0：无投票权。</li></ul>
	Votes *int64 `json:"Votes,omitnil" name:"Votes"`

	// 节点标签。
	Tags []*NodeTag `json:"Tags,omitnil" name:"Tags"`
}

type DescribeDBInstanceNodePropertyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点ID。
	NodeIds []*string `json:"NodeIds,omitnil" name:"NodeIds"`

	// 节点角色。可选值包括：
	// <ul><li>PRIMARY：主节点。</li><li>SECONDARY：从节点。</li><li>READONLY：只读节点。</li><li>ARBITER：仲裁节点。</li></ul>
	Roles []*string `json:"Roles,omitnil" name:"Roles"`

	// 该参数指定节点是否为Hidden节点，默认为false。
	OnlyHidden *bool `json:"OnlyHidden,omitnil" name:"OnlyHidden"`

	// 该参数指定选举新主节点的优先级。其取值范围为[0,100]，数值越高，优先级越高。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 该参数指定节点投票权。
	// <ul><li>1：具有投票权。</li><li>0：无投票权。</li></ul>
	Votes *int64 `json:"Votes,omitnil" name:"Votes"`

	// 节点标签。
	Tags []*NodeTag `json:"Tags,omitnil" name:"Tags"`
}

func (r *DescribeDBInstanceNodePropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceNodePropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeIds")
	delete(f, "Roles")
	delete(f, "OnlyHidden")
	delete(f, "Priority")
	delete(f, "Votes")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceNodePropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceNodePropertyResponseParams struct {
	// Mongos节点属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mongos []*NodeProperty `json:"Mongos,omitnil" name:"Mongos"`

	// 副本集节点信息。
	ReplicateSets []*ReplicateSetInfo `json:"ReplicateSets,omitnil" name:"ReplicateSets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBInstanceNodePropertyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceNodePropertyResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceNodePropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceNodePropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 实例 ID 列表。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 实例类型。取值范围如下：<ul><li>0：所有实例。</li><li>1：正式实例。</li><li>2：临时实例。</li><li>3：只读实例。</li><li>-1：正式实例、只读、灾备实例。</li></ul>
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// 集群类型，取值范围如下：<ul><li>0：副本集实例。</li><li>1：分片实例。</li><li>-1：所有实例。</li></ul>
	ClusterType *int64 `json:"ClusterType,omitnil" name:"ClusterType"`

	// 实例状态，取值范围如下所示：<ul><li>0：待初始化。</li><li>1：流程执行中。</li><li>2：实例有效。</li><li>-2：已隔离（包年包月实例）。</li><li>-3：已隔离（按量计费实例）。</li></ul>
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// 私有网络的 ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其 ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络的子网ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其子网 ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 付费类型，取值范围如下：<ul><li>0：查询按量计费实例。</li><li>1：查询包年包月实例。</li><li>-1：查询按量计费与包年包月实例。</li></ul>
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 单次请求返回的数量。默认值为20，取值范围为[1,100]。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 配置返回结果排序依据的字段。目前支持依据"ProjectId"、"InstanceName"、"CreateTime"排序。
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 配置返回结果的排序方式。
	// - ASC：升序。
	// - DESC：降序。
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// 项目 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在右上角的账户信息下拉菜单中，选择项目管理查询项目。
	ProjectIds []*uint64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// 配置查询搜索的关键词。支持配置为实例ID、实例名称或者内网 IP 地址。
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// 标签信息，包含标签键与标签值。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 实例类型。取值范围如下：<ul><li>0：所有实例。</li><li>1：正式实例。</li><li>2：临时实例。</li><li>3：只读实例。</li><li>-1：正式实例、只读、灾备实例。</li></ul>
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// 集群类型，取值范围如下：<ul><li>0：副本集实例。</li><li>1：分片实例。</li><li>-1：所有实例。</li></ul>
	ClusterType *int64 `json:"ClusterType,omitnil" name:"ClusterType"`

	// 实例状态，取值范围如下所示：<ul><li>0：待初始化。</li><li>1：流程执行中。</li><li>2：实例有效。</li><li>-2：已隔离（包年包月实例）。</li><li>-3：已隔离（按量计费实例）。</li></ul>
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// 私有网络的 ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其 ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络的子网ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其子网 ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 付费类型，取值范围如下：<ul><li>0：查询按量计费实例。</li><li>1：查询包年包月实例。</li><li>-1：查询按量计费与包年包月实例。</li></ul>
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 单次请求返回的数量。默认值为20，取值范围为[1,100]。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 配置返回结果排序依据的字段。目前支持依据"ProjectId"、"InstanceName"、"CreateTime"排序。
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 配置返回结果的排序方式。
	// - ASC：升序。
	// - DESC：降序。
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// 项目 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在右上角的账户信息下拉菜单中，选择项目管理查询项目。
	ProjectIds []*uint64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// 配置查询搜索的关键词。支持配置为实例ID、实例名称或者内网 IP 地址。
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// 标签信息，包含标签键与标签值。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ClusterType")
	delete(f, "Status")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "ProjectIds")
	delete(f, "SearchKey")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例详细信息列表。
	InstanceDetails []*InstanceDetail `json:"InstanceDetails,omitnil" name:"InstanceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// 指定待查询参数列表的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 指定待查询参数列表的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 参数值为枚举类型参数集合。
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitnil" name:"InstanceEnumParam"`

	// 参数值为 Integer 类型参数集合。
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitnil" name:"InstanceIntegerParam"`

	// 参数值为 Text 类型的参数集合。
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitnil" name:"InstanceTextParam"`

	// 参数值为混合类型的参数集合。
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitnil" name:"InstanceMultiParam"`

	// 当前实例支持修改的参数个数统计 如0
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRequestParams struct {
	// 实例 ID。例如：cmgo-p8vn****。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。例如：cmgo-p8vn****。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupResponseParams struct {
	// 实例绑定的安全组信息。
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogPatternsRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 慢日志返回格式，可设置为json，不传默认返回原生慢日志格式。
	Format *string `json:"Format,omitnil" name:"Format"`
}

type DescribeSlowLogPatternsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 慢日志返回格式，可设置为json，不传默认返回原生慢日志格式。
	Format *string `json:"Format,omitnil" name:"Format"`
}

func (r *DescribeSlowLogPatternsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogPatternsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogPatternsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogPatternsResponseParams struct {
	// 慢日志统计信息总数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 慢日志统计信息
	SlowLogPatterns []*SlowLogPattern `json:"SlowLogPatterns,omitnil" name:"SlowLogPatterns"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSlowLogPatternsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogPatternsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogPatternsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogPatternsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 慢日志返回格式。默认返回原生慢日志格式，4.4及以上版本可设置为json。
	Format *string `json:"Format,omitnil" name:"Format"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 慢日志返回格式。默认返回原生慢日志格式，4.4及以上版本可设置为json。
	Format *string `json:"Format,omitnil" name:"Format"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// 慢日志总数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 慢日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowLogs []*string `json:"SlowLogs,omitnil" name:"SlowLogs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoRequestParams struct {
	// 待查询可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest
	
	// 待查询可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

func (r *DescribeSpecInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoResponseParams struct {
	// 实例售卖规格信息列表
	SpecInfoList []*SpecificationInfo `json:"SpecInfoList,omitnil" name:"SpecInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpecInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecInfoResponseParams `json:"Response"`
}

func (r *DescribeSpecInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlushInstanceRouterConfigRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type FlushInstanceRouterConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *FlushInstanceRouterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushInstanceRouterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlushInstanceRouterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlushInstanceRouterConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FlushInstanceRouterConfigResponse struct {
	*tchttp.BaseResponse
	Response *FlushInstanceRouterConfigResponseParams `json:"Response"`
}

func (r *FlushInstanceRouterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushInstanceRouterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDBInstancesRequestParams struct {
	// 实例所属区域及可用区信息。格式：ap-guangzhou-2。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 每个分片的主从节点数量。
	// 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinNodeNum与MaxNodeNum分别对应其最小值与最大值。
	NodeNum *int64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例内存大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数CPU与Memory分别对应CPU核数与内存规格。
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 实例硬盘大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinStorage与MaxStorage分别对应其最小磁盘规格与最大磁盘规格。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// 实例版本信息。具体支持的版本，请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MongoVersionCode为实例所支持的版本信息。版本信息与版本号对应关系如下：
	// - MONGO_3_WT：MongoDB 3.2 WiredTiger存储引擎版本。
	// - MONGO_3_ROCKS：MongoDB 3.2 RocksDB存储引擎版本。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 机器类型。
	// - HIO：高IO型。
	// - HIO10G：高IO万兆型。
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 实例数量，取值范围为[1,10]。
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 实例类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	// - STANDALONE：单节点。
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 副本集个数。
	// - 创建副本集实例时，该参数固定设置为1。
	// - 创建分片集群时，指分片数量，请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	// - 若为单节点实例，该参数固定设置为0。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// - 选择包年包月计费模式，即 <b>InstanceChargeType </b>设定为<b>PREPAID</b>时，需设定购买实例的时长。该参数取值可选：[1,2,3,4,5,6,7,8,9,10,11,12,24,36]；单位：月。
	// -选择按量计费，即 <b>InstanceChargeType</b> 设定为 **POSTPAID_BY_HOUR** 时，该参数仅可配置为 1。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 实例付费方式。
	// - PREPAID：包年包月计费。
	// - POSTPAID_BY_HOUR：按量计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 分片实例询价必填参数，指 Mongos CPU核数，取值范围为[1,16]。
	MongosCpu *uint64 `json:"MongosCpu,omitnil" name:"MongosCpu"`

	// 分片实例询价必填参数，指 Mongos 内存，取值范围为[2,32]，单位：GB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// 分片实例询价必填参数，指 Mongos 个数，取值范围为[3,32]。
	MongosNum *uint64 `json:"MongosNum,omitnil" name:"MongosNum"`

	// 分片实例询价必填参数，指 ConfigServer CPU核数，取值为1，单位：GB。
	ConfigServerCpu *uint64 `json:"ConfigServerCpu,omitnil" name:"ConfigServerCpu"`

	// 分片实例询价必填参数，指 ConfigServer 内存大小，取值为2，单位：GB。
	ConfigServerMemory *uint64 `json:"ConfigServerMemory,omitnil" name:"ConfigServerMemory"`

	// 分片实例询价必填参数，指 ConfigServer 磁盘大小，取值为 20，单位：GB。
	ConfigServerVolume *uint64 `json:"ConfigServerVolume,omitnil" name:"ConfigServerVolume"`
}

type InquirePriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属区域及可用区信息。格式：ap-guangzhou-2。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 每个分片的主从节点数量。
	// 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinNodeNum与MaxNodeNum分别对应其最小值与最大值。
	NodeNum *int64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例内存大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数CPU与Memory分别对应CPU核数与内存规格。
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 实例硬盘大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinStorage与MaxStorage分别对应其最小磁盘规格与最大磁盘规格。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// 实例版本信息。具体支持的版本，请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MongoVersionCode为实例所支持的版本信息。版本信息与版本号对应关系如下：
	// - MONGO_3_WT：MongoDB 3.2 WiredTiger存储引擎版本。
	// - MONGO_3_ROCKS：MongoDB 3.2 RocksDB存储引擎版本。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 机器类型。
	// - HIO：高IO型。
	// - HIO10G：高IO万兆型。
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 实例数量，取值范围为[1,10]。
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 实例类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	// - STANDALONE：单节点。
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 副本集个数。
	// - 创建副本集实例时，该参数固定设置为1。
	// - 创建分片集群时，指分片数量，请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	// - 若为单节点实例，该参数固定设置为0。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// - 选择包年包月计费模式，即 <b>InstanceChargeType </b>设定为<b>PREPAID</b>时，需设定购买实例的时长。该参数取值可选：[1,2,3,4,5,6,7,8,9,10,11,12,24,36]；单位：月。
	// -选择按量计费，即 <b>InstanceChargeType</b> 设定为 **POSTPAID_BY_HOUR** 时，该参数仅可配置为 1。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 实例付费方式。
	// - PREPAID：包年包月计费。
	// - POSTPAID_BY_HOUR：按量计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 分片实例询价必填参数，指 Mongos CPU核数，取值范围为[1,16]。
	MongosCpu *uint64 `json:"MongosCpu,omitnil" name:"MongosCpu"`

	// 分片实例询价必填参数，指 Mongos 内存，取值范围为[2,32]，单位：GB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// 分片实例询价必填参数，指 Mongos 个数，取值范围为[3,32]。
	MongosNum *uint64 `json:"MongosNum,omitnil" name:"MongosNum"`

	// 分片实例询价必填参数，指 ConfigServer CPU核数，取值为1，单位：GB。
	ConfigServerCpu *uint64 `json:"ConfigServerCpu,omitnil" name:"ConfigServerCpu"`

	// 分片实例询价必填参数，指 ConfigServer 内存大小，取值为2，单位：GB。
	ConfigServerMemory *uint64 `json:"ConfigServerMemory,omitnil" name:"ConfigServerMemory"`

	// 分片实例询价必填参数，指 ConfigServer 磁盘大小，取值为 20，单位：GB。
	ConfigServerVolume *uint64 `json:"ConfigServerVolume,omitnil" name:"ConfigServerVolume"`
}

func (r *InquirePriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "NodeNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "ClusterType")
	delete(f, "ReplicateSetNum")
	delete(f, "Period")
	delete(f, "InstanceChargeType")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNum")
	delete(f, "ConfigServerCpu")
	delete(f, "ConfigServerMemory")
	delete(f, "ConfigServerVolume")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDBInstancesResponseParams struct {
	// 价格
	Price *DBInstancePrice `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDBInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDBInstanceSpecRequestParams struct {
	// 实例 ID，格式如：cmgo-p8vn****。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 变更配置后实例内存大小，单位：GB。
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 变更配置后实例磁盘大小，单位：GB。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// 实例节点数。默认为不变更节点数，暂不支持变更。
	NodeNum *int64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例分片数。默认为不变更分片数，暂不支持变更。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`
}

type InquirePriceModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cmgo-p8vn****。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 变更配置后实例内存大小，单位：GB。
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 变更配置后实例磁盘大小，单位：GB。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// 实例节点数。默认为不变更节点数，暂不支持变更。
	NodeNum *int64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例分片数。默认为不变更分片数，暂不支持变更。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`
}

func (r *InquirePriceModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "NodeNum")
	delete(f, "ReplicateSetNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDBInstanceSpecResponseParams struct {
	// 价格。
	Price *DBInstancePrice `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDBInstancesRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同，接口单次最多只支持5个实例进行操作。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 预付费模式（即包年包月）相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

type InquirePriceRenewDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同，接口单次最多只支持5个实例进行操作。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 预付费模式（即包年包月）相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDBInstancesResponseParams struct {
	// 价格
	Price *DBInstancePrice `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceRenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewDBInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceRenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	// （InquirePriceRenewDBInstances，RenewDBInstances调用时必填）
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 自动续费标识。取值范围：
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	// （InquirePriceRenewDBInstances，RenewDBInstances调用时必填）
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type InstanceDetail struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 付费类型，可能的返回值：1-包年包月；0-按量计费
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 集群类型，可能的返回值：0-副本集实例，1-分片实例。
	ClusterType *uint64 `json:"ClusterType,omitnil" name:"ClusterType"`

	// 地域信息。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 可用区信息。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 网络类型，可能的返回值：0-基础网络，1-私有网络
	NetType *uint64 `json:"NetType,omitnil" name:"NetType"`

	// 私有网络的ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络的子网ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例状态，可能的返回值：0-待初始化，1-流程处理中，2-运行中，-2-实例已过期。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 实例IP。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 端口号。
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`

	// 实例创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 实例到期时间。
	DeadLine *string `json:"DeadLine,omitnil" name:"DeadLine"`

	// 实例版本信息。
	MongoVersion *string `json:"MongoVersion,omitnil" name:"MongoVersion"`

	// 实例内存规格，单位为MB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例磁盘规格，单位为MB。
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// 实例CPU核心数。
	CpuNum *uint64 `json:"CpuNum,omitnil" name:"CpuNum"`

	// 实例机器类型。
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`

	// 实例从节点数。
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil" name:"SecondaryNum"`

	// 实例分片数。
	ReplicationSetNum *uint64 `json:"ReplicationSetNum,omitnil" name:"ReplicationSetNum"`

	// 实例自动续费标志，可能的返回值：0-手动续费，1-自动续费，2-确认不续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 已用容量，单位MB。
	UsedVolume *uint64 `json:"UsedVolume,omitnil" name:"UsedVolume"`

	// 维护窗口起始时间。
	MaintenanceStart *string `json:"MaintenanceStart,omitnil" name:"MaintenanceStart"`

	// 维护窗口结束时间。
	MaintenanceEnd *string `json:"MaintenanceEnd,omitnil" name:"MaintenanceEnd"`

	// 分片信息。
	ReplicaSets []*ShardInfo `json:"ReplicaSets,omitnil" name:"ReplicaSets"`

	// 只读实例信息。
	ReadonlyInstances []*DBInstanceInfo `json:"ReadonlyInstances,omitnil" name:"ReadonlyInstances"`

	// 灾备实例信息。
	StandbyInstances []*DBInstanceInfo `json:"StandbyInstances,omitnil" name:"StandbyInstances"`

	// 临时实例信息。
	CloneInstances []*DBInstanceInfo `json:"CloneInstances,omitnil" name:"CloneInstances"`

	// 关联实例信息，对于正式实例，该字段表示它的临时实例信息；对于临时实例，则表示它的正式实例信息;如果为只读/灾备实例,则表示他的主实例信息。
	RelatedInstance *DBInstanceInfo `json:"RelatedInstance,omitnil" name:"RelatedInstance"`

	// 实例标签信息集合。
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 实例版本标记。
	InstanceVer *uint64 `json:"InstanceVer,omitnil" name:"InstanceVer"`

	// 实例版本标记。
	ClusterVer *uint64 `json:"ClusterVer,omitnil" name:"ClusterVer"`

	// 协议信息，可能的返回值：1-mongodb，2-dynamodb。
	Protocol *uint64 `json:"Protocol,omitnil" name:"Protocol"`

	// 实例类型，可能的返回值，1-正式实例，2-临时实例，3-只读实例，4-灾备实例
	InstanceType *uint64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例状态描述
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil" name:"InstanceStatusDesc"`

	// 实例对应的物理实例id，回档并替换过的实例有不同的InstanceId和RealInstanceId，从barad获取监控数据等场景下需要用物理id获取
	RealInstanceId *string `json:"RealInstanceId,omitnil" name:"RealInstanceId"`

	// 实例当前可用区信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneList []*string `json:"ZoneList,omitnil" name:"ZoneList"`

	// mongos节点个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil" name:"MongosNodeNum"`

	// mongos节点内存。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongosMemory *uint64 `json:"MongosMemory,omitnil" name:"MongosMemory"`

	// mongos节点CPU核数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongosCpuNum *uint64 `json:"MongosCpuNum,omitnil" name:"MongosCpuNum"`

	// Config Server节点个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigServerNodeNum *uint64 `json:"ConfigServerNodeNum,omitnil" name:"ConfigServerNodeNum"`

	// Config Server节点内存。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigServerMemory *uint64 `json:"ConfigServerMemory,omitnil" name:"ConfigServerMemory"`

	// Config Server节点磁盘大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigServerVolume *uint64 `json:"ConfigServerVolume,omitnil" name:"ConfigServerVolume"`

	// Config Server节点CPU核数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigServerCpuNum *uint64 `json:"ConfigServerCpuNum,omitnil" name:"ConfigServerCpuNum"`

	// readonly节点个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil" name:"ReadonlyNodeNum"`
}

type InstanceEnumParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// 枚举值，所有支持的值。
	EnumValue []*string `json:"EnumValue,omitnil" name:"EnumValue"`

	// 参数修改之后是否需要重启生效。
	// - 1：需要重启后生效。
	// - 0：无需重启，设置成功即可生效。
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil" name:"Tips"`

	// 参数值类型说明。
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// 是否为运行中参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type InstanceIntegerParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// 参数最大值。
	Max *string `json:"Max,omitnil" name:"Max"`

	// 最小值。
	Min *string `json:"Min,omitnil" name:"Min"`

	// 参数修改之后是否需要重启生效。
	// - 1:需要重启后生效。
	// - 0：无需重启，设置成功即可生效。
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil" name:"Tips"`

	// 参数类型。
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// 是否为运行中参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 冗余字段，可忽略。
	Unit *string `json:"Unit,omitnil" name:"Unit"`
}

type InstanceMultiParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// 参考值范围。
	EnumValue []*string `json:"EnumValue,omitnil" name:"EnumValue"`

	// 参数修改后是否需要重启才会生效。
	// - 1：需要重启后生效。
	// - 0：无需重启，设置成功即可生效。
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// 是否为运行中参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil" name:"Tips"`

	// 当前值的类型描述，默认为multi。
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`
}

type InstanceTextParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// 修改参数值之后是否需要重启。
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// Text 类型参数对应的值。
	TextValue *string `json:"TextValue,omitnil" name:"TextValue"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil" name:"Tips"`

	// 参数值类型说明。
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// 是否为运行中的参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillOpsRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 待终止的操作
	Operations []*Operation `json:"Operations,omitnil" name:"Operations"`
}

type KillOpsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 待终止的操作
	Operations []*Operation `json:"Operations,omitnil" name:"Operations"`
}

func (r *KillOpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillOpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Operations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillOpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillOpsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type KillOpsResponse struct {
	*tchttp.BaseResponse
	Response *KillOpsResponseParams `json:"Response"`
}

func (r *KillOpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillOpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkAddressRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 原IP保留时长，单位为分钟；原IP会在约定时间后释放，在释放前原IP和新IP均可访问；0表示立即回收原IP
	OldIpExpiredTime *uint64 `json:"OldIpExpiredTime,omitnil" name:"OldIpExpiredTime"`

	// 切换后IP地址的归属私有网络统一ID，若为基础网络，该字段为空
	NewUniqVpcId *string `json:"NewUniqVpcId,omitnil" name:"NewUniqVpcId"`

	// 切换后IP地址的归属子网统一ID，若为基础网络，该字段为空
	NewUniqSubnetId *string `json:"NewUniqSubnetId,omitnil" name:"NewUniqSubnetId"`

	// 待修改IP信息
	NetworkAddresses []*ModifyNetworkAddress `json:"NetworkAddresses,omitnil" name:"NetworkAddresses"`
}

type ModifyDBInstanceNetworkAddressRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 原IP保留时长，单位为分钟；原IP会在约定时间后释放，在释放前原IP和新IP均可访问；0表示立即回收原IP
	OldIpExpiredTime *uint64 `json:"OldIpExpiredTime,omitnil" name:"OldIpExpiredTime"`

	// 切换后IP地址的归属私有网络统一ID，若为基础网络，该字段为空
	NewUniqVpcId *string `json:"NewUniqVpcId,omitnil" name:"NewUniqVpcId"`

	// 切换后IP地址的归属子网统一ID，若为基础网络，该字段为空
	NewUniqSubnetId *string `json:"NewUniqSubnetId,omitnil" name:"NewUniqSubnetId"`

	// 待修改IP信息
	NetworkAddresses []*ModifyNetworkAddress `json:"NetworkAddresses,omitnil" name:"NetworkAddresses"`
}

func (r *ModifyDBInstanceNetworkAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldIpExpiredTime")
	delete(f, "NewUniqVpcId")
	delete(f, "NewUniqSubnetId")
	delete(f, "NetworkAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNetworkAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstanceNetworkAddressResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNetworkAddressResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNetworkAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupRequestParams struct {
	// 实例 ID。例如：cmgo-7pje****。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 目标安全组 ID。请通过接口[DescribeSecurityGroup](https://cloud.tencent.com/document/product/240/55675)查看具体的安全组 ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。例如：cmgo-7pje****。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 目标安全组 ID。请通过接口[DescribeSecurityGroup](https://cloud.tencent.com/document/product/240/55675)查看具体的安全组 ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *ModifyDBInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecRequestParams struct {
	// 实例 ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例配置变更后的内存大小。
	// - 单位：GB。
	// - 内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例配置变更后的硬盘大小，单位：GB。<ul><li>内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。</li><li>降配时，变更后的磁盘容量必须大于已用磁盘容量的1.2倍。</li></ul>
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// (已废弃) 请使用ResizeOplog独立接口完成。
	// 
	// 实例配置变更后 Oplog 的大小。
	// - 单位：GB。
	// - 默认 Oplog 占用容量为磁盘空间的10%。系统允许设置的 Oplog 容量范围为磁盘空间的[10%,90%]。
	OplogSize *uint64 `json:"OplogSize,omitnil" name:"OplogSize"`

	// 实例变更后的节点数(mongod节点或mongos节点或readonly节点调整后的节点数，具体类型取决于AddNodeList或RemoveNodeList参数的节点类型)。
	// <ul><li>副本集：取值范围请通过云数据库的售卖规格（DescribeSpecInfo）接口返回的参数 MinNodeNum 与 MaxNodeNum 获取。</li><li>分片集群：取值范围请通过云数据库的售卖规格（DescribeSpecInfo）接口返回的参数 MinReplicateSetNodeNum 与 MaxReplicateSetNodeNum 获取。</li></ul>
	NodeNum *uint64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例变更后的分片数。<ul><li>取值范围请通过云数据库的售卖规格（DescribeSpecInfo）接口返回的参数MinReplicateSetNum与MaxReplicateSetNum获取。</li><li>该参数只能增加不能减少。</li></ul>
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// 实例配置变更的切换时间。<ul><li>0：调整完成时，立即执行变配任务。默认为0。</li><li>1：在维护时间窗内，执行变配任务。
	// <b>说明</b>：调整节点数和分片数不支持在<b>维护时间窗内</b>变更。</li></ul>
	InMaintenance *uint64 `json:"InMaintenance,omitnil" name:"InMaintenance"`

	// 新增节点属性列表。
	AddNodeList []*AddNodeList `json:"AddNodeList,omitnil" name:"AddNodeList"`

	// 删除节点属性列表。
	RemoveNodeList []*RemoveNodeList `json:"RemoveNodeList,omitnil" name:"RemoveNodeList"`
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例配置变更后的内存大小。
	// - 单位：GB。
	// - 内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 实例配置变更后的硬盘大小，单位：GB。<ul><li>内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。</li><li>降配时，变更后的磁盘容量必须大于已用磁盘容量的1.2倍。</li></ul>
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// (已废弃) 请使用ResizeOplog独立接口完成。
	// 
	// 实例配置变更后 Oplog 的大小。
	// - 单位：GB。
	// - 默认 Oplog 占用容量为磁盘空间的10%。系统允许设置的 Oplog 容量范围为磁盘空间的[10%,90%]。
	OplogSize *uint64 `json:"OplogSize,omitnil" name:"OplogSize"`

	// 实例变更后的节点数(mongod节点或mongos节点或readonly节点调整后的节点数，具体类型取决于AddNodeList或RemoveNodeList参数的节点类型)。
	// <ul><li>副本集：取值范围请通过云数据库的售卖规格（DescribeSpecInfo）接口返回的参数 MinNodeNum 与 MaxNodeNum 获取。</li><li>分片集群：取值范围请通过云数据库的售卖规格（DescribeSpecInfo）接口返回的参数 MinReplicateSetNodeNum 与 MaxReplicateSetNodeNum 获取。</li></ul>
	NodeNum *uint64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// 实例变更后的分片数。<ul><li>取值范围请通过云数据库的售卖规格（DescribeSpecInfo）接口返回的参数MinReplicateSetNum与MaxReplicateSetNum获取。</li><li>该参数只能增加不能减少。</li></ul>
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil" name:"ReplicateSetNum"`

	// 实例配置变更的切换时间。<ul><li>0：调整完成时，立即执行变配任务。默认为0。</li><li>1：在维护时间窗内，执行变配任务。
	// <b>说明</b>：调整节点数和分片数不支持在<b>维护时间窗内</b>变更。</li></ul>
	InMaintenance *uint64 `json:"InMaintenance,omitnil" name:"InMaintenance"`

	// 新增节点属性列表。
	AddNodeList []*AddNodeList `json:"AddNodeList,omitnil" name:"AddNodeList"`

	// 删除节点属性列表。
	RemoveNodeList []*RemoveNodeList `json:"RemoveNodeList,omitnil" name:"RemoveNodeList"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "OplogSize")
	delete(f, "NodeNum")
	delete(f, "ReplicateSetNum")
	delete(f, "InMaintenance")
	delete(f, "AddNodeList")
	delete(f, "RemoveNodeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecResponseParams struct {
	// 订单 ID。
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAddress struct {
	// 新IP地址。
	NewIPAddress *string `json:"NewIPAddress,omitnil" name:"NewIPAddress"`

	// 原IP地址。
	OldIpAddress *string `json:"OldIpAddress,omitnil" name:"OldIpAddress"`
}

type NodeProperty struct {
	// 节点所在的可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 节点名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 节点访问地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil" name:"Address"`

	// 角色。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 是否为Hidden节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hidden *bool `json:"Hidden,omitnil" name:"Hidden"`

	// 节点状态，包括：ORMAL/STARTUP/RECOVERING/STARTUP2/UNKNOWN/DOWN/ROLLBACK/REMOVED等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 主从延迟，单位秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveDelay *int64 `json:"SlaveDelay,omitnil" name:"SlaveDelay"`

	// 节点优先级。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 节点投票权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Votes *int64 `json:"Votes,omitnil" name:"Votes"`

	// 节点标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*NodeTag `json:"Tags,omitnil" name:"Tags"`

	// 副本集Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicateSetId *string `json:"ReplicateSetId,omitnil" name:"ReplicateSetId"`
}

type NodeTag struct {
	// 节点Tag key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 节点Tag Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

// Predefined struct for user
type OfflineIsolatedDBInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type OfflineIsolatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *OfflineIsolatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedDBInstanceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OfflineIsolatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedDBInstanceResponseParams `json:"Response"`
}

func (r *OfflineIsolatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Operation struct {
	// 操作所在的分片名
	ReplicaSetName *string `json:"ReplicaSetName,omitnil" name:"ReplicaSetName"`

	// 操作所在的节点名
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 操作序号
	OpId *int64 `json:"OpId,omitnil" name:"OpId"`
}

type RemoveNodeList struct {
	// 需要删除的节点角色。
	// - SECONDARY：Mongod 节点。
	// - READONLY：只读节点。
	// - MONGOS：Mongos 节点。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 要删除的节点 ID。分片集群须指定一组分片要删除的节点名称即可，其余分片对改组对齐。
	// 
	// - 获取方式：登录 [MongoDB控制台](https://console.cloud.tencent.com/)，在**节点管理**页签，可获取**节点 ID**。
	// - 特别说明：分片集群同一节点上的分片，仅需指定0分片节点 ID 即可。例如：cmgo-6hfk****_0-node-primary。
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// 节点所对应的可用区。
	// - 单可用区，所有节点在同一可用区。
	// - 多可用区：当前标准规格是三可用区分布，主从节点不在同一可用区，需注意配置所删除节点对应的可用区，且删除后必须满足任意2个可用区节点数大于第3个可用区原则。
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

// Predefined struct for user
type RenameInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 自定义实例名称，名称只支持长度为60个字符的中文、英文、数字、下划线_、分隔符 -
	NewName *string `json:"NewName,omitnil" name:"NewName"`
}

type RenameInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 自定义实例名称，名称只支持长度为60个字符的中文、英文、数字、下划线_、分隔符 -
	NewName *string `json:"NewName,omitnil" name:"NewName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenameInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenameInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenameInstanceResponseParams `json:"Response"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

type RenewDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

func (r *RenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBInstancesResponseParams `json:"Response"`
}

func (r *RenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicaSetInfo struct {
	// 副本集ID
	ReplicaSetId *string `json:"ReplicaSetId,omitnil" name:"ReplicaSetId"`
}

type ReplicateSetInfo struct {
	// 节点属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*NodeProperty `json:"Nodes,omitnil" name:"Nodes"`
}

// Predefined struct for user
type ResetDBInstancePasswordRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例账号名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 新密码，新密码长度不能少于8位
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ResetDBInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例账号名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 新密码，新密码长度不能少于8位
	Password *string `json:"Password,omitnil" name:"Password"`
}

func (r *ResetDBInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDBInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDBInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDBInstancePasswordResponseParams struct {
	// 异步请求Id，用户查询该流程的运行状态
	AsyncRequestId *string `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetDBInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetDBInstancePasswordResponseParams `json:"Response"`
}

func (r *ResetDBInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDBInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 所属项目id
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 入站规则
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil" name:"Inbound"`

	// 出站规则
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil" name:"Outbound"`

	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitnil" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil" name:"SecurityGroupRemark"`
}

type SecurityGroupBound struct {
	// 执行规则。ACCEPT或DROP
	Action *string `json:"Action,omitnil" name:"Action"`

	// ip段。
	CidrIp *string `json:"CidrIp,omitnil" name:"CidrIp"`

	// 端口范围
	PortRange *string `json:"PortRange,omitnil" name:"PortRange"`

	// 传输层协议。tcp，udp或ALL
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 安全组id代表的地址集合
	Id *string `json:"Id,omitnil" name:"Id"`

	// 地址组id代表的地址集合
	AddressModule *string `json:"AddressModule,omitnil" name:"AddressModule"`

	// 服务组id代表的协议和端口集合
	ServiceModule *string `json:"ServiceModule,omitnil" name:"ServiceModule"`

	// 描述
	Desc *string `json:"Desc,omitnil" name:"Desc"`
}

// Predefined struct for user
type SetAccountUserPrivilegeRequestParams struct {
	// 指定待设置账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 设置账号名称。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 设置权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil" name:"AuthRole"`
}

type SetAccountUserPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// 指定待设置账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 设置账号名称。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 设置权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil" name:"AuthRole"`
}

func (r *SetAccountUserPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAccountUserPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "AuthRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAccountUserPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAccountUserPrivilegeResponseParams struct {
	// 任务ID。
	FlowId *uint64 `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetAccountUserPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *SetAccountUserPrivilegeResponseParams `json:"Response"`
}

func (r *SetAccountUserPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAccountUserPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {
	// 分片已使用容量
	UsedVolume *float64 `json:"UsedVolume,omitnil" name:"UsedVolume"`

	// 分片ID
	ReplicaSetId *string `json:"ReplicaSetId,omitnil" name:"ReplicaSetId"`

	// 分片名
	ReplicaSetName *string `json:"ReplicaSetName,omitnil" name:"ReplicaSetName"`

	// 分片内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 分片磁盘规格，单位为MB
	Volume *uint64 `json:"Volume,omitnil" name:"Volume"`

	// 分片Oplog大小，单位为MB
	OplogSize *uint64 `json:"OplogSize,omitnil" name:"OplogSize"`

	// 分片从节点数
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil" name:"SecondaryNum"`

	// 分片物理id
	RealReplicaSetId *string `json:"RealReplicaSetId,omitnil" name:"RealReplicaSetId"`
}

type SlowLogPattern struct {
	// 慢日志模式
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`

	// 最大执行时间
	MaxTime *uint64 `json:"MaxTime,omitnil" name:"MaxTime"`

	// 平均执行时间
	AverageTime *uint64 `json:"AverageTime,omitnil" name:"AverageTime"`

	// 该模式慢日志条数
	Total *uint64 `json:"Total,omitnil" name:"Total"`
}

type SpecItem struct {
	// 规格信息标识
	SpecCode *string `json:"SpecCode,omitnil" name:"SpecCode"`

	// 规格有效标志，取值：0-停止售卖，1-开放售卖
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 计算资源规格，单位为CPU核心数
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 默认磁盘规格，单位MB
	DefaultStorage *uint64 `json:"DefaultStorage,omitnil" name:"DefaultStorage"`

	// 最大磁盘规格，单位MB
	MaxStorage *uint64 `json:"MaxStorage,omitnil" name:"MaxStorage"`

	// 最小磁盘规格，单位MB
	MinStorage *uint64 `json:"MinStorage,omitnil" name:"MinStorage"`

	// 可承载qps信息
	Qps *uint64 `json:"Qps,omitnil" name:"Qps"`

	// 连接数限制
	Conns *uint64 `json:"Conns,omitnil" name:"Conns"`

	// 实例mongodb版本信息
	MongoVersionCode *string `json:"MongoVersionCode,omitnil" name:"MongoVersionCode"`

	// 实例mongodb版本号
	MongoVersionValue *uint64 `json:"MongoVersionValue,omitnil" name:"MongoVersionValue"`

	// 实例mongodb版本号（短）
	Version *string `json:"Version,omitnil" name:"Version"`

	// 存储引擎
	EngineName *string `json:"EngineName,omitnil" name:"EngineName"`

	// 集群类型，取值：1-分片集群，0-副本集集群
	ClusterType *uint64 `json:"ClusterType,omitnil" name:"ClusterType"`

	// 最小副本集从节点数
	MinNodeNum *uint64 `json:"MinNodeNum,omitnil" name:"MinNodeNum"`

	// 最大副本集从节点数
	MaxNodeNum *uint64 `json:"MaxNodeNum,omitnil" name:"MaxNodeNum"`

	// 最小分片数
	MinReplicateSetNum *uint64 `json:"MinReplicateSetNum,omitnil" name:"MinReplicateSetNum"`

	// 最大分片数
	MaxReplicateSetNum *uint64 `json:"MaxReplicateSetNum,omitnil" name:"MaxReplicateSetNum"`

	// 最小分片从节点数
	MinReplicateSetNodeNum *uint64 `json:"MinReplicateSetNodeNum,omitnil" name:"MinReplicateSetNodeNum"`

	// 最大分片从节点数
	MaxReplicateSetNodeNum *uint64 `json:"MaxReplicateSetNodeNum,omitnil" name:"MaxReplicateSetNodeNum"`

	// 机器类型，取值：0-HIO，4-HIO10G
	MachineType *string `json:"MachineType,omitnil" name:"MachineType"`
}

type SpecificationInfo struct {
	// 地域信息
	Region *string `json:"Region,omitnil" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 售卖规格信息
	SpecItems []*SpecItem `json:"SpecItems,omitnil" name:"SpecItems"`

	// 是否支持跨可用区部署 1-支持，0-不支持
	SupportMultiAZ *int64 `json:"SupportMultiAZ,omitnil" name:"SupportMultiAZ"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

// Predefined struct for user
type TerminateDBInstancesRequestParams struct {
	// 指定预隔离实例ID。格式如：cmgo-p8vnipr5。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type TerminateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 指定预隔离实例ID。格式如：cmgo-p8vnipr5。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *TerminateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDBInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDBInstancesResponseParams `json:"Response"`
}

func (r *TerminateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 账号名。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 账号权限详情。
	AuthRole []*Auth `json:"AuthRole,omitnil" name:"AuthRole"`

	// 账号创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 账号更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 备注信息。
	UserDesc *string `json:"UserDesc,omitnil" name:"UserDesc"`
}