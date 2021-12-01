package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AllowedDataLocation 
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
// NewAllowedDataLocation instantiates a new allowedDataLocation and sets the default values.
func NewAllowedDataLocation()(*AllowedDataLocation) {
    m := &AllowedDataLocation{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppId gets the appId property value. 
func (m *AllowedDataLocation) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetDomain gets the domain property value. 
func (m *AllowedDataLocation) GetDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domain
    }
}
// GetIsDefault gets the isDefault property value. 
func (m *AllowedDataLocation) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetLocation gets the location property value. 
func (m *AllowedDataLocation) GetLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AllowedDataLocation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["domain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    return res
}
func (m *AllowedDataLocation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAppId sets the appId property value. 
func (m *AllowedDataLocation) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetDomain sets the domain property value. 
func (m *AllowedDataLocation) SetDomain(value *string)() {
    if m != nil {
        m.domain = value
    }
}
// SetIsDefault sets the isDefault property value. 
func (m *AllowedDataLocation) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetLocation sets the location property value. 
func (m *AllowedDataLocation) SetLocation(value *string)() {
    if m != nil {
        m.location = value
    }
}
