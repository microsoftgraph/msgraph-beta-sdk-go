package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreAuthorizedApplication 
type PreAuthorizedApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier for the application.
    appId *string
    // The unique identifier for the oauth2PermissionScopes the application requires.
    permissionIds []string
}
// NewPreAuthorizedApplication instantiates a new preAuthorizedApplication and sets the default values.
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePreAuthorizedApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePreAuthorizedApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreAuthorizedApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppId gets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["permissionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPermissionIds(res)
        }
        return nil
    }
    return res
}
// GetPermissionIds gets the permissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) GetPermissionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissionIds
    }
}
// Serialize serializes information the current object
func (m *PreAuthorizedApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionIds() != nil {
        err := writer.WriteCollectionOfStringValues("permissionIds", m.GetPermissionIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppId sets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetPermissionIds sets the permissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) SetPermissionIds(value []string)() {
    if m != nil {
        m.permissionIds = value
    }
}
