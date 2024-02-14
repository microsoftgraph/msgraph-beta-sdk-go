package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternallyAccessibleAwsStorageBucketFinding struct {
    Finding
}
// NewExternallyAccessibleAwsStorageBucketFinding instantiates a new ExternallyAccessibleAwsStorageBucketFinding and sets the default values.
func NewExternallyAccessibleAwsStorageBucketFinding()(*ExternallyAccessibleAwsStorageBucketFinding) {
    m := &ExternallyAccessibleAwsStorageBucketFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateExternallyAccessibleAwsStorageBucketFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternallyAccessibleAwsStorageBucketFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternallyAccessibleAwsStorageBucketFinding(), nil
}
// GetAccessibility gets the accessibility property value. The accessibility property
// returns a *AwsAccessType when successful
func (m *ExternallyAccessibleAwsStorageBucketFinding) GetAccessibility()(*AwsAccessType) {
    val, err := m.GetBackingStore().Get("accessibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsAccessType)
    }
    return nil
}
// GetAccountsWithAccess gets the accountsWithAccess property value. The accountsWithAccess property
// returns a AccountsWithAccessable when successful
func (m *ExternallyAccessibleAwsStorageBucketFinding) GetAccountsWithAccess()(AccountsWithAccessable) {
    val, err := m.GetBackingStore().Get("accountsWithAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccountsWithAccessable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternallyAccessibleAwsStorageBucketFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibility(val.(*AwsAccessType))
        }
        return nil
    }
    res["accountsWithAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountsWithAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountsWithAccess(val.(AccountsWithAccessable))
        }
        return nil
    }
    res["storageBucket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBucket(val.(AuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetStorageBucket gets the storageBucket property value. The storageBucket property
// returns a AuthorizationSystemResourceable when successful
func (m *ExternallyAccessibleAwsStorageBucketFinding) GetStorageBucket()(AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("storageBucket")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExternallyAccessibleAwsStorageBucketFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessibility() != nil {
        cast := (*m.GetAccessibility()).String()
        err = writer.WriteStringValue("accessibility", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accountsWithAccess", m.GetAccountsWithAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storageBucket", m.GetStorageBucket())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibility sets the accessibility property value. The accessibility property
func (m *ExternallyAccessibleAwsStorageBucketFinding) SetAccessibility(value *AwsAccessType)() {
    err := m.GetBackingStore().Set("accessibility", value)
    if err != nil {
        panic(err)
    }
}
// SetAccountsWithAccess sets the accountsWithAccess property value. The accountsWithAccess property
func (m *ExternallyAccessibleAwsStorageBucketFinding) SetAccountsWithAccess(value AccountsWithAccessable)() {
    err := m.GetBackingStore().Set("accountsWithAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBucket sets the storageBucket property value. The storageBucket property
func (m *ExternallyAccessibleAwsStorageBucketFinding) SetStorageBucket(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("storageBucket", value)
    if err != nil {
        panic(err)
    }
}
type ExternallyAccessibleAwsStorageBucketFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibility()(*AwsAccessType)
    GetAccountsWithAccess()(AccountsWithAccessable)
    GetStorageBucket()(AuthorizationSystemResourceable)
    SetAccessibility(value *AwsAccessType)()
    SetAccountsWithAccess(value AccountsWithAccessable)()
    SetStorageBucket(value AuthorizationSystemResourceable)()
}
