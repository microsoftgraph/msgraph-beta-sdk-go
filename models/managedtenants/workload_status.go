package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkloadStatus 
type WorkloadStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name for the workload. Required. Read-only.
    displayName *string
    // The date and time the workload was offboarded. Optional. Read-only.
    offboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time the workload was onboarded. Optional. Read-only.
    onboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The onboarding status for the workload. Possible values are: notOnboarded, onboarded, unknownFutureValue. Optional. Read-only.
    onboardingStatus *WorkloadOnboardingStatus
}
// NewWorkloadStatus instantiates a new workloadStatus and sets the default values.
func NewWorkloadStatus()(*WorkloadStatus) {
    m := &WorkloadStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkloadStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkloadStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkloadStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The display name for the workload. Required. Read-only.
func (m *WorkloadStatus) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkloadStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["offboardedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedDateTime(val)
        }
        return nil
    }
    res["onboardedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkloadOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*WorkloadOnboardingStatus))
        }
        return nil
    }
    return res
}
// GetOffboardedDateTime gets the offboardedDateTime property value. The date and time the workload was offboarded. Optional. Read-only.
func (m *WorkloadStatus) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.offboardedDateTime
    }
}
// GetOnboardedDateTime gets the onboardedDateTime property value. The date and time the workload was onboarded. Optional. Read-only.
func (m *WorkloadStatus) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onboardedDateTime
    }
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboarding status for the workload. Possible values are: notOnboarded, onboarded, unknownFutureValue. Optional. Read-only.
func (m *WorkloadStatus) GetOnboardingStatus()(*WorkloadOnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// Serialize serializes information the current object
func (m *WorkloadStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetOnboardingStatus()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the workload. Required. Read-only.
func (m *WorkloadStatus) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetOffboardedDateTime sets the offboardedDateTime property value. The date and time the workload was offboarded. Optional. Read-only.
func (m *WorkloadStatus) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.offboardedDateTime = value
    }
}
// SetOnboardedDateTime sets the onboardedDateTime property value. The date and time the workload was onboarded. Optional. Read-only.
func (m *WorkloadStatus) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.onboardedDateTime = value
    }
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboarding status for the workload. Possible values are: notOnboarded, onboarded, unknownFutureValue. Optional. Read-only.
func (m *WorkloadStatus) SetOnboardingStatus(value *WorkloadOnboardingStatus)() {
    if m != nil {
        m.onboardingStatus = value
    }
}
