package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OemWarrantyInformationOnboarding 
type OemWarrantyInformationOnboarding struct {
    Entity
    // Specifies whether warranty API is available. This property is read-only.
    available *bool;
    // Specifies whether warranty query is enabled for given OEM. This property is read-only.
    enabled *bool;
    // OEM name. This property is read-only.
    oemName *string;
}
// NewOemWarrantyInformationOnboarding instantiates a new oemWarrantyInformationOnboarding and sets the default values.
func NewOemWarrantyInformationOnboarding()(*OemWarrantyInformationOnboarding) {
    m := &OemWarrantyInformationOnboarding{
        Entity: *NewEntity(),
    }
    return m
}
// GetAvailable gets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.available
    }
}
// GetEnabled gets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetOemName gets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetOemName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oemName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarrantyInformationOnboarding) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["available"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailable(val)
        }
        return nil
    }
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["oemName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOemName(val)
        }
        return nil
    }
    return res
}
func (m *OemWarrantyInformationOnboarding) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OemWarrantyInformationOnboarding) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("available", m.GetAvailable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("oemName", m.GetOemName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailable sets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetAvailable(value *bool)() {
    if m != nil {
        m.available = value
    }
}
// SetEnabled sets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetOemName sets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetOemName(value *string)() {
    if m != nil {
        m.oemName = value
    }
}
