package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LinkedResource_v2 provides operations to manage the compliance singleton.
type LinkedResource_v2 struct {
    Entity
    // Field indicating the app name of the source that is sending the linkedResource.
    applicationName *string;
    // Field indicating the title of the linkedResource.
    displayName *string;
    // Id of the object that is associated with this task on the third-party/partner system.
    externalId *string;
    // Deep link to the linkedResource.
    webUrl *string;
}
// NewLinkedResource_v2 instantiates a new linkedResource_v2 and sets the default values.
func NewLinkedResource_v2()(*LinkedResource_v2) {
    m := &LinkedResource_v2{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLinkedResource_v2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLinkedResource_v2FromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLinkedResource_v2(), nil
}
// GetApplicationName gets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
func (m *LinkedResource_v2) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
// GetDisplayName gets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource_v2) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalId gets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource_v2) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LinkedResource_v2) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource_v2) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *LinkedResource_v2) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LinkedResource_v2) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationName", m.GetApplicationName())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationName sets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
func (m *LinkedResource_v2) SetApplicationName(value *string)() {
    if m != nil {
        m.applicationName = value
    }
}
// SetDisplayName sets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource_v2) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalId sets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource_v2) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetWebUrl sets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource_v2) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
