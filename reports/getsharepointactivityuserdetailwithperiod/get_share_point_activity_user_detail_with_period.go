package getsharepointactivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getSharePointActivityUserDetailWithPeriod 
type GetSharePointActivityUserDetailWithPeriod struct {
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
    // 
    visitedPageCount *int64;
}
// NewGetSharePointActivityUserDetailWithPeriod instantiates a new getSharePointActivityUserDetailWithPeriod and sets the default values.
func NewGetSharePointActivityUserDetailWithPeriod()(*GetSharePointActivityUserDetailWithPeriod) {
    m := &GetSharePointActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAssignedProducts gets the assignedProducts property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// GetDeletedDate gets the deletedDate property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharedExternallyFileCount gets the sharedExternallyFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetSharedExternallyFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternallyFileCount
    }
}
// GetSharedInternallyFileCount gets the sharedInternallyFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetSharedInternallyFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternallyFileCount
    }
}
// GetSyncedFileCount gets the syncedFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetSyncedFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.syncedFileCount
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetViewedOrEditedFileCount gets the viewedOrEditedFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetViewedOrEditedFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEditedFileCount
    }
}
// GetVisitedPageCount gets the visitedPageCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) GetVisitedPageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPageCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSharePointActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["visitedPageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisitedPageCount(val)
        }
        return nil
    }
    return res
}
func (m *GetSharePointActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSharePointActivityUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err = writer.WriteInt64Value("visitedPageCount", m.GetVisitedPageCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedProducts sets the assignedProducts property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// SetDeletedDate sets the deletedDate property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetSharedExternallyFileCount sets the sharedExternallyFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetSharedExternallyFileCount(value *int64)() {
    m.sharedExternallyFileCount = value
}
// SetSharedInternallyFileCount sets the sharedInternallyFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetSharedInternallyFileCount(value *int64)() {
    m.sharedInternallyFileCount = value
}
// SetSyncedFileCount sets the syncedFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetSyncedFileCount(value *int64)() {
    m.syncedFileCount = value
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetViewedOrEditedFileCount sets the viewedOrEditedFileCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetViewedOrEditedFileCount(value *int64)() {
    m.viewedOrEditedFileCount = value
}
// SetVisitedPageCount sets the visitedPageCount property value. 
func (m *GetSharePointActivityUserDetailWithPeriod) SetVisitedPageCount(value *int64)() {
    m.visitedPageCount = value
}
