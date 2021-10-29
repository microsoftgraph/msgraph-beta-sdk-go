package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintJobStatus struct {
    // 
    acquiredByPrinter *bool;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A human-readable description of the print job's current processing state. Read-only.
    description *string;
    // Additional details for print job state. Valid values are described in the following table. Read-only.
    details []PrintJobStateDetail;
    // True if the job was acknowledged by a printer; false otherwise. Read-only.
    isAcquiredByPrinter *bool;
    // 
    processingState *PrintJobProcessingState;
    // 
    processingStateDescription *string;
    // The print job's current processing state. Valid values are described in the following table. Read-only.
    state *PrintJobProcessingState;
}
// Instantiates a new printJobStatus and sets the default values.
func NewPrintJobStatus()(*PrintJobStatus) {
    m := &PrintJobStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the acquiredByPrinter property value. 
func (m *PrintJobStatus) GetAcquiredByPrinter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acquiredByPrinter
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintJobStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. A human-readable description of the print job's current processing state. Read-only.
func (m *PrintJobStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the details property value. Additional details for print job state. Valid values are described in the following table. Read-only.
func (m *PrintJobStatus) GetDetails()([]PrintJobStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// Gets the isAcquiredByPrinter property value. True if the job was acknowledged by a printer; false otherwise. Read-only.
func (m *PrintJobStatus) GetIsAcquiredByPrinter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAcquiredByPrinter
    }
}
// Gets the processingState property value. 
func (m *PrintJobStatus) GetProcessingState()(*PrintJobProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.processingState
    }
}
// Gets the processingStateDescription property value. 
func (m *PrintJobStatus) GetProcessingStateDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processingStateDescription
    }
}
// Gets the state property value. The print job's current processing state. Valid values are described in the following table. Read-only.
func (m *PrintJobStatus) GetState()(*PrintJobProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *PrintJobStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acquiredByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAcquiredByPrinter(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintJobStateDetail)
        if err != nil {
            return err
        }
        res := make([]PrintJobStateDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintJobStateDetail))
        }
        m.SetDetails(res)
        return nil
    }
    res["isAcquiredByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAcquiredByPrinter(val)
        return nil
    }
    res["processingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintJobProcessingState)
        if err != nil {
            return err
        }
        cast := val.(PrintJobProcessingState)
        m.SetProcessingState(&cast)
        return nil
    }
    res["processingStateDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProcessingStateDescription(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintJobProcessingState)
        if err != nil {
            return err
        }
        cast := val.(PrintJobProcessingState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *PrintJobStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintJobStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acquiredByPrinter", m.GetAcquiredByPrinter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("details", SerializePrintJobStateDetail(m.GetDetails()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAcquiredByPrinter", m.GetIsAcquiredByPrinter())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingState() != nil {
        cast := m.GetProcessingState().String()
        err := writer.WriteStringValue("processingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processingStateDescription", m.GetProcessingStateDescription())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
// Sets the acquiredByPrinter property value. 
// Parameters:
//  - value : Value to set for the acquiredByPrinter property.
func (m *PrintJobStatus) SetAcquiredByPrinter(value *bool)() {
    m.acquiredByPrinter = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PrintJobStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. A human-readable description of the print job's current processing state. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *PrintJobStatus) SetDescription(value *string)() {
    m.description = value
}
// Sets the details property value. Additional details for print job state. Valid values are described in the following table. Read-only.
// Parameters:
//  - value : Value to set for the details property.
func (m *PrintJobStatus) SetDetails(value []PrintJobStateDetail)() {
    m.details = value
}
// Sets the isAcquiredByPrinter property value. True if the job was acknowledged by a printer; false otherwise. Read-only.
// Parameters:
//  - value : Value to set for the isAcquiredByPrinter property.
func (m *PrintJobStatus) SetIsAcquiredByPrinter(value *bool)() {
    m.isAcquiredByPrinter = value
}
// Sets the processingState property value. 
// Parameters:
//  - value : Value to set for the processingState property.
func (m *PrintJobStatus) SetProcessingState(value *PrintJobProcessingState)() {
    m.processingState = value
}
// Sets the processingStateDescription property value. 
// Parameters:
//  - value : Value to set for the processingStateDescription property.
func (m *PrintJobStatus) SetProcessingStateDescription(value *string)() {
    m.processingStateDescription = value
}
// Sets the state property value. The print job's current processing state. Valid values are described in the following table. Read-only.
// Parameters:
//  - value : Value to set for the state property.
func (m *PrintJobStatus) SetState(value *PrintJobProcessingState)() {
    m.state = value
}
