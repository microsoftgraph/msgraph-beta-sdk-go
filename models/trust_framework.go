package models

import (
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
    odataTypeValue := "#microsoft.graph.trustFramework";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTrustFrameworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFramework(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFramework) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFramework) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keySets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkKeySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKeySetable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkKeySetable)
            }
            m.SetKeySets(res)
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
    res["policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkPolicyable)
            }
            m.SetPolicies(res)
        }
        return nil
    }
    return res
}
// GetKeySets gets the keySets property value. The keySets property
func (m *TrustFramework) GetKeySets()([]TrustFrameworkKeySetable) {
    if m == nil {
        return nil
    } else {
        return m.keySets
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TrustFramework) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetPolicies gets the policies property value. The policies property
func (m *TrustFramework) GetPolicies()([]TrustFrameworkPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.policies
    }
}
// Serialize serializes information the current object
func (m *TrustFramework) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeySets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeySets()))
        for i, v := range m.GetKeySets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetKeySets sets the keySets property value. The keySets property
func (m *TrustFramework) SetKeySets(value []TrustFrameworkKeySetable)() {
    if m != nil {
        m.keySets = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrustFramework) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetPolicies sets the policies property value. The policies property
func (m *TrustFramework) SetPolicies(value []TrustFrameworkPolicyable)() {
    if m != nil {
        m.policies = value
    }
}
