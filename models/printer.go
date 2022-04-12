package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Printer 
type Printer struct {
    PrinterBase
    // The acceptingJobs property
    acceptingJobs *bool
    // The connectors that are associated with the printer.
    connectors []PrintConnectorable
    // True if the printer has a physical device for printing. Read-only.
    hasPhysicalDevice *bool
    // True if the printer is shared; false otherwise. Read-only.
    isShared *bool
    // The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DateTimeOffset when the printer was registered. Read-only.
    registeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The share property
    share PrinterShareable
    // The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable.
    shares []PrinterShareable
    // A list of task triggers that are associated with the printer.
    taskTriggers []PrintTaskTriggerable
}
// NewPrinter instantiates a new printer and sets the default values.
func NewPrinter()(*Printer) {
    m := &Printer{
        PrinterBase: *NewPrinterBase(),
    }
    return m
}
// CreatePrinterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinter(), nil
}
// GetAcceptingJobs gets the acceptingJobs property value. The acceptingJobs property
func (m *Printer) GetAcceptingJobs()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acceptingJobs
    }
}
// GetConnectors gets the connectors property value. The connectors that are associated with the printer.
func (m *Printer) GetConnectors()([]PrintConnectorable) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Printer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrinterBase.GetFieldDeserializers()
    res["acceptingJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptingJobs(val)
        }
        return nil
    }
    res["connectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(PrintConnectorable)
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["hasPhysicalDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPhysicalDevice(val)
        }
        return nil
    }
    res["isShared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["registeredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredDateTime(val)
        }
        return nil
    }
    res["share"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShare(val.(PrinterShareable))
        }
        return nil
    }
    res["shares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrinterShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterShareable, len(val))
            for i, v := range val {
                res[i] = v.(PrinterShareable)
            }
            m.SetShares(res)
        }
        return nil
    }
    res["taskTriggers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintTaskTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskTriggerable, len(val))
            for i, v := range val {
                res[i] = v.(PrintTaskTriggerable)
            }
            m.SetTaskTriggers(res)
        }
        return nil
    }
    return res
}
// GetHasPhysicalDevice gets the hasPhysicalDevice property value. True if the printer has a physical device for printing. Read-only.
func (m *Printer) GetHasPhysicalDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasPhysicalDevice
    }
}
// GetIsShared gets the isShared property value. True if the printer is shared; false otherwise. Read-only.
func (m *Printer) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
func (m *Printer) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// GetRegisteredDateTime gets the registeredDateTime property value. The DateTimeOffset when the printer was registered. Read-only.
func (m *Printer) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registeredDateTime
    }
}
// GetShare gets the share property value. The share property
func (m *Printer) GetShare()(PrinterShareable) {
    if m == nil {
        return nil
    } else {
        return m.share
    }
}
// GetShares gets the shares property value. The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable.
func (m *Printer) GetShares()([]PrinterShareable) {
    if m == nil {
        return nil
    } else {
        return m.shares
    }
}
// GetTaskTriggers gets the taskTriggers property value. A list of task triggers that are associated with the printer.
func (m *Printer) GetTaskTriggers()([]PrintTaskTriggerable) {
    if m == nil {
        return nil
    } else {
        return m.taskTriggers
    }
}
// Serialize serializes information the current object
func (m *Printer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrinterBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("acceptingJobs", m.GetAcceptingJobs())
        if err != nil {
            return err
        }
    }
    if m.GetConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("connectors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasPhysicalDevice", m.GetHasPhysicalDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registeredDateTime", m.GetRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("share", m.GetShare())
        if err != nil {
            return err
        }
    }
    if m.GetShares() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShares()))
        for i, v := range m.GetShares() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("shares", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskTriggers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskTriggers()))
        for i, v := range m.GetTaskTriggers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskTriggers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptingJobs sets the acceptingJobs property value. The acceptingJobs property
func (m *Printer) SetAcceptingJobs(value *bool)() {
    if m != nil {
        m.acceptingJobs = value
    }
}
// SetConnectors sets the connectors property value. The connectors that are associated with the printer.
func (m *Printer) SetConnectors(value []PrintConnectorable)() {
    if m != nil {
        m.connectors = value
    }
}
// SetHasPhysicalDevice sets the hasPhysicalDevice property value. True if the printer has a physical device for printing. Read-only.
func (m *Printer) SetHasPhysicalDevice(value *bool)() {
    if m != nil {
        m.hasPhysicalDevice = value
    }
}
// SetIsShared sets the isShared property value. True if the printer is shared; false otherwise. Read-only.
func (m *Printer) SetIsShared(value *bool)() {
    if m != nil {
        m.isShared = value
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
func (m *Printer) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSeenDateTime = value
    }
}
// SetRegisteredDateTime sets the registeredDateTime property value. The DateTimeOffset when the printer was registered. Read-only.
func (m *Printer) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.registeredDateTime = value
    }
}
// SetShare sets the share property value. The share property
func (m *Printer) SetShare(value PrinterShareable)() {
    if m != nil {
        m.share = value
    }
}
// SetShares sets the shares property value. The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable.
func (m *Printer) SetShares(value []PrinterShareable)() {
    if m != nil {
        m.shares = value
    }
}
// SetTaskTriggers sets the taskTriggers property value. A list of task triggers that are associated with the printer.
func (m *Printer) SetTaskTriggers(value []PrintTaskTriggerable)() {
    if m != nil {
        m.taskTriggers = value
    }
}
