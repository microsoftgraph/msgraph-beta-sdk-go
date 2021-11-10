package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OfficeUserCheckinSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total failed user check ins for the last 3 months.
    failedUserCount *int32;
    // Total successful user check ins for the last 3 months.
    succeededUserCount *int32;
}
// Instantiates a new officeUserCheckinSummary and sets the default values.
func NewOfficeUserCheckinSummary()(*OfficeUserCheckinSummary) {
    m := &OfficeUserCheckinSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeUserCheckinSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the failedUserCount property value. Total failed user check ins for the last 3 months.
func (m *OfficeUserCheckinSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
// Gets the succeededUserCount property value. Total successful user check ins for the last 3 months.
func (m *OfficeUserCheckinSummary) GetSucceededUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.succeededUserCount
    }
}
// The deserialization information for the current model
func (m *OfficeUserCheckinSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUserCount(val)
        }
        return nil
    }
    res["succeededUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSucceededUserCount(val)
        }
        return nil
    }
    return res
}
func (m *OfficeUserCheckinSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OfficeUserCheckinSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("succeededUserCount", m.GetSucceededUserCount())
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
func (m *OfficeUserCheckinSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the failedUserCount property value. Total failed user check ins for the last 3 months.
// Parameters:
//  - value : Value to set for the failedUserCount property.
func (m *OfficeUserCheckinSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
// Sets the succeededUserCount property value. Total successful user check ins for the last 3 months.
// Parameters:
//  - value : Value to set for the succeededUserCount property.
func (m *OfficeUserCheckinSummary) SetSucceededUserCount(value *int32)() {
    m.succeededUserCount = value
}
