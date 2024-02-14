package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type GcpAssociatedIdentities struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewGcpAssociatedIdentities instantiates a new GcpAssociatedIdentities and sets the default values.
func NewGcpAssociatedIdentities()(*GcpAssociatedIdentities) {
    m := &GcpAssociatedIdentities{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGcpAssociatedIdentitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGcpAssociatedIdentitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpAssociatedIdentities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GcpAssociatedIdentities) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAll gets the all property value. The all property
// returns a []GcpIdentityable when successful
func (m *GcpAssociatedIdentities) GetAll()([]GcpIdentityable) {
    val, err := m.GetBackingStore().Get("all")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpIdentityable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *GcpAssociatedIdentities) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GcpAssociatedIdentities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["all"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpIdentityable)
                }
            }
            m.SetAll(res)
        }
        return nil
    }
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
    res["serviceAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpServiceAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpServiceAccountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpServiceAccountable)
                }
            }
            m.SetServiceAccounts(res)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGcpUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GcpUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GcpUserable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *GcpAssociatedIdentities) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServiceAccounts gets the serviceAccounts property value. The serviceAccounts property
// returns a []GcpServiceAccountable when successful
func (m *GcpAssociatedIdentities) GetServiceAccounts()([]GcpServiceAccountable) {
    val, err := m.GetBackingStore().Get("serviceAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpServiceAccountable)
    }
    return nil
}
// GetUsers gets the users property value. The users property
// returns a []GcpUserable when successful
func (m *GcpAssociatedIdentities) GetUsers()([]GcpUserable) {
    val, err := m.GetBackingStore().Get("users")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GcpUserable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GcpAssociatedIdentities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAll() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAll()))
        for i, v := range m.GetAll() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("all", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetServiceAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceAccounts()))
        for i, v := range m.GetServiceAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("serviceAccounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("users", cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GcpAssociatedIdentities) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAll sets the all property value. The all property
func (m *GcpAssociatedIdentities) SetAll(value []GcpIdentityable)() {
    err := m.GetBackingStore().Set("all", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *GcpAssociatedIdentities) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GcpAssociatedIdentities) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceAccounts sets the serviceAccounts property value. The serviceAccounts property
func (m *GcpAssociatedIdentities) SetServiceAccounts(value []GcpServiceAccountable)() {
    err := m.GetBackingStore().Set("serviceAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetUsers sets the users property value. The users property
func (m *GcpAssociatedIdentities) SetUsers(value []GcpUserable)() {
    err := m.GetBackingStore().Set("users", value)
    if err != nil {
        panic(err)
    }
}
type GcpAssociatedIdentitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAll()([]GcpIdentityable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetServiceAccounts()([]GcpServiceAccountable)
    GetUsers()([]GcpUserable)
    SetAll(value []GcpIdentityable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetServiceAccounts(value []GcpServiceAccountable)()
    SetUsers(value []GcpUserable)()
}
