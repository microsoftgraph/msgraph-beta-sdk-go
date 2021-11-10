package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssignmentFilterEvaluateRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Platform type of the devices on which the Assignment Filter will be applicable. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
    platform *DevicePlatformType;
    // Rule definition of the Assignment Filter.
    rule *string;
    // Number of records to skip. Default value is 0
    skip *int32;
    // Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
    top *int32;
}
// Instantiates a new assignmentFilterEvaluateRequest and sets the default values.
func NewAssignmentFilterEvaluateRequest()(*AssignmentFilterEvaluateRequest) {
    m := &AssignmentFilterEvaluateRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluateRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the platform property value. Platform type of the devices on which the Assignment Filter will be applicable. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *AssignmentFilterEvaluateRequest) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Gets the rule property value. Rule definition of the Assignment Filter.
func (m *AssignmentFilterEvaluateRequest) GetRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rule
    }
}
// Gets the skip property value. Number of records to skip. Default value is 0
func (m *AssignmentFilterEvaluateRequest) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// Gets the top property value. Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
func (m *AssignmentFilterEvaluateRequest) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// The deserialization information for the current model
func (m *AssignmentFilterEvaluateRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DevicePlatformType)
            m.SetPlatform(&cast)
        }
        return nil
    }
    res["rule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val)
        }
        return nil
    }
    res["skip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
func (m *AssignmentFilterEvaluateRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssignmentFilterEvaluateRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetPlatform() != nil {
        cast := m.GetPlatform().String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rule", m.GetRule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *AssignmentFilterEvaluateRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the platform property value. Platform type of the devices on which the Assignment Filter will be applicable. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
// Parameters:
//  - value : Value to set for the platform property.
func (m *AssignmentFilterEvaluateRequest) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
// Sets the rule property value. Rule definition of the Assignment Filter.
// Parameters:
//  - value : Value to set for the rule property.
func (m *AssignmentFilterEvaluateRequest) SetRule(value *string)() {
    m.rule = value
}
// Sets the skip property value. Number of records to skip. Default value is 0
// Parameters:
//  - value : Value to set for the skip property.
func (m *AssignmentFilterEvaluateRequest) SetSkip(value *int32)() {
    m.skip = value
}
// Sets the top property value. Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
// Parameters:
//  - value : Value to set for the top property.
func (m *AssignmentFilterEvaluateRequest) SetTop(value *int32)() {
    m.top = value
}
