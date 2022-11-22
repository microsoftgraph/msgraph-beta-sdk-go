package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Contract provides operations to manage the collection of accessReview entities.
type Contract struct {
    DirectoryObject
    // The contractType property
    contractType *string
    // The customerId property
    customerId *string
    // The defaultDomainName property
    defaultDomainName *string
    // The displayName property
    displayName *string
}
// NewContract instantiates a new contract and sets the default values.
func NewContract()(*Contract) {
    m := &Contract{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.contract";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateContractFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContractFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContract(), nil
}
// GetContractType gets the contractType property value. The contractType property
func (m *Contract) GetContractType()(*string) {
    return m.contractType
}
// GetCustomerId gets the customerId property value. The customerId property
func (m *Contract) GetCustomerId()(*string) {
    return m.customerId
}
// GetDefaultDomainName gets the defaultDomainName property value. The defaultDomainName property
func (m *Contract) GetDefaultDomainName()(*string) {
    return m.defaultDomainName
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Contract) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Contract) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["contractType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContractType)
    res["customerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomerId)
    res["defaultDomainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultDomainName)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    return res
}
// Serialize serializes information the current object
func (m *Contract) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contractType", m.GetContractType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
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
    return nil
}
// SetContractType sets the contractType property value. The contractType property
func (m *Contract) SetContractType(value *string)() {
    m.contractType = value
}
// SetCustomerId sets the customerId property value. The customerId property
func (m *Contract) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDefaultDomainName sets the defaultDomainName property value. The defaultDomainName property
func (m *Contract) SetDefaultDomainName(value *string)() {
    m.defaultDomainName = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Contract) SetDisplayName(value *string)() {
    m.displayName = value
}
