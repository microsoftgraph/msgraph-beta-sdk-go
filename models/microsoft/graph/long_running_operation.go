package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LongRunningOperation 
type LongRunningOperation struct {
    Entity
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    resourceLocation *string;
    // 
    status *LongRunningOperationStatus;
    // 
    statusDetail *string;
}
// NewLongRunningOperation instantiates a new longRunningOperation and sets the default values.
func NewLongRunningOperation()(*LongRunningOperation) {
    m := &LongRunningOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *LongRunningOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetLastActionDateTime gets the lastActionDateTime property value. 
func (m *LongRunningOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetResourceLocation gets the resourceLocation property value. 
func (m *LongRunningOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
// GetStatus gets the status property value. 
func (m *LongRunningOperation) GetStatus()(*LongRunningOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetStatusDetail gets the statusDetail property value. 
func (m *LongRunningOperation) GetStatusDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statusDetail
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LongRunningOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["resourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceLocation(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLongRunningOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*LongRunningOperationStatus))
        }
        return nil
    }
    res["statusDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetail(val)
        }
        return nil
    }
    return res
}
func (m *LongRunningOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LongRunningOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceLocation", m.GetResourceLocation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("statusDetail", m.GetStatusDetail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *LongRunningOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. 
func (m *LongRunningOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetResourceLocation sets the resourceLocation property value. 
func (m *LongRunningOperation) SetResourceLocation(value *string)() {
    if m != nil {
        m.resourceLocation = value
    }
}
// SetStatus sets the status property value. 
func (m *LongRunningOperation) SetStatus(value *LongRunningOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetStatusDetail sets the statusDetail property value. 
func (m *LongRunningOperation) SetStatusDetail(value *string)() {
    if m != nil {
        m.statusDetail = value
    }
}
