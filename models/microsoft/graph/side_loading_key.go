package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SideLoadingKey 
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
// NewSideLoadingKey instantiates a new sideLoadingKey and sets the default values.
func NewSideLoadingKey()(*SideLoadingKey) {
    m := &SideLoadingKey{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Side Loading Key description displayed to the ITPro Admins..
func (m *SideLoadingKey) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Side Loading Key Name displayed to the ITPro Admins.
func (m *SideLoadingKey) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Side Loading Key Last Updated Date displayed to the ITPro Admins.
func (m *SideLoadingKey) GetLastUpdatedDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetTotalActivation gets the totalActivation property value. Side Loading Key Total Activation displayed to the ITPro Admins.
func (m *SideLoadingKey) GetTotalActivation()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalActivation
    }
}
// GetValue gets the value property value. Side Loading Key Value, it is 5x5 value, seperated by hiphens.
func (m *SideLoadingKey) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SideLoadingKey) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["totalActivation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActivation(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
func (m *SideLoadingKey) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDescription sets the description property value. Side Loading Key description displayed to the ITPro Admins..
func (m *SideLoadingKey) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Side Loading Key Name displayed to the ITPro Admins.
func (m *SideLoadingKey) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Side Loading Key Last Updated Date displayed to the ITPro Admins.
func (m *SideLoadingKey) SetLastUpdatedDateTime(value *string)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetTotalActivation sets the totalActivation property value. Side Loading Key Total Activation displayed to the ITPro Admins.
func (m *SideLoadingKey) SetTotalActivation(value *int32)() {
    if m != nil {
        m.totalActivation = value
    }
}
// SetValue sets the value property value. Side Loading Key Value, it is 5x5 value, seperated by hiphens.
func (m *SideLoadingKey) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
