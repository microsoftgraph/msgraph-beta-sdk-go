package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

// DeploymentState 
type DeploymentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the reasons the deployment has its state value. Read-only.
    reasons []DeploymentStateReason;
    // Specifies the requested state of the deployment. Supports a subset of the values for requestedDeploymentStateValue. Possible values are: none, paused, unknownFutureValue.
    requestedValue *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue;
    // Specifies the state of the deployment. Supports a subset of the values for deploymentStateValue. Possible values are: scheduled, offering, paused, unknownFutureValue. Read-only.
    value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue;
}
// NewDeploymentState instantiates a new deploymentState and sets the default values.
func NewDeploymentState()(*DeploymentState) {
    m := &DeploymentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetReasons gets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) GetReasons()([]DeploymentStateReason) {
    if m == nil {
        return nil
    } else {
        return m.reasons
    }
}
// GetRequestedValue gets the requestedValue property value. Specifies the requested state of the deployment. Supports a subset of the values for requestedDeploymentStateValue. Possible values are: none, paused, unknownFutureValue.
func (m *DeploymentState) GetRequestedValue()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue) {
    if m == nil {
        return nil
    } else {
        return m.requestedValue
    }
}
// GetValue gets the value property value. Specifies the state of the deployment. Supports a subset of the values for deploymentStateValue. Possible values are: scheduled, offering, paused, unknownFutureValue. Read-only.
func (m *DeploymentState) GetValue()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeploymentStateReason() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeploymentStateReason, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeploymentStateReason))
            }
            m.SetReasons(res)
        }
        return nil
    }
    res["requestedValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseRequestedDeploymentStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedValue(val.(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue))
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseDeploymentStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue))
        }
        return nil
    }
    return res
}
func (m *DeploymentState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeploymentState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetReasons() != nil {
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
        cast := (*m.GetRequestedValue()).String()
        err := writer.WriteStringValue("requestedValue", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := (*m.GetValue()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReasons sets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) SetReasons(value []DeploymentStateReason)() {
    if m != nil {
        m.reasons = value
    }
}
// SetRequestedValue sets the requestedValue property value. Specifies the requested state of the deployment. Supports a subset of the values for requestedDeploymentStateValue. Possible values are: none, paused, unknownFutureValue.
func (m *DeploymentState) SetRequestedValue(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.RequestedDeploymentStateValue)() {
    if m != nil {
        m.requestedValue = value
    }
}
// SetValue sets the value property value. Specifies the state of the deployment. Supports a subset of the values for deploymentStateValue. Possible values are: scheduled, offering, paused, unknownFutureValue. Read-only.
func (m *DeploymentState) SetValue(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.DeploymentStateValue)() {
    if m != nil {
        m.value = value
    }
}
