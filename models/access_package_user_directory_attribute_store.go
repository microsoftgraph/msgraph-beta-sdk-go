package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageUserDirectoryAttributeStore 
type AccessPackageUserDirectoryAttributeStore struct {
    AccessPackageResourceAttributeDestination
}
// NewAccessPackageUserDirectoryAttributeStore instantiates a new accessPackageUserDirectoryAttributeStore and sets the default values.
func NewAccessPackageUserDirectoryAttributeStore()(*AccessPackageUserDirectoryAttributeStore) {
    m := &AccessPackageUserDirectoryAttributeStore{
        AccessPackageResourceAttributeDestination: *NewAccessPackageResourceAttributeDestination(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageUserDirectoryAttributeStore"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessPackageUserDirectoryAttributeStoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageUserDirectoryAttributeStoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageUserDirectoryAttributeStore(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageUserDirectoryAttributeStore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageResourceAttributeDestination.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageUserDirectoryAttributeStore) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageUserDirectoryAttributeStore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageResourceAttributeDestination.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageUserDirectoryAttributeStore) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageUserDirectoryAttributeStoreable 
type AccessPackageUserDirectoryAttributeStoreable interface {
    AccessPackageResourceAttributeDestinationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
