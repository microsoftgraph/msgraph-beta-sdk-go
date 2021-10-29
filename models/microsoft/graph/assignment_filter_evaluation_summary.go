package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssignmentFilterEvaluationSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The admin defined name for assignment filter.
    assignmentFilterDisplayName *string;
    // Unique identifier for the assignment filter object
    assignmentFilterId *string;
    // The time the assignment filter was last modified.
    assignmentFilterLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The platform for which this assignment filter is created. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
    assignmentFilterPlatform *DevicePlatformType;
    // Indicate filter type either include or exclude. Possible values are: none, include, exclude.
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType;
    // A collection of filter types and their corresponding evaluation results.
    assignmentFilterTypeAndEvaluationResults []AssignmentFilterTypeAndEvaluationResult;
    // The time assignment filter was evaluated.
    evaluationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Assignment filter evaluation result. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
    evaluationResult *AssignmentFilterEvaluationResult;
}
// Instantiates a new assignmentFilterEvaluationSummary and sets the default values.
func NewAssignmentFilterEvaluationSummary()(*AssignmentFilterEvaluationSummary) {
    m := &AssignmentFilterEvaluationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluationSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterDisplayName
    }
}
// Gets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterId
    }
}
// Gets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterLastModifiedDateTime
    }
}
// Gets the assignmentFilterPlatform property value. The platform for which this assignment filter is created. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterPlatform
    }
}
// Gets the assignmentFilterType property value. Indicate filter type either include or exclude. Possible values are: none, include, exclude.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterType
    }
}
// Gets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterTypeAndEvaluationResults
    }
}
// Gets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.evaluationDateTime
    }
}
// Gets the evaluationResult property value. Assignment filter evaluation result. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.evaluationResult
    }
}
// The deserialization information for the current model
func (m *AssignmentFilterEvaluationSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentFilterDisplayName(val)
        return nil
    }
    res["assignmentFilterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentFilterId(val)
        return nil
    }
    res["assignmentFilterLastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAssignmentFilterLastModifiedDateTime(val)
        return nil
    }
    res["assignmentFilterPlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        cast := val.(DevicePlatformType)
        m.SetAssignmentFilterPlatform(&cast)
        return nil
    }
    res["assignmentFilterType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        cast := val.(DeviceAndAppManagementAssignmentFilterType)
        m.SetAssignmentFilterType(&cast)
        return nil
    }
    res["assignmentFilterTypeAndEvaluationResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentFilterTypeAndEvaluationResult() })
        if err != nil {
            return err
        }
        res := make([]AssignmentFilterTypeAndEvaluationResult, len(val))
        for i, v := range val {
            res[i] = *(v.(*AssignmentFilterTypeAndEvaluationResult))
        }
        m.SetAssignmentFilterTypeAndEvaluationResults(res)
        return nil
    }
    res["evaluationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEvaluationDateTime(val)
        return nil
    }
    res["evaluationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterEvaluationResult)
        if err != nil {
            return err
        }
        cast := val.(AssignmentFilterEvaluationResult)
        m.SetEvaluationResult(&cast)
        return nil
    }
    return res
}
func (m *AssignmentFilterEvaluationSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssignmentFilterEvaluationSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignmentFilterDisplayName", m.GetAssignmentFilterDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignmentFilterId", m.GetAssignmentFilterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("assignmentFilterLastModifiedDateTime", m.GetAssignmentFilterLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterPlatform() != nil {
        cast := m.GetAssignmentFilterPlatform().String()
        err := writer.WriteStringValue("assignmentFilterPlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterType() != nil {
        cast := m.GetAssignmentFilterType().String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentFilterTypeAndEvaluationResults()))
        for i, v := range m.GetAssignmentFilterTypeAndEvaluationResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("assignmentFilterTypeAndEvaluationResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("evaluationDateTime", m.GetEvaluationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := m.GetEvaluationResult().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AssignmentFilterEvaluationSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
// Parameters:
//  - value : Value to set for the assignmentFilterDisplayName property.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterDisplayName(value *string)() {
    m.assignmentFilterDisplayName = value
}
// Sets the assignmentFilterId property value. Unique identifier for the assignment filter object
// Parameters:
//  - value : Value to set for the assignmentFilterId property.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterId(value *string)() {
    m.assignmentFilterId = value
}
// Sets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
// Parameters:
//  - value : Value to set for the assignmentFilterLastModifiedDateTime property.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignmentFilterLastModifiedDateTime = value
}
// Sets the assignmentFilterPlatform property value. The platform for which this assignment filter is created. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
// Parameters:
//  - value : Value to set for the assignmentFilterPlatform property.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterPlatform(value *DevicePlatformType)() {
    m.assignmentFilterPlatform = value
}
// Sets the assignmentFilterType property value. Indicate filter type either include or exclude. Possible values are: none, include, exclude.
// Parameters:
//  - value : Value to set for the assignmentFilterType property.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
// Sets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
// Parameters:
//  - value : Value to set for the assignmentFilterTypeAndEvaluationResults property.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResult)() {
    m.assignmentFilterTypeAndEvaluationResults = value
}
// Sets the evaluationDateTime property value. The time assignment filter was evaluated.
// Parameters:
//  - value : Value to set for the evaluationDateTime property.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.evaluationDateTime = value
}
// Sets the evaluationResult property value. Assignment filter evaluation result. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
// Parameters:
//  - value : Value to set for the evaluationResult property.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    m.evaluationResult = value
}
