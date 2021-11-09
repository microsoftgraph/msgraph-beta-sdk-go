package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UsageRight struct {
    Entity
    // Product id corresponding to the usage right.
    catalogId *string;
    // Identifier of the service corresponding to the usage right.
    serviceIdentifier *string;
    // The state of the usage right. Possible values are: active, inactive, warning, suspended.
    state *UsageRightState;
}
// Instantiates a new usageRight and sets the default values.
func NewUsageRight()(*UsageRight) {
    m := &UsageRight{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the catalogId property value. Product id corresponding to the usage right.
func (m *UsageRight) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
// Gets the serviceIdentifier property value. Identifier of the service corresponding to the usage right.
func (m *UsageRight) GetServiceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceIdentifier
    }
}
// Gets the state property value. The state of the usage right. Possible values are: active, inactive, warning, suspended.
func (m *UsageRight) GetState()(*UsageRightState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
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
            cast := val.(UsageRightState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *UsageRight) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the catalogId property value. Product id corresponding to the usage right.
// Parameters:
//  - value : Value to set for the catalogId property.
func (m *UsageRight) SetCatalogId(value *string)() {
    m.catalogId = value
}
// Sets the serviceIdentifier property value. Identifier of the service corresponding to the usage right.
// Parameters:
//  - value : Value to set for the serviceIdentifier property.
func (m *UsageRight) SetServiceIdentifier(value *string)() {
    m.serviceIdentifier = value
}
// Sets the state property value. The state of the usage right. Possible values are: active, inactive, warning, suspended.
// Parameters:
//  - value : Value to set for the state property.
func (m *UsageRight) SetState(value *UsageRightState)() {
    m.state = value
}
