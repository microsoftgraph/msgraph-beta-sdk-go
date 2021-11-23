package getyammeractivityuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getYammerActivityUserDetailWithDate 
type GetYammerActivityUserDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    assignedProducts []string;
    // 
    displayName *string;
    // 
    lastActivityDate *string;
    // 
    likedCount *int64;
    // 
    postedCount *int64;
    // 
    readCount *int64;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    stateChangeDate *string;
    // 
    userPrincipalName *string;
    // 
    userState *string;
}
// NewGetYammerActivityUserDetailWithDate instantiates a new getYammerActivityUserDetailWithDate and sets the default values.
func NewGetYammerActivityUserDetailWithDate()(*GetYammerActivityUserDetailWithDate) {
    m := &GetYammerActivityUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAssignedProducts gets the assignedProducts property value. 
func (m *GetYammerActivityUserDetailWithDate) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// GetDisplayName gets the displayName property value. 
func (m *GetYammerActivityUserDetailWithDate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetYammerActivityUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetLikedCount gets the likedCount property value. 
func (m *GetYammerActivityUserDetailWithDate) GetLikedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.likedCount
    }
}
// GetPostedCount gets the postedCount property value. 
func (m *GetYammerActivityUserDetailWithDate) GetPostedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.postedCount
    }
}
// GetReadCount gets the readCount property value. 
func (m *GetYammerActivityUserDetailWithDate) GetReadCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.readCount
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetYammerActivityUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetYammerActivityUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetStateChangeDate gets the stateChangeDate property value. 
func (m *GetYammerActivityUserDetailWithDate) GetStateChangeDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateChangeDate
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetYammerActivityUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserState gets the userState property value. 
func (m *GetYammerActivityUserDetailWithDate) GetUserState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetYammerActivityUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["stateChangeDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateChangeDate(val)
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
    res["userState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserState(val)
        }
        return nil
    }
    return res
}
func (m *GetYammerActivityUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetYammerActivityUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    {
        err = writer.WriteStringValue("stateChangeDate", m.GetStateChangeDate())
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
        err = writer.WriteStringValue("userState", m.GetUserState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedProducts sets the assignedProducts property value. 
func (m *GetYammerActivityUserDetailWithDate) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// SetDisplayName sets the displayName property value. 
func (m *GetYammerActivityUserDetailWithDate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetYammerActivityUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// SetLikedCount sets the likedCount property value. 
func (m *GetYammerActivityUserDetailWithDate) SetLikedCount(value *int64)() {
    m.likedCount = value
}
// SetPostedCount sets the postedCount property value. 
func (m *GetYammerActivityUserDetailWithDate) SetPostedCount(value *int64)() {
    m.postedCount = value
}
// SetReadCount sets the readCount property value. 
func (m *GetYammerActivityUserDetailWithDate) SetReadCount(value *int64)() {
    m.readCount = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetYammerActivityUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetYammerActivityUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetStateChangeDate sets the stateChangeDate property value. 
func (m *GetYammerActivityUserDetailWithDate) SetStateChangeDate(value *string)() {
    m.stateChangeDate = value
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetYammerActivityUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserState sets the userState property value. 
func (m *GetYammerActivityUserDetailWithDate) SetUserState(value *string)() {
    m.userState = value
}
