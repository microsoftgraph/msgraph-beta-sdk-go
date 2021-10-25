package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

type MonitoringRule struct {
    action *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringAction;
    additionalData map[string]interface{};
    signal *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringSignal;
    threshold *int32;
}
func NewMonitoringRule()(*MonitoringRule) {
    m := &MonitoringRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MonitoringRule) GetAction()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *MonitoringRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MonitoringRule) GetSignal()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringSignal) {
    if m == nil {
        return nil
    } else {
        return m.signal
    }
}
func (m *MonitoringRule) GetThreshold()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.threshold
    }
}
func (m *MonitoringRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseMonitoringAction)
        if err != nil {
            return err
        }
        cast := val.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringAction)
        m.SetAction(&cast)
        return nil
    }
    res["signal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseMonitoringSignal)
        if err != nil {
            return err
        }
        cast := val.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringSignal)
        m.SetSignal(&cast)
        return nil
    }
    res["threshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetThreshold(val)
        return nil
    }
    return res
}
func (m *MonitoringRule) IsNil()(bool) {
    return m == nil
}
func (m *MonitoringRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignal() != nil {
        cast := m.GetSignal().String()
        err := writer.WriteStringValue("signal", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("threshold", m.GetThreshold())
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
func (m *MonitoringRule) SetAction(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringAction)() {
    m.action = value
}
func (m *MonitoringRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MonitoringRule) SetSignal(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.MonitoringSignal)() {
    m.signal = value
}
func (m *MonitoringRule) SetThreshold(value *int32)() {
    m.threshold = value
}
