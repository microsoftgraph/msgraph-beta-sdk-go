package getonedriveactivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOneDriveActivityUserDetailWithPeriod 
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
// NewGetOneDriveActivityUserDetailWithPeriod instantiates a new getOneDriveActivityUserDetailWithPeriod and sets the default values.
func NewGetOneDriveActivityUserDetailWithPeriod()(*GetOneDriveActivityUserDetailWithPeriod) {
    m := &GetOneDriveActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAssignedProducts gets the assignedProducts property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// GetDeletedDate gets the deletedDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharedExternallyFileCount gets the sharedExternallyFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetSharedExternallyFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternallyFileCount
    }
}
// GetSharedInternallyFileCount gets the sharedInternallyFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetSharedInternallyFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternallyFileCount
    }
}
// GetSyncedFileCount gets the syncedFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetSyncedFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.syncedFileCount
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetViewedOrEditedFileCount gets the viewedOrEditedFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) GetViewedOrEditedFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEditedFileCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOneDriveActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignedProducts(res)
        }
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDate(val)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharedExternallyFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedExternallyFileCount(val)
        }
        return nil
    }
    res["sharedInternallyFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedInternallyFileCount(val)
        }
        return nil
    }
    res["syncedFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncedFileCount(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["viewedOrEditedFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewedOrEditedFileCount(val)
        }
        return nil
    }
    return res
}
func (m *GetOneDriveActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAssignedProducts sets the assignedProducts property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    if m != nil {
        m.assignedProducts = value
    }
}
// SetDeletedDate sets the deletedDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetDeletedDate(value *string)() {
    if m != nil {
        m.deletedDate = value
    }
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSharedExternallyFileCount sets the sharedExternallyFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetSharedExternallyFileCount(value *int64)() {
    if m != nil {
        m.sharedExternallyFileCount = value
    }
}
// SetSharedInternallyFileCount sets the sharedInternallyFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetSharedInternallyFileCount(value *int64)() {
    if m != nil {
        m.sharedInternallyFileCount = value
    }
}
// SetSyncedFileCount sets the syncedFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetSyncedFileCount(value *int64)() {
    if m != nil {
        m.syncedFileCount = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetViewedOrEditedFileCount sets the viewedOrEditedFileCount property value. 
func (m *GetOneDriveActivityUserDetailWithPeriod) SetViewedOrEditedFileCount(value *int64)() {
    if m != nil {
        m.viewedOrEditedFileCount = value
    }
}
