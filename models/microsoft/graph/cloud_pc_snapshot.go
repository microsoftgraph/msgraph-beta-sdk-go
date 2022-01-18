package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcSnapshot 
type CloudPcSnapshot struct {
    Entity
    // 
    cloudPcId *string;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastRestoredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    status *CloudPcSnapshotStatus;
}
// NewCloudPcSnapshot instantiates a new cloudPcSnapshot and sets the default values.
func NewCloudPcSnapshot()(*CloudPcSnapshot) {
    m := &CloudPcSnapshot{
        Entity: *NewEntity(),
    }
    return m
}
// GetCloudPcId gets the cloudPcId property value. 
func (m *CloudPcSnapshot) GetCloudPcId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *CloudPcSnapshot) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetLastRestoredDateTime gets the lastRestoredDateTime property value. 
func (m *CloudPcSnapshot) GetLastRestoredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRestoredDateTime
    }
}
// GetStatus gets the status property value. 
func (m *CloudPcSnapshot) GetStatus()(*CloudPcSnapshotStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSnapshot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcId(val)
        }
        return nil
    }
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
    res["lastRestoredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRestoredDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcSnapshotStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcSnapshotStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *CloudPcSnapshot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcSnapshot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRestoredDateTime", m.GetLastRestoredDateTime())
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
    return nil
}
// SetCloudPcId sets the cloudPcId property value. 
func (m *CloudPcSnapshot) SetCloudPcId(value *string)() {
    if m != nil {
        m.cloudPcId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *CloudPcSnapshot) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetLastRestoredDateTime sets the lastRestoredDateTime property value. 
func (m *CloudPcSnapshot) SetLastRestoredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRestoredDateTime = value
    }
}
// SetStatus sets the status property value. 
func (m *CloudPcSnapshot) SetStatus(value *CloudPcSnapshotStatus)() {
    if m != nil {
        m.status = value
    }
}
