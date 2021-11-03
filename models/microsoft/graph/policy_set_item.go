package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PolicySetItem struct {
    Entity
    // Creation time of the PolicySetItem.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // DisplayName of the PolicySetItem.
    displayName *string;
    // Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
    errorCode *ErrorCode;
    // Tags of the guided deployment
    guidedDeploymentTags []string;
    // policySetType of the PolicySetItem.
    itemType *string;
    // Last modified time of the PolicySetItem.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // PayloadId of the PolicySetItem.
    payloadId *string;
    // Status of the PolicySetItem. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
    status *PolicySetStatus;
}
// Instantiates a new policySetItem and sets the default values.
func NewPolicySetItem()(*PolicySetItem) {
    m := &PolicySetItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Creation time of the PolicySetItem.
func (m *PolicySetItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. DisplayName of the PolicySetItem.
func (m *PolicySetItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the errorCode property value. Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
func (m *PolicySetItem) GetErrorCode()(*ErrorCode) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySetItem) GetGuidedDeploymentTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.guidedDeploymentTags
    }
}
// Gets the itemType property value. policySetType of the PolicySetItem.
func (m *PolicySetItem) GetItemType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemType
    }
}
// Gets the lastModifiedDateTime property value. Last modified time of the PolicySetItem.
func (m *PolicySetItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the payloadId property value. PayloadId of the PolicySetItem.
func (m *PolicySetItem) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
// Gets the status property value. Status of the PolicySetItem. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
func (m *PolicySetItem) GetStatus()(*PolicySetStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *PolicySetItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseErrorCode)
        if err != nil {
            return err
        }
        cast := val.(ErrorCode)
        m.SetErrorCode(&cast)
        return nil
    }
    res["guidedDeploymentTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetGuidedDeploymentTags(res)
        return nil
    }
    res["itemType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetItemType(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayloadId(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicySetStatus)
        if err != nil {
            return err
        }
        cast := val.(PolicySetStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *PolicySetItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PolicySetItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCode() != nil {
        cast := m.GetErrorCode().String()
        err = writer.WriteStringValue("errorCode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("guidedDeploymentTags", m.GetGuidedDeploymentTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("itemType", m.GetItemType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. Creation time of the PolicySetItem.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *PolicySetItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. DisplayName of the PolicySetItem.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PolicySetItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the errorCode property value. Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *PolicySetItem) SetErrorCode(value *ErrorCode)() {
    m.errorCode = value
}
// Sets the guidedDeploymentTags property value. Tags of the guided deployment
// Parameters:
//  - value : Value to set for the guidedDeploymentTags property.
func (m *PolicySetItem) SetGuidedDeploymentTags(value []string)() {
    m.guidedDeploymentTags = value
}
// Sets the itemType property value. policySetType of the PolicySetItem.
// Parameters:
//  - value : Value to set for the itemType property.
func (m *PolicySetItem) SetItemType(value *string)() {
    m.itemType = value
}
// Sets the lastModifiedDateTime property value. Last modified time of the PolicySetItem.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *PolicySetItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the payloadId property value. PayloadId of the PolicySetItem.
// Parameters:
//  - value : Value to set for the payloadId property.
func (m *PolicySetItem) SetPayloadId(value *string)() {
    m.payloadId = value
}
// Sets the status property value. Status of the PolicySetItem. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
// Parameters:
//  - value : Value to set for the status property.
func (m *PolicySetItem) SetStatus(value *PolicySetStatus)() {
    m.status = value
}
