package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TenantTag struct {
    Entity
    // The identifier for the account that created the tenant tag. Required. Read-only.
    createdByUserId *string;
    // The date and time when the tenant tag was created. Required. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time when the tenant tag was deleted. Required. Read-only.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description for the tenant tag. Optional. Read-only.
    description *string;
    // The display name for the tenant tag. Required. Read-only.
    displayName *string;
    // The identifier for the account that lasted on the tenant tag. Optional. Read-only.
    lastActionByUserId *string;
    // The date and time the last action was performed against the tenant tag. Optional. Read-only.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of managed tenants associated with the tenant tag. Optional.
    tenants []TenantInfo;
}
// Instantiates a new tenantTag and sets the default values.
func NewTenantTag()(*TenantTag) {
    m := &TenantTag{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdByUserId property value. The identifier for the account that created the tenant tag. Required. Read-only.
func (m *TenantTag) GetCreatedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByUserId
    }
}
// Gets the createdDateTime property value. The date and time when the tenant tag was created. Required. Read-only.
func (m *TenantTag) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deletedDateTime property value. The date and time when the tenant tag was deleted. Required. Read-only.
func (m *TenantTag) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deletedDateTime
    }
}
// Gets the description property value. The description for the tenant tag. Optional. Read-only.
func (m *TenantTag) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the tenant tag. Required. Read-only.
func (m *TenantTag) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastActionByUserId property value. The identifier for the account that lasted on the tenant tag. Optional. Read-only.
func (m *TenantTag) GetLastActionByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActionByUserId
    }
}
// Gets the lastActionDateTime property value. The date and time the last action was performed against the tenant tag. Optional. Read-only.
func (m *TenantTag) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// Gets the tenants property value. The collection of managed tenants associated with the tenant tag. Optional.
func (m *TenantTag) GetTenants()([]TenantInfo) {
    if m == nil {
        return nil
    } else {
        return m.tenants
    }
}
// The deserialization information for the current model
func (m *TenantTag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatedByUserId(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deletedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeletedDateTime(val)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastActionByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActionByUserId(val)
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
    res["tenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantInfo() })
        if err != nil {
            return err
        }
        res := make([]TenantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*TenantInfo))
        }
        m.SetTenants(res)
        return nil
    }
    return res
}
func (m *TenantTag) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TenantTag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
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
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdByUserId property value. The identifier for the account that created the tenant tag. Required. Read-only.
// Parameters:
//  - value : Value to set for the createdByUserId property.
func (m *TenantTag) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// Sets the createdDateTime property value. The date and time when the tenant tag was created. Required. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *TenantTag) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deletedDateTime property value. The date and time when the tenant tag was deleted. Required. Read-only.
// Parameters:
//  - value : Value to set for the deletedDateTime property.
func (m *TenantTag) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// Sets the description property value. The description for the tenant tag. Optional. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *TenantTag) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the tenant tag. Required. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TenantTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastActionByUserId property value. The identifier for the account that lasted on the tenant tag. Optional. Read-only.
// Parameters:
//  - value : Value to set for the lastActionByUserId property.
func (m *TenantTag) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// Sets the lastActionDateTime property value. The date and time the last action was performed against the tenant tag. Optional. Read-only.
// Parameters:
//  - value : Value to set for the lastActionDateTime property.
func (m *TenantTag) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// Sets the tenants property value. The collection of managed tenants associated with the tenant tag. Optional.
// Parameters:
//  - value : Value to set for the tenants property.
func (m *TenantTag) SetTenants(value []TenantInfo)() {
    m.tenants = value
}
