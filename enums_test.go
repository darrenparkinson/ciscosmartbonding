package ciscosmartbonding

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestComplexityEnum(t *testing.T) {
	var tests = []struct {
		name string
		enum CloseTicketComplexity
		text string
	}{
		{"procedural", Level0Procedural, "0 Level -Procedural"},
		{"basic", Level1Basic, "1 Level -Basic"},
		{"advanced", Level2Advanced, "2 Level -Advanced"},
		{"complex", Level3ExceptionallyComplex, "3 Level -Exceptionally Complex"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("marshal_%s", tt.name)
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tt.enum)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			want, err := json.Marshal(tt.text)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			if string(got) != string(want) {
				t.Errorf("got %+v, want %+v", string(got), string(want))
			}
		})
	}
	for _, tt := range tests {
		name := fmt.Sprintf("unmarshal_%s", tt.name)
		t.Run(name, func(t *testing.T) {
			text, _ := json.Marshal(tt.text)
			var got CloseTicketComplexity
			err := json.Unmarshal(text, &got)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			if got != tt.enum {
				t.Errorf("got %+v, want %+v", got, tt.enum)
			}
		})
	}
}

func TestResolutionCodeEnum(t *testing.T) {
	var tests = []struct {
		name string
		enum CloseTicketResolutionCode
		text string
	}{
		{"Undefined", ResolutionCode_Undefined, "Undefined"},
		{"Closed-Not Entitled", ResolutionCode_ClosedNotEntitled, "Closed-Not Entitled"},
		{"Customer Education", ResolutionCode_CustomerEducation, "Customer Education"},
		{"Documentation Provided", ResolutionCode_DocumentationProvided, "Documentation Provided"},
		{"Duplicate", ResolutionCode_Duplicate, "Duplicate"},
		{"Hardware Replacement (RMA/SVO)", ResolutionCode_HardwareReplacementRMASVO, "Hardware Replacement (RMA/SVO)"},
		{"Hardware Upgrade - Existing Defect", ResolutionCode_HardwareUpgradeExistingDefect, "Hardware Upgrade - Existing Defect"},
		{"Hardware Upgrade - New Defect", ResolutionCode_HardwareUpgradeNewDefect, "Hardware Upgrade - New Defect"},
		{"Hardware Upgrade - New Feature/Functionality", ResolutionCode_HardwareUpgradeNewFeatureFunctionality, "Hardware Upgrade - New Feature/Functionality"},
		{"License Issues", ResolutionCode_LicenseIssues, "License Issues"},
		{"Network Redesign", ResolutionCode_NetworkRedesign, "Network Redesign"},
		{"No Response from Customer", ResolutionCode_NoResponseFromCustomer, "No Response from Customer"},
		{"Opened In Error", ResolutionCode_OpenedInError, "Opened In Error"},
		{"PARTNER_SOLVED", ResolutionCode_PartnerSolved, "PARTNER_SOLVED"},
		{"Redirect-3rd Party Support", ResolutionCode_Redirect3rdPartySupport, "Redirect-3rd Party Support"},
		{"Redirect-Partner/Reseller Support", ResolutionCode_RedirectPartnerResellerSupport, "Redirect-Partner/Reseller Support"},
		{"Resolv_By_Cust", ResolutionCode_ResolvedByCustomer, "Resolv_By_Cust"},
		{"Software Configuration", ResolutionCode_SoftwareConfiguration, "Software Configuration"},
		{"Software Upgrade - Existing Defect", ResolutionCode_SoftwareUpgradeExistingDefect, "Software Upgrade - Existing Defect"},
		{"Software Upgrade - New Defect", ResolutionCode_SoftwareUpgradeNewDefect, "Software Upgrade - New Defect"},
		{"Software Upgrade - New Feature/Functionality", ResolutionCode_SoftwareUpgradeNewFeatureFunctionality, "Software Upgrade - New Feature/Functionality"},
		{"Unreproducible Problem", ResolutionCode_UnreproducibleProblem, "Unreproducible Problem"},
		{"Unresolved Bug", ResolutionCode_UnresolvedBug, "Unresolved Bug"},
		{"New user Access", ResolutionCode_NewUserAccess, "New user Access"},
		{"Customer Activation", ResolutionCode_CustomerActivation, "Customer Activation"},
		{"Reports", ResolutionCode_Reports, "Reports"},
		{"Sandbox Upload", ResolutionCode_SandboxUpload, "Sandbox Upload"},
		{"Application error - WPC", ResolutionCode_ApplicationErrorWPC, "Application error - WPC"},
		{"Content Delivered", ResolutionCode_ContentDelivered, "Content Delivered"},
		{"New Feature Request", ResolutionCode_NewFeatureRequest, "New Feature Request"},
		{"Content Request", ResolutionCode_ContentRequest, "Content Request"},
		{"Content not Available", ResolutionCode_ContentNotAvailable, "Content not Available"},
		{"Reset Password", ResolutionCode_ResetPassword, "Reset Password"},
		{"Mobile - IOS", ResolutionCode_MobileIOS, "Mobile - IOS"},
		{"Mobile - Android", ResolutionCode_MobileAndroid, "Mobile - Android"},
		{"Viewer - Installed", ResolutionCode_ViewerInstalled, "Viewer - Installed"},
		{"Viewer - Online", ResolutionCode_ViewerOnline, "Viewer - Online"},
		{"Application error - MAC", ResolutionCode_ApplicationErrorMAC, "Application error - MAC"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("marshal_%s", tt.name)
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tt.enum)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			want, err := json.Marshal(tt.text)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			if string(got) != string(want) {
				t.Errorf("got %+v, want %+v", string(got), string(want))
			}
		})
	}
	for _, tt := range tests {
		name := fmt.Sprintf("unmarshal_%s", tt.name)
		t.Run(name, func(t *testing.T) {
			text, _ := json.Marshal(tt.text)
			var got CloseTicketResolutionCode
			err := json.Unmarshal(text, &got)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			if got != tt.enum {
				t.Errorf("got %+v, want %+v", got, tt.enum)
			}
		})
	}
}

func TestUnderlyingCauseEnum(t *testing.T) {
	var tests = []struct {
		name string
		enum CloseTicketUnderlyingCause
		text string
	}{
		{"Unknown Cause/Other", UnderlyingCause_UnknownCauseOther, "Unknown Cause/Other"},
		{"Licensing", UnderlyingCause_Licensing, "Licensing"},
		{"Hardware Failure", UnderlyingCause_HardwareFailure, "Hardware Failure"},
		{"Hardware - non-failure", UnderlyingCause_HardwareNonFailure, "Hardware - non-failure"},
		{"Software Bug", UnderlyingCause_SoftwareBug, "Software Bug"},
		{"Software -not a bug", UnderlyingCause_SoftwareNotABug, "Software -not a bug"},
		{"Documentation/Tools", UnderlyingCause_DocumentationTools, "Documentation/Tools"},
		{"Debug/Diagnostic Capabilities", UnderlyingCause_DebugDiagnosticCapabilities, "Debug/Diagnostic Capabilities"},
		{"Design Assistance Needed", UnderlyingCause_DesignAssistanceNeeded, "Design Assistance Needed"},
		{"Configuration Assistance", UnderlyingCause_ConfigurationAssistance, "Configuration Assistance"},
		{"Usability -other than config", UnderlyingCause_UsabilityOtherThanConfig, "Usability -other than config"},
		{"Interoperability/Compatibility", UnderlyingCause_InteroperabilityCompatibility, "Interoperability/Compatibility"},
		{"Non-Cisco product or service problem", UnderlyingCause_NonCiscoProductOrServiceProblem, "Non-Cisco product or service problem"},
		{"External Environment issue", UnderlyingCause_ExternalEnvironmentIssue, "External Environment issue"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("marshal_%s", tt.name)
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tt.enum)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			want, err := json.Marshal(tt.text)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			if string(got) != string(want) {
				t.Errorf("got %+v, want %+v", string(got), string(want))
			}
		})
	}
	for _, tt := range tests {
		name := fmt.Sprintf("unmarshal_%s", tt.name)
		t.Run(name, func(t *testing.T) {
			text, _ := json.Marshal(tt.text)
			var got CloseTicketUnderlyingCause
			err := json.Unmarshal(text, &got)
			if err != nil {
				t.Fatalf("received unexpected error: %v", err)
			}
			if got != tt.enum {
				t.Errorf("got %+v, want %+v", got, tt.enum)
			}
		})
	}
}
