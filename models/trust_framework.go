package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFramework 
type TrustFramework struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The keySets property
    keySets []TrustFrameworkKeySetable
    // The OdataType property
    odataType *string
    // The policies property
    policies []TrustFrameworkPolicyable
}
// NewTrustFramework instantiates a new TrustFramework and sets the default values.
func NewTrustFramework()(*TrustFramework) {
    m := &TrustFramework{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTrustFrameworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFramework(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFramework) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFramework) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keySets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTrustFrameworkKeySetFromDiscriminatorValue , m.SetKeySets)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["policies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTrustFrameworkPolicyFromDiscriminatorValue , m.SetPolicies)
    return res
}
// GetKeySets gets the keySets property value. The keySets property
func (m *TrustFramework) GetKeySets()([]TrustFrameworkKeySetable) {
    return m.keySets
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TrustFramework) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicies gets the policies property value. The policies property
func (m *TrustFramework) GetPolicies()([]TrustFrameworkPolicyable) {
    return m.policies
}
// Serialize serializes information the current object
func (m *TrustFramework) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeySets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetKeySets())
        err := writer.WriteCollectionOfObjectValues("keySets", cast)
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
    if m.GetPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPolicies())
        err := writer.WriteCollectionOfObjectValues("policies", cast)
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
func (m *TrustFramework) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeySets sets the keySets property value. The keySets property
func (m *TrustFramework) SetKeySets(value []TrustFrameworkKeySetable)() {
    m.keySets = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrustFramework) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicies sets the policies property value. The policies property
func (m *TrustFramework) SetPolicies(value []TrustFrameworkPolicyable)() {
    m.policies = value
}
