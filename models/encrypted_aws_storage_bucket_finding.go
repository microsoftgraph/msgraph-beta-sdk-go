package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptedAwsStorageBucketFinding 
type EncryptedAwsStorageBucketFinding struct {
    Finding
}
// NewEncryptedAwsStorageBucketFinding instantiates a new encryptedAwsStorageBucketFinding and sets the default values.
func NewEncryptedAwsStorageBucketFinding()(*EncryptedAwsStorageBucketFinding) {
    m := &EncryptedAwsStorageBucketFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateEncryptedAwsStorageBucketFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEncryptedAwsStorageBucketFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptedAwsStorageBucketFinding(), nil
}
// GetAccessibility gets the accessibility property value. The accessibility property
func (m *EncryptedAwsStorageBucketFinding) GetAccessibility()(*AwsAccessType) {
    val, err := m.GetBackingStore().Get("accessibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsAccessType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptedAwsStorageBucketFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *EncryptedAwsStorageBucketFinding) GetStorageBucket()(AuthorizationSystemResourceable) {
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
func (m *EncryptedAwsStorageBucketFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("storageBucket", m.GetStorageBucket())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibility sets the accessibility property value. The accessibility property
func (m *EncryptedAwsStorageBucketFinding) SetAccessibility(value *AwsAccessType)() {
    err := m.GetBackingStore().Set("accessibility", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBucket sets the storageBucket property value. The storageBucket property
func (m *EncryptedAwsStorageBucketFinding) SetStorageBucket(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("storageBucket", value)
    if err != nil {
        panic(err)
    }
}
// EncryptedAwsStorageBucketFindingable 
type EncryptedAwsStorageBucketFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibility()(*AwsAccessType)
    GetStorageBucket()(AuthorizationSystemResourceable)
    SetAccessibility(value *AwsAccessType)()
    SetStorageBucket(value AuthorizationSystemResourceable)()
}
