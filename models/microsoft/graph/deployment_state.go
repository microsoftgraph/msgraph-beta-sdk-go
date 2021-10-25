package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

type DeploymentState struct {
    additionalData map[string]interface{};
    reasons []DeploymentStateReason;
    requestedValue *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue;
    value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue;
}
func NewDeploymentState()(*DeploymentState) {
    m := &DeploymentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeploymentState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeploymentState) GetReasons()([]DeploymentStateReason) {
    if m == nil {
        return nil
    } else {
        return m.reasons
    }
}
func (m *DeploymentState) GetRequestedValue()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue) {
    if m == nil {
        return nil
    } else {
        return m.requestedValue
    }
}
func (m *DeploymentState) GetValue()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *DeploymentState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeploymentStateReason() })
        if err != nil {
            return err
        }
        res := make([]DeploymentStateReason, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeploymentStateReason))
        }
        m.SetReasons(res)
        return nil
    }
    res["requestedValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseRequestedDeploymentStateValue)
        if err != nil {
            return err
        }
        cast := val.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue)
        m.SetRequestedValue(&cast)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseDeploymentStateValue)
        if err != nil {
            return err
        }
        cast := val.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue)
        m.SetValue(&cast)
        return nil
    }
    return res
}
func (m *DeploymentState) IsNil()(bool) {
    return m == nil
}
func (m *DeploymentState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReasons()))
        for i, v := range m.GetReasons() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("reasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestedValue() != nil {
        cast := m.GetRequestedValue().String()
        err := writer.WriteStringValue("requestedValue", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := m.GetValue().String()
        err := writer.WriteStringValue("value", &cast)
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
func (m *DeploymentState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeploymentState) SetReasons(value []DeploymentStateReason)() {
    m.reasons = value
}
func (m *DeploymentState) SetRequestedValue(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue)() {
    m.requestedValue = value
}
func (m *DeploymentState) SetValue(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue)() {
    m.value = value
}
