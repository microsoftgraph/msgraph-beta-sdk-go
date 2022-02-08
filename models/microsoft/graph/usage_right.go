package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UsageRight 
type UsageRight struct {
    Entity
    // Product id corresponding to the usage right.
    catalogId *string;
    // Identifier of the service corresponding to the usage right.
    serviceIdentifier *string;
    // The state of the usage right. Possible values are: active, inactive, warning, suspended.
    state *UsageRightState;
}
// NewUsageRight instantiates a new usageRight and sets the default values.
func NewUsageRight()(*UsageRight) {
    m := &UsageRight{
        Entity: *NewEntity(),
    }
    return m
}
// GetCatalogId gets the catalogId property value. Product id corresponding to the usage right.
func (m *UsageRight) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// GetServiceIdentifier gets the serviceIdentifier property value. Identifier of the service corresponding to the usage right.
func (m *UsageRight) GetServiceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceIdentifier
    }
}
// GetState gets the state property value. The state of the usage right. Possible values are: active, inactive, warning, suspended.
func (m *UsageRight) GetState()(*UsageRightState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsageRight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["serviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceIdentifier(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageRightState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*UsageRightState))
        }
        return nil
    }
    return res
}
func (m *UsageRight) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UsageRight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceIdentifier", m.GetServiceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalogId sets the catalogId property value. Product id corresponding to the usage right.
func (m *UsageRight) SetCatalogId(value *string)() {
    if m != nil {
        m.catalogId = value
    }
}
// SetServiceIdentifier sets the serviceIdentifier property value. Identifier of the service corresponding to the usage right.
func (m *UsageRight) SetServiceIdentifier(value *string)() {
    if m != nil {
        m.serviceIdentifier = value
    }
}
// SetState sets the state property value. The state of the usage right. Possible values are: active, inactive, warning, suspended.
func (m *UsageRight) SetState(value *UsageRightState)() {
    if m != nil {
        m.state = value
    }
}
