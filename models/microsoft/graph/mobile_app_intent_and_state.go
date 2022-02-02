package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppIntentAndState 
type MobileAppIntentAndState struct {
    Entity
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string;
    // The list of payload intents and states for the tenant.
    mobileAppList []MobileAppIntentAndStateDetail;
    // Identifier for the user that tried to enroll the device.
    userId *string;
}
// NewMobileAppIntentAndState instantiates a new mobileAppIntentAndState and sets the default values.
func NewMobileAppIntentAndState()(*MobileAppIntentAndState) {
    m := &MobileAppIntentAndState{
        Entity: *NewEntity(),
    }
    return m
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppIntentAndState) GetManagedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceIdentifier
    }
}
// GetMobileAppList gets the mobileAppList property value. The list of payload intents and states for the tenant.
func (m *MobileAppIntentAndState) GetMobileAppList()([]MobileAppIntentAndStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppList
    }
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppIntentAndState) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIntentAndState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceIdentifier(val)
        }
        return nil
    }
    res["mobileAppList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppIntentAndStateDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppIntentAndStateDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppIntentAndStateDetail))
            }
            m.SetMobileAppList(res)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *MobileAppIntentAndState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetMobileAppList() != nil {
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
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppIntentAndState) SetManagedDeviceIdentifier(value *string)() {
    if m != nil {
        m.managedDeviceIdentifier = value
    }
}
// SetMobileAppList sets the mobileAppList property value. The list of payload intents and states for the tenant.
func (m *MobileAppIntentAndState) SetMobileAppList(value []MobileAppIntentAndStateDetail)() {
    if m != nil {
        m.mobileAppList = value
    }
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppIntentAndState) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
