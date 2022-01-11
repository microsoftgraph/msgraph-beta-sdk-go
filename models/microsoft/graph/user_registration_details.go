package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserRegistrationDetails 
type UserRegistrationDetails struct {
    Entity
    // 
    isMfaCapable *bool;
    // 
    isMfaRegistered *bool;
    // 
    isPasswordlessCapable *bool;
    // 
    isSsprCapable *bool;
    // 
    isSsprEnabled *bool;
    // 
    isSsprRegistered *bool;
    // 
    methodsRegistered []string;
    // 
    userDisplayName *string;
    // 
    userPrincipalName *string;
}
// NewUserRegistrationDetails instantiates a new userRegistrationDetails and sets the default values.
func NewUserRegistrationDetails()(*UserRegistrationDetails) {
    m := &UserRegistrationDetails{
        Entity: *NewEntity(),
    }
    return m
}
// GetIsMfaCapable gets the isMfaCapable property value. 
func (m *UserRegistrationDetails) GetIsMfaCapable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaCapable
    }
}
// GetIsMfaRegistered gets the isMfaRegistered property value. 
func (m *UserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaRegistered
    }
}
// GetIsPasswordlessCapable gets the isPasswordlessCapable property value. 
func (m *UserRegistrationDetails) GetIsPasswordlessCapable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPasswordlessCapable
    }
}
// GetIsSsprCapable gets the isSsprCapable property value. 
func (m *UserRegistrationDetails) GetIsSsprCapable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSsprCapable
    }
}
// GetIsSsprEnabled gets the isSsprEnabled property value. 
func (m *UserRegistrationDetails) GetIsSsprEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSsprEnabled
    }
}
// GetIsSsprRegistered gets the isSsprRegistered property value. 
func (m *UserRegistrationDetails) GetIsSsprRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSsprRegistered
    }
}
// GetMethodsRegistered gets the methodsRegistered property value. 
func (m *UserRegistrationDetails) GetMethodsRegistered()([]string) {
    if m == nil {
        return nil
    } else {
        return m.methodsRegistered
    }
}
// GetUserDisplayName gets the userDisplayName property value. 
func (m *UserRegistrationDetails) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *UserRegistrationDetails) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isMfaCapable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaCapable(val)
        }
        return nil
    }
    res["isMfaRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaRegistered(val)
        }
        return nil
    }
    res["isPasswordlessCapable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordlessCapable(val)
        }
        return nil
    }
    res["isSsprCapable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprCapable(val)
        }
        return nil
    }
    res["isSsprEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprEnabled(val)
        }
        return nil
    }
    res["isSsprRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprRegistered(val)
        }
        return nil
    }
    res["methodsRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMethodsRegistered(res)
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *UserRegistrationDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserRegistrationDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isMfaCapable", m.GetIsMfaCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaRegistered", m.GetIsMfaRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPasswordlessCapable", m.GetIsPasswordlessCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprCapable", m.GetIsSsprCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprEnabled", m.GetIsSsprEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprRegistered", m.GetIsSsprRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("methodsRegistered", m.GetMethodsRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsMfaCapable sets the isMfaCapable property value. 
func (m *UserRegistrationDetails) SetIsMfaCapable(value *bool)() {
    if m != nil {
        m.isMfaCapable = value
    }
}
// SetIsMfaRegistered sets the isMfaRegistered property value. 
func (m *UserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    if m != nil {
        m.isMfaRegistered = value
    }
}
// SetIsPasswordlessCapable sets the isPasswordlessCapable property value. 
func (m *UserRegistrationDetails) SetIsPasswordlessCapable(value *bool)() {
    if m != nil {
        m.isPasswordlessCapable = value
    }
}
// SetIsSsprCapable sets the isSsprCapable property value. 
func (m *UserRegistrationDetails) SetIsSsprCapable(value *bool)() {
    if m != nil {
        m.isSsprCapable = value
    }
}
// SetIsSsprEnabled sets the isSsprEnabled property value. 
func (m *UserRegistrationDetails) SetIsSsprEnabled(value *bool)() {
    if m != nil {
        m.isSsprEnabled = value
    }
}
// SetIsSsprRegistered sets the isSsprRegistered property value. 
func (m *UserRegistrationDetails) SetIsSsprRegistered(value *bool)() {
    if m != nil {
        m.isSsprRegistered = value
    }
}
// SetMethodsRegistered sets the methodsRegistered property value. 
func (m *UserRegistrationDetails) SetMethodsRegistered(value []string)() {
    if m != nil {
        m.methodsRegistered = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. 
func (m *UserRegistrationDetails) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *UserRegistrationDetails) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
