package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobileAppRelationship struct {
    Entity
    targetDisplayName *string;
    targetDisplayVersion *string;
    targetId *string;
    targetPublisher *string;
    targetType *MobileAppRelationshipType;
}
func NewMobileAppRelationship()(*MobileAppRelationship) {
    m := &MobileAppRelationship{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MobileAppRelationship) GetTargetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayName
    }
}
func (m *MobileAppRelationship) GetTargetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayVersion
    }
}
func (m *MobileAppRelationship) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
func (m *MobileAppRelationship) GetTargetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetPublisher
    }
}
func (m *MobileAppRelationship) GetTargetType()(*MobileAppRelationshipType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
func (m *MobileAppRelationship) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetDisplayName(val)
        return nil
    }
    res["targetDisplayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetDisplayVersion(val)
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetId(val)
        return nil
    }
    res["targetPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetPublisher(val)
        return nil
    }
    res["targetType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppRelationshipType)
        if err != nil {
            return err
        }
        cast := val.(MobileAppRelationshipType)
        m.SetTargetType(&cast)
        return nil
    }
    return res
}
func (m *MobileAppRelationship) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetTargetType().String()
        err = writer.WriteStringValue("targetType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MobileAppRelationship) SetTargetDisplayName(value *string)() {
    m.targetDisplayName = value
}
func (m *MobileAppRelationship) SetTargetDisplayVersion(value *string)() {
    m.targetDisplayVersion = value
}
func (m *MobileAppRelationship) SetTargetId(value *string)() {
    m.targetId = value
}
func (m *MobileAppRelationship) SetTargetPublisher(value *string)() {
    m.targetPublisher = value
}
func (m *MobileAppRelationship) SetTargetType(value *MobileAppRelationshipType)() {
    m.targetType = value
}
