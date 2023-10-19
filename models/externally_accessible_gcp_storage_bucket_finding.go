package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternallyAccessibleGcpStorageBucketFinding 
type ExternallyAccessibleGcpStorageBucketFinding struct {
    Finding
}
// NewExternallyAccessibleGcpStorageBucketFinding instantiates a new externallyAccessibleGcpStorageBucketFinding and sets the default values.
func NewExternallyAccessibleGcpStorageBucketFinding()(*ExternallyAccessibleGcpStorageBucketFinding) {
    m := &ExternallyAccessibleGcpStorageBucketFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateExternallyAccessibleGcpStorageBucketFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternallyAccessibleGcpStorageBucketFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternallyAccessibleGcpStorageBucketFinding(), nil
}
// GetAccessibility gets the accessibility property value. The accessibility property
func (m *ExternallyAccessibleGcpStorageBucketFinding) GetAccessibility()(*GcpAccessType) {
    val, err := m.GetBackingStore().Get("accessibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GcpAccessType)
    }
    return nil
}
// GetEncryptionManagedBy gets the encryptionManagedBy property value. The encryptionManagedBy property
func (m *ExternallyAccessibleGcpStorageBucketFinding) GetEncryptionManagedBy()(*GcpEncryption) {
    val, err := m.GetBackingStore().Get("encryptionManagedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GcpEncryption)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternallyAccessibleGcpStorageBucketFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGcpAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibility(val.(*GcpAccessType))
        }
        return nil
    }
    res["encryptionManagedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGcpEncryption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionManagedBy(val.(*GcpEncryption))
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
func (m *ExternallyAccessibleGcpStorageBucketFinding) GetStorageBucket()(AuthorizationSystemResourceable) {
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
func (m *ExternallyAccessibleGcpStorageBucketFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetEncryptionManagedBy() != nil {
        cast := (*m.GetEncryptionManagedBy()).String()
        err = writer.WriteStringValue("encryptionManagedBy", &cast)
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
func (m *ExternallyAccessibleGcpStorageBucketFinding) SetAccessibility(value *GcpAccessType)() {
    err := m.GetBackingStore().Set("accessibility", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionManagedBy sets the encryptionManagedBy property value. The encryptionManagedBy property
func (m *ExternallyAccessibleGcpStorageBucketFinding) SetEncryptionManagedBy(value *GcpEncryption)() {
    err := m.GetBackingStore().Set("encryptionManagedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBucket sets the storageBucket property value. The storageBucket property
func (m *ExternallyAccessibleGcpStorageBucketFinding) SetStorageBucket(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("storageBucket", value)
    if err != nil {
        panic(err)
    }
}
// ExternallyAccessibleGcpStorageBucketFindingable 
type ExternallyAccessibleGcpStorageBucketFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibility()(*GcpAccessType)
    GetEncryptionManagedBy()(*GcpEncryption)
    GetStorageBucket()(AuthorizationSystemResourceable)
    SetAccessibility(value *GcpAccessType)()
    SetEncryptionManagedBy(value *GcpEncryption)()
    SetStorageBucket(value AuthorizationSystemResourceable)()
}
