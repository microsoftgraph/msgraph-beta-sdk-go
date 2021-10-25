package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConfigurationManagerClientHealthState struct {
    additionalData map[string]interface{};
    errorCode *int32;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    state *ConfigurationManagerClientState;
}
func NewConfigurationManagerClientHealthState()(*ConfigurationManagerClientHealthState) {
    m := &ConfigurationManagerClientHealthState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConfigurationManagerClientHealthState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConfigurationManagerClientHealthState) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *ConfigurationManagerClientHealthState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *ConfigurationManagerClientHealthState) GetState()(*ConfigurationManagerClientState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *ConfigurationManagerClientHealthState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerClientState)
        if err != nil {
            return err
        }
        cast := val.(ConfigurationManagerClientState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *ConfigurationManagerClientHealthState) IsNil()(bool) {
    return m == nil
}
func (m *ConfigurationManagerClientHealthState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ConfigurationManagerClientHealthState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConfigurationManagerClientHealthState) SetErrorCode(value *int32)() {
    m.errorCode = value
}
func (m *ConfigurationManagerClientHealthState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *ConfigurationManagerClientHealthState) SetState(value *ConfigurationManagerClientState)() {
    m.state = value
}
