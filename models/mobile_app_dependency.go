package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppDependency describes a dependency type between two mobile apps.
type MobileAppDependency struct {
    MobileAppRelationship
}
// NewMobileAppDependency instantiates a new MobileAppDependency and sets the default values.
func NewMobileAppDependency()(*MobileAppDependency) {
    m := &MobileAppDependency{
        MobileAppRelationship: *NewMobileAppRelationship(),
    }
    odataTypeValue := "#microsoft.graph.mobileAppDependency"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMobileAppDependencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppDependencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppDependency(), nil
}
// GetDependencyType gets the dependencyType property value. Indicates the dependency type associated with a relationship between two mobile apps.
// returns a *MobileAppDependencyType when successful
func (m *MobileAppDependency) GetDependencyType()(*MobileAppDependencyType) {
    val, err := m.GetBackingStore().Get("dependencyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileAppDependencyType)
    }
    return nil
}
// GetDependentAppCount gets the dependentAppCount property value. The total number of apps that directly or indirectly depend on the parent app. Read-Only. This property is read-only.
// returns a *int32 when successful
func (m *MobileAppDependency) GetDependentAppCount()(*int32) {
    val, err := m.GetBackingStore().Get("dependentAppCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDependsOnAppCount gets the dependsOnAppCount property value. The total number of apps the child app directly or indirectly depends on. Read-Only. This property is read-only.
// returns a *int32 when successful
func (m *MobileAppDependency) GetDependsOnAppCount()(*int32) {
    val, err := m.GetBackingStore().Get("dependsOnAppCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MobileAppDependency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppRelationship.GetFieldDeserializers()
    res["dependencyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppDependencyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyType(val.(*MobileAppDependencyType))
        }
        return nil
    }
    res["dependentAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependentAppCount(val)
        }
        return nil
    }
    res["dependsOnAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependsOnAppCount(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MobileAppDependency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppRelationship.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDependencyType() != nil {
        cast := (*m.GetDependencyType()).String()
        err = writer.WriteStringValue("dependencyType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDependencyType sets the dependencyType property value. Indicates the dependency type associated with a relationship between two mobile apps.
func (m *MobileAppDependency) SetDependencyType(value *MobileAppDependencyType)() {
    err := m.GetBackingStore().Set("dependencyType", value)
    if err != nil {
        panic(err)
    }
}
// SetDependentAppCount sets the dependentAppCount property value. The total number of apps that directly or indirectly depend on the parent app. Read-Only. This property is read-only.
func (m *MobileAppDependency) SetDependentAppCount(value *int32)() {
    err := m.GetBackingStore().Set("dependentAppCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDependsOnAppCount sets the dependsOnAppCount property value. The total number of apps the child app directly or indirectly depends on. Read-Only. This property is read-only.
func (m *MobileAppDependency) SetDependsOnAppCount(value *int32)() {
    err := m.GetBackingStore().Set("dependsOnAppCount", value)
    if err != nil {
        panic(err)
    }
}
type MobileAppDependencyable interface {
    MobileAppRelationshipable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDependencyType()(*MobileAppDependencyType)
    GetDependentAppCount()(*int32)
    GetDependsOnAppCount()(*int32)
    SetDependencyType(value *MobileAppDependencyType)()
    SetDependentAppCount(value *int32)()
    SetDependsOnAppCount(value *int32)()
}
