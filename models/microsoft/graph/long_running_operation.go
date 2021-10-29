package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new longRunningOperation and sets the default values.
func NewLongRunningOperation()(*LongRunningOperation) {
    m := &LongRunningOperation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. 
func (m *LongRunningOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the lastActionDateTime property value. 
func (m *LongRunningOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// Gets the resourceLocation property value. 
func (m *LongRunningOperation) GetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLocation
    }
}
// Gets the status property value. 
func (m *LongRunningOperation) GetStatus()(*LongRunningOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the statusDetail property value. 
func (m *LongRunningOperation) GetStatusDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statusDetail
    }
}
// The deserialization information for the current model
func (m *LongRunningOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastActionDateTime(val)
        return nil
    }
    res["resourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceLocation(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLongRunningOperationStatus)
        if err != nil {
            return err
        }
        cast := val.(LongRunningOperationStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["statusDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatusDetail(val)
        return nil
    }
    return res
}
func (m *LongRunningOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetStatus().String()
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
// Sets the createdDateTime property value. 
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *LongRunningOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the lastActionDateTime property value. 
// Parameters:
//  - value : Value to set for the lastActionDateTime property.
func (m *LongRunningOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// Sets the resourceLocation property value. 
// Parameters:
//  - value : Value to set for the resourceLocation property.
func (m *LongRunningOperation) SetResourceLocation(value *string)() {
    m.resourceLocation = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *LongRunningOperation) SetStatus(value *LongRunningOperationStatus)() {
    m.status = value
}
// Sets the statusDetail property value. 
// Parameters:
//  - value : Value to set for the statusDetail property.
func (m *LongRunningOperation) SetStatusDetail(value *string)() {
    m.statusDetail = value
}
