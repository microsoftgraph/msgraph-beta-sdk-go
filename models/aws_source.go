package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsSource struct {
    AuthorizationSystemIdentitySource
}
// NewAwsSource instantiates a new AwsSource and sets the default values.
func NewAwsSource()(*AwsSource) {
    m := &AwsSource{
        AuthorizationSystemIdentitySource: *NewAuthorizationSystemIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.awsSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsSource(), nil
}
// GetAccountId gets the accountId property value. AWS account ID.
// returns a *string when successful
func (m *AwsSource) GetAccountId()(*string) {
    val, err := m.GetBackingStore().Get("accountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentitySource.GetFieldDeserializers()
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AwsSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountId sets the accountId property value. AWS account ID.
func (m *AwsSource) SetAccountId(value *string)() {
    err := m.GetBackingStore().Set("accountId", value)
    if err != nil {
        panic(err)
    }
}
type AwsSourceable interface {
    AuthorizationSystemIdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountId()(*string)
    SetAccountId(value *string)()
}
