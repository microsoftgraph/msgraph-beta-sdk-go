package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SideLoadingKey struct {
    Entity
    // Side Loading Key description displayed to the ITPro Admins..
    description *string;
    // Side Loading Key Name displayed to the ITPro Admins.
    displayName *string;
    // Side Loading Key Last Updated Date displayed to the ITPro Admins.
    lastUpdatedDateTime *string;
    // Side Loading Key Total Activation displayed to the ITPro Admins.
    totalActivation *int32;
    // Side Loading Key Value, it is 5x5 value, seperated by hiphens.
    value *string;
}
// Instantiates a new sideLoadingKey and sets the default values.
func NewSideLoadingKey()(*SideLoadingKey) {
    m := &SideLoadingKey{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Side Loading Key description displayed to the ITPro Admins..
func (m *SideLoadingKey) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Side Loading Key Name displayed to the ITPro Admins.
func (m *SideLoadingKey) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastUpdatedDateTime property value. Side Loading Key Last Updated Date displayed to the ITPro Admins.
func (m *SideLoadingKey) GetLastUpdatedDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the totalActivation property value. Side Loading Key Total Activation displayed to the ITPro Admins.
func (m *SideLoadingKey) GetTotalActivation()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalActivation
    }
}
// Gets the value property value. Side Loading Key Value, it is 5x5 value, seperated by hiphens.
func (m *SideLoadingKey) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *SideLoadingKey) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["totalActivation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalActivation(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *SideLoadingKey) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SideLoadingKey) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    {
        err = writer.WriteStringValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalActivation", m.GetTotalActivation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. Side Loading Key description displayed to the ITPro Admins..
// Parameters:
//  - value : Value to set for the description property.
func (m *SideLoadingKey) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Side Loading Key Name displayed to the ITPro Admins.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SideLoadingKey) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastUpdatedDateTime property value. Side Loading Key Last Updated Date displayed to the ITPro Admins.
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *SideLoadingKey) SetLastUpdatedDateTime(value *string)() {
    m.lastUpdatedDateTime = value
}
// Sets the totalActivation property value. Side Loading Key Total Activation displayed to the ITPro Admins.
// Parameters:
//  - value : Value to set for the totalActivation property.
func (m *SideLoadingKey) SetTotalActivation(value *int32)() {
    m.totalActivation = value
}
// Sets the value property value. Side Loading Key Value, it is 5x5 value, seperated by hiphens.
// Parameters:
//  - value : Value to set for the value property.
func (m *SideLoadingKey) SetValue(value *string)() {
    m.value = value
}
