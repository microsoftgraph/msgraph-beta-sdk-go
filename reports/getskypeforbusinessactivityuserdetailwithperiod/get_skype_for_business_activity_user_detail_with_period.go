package getskypeforbusinessactivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSkypeForBusinessActivityUserDetailWithPeriod 
type GetSkypeForBusinessActivityUserDetailWithPeriod struct {
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
// NewGetSkypeForBusinessActivityUserDetailWithPeriod instantiates a new getSkypeForBusinessActivityUserDetailWithPeriod and sets the default values.
func NewGetSkypeForBusinessActivityUserDetailWithPeriod()(*GetSkypeForBusinessActivityUserDetailWithPeriod) {
    m := &GetSkypeForBusinessActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAssignedProducts gets the assignedProducts property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// GetDeletedDate gets the deletedDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetOrganizedConferenceAppSharingCount gets the organizedConferenceAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAppSharingCount
    }
}
// GetOrganizedConferenceAudioVideoCount gets the organizedConferenceAudioVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceAudioVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAudioVideoCount
    }
}
// GetOrganizedConferenceAudioVideoMinutes gets the organizedConferenceAudioVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceAudioVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAudioVideoMinutes
    }
}
// GetOrganizedConferenceCloudDialInMicrosoftMinutes gets the organizedConferenceCloudDialInMicrosoftMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceCloudDialInMicrosoftMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialInMicrosoftMinutes
    }
}
// GetOrganizedConferenceCloudDialInOutMicrosoftCount gets the organizedConferenceCloudDialInOutMicrosoftCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceCloudDialInOutMicrosoftCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialInOutMicrosoftCount
    }
}
// GetOrganizedConferenceCloudDialOutMicrosoftMinutes gets the organizedConferenceCloudDialOutMicrosoftMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceCloudDialOutMicrosoftMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialOutMicrosoftMinutes
    }
}
// GetOrganizedConferenceDialInOut3rdPartyCount gets the organizedConferenceDialInOut3rdPartyCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceDialInOut3rdPartyCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceDialInOut3rdPartyCount
    }
}
// GetOrganizedConferenceIMCount gets the organizedConferenceIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceIMCount
    }
}
// GetOrganizedConferenceLastActivityDate gets the organizedConferenceLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceLastActivityDate
    }
}
// GetOrganizedConferenceWebCount gets the organizedConferenceWebCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceWebCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceWebCount
    }
}
// GetParticipatedConferenceAppSharingCount gets the participatedConferenceAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAppSharingCount
    }
}
// GetParticipatedConferenceAudioVideoCount gets the participatedConferenceAudioVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceAudioVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAudioVideoCount
    }
}
// GetParticipatedConferenceAudioVideoMinutes gets the participatedConferenceAudioVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceAudioVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAudioVideoMinutes
    }
}
// GetParticipatedConferenceDialInOut3rdPartyCount gets the participatedConferenceDialInOut3rdPartyCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceDialInOut3rdPartyCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceDialInOut3rdPartyCount
    }
}
// GetParticipatedConferenceIMCount gets the participatedConferenceIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceIMCount
    }
}
// GetParticipatedConferenceLastActivityDate gets the participatedConferenceLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceLastActivityDate
    }
}
// GetParticipatedConferenceWebCount gets the participatedConferenceWebCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceWebCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceWebCount
    }
}
// GetPeerToPeerAppSharingCount gets the peerToPeerAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAppSharingCount
    }
}
// GetPeerToPeerAudioCount gets the peerToPeerAudioCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerAudioCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAudioCount
    }
}
// GetPeerToPeerAudioMinutes gets the peerToPeerAudioMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerAudioMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAudioMinutes
    }
}
// GetPeerToPeerFileTransferCount gets the peerToPeerFileTransferCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerFileTransferCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerFileTransferCount
    }
}
// GetPeerToPeerIMCount gets the peerToPeerIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerIMCount
    }
}
// GetPeerToPeerLastActivityDate gets the peerToPeerLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerLastActivityDate
    }
}
// GetPeerToPeerVideoCount gets the peerToPeerVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerVideoCount
    }
}
// GetPeerToPeerVideoMinutes gets the peerToPeerVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerVideoMinutes
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetTotalOrganizedConferenceCount gets the totalOrganizedConferenceCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetTotalOrganizedConferenceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalOrganizedConferenceCount
    }
}
// GetTotalParticipatedConferenceCount gets the totalParticipatedConferenceCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetTotalParticipatedConferenceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalParticipatedConferenceCount
    }
}
// GetTotalPeerToPeerSessionCount gets the totalPeerToPeerSessionCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetTotalPeerToPeerSessionCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalPeerToPeerSessionCount
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["organizedConferenceAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceAppSharingCount(val)
        }
        return nil
    }
    res["organizedConferenceAudioVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceAudioVideoCount(val)
        }
        return nil
    }
    res["organizedConferenceAudioVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceAudioVideoMinutes(val)
        }
        return nil
    }
    res["organizedConferenceCloudDialInMicrosoftMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceCloudDialInMicrosoftMinutes(val)
        }
        return nil
    }
    res["organizedConferenceCloudDialInOutMicrosoftCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceCloudDialInOutMicrosoftCount(val)
        }
        return nil
    }
    res["organizedConferenceCloudDialOutMicrosoftMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceCloudDialOutMicrosoftMinutes(val)
        }
        return nil
    }
    res["organizedConferenceDialInOut3rdPartyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceDialInOut3rdPartyCount(val)
        }
        return nil
    }
    res["organizedConferenceIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceIMCount(val)
        }
        return nil
    }
    res["organizedConferenceLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceLastActivityDate(val)
        }
        return nil
    }
    res["organizedConferenceWebCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizedConferenceWebCount(val)
        }
        return nil
    }
    res["participatedConferenceAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceAppSharingCount(val)
        }
        return nil
    }
    res["participatedConferenceAudioVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceAudioVideoCount(val)
        }
        return nil
    }
    res["participatedConferenceAudioVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceAudioVideoMinutes(val)
        }
        return nil
    }
    res["participatedConferenceDialInOut3rdPartyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceDialInOut3rdPartyCount(val)
        }
        return nil
    }
    res["participatedConferenceIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceIMCount(val)
        }
        return nil
    }
    res["participatedConferenceLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceLastActivityDate(val)
        }
        return nil
    }
    res["participatedConferenceWebCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipatedConferenceWebCount(val)
        }
        return nil
    }
    res["peerToPeerAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerAppSharingCount(val)
        }
        return nil
    }
    res["peerToPeerAudioCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerAudioCount(val)
        }
        return nil
    }
    res["peerToPeerAudioMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerAudioMinutes(val)
        }
        return nil
    }
    res["peerToPeerFileTransferCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerFileTransferCount(val)
        }
        return nil
    }
    res["peerToPeerIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerIMCount(val)
        }
        return nil
    }
    res["peerToPeerLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerLastActivityDate(val)
        }
        return nil
    }
    res["peerToPeerVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerVideoCount(val)
        }
        return nil
    }
    res["peerToPeerVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeerToPeerVideoMinutes(val)
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
    res["totalOrganizedConferenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalOrganizedConferenceCount(val)
        }
        return nil
    }
    res["totalParticipatedConferenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalParticipatedConferenceCount(val)
        }
        return nil
    }
    res["totalPeerToPeerSessionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPeerToPeerSessionCount(val)
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
    return res
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAssignedProducts sets the assignedProducts property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    if m != nil {
        m.assignedProducts = value
    }
}
// SetDeletedDate sets the deletedDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetDeletedDate(value *string)() {
    if m != nil {
        m.deletedDate = value
    }
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetOrganizedConferenceAppSharingCount sets the organizedConferenceAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceAppSharingCount(value *int64)() {
    if m != nil {
        m.organizedConferenceAppSharingCount = value
    }
}
// SetOrganizedConferenceAudioVideoCount sets the organizedConferenceAudioVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceAudioVideoCount(value *int64)() {
    if m != nil {
        m.organizedConferenceAudioVideoCount = value
    }
}
// SetOrganizedConferenceAudioVideoMinutes sets the organizedConferenceAudioVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceAudioVideoMinutes(value *int64)() {
    if m != nil {
        m.organizedConferenceAudioVideoMinutes = value
    }
}
// SetOrganizedConferenceCloudDialInMicrosoftMinutes sets the organizedConferenceCloudDialInMicrosoftMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceCloudDialInMicrosoftMinutes(value *int64)() {
    if m != nil {
        m.organizedConferenceCloudDialInMicrosoftMinutes = value
    }
}
// SetOrganizedConferenceCloudDialInOutMicrosoftCount sets the organizedConferenceCloudDialInOutMicrosoftCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceCloudDialInOutMicrosoftCount(value *int64)() {
    if m != nil {
        m.organizedConferenceCloudDialInOutMicrosoftCount = value
    }
}
// SetOrganizedConferenceCloudDialOutMicrosoftMinutes sets the organizedConferenceCloudDialOutMicrosoftMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceCloudDialOutMicrosoftMinutes(value *int64)() {
    if m != nil {
        m.organizedConferenceCloudDialOutMicrosoftMinutes = value
    }
}
// SetOrganizedConferenceDialInOut3rdPartyCount sets the organizedConferenceDialInOut3rdPartyCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceDialInOut3rdPartyCount(value *int64)() {
    if m != nil {
        m.organizedConferenceDialInOut3rdPartyCount = value
    }
}
// SetOrganizedConferenceIMCount sets the organizedConferenceIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceIMCount(value *int64)() {
    if m != nil {
        m.organizedConferenceIMCount = value
    }
}
// SetOrganizedConferenceLastActivityDate sets the organizedConferenceLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceLastActivityDate(value *string)() {
    if m != nil {
        m.organizedConferenceLastActivityDate = value
    }
}
// SetOrganizedConferenceWebCount sets the organizedConferenceWebCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceWebCount(value *int64)() {
    if m != nil {
        m.organizedConferenceWebCount = value
    }
}
// SetParticipatedConferenceAppSharingCount sets the participatedConferenceAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceAppSharingCount(value *int64)() {
    if m != nil {
        m.participatedConferenceAppSharingCount = value
    }
}
// SetParticipatedConferenceAudioVideoCount sets the participatedConferenceAudioVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceAudioVideoCount(value *int64)() {
    if m != nil {
        m.participatedConferenceAudioVideoCount = value
    }
}
// SetParticipatedConferenceAudioVideoMinutes sets the participatedConferenceAudioVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceAudioVideoMinutes(value *int64)() {
    if m != nil {
        m.participatedConferenceAudioVideoMinutes = value
    }
}
// SetParticipatedConferenceDialInOut3rdPartyCount sets the participatedConferenceDialInOut3rdPartyCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceDialInOut3rdPartyCount(value *int64)() {
    if m != nil {
        m.participatedConferenceDialInOut3rdPartyCount = value
    }
}
// SetParticipatedConferenceIMCount sets the participatedConferenceIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceIMCount(value *int64)() {
    if m != nil {
        m.participatedConferenceIMCount = value
    }
}
// SetParticipatedConferenceLastActivityDate sets the participatedConferenceLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceLastActivityDate(value *string)() {
    if m != nil {
        m.participatedConferenceLastActivityDate = value
    }
}
// SetParticipatedConferenceWebCount sets the participatedConferenceWebCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceWebCount(value *int64)() {
    if m != nil {
        m.participatedConferenceWebCount = value
    }
}
// SetPeerToPeerAppSharingCount sets the peerToPeerAppSharingCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerAppSharingCount(value *int64)() {
    if m != nil {
        m.peerToPeerAppSharingCount = value
    }
}
// SetPeerToPeerAudioCount sets the peerToPeerAudioCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerAudioCount(value *int64)() {
    if m != nil {
        m.peerToPeerAudioCount = value
    }
}
// SetPeerToPeerAudioMinutes sets the peerToPeerAudioMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerAudioMinutes(value *int64)() {
    if m != nil {
        m.peerToPeerAudioMinutes = value
    }
}
// SetPeerToPeerFileTransferCount sets the peerToPeerFileTransferCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerFileTransferCount(value *int64)() {
    if m != nil {
        m.peerToPeerFileTransferCount = value
    }
}
// SetPeerToPeerIMCount sets the peerToPeerIMCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerIMCount(value *int64)() {
    if m != nil {
        m.peerToPeerIMCount = value
    }
}
// SetPeerToPeerLastActivityDate sets the peerToPeerLastActivityDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerLastActivityDate(value *string)() {
    if m != nil {
        m.peerToPeerLastActivityDate = value
    }
}
// SetPeerToPeerVideoCount sets the peerToPeerVideoCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerVideoCount(value *int64)() {
    if m != nil {
        m.peerToPeerVideoCount = value
    }
}
// SetPeerToPeerVideoMinutes sets the peerToPeerVideoMinutes property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerVideoMinutes(value *int64)() {
    if m != nil {
        m.peerToPeerVideoMinutes = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetTotalOrganizedConferenceCount sets the totalOrganizedConferenceCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetTotalOrganizedConferenceCount(value *int64)() {
    if m != nil {
        m.totalOrganizedConferenceCount = value
    }
}
// SetTotalParticipatedConferenceCount sets the totalParticipatedConferenceCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetTotalParticipatedConferenceCount(value *int64)() {
    if m != nil {
        m.totalParticipatedConferenceCount = value
    }
}
// SetTotalPeerToPeerSessionCount sets the totalPeerToPeerSessionCount property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetTotalPeerToPeerSessionCount(value *int64)() {
    if m != nil {
        m.totalPeerToPeerSessionCount = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
