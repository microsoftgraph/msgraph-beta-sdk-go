package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CompanyPortalBlockedAction struct {
    action *CompanyPortalAction;
    additionalData map[string]interface{};
    ownerType *OwnerType;
    platform *DevicePlatformType;
}
func NewCompanyPortalBlockedAction()(*CompanyPortalBlockedAction) {
    m := &CompanyPortalBlockedAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CompanyPortalBlockedAction) GetAction()(*CompanyPortalAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *CompanyPortalBlockedAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CompanyPortalBlockedAction) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
func (m *CompanyPortalBlockedAction) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
func (m *CompanyPortalBlockedAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCompanyPortalAction)
        if err != nil {
            return err
        }
        cast := val.(CompanyPortalAction)
        m.SetAction(&cast)
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        cast := val.(OwnerType)
        m.SetOwnerType(&cast)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        cast := val.(DevicePlatformType)
        m.SetPlatform(&cast)
        return nil
    }
    return res
}
func (m *CompanyPortalBlockedAction) IsNil()(bool) {
    return m == nil
}
func (m *CompanyPortalBlockedAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnerType() != nil {
        cast := m.GetOwnerType().String()
        err := writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := m.GetPlatform().String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CompanyPortalBlockedAction) SetAction(value *CompanyPortalAction)() {
    m.action = value
}
func (m *CompanyPortalBlockedAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CompanyPortalBlockedAction) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
func (m *CompanyPortalBlockedAction) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
