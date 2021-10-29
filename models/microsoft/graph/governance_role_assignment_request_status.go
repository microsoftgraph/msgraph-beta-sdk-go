package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernanceRoleAssignmentRequestStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    status *string;
    // 
    statusDetails []KeyValue;
    // 
    subStatus *string;
}
// Instantiates a new governanceRoleAssignmentRequestStatus and sets the default values.
func NewGovernanceRoleAssignmentRequestStatus()(*GovernanceRoleAssignmentRequestStatus) {
    m := &GovernanceRoleAssignmentRequestStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRoleAssignmentRequestStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the status property value. 
func (m *GovernanceRoleAssignmentRequestStatus) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the statusDetails property value. 
func (m *GovernanceRoleAssignmentRequestStatus) GetStatusDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// Gets the subStatus property value. 
func (m *GovernanceRoleAssignmentRequestStatus) GetSubStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subStatus
    }
}
// The deserialization information for the current model
func (m *GovernanceRoleAssignmentRequestStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        res := make([]KeyValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValue))
        }
        m.SetStatusDetails(res)
        return nil
    }
    res["subStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubStatus(val)
        return nil
    }
    return res
}
func (m *GovernanceRoleAssignmentRequestStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GovernanceRoleAssignmentRequestStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStatusDetails()))
        for i, v := range m.GetStatusDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("statusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subStatus", m.GetSubStatus())
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
func (m *GovernanceRoleAssignmentRequestStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *GovernanceRoleAssignmentRequestStatus) SetStatus(value *string)() {
    m.status = value
}
// Sets the statusDetails property value. 
// Parameters:
//  - value : Value to set for the statusDetails property.
func (m *GovernanceRoleAssignmentRequestStatus) SetStatusDetails(value []KeyValue)() {
    m.statusDetails = value
}
// Sets the subStatus property value. 
// Parameters:
//  - value : Value to set for the subStatus property.
func (m *GovernanceRoleAssignmentRequestStatus) SetSubStatus(value *string)() {
    m.subStatus = value
}
