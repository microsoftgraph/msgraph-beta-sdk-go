package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllAccountsWithAccess struct {
    AccountsWithAccess
}
// NewAllAccountsWithAccess instantiates a new AllAccountsWithAccess and sets the default values.
func NewAllAccountsWithAccess()(*AllAccountsWithAccess) {
    m := &AllAccountsWithAccess{
        AccountsWithAccess: *NewAccountsWithAccess(),
    }
    odataTypeValue := "#microsoft.graph.allAccountsWithAccess"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAllAccountsWithAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllAccountsWithAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllAccountsWithAccess(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllAccountsWithAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccountsWithAccess.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AllAccountsWithAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccountsWithAccess.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AllAccountsWithAccessable interface {
    AccountsWithAccessable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
