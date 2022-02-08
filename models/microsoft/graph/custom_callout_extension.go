package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomCalloutExtension 
type CustomCalloutExtension struct {
    Entity
    // 
    authenticationConfiguration *CustomExtensionAuthenticationConfiguration;
    // 
    clientConfiguration *CustomExtensionClientConfiguration;
    // 
    description *string;
    // 
    displayName *string;
    // 
    endpointConfiguration *CustomExtensionEndpointConfiguration;
}
// NewCustomCalloutExtension instantiates a new customCalloutExtension and sets the default values.
func NewCustomCalloutExtension()(*CustomCalloutExtension) {
    m := &CustomCalloutExtension{
        Entity: *NewEntity(),
    }
    return m
}
// GetAuthenticationConfiguration gets the authenticationConfiguration property value. 
func (m *CustomCalloutExtension) GetAuthenticationConfiguration()(*CustomExtensionAuthenticationConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.authenticationConfiguration
    }
}
// GetClientConfiguration gets the clientConfiguration property value. 
func (m *CustomCalloutExtension) GetClientConfiguration()(*CustomExtensionClientConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.clientConfiguration
    }
}
// GetDescription gets the description property value. 
func (m *CustomCalloutExtension) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *CustomCalloutExtension) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndpointConfiguration gets the endpointConfiguration property value. 
func (m *CustomCalloutExtension) GetEndpointConfiguration()(*CustomExtensionEndpointConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.endpointConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomCalloutExtension) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomExtensionAuthenticationConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationConfiguration(val.(*CustomExtensionAuthenticationConfiguration))
        }
        return nil
    }
    res["clientConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomExtensionClientConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientConfiguration(val.(*CustomExtensionClientConfiguration))
        }
        return nil
    }
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
    res["endpointConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomExtensionEndpointConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointConfiguration(val.(*CustomExtensionEndpointConfiguration))
        }
        return nil
    }
    return res
}
func (m *CustomCalloutExtension) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomCalloutExtension) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authenticationConfiguration", m.GetAuthenticationConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clientConfiguration", m.GetClientConfiguration())
        if err != nil {
            return err
        }
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
        err = writer.WriteObjectValue("endpointConfiguration", m.GetEndpointConfiguration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationConfiguration sets the authenticationConfiguration property value. 
func (m *CustomCalloutExtension) SetAuthenticationConfiguration(value *CustomExtensionAuthenticationConfiguration)() {
    if m != nil {
        m.authenticationConfiguration = value
    }
}
// SetClientConfiguration sets the clientConfiguration property value. 
func (m *CustomCalloutExtension) SetClientConfiguration(value *CustomExtensionClientConfiguration)() {
    if m != nil {
        m.clientConfiguration = value
    }
}
// SetDescription sets the description property value. 
func (m *CustomCalloutExtension) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *CustomCalloutExtension) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndpointConfiguration sets the endpointConfiguration property value. 
func (m *CustomCalloutExtension) SetEndpointConfiguration(value *CustomExtensionEndpointConfiguration)() {
    if m != nil {
        m.endpointConfiguration = value
    }
}
