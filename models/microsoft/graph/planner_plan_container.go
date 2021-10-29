package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerPlanContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identifier of the resource that contains the plan.
    containerId *string;
    // The type of the resource that contains the plan. See the previous table for supported types. Possible values are: group, unknownFutureValue, roster. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: roster.
    type_escaped *PlannerContainerType;
    // The full canonical URL of the container.
    url *string;
}
// Instantiates a new plannerPlanContainer and sets the default values.
func NewPlannerPlanContainer()(*PlannerPlanContainer) {
    m := &PlannerPlanContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerPlanContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the containerId property value. The identifier of the resource that contains the plan.
func (m *PlannerPlanContainer) GetContainerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerId
    }
}
// Gets the type_escaped property value. The type of the resource that contains the plan. See the previous table for supported types. Possible values are: group, unknownFutureValue, roster. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: roster.
func (m *PlannerPlanContainer) GetType_escaped()(*PlannerContainerType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the url property value. The full canonical URL of the container.
func (m *PlannerPlanContainer) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PlannerPlanContainer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the containerId property value. The identifier of the resource that contains the plan.
// Parameters:
//  - value : Value to set for the containerId property.
func (m *PlannerPlanContainer) SetContainerId(value *string)() {
    m.containerId = value
}
// Sets the type_escaped property value. The type of the resource that contains the plan. See the previous table for supported types. Possible values are: group, unknownFutureValue, roster. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: roster.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *PlannerPlanContainer) SetType_escaped(value *PlannerContainerType)() {
    m.type_escaped = value
}
// Sets the url property value. The full canonical URL of the container.
// Parameters:
//  - value : Value to set for the url property.
func (m *PlannerPlanContainer) SetUrl(value *string)() {
    m.url = value
}
