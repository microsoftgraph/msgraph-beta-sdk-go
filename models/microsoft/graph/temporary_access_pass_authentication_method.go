package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TemporaryAccessPassAuthenticationMethod struct {
    AuthenticationMethod
    // The date and time when the temporaryAccessPass was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The state of the authentication method that indicates whether it's currently usable by the user.
    isUsable *bool;
    // Determines whether the pass is limited to a one time use. If true, the pass can be used once; if false, the pass can be used multiple times within the temporaryAccessPass lifetime.
    isUsableOnce *bool;
    // The lifetime of the temporaryAccessPass in minutes starting at startDateTime. Minimum 10, Maximum 43200 (equivalent to 30 days).
    lifetimeInMinutes *int32;
    // Details about usability state (isUsable). Reasons can include: enabledByPolicy, disabledByPolicy, expired, notYetValid, oneTimeUsed.
    methodUsabilityReason *string;
    // The date and time when the temporaryAccessPass becomes available to use.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The temporaryAccessPass used to authenticate. Returned only on creation of a new temporaryAccessPass; returned as NULL with GET.
    temporaryAccessPass *string;
}
// Instantiates a new temporaryAccessPassAuthenticationMethod and sets the default values.
func NewTemporaryAccessPassAuthenticationMethod()(*TemporaryAccessPassAuthenticationMethod) {
    m := &TemporaryAccessPassAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// Gets the createdDateTime property value. The date and time when the temporaryAccessPass was created.
func (m *TemporaryAccessPassAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the isUsable property value. The state of the authentication method that indicates whether it's currently usable by the user.
func (m *TemporaryAccessPassAuthenticationMethod) GetIsUsable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUsable
    }
}
// Gets the isUsableOnce property value. Determines whether the pass is limited to a one time use. If true, the pass can be used once; if false, the pass can be used multiple times within the temporaryAccessPass lifetime.
func (m *TemporaryAccessPassAuthenticationMethod) GetIsUsableOnce()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUsableOnce
    }
}
// Gets the lifetimeInMinutes property value. The lifetime of the temporaryAccessPass in minutes starting at startDateTime. Minimum 10, Maximum 43200 (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethod) GetLifetimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lifetimeInMinutes
    }
}
// Gets the methodUsabilityReason property value. Details about usability state (isUsable). Reasons can include: enabledByPolicy, disabledByPolicy, expired, notYetValid, oneTimeUsed.
func (m *TemporaryAccessPassAuthenticationMethod) GetMethodUsabilityReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.methodUsabilityReason
    }
}
// Gets the startDateTime property value. The date and time when the temporaryAccessPass becomes available to use.
func (m *TemporaryAccessPassAuthenticationMethod) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the temporaryAccessPass property value. The temporaryAccessPass used to authenticate. Returned only on creation of a new temporaryAccessPass; returned as NULL with GET.
func (m *TemporaryAccessPassAuthenticationMethod) GetTemporaryAccessPass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.temporaryAccessPass
    }
}
// The deserialization information for the current model
func (m *TemporaryAccessPassAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isUsable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsable(val)
        }
        return nil
    }
    res["isUsableOnce"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsableOnce(val)
        }
        return nil
    }
    res["lifetimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifetimeInMinutes(val)
        }
        return nil
    }
    res["methodUsabilityReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodUsabilityReason(val)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["temporaryAccessPass"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporaryAccessPass(val)
        }
        return nil
    }
    return res
}
func (m *TemporaryAccessPassAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TemporaryAccessPassAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isUsable", m.GetIsUsable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isUsableOnce", m.GetIsUsableOnce())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("lifetimeInMinutes", m.GetLifetimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("methodUsabilityReason", m.GetMethodUsabilityReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("temporaryAccessPass", m.GetTemporaryAccessPass())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. The date and time when the temporaryAccessPass was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *TemporaryAccessPassAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the isUsable property value. The state of the authentication method that indicates whether it's currently usable by the user.
// Parameters:
//  - value : Value to set for the isUsable property.
func (m *TemporaryAccessPassAuthenticationMethod) SetIsUsable(value *bool)() {
    m.isUsable = value
}
// Sets the isUsableOnce property value. Determines whether the pass is limited to a one time use. If true, the pass can be used once; if false, the pass can be used multiple times within the temporaryAccessPass lifetime.
// Parameters:
//  - value : Value to set for the isUsableOnce property.
func (m *TemporaryAccessPassAuthenticationMethod) SetIsUsableOnce(value *bool)() {
    m.isUsableOnce = value
}
// Sets the lifetimeInMinutes property value. The lifetime of the temporaryAccessPass in minutes starting at startDateTime. Minimum 10, Maximum 43200 (equivalent to 30 days).
// Parameters:
//  - value : Value to set for the lifetimeInMinutes property.
func (m *TemporaryAccessPassAuthenticationMethod) SetLifetimeInMinutes(value *int32)() {
    m.lifetimeInMinutes = value
}
// Sets the methodUsabilityReason property value. Details about usability state (isUsable). Reasons can include: enabledByPolicy, disabledByPolicy, expired, notYetValid, oneTimeUsed.
// Parameters:
//  - value : Value to set for the methodUsabilityReason property.
func (m *TemporaryAccessPassAuthenticationMethod) SetMethodUsabilityReason(value *string)() {
    m.methodUsabilityReason = value
}
// Sets the startDateTime property value. The date and time when the temporaryAccessPass becomes available to use.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *TemporaryAccessPassAuthenticationMethod) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the temporaryAccessPass property value. The temporaryAccessPass used to authenticate. Returned only on creation of a new temporaryAccessPass; returned as NULL with GET.
// Parameters:
//  - value : Value to set for the temporaryAccessPass property.
func (m *TemporaryAccessPassAuthenticationMethod) SetTemporaryAccessPass(value *string)() {
    m.temporaryAccessPass = value
}
