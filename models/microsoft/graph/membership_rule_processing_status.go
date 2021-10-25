package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MembershipRuleProcessingStatus struct {
    additionalData map[string]interface{};
    errorMessage *string;
    lastMembershipUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    status *MembershipRuleProcessingStatusDetails;
}
func NewMembershipRuleProcessingStatus()(*MembershipRuleProcessingStatus) {
    m := &MembershipRuleProcessingStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MembershipRuleProcessingStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MembershipRuleProcessingStatus) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
func (m *MembershipRuleProcessingStatus) GetLastMembershipUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastMembershipUpdated
    }
}
func (m *MembershipRuleProcessingStatus) GetStatus()(*MembershipRuleProcessingStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *MembershipRuleProcessingStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorMessage(val)
        return nil
    }
    res["lastMembershipUpdated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastMembershipUpdated(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMembershipRuleProcessingStatusDetails)
        if err != nil {
            return err
        }
        cast := val.(MembershipRuleProcessingStatusDetails)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *MembershipRuleProcessingStatus) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetStatus().String()
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
func (m *MembershipRuleProcessingStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MembershipRuleProcessingStatus) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
func (m *MembershipRuleProcessingStatus) SetLastMembershipUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastMembershipUpdated = value
}
func (m *MembershipRuleProcessingStatus) SetStatus(value *MembershipRuleProcessingStatusDetails)() {
    m.status = value
}
