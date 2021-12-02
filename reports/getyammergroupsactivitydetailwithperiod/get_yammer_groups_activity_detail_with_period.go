package getyammergroupsactivitydetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetYammerGroupsActivityDetailWithPeriod 
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
// NewGetYammerGroupsActivityDetailWithPeriod instantiates a new getYammerGroupsActivityDetailWithPeriod and sets the default values.
func NewGetYammerGroupsActivityDetailWithPeriod()(*GetYammerGroupsActivityDetailWithPeriod) {
    m := &GetYammerGroupsActivityDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetGroupDisplayName gets the groupDisplayName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupDisplayName
    }
}
// GetGroupType gets the groupType property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetLikedCount gets the likedCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetLikedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.likedCount
    }
}
// GetMemberCount gets the memberCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
// GetNetworkDisplayName gets the networkDisplayName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetNetworkDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDisplayName
    }
}
// GetOffice365Connected gets the office365Connected property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetOffice365Connected()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.office365Connected
    }
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// GetPostedCount gets the postedCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetPostedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.postedCount
    }
}
// GetReadCount gets the readCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetReadCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.readCount
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetGroupDisplayName sets the groupDisplayName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetGroupDisplayName(value *string)() {
    if m != nil {
        m.groupDisplayName = value
    }
}
// SetGroupType sets the groupType property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetGroupType(value *string)() {
    if m != nil {
        m.groupType = value
    }
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetLastActivityDate(value *string)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetLikedCount sets the likedCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetLikedCount(value *int64)() {
    if m != nil {
        m.likedCount = value
    }
}
// SetMemberCount sets the memberCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetMemberCount(value *int64)() {
    if m != nil {
        m.memberCount = value
    }
}
// SetNetworkDisplayName sets the networkDisplayName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetNetworkDisplayName(value *string)() {
    if m != nil {
        m.networkDisplayName = value
    }
}
// SetOffice365Connected sets the office365Connected property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetOffice365Connected(value *bool)() {
    if m != nil {
        m.office365Connected = value
    }
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    if m != nil {
        m.ownerPrincipalName = value
    }
}
// SetPostedCount sets the postedCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetPostedCount(value *int64)() {
    if m != nil {
        m.postedCount = value
    }
}
// SetReadCount sets the readCount property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetReadCount(value *int64)() {
    if m != nil {
        m.readCount = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetYammerGroupsActivityDetailWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
