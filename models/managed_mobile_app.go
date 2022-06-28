package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedMobileApp the identifier for the deployment an app.
type ManagedMobileApp struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier for an app with it's operating system type.
    mobileAppIdentifier MobileAppIdentifierable
    // Version of the entity.
    version *string
}
// NewManagedMobileApp instantiates a new managedMobileApp and sets the default values.
func NewManagedMobileApp()(*ManagedMobileApp) {
    m := &ManagedMobileApp{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedMobileAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedMobileAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedMobileApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedMobileApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedMobileApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["mobileAppIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIdentifier(val.(MobileAppIdentifierable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetMobileAppIdentifier gets the mobileAppIdentifier property value. The identifier for an app with it's operating system type.
func (m *ManagedMobileApp) GetMobileAppIdentifier()(MobileAppIdentifierable) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIdentifier
    }
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedMobileApp) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *ManagedMobileApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("mobileAppIdentifier", m.GetMobileAppIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
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
func (m *ManagedMobileApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMobileAppIdentifier sets the mobileAppIdentifier property value. The identifier for an app with it's operating system type.
func (m *ManagedMobileApp) SetMobileAppIdentifier(value MobileAppIdentifierable)() {
    if m != nil {
        m.mobileAppIdentifier = value
    }
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedMobileApp) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
