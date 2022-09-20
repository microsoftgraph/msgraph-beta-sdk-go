package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LinkedResource_v2 provides operations to manage the collection of activityStatistics entities.
type LinkedResource_v2 struct {
    Entity
    // Field indicating the app name of the source that is sending the linkedResource.
    applicationName *string
    // Field indicating the title of the linkedResource.
    displayName *string
    // Id of the object that is associated with this task on the third-party/partner system.
    externalId *string
    // Deep link to the linkedResource.
    webUrl *string
}
// NewLinkedResource_v2 instantiates a new linkedResource_v2 and sets the default values.
func NewLinkedResource_v2()(*LinkedResource_v2) {
    m := &LinkedResource_v2{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.linkedResource_v2";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLinkedResource_v2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLinkedResource_v2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLinkedResource_v2(), nil
}
// GetApplicationName gets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
func (m *LinkedResource_v2) GetApplicationName()(*string) {
    return m.applicationName
}
// GetDisplayName gets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource_v2) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource_v2) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LinkedResource_v2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    return m.webUrl
}
// Serialize serializes information the current object
func (m *LinkedResource_v2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    m.applicationName = value
}
// SetDisplayName sets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource_v2) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource_v2) SetExternalId(value *string)() {
    m.externalId = value
}
// SetWebUrl sets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource_v2) SetWebUrl(value *string)() {
    m.webUrl = value
}
