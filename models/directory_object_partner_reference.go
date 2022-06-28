package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryObjectPartnerReference 
type DirectoryObjectPartnerReference struct {
    DirectoryObject
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Description of the object returned. Read-only.
    description *string
    // Name of directory object being returned, like group or application. Read-only.
    displayName *string
    // The tenant identifier for the partner tenant. Read-only.
    externalPartnerTenantId *string
    // The type of the referenced object in the partner tenant. Read-only.
    objectType *string
}
// NewDirectoryObjectPartnerReference instantiates a new DirectoryObjectPartnerReference and sets the default values.
func NewDirectoryObjectPartnerReference()(*DirectoryObjectPartnerReference) {
    m := &DirectoryObjectPartnerReference{
        DirectoryObject: *NewDirectoryObject(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDirectoryObjectPartnerReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryObjectPartnerReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryObjectPartnerReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectoryObjectPartnerReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Description of the object returned. Read-only.
func (m *DirectoryObjectPartnerReference) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of directory object being returned, like group or application. Read-only.
func (m *DirectoryObjectPartnerReference) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalPartnerTenantId gets the externalPartnerTenantId property value. The tenant identifier for the partner tenant. Read-only.
func (m *DirectoryObjectPartnerReference) GetExternalPartnerTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalPartnerTenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryObjectPartnerReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["externalPartnerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalPartnerTenantId(val)
        }
        return nil
    }
    res["objectType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectType(val)
        }
        return nil
    }
    return res
}
// GetObjectType gets the objectType property value. The type of the referenced object in the partner tenant. Read-only.
func (m *DirectoryObjectPartnerReference) GetObjectType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectType
    }
}
// Serialize serializes information the current object
func (m *DirectoryObjectPartnerReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
        err = writer.WriteStringValue("externalPartnerTenantId", m.GetExternalPartnerTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectType", m.GetObjectType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectoryObjectPartnerReference) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Description of the object returned. Read-only.
func (m *DirectoryObjectPartnerReference) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of directory object being returned, like group or application. Read-only.
func (m *DirectoryObjectPartnerReference) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalPartnerTenantId sets the externalPartnerTenantId property value. The tenant identifier for the partner tenant. Read-only.
func (m *DirectoryObjectPartnerReference) SetExternalPartnerTenantId(value *string)() {
    if m != nil {
        m.externalPartnerTenantId = value
    }
}
// SetObjectType sets the objectType property value. The type of the referenced object in the partner tenant. Read-only.
func (m *DirectoryObjectPartnerReference) SetObjectType(value *string)() {
    if m != nil {
        m.objectType = value
    }
}
