package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PolicySetItem 
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
// NewPolicySetItem instantiates a new policySetItem and sets the default values.
func NewPolicySetItem()(*PolicySetItem) {
    m := &PolicySetItem{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. Creation time of the PolicySetItem.
func (m *PolicySetItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. DisplayName of the PolicySetItem.
func (m *PolicySetItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetErrorCode gets the errorCode property value. Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
func (m *PolicySetItem) GetErrorCode()(*ErrorCode) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetGuidedDeploymentTags gets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySetItem) GetGuidedDeploymentTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.guidedDeploymentTags
    }
}
// GetItemType gets the itemType property value. policySetType of the PolicySetItem.
func (m *PolicySetItem) GetItemType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemType
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified time of the PolicySetItem.
func (m *PolicySetItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPayloadId gets the payloadId property value. PayloadId of the PolicySetItem.
func (m *PolicySetItem) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
// GetStatus gets the status property value. Status of the PolicySetItem. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
func (m *PolicySetItem) GetStatus()(*PolicySetStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicySetItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
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
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseErrorCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val.(*ErrorCode))
        }
        return nil
    }
    res["guidedDeploymentTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGuidedDeploymentTags(res)
        }
        return nil
    }
    res["itemType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemType(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicySetStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PolicySetStatus))
        }
        return nil
    }
    return res
}
func (m *PolicySetItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetErrorCode()).String()
        err = writer.WriteStringValue("errorCode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetGuidedDeploymentTags() != nil {
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
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Creation time of the PolicySetItem.
func (m *PolicySetItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. DisplayName of the PolicySetItem.
func (m *PolicySetItem) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetErrorCode sets the errorCode property value. Error code if any occured. Possible values are: noError, unauthorized, notFound, deleted.
func (m *PolicySetItem) SetErrorCode(value *ErrorCode)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetGuidedDeploymentTags sets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySetItem) SetGuidedDeploymentTags(value []string)() {
    if m != nil {
        m.guidedDeploymentTags = value
    }
}
// SetItemType sets the itemType property value. policySetType of the PolicySetItem.
func (m *PolicySetItem) SetItemType(value *string)() {
    if m != nil {
        m.itemType = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified time of the PolicySetItem.
func (m *PolicySetItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPayloadId sets the payloadId property value. PayloadId of the PolicySetItem.
func (m *PolicySetItem) SetPayloadId(value *string)() {
    if m != nil {
        m.payloadId = value
    }
}
// SetStatus sets the status property value. Status of the PolicySetItem. Possible values are: unknown, validating, partialSuccess, success, error, notAssigned.
func (m *PolicySetItem) SetStatus(value *PolicySetStatus)() {
    if m != nil {
        m.status = value
    }
}
