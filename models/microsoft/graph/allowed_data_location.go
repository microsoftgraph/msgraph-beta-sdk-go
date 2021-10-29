package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AllowedDataLocation struct {
    Entity
    // 
    appId *string;
    // 
    domain *string;
    // 
    isDefault *bool;
    // 
    location *string;
}
// Instantiates a new allowedDataLocation and sets the default values.
func NewAllowedDataLocation()(*AllowedDataLocation) {
    m := &AllowedDataLocation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appId property value. 
func (m *AllowedDataLocation) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the domain property value. 
func (m *AllowedDataLocation) GetDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domain
    }
}
// Gets the isDefault property value. 
func (m *AllowedDataLocation) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the location property value. 
func (m *AllowedDataLocation) GetLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// The deserialization information for the current model
func (m *AllowedDataLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["domain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDomain(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocation(val)
        return nil
    }
    return res
}
func (m *AllowedDataLocation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AllowedDataLocation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appId property value. 
// Parameters:
//  - value : Value to set for the appId property.
func (m *AllowedDataLocation) SetAppId(value *string)() {
    m.appId = value
}
// Sets the domain property value. 
// Parameters:
//  - value : Value to set for the domain property.
func (m *AllowedDataLocation) SetDomain(value *string)() {
    m.domain = value
}
// Sets the isDefault property value. 
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *AllowedDataLocation) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the location property value. 
// Parameters:
//  - value : Value to set for the location property.
func (m *AllowedDataLocation) SetLocation(value *string)() {
    m.location = value
}
