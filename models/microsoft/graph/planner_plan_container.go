package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerPlanContainer struct {
    additionalData map[string]interface{};
    containerId *string;
    type_escaped *PlannerContainerType;
    url *string;
}
func NewPlannerPlanContainer()(*PlannerPlanContainer) {
    m := &PlannerPlanContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PlannerPlanContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PlannerPlanContainer) GetContainerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerId
    }
}
func (m *PlannerPlanContainer) GetType_escaped()(*PlannerContainerType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *PlannerPlanContainer) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *PlannerPlanContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["containerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContainerId(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerContainerType)
        if err != nil {
            return err
        }
        cast := val.(PlannerContainerType)
        m.SetType_escaped(&cast)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *PlannerPlanContainer) IsNil()(bool) {
    return m == nil
}
func (m *PlannerPlanContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerId", m.GetContainerId())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *PlannerPlanContainer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PlannerPlanContainer) SetContainerId(value *string)() {
    m.containerId = value
}
func (m *PlannerPlanContainer) SetType_escaped(value *PlannerContainerType)() {
    m.type_escaped = value
}
func (m *PlannerPlanContainer) SetUrl(value *string)() {
    m.url = value
}
