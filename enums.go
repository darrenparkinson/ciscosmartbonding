package ciscosmartbonding

import (
	"encoding/json"
	"strings"
)

// CloseTicketComplexity provides an enum for the ticket complexity value
type CloseTicketComplexity int64

const (
	Level0Procedural CloseTicketComplexity = iota // assume zero value is fine
	Level1Basic
	Level2Advanced
	Level3ExceptionallyComplex
)

var (
	CloseTicketComplexity_name = map[int64]string{
		0: "0 Level -Procedural",
		1: "1 Level -Basic",
		2: "2 Level -Advanced",
		3: "3 Level -Exceptionally Complex",
	}
	CloseTicketComplexity_value = map[string]int64{
		"0 Level -Procedural":            0,
		"1 Level -Basic":                 1,
		"2 Level -Advanced":              2,
		"3 Level -Exceptionally Complex": 3,
	}
)

func (c CloseTicketComplexity) String() string {
	return CloseTicketComplexity_name[int64(c)]
}

func (c CloseTicketComplexity) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CloseTicketComplexity) UnmarshalJSON(data []byte) error {
	var err error
	var complexity string
	if err := json.Unmarshal(data, &complexity); err != nil {
		return err
	}
	if *c, err = ParseComplexity(complexity); err != nil {
		return err
	}
	return nil
}

func ParseComplexity(c string) (CloseTicketComplexity, error) {
	c = strings.TrimSpace(c)
	val, ok := CloseTicketComplexity_value[c]
	if !ok {
		return CloseTicketComplexity(0), nil
		// return CloseTicketComplexity(0), fmt.Errorf("%q is not a valid CloseTicketComplexity", c)
	}
	return CloseTicketComplexity(val), nil
}

// ClostTicketResolutionCode provides an enum for the ticket resolution code when closing a ticket
type CloseTicketResolutionCode int64

const (
	ResolutionCode_Undefined = iota // to catch zero values
	ResolutionCode_ClosedNotEntitled
	ResolutionCode_CustomerEducation
	ResolutionCode_DocumentationProvided
	ResolutionCode_Duplicate
	ResolutionCode_HardwareReplacementRMASVO
	ResolutionCode_HardwareUpgradeExistingDefect
	ResolutionCode_HardwareUpgradeNewDefect
	ResolutionCode_HardwareUpgradeNewFeatureFunctionality
	ResolutionCode_LicenseIssues
	ResolutionCode_NetworkRedesign
	ResolutionCode_NoResponseFromCustomer
	ResolutionCode_OpenedInError
	ResolutionCode_PartnerSolved
	ResolutionCode_Redirect3rdPartySupport
	ResolutionCode_RedirectPartnerResellerSupport
	ResolutionCode_ResolvedByCustomer
	ResolutionCode_SoftwareConfiguration
	ResolutionCode_SoftwareUpgradeExistingDefect
	ResolutionCode_SoftwareUpgradeNewDefect
	ResolutionCode_SoftwareUpgradeNewFeatureFunctionality
	ResolutionCode_UnreproducibleProblem
	ResolutionCode_UnresolvedBug
	ResolutionCode_NewUserAccess
	ResolutionCode_CustomerActivation
	ResolutionCode_Reports
	ResolutionCode_SandboxUpload
	ResolutionCode_ApplicationErrorWPC
	ResolutionCode_ContentDelivered
	ResolutionCode_NewFeatureRequest
	ResolutionCode_ContentRequest
	ResolutionCode_ContentNotAvailable
	ResolutionCode_ResetPassword
	ResolutionCode_MobileIOS
	ResolutionCode_MobileAndroid
	ResolutionCode_ViewerInstalled
	ResolutionCode_ViewerOnline
	ResolutionCode_ApplicationErrorMAC
)

var (
	CloseTicketResolutionCode_name = map[int64]string{
		0:  "Undefined",
		1:  "Closed-Not Entitled",
		2:  "Customer Education",
		3:  "Documentation Provided",
		4:  "Duplicate",
		5:  "Hardware Replacement (RMA/SVO)",
		6:  "Hardware Upgrade - Existing Defect",
		7:  "Hardware Upgrade - New Defect",
		8:  "Hardware Upgrade - New Feature/Functionality",
		9:  "License Issues",
		10: "Network Redesign",
		11: "No Response from Customer",
		12: "Opened In Error",
		13: "PARTNER_SOLVED",
		14: "Redirect-3rd Party Support",
		15: "Redirect-Partner/Reseller Support",
		16: "Resolv_By_Cust",
		17: "Software Configuration",
		18: "Software Upgrade - Existing Defect",
		19: "Software Upgrade - New Defect",
		20: "Software Upgrade - New Feature/Functionality",
		21: "Unreproducible Problem",
		22: "Unresolved Bug",
		23: "New user Access",
		24: "Customer Activation",
		25: "Reports",
		26: "Sandbox Upload",
		27: "Application error - WPC",
		28: "Content Delivered",
		29: "New Feature Request",
		30: "Content Request",
		31: "Content not Available",
		32: "Reset Password",
		33: "Mobile - IOS",
		34: "Mobile - Android",
		35: "Viewer - Installed",
		36: "Viewer - Online",
		37: "Application error - MAC",
	}
	CloseTicketResolutionCode_value = map[string]int64{
		"Undefined":                                    0,
		"Closed-Not Entitled":                          1,
		"Customer Education":                           2,
		"Documentation Provided":                       3,
		"Duplicate":                                    4,
		"Hardware Replacement (RMA/SVO)":               5,
		"Hardware Upgrade - Existing Defect":           6,
		"Hardware Upgrade - New Defect":                7,
		"Hardware Upgrade - New Feature/Functionality": 8,
		"License Issues":                               9,
		"Network Redesign":                             10,
		"No Response from Customer":                    11,
		"Opened In Error":                              12,
		"PARTNER_SOLVED":                               13,
		"Redirect-3rd Party Support":                   14,
		"Redirect-Partner/Reseller Support":            15,
		"Resolv_By_Cust":                               16,
		"Software Configuration":                       17,
		"Software Upgrade - Existing Defect":           18,
		"Software Upgrade - New Defect":                19,
		"Software Upgrade - New Feature/Functionality": 20,
		"Unreproducible Problem":                       21,
		"Unresolved Bug":                               22,
		"New user Access":                              23,
		"Customer Activation":                          24,
		"Reports":                                      25,
		"Sandbox Upload":                               26,
		"Application error - WPC":                      27,
		"Content Delivered":                            28,
		"New Feature Request":                          29,
		"Content Request":                              30,
		"Content not Available":                        31,
		"Reset Password":                               32,
		"Mobile - IOS":                                 33,
		"Mobile - Android":                             34,
		"Viewer - Installed":                           35,
		"Viewer - Online":                              36,
		"Application error - MAC":                      37,
	}
)

func (r CloseTicketResolutionCode) String() string {
	return CloseTicketResolutionCode_name[int64(r)]
}

func (r CloseTicketResolutionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *CloseTicketResolutionCode) UnmarshalJSON(data []byte) error {
	var err error
	var resolution string
	if err := json.Unmarshal(data, &resolution); err != nil {
		return err
	}
	if *r, err = ParseResolution(resolution); err != nil {
		return err
	}
	return nil
}

func ParseResolution(r string) (CloseTicketResolutionCode, error) {
	r = strings.TrimSpace(r)
	val, ok := CloseTicketResolutionCode_value[r]
	if !ok {
		return CloseTicketResolutionCode(0), nil
		// return CloseTicketResolutionCode(0), fmt.Errorf("%q is not a valid CloseTicketResolutionCode", r)
	}
	return CloseTicketResolutionCode(val), nil
}

// CloseTicketUnderlyingCause provides an enum for the ticket underlying cause value on closure
type CloseTicketUnderlyingCause int64

const (
	UnderlyingCause_UnknownCauseOther CloseTicketUnderlyingCause = iota // (Use ONLY if none of the above apply)
	UnderlyingCause_Licensing
	UnderlyingCause_HardwareFailure
	UnderlyingCause_HardwareNonFailure // (limits exceeded, not enough memory, missing parts...)
	UnderlyingCause_SoftwareBug
	UnderlyingCause_SoftwareNotABug                 //(scalability, version selection, install/upgrade help...)
	UnderlyingCause_DocumentationTools              // (incomplete, too complex...)
	UnderlyingCause_DebugDiagnosticCapabilities     // (missing, incomplete, cryptic...)
	UnderlyingCause_DesignAssistanceNeeded          // (best practices, deployment advice, redesign...)
	UnderlyingCause_ConfigurationAssistance         // (process not intuitive, too complex, inconsistent...)
	UnderlyingCause_UsabilityOtherThanConfig        // (product hard to use, no console port...)
	UnderlyingCause_InteroperabilityCompatibility   // (Cisco to Cisco or Cisco to 3rd Party)
	UnderlyingCause_NonCiscoProductOrServiceProblem // (third party failure, telco...)
	UnderlyingCause_ExternalEnvironmentIssue        // (power outage, heat, lightning...)
)

var (
	CloseTicketUnderlyingCause_name = map[int64]string{
		0:  "Unknown Cause/Other (Use ONLY if none of the above apply)",
		1:  "Licensing",
		2:  "Hardware Failure",
		3:  "Hardware - non-failure (limits exceeded, not enough memory, missing parts...)",
		4:  "Software Bug",
		5:  "Software -not a bug (scalability, version selection, install/upgrade help...)",
		6:  "Documentation/Tools (incomplete, too complex...)",
		7:  "Debug/Diagnostic Capability (missing, incomplete, cryptic...)",
		8:  "Design Assistance Needed (best practices, deployment advice, redesign...)",
		9:  "Configuration Assistance (process not intuitive, too complex, inconsistent...)",
		10: "Usability -other than config (product hard to use, no console port...)",
		11: "Interoperability/Compatibility (Cisco to Cisco or Cisco to 3rd Party)",
		12: "Non-Cisco product or service problem (third party failure, telco...)",
		13: "External Environment Issue (power outage, heat, lightning...)",
	}
	CloseTicketUnderlyingCause_value = map[string]int64{
		"Unknown Cause/Other": 0,
		"Unknown Cause/Other (Use ONLY if none of the above apply)": 0,
		"Licensing":              1,
		"Hardware Failure":       2,
		"Hardware - non-failure": 3,
		"Hardware - non-failure (limits exceeded, not enough memory, missing parts...)": 3,
		"Software Bug":        4,
		"Software -not a bug": 5,
		"Software -not a bug (scalability, version selection, install/upgrade help...)": 5,
		"Documentation/Tools":                                                            6,
		"Documentation/Tools (incomplete, too complex...)":                               6,
		"Debug/Diagnostic Capabilities":                                                  7,
		"Debug/Diagnostic Capability (missing, incomplete, cryptic...)":                  7,
		"Design Assistance Needed":                                                       8,
		"Design Assistance Needed (best practices, deployment advice, redesign...)":      8,
		"Configuration Assistance":                                                       9,
		"Configuration Assistance (process not intuitive, too complex, inconsistent...)": 9,
		"Usability -other than config":                                                   10,
		"Usability -other than config (product hard to use, no console port...)":         10,
		"Interoperability/Compatibility":                                                 11,
		"Interoperability/Compatibility (Cisco to Cisco or Cisco to 3rd Party)":          11,
		"Non-Cisco product or service problem":                                           12,
		"Non-Cisco product or service problem (third party failure, telco...)":           12,
		"External Environment issue":                                                     13,
		"External Environment Issue (power outage, heat, lightning...)":                  13,
	}
)

func (u CloseTicketUnderlyingCause) String() string {
	return CloseTicketUnderlyingCause_name[int64(u)]
}

func (u CloseTicketUnderlyingCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *CloseTicketUnderlyingCause) UnmarshalJSON(data []byte) error {
	var err error
	var cause string
	if err := json.Unmarshal(data, &cause); err != nil {
		return err
	}
	if *u, err = ParseCause(cause); err != nil {
		return err
	}
	return nil
}

func ParseCause(u string) (CloseTicketUnderlyingCause, error) {
	u = strings.TrimSpace(u)
	val, ok := CloseTicketUnderlyingCause_value[u]
	if !ok {
		return CloseTicketUnderlyingCause(0), nil
		// return CloseTicketUnderlyingCause(0), fmt.Errorf("%q is not a valid CloseTicketUnderlyingCause", u)
	}
	return CloseTicketUnderlyingCause(val), nil
}
