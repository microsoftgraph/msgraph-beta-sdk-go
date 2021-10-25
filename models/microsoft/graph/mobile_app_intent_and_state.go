package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobileAppIntentAndState struct {
    Entity
    managedDeviceIdentifier *string;
    mobileAppList []MobileAppIntentAndStateDetail;
    userId *string;
}
func NewMobileAppIntentAndState()(*MobileAppIntentAndState) {
    m := &MobileAppIntentAndState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MobileAppIntentAndState) GetManagedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceIdentifier
    }
}
func (m *MobileAppIntentAndState) GetMobileAppList()([]MobileAppIntentAndStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppList
    }
}
func (m *MobileAppIntentAndState) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *MobileAppIntentAndState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceIdentifier(val)
        return nil
    }
    res["mobileAppList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppIntentAndStateDetail() })
        if err != nil {
            return err
        }
        res := make([]MobileAppIntentAndStateDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppIntentAndStateDetail))
        }
        m.SetMobileAppList(res)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *MobileAppIntentAndState) IsNil()(bool) {
    return m == nil
}
func (m *MobileAppIntentAndState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppList()))
        for i, v := range m.GetMobileAppList() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MobileAppIntentAndState) SetManagedDeviceIdentifier(value *string)() {
    m.managedDeviceIdentifier = value
}
func (m *MobileAppIntentAndState) SetMobileAppList(value []MobileAppIntentAndStateDetail)() {
    m.mobileAppList = value
}
func (m *MobileAppIntentAndState) SetUserId(value *string)() {
    m.userId = value
}
