package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterTypeAndEvaluationResult 
type AssignmentFilterTypeAndEvaluationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents the filter type. Possible values are: none, include, exclude.
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType;
    // Represents the evalaution result of the filter. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
    evaluationResult *AssignmentFilterEvaluationResult;
}
// NewAssignmentFilterTypeAndEvaluationResult instantiates a new assignmentFilterTypeAndEvaluationResult and sets the default values.
func NewAssignmentFilterTypeAndEvaluationResult()(*AssignmentFilterTypeAndEvaluationResult) {
    m := &AssignmentFilterTypeAndEvaluationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents the filter type. Possible values are: none, include, exclude.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterType
    }
}
// GetEvaluationResult gets the evaluationResult property value. Represents the evalaution result of the filter. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
func (m *AssignmentFilterTypeAndEvaluationResult) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.evaluationResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterTypeAndEvaluationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["evaluationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterEvaluationResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationResult(val.(*AssignmentFilterEvaluationResult))
        }
        return nil
    }
    return res
}
func (m *AssignmentFilterTypeAndEvaluationResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterTypeAndEvaluationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := (*m.GetEvaluationResult()).String()
        err := writer.WriteStringValue("evaluationResult", &cast)
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
func (m *AssignmentFilterTypeAndEvaluationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents the filter type. Possible values are: none, include, exclude.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    if m != nil {
        m.assignmentFilterType = value
    }
}
// SetEvaluationResult sets the evaluationResult property value. Represents the evalaution result of the filter. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
func (m *AssignmentFilterTypeAndEvaluationResult) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    if m != nil {
        m.evaluationResult = value
    }
}
