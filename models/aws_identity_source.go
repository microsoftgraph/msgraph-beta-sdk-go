package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsIdentitySource struct {
    PermissionsDefinitionIdentitySource
}
// NewAwsIdentitySource instantiates a new AwsIdentitySource and sets the default values.
func NewAwsIdentitySource()(*AwsIdentitySource) {
    m := &AwsIdentitySource{
        PermissionsDefinitionIdentitySource: *NewPermissionsDefinitionIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.awsIdentitySource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsIdentitySourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsIdentitySourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsIdentitySource(), nil
}
// GetAuthorizationSystemInfo gets the authorizationSystemInfo property value. The authorizationSystemInfo property
// returns a PermissionsDefinitionAuthorizationSystemable when successful
func (m *AwsIdentitySource) GetAuthorizationSystemInfo()(PermissionsDefinitionAuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("authorizationSystemInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsDefinitionAuthorizationSystemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsIdentitySource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinitionIdentitySource.GetFieldDeserializers()
    res["authorizationSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsDefinitionAuthorizationSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystemInfo(val.(PermissionsDefinitionAuthorizationSystemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AwsIdentitySource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinitionIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authorizationSystemInfo", m.GetAuthorizationSystemInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizationSystemInfo sets the authorizationSystemInfo property value. The authorizationSystemInfo property
func (m *AwsIdentitySource) SetAuthorizationSystemInfo(value PermissionsDefinitionAuthorizationSystemable)() {
    err := m.GetBackingStore().Set("authorizationSystemInfo", value)
    if err != nil {
        panic(err)
    }
}
type AwsIdentitySourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionIdentitySourceable
    GetAuthorizationSystemInfo()(PermissionsDefinitionAuthorizationSystemable)
    SetAuthorizationSystemInfo(value PermissionsDefinitionAuthorizationSystemable)()
}
