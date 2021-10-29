package getyammeractivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetYammerActivityUserDetailWithPeriod struct {
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
// Instantiates a new getYammerActivityUserDetailWithPeriod and sets the default values.
func NewGetYammerActivityUserDetailWithPeriod()(*GetYammerActivityUserDetailWithPeriod) {
    m := &GetYammerActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the assignedProducts property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// Gets the displayName property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastActivityDate property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the likedCount property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetLikedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.likedCount
    }
}
// Gets the postedCount property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetPostedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.postedCount
    }
}
// Gets the readCount property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetReadCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.readCount
    }
}
// Gets the reportPeriod property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the stateChangeDate property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetStateChangeDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateChangeDate
    }
}
// Gets the userPrincipalName property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the userState property value. 
func (m *GetYammerActivityUserDetailWithPeriod) GetUserState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userState
    }
}
// The deserialization information for the current model
func (m *GetYammerActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignedProducts(res)
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
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["likedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLikedCount(val)
        return nil
    }
    res["postedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPostedCount(val)
        return nil
    }
    res["readCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetReadCount(val)
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
    res["stateChangeDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStateChangeDate(val)
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
    res["userState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserState(val)
        return nil
    }
    return res
}
func (m *GetYammerActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetYammerActivityUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the assignedProducts property value. 
// Parameters:
//  - value : Value to set for the assignedProducts property.
func (m *GetYammerActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GetYammerActivityUserDetailWithPeriod) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetYammerActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the likedCount property value. 
// Parameters:
//  - value : Value to set for the likedCount property.
func (m *GetYammerActivityUserDetailWithPeriod) SetLikedCount(value *int64)() {
    m.likedCount = value
}
// Sets the postedCount property value. 
// Parameters:
//  - value : Value to set for the postedCount property.
func (m *GetYammerActivityUserDetailWithPeriod) SetPostedCount(value *int64)() {
    m.postedCount = value
}
// Sets the readCount property value. 
// Parameters:
//  - value : Value to set for the readCount property.
func (m *GetYammerActivityUserDetailWithPeriod) SetReadCount(value *int64)() {
    m.readCount = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetYammerActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetYammerActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the stateChangeDate property value. 
// Parameters:
//  - value : Value to set for the stateChangeDate property.
func (m *GetYammerActivityUserDetailWithPeriod) SetStateChangeDate(value *string)() {
    m.stateChangeDate = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetYammerActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the userState property value. 
// Parameters:
//  - value : Value to set for the userState property.
func (m *GetYammerActivityUserDetailWithPeriod) SetUserState(value *string)() {
    m.userState = value
}
