package ciscosmartbonding

// AttachmentsHolder defines model for AttachmentsHolder.
type AttachmentsHolder struct {
	DataBase64 *string  `json:"DataBase64,omitempty"`
	FileName   *string  `json:"FileName,omitempty"`
	NR         *float32 `json:"NR,omitempty"`
}

// CallActivitiesHolder defines model for CallActivitiesHolder.
type CallActivitiesHolder struct {
	ActivityType      *EntityTypesHolder `json:"ActivityType,omitempty"`
	BPPShortName      *string            `json:"BPPShortName,omitempty"`
	BackTime          *CiscoDateTime     `json:"BackTime,omitempty"`
	BusyTime          *CiscoDateTime     `json:"BusyTime,omitempty"`
	Distance          *int32             `json:"Distance,omitempty"`
	TravelBackMinutes *int32             `json:"TravelBackMinutes,omitempty"`
	TravelBackRemarks *string            `json:"TravelBackRemarks,omitempty"`
	TravelMinutes     *int32             `json:"TravelMinutes,omitempty"`
	TravelRemarks     *string            `json:"TravelRemarks,omitempty"`
	WorkEndTime       *CiscoDateTime     `json:"WorkEndTime,omitempty"`
	WorkMinutes       *int32             `json:"WorkMinutes,omitempty"`
	WorkRemarks       *string            `json:"WorkRemarks,omitempty"`
	WorkStartTime     *CiscoDateTime     `json:"WorkStartTime,omitempty"`
}

// CallAdditionalsHolder defines model for CallAdditionalsHolder.
type CallAdditionalsHolder struct {
	Field1       *StringKeyField  `json:"Field1,omitempty"`
	Field10      *StringKeyField  `json:"Field10,omitempty"`
	Field11      *StringKeyField  `json:"Field11,omitempty"`
	Field12      *StringKeyField  `json:"Field12,omitempty"`
	Field13      *StringKeyField  `json:"Field13,omitempty"`
	Field14      *StringKeyField  `json:"Field14,omitempty"`
	Field15      *StringKeyField  `json:"Field15,omitempty"`
	Field16      *StringKeyField  `json:"Field16,omitempty"`
	Field17      *StringKeyField  `json:"Field17,omitempty"`
	Field18      *StringKeyField  `json:"Field18,omitempty"`
	Field19      *StringKeyField  `json:"Field19,omitempty"`
	Field2       *StringKeyField  `json:"Field2,omitempty"`
	Field20      *StringKeyField  `json:"Field20,omitempty"`
	Field21      *StringKeyField  `json:"Field21,omitempty"`
	Field22      *StringKeyField  `json:"Field22,omitempty"`
	Field23      *StringKeyField  `json:"Field23,omitempty"`
	Field24      *StringKeyField  `json:"Field24,omitempty"`
	Field25      *StringKeyField  `json:"Field25,omitempty"`
	Field26      *StringKeyField  `json:"Field26,omitempty"`
	Field27      *StringKeyField  `json:"Field27,omitempty"`
	Field28      *StringKeyField  `json:"Field28,omitempty"`
	Field29      *StringKeyField  `json:"Field29,omitempty"`
	Field3       *StringKeyField  `json:"Field3,omitempty"`
	Field30      *StringKeyField  `json:"Field30,omitempty"`
	Field31      *StringKeyField  `json:"Field31,omitempty"`
	Field32      *StringKeyField  `json:"Field32,omitempty"`
	Field33      *StringKeyField  `json:"Field33,omitempty"`
	Field34      *StringKeyField  `json:"Field34,omitempty"`
	Field35      *StringKeyField  `json:"Field35,omitempty"`
	Field36      *StringKeyField  `json:"Field36,omitempty"`
	Field37      *StringKeyField  `json:"Field37,omitempty"`
	Field38      *StringKeyField  `json:"Field38,omitempty"`
	Field39      *StringKeyField  `json:"Field39,omitempty"`
	Field4       *StringKeyField  `json:"Field4,omitempty"`
	Field40      *StringKeyField  `json:"Field40,omitempty"`
	Field41      *StringKeyField  `json:"Field41,omitempty"`
	Field42      *StringKeyField  `json:"Field42,omitempty"`
	Field43      *StringKeyField  `json:"Field43,omitempty"`
	Field44      *StringKeyField  `json:"Field44,omitempty"`
	Field45      *StringKeyField  `json:"Field45,omitempty"`
	Field46      *StringKeyField  `json:"Field46,omitempty"`
	Field47      *StringKeyField  `json:"Field47,omitempty"`
	Field48      *StringKeyField  `json:"Field48,omitempty"`
	Field49      *StringKeyField  `json:"Field49,omitempty"`
	Field5       *StringKeyField  `json:"Field5,omitempty"`
	Field50      *StringKeyField  `json:"Field50,omitempty"`
	Field51      *StringKeyField  `json:"Field51,omitempty"`
	Field52      *StringKeyField  `json:"Field52,omitempty"`
	Field53      *StringKeyField  `json:"Field53,omitempty"`
	Field54      *StringKeyField  `json:"Field54,omitempty"`
	Field55      *StringKeyField  `json:"Field55,omitempty"`
	Field56      *StringKeyField  `json:"Field56,omitempty"`
	Field57      *StringKeyField  `json:"Field57,omitempty"`
	Field58      *StringKeyField  `json:"Field58,omitempty"`
	Field59      *StringKeyField  `json:"Field59,omitempty"`
	Field6       *StringKeyField  `json:"Field6,omitempty"`
	Field60      *StringKeyField  `json:"Field60,omitempty"`
	Field61      *StringKeyField  `json:"Field61,omitempty"`
	Field62      *StringKeyField  `json:"Field62,omitempty"`
	Field63      *StringKeyField  `json:"Field63,omitempty"`
	Field64      *StringKeyField  `json:"Field64,omitempty"`
	Field7       *StringKeyField  `json:"Field7,omitempty"`
	Field8       *StringKeyField  `json:"Field8,omitempty"`
	Field9       *StringKeyField  `json:"Field9,omitempty"`
	FunctionPort *StringKeyField  `json:"FunctionPort,omitempty"`
	PosNr        *IntegerKeyField `json:"PosNr,omitempty"`
	Url1         *string          `json:"Url1,omitempty"`
	Url2         *string          `json:"Url2,omitempty"`
	Url3         *string          `json:"Url3,omitempty"`
	Url4         *string          `json:"Url4,omitempty"`
	Url5         *string          `json:"Url5,omitempty"`
	Url6         *string          `json:"Url6,omitempty"`
	Url7         *string          `json:"Url7,omitempty"`
	Url8         *string          `json:"Url8,omitempty"`
	UrlName1     *string          `json:"UrlName1,omitempty"`
	UrlName2     *string          `json:"UrlName2,omitempty"`
	UrlName3     *string          `json:"UrlName3,omitempty"`
	UrlName4     *string          `json:"UrlName4,omitempty"`
	UrlName5     *string          `json:"UrlName5,omitempty"`
	UrlName6     *string          `json:"UrlName6,omitempty"`
	UrlName7     *string          `json:"UrlName7,omitempty"`
	UrlName8     *string          `json:"UrlName8,omitempty"`
}

// CallCalculationsHolder defines model for CallCalculationsHolder.
type CallCalculationsHolder struct {
	ChargingTypes        *ChargingTypesHolder `json:"ChargingTypes,omitempty"`
	ContractShortName    *string              `json:"ContractShortName,omitempty"`
	RequestedWorkEndTime *CiscoDateTime       `json:"RequestedWorkEndTime,omitempty"`
	Revenue              *float32             `json:"Revenue,omitempty"`
	ServiceType          *string              `json:"ServiceType,omitempty"`
}

// CallData defines model for CallData.
type CallData struct {
	Attachments      *[]AttachmentsHolder           `json:"Attachments,omitempty"`
	CallActivities   *[]CallActivitiesHolder        `json:"CallActivities,omitempty"`
	CallAdditionals  *[]CallAdditionalsHolder       `json:"CallAdditionals,omitempty"`
	CallCalculations *CallCalculationsHolder        `json:"CallCalculations,omitempty"`
	CallStates       *CallSystemCodesHolder         `json:"CallStates,omitempty"`
	Calls            InboundCallsHolder             `json:"Calls"`
	ContractElements *InboundContractElementsHolder `json:"ContractElements,omitempty"`
	Contracts        *InboundContractsHolder        `json:"Contracts,omitempty"`
	Control          *ControlFlagsHolder            `json:"Control,omitempty"`
	DeviceMovements  *[]DeviceMovementsHolder       `json:"DeviceMovements,omitempty"`
	Devices          *DevicesHolder                 `json:"Devices,omitempty"`
	ExtTableValues   *CallExtensionsHolder          `json:"ExtTableValues,omitempty"`
	FailureTypes     *CallSystemCodesHolder         `json:"FailureTypes,omitempty"`
	Impacts          *CallSystemCodesHolder         `json:"Impacts,omitempty"`
	Locations        *LocationsHolder               `json:"Locations,omitempty"`
	ParentCall       *ParentCallsHolder             `json:"ParentCall,omitempty"`
	Priorities       *CallSystemCodesHolder         `json:"Priorities,omitempty"`
	ProblemTypes     *CallSystemCodesHolder         `json:"ProblemTypes,omitempty"`
	Queues1          *QueuesHolder                  `json:"Queues1,omitempty"`
	Queues2          *QueuesHolder                  `json:"Queues2,omitempty"`
	Queues3          *QueuesHolder                  `json:"Queues3,omitempty"`
	RequestTypes     *CallSystemCodesHolder         `json:"RequestTypes,omitempty"`
	Severities       *CallSystemCodesHolder         `json:"Severities,omitempty"`
	Urgency          *CallSystemCodesHolder         `json:"Urgency,omitempty"`
}

// CallExtensionsHolder defines model for CallExtensionsHolder.
type CallExtensionsHolder struct {
	Field1   *string `json:"Field1,omitempty"`
	Field10  *string `json:"Field10,omitempty"`
	Field100 *string `json:"Field100,omitempty"`
	Field101 *string `json:"Field101,omitempty"`
	Field102 *string `json:"Field102,omitempty"`
	Field103 *string `json:"Field103,omitempty"`
	Field104 *string `json:"Field104,omitempty"`
	Field105 *string `json:"Field105,omitempty"`
	Field106 *string `json:"Field106,omitempty"`
	Field107 *string `json:"Field107,omitempty"`
	Field108 *string `json:"Field108,omitempty"`
	Field109 *string `json:"Field109,omitempty"`
	Field11  *string `json:"Field11,omitempty"`
	Field110 *string `json:"Field110,omitempty"`
	Field111 *string `json:"Field111,omitempty"`
	Field112 *string `json:"Field112,omitempty"`
	Field113 *string `json:"Field113,omitempty"`
	Field114 *string `json:"Field114,omitempty"`
	Field115 *string `json:"Field115,omitempty"`
	Field116 *string `json:"Field116,omitempty"`
	Field117 *string `json:"Field117,omitempty"`
	Field118 *string `json:"Field118,omitempty"`
	Field119 *string `json:"Field119,omitempty"`
	Field12  *string `json:"Field12,omitempty"`
	Field120 *string `json:"Field120,omitempty"`
	Field121 *string `json:"Field121,omitempty"`
	Field122 *string `json:"Field122,omitempty"`
	Field123 *string `json:"Field123,omitempty"`
	Field124 *string `json:"Field124,omitempty"`
	Field125 *string `json:"Field125,omitempty"`
	Field126 *string `json:"Field126,omitempty"`
	Field127 *string `json:"Field127,omitempty"`
	Field128 *string `json:"Field128,omitempty"`
	Field13  *string `json:"Field13,omitempty"`
	Field14  *string `json:"Field14,omitempty"`
	Field15  *string `json:"Field15,omitempty"`
	Field16  *string `json:"Field16,omitempty"`
	Field17  *string `json:"Field17,omitempty"`
	Field18  *string `json:"Field18,omitempty"`
	Field19  *string `json:"Field19,omitempty"`
	Field2   *string `json:"Field2,omitempty"`
	Field20  *string `json:"Field20,omitempty"`
	Field21  *string `json:"Field21,omitempty"`
	Field22  *string `json:"Field22,omitempty"`
	Field23  *string `json:"Field23,omitempty"`
	Field24  *string `json:"Field24,omitempty"`
	Field25  *string `json:"Field25,omitempty"`
	Field26  *string `json:"Field26,omitempty"`
	Field27  *string `json:"Field27,omitempty"`
	Field28  *string `json:"Field28,omitempty"`
	Field29  *string `json:"Field29,omitempty"`
	Field3   *string `json:"Field3,omitempty"`
	Field30  *string `json:"Field30,omitempty"`
	Field31  *string `json:"Field31,omitempty"`
	Field32  *string `json:"Field32,omitempty"`
	Field33  *string `json:"Field33,omitempty"`
	Field34  *string `json:"Field34,omitempty"`
	Field35  *string `json:"Field35,omitempty"`
	Field36  *string `json:"Field36,omitempty"`
	Field37  *string `json:"Field37,omitempty"`
	Field38  *string `json:"Field38,omitempty"`
	Field39  *string `json:"Field39,omitempty"`
	Field4   *string `json:"Field4,omitempty"`
	Field40  *string `json:"Field40,omitempty"`
	Field41  *string `json:"Field41,omitempty"`
	Field42  *string `json:"Field42,omitempty"`
	Field43  *string `json:"Field43,omitempty"`
	Field44  *string `json:"Field44,omitempty"`
	Field45  *string `json:"Field45,omitempty"`
	Field46  *string `json:"Field46,omitempty"`
	Field47  *string `json:"Field47,omitempty"`
	Field48  *string `json:"Field48,omitempty"`
	Field49  *string `json:"Field49,omitempty"`
	Field5   *string `json:"Field5,omitempty"`
	Field50  *string `json:"Field50,omitempty"`
	Field51  *string `json:"Field51,omitempty"`
	Field52  *string `json:"Field52,omitempty"`
	Field53  *string `json:"Field53,omitempty"`
	Field54  *string `json:"Field54,omitempty"`
	Field55  *string `json:"Field55,omitempty"`
	Field56  *string `json:"Field56,omitempty"`
	Field57  *string `json:"Field57,omitempty"`
	Field58  *string `json:"Field58,omitempty"`
	Field59  *string `json:"Field59,omitempty"`
	Field6   *string `json:"Field6,omitempty"`
	Field60  *string `json:"Field60,omitempty"`
	Field61  *string `json:"Field61,omitempty"`
	Field62  *string `json:"Field62,omitempty"`
	Field63  *string `json:"Field63,omitempty"`
	Field64  *string `json:"Field64,omitempty"`
	Field65  *string `json:"Field65,omitempty"`
	Field66  *string `json:"Field66,omitempty"`
	Field67  *string `json:"Field67,omitempty"`
	Field68  *string `json:"Field68,omitempty"`
	Field69  *string `json:"Field69,omitempty"`
	Field7   *string `json:"Field7,omitempty"`
	Field70  *string `json:"Field70,omitempty"`
	Field71  *string `json:"Field71,omitempty"`
	Field72  *string `json:"Field72,omitempty"`
	Field73  *string `json:"Field73,omitempty"`
	Field74  *string `json:"Field74,omitempty"`
	Field75  *string `json:"Field75,omitempty"`
	Field76  *string `json:"Field76,omitempty"`
	Field77  *string `json:"Field77,omitempty"`
	Field78  *string `json:"Field78,omitempty"`
	Field79  *string `json:"Field79,omitempty"`
	Field8   *string `json:"Field8,omitempty"`
	Field80  *string `json:"Field80,omitempty"`
	Field81  *string `json:"Field81,omitempty"`
	Field82  *string `json:"Field82,omitempty"`
	Field83  *string `json:"Field83,omitempty"`
	Field84  *string `json:"Field84,omitempty"`
	Field85  *string `json:"Field85,omitempty"`
	Field86  *string `json:"Field86,omitempty"`
	Field87  *string `json:"Field87,omitempty"`
	Field88  *string `json:"Field88,omitempty"`
	Field89  *string `json:"Field89,omitempty"`
	Field9   *string `json:"Field9,omitempty"`
	Field90  *string `json:"Field90,omitempty"`
	Field91  *string `json:"Field91,omitempty"`
	Field92  *string `json:"Field92,omitempty"`
	Field93  *string `json:"Field93,omitempty"`
	Field94  *string `json:"Field94,omitempty"`
	Field95  *string `json:"Field95,omitempty"`
	Field96  *string `json:"Field96,omitempty"`
	Field97  *string `json:"Field97,omitempty"`
	Field98  *string `json:"Field98,omitempty"`
	Field99  *string `json:"Field99,omitempty"`
}

// CallNotesHolder defines model for CallNotesHolder.
type CallNotesHolder struct {
	IsPrivate *string `json:"IsPrivate,omitempty"`
	Type      *string `json:"Type,omitempty"`
	Value     *string `json:"value,omitempty"`
}

// CallSystemCodesHolder defines model for CallSystemCodesHolder.
type CallSystemCodesHolder struct {
	Name      *string `json:"Name,omitempty"`
	ShortName *string `json:"ShortName,omitempty"`
}

// ChargingTypesHolder defines model for ChargingTypesHolder.
type ChargingTypesHolder struct {
	ShortName *string `json:"ShortName,omitempty"`
}

// ComponentsHolder defines model for ComponentsHolder.
type ComponentsHolder struct {
	Component           *string  `json:"Component,omitempty"`
	Description         *string  `json:"Description,omitempty"`
	Hostname            *string  `json:"Hostname,omitempty"`
	IPAddress           *string  `json:"IPAddress,omitempty"`
	InvNr               *string  `json:"InvNr,omitempty"`
	Location            *string  `json:"Location,omitempty"`
	LocationCategory    *string  `json:"LocationCategory,omitempty"`
	LocationCity        *string  `json:"LocationCity,omitempty"`
	LocationCountry     *string  `json:"LocationCountry,omitempty"`
	LocationDescription *string  `json:"LocationDescription,omitempty"`
	LocationLevel       *float32 `json:"LocationLevel,omitempty"`
	LocationName        *string  `json:"LocationName,omitempty"`
	LocationProvince    *string  `json:"LocationProvince,omitempty"`
	LocationRegion      *string  `json:"LocationRegion,omitempty"`
	LocationStreet      *string  `json:"LocationStreet,omitempty"`
	LocationTel         *string  `json:"LocationTel,omitempty"`
	LocationZip         *string  `json:"LocationZip,omitempty"`
	MACAddress          *string  `json:"MACAddress,omitempty"`
	Manufacturer        *string  `json:"Manufacturer,omitempty"`
	Model               *string  `json:"Model,omitempty"`
	Name                *string  `json:"Name,omitempty"`
	OpSys               *string  `json:"OpSys,omitempty"`
	Room                *string  `json:"Room,omitempty"`
	SerNr               *string  `json:"SerNr,omitempty"`
	SerNrProv           *string  `json:"SerNrProv,omitempty"`
	ShortName           *string  `json:"ShortName,omitempty"`
	Type                *string  `json:"Type,omitempty"`
}

// ControlFlagsHolder defines model for ControlFlagsHolder.
type ControlFlagsHolder struct {
	DeactivateOutboundTriggers                     *string `json:"DeactivateOutboundTriggers,omitempty"`
	ForwardAfterClose                              *string `json:"ForwardAfterClose,omitempty"`
	MakePreSelection                               *string `json:"MakePreSelection,omitempty"`
	MergeDefaultsOnContractOrContractElementChange *string `json:"MergeDefaultsOnContractOrContractElementChange,omitempty"`
	MergeLocationFromCaller                        *string `json:"MergeLocationFromCaller,omitempty"`
	SetCurrentCallState                            *string `json:"SetCurrentCallState,omitempty"`
	UpdateAfterCloseAllowed                        *string `json:"UpdateAfterCloseAllowed,omitempty"`
	UseCodeDefaults                                *string `json:"UseCodeDefaults,omitempty"`
	UseIndependentDeviceRef                        *string `json:"UseIndependentDeviceRef,omitempty"`
	UsePartnerSideSuccessors                       *string `json:"UsePartnerSideSuccessors,omitempty"`
	UseSuccessors                                  *string `json:"UseSuccessors,omitempty"`
}

// DeviceMovementsDeviceHolder defines model for DeviceMovementsDeviceHolder.
type DeviceMovementsDeviceHolder struct {
	Component    *string `json:"Component,omitempty"`
	InvNr        *string `json:"InvNr,omitempty"`
	Manufacturer *string `json:"Manufacturer,omitempty"`
	Model        *string `json:"Model,omitempty"`
	Owner        *string `json:"Owner,omitempty"`
	SerNr        *string `json:"SerNr,omitempty"`
	ShortName    *string `json:"ShortName,omitempty"`
	Type         *string `json:"Type,omitempty"`
	UNSPSC       *string `json:"UNSPSC,omitempty"`
}

// DeviceMovementsHolder defines model for DeviceMovementsHolder.
type DeviceMovementsHolder struct {
	Device       *DeviceMovementsDeviceHolder `json:"Device,omitempty"`
	FromLocation *LocationsHolder             `json:"FromLocation,omitempty"`
	ToLocation   *LocationsHolder             `json:"ToLocation,omitempty"`
}

// DevicesHolder defines model for DevicesHolder.
type DevicesHolder struct {
	Component    *string `json:"Component,omitempty"`
	InvNr        *string `json:"InvNr,omitempty"`
	Manufacturer *string `json:"Manufacturer,omitempty"`
	Model        *string `json:"Model,omitempty"`
	SerNr        *string `json:"SerNr,omitempty"`
	ShortName    *string `json:"ShortName,omitempty"`
	Type         *string `json:"Type,omitempty"`
	UNSPSC       *string `json:"UNSPSC,omitempty"`
}

// EntityTypesHolder defines model for EntityTypesHolder.
type EntityTypesHolder struct {
	Name      *string `json:"Name,omitempty"`
	ShortName *string `json:"ShortName,omitempty"`
}

// InboundCallsHolder defines model for InboundCallsHolder.
type InboundCallsHolder struct {
	AddRemarksToSummary        *string           `json:"AddRemarksToSummary,omitempty"`
	CCP                        *PersonsHolder    `json:"CCP,omitempty"`
	CHD                        *PersonsHolder    `json:"CHD,omitempty"`
	CallAcknowledgeTime        *CiscoDateTime    `json:"CallAcknowledgeTime,omitempty"`
	CallCloseTime              *CiscoDateTime    `json:"CallCloseTime,omitempty"`
	CallOpenTime               *CiscoDateTime    `json:"CallOpenTime,omitempty"`
	CallRecoveryTime           *CiscoDateTime    `json:"CallRecoveryTime,omitempty"`
	CallResponseTime           *CiscoDateTime    `json:"CallResponseTime,omitempty"`
	CallSendTime               *CiscoDateTime    `json:"CallSendTime,omitempty"`
	CallStartSLATime           *CiscoDateTime    `json:"CallStartSLATime,omitempty"`
	Caller                     *PersonsHolder    `json:"Caller,omitempty"`
	CustCallID                 *string           `json:"CustCallID,omitempty"`
	CustomerCategory1          *string           `json:"CustomerCategory1,omitempty"`
	CustomerCategory2          *string           `json:"CustomerCategory2,omitempty"`
	CustomerCategory3          *string           `json:"CustomerCategory3,omitempty"`
	CustomerCategory4          *string           `json:"CustomerCategory4,omitempty"`
	CustomerCategory5          *string           `json:"CustomerCategory5,omitempty"`
	CustomerReasonCategory1    *string           `json:"CustomerReasonCategory1,omitempty"`
	CustomerReasonCategory2    *string           `json:"CustomerReasonCategory2,omitempty"`
	CustomerReasonCategory3    *string           `json:"CustomerReasonCategory3,omitempty"`
	CustomerReasonCategory4    *string           `json:"CustomerReasonCategory4,omitempty"`
	CustomerReasonCategory5    *string           `json:"CustomerReasonCategory5,omitempty"`
	CustomerRequestedEndTime   *CiscoDateTime    `json:"CustomerRequestedEndTime,omitempty"`
	CustomerRequestedStartTime *CiscoDateTime    `json:"CustomerRequestedStartTime,omitempty"`
	Description                *string           `json:"Description,omitempty"`
	Diagnosis                  *string           `json:"Diagnosis,omitempty"`
	MainComp                   *ComponentsHolder `json:"MainComp,omitempty"`
	Notes                      *CallNotesHolder  `json:"Notes,omitempty"`
	Ownership                  *string           `json:"Ownership,omitempty"`
	PartnerCoreTicketId        *float32          `json:"PartnerCoreTicketId,omitempty"`
	ProblemStartTime           *CiscoDateTime    `json:"ProblemStartTime,omitempty"`
	ProviderCategory1          *string           `json:"ProviderCategory1,omitempty"`
	ProviderCategory2          *string           `json:"ProviderCategory2,omitempty"`
	ProviderCategory3          *string           `json:"ProviderCategory3,omitempty"`
	ProviderCategory4          *string           `json:"ProviderCategory4,omitempty"`
	ProviderCategory5          *string           `json:"ProviderCategory5,omitempty"`
	ProviderReasonCategory1    *string           `json:"ProviderReasonCategory1,omitempty"`
	ProviderReasonCategory2    *string           `json:"ProviderReasonCategory2,omitempty"`
	ProviderReasonCategory3    *string           `json:"ProviderReasonCategory3,omitempty"`
	ProviderReasonCategory4    *string           `json:"ProviderReasonCategory4,omitempty"`
	ProviderReasonCategory5    *string           `json:"ProviderReasonCategory5,omitempty"`
	ProviderScheduledEndTime   *CiscoDateTime    `json:"ProviderScheduledEndTime,omitempty"`
	ProviderScheduledStartTime *CiscoDateTime    `json:"ProviderScheduledStartTime,omitempty"`
	Remarks                    *string           `json:"Remarks,omitempty"`
	SDCallID                   *float32          `json:"SDCallID,omitempty"`
	SPCallID                   *string           `json:"SPCallID,omitempty"`
	ShortDescription           *string           `json:"ShortDescription,omitempty"`
	Solution                   *string           `json:"Solution,omitempty"`
	SubComp                    *ComponentsHolder `json:"SubComp,omitempty"`
	SysSpecField1              *string           `json:"SysSpecField1,omitempty"`
	SysSpecField10             *string           `json:"SysSpecField10,omitempty"`
	SysSpecField2              *string           `json:"SysSpecField2,omitempty"`
	SysSpecField3              *string           `json:"SysSpecField3,omitempty"`
	SysSpecField4              *string           `json:"SysSpecField4,omitempty"`
	SysSpecField5              *string           `json:"SysSpecField5,omitempty"`
	SysSpecField6              *string           `json:"SysSpecField6,omitempty"`
	SysSpecField7              *string           `json:"SysSpecField7,omitempty"`
	SysSpecField8              *string           `json:"SysSpecField8,omitempty"`
	SysSpecField9              *string           `json:"SysSpecField9,omitempty"`
}

// InboundContractElementsHolder defines model for InboundContractElementsHolder.
type InboundContractElementsHolder struct {
	CompLocation      *string `json:"CompLocation,omitempty"`
	Component         *string `json:"Component,omitempty"`
	ContractIDCust    *string `json:"ContractIDCust,omitempty"`
	ContractIDProv    *string `json:"ContractIDProv,omitempty"`
	CustomerCategory1 *string `json:"CustomerCategory1,omitempty"`
	CustomerCategory2 *string `json:"CustomerCategory2,omitempty"`
	CustomerCategory3 *string `json:"CustomerCategory3,omitempty"`
	CustomerCategory4 *string `json:"CustomerCategory4,omitempty"`
	CustomerCategory5 *string `json:"CustomerCategory5,omitempty"`
	Name              *string `json:"Name,omitempty"`
	NewName           *string `json:"NewName,omitempty"`
	NewShortName      *string `json:"NewShortName,omitempty"`
	ProviderCategory1 *string `json:"ProviderCategory1,omitempty"`
	ProviderCategory2 *string `json:"ProviderCategory2,omitempty"`
	ProviderCategory3 *string `json:"ProviderCategory3,omitempty"`
	ProviderCategory4 *string `json:"ProviderCategory4,omitempty"`
	ProviderCategory5 *string `json:"ProviderCategory5,omitempty"`
	ShortName         *string `json:"ShortName,omitempty"`
}

// InboundContractsHolder defines model for InboundContractsHolder.
type InboundContractsHolder struct {
	BPOrganizationShortNameCust *string `json:"BPOrganizationShortNameCust,omitempty"`
	BPOrganizationShortNameProv *string `json:"BPOrganizationShortNameProv,omitempty"`
	BPartnerShortNameCust       *string `json:"BPartnerShortNameCust,omitempty"`
	BPartnerShortNameProv       *string `json:"BPartnerShortNameProv,omitempty"`
	ContractIDCust              *string `json:"ContractIDCust,omitempty"`
	ContractIDProv              *string `json:"ContractIDProv,omitempty"`
	CustIDProv                  *string `json:"CustIDProv,omitempty"`
	Name                        *string `json:"Name,omitempty"`
	NewName                     *string `json:"NewName,omitempty"`
	NewShortName                *string `json:"NewShortName,omitempty"`
	ProvIDCust                  *string `json:"ProvIDCust,omitempty"`
	ShortName                   *string `json:"ShortName,omitempty"`
}

// IntegerKeyField defines model for IntegerKeyField.
type IntegerKeyField struct {
	Key   *string `json:"key,omitempty"`
	Value *int32  `json:"value,omitempty"`
}

// LocationsHolder defines model for LocationsHolder.
type LocationsHolder struct {
	ShortName *string `json:"ShortName,omitempty"`
}

// ParentCallsHolder defines model for ParentCallsHolder.
type ParentCallsHolder struct {
	CustCallID *string  `json:"CustCallID,omitempty"`
	SDCallID   *float32 `json:"SDCallID,omitempty"`
	SPCallID   *string  `json:"SPCallID,omitempty"`
}

// PersonsHolder defines model for PersonsHolder.
type PersonsHolder struct {
	Category         *string `json:"Category,omitempty"`
	Department       *string `json:"Department,omitempty"`
	Description      *string `json:"Description,omitempty"`
	EMail            *string `json:"EMail,omitempty"`
	Fax              *string `json:"Fax,omitempty"`
	FirstName        *string `json:"FirstName,omitempty"`
	Language         *string `json:"Language,omitempty"`
	LastName         *string `json:"LastName,omitempty"`
	Location         *string `json:"Location,omitempty"`
	LocationCity     *string `json:"LocationCity,omitempty"`
	LocationCountry  *string `json:"LocationCountry,omitempty"`
	LocationProvince *string `json:"LocationProvince,omitempty"`
	LocationStreet   *string `json:"LocationStreet,omitempty"`
	LocationZip      *string `json:"LocationZip,omitempty"`
	MobileTel        *string `json:"MobileTel,omitempty"`
	PIN              *string `json:"PIN,omitempty"`
	Room             *string `json:"Room,omitempty"`
	Salutation       *string `json:"Salutation,omitempty"`
	ShortName        *string `json:"ShortName,omitempty"`
	Sign             *string `json:"Sign,omitempty"`
	Tel              *string `json:"Tel,omitempty"`
	Tel2             *string `json:"Tel2,omitempty"`
	Title            *string `json:"Title,omitempty"`
}

// QueuesHolder defines model for QueuesHolder.
type QueuesHolder struct {
	ShortName *string `json:"ShortName,omitempty"`
}

// StringKeyField defines model for StringKeyField.
type StringKeyField struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// PutCallJSONBody defines parameters for PutCall.
type PutCallJSONBody CallData

// PutCallJSONRequestBody defines body for PutCall for application/json ContentType.
type PutCallJSONRequestBody PutCallJSONBody
