package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type AnalyzedEmail struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAnalyzedEmail instantiates a new AnalyzedEmail and sets the default values.
func NewAnalyzedEmail()(*AnalyzedEmail) {
    m := &AnalyzedEmail{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAnalyzedEmailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnalyzedEmailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnalyzedEmail(), nil
}
// GetAlertIds gets the alertIds property value. A collection of values that contain the IDs of any alerts associated with the email.
// returns a []string when successful
func (m *AnalyzedEmail) GetAlertIds()([]string) {
    val, err := m.GetBackingStore().Get("alertIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAttachments gets the attachments property value. A collection of the attachments in the email.
// returns a []AnalyzedEmailAttachmentable when successful
func (m *AnalyzedEmail) GetAttachments()([]AnalyzedEmailAttachmentable) {
    val, err := m.GetBackingStore().Get("attachments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AnalyzedEmailAttachmentable)
    }
    return nil
}
// GetAttachmentsCount gets the attachmentsCount property value. The number of attachments in the email.
// returns a *int32 when successful
func (m *AnalyzedEmail) GetAttachmentsCount()(*int32) {
    val, err := m.GetBackingStore().Get("attachmentsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAuthenticationDetails gets the authenticationDetails property value. The authentication details associated with the email.
// returns a AnalyzedEmailAuthenticationDetailable when successful
func (m *AnalyzedEmail) GetAuthenticationDetails()(AnalyzedEmailAuthenticationDetailable) {
    val, err := m.GetBackingStore().Get("authenticationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AnalyzedEmailAuthenticationDetailable)
    }
    return nil
}
// GetBulkComplaintLevel gets the bulkComplaintLevel property value. The bulk complaint level of the email. A higher level is more likely to be spam.
// returns a *string when successful
func (m *AnalyzedEmail) GetBulkComplaintLevel()(*string) {
    val, err := m.GetBackingStore().Get("bulkComplaintLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContexts gets the contexts property value. Provides context of the email.
// returns a []string when successful
func (m *AnalyzedEmail) GetContexts()([]string) {
    val, err := m.GetBackingStore().Get("contexts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDetectionMethods gets the detectionMethods property value. The methods of detection used.
// returns a []string when successful
func (m *AnalyzedEmail) GetDetectionMethods()([]string) {
    val, err := m.GetBackingStore().Get("detectionMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDirectionality gets the directionality property value. The direction of the emails. The possible values are: unknown, inbound, outbound, intraOrg, unknownFutureValue.
// returns a *AntispamDirectionality when successful
func (m *AnalyzedEmail) GetDirectionality()(*AntispamDirectionality) {
    val, err := m.GetBackingStore().Get("directionality")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AntispamDirectionality)
    }
    return nil
}
// GetDistributionList gets the distributionList property value. The distribution list details to which the email was sent.
// returns a *string when successful
func (m *AnalyzedEmail) GetDistributionList()(*string) {
    val, err := m.GetBackingStore().Get("distributionList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailClusterId gets the emailClusterId property value. The identifier for the group of similar emails clustered based on heuristic analysis of their content.
// returns a *string when successful
func (m *AnalyzedEmail) GetEmailClusterId()(*string) {
    val, err := m.GetBackingStore().Get("emailClusterId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExchangeTransportRules gets the exchangeTransportRules property value. The name of the Exchange transport rules (ETRs) associated with the email.
// returns a []AnalyzedEmailExchangeTransportRuleInfoable when successful
func (m *AnalyzedEmail) GetExchangeTransportRules()([]AnalyzedEmailExchangeTransportRuleInfoable) {
    val, err := m.GetBackingStore().Get("exchangeTransportRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AnalyzedEmailExchangeTransportRuleInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AnalyzedEmail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAlertIds(res)
        }
        return nil
    }
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAnalyzedEmailAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnalyzedEmailAttachmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AnalyzedEmailAttachmentable)
                }
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["attachmentsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentsCount(val)
        }
        return nil
    }
    res["authenticationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAnalyzedEmailAuthenticationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationDetails(val.(AnalyzedEmailAuthenticationDetailable))
        }
        return nil
    }
    res["bulkComplaintLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBulkComplaintLevel(val)
        }
        return nil
    }
    res["contexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetContexts(res)
        }
        return nil
    }
    res["detectionMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDetectionMethods(res)
        }
        return nil
    }
    res["directionality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAntispamDirectionality)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectionality(val.(*AntispamDirectionality))
        }
        return nil
    }
    res["distributionList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDistributionList(val)
        }
        return nil
    }
    res["emailClusterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailClusterId(val)
        }
        return nil
    }
    res["exchangeTransportRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAnalyzedEmailExchangeTransportRuleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnalyzedEmailExchangeTransportRuleInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AnalyzedEmailExchangeTransportRuleInfoable)
                }
            }
            m.SetExchangeTransportRules(res)
        }
        return nil
    }
    res["internetMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["latestDelivery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAnalyzedEmailDeliveryDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestDelivery(val.(AnalyzedEmailDeliveryDetailable))
        }
        return nil
    }
    res["loggedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggedDateTime(val)
        }
        return nil
    }
    res["networkMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkMessageId(val)
        }
        return nil
    }
    res["originalDelivery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAnalyzedEmailDeliveryDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalDelivery(val.(AnalyzedEmailDeliveryDetailable))
        }
        return nil
    }
    res["overrideSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetOverrideSources(res)
        }
        return nil
    }
    res["phishConfidenceLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhishConfidenceLevel(val)
        }
        return nil
    }
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val)
        }
        return nil
    }
    res["policyAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyAction(val)
        }
        return nil
    }
    res["recipientEmailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRecipientEmailAddresses(res)
        }
        return nil
    }
    res["returnPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnPath(val)
        }
        return nil
    }
    res["senderDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAnalyzedEmailSenderDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderDetail(val.(AnalyzedEmailSenderDetailable))
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["spamConfidenceLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpamConfidenceLevel(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["threatType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatType(val.(*ThreatType))
        }
        return nil
    }
    res["urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAnalyzedEmailUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnalyzedEmailUrlable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AnalyzedEmailUrlable)
                }
            }
            m.SetUrls(res)
        }
        return nil
    }
    res["urlsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrlsCount(val)
        }
        return nil
    }
    return res
}
// GetInternetMessageId gets the internetMessageId property value. A public-facing identifier for the email that is sent. The message ID is in the format specified by RFC2822.
// returns a *string when successful
func (m *AnalyzedEmail) GetInternetMessageId()(*string) {
    val, err := m.GetBackingStore().Get("internetMessageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLanguage gets the language property value. The detected language of the email content.
// returns a *string when successful
func (m *AnalyzedEmail) GetLanguage()(*string) {
    val, err := m.GetBackingStore().Get("language")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLatestDelivery gets the latestDelivery property value. The latest delivery details of the email.
// returns a AnalyzedEmailDeliveryDetailable when successful
func (m *AnalyzedEmail) GetLatestDelivery()(AnalyzedEmailDeliveryDetailable) {
    val, err := m.GetBackingStore().Get("latestDelivery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AnalyzedEmailDeliveryDetailable)
    }
    return nil
}
// GetLoggedDateTime gets the loggedDateTime property value. Date-time when the email record was logged.
// returns a *Time when successful
func (m *AnalyzedEmail) GetLoggedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("loggedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNetworkMessageId gets the networkMessageId property value. An internal identifier for the email generated by Microsoft 365.
// returns a *string when successful
func (m *AnalyzedEmail) GetNetworkMessageId()(*string) {
    val, err := m.GetBackingStore().Get("networkMessageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginalDelivery gets the originalDelivery property value. The original delivery details of the email.
// returns a AnalyzedEmailDeliveryDetailable when successful
func (m *AnalyzedEmail) GetOriginalDelivery()(AnalyzedEmailDeliveryDetailable) {
    val, err := m.GetBackingStore().Get("originalDelivery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AnalyzedEmailDeliveryDetailable)
    }
    return nil
}
// GetOverrideSources gets the overrideSources property value. An aggregated list of all overrides with source on email.
// returns a []string when successful
func (m *AnalyzedEmail) GetOverrideSources()([]string) {
    val, err := m.GetBackingStore().Get("overrideSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPhishConfidenceLevel gets the phishConfidenceLevel property value. The phish confidence level associated with the email
// returns a *string when successful
func (m *AnalyzedEmail) GetPhishConfidenceLevel()(*string) {
    val, err := m.GetBackingStore().Get("phishConfidenceLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicy gets the policy property value. The action policy that took effect.
// returns a *string when successful
func (m *AnalyzedEmail) GetPolicy()(*string) {
    val, err := m.GetBackingStore().Get("policy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyAction gets the policyAction property value. The action taken on the email based on the configured policy.
// returns a *string when successful
func (m *AnalyzedEmail) GetPolicyAction()(*string) {
    val, err := m.GetBackingStore().Get("policyAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecipientEmailAddresses gets the recipientEmailAddresses property value. Contains the email addresses of the recipients.
// returns a []string when successful
func (m *AnalyzedEmail) GetRecipientEmailAddresses()([]string) {
    val, err := m.GetBackingStore().Get("recipientEmailAddresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetReturnPath gets the returnPath property value. A field that indicates where and how bounced emails are processed.
// returns a *string when successful
func (m *AnalyzedEmail) GetReturnPath()(*string) {
    val, err := m.GetBackingStore().Get("returnPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSenderDetail gets the senderDetail property value. Sender details of the email.
// returns a AnalyzedEmailSenderDetailable when successful
func (m *AnalyzedEmail) GetSenderDetail()(AnalyzedEmailSenderDetailable) {
    val, err := m.GetBackingStore().Get("senderDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AnalyzedEmailSenderDetailable)
    }
    return nil
}
// GetSizeInBytes gets the sizeInBytes property value. Size of the email in bytes.
// returns a *int32 when successful
func (m *AnalyzedEmail) GetSizeInBytes()(*int32) {
    val, err := m.GetBackingStore().Get("sizeInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSpamConfidenceLevel gets the spamConfidenceLevel property value. Spam confidence of the email.
// returns a *string when successful
func (m *AnalyzedEmail) GetSpamConfidenceLevel()(*string) {
    val, err := m.GetBackingStore().Get("spamConfidenceLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubject gets the subject property value. Subject of the email.
// returns a *string when successful
func (m *AnalyzedEmail) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreatType gets the threatType property value. Indicates the threat types. The possible values are: unknown, spam, malware, phishing, none, unknownFutureValue.
// returns a *ThreatType when successful
func (m *AnalyzedEmail) GetThreatType()(*ThreatType) {
    val, err := m.GetBackingStore().Get("threatType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ThreatType)
    }
    return nil
}
// GetUrls gets the urls property value. A collection of the URLs in the email.
// returns a []AnalyzedEmailUrlable when successful
func (m *AnalyzedEmail) GetUrls()([]AnalyzedEmailUrlable) {
    val, err := m.GetBackingStore().Get("urls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AnalyzedEmailUrlable)
    }
    return nil
}
// GetUrlsCount gets the urlsCount property value. The number of URLs in the email.
// returns a *int32 when successful
func (m *AnalyzedEmail) GetUrlsCount()(*int32) {
    val, err := m.GetBackingStore().Get("urlsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AnalyzedEmail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertIds() != nil {
        err = writer.WriteCollectionOfStringValues("alertIds", m.GetAlertIds())
        if err != nil {
            return err
        }
    }
    if m.GetAttachments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("attachmentsCount", m.GetAttachmentsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationDetails", m.GetAuthenticationDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bulkComplaintLevel", m.GetBulkComplaintLevel())
        if err != nil {
            return err
        }
    }
    if m.GetContexts() != nil {
        err = writer.WriteCollectionOfStringValues("contexts", m.GetContexts())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionMethods() != nil {
        err = writer.WriteCollectionOfStringValues("detectionMethods", m.GetDetectionMethods())
        if err != nil {
            return err
        }
    }
    if m.GetDirectionality() != nil {
        cast := (*m.GetDirectionality()).String()
        err = writer.WriteStringValue("directionality", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("distributionList", m.GetDistributionList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailClusterId", m.GetEmailClusterId())
        if err != nil {
            return err
        }
    }
    if m.GetExchangeTransportRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExchangeTransportRules()))
        for i, v := range m.GetExchangeTransportRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exchangeTransportRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("latestDelivery", m.GetLatestDelivery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("loggedDateTime", m.GetLoggedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkMessageId", m.GetNetworkMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("originalDelivery", m.GetOriginalDelivery())
        if err != nil {
            return err
        }
    }
    if m.GetOverrideSources() != nil {
        err = writer.WriteCollectionOfStringValues("overrideSources", m.GetOverrideSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phishConfidenceLevel", m.GetPhishConfidenceLevel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyAction", m.GetPolicyAction())
        if err != nil {
            return err
        }
    }
    if m.GetRecipientEmailAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("recipientEmailAddresses", m.GetRecipientEmailAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("returnPath", m.GetReturnPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("senderDetail", m.GetSenderDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("spamConfidenceLevel", m.GetSpamConfidenceLevel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetThreatType() != nil {
        cast := (*m.GetThreatType()).String()
        err = writer.WriteStringValue("threatType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUrls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUrls()))
        for i, v := range m.GetUrls() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("urls", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("urlsCount", m.GetUrlsCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertIds sets the alertIds property value. A collection of values that contain the IDs of any alerts associated with the email.
func (m *AnalyzedEmail) SetAlertIds(value []string)() {
    err := m.GetBackingStore().Set("alertIds", value)
    if err != nil {
        panic(err)
    }
}
// SetAttachments sets the attachments property value. A collection of the attachments in the email.
func (m *AnalyzedEmail) SetAttachments(value []AnalyzedEmailAttachmentable)() {
    err := m.GetBackingStore().Set("attachments", value)
    if err != nil {
        panic(err)
    }
}
// SetAttachmentsCount sets the attachmentsCount property value. The number of attachments in the email.
func (m *AnalyzedEmail) SetAttachmentsCount(value *int32)() {
    err := m.GetBackingStore().Set("attachmentsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationDetails sets the authenticationDetails property value. The authentication details associated with the email.
func (m *AnalyzedEmail) SetAuthenticationDetails(value AnalyzedEmailAuthenticationDetailable)() {
    err := m.GetBackingStore().Set("authenticationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetBulkComplaintLevel sets the bulkComplaintLevel property value. The bulk complaint level of the email. A higher level is more likely to be spam.
func (m *AnalyzedEmail) SetBulkComplaintLevel(value *string)() {
    err := m.GetBackingStore().Set("bulkComplaintLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetContexts sets the contexts property value. Provides context of the email.
func (m *AnalyzedEmail) SetContexts(value []string)() {
    err := m.GetBackingStore().Set("contexts", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionMethods sets the detectionMethods property value. The methods of detection used.
func (m *AnalyzedEmail) SetDetectionMethods(value []string)() {
    err := m.GetBackingStore().Set("detectionMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetDirectionality sets the directionality property value. The direction of the emails. The possible values are: unknown, inbound, outbound, intraOrg, unknownFutureValue.
func (m *AnalyzedEmail) SetDirectionality(value *AntispamDirectionality)() {
    err := m.GetBackingStore().Set("directionality", value)
    if err != nil {
        panic(err)
    }
}
// SetDistributionList sets the distributionList property value. The distribution list details to which the email was sent.
func (m *AnalyzedEmail) SetDistributionList(value *string)() {
    err := m.GetBackingStore().Set("distributionList", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailClusterId sets the emailClusterId property value. The identifier for the group of similar emails clustered based on heuristic analysis of their content.
func (m *AnalyzedEmail) SetEmailClusterId(value *string)() {
    err := m.GetBackingStore().Set("emailClusterId", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeTransportRules sets the exchangeTransportRules property value. The name of the Exchange transport rules (ETRs) associated with the email.
func (m *AnalyzedEmail) SetExchangeTransportRules(value []AnalyzedEmailExchangeTransportRuleInfoable)() {
    err := m.GetBackingStore().Set("exchangeTransportRules", value)
    if err != nil {
        panic(err)
    }
}
// SetInternetMessageId sets the internetMessageId property value. A public-facing identifier for the email that is sent. The message ID is in the format specified by RFC2822.
func (m *AnalyzedEmail) SetInternetMessageId(value *string)() {
    err := m.GetBackingStore().Set("internetMessageId", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguage sets the language property value. The detected language of the email content.
func (m *AnalyzedEmail) SetLanguage(value *string)() {
    err := m.GetBackingStore().Set("language", value)
    if err != nil {
        panic(err)
    }
}
// SetLatestDelivery sets the latestDelivery property value. The latest delivery details of the email.
func (m *AnalyzedEmail) SetLatestDelivery(value AnalyzedEmailDeliveryDetailable)() {
    err := m.GetBackingStore().Set("latestDelivery", value)
    if err != nil {
        panic(err)
    }
}
// SetLoggedDateTime sets the loggedDateTime property value. Date-time when the email record was logged.
func (m *AnalyzedEmail) SetLoggedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("loggedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkMessageId sets the networkMessageId property value. An internal identifier for the email generated by Microsoft 365.
func (m *AnalyzedEmail) SetNetworkMessageId(value *string)() {
    err := m.GetBackingStore().Set("networkMessageId", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginalDelivery sets the originalDelivery property value. The original delivery details of the email.
func (m *AnalyzedEmail) SetOriginalDelivery(value AnalyzedEmailDeliveryDetailable)() {
    err := m.GetBackingStore().Set("originalDelivery", value)
    if err != nil {
        panic(err)
    }
}
// SetOverrideSources sets the overrideSources property value. An aggregated list of all overrides with source on email.
func (m *AnalyzedEmail) SetOverrideSources(value []string)() {
    err := m.GetBackingStore().Set("overrideSources", value)
    if err != nil {
        panic(err)
    }
}
// SetPhishConfidenceLevel sets the phishConfidenceLevel property value. The phish confidence level associated with the email
func (m *AnalyzedEmail) SetPhishConfidenceLevel(value *string)() {
    err := m.GetBackingStore().Set("phishConfidenceLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicy sets the policy property value. The action policy that took effect.
func (m *AnalyzedEmail) SetPolicy(value *string)() {
    err := m.GetBackingStore().Set("policy", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyAction sets the policyAction property value. The action taken on the email based on the configured policy.
func (m *AnalyzedEmail) SetPolicyAction(value *string)() {
    err := m.GetBackingStore().Set("policyAction", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipientEmailAddresses sets the recipientEmailAddresses property value. Contains the email addresses of the recipients.
func (m *AnalyzedEmail) SetRecipientEmailAddresses(value []string)() {
    err := m.GetBackingStore().Set("recipientEmailAddresses", value)
    if err != nil {
        panic(err)
    }
}
// SetReturnPath sets the returnPath property value. A field that indicates where and how bounced emails are processed.
func (m *AnalyzedEmail) SetReturnPath(value *string)() {
    err := m.GetBackingStore().Set("returnPath", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderDetail sets the senderDetail property value. Sender details of the email.
func (m *AnalyzedEmail) SetSenderDetail(value AnalyzedEmailSenderDetailable)() {
    err := m.GetBackingStore().Set("senderDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeInBytes sets the sizeInBytes property value. Size of the email in bytes.
func (m *AnalyzedEmail) SetSizeInBytes(value *int32)() {
    err := m.GetBackingStore().Set("sizeInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetSpamConfidenceLevel sets the spamConfidenceLevel property value. Spam confidence of the email.
func (m *AnalyzedEmail) SetSpamConfidenceLevel(value *string)() {
    err := m.GetBackingStore().Set("spamConfidenceLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. Subject of the email.
func (m *AnalyzedEmail) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatType sets the threatType property value. Indicates the threat types. The possible values are: unknown, spam, malware, phishing, none, unknownFutureValue.
func (m *AnalyzedEmail) SetThreatType(value *ThreatType)() {
    err := m.GetBackingStore().Set("threatType", value)
    if err != nil {
        panic(err)
    }
}
// SetUrls sets the urls property value. A collection of the URLs in the email.
func (m *AnalyzedEmail) SetUrls(value []AnalyzedEmailUrlable)() {
    err := m.GetBackingStore().Set("urls", value)
    if err != nil {
        panic(err)
    }
}
// SetUrlsCount sets the urlsCount property value. The number of URLs in the email.
func (m *AnalyzedEmail) SetUrlsCount(value *int32)() {
    err := m.GetBackingStore().Set("urlsCount", value)
    if err != nil {
        panic(err)
    }
}
type AnalyzedEmailable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertIds()([]string)
    GetAttachments()([]AnalyzedEmailAttachmentable)
    GetAttachmentsCount()(*int32)
    GetAuthenticationDetails()(AnalyzedEmailAuthenticationDetailable)
    GetBulkComplaintLevel()(*string)
    GetContexts()([]string)
    GetDetectionMethods()([]string)
    GetDirectionality()(*AntispamDirectionality)
    GetDistributionList()(*string)
    GetEmailClusterId()(*string)
    GetExchangeTransportRules()([]AnalyzedEmailExchangeTransportRuleInfoable)
    GetInternetMessageId()(*string)
    GetLanguage()(*string)
    GetLatestDelivery()(AnalyzedEmailDeliveryDetailable)
    GetLoggedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNetworkMessageId()(*string)
    GetOriginalDelivery()(AnalyzedEmailDeliveryDetailable)
    GetOverrideSources()([]string)
    GetPhishConfidenceLevel()(*string)
    GetPolicy()(*string)
    GetPolicyAction()(*string)
    GetRecipientEmailAddresses()([]string)
    GetReturnPath()(*string)
    GetSenderDetail()(AnalyzedEmailSenderDetailable)
    GetSizeInBytes()(*int32)
    GetSpamConfidenceLevel()(*string)
    GetSubject()(*string)
    GetThreatType()(*ThreatType)
    GetUrls()([]AnalyzedEmailUrlable)
    GetUrlsCount()(*int32)
    SetAlertIds(value []string)()
    SetAttachments(value []AnalyzedEmailAttachmentable)()
    SetAttachmentsCount(value *int32)()
    SetAuthenticationDetails(value AnalyzedEmailAuthenticationDetailable)()
    SetBulkComplaintLevel(value *string)()
    SetContexts(value []string)()
    SetDetectionMethods(value []string)()
    SetDirectionality(value *AntispamDirectionality)()
    SetDistributionList(value *string)()
    SetEmailClusterId(value *string)()
    SetExchangeTransportRules(value []AnalyzedEmailExchangeTransportRuleInfoable)()
    SetInternetMessageId(value *string)()
    SetLanguage(value *string)()
    SetLatestDelivery(value AnalyzedEmailDeliveryDetailable)()
    SetLoggedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNetworkMessageId(value *string)()
    SetOriginalDelivery(value AnalyzedEmailDeliveryDetailable)()
    SetOverrideSources(value []string)()
    SetPhishConfidenceLevel(value *string)()
    SetPolicy(value *string)()
    SetPolicyAction(value *string)()
    SetRecipientEmailAddresses(value []string)()
    SetReturnPath(value *string)()
    SetSenderDetail(value AnalyzedEmailSenderDetailable)()
    SetSizeInBytes(value *int32)()
    SetSpamConfidenceLevel(value *string)()
    SetSubject(value *string)()
    SetThreatType(value *ThreatType)()
    SetUrls(value []AnalyzedEmailUrlable)()
    SetUrlsCount(value *int32)()
}
