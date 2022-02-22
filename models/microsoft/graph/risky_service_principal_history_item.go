package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyServicePrincipalHistoryItem 
type RiskyServicePrincipalHistoryItem struct {
    RiskyServicePrincipal
    // The activity related to service principal risk level change.
    activity *RiskServicePrincipalActivity;
    // The identifier of the actor of the operation.
    initiatedBy *string;
    // The identifier of the service principal.
    servicePrincipalId *string;
}
// NewRiskyServicePrincipalHistoryItem instantiates a new riskyServicePrincipalHistoryItem and sets the default values.
func NewRiskyServicePrincipalHistoryItem()(*RiskyServicePrincipalHistoryItem) {
    m := &RiskyServicePrincipalHistoryItem{
        RiskyServicePrincipal: *NewRiskyServicePrincipal(),
    }
    return m
}
// GetActivity gets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) GetActivity()(*RiskServicePrincipalActivity) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetInitiatedBy gets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) GetInitiatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
// GetServicePrincipalId gets the servicePrincipalId property value. The identifier of the service principal.
func (m *RiskyServicePrincipalHistoryItem) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyServicePrincipalHistoryItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.RiskyServicePrincipal.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskServicePrincipalActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(*RiskServicePrincipalActivity))
        }
        return nil
    }
    res["initiatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    return res
}
func (m *RiskyServicePrincipalHistoryItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RiskyServicePrincipalHistoryItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.RiskyServicePrincipal.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) SetActivity(value *RiskServicePrincipalActivity)() {
    if m != nil {
        m.activity = value
    }
}
// SetInitiatedBy sets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) SetInitiatedBy(value *string)() {
    if m != nil {
        m.initiatedBy = value
    }
}
// SetServicePrincipalId sets the servicePrincipalId property value. The identifier of the service principal.
func (m *RiskyServicePrincipalHistoryItem) SetServicePrincipalId(value *string)() {
    if m != nil {
        m.servicePrincipalId = value
    }
}
