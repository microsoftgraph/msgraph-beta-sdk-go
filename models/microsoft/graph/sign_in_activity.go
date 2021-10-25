package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SignInActivity struct {
    additionalData map[string]interface{};
    lastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSignInRequestId *string;
}
func NewSignInActivity()(*SignInActivity) {
    m := &SignInActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SignInActivity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SignInActivity) GetLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSignInDateTime
    }
}
func (m *SignInActivity) GetLastSignInRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastSignInRequestId
    }
}
func (m *SignInActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["lastSignInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSignInDateTime(val)
        return nil
    }
    res["lastSignInRequestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastSignInRequestId(val)
        return nil
    }
    return res
}
func (m *SignInActivity) IsNil()(bool) {
    return m == nil
}
func (m *SignInActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastSignInDateTime", m.GetLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastSignInRequestId", m.GetLastSignInRequestId())
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
func (m *SignInActivity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SignInActivity) SetLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSignInDateTime = value
}
func (m *SignInActivity) SetLastSignInRequestId(value *string)() {
    m.lastSignInRequestId = value
}
