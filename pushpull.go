package ciscosmartbonding

import (
	"encoding/json"
)

// CiscoRemarksError is an error type that is sometimes returned in the Remarks field!! IKR!
type CiscoRemarksError struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
}

// StringOrSliceOfErrors caters for Cisco returning either a string or a slice of errors in the Remarks field
// Yep, I'm serious!
// We provide a custom MarshalJSON to only send the string back to Cisco and a custom UnmarshalJSON to get the
// errors or a string
type StringOrSliceOfErrors struct {
	Errors       []CiscoRemarksError
	RemarkString string // used to hold the remarks if we have a single string instead of errors
}

// MarshalJSON lets us send just the string
func (se *StringOrSliceOfErrors) MarshalJSON() ([]byte, error) {
	if se == nil {
		return []byte(""), nil
	}
	return json.Marshal(se.RemarkString)
}

// UnmarshalJSON convert JSON object array of errors or string to object
func (se *StringOrSliceOfErrors) UnmarshalJSON(data []byte) error {
	var jsonObj interface{}
	err := json.Unmarshal(data, &jsonObj)
	if err != nil {
		return err
	}
	switch obj := jsonObj.(type) {
	case string:
		*se = StringOrSliceOfErrors{RemarkString: obj}
		return nil
	case []interface{}:
		var errSlice []CiscoRemarksError
		err := json.Unmarshal(data, &errSlice)
		if err != nil {
			return err
		}
		*se = StringOrSliceOfErrors{Errors: errSlice}
		return nil
	}
	// fmt.Println("GOT TO END")
	return ErrUnsupportedType
}

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
	CallStates       *CallSystemCodesHolder         `json:"CallStates,omitempty"`    // Ticket Status of the ticket/case e.g. "In Progress" (from list). M for Create/Update.
	CallStatesSPR    *CallSystemCodesHolder         `json:"CallStatesSPR,omitempty"` // TAC Ticket Status
	Calls            *InboundCallsHolder            `json:"Calls,omitempty"`
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
	Priorities       *CallSystemCodesHolder         `json:"Priorities,omitempty"` // Urgency of the ticket/case (shadow or escalated case) e.g. "Shadow" (from list). M on Create.
	ProblemTypes     *CallSystemCodesHolder         `json:"ProblemTypes,omitempty"`
	Queues1          *QueuesHolder                  `json:"Queues1,omitempty"`
	Queues2          *QueuesHolder                  `json:"Queues2,omitempty"`
	Queues3          *QueuesHolder                  `json:"Queues3,omitempty"`
	RequestTypes     *CallSystemCodesHolder         `json:"RequestTypes,omitempty"`
	Severities       *CallSystemCodesHolder         `json:"Severities,omitempty"` // Severity of the ticket/case e.g. "MED" (from list). M for Create.
	Urgency          *CallSystemCodesHolder         `json:"Urgency,omitempty"`

	// They may also send us an "error" if there is no data to fetch, also with a 200 status.
	// So we'll add the fields here so we can check against them
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

// CallExtensionsHolder defines model for CallExtensionsHolder.
type CallExtensionsHolder struct {
	Field1   *string `json:"Field1,omitempty"`
	Field10  *string `json:"Field10,omitempty"`
	Field100 *string `json:"Field100,omitempty"`
	Field101 *string `json:"Field101,omitempty"`
	Field102 *string `json:"Field102,omitempty"`
	Field103 *string `json:"Field103,omitempty"`
	Field104 *string `json:"Field104,omitempty"` // Partner Ticket Open Date & Time. M for Create.
	Field105 *string `json:"Field105,omitempty"` // Partner ticket acceptance time. M for Update.
	Field106 *string `json:"Field106,omitempty"` // Partner Ticket Closed Date & Time. M for Update (on closure).
	Field107 *string `json:"Field107,omitempty"` // RMA # associated with partner ticket.
	Field108 *string `json:"Field108,omitempty"` // 3rd Party Case Number e.g. "VendINC123". M for Create.
	Field109 *string `json:"Field109,omitempty"` // 3rd Party Product Vendor Name e.g. "vendor". M for Create.
	Field11  *string `json:"Field11,omitempty"`  // Recipient / Site Contact. M for Create (Premium).
	Field110 *string `json:"Field110,omitempty"` // Customer 3rd Party Support Contract Active (customer/partner) e.g. "Yes".
	Field111 *string `json:"Field111,omitempty"` // 3rd Party Suspected Product Name e.g. "Router Cat9K". M for Create.
	Field112 *string `json:"Field112,omitempty"` // 3rd Party Suspected Product Issue e.g. "Power". M for Create.
	Field113 *string `json:"Field113,omitempty"` // 3rd Party Main Contact First Name (case owner) e.g. "Simon". M for Create.
	Field114 *string `json:"Field114,omitempty"` // 3rd Party Main Contact Last Name(case owner) e.g. "Smith". M for Create.
	Field115 *string `json:"Field115,omitempty"` // 3rd Party Contact Telephone (case owner) e.g. "1-855-000-0000". M for Create.
	Field116 *string `json:"Field116,omitempty"` // 3rd Party Contact Email (case owner) e.g. simon.smith@vendor.com. M for Create.
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
	Field13  *string `json:"Field13,omitempty"` // Site Availability From in  yyyy-MM-dd HH:mm:ss | | dd-MM-yyyy HH:mm:ss format. M for Create (Premium)
	Field14  *string `json:"Field14,omitempty"` // Site Availability To in  yyyy-MM-dd HH:mm:ss | | dd-MM-yyyy HH:mm:ss format. M for Create (Premium)
	Field15  *string `json:"Field15,omitempty"`
	Field16  *string `json:"Field16,omitempty"`
	Field17  *string `json:"Field17,omitempty"`
	Field18  *string `json:"Field18,omitempty"`
	Field19  *string `json:"Field19,omitempty"`
	Field2   *string `json:"Field2,omitempty"`
	Field20  *string `json:"Field20,omitempty"`
	Field21  *string `json:"Field21,omitempty"` // Security Clearance Code e.g. "None". M for Create (Premium).
	Field22  *string `json:"Field22,omitempty"`
	Field23  *string `json:"Field23,omitempty"`
	Field24  *string `json:"Field24,omitempty"`
	Field25  *string `json:"Field25,omitempty"` // Ship To Address - Country. M for Create.
	Field26  *string `json:"Field26,omitempty"` // Ship To Country Code e.g. 1. M for Create.
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
	InvNr               *string  `json:"InvNr,omitempty"`            // Cisco Contract ID (Partner Contract Number). Must match Cisco DB.
	Location            *string  `json:"Location,omitempty"`         // Install Site ID
	LocationCategory    *string  `json:"LocationCategory,omitempty"` // Site Type e.g. Data Center. M for Create (Premium).
	LocationCity        *string  `json:"LocationCity,omitempty"`     // For SubComp, "Ship To Address - City". M for Create.
	LocationCountry     *string  `json:"LocationCountry,omitempty"`
	LocationDescription *string  `json:"LocationDescription,omitempty"`
	LocationLevel       *float32 `json:"LocationLevel,omitempty"`
	LocationName        *string  `json:"LocationName,omitempty"`
	LocationProvince    *string  `json:"LocationProvince,omitempty"` // Ship To Address - State/Territory. M for Create.
	LocationRegion      *string  `json:"LocationRegion,omitempty"`
	LocationStreet      *string  `json:"LocationStreet,omitempty"` // For SubComp, Ship To Address - Line 1. M for Create.
	LocationTel         *string  `json:"LocationTel,omitempty"`
	LocationZip         *string  `json:"LocationZip,omitempty"` // Ship To Address - Postal Code. M for Create.
	MACAddress          *string  `json:"MACAddress,omitempty"`
	Manufacturer        *string  `json:"Manufacturer,omitempty"`
	Model               *string  `json:"Model,omitempty"`
	Name                *string  `json:"Name,omitempty"`
	OpSys               *string  `json:"OpSys,omitempty"`
	Room                *string  `json:"Room,omitempty"` // Product ID for Entitlement (Software). Must match Cisco DB.
	SerNr               *string  `json:"SerNr,omitempty"`
	SerNrProv           *string  `json:"SerNrProv,omitempty"` // Serial Number used for Entitlement (Hardware). Must match Cisco DB.
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
	AddRemarksToSummary        *string            `json:"AddRemarksToSummary,omitempty"`
	CCP                        *PersonsHolder     `json:"CCP,omitempty"` // TAC Engineer Details
	CHD                        *PersonsHolder     `json:"CHD,omitempty"` // Partner Primary Contact / Information. First, Last, Tel, Email, Sign M for Create.
	CallAcknowledgeTime        *CiscoDateTime     `json:"CallAcknowledgeTime,omitempty"`
	CallCloseTime              *CiscoDateTime     `json:"CallCloseTime,omitempty"`
	CallOpenTime               *CiscoDateTime     `json:"CallOpenTime,omitempty"`
	CallRecoveryTime           *CiscoDateTime     `json:"CallRecoveryTime,omitempty"`
	CallResponseTime           *CiscoDateTime     `json:"CallResponseTime,omitempty"`
	CallSendTime               *CiscoDateTime     `json:"CallSendTime,omitempty"`
	CallStartSLATime           *CiscoDateTime     `json:"CallStartSLATime,omitempty"`
	Caller                     *PersonsHolder     `json:"Caller,omitempty"`     // End Customer Information. First, Last and Email M for Create.
	CustCallID                 *string            `json:"CustCallID,omitempty"` // PARTNER Networks ticket ID e.g. INC000456. M for Create/Update.
	CustomerCategory1          *string            `json:"CustomerCategory1,omitempty"`
	CustomerCategory2          *string            `json:"CustomerCategory2,omitempty"`
	CustomerCategory3          *string            `json:"CustomerCategory3,omitempty"`
	CustomerCategory4          *string            `json:"CustomerCategory4,omitempty"`
	CustomerCategory5          *string            `json:"CustomerCategory5,omitempty"`
	CustomerReasonCategory1    *string            `json:"CustomerReasonCategory1,omitempty"` // Technology Code for entitlement - can be sent as number or text e.g. "13" or "LAN Switching" (from list). M for Create.
	CustomerReasonCategory2    *string            `json:"CustomerReasonCategory2,omitempty"` // Sub-Technology Code for entitlement - can be sent as number or text e.g. "5190" or Cat9500X" (from list). M for Create.
	CustomerReasonCategory3    *string            `json:"CustomerReasonCategory3,omitempty"` // Problem Code for entitlement e.g. "INSTLL_UNSTLL_UPGRD" (from list). M for Create.
	CustomerReasonCategory4    *string            `json:"CustomerReasonCategory4,omitempty"`
	CustomerReasonCategory5    *string            `json:"CustomerReasonCategory5,omitempty"`
	CustomerRequestedEndTime   *CiscoDateTime     `json:"CustomerRequestedEndTime,omitempty"`
	CustomerRequestedStartTime *CiscoDateTime     `json:"CustomerRequestedStartTime,omitempty"`
	Description                *string            `json:"Description,omitempty"` // Description of Ticket e.g "No internet connection on Location XXX, Router shows no green lights, Error Message YYY on interface". M for Create. M for Close.
	Diagnosis                  *string            `json:"Diagnosis,omitempty"`   // Customer Symptom. M for Close.
	MainComp                   *ComponentsHolder  `json:"MainComp,omitempty"`
	Notes                      []*CallNotesHolder `json:"Notes,omitempty"`
	Ownership                  *string            `json:"Ownership,omitempty"`
	PartnerCoreTicketId        *float32           `json:"PartnerCoreTicketId,omitempty"`
	ProblemStartTime           *CiscoDateTime     `json:"ProblemStartTime,omitempty"`
	ProviderCategory1          *string            `json:"ProviderCategory1,omitempty"`
	ProviderCategory2          *string            `json:"ProviderCategory2,omitempty"`
	ProviderCategory3          *string            `json:"ProviderCategory3,omitempty"`
	ProviderCategory4          *string            `json:"ProviderCategory4,omitempty"`
	ProviderCategory5          *string            `json:"ProviderCategory5,omitempty"`
	ProviderReasonCategory1    *string            `json:"ProviderReasonCategory1,omitempty"`
	ProviderReasonCategory2    *string            `json:"ProviderReasonCategory2,omitempty"`
	ProviderReasonCategory3    *string            `json:"ProviderReasonCategory3,omitempty"`
	ProviderReasonCategory4    *string            `json:"ProviderReasonCategory4,omitempty"`
	ProviderReasonCategory5    *string            `json:"ProviderReasonCategory5,omitempty"`
	ProviderScheduledEndTime   *CiscoDateTime     `json:"ProviderScheduledEndTime,omitempty"`
	ProviderScheduledStartTime *CiscoDateTime     `json:"ProviderScheduledStartTime,omitempty"`
	// Remarks                    *string            `json:"Remarks,omitempty"`          // Comments sent during ticket updates (e.g. work log) TODO: May be an array of strings too!!! So perhaps need to check for this.
	Remarks          *StringOrSliceOfErrors `json:"Remarks,omitempty"`          // Comments sent during ticket updates (e.g. work log) TODO: May be an array of strings too!!! So perhaps need to check for this.
	SDCallID         *string                `json:"SDCallID,omitempty"`         // ServiceGrid ticket ID e.g. 1340036031
	SPCallID         *string                `json:"SPCallID,omitempty"`         // Cisco TAC CSOne ticket ID e.g. 692072147
	ShortDescription *string                `json:"ShortDescription,omitempty"` // Subject of Ticket when Created e.g. "Network issue in Location XXX". M for Create.
	Solution         *string                `json:"Solution,omitempty"`         // The resolution notes sent over when ticket is RESOLVED. M on Update Resolved.
	SubComp          *ComponentsHolder      `json:"SubComp,omitempty"`
	SysSpecField1    *string                `json:"SysSpecField1,omitempty"`
	SysSpecField10   *string                `json:"SysSpecField10,omitempty"`
	SysSpecField2    *string                `json:"SysSpecField2,omitempty"`
	SysSpecField3    *string                `json:"SysSpecField3,omitempty"`
	SysSpecField4    *string                `json:"SysSpecField4,omitempty"`
	SysSpecField5    *string                `json:"SysSpecField5,omitempty"`
	SysSpecField6    *string                `json:"SysSpecField6,omitempty"`
	SysSpecField7    *string                `json:"SysSpecField7,omitempty"`
	SysSpecField8    *string                `json:"SysSpecField8,omitempty"`
	SysSpecField9    *string                `json:"SysSpecField9,omitempty"`
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
	ShortName         *string `json:"ShortName,omitempty"` // A unique value required in SG (will be provided by SG to PARTNER) e.g. 1190048357. M for Create.
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
	ShortName                   *string `json:"ShortName,omitempty"` // A unique value required in SG (will be provided by SG to PARTNER) e.g. 1190048450. M for Create.
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
	Department       *string `json:"Department,omitempty"` // For CCP, Software Product ID. M on Close.
	Description      *string `json:"Description,omitempty"`
	EMail            *string `json:"EMail,omitempty"` // For CCP, Underlying Cause. M on Close.
	Fax              *string `json:"Fax,omitempty"`   // For CCP, Resolution Code. M on Close.
	FirstName        *string `json:"FirstName,omitempty"`
	Language         *string `json:"Language,omitempty"`
	LastName         *string `json:"LastName,omitempty"`
	Location         *string `json:"Location,omitempty"`
	LocationCity     *string `json:"LocationCity,omitempty"` // For CCP, Complexity. M on Close.
	LocationCountry  *string `json:"LocationCountry,omitempty"`
	LocationProvince *string `json:"LocationProvince,omitempty"`
	LocationStreet   *string `json:"LocationStreet,omitempty"` // For CCP, Resolution Summary. M on Close.
	LocationZip      *string `json:"LocationZip,omitempty"`
	MobileTel        *string `json:"MobileTel,omitempty"`
	PIN              *string `json:"PIN,omitempty"`  // For CHD, CCO ID for the case used for entitlement. M for Create. For CCP, Hardware Product ID. M for Close.
	Room             *string `json:"Room,omitempty"` // For CCP, Subscription ID (if needed to pass entitlement).
	Salutation       *string `json:"Salutation,omitempty"`
	ShortName        *string `json:"ShortName,omitempty"`
	Sign             *string `json:"Sign,omitempty"` // For CHD, Preferred communication method. M for Create.
	Tel              *string `json:"Tel,omitempty"`
	Tel2             *string `json:"Tel2,omitempty"`
	Title            *string `json:"Title,omitempty"` // For Caller, this is apparently "Ship To Customer Company"?
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
