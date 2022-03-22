package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomCalloutExtension 
type CustomCalloutExtension struct {
    Entity
    // Configuration for securing the API call to the logic app. For example, using OAuth client credentials flow.
    authenticationConfiguration CustomExtensionAuthenticationConfigurationable;
    // HTTP connection settings that define how long Azure AD can wait for a connection to a logic app, how many times you can retry a timed-out connection and the exception scenarios when retries are allowed.
    clientConfiguration CustomExtensionClientConfigurationable;
    // Description for the customCalloutExtension object.
    description *string;
    // Display name for the customCalloutExtension object.
    displayName *string;
    // The type and details for configuring the endpoint to call the logic app's workflow.
    endpointConfiguration CustomExtensionEndpointConfigurationable;
}
// NewCustomCalloutExtension instantiates a new customCalloutExtension and sets the default values.
func NewCustomCalloutExtension()(*CustomCalloutExtension) {
    m := &CustomCalloutExtension{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomCalloutExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomCalloutExtensionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCustomCalloutExtension(), nil
}
// GetAuthenticationConfiguration gets the authenticationConfiguration property value. Configuration for securing the API call to the logic app. For example, using OAuth client credentials flow.
func (m *CustomCalloutExtension) GetAuthenticationConfiguration()(CustomExtensionAuthenticationConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.authenticationConfiguration
    }
}
// GetClientConfiguration gets the clientConfiguration property value. HTTP connection settings that define how long Azure AD can wait for a connection to a logic app, how many times you can retry a timed-out connection and the exception scenarios when retries are allowed.
func (m *CustomCalloutExtension) GetClientConfiguration()(CustomExtensionClientConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.clientConfiguration
    }
}
// GetDescription gets the description property value. Description for the customCalloutExtension object.
func (m *CustomCalloutExtension) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name for the customCalloutExtension object.
func (m *CustomCalloutExtension) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndpointConfiguration gets the endpointConfiguration property value. The type and details for configuring the endpoint to call the logic app's workflow.
func (m *CustomCalloutExtension) GetEndpointConfiguration()(CustomExtensionEndpointConfigurationable) {
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
        val, err := n.GetObjectValue(CreateCustomExtensionAuthenticationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationConfiguration(val.(CustomExtensionAuthenticationConfigurationable))
        }
        return nil
    }
    res["clientConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionClientConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientConfiguration(val.(CustomExtensionClientConfigurationable))
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
        val, err := n.GetObjectValue(CreateCustomExtensionEndpointConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointConfiguration(val.(CustomExtensionEndpointConfigurationable))
        }
        return nil
    }
    return res
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
// SetAuthenticationConfiguration sets the authenticationConfiguration property value. Configuration for securing the API call to the logic app. For example, using OAuth client credentials flow.
func (m *CustomCalloutExtension) SetAuthenticationConfiguration(value CustomExtensionAuthenticationConfigurationable)() {
    if m != nil {
        m.authenticationConfiguration = value
    }
}
// SetClientConfiguration sets the clientConfiguration property value. HTTP connection settings that define how long Azure AD can wait for a connection to a logic app, how many times you can retry a timed-out connection and the exception scenarios when retries are allowed.
func (m *CustomCalloutExtension) SetClientConfiguration(value CustomExtensionClientConfigurationable)() {
    if m != nil {
        m.clientConfiguration = value
    }
}
// SetDescription sets the description property value. Description for the customCalloutExtension object.
func (m *CustomCalloutExtension) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name for the customCalloutExtension object.
func (m *CustomCalloutExtension) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndpointConfiguration sets the endpointConfiguration property value. The type and details for configuring the endpoint to call the logic app's workflow.
func (m *CustomCalloutExtension) SetEndpointConfiguration(value CustomExtensionEndpointConfigurationable)() {
    if m != nil {
        m.endpointConfiguration = value
    }
}
