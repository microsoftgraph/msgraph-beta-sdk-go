package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppRelationship describes a relationship between two mobile apps.
type MobileAppRelationship struct {
    Entity
}
// NewMobileAppRelationship instantiates a new MobileAppRelationship and sets the default values.
func NewMobileAppRelationship()(*MobileAppRelationship) {
    m := &MobileAppRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppRelationshipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MobileAppRelationship) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["sourceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceDisplayName(val)
        }
        return nil
    }
    res["sourceDisplayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceDisplayVersion(val)
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["sourcePublisherDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourcePublisherDisplayName(val)
        }
        return nil
    }
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
    res["targetPublisherDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPublisherDisplayName(val)
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
// GetSourceDisplayName gets the sourceDisplayName property value. The display name of the app that is the source of the mobile app relationship entity. For example: Orca. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
// returns a *string when successful
func (m *MobileAppRelationship) GetSourceDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("sourceDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourceDisplayVersion gets the sourceDisplayVersion property value. The display version of the app that is the source of the mobile app relationship entity. For example 1.0.12 or 1.2203.156 or 3. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
// returns a *string when successful
func (m *MobileAppRelationship) GetSourceDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("sourceDisplayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourceId gets the sourceId property value. The unique app identifier of the source of the mobile app relationship entity. For example: 2dbc75b9-e993-4e4d-a071-91ac5a218672. If null during relationship creation, then it will be populated with parent Id. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
// returns a *string when successful
func (m *MobileAppRelationship) GetSourceId()(*string) {
    val, err := m.GetBackingStore().Get("sourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourcePublisherDisplayName gets the sourcePublisherDisplayName property value. The publisher display name of the app that is the source of the mobile app relationship entity. For example: Fabrikam. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
// returns a *string when successful
func (m *MobileAppRelationship) GetSourcePublisherDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("sourcePublisherDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetDisplayName gets the targetDisplayName property value. The display name of the app that is the target of the mobile app relationship entity. Read-Only. This property is read-only.
// returns a *string when successful
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
// GetTargetDisplayVersion gets the targetDisplayVersion property value. The display version of the app that is the target of the mobile app relationship entity. Read-Only. This property is read-only.
// returns a *string when successful
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
// GetTargetId gets the targetId property value. App ID of the app that is the target of the mobile app relationship entity. Read-Only
// returns a *string when successful
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
// GetTargetPublisher gets the targetPublisher property value. The publisher of the app that is the target of the mobile app relationship entity. Read-Only. This property is read-only.
// returns a *string when successful
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
// GetTargetPublisherDisplayName gets the targetPublisherDisplayName property value. The publisher display name of the app that is the target of the mobile app relationship entity. For example: Fabrikam. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
// returns a *string when successful
func (m *MobileAppRelationship) GetTargetPublisherDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("targetPublisherDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetType gets the targetType property value. Indicates whether the target of a relationship is the parent or the child in the relationship.
// returns a *MobileAppRelationshipType when successful
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
        err = writer.WriteStringValue("targetId", m.GetTargetId())
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
// SetSourceDisplayName sets the sourceDisplayName property value. The display name of the app that is the source of the mobile app relationship entity. For example: Orca. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
func (m *MobileAppRelationship) SetSourceDisplayName(value *string)() {
    err := m.GetBackingStore().Set("sourceDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceDisplayVersion sets the sourceDisplayVersion property value. The display version of the app that is the source of the mobile app relationship entity. For example 1.0.12 or 1.2203.156 or 3. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
func (m *MobileAppRelationship) SetSourceDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("sourceDisplayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceId sets the sourceId property value. The unique app identifier of the source of the mobile app relationship entity. For example: 2dbc75b9-e993-4e4d-a071-91ac5a218672. If null during relationship creation, then it will be populated with parent Id. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
func (m *MobileAppRelationship) SetSourceId(value *string)() {
    err := m.GetBackingStore().Set("sourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetSourcePublisherDisplayName sets the sourcePublisherDisplayName property value. The publisher display name of the app that is the source of the mobile app relationship entity. For example: Fabrikam. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
func (m *MobileAppRelationship) SetSourcePublisherDisplayName(value *string)() {
    err := m.GetBackingStore().Set("sourcePublisherDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetDisplayName sets the targetDisplayName property value. The display name of the app that is the target of the mobile app relationship entity. Read-Only. This property is read-only.
func (m *MobileAppRelationship) SetTargetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("targetDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetDisplayVersion sets the targetDisplayVersion property value. The display version of the app that is the target of the mobile app relationship entity. Read-Only. This property is read-only.
func (m *MobileAppRelationship) SetTargetDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("targetDisplayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetId sets the targetId property value. App ID of the app that is the target of the mobile app relationship entity. Read-Only
func (m *MobileAppRelationship) SetTargetId(value *string)() {
    err := m.GetBackingStore().Set("targetId", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetPublisher sets the targetPublisher property value. The publisher of the app that is the target of the mobile app relationship entity. Read-Only. This property is read-only.
func (m *MobileAppRelationship) SetTargetPublisher(value *string)() {
    err := m.GetBackingStore().Set("targetPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetPublisherDisplayName sets the targetPublisherDisplayName property value. The publisher display name of the app that is the target of the mobile app relationship entity. For example: Fabrikam. Maximum length is 500 characters. Read-Only. Supports: $select. Does not support $search, $filter, $orderBy. This property is read-only.
func (m *MobileAppRelationship) SetTargetPublisherDisplayName(value *string)() {
    err := m.GetBackingStore().Set("targetPublisherDisplayName", value)
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
type MobileAppRelationshipable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSourceDisplayName()(*string)
    GetSourceDisplayVersion()(*string)
    GetSourceId()(*string)
    GetSourcePublisherDisplayName()(*string)
    GetTargetDisplayName()(*string)
    GetTargetDisplayVersion()(*string)
    GetTargetId()(*string)
    GetTargetPublisher()(*string)
    GetTargetPublisherDisplayName()(*string)
    GetTargetType()(*MobileAppRelationshipType)
    SetSourceDisplayName(value *string)()
    SetSourceDisplayVersion(value *string)()
    SetSourceId(value *string)()
    SetSourcePublisherDisplayName(value *string)()
    SetTargetDisplayName(value *string)()
    SetTargetDisplayVersion(value *string)()
    SetTargetId(value *string)()
    SetTargetPublisher(value *string)()
    SetTargetPublisherDisplayName(value *string)()
    SetTargetType(value *MobileAppRelationshipType)()
}
