package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemporaryAccessPassAuthenticationMethod 
type TemporaryAccessPassAuthenticationMethod struct {
    AuthenticationMethod
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The date and time when the Temporary Access Pass was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The state of the authentication method that indicates whether it's currently usable by the user.
    isUsable *bool
    // Determines whether the pass is limited to a one-time use. If true, the pass can be used once; if false, the pass can be used multiple times within the Temporary Access Pass lifetime.
    isUsableOnce *bool
    // The lifetime of the Temporary Access Pass in minutes starting at startDateTime. Must be between 10 and 43200 inclusive (equivalent to 30 days).
    lifetimeInMinutes *int32
    // Details about the usability state (isUsable). Reasons can include: EnabledByPolicy, DisabledByPolicy, Expired, NotYetValid, OneTimeUsed.
    methodUsabilityReason *string
    // The date and time when the Temporary Access Pass becomes available to use and when isUsable is true is enforced.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Temporary Access Pass used to authenticate. Returned only on creation of a new temporaryAccessPassAuthenticationMethod object; Hidden in subsequent read operations and returned as null with GET.
    temporaryAccessPass *string
}
// NewTemporaryAccessPassAuthenticationMethod instantiates a new TemporaryAccessPassAuthenticationMethod and sets the default values.
func NewTemporaryAccessPassAuthenticationMethod()(*TemporaryAccessPassAuthenticationMethod) {
    m := &TemporaryAccessPassAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemporaryAccessPassAuthenticationMethod(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemporaryAccessPassAuthenticationMethod) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the Temporary Access Pass was created.
func (m *TemporaryAccessPassAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TemporaryAccessPassAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isUsable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsable(val)
        }
        return nil
    }
    res["isUsableOnce"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsableOnce(val)
        }
        return nil
    }
    res["lifetimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifetimeInMinutes(val)
        }
        return nil
    }
    res["methodUsabilityReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodUsabilityReason(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["temporaryAccessPass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIsUsable gets the isUsable property value. The state of the authentication method that indicates whether it's currently usable by the user.
func (m *TemporaryAccessPassAuthenticationMethod) GetIsUsable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUsable
    }
}
// GetIsUsableOnce gets the isUsableOnce property value. Determines whether the pass is limited to a one-time use. If true, the pass can be used once; if false, the pass can be used multiple times within the Temporary Access Pass lifetime.
func (m *TemporaryAccessPassAuthenticationMethod) GetIsUsableOnce()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUsableOnce
    }
}
// GetLifetimeInMinutes gets the lifetimeInMinutes property value. The lifetime of the Temporary Access Pass in minutes starting at startDateTime. Must be between 10 and 43200 inclusive (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethod) GetLifetimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lifetimeInMinutes
    }
}
// GetMethodUsabilityReason gets the methodUsabilityReason property value. Details about the usability state (isUsable). Reasons can include: EnabledByPolicy, DisabledByPolicy, Expired, NotYetValid, OneTimeUsed.
func (m *TemporaryAccessPassAuthenticationMethod) GetMethodUsabilityReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.methodUsabilityReason
    }
}
// GetStartDateTime gets the startDateTime property value. The date and time when the Temporary Access Pass becomes available to use and when isUsable is true is enforced.
func (m *TemporaryAccessPassAuthenticationMethod) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetTemporaryAccessPass gets the temporaryAccessPass property value. The Temporary Access Pass used to authenticate. Returned only on creation of a new temporaryAccessPassAuthenticationMethod object; Hidden in subsequent read operations and returned as null with GET.
func (m *TemporaryAccessPassAuthenticationMethod) GetTemporaryAccessPass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.temporaryAccessPass
    }
}
// Serialize serializes information the current object
func (m *TemporaryAccessPassAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemporaryAccessPassAuthenticationMethod) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the Temporary Access Pass was created.
func (m *TemporaryAccessPassAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetIsUsable sets the isUsable property value. The state of the authentication method that indicates whether it's currently usable by the user.
func (m *TemporaryAccessPassAuthenticationMethod) SetIsUsable(value *bool)() {
    if m != nil {
        m.isUsable = value
    }
}
// SetIsUsableOnce sets the isUsableOnce property value. Determines whether the pass is limited to a one-time use. If true, the pass can be used once; if false, the pass can be used multiple times within the Temporary Access Pass lifetime.
func (m *TemporaryAccessPassAuthenticationMethod) SetIsUsableOnce(value *bool)() {
    if m != nil {
        m.isUsableOnce = value
    }
}
// SetLifetimeInMinutes sets the lifetimeInMinutes property value. The lifetime of the Temporary Access Pass in minutes starting at startDateTime. Must be between 10 and 43200 inclusive (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethod) SetLifetimeInMinutes(value *int32)() {
    if m != nil {
        m.lifetimeInMinutes = value
    }
}
// SetMethodUsabilityReason sets the methodUsabilityReason property value. Details about the usability state (isUsable). Reasons can include: EnabledByPolicy, DisabledByPolicy, Expired, NotYetValid, OneTimeUsed.
func (m *TemporaryAccessPassAuthenticationMethod) SetMethodUsabilityReason(value *string)() {
    if m != nil {
        m.methodUsabilityReason = value
    }
}
// SetStartDateTime sets the startDateTime property value. The date and time when the Temporary Access Pass becomes available to use and when isUsable is true is enforced.
func (m *TemporaryAccessPassAuthenticationMethod) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetTemporaryAccessPass sets the temporaryAccessPass property value. The Temporary Access Pass used to authenticate. Returned only on creation of a new temporaryAccessPassAuthenticationMethod object; Hidden in subsequent read operations and returned as null with GET.
func (m *TemporaryAccessPassAuthenticationMethod) SetTemporaryAccessPass(value *string)() {
    if m != nil {
        m.temporaryAccessPass = value
    }
}
