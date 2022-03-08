package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtectionable 
type InformationProtectionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBitlocker()(Bitlockerable)
    GetDataLossPreventionPolicies()([]DataLossPreventionPolicyable)
    GetPolicy()(InformationProtectionPolicyable)
    GetSensitivityLabels()([]SensitivityLabelable)
    GetSensitivityPolicySettings()(SensitivityPolicySettingsable)
    GetThreatAssessmentRequests()([]ThreatAssessmentRequestable)
    SetBitlocker(value Bitlockerable)()
    SetDataLossPreventionPolicies(value []DataLossPreventionPolicyable)()
    SetPolicy(value InformationProtectionPolicyable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
    SetSensitivityPolicySettings(value SensitivityPolicySettingsable)()
    SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)()
}
