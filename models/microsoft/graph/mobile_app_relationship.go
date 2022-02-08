package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppRelationship 
type MobileAppRelationship struct {
    Entity
    // The target mobile app's display name.
    targetDisplayName *string;
    // The target mobile app's display version.
    targetDisplayVersion *string;
    // The target mobile app's app id.
    targetId *string;
    // The target mobile app's publisher.
    targetPublisher *string;
    // The type of relationship indicating whether the target is a parent or child. Possible values are: child, parent.
    targetType *MobileAppRelationshipType;
}
// NewMobileAppRelationship instantiates a new mobileAppRelationship and sets the default values.
func NewMobileAppRelationship()(*MobileAppRelationship) {
    m := &MobileAppRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// GetTargetDisplayName gets the targetDisplayName property value. The target mobile app's display name.
func (m *MobileAppRelationship) GetTargetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayName
    }
}
// GetTargetDisplayVersion gets the targetDisplayVersion property value. The target mobile app's display version.
func (m *MobileAppRelationship) GetTargetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayVersion
    }
}
// GetTargetId gets the targetId property value. The target mobile app's app id.
func (m *MobileAppRelationship) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
// GetTargetPublisher gets the targetPublisher property value. The target mobile app's publisher.
func (m *MobileAppRelationship) GetTargetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetPublisher
    }
}
// GetTargetType gets the targetType property value. The type of relationship indicating whether the target is a parent or child. Possible values are: child, parent.
func (m *MobileAppRelationship) GetTargetType()(*MobileAppRelationshipType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppRelationship) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayName(val)
        }
        return nil
    }
    res["targetDisplayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayVersion(val)
        }
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["targetPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPublisher(val)
        }
        return nil
    }
    res["targetType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MobileAppRelationship) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobileAppRelationship) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.targetDisplayName = value
    }
}
// SetTargetDisplayVersion sets the targetDisplayVersion property value. The target mobile app's display version.
func (m *MobileAppRelationship) SetTargetDisplayVersion(value *string)() {
    if m != nil {
        m.targetDisplayVersion = value
    }
}
// SetTargetId sets the targetId property value. The target mobile app's app id.
func (m *MobileAppRelationship) SetTargetId(value *string)() {
    if m != nil {
        m.targetId = value
    }
}
// SetTargetPublisher sets the targetPublisher property value. The target mobile app's publisher.
func (m *MobileAppRelationship) SetTargetPublisher(value *string)() {
    if m != nil {
        m.targetPublisher = value
    }
}
// SetTargetType sets the targetType property value. The type of relationship indicating whether the target is a parent or child. Possible values are: child, parent.
func (m *MobileAppRelationship) SetTargetType(value *MobileAppRelationshipType)() {
    if m != nil {
        m.targetType = value
    }
}
