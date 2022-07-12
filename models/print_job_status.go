package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJobStatus 
type PrintJobStatus struct {
    // The acquiredByPrinter property
    acquiredByPrinter *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A human-readable description of the print job's current processing state. Read-only.
    description *string
    // Additional details for print job state. Valid values are described in the following table. Read-only.
    details []string
    // True if the job was acknowledged by a printer; false otherwise. Read-only.
    isAcquiredByPrinter *bool
    // The processingState property
    processingState *PrintJobProcessingState
    // The processingStateDescription property
    processingStateDescription *string
    // The state property
    state *PrintJobProcessingState
}
// NewPrintJobStatus instantiates a new printJobStatus and sets the default values.
func NewPrintJobStatus()(*PrintJobStatus) {
    m := &PrintJobStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintJobStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintJobStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintJobStatus(), nil
}
// GetAcquiredByPrinter gets the acquiredByPrinter property value. The acquiredByPrinter property
func (m *PrintJobStatus) GetAcquiredByPrinter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acquiredByPrinter
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintJobStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. A human-readable description of the print job's current processing state. Read-only.
func (m *PrintJobStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetails gets the details property value. Additional details for print job state. Valid values are described in the following table. Read-only.
func (m *PrintJobStatus) GetDetails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintJobStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acquiredByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcquiredByPrinter(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["isAcquiredByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAcquiredByPrinter(val)
        }
        return nil
    }
    res["processingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintJobProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingState(val.(*PrintJobProcessingState))
        }
        return nil
    }
    res["processingStateDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingStateDescription(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintJobProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PrintJobProcessingState))
        }
        return nil
    }
    return res
}
// GetIsAcquiredByPrinter gets the isAcquiredByPrinter property value. True if the job was acknowledged by a printer; false otherwise. Read-only.
func (m *PrintJobStatus) GetIsAcquiredByPrinter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAcquiredByPrinter
    }
}
// GetProcessingState gets the processingState property value. The processingState property
func (m *PrintJobStatus) GetProcessingState()(*PrintJobProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.processingState
    }
}
// GetProcessingStateDescription gets the processingStateDescription property value. The processingStateDescription property
func (m *PrintJobStatus) GetProcessingStateDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processingStateDescription
    }
}
// GetState gets the state property value. The state property
func (m *PrintJobStatus) GetState()(*PrintJobProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *PrintJobStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDetails() != nil {
        err := writer.WriteCollectionOfStringValues("details", m.GetDetails())
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
        cast := (*m.GetProcessingState()).String()
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
        cast := (*m.GetState()).String()
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
// SetAcquiredByPrinter sets the acquiredByPrinter property value. The acquiredByPrinter property
func (m *PrintJobStatus) SetAcquiredByPrinter(value *bool)() {
    if m != nil {
        m.acquiredByPrinter = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintJobStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. A human-readable description of the print job's current processing state. Read-only.
func (m *PrintJobStatus) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetails sets the details property value. Additional details for print job state. Valid values are described in the following table. Read-only.
func (m *PrintJobStatus) SetDetails(value []string)() {
    if m != nil {
        m.details = value
    }
}
// SetIsAcquiredByPrinter sets the isAcquiredByPrinter property value. True if the job was acknowledged by a printer; false otherwise. Read-only.
func (m *PrintJobStatus) SetIsAcquiredByPrinter(value *bool)() {
    if m != nil {
        m.isAcquiredByPrinter = value
    }
}
// SetProcessingState sets the processingState property value. The processingState property
func (m *PrintJobStatus) SetProcessingState(value *PrintJobProcessingState)() {
    if m != nil {
        m.processingState = value
    }
}
// SetProcessingStateDescription sets the processingStateDescription property value. The processingStateDescription property
func (m *PrintJobStatus) SetProcessingStateDescription(value *string)() {
    if m != nil {
        m.processingStateDescription = value
    }
}
// SetState sets the state property value. The state property
func (m *PrintJobStatus) SetState(value *PrintJobProcessingState)() {
    if m != nil {
        m.state = value
    }
}
