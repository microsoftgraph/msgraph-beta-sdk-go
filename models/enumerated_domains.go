package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnumeratedDomains struct {
    ValidatingDomains
}
// NewEnumeratedDomains instantiates a new EnumeratedDomains and sets the default values.
func NewEnumeratedDomains()(*EnumeratedDomains) {
    m := &EnumeratedDomains{
        ValidatingDomains: *NewValidatingDomains(),
    }
    odataTypeValue := "#microsoft.graph.enumeratedDomains"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEnumeratedDomainsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnumeratedDomainsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnumeratedDomains(), nil
}
// GetDomainNames gets the domainNames property value. List of federated or managed root domains that Microsoft Entra ID validates.
// returns a []string when successful
func (m *EnumeratedDomains) GetDomainNames()([]string) {
    val, err := m.GetBackingStore().Get("domainNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnumeratedDomains) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ValidatingDomains.GetFieldDeserializers()
    res["domainNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomainNames(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EnumeratedDomains) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ValidatingDomains.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDomainNames() != nil {
        err = writer.WriteCollectionOfStringValues("domainNames", m.GetDomainNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDomainNames sets the domainNames property value. List of federated or managed root domains that Microsoft Entra ID validates.
func (m *EnumeratedDomains) SetDomainNames(value []string)() {
    err := m.GetBackingStore().Set("domainNames", value)
    if err != nil {
        panic(err)
    }
}
type EnumeratedDomainsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ValidatingDomainsable
    GetDomainNames()([]string)
    SetDomainNames(value []string)()
}
