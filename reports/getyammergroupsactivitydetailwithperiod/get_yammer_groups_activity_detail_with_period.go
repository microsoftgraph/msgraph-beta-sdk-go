package getyammergroupsactivitydetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetYammerGroupsActivityDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    groupDisplayName *string;
    // 
    groupType *string;
    // 
    isDeleted *bool;
    // 
    lastActivityDate *string;
    // 
    likedCount *int64;
    // 
    memberCount *int64;
    // 
    networkDisplayName *string;
    // 
    office365Connected *bool;
    // 
    ownerPrincipalName *string;
    // 
    postedCount *int64;
    // 
    readCount *int64;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
}
// Instantiates a new getYammerGroupsActivityDetailWithPeriod and sets the default values.
func NewGetYammerGroupsActivityDetailWithPeriod()(*GetYammerGroupsActivityDetailWithPeriod) {
    m := &GetYammerGroupsActivityDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the groupDisplayName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupDisplayName
    }
}
// Gets the groupType property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
// Gets the isDeleted property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the likedCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetLikedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.likedCount
    }
}
// Gets the memberCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
// Gets the networkDisplayName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetNetworkDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDisplayName
    }
}
// Gets the office365Connected property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetOffice365Connected()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.office365Connected
    }
}
// Gets the ownerPrincipalName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// Gets the postedCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetPostedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.postedCount
    }
}
// Gets the readCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetReadCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.readCount
    }
}
// Gets the reportPeriod property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// The deserialization information for the current model
func (m *GetYammerGroupsActivityDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["groupDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDisplayName(val)
        }
        return nil
    }
    res["groupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupType(val)
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
    res["likedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLikedCount(val)
        }
        return nil
    }
    res["memberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberCount(val)
        }
        return nil
    }
    res["networkDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDisplayName(val)
        }
        return nil
    }
    res["office365Connected"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffice365Connected(val)
        }
        return nil
    }
    res["ownerPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerPrincipalName(val)
        }
        return nil
    }
    res["postedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostedCount(val)
        }
        return nil
    }
    res["readCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadCount(val)
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
    return res
}
func (m *GetYammerGroupsActivityDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetYammerGroupsActivityDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupDisplayName", m.GetGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupType", m.GetGroupType())
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
        err = writer.WriteInt64Value("likedCount", m.GetLikedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("memberCount", m.GetMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkDisplayName", m.GetNetworkDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("office365Connected", m.GetOffice365Connected())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerPrincipalName", m.GetOwnerPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("postedCount", m.GetPostedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("readCount", m.GetReadCount())
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
    return nil
}
// Sets the groupDisplayName property value. 
// Parameters:
//  - value : Value to set for the groupDisplayName property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetGroupDisplayName(value *string)() {
    m.groupDisplayName = value
}
// Sets the groupType property value. 
// Parameters:
//  - value : Value to set for the groupType property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetGroupType(value *string)() {
    m.groupType = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the likedCount property value. 
// Parameters:
//  - value : Value to set for the likedCount property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetLikedCount(value *int64)() {
    m.likedCount = value
}
// Sets the memberCount property value. 
// Parameters:
//  - value : Value to set for the memberCount property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetMemberCount(value *int64)() {
    m.memberCount = value
}
// Sets the networkDisplayName property value. 
// Parameters:
//  - value : Value to set for the networkDisplayName property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetNetworkDisplayName(value *string)() {
    m.networkDisplayName = value
}
// Sets the office365Connected property value. 
// Parameters:
//  - value : Value to set for the office365Connected property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetOffice365Connected(value *bool)() {
    m.office365Connected = value
}
// Sets the ownerPrincipalName property value. 
// Parameters:
//  - value : Value to set for the ownerPrincipalName property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// Sets the postedCount property value. 
// Parameters:
//  - value : Value to set for the postedCount property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetPostedCount(value *int64)() {
    m.postedCount = value
}
// Sets the readCount property value. 
// Parameters:
//  - value : Value to set for the readCount property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetReadCount(value *int64)() {
    m.readCount = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetYammerGroupsActivityDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
