package ciscosmartbonding

// CloseTicketComplexity provides an enum for the ticket complexity value
type CloseTicketComplexity int64

const (
	Level0Procedural CloseTicketComplexity = iota
	Level1Basic
	Level2Advanced
	Level3ExceptionallyComplex
)

func (c CloseTicketComplexity) String() string {
	switch c {
	case Level0Procedural:
		return "0 Level -Procedural"
	case Level1Basic:
		return "1 Level -Basic"
	case Level2Advanced:
		return "2 Level -Advanced"
	case Level3ExceptionallyComplex:
		return "3 Level -Exceptionally Complex"
	}
	return "unknown"
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

func (r CloseTicketResolutionCode) String() string {
	switch r {
	case ResolutionCode_ClosedNotEntitled:
		return "Closed-Not Entitled"
	case ResolutionCode_CustomerEducation:
		return "Customer Education"
	case ResolutionCode_DocumentationProvided:
		return "Documentation Provided"
	case ResolutionCode_Duplicate:
		return "Duplicate"
	case ResolutionCode_HardwareReplacementRMASVO:
		return "Hardware Replacement (RMA/SVO)"
	case ResolutionCode_HardwareUpgradeExistingDefect:
		return "Hardware Upgrade - Existing Defect"
	case ResolutionCode_HardwareUpgradeNewDefect:
		return "Hardware Upgrade - New Defect"
	case ResolutionCode_HardwareUpgradeNewFeatureFunctionality:
		return "Hardware Upgrade - New Feature/Functionality"
	case ResolutionCode_LicenseIssues:
		return "License Issues"
	case ResolutionCode_NetworkRedesign:
		return "Network Redesign"
	case ResolutionCode_NoResponseFromCustomer:
		return "No Response from Customer"
	case ResolutionCode_OpenedInError:
		return "Opened In Error"
	case ResolutionCode_PartnerSolved:
		return "PARTNER_SOLVED"
	case ResolutionCode_Redirect3rdPartySupport:
		return "Redirect-3rd Party Support"
	case ResolutionCode_RedirectPartnerResellerSupport:
		return "Redirect-Partner/Reseller Support"
	case ResolutionCode_ResolvedByCustomer:
		return "Resolv_By_Cust"
	case ResolutionCode_SoftwareConfiguration:
		return "Software Configuration"
	case ResolutionCode_SoftwareUpgradeExistingDefect:
		return "Software Upgrade - Existing Defect"
	case ResolutionCode_SoftwareUpgradeNewDefect:
		return "Software Upgrade - New Defect"
	case ResolutionCode_SoftwareUpgradeNewFeatureFunctionality:
		return "Software Upgrade - New Feature/Functionality"
	case ResolutionCode_UnreproducibleProblem:
		return "Unreproducible Problem"
	case ResolutionCode_UnresolvedBug:
		return "Unresolved Bug"
	case ResolutionCode_NewUserAccess:
		return "New user Access"
	case ResolutionCode_CustomerActivation:
		return "Customer Activation"
	case ResolutionCode_Reports:
		return "Reports"
	case ResolutionCode_SandboxUpload:
		return "Sandbox Upload"
	case ResolutionCode_ApplicationErrorWPC:
		return "Application error - WPC"
	case ResolutionCode_ContentDelivered:
		return "Content Delivered"
	case ResolutionCode_NewFeatureRequest:
		return "New Feature Request"
	case ResolutionCode_ContentRequest:
		return "Content Request"
	case ResolutionCode_ContentNotAvailable:
		return "Content not Available"
	case ResolutionCode_ResetPassword:
		return "Reset Password"
	case ResolutionCode_MobileIOS:
		return "Mobile - IOS"
	case ResolutionCode_MobileAndroid:
		return "Mobile - Android"
	case ResolutionCode_ViewerInstalled:
		return "Viewer - Installed"
	case ResolutionCode_ViewerOnline:
		return "Viewer - Online"
	case ResolutionCode_ApplicationErrorMAC:
		return "Application error - MAC"
	}
	return "unknown"
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

func (u CloseTicketUnderlyingCause) String() string {
	switch u {
	case UnderlyingCause_Licensing:
		return "Licensing"
	case UnderlyingCause_HardwareFailure:
		return "Hardware Failure"
	case UnderlyingCause_HardwareNonFailure:
		return "Hardware - non-failure"
	case UnderlyingCause_SoftwareBug:
		return "Software Bug"
	case UnderlyingCause_SoftwareNotABug:
		return "Software -not a bug"
	case UnderlyingCause_DocumentationTools:
		return "Documentation/Tools"
	case UnderlyingCause_DebugDiagnosticCapabilities:
		return "Debug/Diagnostic Capabilities"
	case UnderlyingCause_DesignAssistanceNeeded:
		return "Design Assistance Needed"
	case UnderlyingCause_ConfigurationAssistance:
		return "Configuration Assistance"
	case UnderlyingCause_UsabilityOtherThanConfig:
		return "Usability -other than config"
	case UnderlyingCause_InteroperabilityCompatibility:
		return "Interoperability/Compatibility"
	case UnderlyingCause_NonCiscoProductOrServiceProblem:
		return "Non-Cisco product or service problem"
	case UnderlyingCause_ExternalEnvironmentIssue:
		return "External Environment issue"
	case UnderlyingCause_UnknownCauseOther:
		return "Unknown Cause/Other"
	}
	return "unknown"
}
