package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppRelationship describes a relationship between two mobile apps.
type MobileAppRelationship struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewMobileAppRelationship instantiates a new mobileAppRelationship and sets the default values.
func NewMobileAppRelationship()(*MobileAppRelationship) {
    m := &MobileAppRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppRelationshipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppRelationshipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.mobileAppDependency":
                        return NewMobileAppDependency(), nil
                    case "#microsoft.graph.mobileAppSupersedence":
                        return NewMobileAppSupersedence(), nil
                }
            }
        }
    }
    return NewMobileAppRelationship(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppRelationship) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayName(val)
        }
        return nil
    }
    res["targetDisplayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayVersion(val)
        }
        return nil
    }
    res["targetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["targetPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPublisher(val)
        }
        return nil
    }
    res["targetType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppRelationshipType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*MobileAppRelationshipType))
        }
        return nil
    }
    return res
}
// GetTargetDisplayName gets the targetDisplayName property value. The target mobile app's display name.
func (m *MobileAppRelationship) GetTargetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("targetDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetDisplayVersion gets the targetDisplayVersion property value. The target mobile app's display version.
func (m *MobileAppRelationship) GetTargetDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("targetDisplayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetId gets the targetId property value. The target mobile app's app id.
func (m *MobileAppRelationship) GetTargetId()(*string) {
    val, err := m.GetBackingStore().Get("targetId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetPublisher gets the targetPublisher property value. The target mobile app's publisher.
func (m *MobileAppRelationship) GetTargetPublisher()(*string) {
    val, err := m.GetBackingStore().Get("targetPublisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetType gets the targetType property value. Indicates whether the target of a relationship is the parent or the child in the relationship.
func (m *MobileAppRelationship) GetTargetType()(*MobileAppRelationshipType) {
    val, err := m.GetBackingStore().Get("targetType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileAppRelationshipType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppRelationship) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetDisplayName", m.GetTargetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetDisplayVersion", m.GetTargetDisplayVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetId", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetPublisher", m.GetTargetPublisher())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err = writer.WriteStringValue("targetType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetDisplayName sets the targetDisplayName property value. The target mobile app's display name.
func (m *MobileAppRelationship) SetTargetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("targetDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetDisplayVersion sets the targetDisplayVersion property value. The target mobile app's display version.
func (m *MobileAppRelationship) SetTargetDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("targetDisplayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetId sets the targetId property value. The target mobile app's app id.
func (m *MobileAppRelationship) SetTargetId(value *string)() {
    err := m.GetBackingStore().Set("targetId", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetPublisher sets the targetPublisher property value. The target mobile app's publisher.
func (m *MobileAppRelationship) SetTargetPublisher(value *string)() {
    err := m.GetBackingStore().Set("targetPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetType sets the targetType property value. Indicates whether the target of a relationship is the parent or the child in the relationship.
func (m *MobileAppRelationship) SetTargetType(value *MobileAppRelationshipType)() {
    err := m.GetBackingStore().Set("targetType", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppRelationshipable 
type MobileAppRelationshipable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTargetDisplayName()(*string)
    GetTargetDisplayVersion()(*string)
    GetTargetId()(*string)
    GetTargetPublisher()(*string)
    GetTargetType()(*MobileAppRelationshipType)
    SetTargetDisplayName(value *string)()
    SetTargetDisplayVersion(value *string)()
    SetTargetId(value *string)()
    SetTargetPublisher(value *string)()
    SetTargetType(value *MobileAppRelationshipType)()
}
