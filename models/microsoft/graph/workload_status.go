package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type WorkloadStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The display name for the workload. Required. Read-only.
    displayName *string;
    // The date and time the workload was offboarded. Optional. Read-only.
    offboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time the workload was onboarded. Optional. Read-only.
    onboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The onboarding status for the workload. Possible values are: notOnboarded, onboarded, unknownFutureValue. Optional. Read-only.
    onboardingStatus *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadOnboardingStatus;
}
// Instantiates a new workloadStatus and sets the default values.
func NewWorkloadStatus()(*WorkloadStatus) {
    m := &WorkloadStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. The display name for the workload. Required. Read-only.
func (m *WorkloadStatus) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the offboardedDateTime property value. The date and time the workload was offboarded. Optional. Read-only.
func (m *WorkloadStatus) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.offboardedDateTime
    }
}
// Gets the onboardedDateTime property value. The date and time the workload was onboarded. Optional. Read-only.
func (m *WorkloadStatus) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onboardedDateTime
    }
}
// Gets the onboardingStatus property value. The onboarding status for the workload. Possible values are: notOnboarded, onboarded, unknownFutureValue. Optional. Read-only.
func (m *WorkloadStatus) GetOnboardingStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadOnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// The deserialization information for the current model
func (m *WorkloadStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["offboardedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedDateTime(val)
        }
        return nil
    }
    res["onboardedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseWorkloadOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadOnboardingStatus)
            m.SetOnboardingStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *WorkloadStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkloadStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("offboardedDateTime", m.GetOffboardedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("onboardedDateTime", m.GetOnboardedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := m.GetOnboardingStatus().String()
        err := writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkloadStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. The display name for the workload. Required. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WorkloadStatus) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the offboardedDateTime property value. The date and time the workload was offboarded. Optional. Read-only.
// Parameters:
//  - value : Value to set for the offboardedDateTime property.
func (m *WorkloadStatus) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.offboardedDateTime = value
}
// Sets the onboardedDateTime property value. The date and time the workload was onboarded. Optional. Read-only.
// Parameters:
//  - value : Value to set for the onboardedDateTime property.
func (m *WorkloadStatus) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onboardedDateTime = value
}
// Sets the onboardingStatus property value. The onboarding status for the workload. Possible values are: notOnboarded, onboarded, unknownFutureValue. Optional. Read-only.
// Parameters:
//  - value : Value to set for the onboardingStatus property.
func (m *WorkloadStatus) SetOnboardingStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadOnboardingStatus)() {
    m.onboardingStatus = value
}
