package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// DataSourceContainer 
type DataSourceContainer struct {
    Entity
    // Created date and time of the dataSourceContainer entity.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Display name of the dataSourceContainer entity.
    displayName *string;
    // 
    holdStatus *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceHoldStatus;
    // 
    lastIndexOperation *CaseIndexOperation;
    // Last modified date and time of the dataSourceContainer.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date and time that the dataSourceContainer was released from the case.
    releasedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Latest status of the dataSourceContainer. Possible values are: Active, Released.
    status *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceContainerStatus;
}
// NewDataSourceContainer instantiates a new dataSourceContainer and sets the default values.
func NewDataSourceContainer()(*DataSourceContainer) {
    m := &DataSourceContainer{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHoldStatus gets the holdStatus property value. 
func (m *DataSourceContainer) GetHoldStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceHoldStatus) {
    if m == nil {
        return nil
    } else {
        return m.holdStatus
    }
}
// GetLastIndexOperation gets the lastIndexOperation property value. 
func (m *DataSourceContainer) GetLastIndexOperation()(*CaseIndexOperation) {
    if m == nil {
        return nil
    } else {
        return m.lastIndexOperation
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetReleasedDateTime gets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) GetReleasedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.releasedDateTime
    }
}
// GetStatus gets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) GetStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceContainerStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataSourceContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["holdStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseDataSourceHoldStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHoldStatus(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceHoldStatus))
        }
        return nil
    }
    res["lastIndexOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseIndexOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastIndexOperation(val.(*CaseIndexOperation))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["releasedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseDataSourceContainerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceContainerStatus))
        }
        return nil
    }
    return res
}
func (m *DataSourceContainer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DataSourceContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetHoldStatus() != nil {
        cast := (*m.GetHoldStatus()).String()
        err = writer.WriteStringValue("holdStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastIndexOperation", m.GetLastIndexOperation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releasedDateTime", m.GetReleasedDateTime())
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
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHoldStatus sets the holdStatus property value. 
func (m *DataSourceContainer) SetHoldStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceHoldStatus)() {
    if m != nil {
        m.holdStatus = value
    }
}
// SetLastIndexOperation sets the lastIndexOperation property value. 
func (m *DataSourceContainer) SetLastIndexOperation(value *CaseIndexOperation)() {
    if m != nil {
        m.lastIndexOperation = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetReleasedDateTime sets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) SetReleasedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.releasedDateTime = value
    }
}
// SetStatus sets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) SetStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceContainerStatus)() {
    if m != nil {
        m.status = value
    }
}
