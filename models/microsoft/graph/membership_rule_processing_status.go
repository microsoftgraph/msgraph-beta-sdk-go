package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MembershipRuleProcessingStatus 
type MembershipRuleProcessingStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
    errorMessage *string;
    // Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
    lastMembershipUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.   Required. Read-only.
    status *MembershipRuleProcessingStatusDetails;
}
// NewMembershipRuleProcessingStatus instantiates a new membershipRuleProcessingStatus and sets the default values.
func NewMembershipRuleProcessingStatus()(*MembershipRuleProcessingStatus) {
    m := &MembershipRuleProcessingStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MembershipRuleProcessingStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorMessage gets the errorMessage property value. Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// GetLastMembershipUpdated gets the lastMembershipUpdated property value. Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) GetLastMembershipUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastMembershipUpdated
    }
}
// GetStatus gets the status property value. Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.   Required. Read-only.
func (m *MembershipRuleProcessingStatus) GetStatus()(*MembershipRuleProcessingStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MembershipRuleProcessingStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["lastMembershipUpdated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastMembershipUpdated(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMembershipRuleProcessingStatusDetails)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MembershipRuleProcessingStatusDetails))
        }
        return nil
    }
    return res
}
func (m *MembershipRuleProcessingStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MembershipRuleProcessingStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastMembershipUpdated", m.GetLastMembershipUpdated())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *MembershipRuleProcessingStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetErrorMessage sets the errorMessage property value. Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) SetErrorMessage(value *string)() {
    if m != nil {
        m.errorMessage = value
    }
}
// SetLastMembershipUpdated sets the lastMembershipUpdated property value. Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) SetLastMembershipUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastMembershipUpdated = value
    }
}
// SetStatus sets the status property value. Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.   Required. Read-only.
func (m *MembershipRuleProcessingStatus) SetStatus(value *MembershipRuleProcessingStatusDetails)() {
    if m != nil {
        m.status = value
    }
}
