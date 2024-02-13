package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TiIndicator struct {
    Entity
}
// NewTiIndicator instantiates a new TiIndicator and sets the default values.
func NewTiIndicator()(*TiIndicator) {
    m := &TiIndicator{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTiIndicatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTiIndicatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicator(), nil
}
// GetAction gets the action property value. The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
// returns a *TiAction when successful
func (m *TiIndicator) GetAction()(*TiAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TiAction)
    }
    return nil
}
// GetActivityGroupNames gets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
// returns a []string when successful
func (m *TiIndicator) GetActivityGroupNames()([]string) {
    val, err := m.GetBackingStore().Get("activityGroupNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdditionalInformation gets the additionalInformation property value. A catchall area for extra data from the indicator that is not specifically covered by other tiIndicator properties. The security tool specified by targetProduct typically does not utilize this data.
// returns a *string when successful
func (m *TiIndicator) GetAdditionalInformation()(*string) {
    val, err := m.GetBackingStore().Get("additionalInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureTenantId gets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Microsoft Entra tenant id of submitting client. Required.
// returns a *string when successful
func (m *TiIndicator) GetAzureTenantId()(*string) {
    val, err := m.GetBackingStore().Get("azureTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfidence gets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
// returns a *int32 when successful
func (m *TiIndicator) GetConfidence()(*int32) {
    val, err := m.GetBackingStore().Get("confidence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDescription gets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
// returns a *string when successful
func (m *TiIndicator) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiamondModel gets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
// returns a *DiamondModel when successful
func (m *TiIndicator) GetDiamondModel()(*DiamondModel) {
    val, err := m.GetBackingStore().Get("diamondModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DiamondModel)
    }
    return nil
}
// GetDomainName gets the domainName property value. The domainName property
// returns a *string when successful
func (m *TiIndicator) GetDomainName()(*string) {
    val, err := m.GetBackingStore().Get("domainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailEncoding gets the emailEncoding property value. The emailEncoding property
// returns a *string when successful
func (m *TiIndicator) GetEmailEncoding()(*string) {
    val, err := m.GetBackingStore().Get("emailEncoding")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailLanguage gets the emailLanguage property value. The emailLanguage property
// returns a *string when successful
func (m *TiIndicator) GetEmailLanguage()(*string) {
    val, err := m.GetBackingStore().Get("emailLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailRecipient gets the emailRecipient property value. The emailRecipient property
// returns a *string when successful
func (m *TiIndicator) GetEmailRecipient()(*string) {
    val, err := m.GetBackingStore().Get("emailRecipient")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailSenderAddress gets the emailSenderAddress property value. The emailSenderAddress property
// returns a *string when successful
func (m *TiIndicator) GetEmailSenderAddress()(*string) {
    val, err := m.GetBackingStore().Get("emailSenderAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailSenderName gets the emailSenderName property value. The emailSenderName property
// returns a *string when successful
func (m *TiIndicator) GetEmailSenderName()(*string) {
    val, err := m.GetBackingStore().Get("emailSenderName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailSourceDomain gets the emailSourceDomain property value. The emailSourceDomain property
// returns a *string when successful
func (m *TiIndicator) GetEmailSourceDomain()(*string) {
    val, err := m.GetBackingStore().Get("emailSourceDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailSourceIpAddress gets the emailSourceIpAddress property value. The emailSourceIpAddress property
// returns a *string when successful
func (m *TiIndicator) GetEmailSourceIpAddress()(*string) {
    val, err := m.GetBackingStore().Get("emailSourceIpAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailSubject gets the emailSubject property value. The emailSubject property
// returns a *string when successful
func (m *TiIndicator) GetEmailSubject()(*string) {
    val, err := m.GetBackingStore().Get("emailSubject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailXMailer gets the emailXMailer property value. The emailXMailer property
// returns a *string when successful
func (m *TiIndicator) GetEmailXMailer()(*string) {
    val, err := m.GetBackingStore().Get("emailXMailer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
// returns a *Time when successful
func (m *TiIndicator) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExternalId gets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (for example, a foreign key).
// returns a *string when successful
func (m *TiIndicator) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TiIndicator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTiAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*TiAction))
        }
        return nil
    }
    res["activityGroupNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetActivityGroupNames(res)
        }
        return nil
    }
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["confidence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["diamondModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiamondModel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiamondModel(val.(*DiamondModel))
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["emailEncoding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailEncoding(val)
        }
        return nil
    }
    res["emailLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailLanguage(val)
        }
        return nil
    }
    res["emailRecipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailRecipient(val)
        }
        return nil
    }
    res["emailSenderAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSenderAddress(val)
        }
        return nil
    }
    res["emailSenderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSenderName(val)
        }
        return nil
    }
    res["emailSourceDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSourceDomain(val)
        }
        return nil
    }
    res["emailSourceIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSourceIpAddress(val)
        }
        return nil
    }
    res["emailSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSubject(val)
        }
        return nil
    }
    res["emailXMailer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailXMailer(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["fileCompileDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileCompileDateTime(val)
        }
        return nil
    }
    res["fileCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileCreatedDateTime(val)
        }
        return nil
    }
    res["fileHashType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileHashType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHashType(val.(*FileHashType))
        }
        return nil
    }
    res["fileHashValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHashValue(val)
        }
        return nil
    }
    res["fileMutexName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileMutexName(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["filePacker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePacker(val)
        }
        return nil
    }
    res["filePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
        }
        return nil
    }
    res["fileSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["fileType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileType(val)
        }
        return nil
    }
    res["ingestedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIngestedDateTime(val)
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["killChain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetKillChain(res)
        }
        return nil
    }
    res["knownFalsePositives"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKnownFalsePositives(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["malwareFamilyNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMalwareFamilyNames(res)
        }
        return nil
    }
    res["networkCidrBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkCidrBlock(val)
        }
        return nil
    }
    res["networkDestinationAsn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationAsn(val)
        }
        return nil
    }
    res["networkDestinationCidrBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationCidrBlock(val)
        }
        return nil
    }
    res["networkDestinationIPv4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationIPv4(val)
        }
        return nil
    }
    res["networkDestinationIPv6"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationIPv6(val)
        }
        return nil
    }
    res["networkDestinationPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationPort(val)
        }
        return nil
    }
    res["networkIPv4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkIPv4(val)
        }
        return nil
    }
    res["networkIPv6"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkIPv6(val)
        }
        return nil
    }
    res["networkPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkPort(val)
        }
        return nil
    }
    res["networkProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProtocol(val)
        }
        return nil
    }
    res["networkSourceAsn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceAsn(val)
        }
        return nil
    }
    res["networkSourceCidrBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceCidrBlock(val)
        }
        return nil
    }
    res["networkSourceIPv4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceIPv4(val)
        }
        return nil
    }
    res["networkSourceIPv6"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceIPv6(val)
        }
        return nil
    }
    res["networkSourcePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourcePort(val)
        }
        return nil
    }
    res["passiveOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassiveOnly(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTags(res)
        }
        return nil
    }
    res["targetProduct"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetProduct(val)
        }
        return nil
    }
    res["threatType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatType(val)
        }
        return nil
    }
    res["tlpLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTlpLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlpLevel(val.(*TlpLevel))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["userAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val)
        }
        return nil
    }
    return res
}
// GetFileCompileDateTime gets the fileCompileDateTime property value. The fileCompileDateTime property
// returns a *Time when successful
func (m *TiIndicator) GetFileCompileDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("fileCompileDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFileCreatedDateTime gets the fileCreatedDateTime property value. The fileCreatedDateTime property
// returns a *Time when successful
func (m *TiIndicator) GetFileCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("fileCreatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFileHashType gets the fileHashType property value. The fileHashType property
// returns a *FileHashType when successful
func (m *TiIndicator) GetFileHashType()(*FileHashType) {
    val, err := m.GetBackingStore().Get("fileHashType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FileHashType)
    }
    return nil
}
// GetFileHashValue gets the fileHashValue property value. The fileHashValue property
// returns a *string when successful
func (m *TiIndicator) GetFileHashValue()(*string) {
    val, err := m.GetBackingStore().Get("fileHashValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileMutexName gets the fileMutexName property value. The fileMutexName property
// returns a *string when successful
func (m *TiIndicator) GetFileMutexName()(*string) {
    val, err := m.GetBackingStore().Get("fileMutexName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileName gets the fileName property value. The fileName property
// returns a *string when successful
func (m *TiIndicator) GetFileName()(*string) {
    val, err := m.GetBackingStore().Get("fileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFilePacker gets the filePacker property value. The filePacker property
// returns a *string when successful
func (m *TiIndicator) GetFilePacker()(*string) {
    val, err := m.GetBackingStore().Get("filePacker")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFilePath gets the filePath property value. The filePath property
// returns a *string when successful
func (m *TiIndicator) GetFilePath()(*string) {
    val, err := m.GetBackingStore().Get("filePath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileSize gets the fileSize property value. The fileSize property
// returns a *int64 when successful
func (m *TiIndicator) GetFileSize()(*int64) {
    val, err := m.GetBackingStore().Get("fileSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFileType gets the fileType property value. The fileType property
// returns a *string when successful
func (m *TiIndicator) GetFileType()(*string) {
    val, err := m.GetBackingStore().Get("fileType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIngestedDateTime gets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *TiIndicator) GetIngestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("ingestedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetIsActive gets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
// returns a *bool when successful
func (m *TiIndicator) GetIsActive()(*bool) {
    val, err := m.GetBackingStore().Get("isActive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKillChain gets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
// returns a []string when successful
func (m *TiIndicator) GetKillChain()([]string) {
    val, err := m.GetBackingStore().Get("killChain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKnownFalsePositives gets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
// returns a *string when successful
func (m *TiIndicator) GetKnownFalsePositives()(*string) {
    val, err := m.GetBackingStore().Get("knownFalsePositives")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *TiIndicator) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastReportedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMalwareFamilyNames gets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible that can be found via the Windows Defender Security Intelligence threat encyclopedia.
// returns a []string when successful
func (m *TiIndicator) GetMalwareFamilyNames()([]string) {
    val, err := m.GetBackingStore().Get("malwareFamilyNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetNetworkCidrBlock gets the networkCidrBlock property value. The networkCidrBlock property
// returns a *string when successful
func (m *TiIndicator) GetNetworkCidrBlock()(*string) {
    val, err := m.GetBackingStore().Get("networkCidrBlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkDestinationAsn gets the networkDestinationAsn property value. The networkDestinationAsn property
// returns a *int64 when successful
func (m *TiIndicator) GetNetworkDestinationAsn()(*int64) {
    val, err := m.GetBackingStore().Get("networkDestinationAsn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetNetworkDestinationCidrBlock gets the networkDestinationCidrBlock property value. The networkDestinationCidrBlock property
// returns a *string when successful
func (m *TiIndicator) GetNetworkDestinationCidrBlock()(*string) {
    val, err := m.GetBackingStore().Get("networkDestinationCidrBlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkDestinationIPv4 gets the networkDestinationIPv4 property value. The networkDestinationIPv4 property
// returns a *string when successful
func (m *TiIndicator) GetNetworkDestinationIPv4()(*string) {
    val, err := m.GetBackingStore().Get("networkDestinationIPv4")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkDestinationIPv6 gets the networkDestinationIPv6 property value. The networkDestinationIPv6 property
// returns a *string when successful
func (m *TiIndicator) GetNetworkDestinationIPv6()(*string) {
    val, err := m.GetBackingStore().Get("networkDestinationIPv6")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkDestinationPort gets the networkDestinationPort property value. The networkDestinationPort property
// returns a *int32 when successful
func (m *TiIndicator) GetNetworkDestinationPort()(*int32) {
    val, err := m.GetBackingStore().Get("networkDestinationPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNetworkIPv4 gets the networkIPv4 property value. The networkIPv4 property
// returns a *string when successful
func (m *TiIndicator) GetNetworkIPv4()(*string) {
    val, err := m.GetBackingStore().Get("networkIPv4")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkIPv6 gets the networkIPv6 property value. The networkIPv6 property
// returns a *string when successful
func (m *TiIndicator) GetNetworkIPv6()(*string) {
    val, err := m.GetBackingStore().Get("networkIPv6")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkPort gets the networkPort property value. The networkPort property
// returns a *int32 when successful
func (m *TiIndicator) GetNetworkPort()(*int32) {
    val, err := m.GetBackingStore().Get("networkPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNetworkProtocol gets the networkProtocol property value. The networkProtocol property
// returns a *int32 when successful
func (m *TiIndicator) GetNetworkProtocol()(*int32) {
    val, err := m.GetBackingStore().Get("networkProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNetworkSourceAsn gets the networkSourceAsn property value. The networkSourceAsn property
// returns a *int64 when successful
func (m *TiIndicator) GetNetworkSourceAsn()(*int64) {
    val, err := m.GetBackingStore().Get("networkSourceAsn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetNetworkSourceCidrBlock gets the networkSourceCidrBlock property value. The networkSourceCidrBlock property
// returns a *string when successful
func (m *TiIndicator) GetNetworkSourceCidrBlock()(*string) {
    val, err := m.GetBackingStore().Get("networkSourceCidrBlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkSourceIPv4 gets the networkSourceIPv4 property value. The networkSourceIPv4 property
// returns a *string when successful
func (m *TiIndicator) GetNetworkSourceIPv4()(*string) {
    val, err := m.GetBackingStore().Get("networkSourceIPv4")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkSourceIPv6 gets the networkSourceIPv6 property value. The networkSourceIPv6 property
// returns a *string when successful
func (m *TiIndicator) GetNetworkSourceIPv6()(*string) {
    val, err := m.GetBackingStore().Get("networkSourceIPv6")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkSourcePort gets the networkSourcePort property value. The networkSourcePort property
// returns a *int32 when successful
func (m *TiIndicator) GetNetworkSourcePort()(*int32) {
    val, err := m.GetBackingStore().Get("networkSourcePort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPassiveOnly gets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools won't notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they'll simply log that a match occurred but won't perform the action. Default value is false.
// returns a *bool when successful
func (m *TiIndicator) GetPassiveOnly()(*bool) {
    val, err := m.GetBackingStore().Get("passiveOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSeverity gets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero isn't severe at all. Default value is 3.
// returns a *int32 when successful
func (m *TiIndicator) GetSeverity()(*int32) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTags gets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
// returns a []string when successful
func (m *TiIndicator) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTargetProduct gets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
// returns a *string when successful
func (m *TiIndicator) GetTargetProduct()(*string) {
    val, err := m.GetBackingStore().Get("targetProduct")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreatType gets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
// returns a *string when successful
func (m *TiIndicator) GetThreatType()(*string) {
    val, err := m.GetBackingStore().Get("threatType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTlpLevel gets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
// returns a *TlpLevel when successful
func (m *TiIndicator) GetTlpLevel()(*TlpLevel) {
    val, err := m.GetBackingStore().Get("tlpLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TlpLevel)
    }
    return nil
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TiIndicator) GetUrl()(*string) {
    val, err := m.GetBackingStore().Get("url")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserAgent gets the userAgent property value. The userAgent property
// returns a *string when successful
func (m *TiIndicator) GetUserAgent()(*string) {
    val, err := m.GetBackingStore().Get("userAgent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TiIndicator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActivityGroupNames() != nil {
        err = writer.WriteCollectionOfStringValues("activityGroupNames", m.GetActivityGroupNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDiamondModel() != nil {
        cast := (*m.GetDiamondModel()).String()
        err = writer.WriteStringValue("diamondModel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailEncoding", m.GetEmailEncoding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailLanguage", m.GetEmailLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailRecipient", m.GetEmailRecipient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailSenderAddress", m.GetEmailSenderAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailSenderName", m.GetEmailSenderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailSourceDomain", m.GetEmailSourceDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailSourceIpAddress", m.GetEmailSourceIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailSubject", m.GetEmailSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailXMailer", m.GetEmailXMailer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("fileCompileDateTime", m.GetFileCompileDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("fileCreatedDateTime", m.GetFileCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetFileHashType() != nil {
        cast := (*m.GetFileHashType()).String()
        err = writer.WriteStringValue("fileHashType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileHashValue", m.GetFileHashValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileMutexName", m.GetFileMutexName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filePacker", m.GetFilePacker())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filePath", m.GetFilePath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("fileSize", m.GetFileSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileType", m.GetFileType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("ingestedDateTime", m.GetIngestedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    if m.GetKillChain() != nil {
        err = writer.WriteCollectionOfStringValues("killChain", m.GetKillChain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("knownFalsePositives", m.GetKnownFalsePositives())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMalwareFamilyNames() != nil {
        err = writer.WriteCollectionOfStringValues("malwareFamilyNames", m.GetMalwareFamilyNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkCidrBlock", m.GetNetworkCidrBlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("networkDestinationAsn", m.GetNetworkDestinationAsn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkDestinationCidrBlock", m.GetNetworkDestinationCidrBlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkDestinationIPv4", m.GetNetworkDestinationIPv4())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkDestinationIPv6", m.GetNetworkDestinationIPv6())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("networkDestinationPort", m.GetNetworkDestinationPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkIPv4", m.GetNetworkIPv4())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkIPv6", m.GetNetworkIPv6())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("networkPort", m.GetNetworkPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("networkProtocol", m.GetNetworkProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("networkSourceAsn", m.GetNetworkSourceAsn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkSourceCidrBlock", m.GetNetworkSourceCidrBlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkSourceIPv4", m.GetNetworkSourceIPv4())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkSourceIPv6", m.GetNetworkSourceIPv6())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("networkSourcePort", m.GetNetworkSourcePort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passiveOnly", m.GetPassiveOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("severity", m.GetSeverity())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetProduct", m.GetTargetProduct())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("threatType", m.GetThreatType())
        if err != nil {
            return err
        }
    }
    if m.GetTlpLevel() != nil {
        cast := (*m.GetTlpLevel()).String()
        err = writer.WriteStringValue("tlpLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userAgent", m.GetUserAgent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
func (m *TiIndicator) SetAction(value *TiAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityGroupNames sets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
func (m *TiIndicator) SetActivityGroupNames(value []string)() {
    err := m.GetBackingStore().Set("activityGroupNames", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalInformation sets the additionalInformation property value. A catchall area for extra data from the indicator that is not specifically covered by other tiIndicator properties. The security tool specified by targetProduct typically does not utilize this data.
func (m *TiIndicator) SetAdditionalInformation(value *string)() {
    err := m.GetBackingStore().Set("additionalInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureTenantId sets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Microsoft Entra tenant id of submitting client. Required.
func (m *TiIndicator) SetAzureTenantId(value *string)() {
    err := m.GetBackingStore().Set("azureTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetConfidence sets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
func (m *TiIndicator) SetConfidence(value *int32)() {
    err := m.GetBackingStore().Set("confidence", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
func (m *TiIndicator) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDiamondModel sets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
func (m *TiIndicator) SetDiamondModel(value *DiamondModel)() {
    err := m.GetBackingStore().Set("diamondModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainName sets the domainName property value. The domainName property
func (m *TiIndicator) SetDomainName(value *string)() {
    err := m.GetBackingStore().Set("domainName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailEncoding sets the emailEncoding property value. The emailEncoding property
func (m *TiIndicator) SetEmailEncoding(value *string)() {
    err := m.GetBackingStore().Set("emailEncoding", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailLanguage sets the emailLanguage property value. The emailLanguage property
func (m *TiIndicator) SetEmailLanguage(value *string)() {
    err := m.GetBackingStore().Set("emailLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailRecipient sets the emailRecipient property value. The emailRecipient property
func (m *TiIndicator) SetEmailRecipient(value *string)() {
    err := m.GetBackingStore().Set("emailRecipient", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSenderAddress sets the emailSenderAddress property value. The emailSenderAddress property
func (m *TiIndicator) SetEmailSenderAddress(value *string)() {
    err := m.GetBackingStore().Set("emailSenderAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSenderName sets the emailSenderName property value. The emailSenderName property
func (m *TiIndicator) SetEmailSenderName(value *string)() {
    err := m.GetBackingStore().Set("emailSenderName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSourceDomain sets the emailSourceDomain property value. The emailSourceDomain property
func (m *TiIndicator) SetEmailSourceDomain(value *string)() {
    err := m.GetBackingStore().Set("emailSourceDomain", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSourceIpAddress sets the emailSourceIpAddress property value. The emailSourceIpAddress property
func (m *TiIndicator) SetEmailSourceIpAddress(value *string)() {
    err := m.GetBackingStore().Set("emailSourceIpAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailSubject sets the emailSubject property value. The emailSubject property
func (m *TiIndicator) SetEmailSubject(value *string)() {
    err := m.GetBackingStore().Set("emailSubject", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailXMailer sets the emailXMailer property value. The emailXMailer property
func (m *TiIndicator) SetEmailXMailer(value *string)() {
    err := m.GetBackingStore().Set("emailXMailer", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
func (m *TiIndicator) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (for example, a foreign key).
func (m *TiIndicator) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetFileCompileDateTime sets the fileCompileDateTime property value. The fileCompileDateTime property
func (m *TiIndicator) SetFileCompileDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("fileCompileDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFileCreatedDateTime sets the fileCreatedDateTime property value. The fileCreatedDateTime property
func (m *TiIndicator) SetFileCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("fileCreatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFileHashType sets the fileHashType property value. The fileHashType property
func (m *TiIndicator) SetFileHashType(value *FileHashType)() {
    err := m.GetBackingStore().Set("fileHashType", value)
    if err != nil {
        panic(err)
    }
}
// SetFileHashValue sets the fileHashValue property value. The fileHashValue property
func (m *TiIndicator) SetFileHashValue(value *string)() {
    err := m.GetBackingStore().Set("fileHashValue", value)
    if err != nil {
        panic(err)
    }
}
// SetFileMutexName sets the fileMutexName property value. The fileMutexName property
func (m *TiIndicator) SetFileMutexName(value *string)() {
    err := m.GetBackingStore().Set("fileMutexName", value)
    if err != nil {
        panic(err)
    }
}
// SetFileName sets the fileName property value. The fileName property
func (m *TiIndicator) SetFileName(value *string)() {
    err := m.GetBackingStore().Set("fileName", value)
    if err != nil {
        panic(err)
    }
}
// SetFilePacker sets the filePacker property value. The filePacker property
func (m *TiIndicator) SetFilePacker(value *string)() {
    err := m.GetBackingStore().Set("filePacker", value)
    if err != nil {
        panic(err)
    }
}
// SetFilePath sets the filePath property value. The filePath property
func (m *TiIndicator) SetFilePath(value *string)() {
    err := m.GetBackingStore().Set("filePath", value)
    if err != nil {
        panic(err)
    }
}
// SetFileSize sets the fileSize property value. The fileSize property
func (m *TiIndicator) SetFileSize(value *int64)() {
    err := m.GetBackingStore().Set("fileSize", value)
    if err != nil {
        panic(err)
    }
}
// SetFileType sets the fileType property value. The fileType property
func (m *TiIndicator) SetFileType(value *string)() {
    err := m.GetBackingStore().Set("fileType", value)
    if err != nil {
        panic(err)
    }
}
// SetIngestedDateTime sets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) SetIngestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("ingestedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsActive sets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
func (m *TiIndicator) SetIsActive(value *bool)() {
    err := m.GetBackingStore().Set("isActive", value)
    if err != nil {
        panic(err)
    }
}
// SetKillChain sets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
func (m *TiIndicator) SetKillChain(value []string)() {
    err := m.GetBackingStore().Set("killChain", value)
    if err != nil {
        panic(err)
    }
}
// SetKnownFalsePositives sets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
func (m *TiIndicator) SetKnownFalsePositives(value *string)() {
    err := m.GetBackingStore().Set("knownFalsePositives", value)
    if err != nil {
        panic(err)
    }
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastReportedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMalwareFamilyNames sets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible that can be found via the Windows Defender Security Intelligence threat encyclopedia.
func (m *TiIndicator) SetMalwareFamilyNames(value []string)() {
    err := m.GetBackingStore().Set("malwareFamilyNames", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkCidrBlock sets the networkCidrBlock property value. The networkCidrBlock property
func (m *TiIndicator) SetNetworkCidrBlock(value *string)() {
    err := m.GetBackingStore().Set("networkCidrBlock", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkDestinationAsn sets the networkDestinationAsn property value. The networkDestinationAsn property
func (m *TiIndicator) SetNetworkDestinationAsn(value *int64)() {
    err := m.GetBackingStore().Set("networkDestinationAsn", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkDestinationCidrBlock sets the networkDestinationCidrBlock property value. The networkDestinationCidrBlock property
func (m *TiIndicator) SetNetworkDestinationCidrBlock(value *string)() {
    err := m.GetBackingStore().Set("networkDestinationCidrBlock", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkDestinationIPv4 sets the networkDestinationIPv4 property value. The networkDestinationIPv4 property
func (m *TiIndicator) SetNetworkDestinationIPv4(value *string)() {
    err := m.GetBackingStore().Set("networkDestinationIPv4", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkDestinationIPv6 sets the networkDestinationIPv6 property value. The networkDestinationIPv6 property
func (m *TiIndicator) SetNetworkDestinationIPv6(value *string)() {
    err := m.GetBackingStore().Set("networkDestinationIPv6", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkDestinationPort sets the networkDestinationPort property value. The networkDestinationPort property
func (m *TiIndicator) SetNetworkDestinationPort(value *int32)() {
    err := m.GetBackingStore().Set("networkDestinationPort", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkIPv4 sets the networkIPv4 property value. The networkIPv4 property
func (m *TiIndicator) SetNetworkIPv4(value *string)() {
    err := m.GetBackingStore().Set("networkIPv4", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkIPv6 sets the networkIPv6 property value. The networkIPv6 property
func (m *TiIndicator) SetNetworkIPv6(value *string)() {
    err := m.GetBackingStore().Set("networkIPv6", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkPort sets the networkPort property value. The networkPort property
func (m *TiIndicator) SetNetworkPort(value *int32)() {
    err := m.GetBackingStore().Set("networkPort", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkProtocol sets the networkProtocol property value. The networkProtocol property
func (m *TiIndicator) SetNetworkProtocol(value *int32)() {
    err := m.GetBackingStore().Set("networkProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkSourceAsn sets the networkSourceAsn property value. The networkSourceAsn property
func (m *TiIndicator) SetNetworkSourceAsn(value *int64)() {
    err := m.GetBackingStore().Set("networkSourceAsn", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkSourceCidrBlock sets the networkSourceCidrBlock property value. The networkSourceCidrBlock property
func (m *TiIndicator) SetNetworkSourceCidrBlock(value *string)() {
    err := m.GetBackingStore().Set("networkSourceCidrBlock", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkSourceIPv4 sets the networkSourceIPv4 property value. The networkSourceIPv4 property
func (m *TiIndicator) SetNetworkSourceIPv4(value *string)() {
    err := m.GetBackingStore().Set("networkSourceIPv4", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkSourceIPv6 sets the networkSourceIPv6 property value. The networkSourceIPv6 property
func (m *TiIndicator) SetNetworkSourceIPv6(value *string)() {
    err := m.GetBackingStore().Set("networkSourceIPv6", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkSourcePort sets the networkSourcePort property value. The networkSourcePort property
func (m *TiIndicator) SetNetworkSourcePort(value *int32)() {
    err := m.GetBackingStore().Set("networkSourcePort", value)
    if err != nil {
        panic(err)
    }
}
// SetPassiveOnly sets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools won't notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they'll simply log that a match occurred but won't perform the action. Default value is false.
func (m *TiIndicator) SetPassiveOnly(value *bool)() {
    err := m.GetBackingStore().Set("passiveOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero isn't severe at all. Default value is 3.
func (m *TiIndicator) SetSeverity(value *int32)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
func (m *TiIndicator) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetProduct sets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
func (m *TiIndicator) SetTargetProduct(value *string)() {
    err := m.GetBackingStore().Set("targetProduct", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatType sets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
func (m *TiIndicator) SetThreatType(value *string)() {
    err := m.GetBackingStore().Set("threatType", value)
    if err != nil {
        panic(err)
    }
}
// SetTlpLevel sets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
func (m *TiIndicator) SetTlpLevel(value *TlpLevel)() {
    err := m.GetBackingStore().Set("tlpLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetUrl sets the url property value. The url property
func (m *TiIndicator) SetUrl(value *string)() {
    err := m.GetBackingStore().Set("url", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAgent sets the userAgent property value. The userAgent property
func (m *TiIndicator) SetUserAgent(value *string)() {
    err := m.GetBackingStore().Set("userAgent", value)
    if err != nil {
        panic(err)
    }
}
type TiIndicatorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*TiAction)
    GetActivityGroupNames()([]string)
    GetAdditionalInformation()(*string)
    GetAzureTenantId()(*string)
    GetConfidence()(*int32)
    GetDescription()(*string)
    GetDiamondModel()(*DiamondModel)
    GetDomainName()(*string)
    GetEmailEncoding()(*string)
    GetEmailLanguage()(*string)
    GetEmailRecipient()(*string)
    GetEmailSenderAddress()(*string)
    GetEmailSenderName()(*string)
    GetEmailSourceDomain()(*string)
    GetEmailSourceIpAddress()(*string)
    GetEmailSubject()(*string)
    GetEmailXMailer()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetFileCompileDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFileCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFileHashType()(*FileHashType)
    GetFileHashValue()(*string)
    GetFileMutexName()(*string)
    GetFileName()(*string)
    GetFilePacker()(*string)
    GetFilePath()(*string)
    GetFileSize()(*int64)
    GetFileType()(*string)
    GetIngestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsActive()(*bool)
    GetKillChain()([]string)
    GetKnownFalsePositives()(*string)
    GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareFamilyNames()([]string)
    GetNetworkCidrBlock()(*string)
    GetNetworkDestinationAsn()(*int64)
    GetNetworkDestinationCidrBlock()(*string)
    GetNetworkDestinationIPv4()(*string)
    GetNetworkDestinationIPv6()(*string)
    GetNetworkDestinationPort()(*int32)
    GetNetworkIPv4()(*string)
    GetNetworkIPv6()(*string)
    GetNetworkPort()(*int32)
    GetNetworkProtocol()(*int32)
    GetNetworkSourceAsn()(*int64)
    GetNetworkSourceCidrBlock()(*string)
    GetNetworkSourceIPv4()(*string)
    GetNetworkSourceIPv6()(*string)
    GetNetworkSourcePort()(*int32)
    GetPassiveOnly()(*bool)
    GetSeverity()(*int32)
    GetTags()([]string)
    GetTargetProduct()(*string)
    GetThreatType()(*string)
    GetTlpLevel()(*TlpLevel)
    GetUrl()(*string)
    GetUserAgent()(*string)
    SetAction(value *TiAction)()
    SetActivityGroupNames(value []string)()
    SetAdditionalInformation(value *string)()
    SetAzureTenantId(value *string)()
    SetConfidence(value *int32)()
    SetDescription(value *string)()
    SetDiamondModel(value *DiamondModel)()
    SetDomainName(value *string)()
    SetEmailEncoding(value *string)()
    SetEmailLanguage(value *string)()
    SetEmailRecipient(value *string)()
    SetEmailSenderAddress(value *string)()
    SetEmailSenderName(value *string)()
    SetEmailSourceDomain(value *string)()
    SetEmailSourceIpAddress(value *string)()
    SetEmailSubject(value *string)()
    SetEmailXMailer(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetFileCompileDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFileCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFileHashType(value *FileHashType)()
    SetFileHashValue(value *string)()
    SetFileMutexName(value *string)()
    SetFileName(value *string)()
    SetFilePacker(value *string)()
    SetFilePath(value *string)()
    SetFileSize(value *int64)()
    SetFileType(value *string)()
    SetIngestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsActive(value *bool)()
    SetKillChain(value []string)()
    SetKnownFalsePositives(value *string)()
    SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareFamilyNames(value []string)()
    SetNetworkCidrBlock(value *string)()
    SetNetworkDestinationAsn(value *int64)()
    SetNetworkDestinationCidrBlock(value *string)()
    SetNetworkDestinationIPv4(value *string)()
    SetNetworkDestinationIPv6(value *string)()
    SetNetworkDestinationPort(value *int32)()
    SetNetworkIPv4(value *string)()
    SetNetworkIPv6(value *string)()
    SetNetworkPort(value *int32)()
    SetNetworkProtocol(value *int32)()
    SetNetworkSourceAsn(value *int64)()
    SetNetworkSourceCidrBlock(value *string)()
    SetNetworkSourceIPv4(value *string)()
    SetNetworkSourceIPv6(value *string)()
    SetNetworkSourcePort(value *int32)()
    SetPassiveOnly(value *bool)()
    SetSeverity(value *int32)()
    SetTags(value []string)()
    SetTargetProduct(value *string)()
    SetThreatType(value *string)()
    SetTlpLevel(value *TlpLevel)()
    SetUrl(value *string)()
    SetUserAgent(value *string)()
}
