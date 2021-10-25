package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentFilterEvaluationSummary struct {
    additionalData map[string]interface{};
    assignmentFilterDisplayName *string;
    assignmentFilterId *string;
    assignmentFilterLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    assignmentFilterPlatform *DevicePlatformType;
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType;
    assignmentFilterTypeAndEvaluationResults []AssignmentFilterTypeAndEvaluationResult;
    evaluationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    evaluationResult *AssignmentFilterEvaluationResult;
}
func NewAssignmentFilterEvaluationSummary()(*AssignmentFilterEvaluationSummary) {
    m := &AssignmentFilterEvaluationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignmentFilterEvaluationSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterDisplayName
    }
}
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterId
    }
}
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterLastModifiedDateTime
    }
}
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterPlatform
    }
}
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterType
    }
}
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterTypeAndEvaluationResults
    }
}
func (m *AssignmentFilterEvaluationSummary) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.evaluationDateTime
    }
}
func (m *AssignmentFilterEvaluationSummary) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.evaluationResult
    }
}
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
func (m *AssignmentFilterEvaluationSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterDisplayName(value *string)() {
    m.assignmentFilterDisplayName = value
}
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterId(value *string)() {
    m.assignmentFilterId = value
}
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignmentFilterLastModifiedDateTime = value
}
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterPlatform(value *DevicePlatformType)() {
    m.assignmentFilterPlatform = value
}
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResult)() {
    m.assignmentFilterTypeAndEvaluationResults = value
}
func (m *AssignmentFilterEvaluationSummary) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.evaluationDateTime = value
}
func (m *AssignmentFilterEvaluationSummary) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    m.evaluationResult = value
}
