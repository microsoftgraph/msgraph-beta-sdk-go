package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnumeratedAccountsWithAccess struct {
    AccountsWithAccess
}
// NewEnumeratedAccountsWithAccess instantiates a new EnumeratedAccountsWithAccess and sets the default values.
func NewEnumeratedAccountsWithAccess()(*EnumeratedAccountsWithAccess) {
    m := &EnumeratedAccountsWithAccess{
        AccountsWithAccess: *NewAccountsWithAccess(),
    }
    odataTypeValue := "#microsoft.graph.enumeratedAccountsWithAccess"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEnumeratedAccountsWithAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnumeratedAccountsWithAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnumeratedAccountsWithAccess(), nil
}
// GetAccounts gets the accounts property value. The accounts property
// returns a []AuthorizationSystemable when successful
func (m *EnumeratedAccountsWithAccess) GetAccounts()([]AuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("accounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorizationSystemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnumeratedAccountsWithAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccountsWithAccess.GetFieldDeserializers()
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationSystemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorizationSystemable)
                }
            }
            m.SetAccounts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EnumeratedAccountsWithAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccountsWithAccess.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *EnumeratedAccountsWithAccess) SetAccounts(value []AuthorizationSystemable)() {
    err := m.GetBackingStore().Set("accounts", value)
    if err != nil {
        panic(err)
    }
}
type EnumeratedAccountsWithAccessable interface {
    AccountsWithAccessable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()([]AuthorizationSystemable)
    SetAccounts(value []AuthorizationSystemable)()
}
