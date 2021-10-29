package getskypeforbusinessactivityuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetSkypeForBusinessActivityUserDetailWithDate struct {
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
    organizedConferenceAppSharingCount *int64;
    // 
    organizedConferenceAudioVideoCount *int64;
    // 
    organizedConferenceAudioVideoMinutes *int64;
    // 
    organizedConferenceCloudDialInMicrosoftMinutes *int64;
    // 
    organizedConferenceCloudDialInOutMicrosoftCount *int64;
    // 
    organizedConferenceCloudDialOutMicrosoftMinutes *int64;
    // 
    organizedConferenceDialInOut3rdPartyCount *int64;
    // 
    organizedConferenceIMCount *int64;
    // 
    organizedConferenceLastActivityDate *string;
    // 
    organizedConferenceWebCount *int64;
    // 
    participatedConferenceAppSharingCount *int64;
    // 
    participatedConferenceAudioVideoCount *int64;
    // 
    participatedConferenceAudioVideoMinutes *int64;
    // 
    participatedConferenceDialInOut3rdPartyCount *int64;
    // 
    participatedConferenceIMCount *int64;
    // 
    participatedConferenceLastActivityDate *string;
    // 
    participatedConferenceWebCount *int64;
    // 
    peerToPeerAppSharingCount *int64;
    // 
    peerToPeerAudioCount *int64;
    // 
    peerToPeerAudioMinutes *int64;
    // 
    peerToPeerFileTransferCount *int64;
    // 
    peerToPeerIMCount *int64;
    // 
    peerToPeerLastActivityDate *string;
    // 
    peerToPeerVideoCount *int64;
    // 
    peerToPeerVideoMinutes *int64;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    totalOrganizedConferenceCount *int64;
    // 
    totalParticipatedConferenceCount *int64;
    // 
    totalPeerToPeerSessionCount *int64;
    // 
    userPrincipalName *string;
}
// Instantiates a new getSkypeForBusinessActivityUserDetailWithDate and sets the default values.
func NewGetSkypeForBusinessActivityUserDetailWithDate()(*GetSkypeForBusinessActivityUserDetailWithDate) {
    m := &GetSkypeForBusinessActivityUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the assignedProducts property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// Gets the deletedDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// Gets the isDeleted property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the organizedConferenceAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAppSharingCount
    }
}
// Gets the organizedConferenceAudioVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceAudioVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAudioVideoCount
    }
}
// Gets the organizedConferenceAudioVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceAudioVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAudioVideoMinutes
    }
}
// Gets the organizedConferenceCloudDialInMicrosoftMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceCloudDialInMicrosoftMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialInMicrosoftMinutes
    }
}
// Gets the organizedConferenceCloudDialInOutMicrosoftCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceCloudDialInOutMicrosoftCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialInOutMicrosoftCount
    }
}
// Gets the organizedConferenceCloudDialOutMicrosoftMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceCloudDialOutMicrosoftMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialOutMicrosoftMinutes
    }
}
// Gets the organizedConferenceDialInOut3rdPartyCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceDialInOut3rdPartyCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceDialInOut3rdPartyCount
    }
}
// Gets the organizedConferenceIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceIMCount
    }
}
// Gets the organizedConferenceLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceLastActivityDate
    }
}
// Gets the organizedConferenceWebCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetOrganizedConferenceWebCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceWebCount
    }
}
// Gets the participatedConferenceAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAppSharingCount
    }
}
// Gets the participatedConferenceAudioVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceAudioVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAudioVideoCount
    }
}
// Gets the participatedConferenceAudioVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceAudioVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAudioVideoMinutes
    }
}
// Gets the participatedConferenceDialInOut3rdPartyCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceDialInOut3rdPartyCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceDialInOut3rdPartyCount
    }
}
// Gets the participatedConferenceIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceIMCount
    }
}
// Gets the participatedConferenceLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceLastActivityDate
    }
}
// Gets the participatedConferenceWebCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetParticipatedConferenceWebCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceWebCount
    }
}
// Gets the peerToPeerAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAppSharingCount
    }
}
// Gets the peerToPeerAudioCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerAudioCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAudioCount
    }
}
// Gets the peerToPeerAudioMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerAudioMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAudioMinutes
    }
}
// Gets the peerToPeerFileTransferCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerFileTransferCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerFileTransferCount
    }
}
// Gets the peerToPeerIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerIMCount
    }
}
// Gets the peerToPeerLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerLastActivityDate
    }
}
// Gets the peerToPeerVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerVideoCount
    }
}
// Gets the peerToPeerVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetPeerToPeerVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerVideoMinutes
    }
}
// Gets the reportPeriod property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the totalOrganizedConferenceCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetTotalOrganizedConferenceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalOrganizedConferenceCount
    }
}
// Gets the totalParticipatedConferenceCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetTotalParticipatedConferenceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalParticipatedConferenceCount
    }
}
// Gets the totalPeerToPeerSessionCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetTotalPeerToPeerSessionCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalPeerToPeerSessionCount
    }
}
// Gets the userPrincipalName property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *GetSkypeForBusinessActivityUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["organizedConferenceAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceAppSharingCount(val)
        return nil
    }
    res["organizedConferenceAudioVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceAudioVideoCount(val)
        return nil
    }
    res["organizedConferenceAudioVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceAudioVideoMinutes(val)
        return nil
    }
    res["organizedConferenceCloudDialInMicrosoftMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceCloudDialInMicrosoftMinutes(val)
        return nil
    }
    res["organizedConferenceCloudDialInOutMicrosoftCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceCloudDialInOutMicrosoftCount(val)
        return nil
    }
    res["organizedConferenceCloudDialOutMicrosoftMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceCloudDialOutMicrosoftMinutes(val)
        return nil
    }
    res["organizedConferenceDialInOut3rdPartyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceDialInOut3rdPartyCount(val)
        return nil
    }
    res["organizedConferenceIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceIMCount(val)
        return nil
    }
    res["organizedConferenceLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceLastActivityDate(val)
        return nil
    }
    res["organizedConferenceWebCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceWebCount(val)
        return nil
    }
    res["participatedConferenceAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceAppSharingCount(val)
        return nil
    }
    res["participatedConferenceAudioVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceAudioVideoCount(val)
        return nil
    }
    res["participatedConferenceAudioVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceAudioVideoMinutes(val)
        return nil
    }
    res["participatedConferenceDialInOut3rdPartyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceDialInOut3rdPartyCount(val)
        return nil
    }
    res["participatedConferenceIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceIMCount(val)
        return nil
    }
    res["participatedConferenceLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceLastActivityDate(val)
        return nil
    }
    res["participatedConferenceWebCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceWebCount(val)
        return nil
    }
    res["peerToPeerAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerAppSharingCount(val)
        return nil
    }
    res["peerToPeerAudioCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerAudioCount(val)
        return nil
    }
    res["peerToPeerAudioMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerAudioMinutes(val)
        return nil
    }
    res["peerToPeerFileTransferCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerFileTransferCount(val)
        return nil
    }
    res["peerToPeerIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerIMCount(val)
        return nil
    }
    res["peerToPeerLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeerToPeerLastActivityDate(val)
        return nil
    }
    res["peerToPeerVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerVideoCount(val)
        return nil
    }
    res["peerToPeerVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerVideoMinutes(val)
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
    res["totalOrganizedConferenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalOrganizedConferenceCount(val)
        return nil
    }
    res["totalParticipatedConferenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalParticipatedConferenceCount(val)
        return nil
    }
    res["totalPeerToPeerSessionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalPeerToPeerSessionCount(val)
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
    return res
}
func (m *GetSkypeForBusinessActivityUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetSkypeForBusinessActivityUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("organizedConferenceAppSharingCount", m.GetOrganizedConferenceAppSharingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceAudioVideoCount", m.GetOrganizedConferenceAudioVideoCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceAudioVideoMinutes", m.GetOrganizedConferenceAudioVideoMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceCloudDialInMicrosoftMinutes", m.GetOrganizedConferenceCloudDialInMicrosoftMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceCloudDialInOutMicrosoftCount", m.GetOrganizedConferenceCloudDialInOutMicrosoftCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceCloudDialOutMicrosoftMinutes", m.GetOrganizedConferenceCloudDialOutMicrosoftMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceDialInOut3rdPartyCount", m.GetOrganizedConferenceDialInOut3rdPartyCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceIMCount", m.GetOrganizedConferenceIMCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizedConferenceLastActivityDate", m.GetOrganizedConferenceLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceWebCount", m.GetOrganizedConferenceWebCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceAppSharingCount", m.GetParticipatedConferenceAppSharingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceAudioVideoCount", m.GetParticipatedConferenceAudioVideoCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceAudioVideoMinutes", m.GetParticipatedConferenceAudioVideoMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceDialInOut3rdPartyCount", m.GetParticipatedConferenceDialInOut3rdPartyCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceIMCount", m.GetParticipatedConferenceIMCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("participatedConferenceLastActivityDate", m.GetParticipatedConferenceLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceWebCount", m.GetParticipatedConferenceWebCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerAppSharingCount", m.GetPeerToPeerAppSharingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerAudioCount", m.GetPeerToPeerAudioCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerAudioMinutes", m.GetPeerToPeerAudioMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerFileTransferCount", m.GetPeerToPeerFileTransferCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerIMCount", m.GetPeerToPeerIMCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("peerToPeerLastActivityDate", m.GetPeerToPeerLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerVideoCount", m.GetPeerToPeerVideoCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerVideoMinutes", m.GetPeerToPeerVideoMinutes())
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
        err = writer.WriteInt64Value("totalOrganizedConferenceCount", m.GetTotalOrganizedConferenceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalParticipatedConferenceCount", m.GetTotalParticipatedConferenceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalPeerToPeerSessionCount", m.GetTotalPeerToPeerSessionCount())
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
    return nil
}
// Sets the assignedProducts property value. 
// Parameters:
//  - value : Value to set for the assignedProducts property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// Sets the deletedDate property value. 
// Parameters:
//  - value : Value to set for the deletedDate property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the organizedConferenceAppSharingCount property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceAppSharingCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceAppSharingCount(value *int64)() {
    m.organizedConferenceAppSharingCount = value
}
// Sets the organizedConferenceAudioVideoCount property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceAudioVideoCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceAudioVideoCount(value *int64)() {
    m.organizedConferenceAudioVideoCount = value
}
// Sets the organizedConferenceAudioVideoMinutes property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceAudioVideoMinutes property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceAudioVideoMinutes(value *int64)() {
    m.organizedConferenceAudioVideoMinutes = value
}
// Sets the organizedConferenceCloudDialInMicrosoftMinutes property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceCloudDialInMicrosoftMinutes property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceCloudDialInMicrosoftMinutes(value *int64)() {
    m.organizedConferenceCloudDialInMicrosoftMinutes = value
}
// Sets the organizedConferenceCloudDialInOutMicrosoftCount property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceCloudDialInOutMicrosoftCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceCloudDialInOutMicrosoftCount(value *int64)() {
    m.organizedConferenceCloudDialInOutMicrosoftCount = value
}
// Sets the organizedConferenceCloudDialOutMicrosoftMinutes property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceCloudDialOutMicrosoftMinutes property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceCloudDialOutMicrosoftMinutes(value *int64)() {
    m.organizedConferenceCloudDialOutMicrosoftMinutes = value
}
// Sets the organizedConferenceDialInOut3rdPartyCount property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceDialInOut3rdPartyCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceDialInOut3rdPartyCount(value *int64)() {
    m.organizedConferenceDialInOut3rdPartyCount = value
}
// Sets the organizedConferenceIMCount property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceIMCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceIMCount(value *int64)() {
    m.organizedConferenceIMCount = value
}
// Sets the organizedConferenceLastActivityDate property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceLastActivityDate property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceLastActivityDate(value *string)() {
    m.organizedConferenceLastActivityDate = value
}
// Sets the organizedConferenceWebCount property value. 
// Parameters:
//  - value : Value to set for the organizedConferenceWebCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetOrganizedConferenceWebCount(value *int64)() {
    m.organizedConferenceWebCount = value
}
// Sets the participatedConferenceAppSharingCount property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceAppSharingCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceAppSharingCount(value *int64)() {
    m.participatedConferenceAppSharingCount = value
}
// Sets the participatedConferenceAudioVideoCount property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceAudioVideoCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceAudioVideoCount(value *int64)() {
    m.participatedConferenceAudioVideoCount = value
}
// Sets the participatedConferenceAudioVideoMinutes property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceAudioVideoMinutes property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceAudioVideoMinutes(value *int64)() {
    m.participatedConferenceAudioVideoMinutes = value
}
// Sets the participatedConferenceDialInOut3rdPartyCount property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceDialInOut3rdPartyCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceDialInOut3rdPartyCount(value *int64)() {
    m.participatedConferenceDialInOut3rdPartyCount = value
}
// Sets the participatedConferenceIMCount property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceIMCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceIMCount(value *int64)() {
    m.participatedConferenceIMCount = value
}
// Sets the participatedConferenceLastActivityDate property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceLastActivityDate property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceLastActivityDate(value *string)() {
    m.participatedConferenceLastActivityDate = value
}
// Sets the participatedConferenceWebCount property value. 
// Parameters:
//  - value : Value to set for the participatedConferenceWebCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetParticipatedConferenceWebCount(value *int64)() {
    m.participatedConferenceWebCount = value
}
// Sets the peerToPeerAppSharingCount property value. 
// Parameters:
//  - value : Value to set for the peerToPeerAppSharingCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerAppSharingCount(value *int64)() {
    m.peerToPeerAppSharingCount = value
}
// Sets the peerToPeerAudioCount property value. 
// Parameters:
//  - value : Value to set for the peerToPeerAudioCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerAudioCount(value *int64)() {
    m.peerToPeerAudioCount = value
}
// Sets the peerToPeerAudioMinutes property value. 
// Parameters:
//  - value : Value to set for the peerToPeerAudioMinutes property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerAudioMinutes(value *int64)() {
    m.peerToPeerAudioMinutes = value
}
// Sets the peerToPeerFileTransferCount property value. 
// Parameters:
//  - value : Value to set for the peerToPeerFileTransferCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerFileTransferCount(value *int64)() {
    m.peerToPeerFileTransferCount = value
}
// Sets the peerToPeerIMCount property value. 
// Parameters:
//  - value : Value to set for the peerToPeerIMCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerIMCount(value *int64)() {
    m.peerToPeerIMCount = value
}
// Sets the peerToPeerLastActivityDate property value. 
// Parameters:
//  - value : Value to set for the peerToPeerLastActivityDate property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerLastActivityDate(value *string)() {
    m.peerToPeerLastActivityDate = value
}
// Sets the peerToPeerVideoCount property value. 
// Parameters:
//  - value : Value to set for the peerToPeerVideoCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerVideoCount(value *int64)() {
    m.peerToPeerVideoCount = value
}
// Sets the peerToPeerVideoMinutes property value. 
// Parameters:
//  - value : Value to set for the peerToPeerVideoMinutes property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetPeerToPeerVideoMinutes(value *int64)() {
    m.peerToPeerVideoMinutes = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the totalOrganizedConferenceCount property value. 
// Parameters:
//  - value : Value to set for the totalOrganizedConferenceCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetTotalOrganizedConferenceCount(value *int64)() {
    m.totalOrganizedConferenceCount = value
}
// Sets the totalParticipatedConferenceCount property value. 
// Parameters:
//  - value : Value to set for the totalParticipatedConferenceCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetTotalParticipatedConferenceCount(value *int64)() {
    m.totalParticipatedConferenceCount = value
}
// Sets the totalPeerToPeerSessionCount property value. 
// Parameters:
//  - value : Value to set for the totalPeerToPeerSessionCount property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetTotalPeerToPeerSessionCount(value *int64)() {
    m.totalPeerToPeerSessionCount = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetSkypeForBusinessActivityUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
