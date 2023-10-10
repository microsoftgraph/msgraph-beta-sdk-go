package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AzureAssociatedIdentities 
type AzureAssociatedIdentities struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAzureAssociatedIdentities instantiates a new azureAssociatedIdentities and sets the default values.
func NewAzureAssociatedIdentities()(*AzureAssociatedIdentities) {
    m := &AzureAssociatedIdentities{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAzureAssociatedIdentitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureAssociatedIdentitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureAssociatedIdentities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureAssociatedIdentities) GetAdditionalData()(map[string]any) {
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
func (m *AzureAssociatedIdentities) GetAll()([]AzureIdentityable) {
    val, err := m.GetBackingStore().Get("all")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureIdentityable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AzureAssociatedIdentities) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureAssociatedIdentities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["all"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureIdentityable)
                }
            }
            m.SetAll(res)
        }
        return nil
    }
    res["managedIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureManagedIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureManagedIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureManagedIdentityable)
                }
            }
            m.SetManagedIdentities(res)
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
    res["servicePrincipals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureServicePrincipalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureServicePrincipalable)
                }
            }
            m.SetServicePrincipals(res)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AzureUserable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetManagedIdentities gets the managedIdentities property value. The managedIdentities property
func (m *AzureAssociatedIdentities) GetManagedIdentities()([]AzureManagedIdentityable) {
    val, err := m.GetBackingStore().Get("managedIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureManagedIdentityable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AzureAssociatedIdentities) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePrincipals gets the servicePrincipals property value. The servicePrincipals property
func (m *AzureAssociatedIdentities) GetServicePrincipals()([]AzureServicePrincipalable) {
    val, err := m.GetBackingStore().Get("servicePrincipals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureServicePrincipalable)
    }
    return nil
}
// GetUsers gets the users property value. The users property
func (m *AzureAssociatedIdentities) GetUsers()([]AzureUserable) {
    val, err := m.GetBackingStore().Get("users")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AzureUserable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AzureAssociatedIdentities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetManagedIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedIdentities()))
        for i, v := range m.GetManagedIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("managedIdentities", cast)
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
    if m.GetServicePrincipals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePrincipals()))
        for i, v := range m.GetServicePrincipals() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("servicePrincipals", cast)
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
func (m *AzureAssociatedIdentities) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAll sets the all property value. The all property
func (m *AzureAssociatedIdentities) SetAll(value []AzureIdentityable)() {
    err := m.GetBackingStore().Set("all", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AzureAssociatedIdentities) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetManagedIdentities sets the managedIdentities property value. The managedIdentities property
func (m *AzureAssociatedIdentities) SetManagedIdentities(value []AzureManagedIdentityable)() {
    err := m.GetBackingStore().Set("managedIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AzureAssociatedIdentities) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipals sets the servicePrincipals property value. The servicePrincipals property
func (m *AzureAssociatedIdentities) SetServicePrincipals(value []AzureServicePrincipalable)() {
    err := m.GetBackingStore().Set("servicePrincipals", value)
    if err != nil {
        panic(err)
    }
}
// SetUsers sets the users property value. The users property
func (m *AzureAssociatedIdentities) SetUsers(value []AzureUserable)() {
    err := m.GetBackingStore().Set("users", value)
    if err != nil {
        panic(err)
    }
}
// AzureAssociatedIdentitiesable 
type AzureAssociatedIdentitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAll()([]AzureIdentityable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetManagedIdentities()([]AzureManagedIdentityable)
    GetOdataType()(*string)
    GetServicePrincipals()([]AzureServicePrincipalable)
    GetUsers()([]AzureUserable)
    SetAll(value []AzureIdentityable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetManagedIdentities(value []AzureManagedIdentityable)()
    SetOdataType(value *string)()
    SetServicePrincipals(value []AzureServicePrincipalable)()
    SetUsers(value []AzureUserable)()
}
