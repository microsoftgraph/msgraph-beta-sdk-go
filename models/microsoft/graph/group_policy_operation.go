package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyOperation 
type GroupPolicyOperation struct {
    Entity
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The group policy operation status. Possible values are: unknown, inProgress, success, failed.
    operationStatus *GroupPolicyOperationStatus;
    // The type of group policy operation. Possible values are: none, upload, uploadNewVersion, addLanguageFiles, removeLanguageFiles, updateLanguageFiles, remove.
    operationType *GroupPolicyOperationType;
    // The group policy operation status detail.
    statusDetails *string;
}
// NewGroupPolicyOperation instantiates a new groupPolicyOperation and sets the default values.
func NewGroupPolicyOperation()(*GroupPolicyOperation) {
    m := &GroupPolicyOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyOperation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOperationStatus gets the operationStatus property value. The group policy operation status. Possible values are: unknown, inProgress, success, failed.
func (m *GroupPolicyOperation) GetOperationStatus()(*GroupPolicyOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.operationStatus
    }
}
// GetOperationType gets the operationType property value. The type of group policy operation. Possible values are: none, upload, uploadNewVersion, addLanguageFiles, removeLanguageFiles, updateLanguageFiles, remove.
func (m *GroupPolicyOperation) GetOperationType()(*GroupPolicyOperationType) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
// GetStatusDetails gets the statusDetails property value. The group policy operation status detail.
func (m *GroupPolicyOperation) GetStatusDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["operationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationStatus(val.(*GroupPolicyOperationStatus))
        }
        return nil
    }
    res["operationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*GroupPolicyOperationType))
        }
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val)
        }
        return nil
    }
    return res
}
func (m *GroupPolicyOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperationStatus() != nil {
        cast := (*m.GetOperationStatus()).String()
        err = writer.WriteStringValue("operationStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("statusDetails", m.GetStatusDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyOperation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOperationStatus sets the operationStatus property value. The group policy operation status. Possible values are: unknown, inProgress, success, failed.
func (m *GroupPolicyOperation) SetOperationStatus(value *GroupPolicyOperationStatus)() {
    if m != nil {
        m.operationStatus = value
    }
}
// SetOperationType sets the operationType property value. The type of group policy operation. Possible values are: none, upload, uploadNewVersion, addLanguageFiles, removeLanguageFiles, updateLanguageFiles, remove.
func (m *GroupPolicyOperation) SetOperationType(value *GroupPolicyOperationType)() {
    if m != nil {
        m.operationType = value
    }
}
// SetStatusDetails sets the statusDetails property value. The group policy operation status detail.
func (m *GroupPolicyOperation) SetStatusDetails(value *string)() {
    if m != nil {
        m.statusDetails = value
    }
}
