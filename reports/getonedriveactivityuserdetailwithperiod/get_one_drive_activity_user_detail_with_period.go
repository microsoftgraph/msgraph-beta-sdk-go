package getonedriveactivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOneDriveActivityUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    assignedProducts []string;
    // 
    deletedDate *string;
    // 
    isDeleted *bool;
    // 
    lastActivityDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    sharedExternallyFileCount *int64;
    // 
    sharedInternallyFileCount *int64;
    // 
    syncedFileCount *int64;
    // 
    userPrincipalName *string;
    // 
    viewedOrEditedFileCount *int64;
}
// Instantiates a new getOneDriveActivityUserDetailWithPeriod and sets the default values.
func NewGetOneDriveActivityUserDetailWithPeriod()(*GetOneDriveActivityUserDetailWithPeriod) {
    m := &GetOneDriveActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the assignedProducts property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// Gets the deletedDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// Gets the isDeleted property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sharedExternallyFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetSharedExternallyFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternallyFileCount
    }
}
// Gets the sharedInternallyFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetSharedInternallyFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternallyFileCount
    }
}
// Gets the syncedFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetSyncedFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.syncedFileCount
    }
}
// Gets the userPrincipalName property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the viewedOrEditedFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetViewedOrEditedFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEditedFileCount
    }
}
// The deserialization information for the current model
func (m *GetOneDriveActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetAssignedProducts(res)
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["sharedExternallyFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedExternallyFileCount(val)
        return nil
    }
    res["sharedInternallyFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedInternallyFileCount(val)
        return nil
    }
    res["syncedFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSyncedFileCount(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["viewedOrEditedFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetViewedOrEditedFileCount(val)
        return nil
    }
    return res
}
func (m *GetOneDriveActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetOneDriveActivityUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("assignedProducts", m.GetAssignedProducts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharedExternallyFileCount", m.GetSharedExternallyFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharedInternallyFileCount", m.GetSharedInternallyFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("syncedFileCount", m.GetSyncedFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("viewedOrEditedFileCount", m.GetViewedOrEditedFileCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignedProducts property value. 
// Parameters:
//  - value : Value to set for the assignedProducts property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// Sets the deletedDate property value. 
// Parameters:
//  - value : Value to set for the deletedDate property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sharedExternallyFileCount property value. 
// Parameters:
//  - value : Value to set for the sharedExternallyFileCount property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetSharedExternallyFileCount(value *int64)() {
    m.sharedExternallyFileCount = value
}
// Sets the sharedInternallyFileCount property value. 
// Parameters:
//  - value : Value to set for the sharedInternallyFileCount property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetSharedInternallyFileCount(value *int64)() {
    m.sharedInternallyFileCount = value
}
// Sets the syncedFileCount property value. 
// Parameters:
//  - value : Value to set for the syncedFileCount property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetSyncedFileCount(value *int64)() {
    m.syncedFileCount = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the viewedOrEditedFileCount property value. 
// Parameters:
//  - value : Value to set for the viewedOrEditedFileCount property.
func (m *GetOneDriveActivityUserDetailWithPeriod) SetViewedOrEditedFileCount(value *int64)() {
    m.viewedOrEditedFileCount = value
}
