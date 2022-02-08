package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterEvaluationSummary 
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
// NewAssignmentFilterEvaluationSummary instantiates a new assignmentFilterEvaluationSummary and sets the default values.
func NewAssignmentFilterEvaluationSummary()(*AssignmentFilterEvaluationSummary) {
    m := &AssignmentFilterEvaluationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluationSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentFilterDisplayName gets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterDisplayName
    }
}
// GetAssignmentFilterId gets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterId
    }
}
// GetAssignmentFilterLastModifiedDateTime gets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterLastModifiedDateTime
    }
}
// GetAssignmentFilterPlatform gets the assignmentFilterPlatform property value. The platform for which this assignment filter is created. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterPlatform
    }
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Indicate filter type either include or exclude. Possible values are: none, include, exclude.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterType
    }
}
// GetAssignmentFilterTypeAndEvaluationResults gets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterTypeAndEvaluationResults
    }
}
// GetEvaluationDateTime gets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.evaluationDateTime
    }
}
// GetEvaluationResult gets the evaluationResult property value. Assignment filter evaluation result. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.evaluationResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterEvaluationSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterDisplayName(val)
        }
        return nil
    }
    res["assignmentFilterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterId(val)
        }
        return nil
    }
    res["assignmentFilterLastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterLastModifiedDateTime(val)
        }
        return nil
    }
    res["assignmentFilterPlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
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
    res["assignmentFilterTypeAndEvaluationResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentFilterTypeAndEvaluationResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterTypeAndEvaluationResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignmentFilterTypeAndEvaluationResult))
            }
            m.SetAssignmentFilterTypeAndEvaluationResults(res)
        }
        return nil
    }
    res["evaluationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationDateTime(val)
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
func (m *AssignmentFilterEvaluationSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetAssignmentFilterPlatform()).String()
        err := writer.WriteStringValue("assignmentFilterPlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterTypeAndEvaluationResults() != nil {
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
func (m *AssignmentFilterEvaluationSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentFilterDisplayName sets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterDisplayName(value *string)() {
    if m != nil {
        m.assignmentFilterDisplayName = value
    }
}
// SetAssignmentFilterId sets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterId(value *string)() {
    if m != nil {
        m.assignmentFilterId = value
    }
}
// SetAssignmentFilterLastModifiedDateTime sets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.assignmentFilterLastModifiedDateTime = value
    }
}
// SetAssignmentFilterPlatform sets the assignmentFilterPlatform property value. The platform for which this assignment filter is created. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, unknown.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterPlatform(value *DevicePlatformType)() {
    if m != nil {
        m.assignmentFilterPlatform = value
    }
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Indicate filter type either include or exclude. Possible values are: none, include, exclude.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    if m != nil {
        m.assignmentFilterType = value
    }
}
// SetAssignmentFilterTypeAndEvaluationResults sets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResult)() {
    if m != nil {
        m.assignmentFilterTypeAndEvaluationResults = value
    }
}
// SetEvaluationDateTime sets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.evaluationDateTime = value
    }
}
// SetEvaluationResult sets the evaluationResult property value. Assignment filter evaluation result. Possible values are: unknown, match, notMatch, inconclusive, failure, notEvaluated.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    if m != nil {
        m.evaluationResult = value
    }
}
