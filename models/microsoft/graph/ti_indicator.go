package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TiIndicator 
type TiIndicator struct {
    Entity
    // The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
    action *TiAction;
    // The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
    activityGroupNames []string;
    // A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
    additionalInformation *string;
    // Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
    azureTenantId *string;
    // An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
    confidence *int32;
    // Brief description (100 characters or less) of the threat represented by the indicator. Required.
    description *string;
    // The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
    diamondModel *DiamondModel;
    // 
    domainName *string;
    // 
    emailEncoding *string;
    // 
    emailLanguage *string;
    // 
    emailRecipient *string;
    // 
    emailSenderAddress *string;
    // 
    emailSenderName *string;
    // 
    emailSourceDomain *string;
    // 
    emailSourceIpAddress *string;
    // 
    emailSubject *string;
    // 
    emailXMailer *string;
    // DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
    externalId *string;
    // 
    fileCompileDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    fileCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    fileHashType *FileHashType;
    // 
    fileHashValue *string;
    // 
    fileMutexName *string;
    // 
    fileName *string;
    // 
    filePacker *string;
    // 
    filePath *string;
    // 
    fileSize *int64;
    // 
    fileType *string;
    // Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    ingestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
    isActive *bool;
    // A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
    killChain []string;
    // Scenarios in which the indicator may cause false positives. This should be human-readable text.
    knownFalsePositives *string;
    // The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
    malwareFamilyNames []string;
    // 
    networkCidrBlock *string;
    // 
    networkDestinationAsn *int64;
    // 
    networkDestinationCidrBlock *string;
    // 
    networkDestinationIPv4 *string;
    // 
    networkDestinationIPv6 *string;
    // 
    networkDestinationPort *int32;
    // 
    networkIPv4 *string;
    // 
    networkIPv6 *string;
    // 
    networkPort *int32;
    // 
    networkProtocol *int32;
    // 
    networkSourceAsn *int64;
    // 
    networkSourceCidrBlock *string;
    // 
    networkSourceIPv4 *string;
    // 
    networkSourceIPv6 *string;
    // 
    networkSourcePort *int32;
    // Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
    passiveOnly *bool;
    // An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
    severity *int32;
    // A JSON array of strings that stores arbitrary tags/keywords.
    tags []string;
    // A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
    targetProduct *string;
    // Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
    threatType *string;
    // Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
    tlpLevel *TlpLevel;
    // 
    url *string;
    // 
    userAgent *string;
}
// NewTiIndicator instantiates a new tiIndicator and sets the default values.
func NewTiIndicator()(*TiIndicator) {
    m := &TiIndicator{
        Entity: *NewEntity(),
    }
    return m
}
// GetAction gets the action property value. The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
func (m *TiIndicator) GetAction()(*TiAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetActivityGroupNames gets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
func (m *TiIndicator) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
// GetAdditionalInformation gets the additionalInformation property value. A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
func (m *TiIndicator) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// GetAzureTenantId gets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
func (m *TiIndicator) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetConfidence gets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
func (m *TiIndicator) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetDescription gets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
func (m *TiIndicator) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDiamondModel gets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
func (m *TiIndicator) GetDiamondModel()(*DiamondModel) {
    if m == nil {
        return nil
    } else {
        return m.diamondModel
    }
}
// GetDomainName gets the domainName property value. 
func (m *TiIndicator) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// GetEmailEncoding gets the emailEncoding property value. 
func (m *TiIndicator) GetEmailEncoding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailEncoding
    }
}
// GetEmailLanguage gets the emailLanguage property value. 
func (m *TiIndicator) GetEmailLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailLanguage
    }
}
// GetEmailRecipient gets the emailRecipient property value. 
func (m *TiIndicator) GetEmailRecipient()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailRecipient
    }
}
// GetEmailSenderAddress gets the emailSenderAddress property value. 
func (m *TiIndicator) GetEmailSenderAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSenderAddress
    }
}
// GetEmailSenderName gets the emailSenderName property value. 
func (m *TiIndicator) GetEmailSenderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSenderName
    }
}
// GetEmailSourceDomain gets the emailSourceDomain property value. 
func (m *TiIndicator) GetEmailSourceDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSourceDomain
    }
}
// GetEmailSourceIpAddress gets the emailSourceIpAddress property value. 
func (m *TiIndicator) GetEmailSourceIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSourceIpAddress
    }
}
// GetEmailSubject gets the emailSubject property value. 
func (m *TiIndicator) GetEmailSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSubject
    }
}
// GetEmailXMailer gets the emailXMailer property value. 
func (m *TiIndicator) GetEmailXMailer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailXMailer
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
func (m *TiIndicator) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetExternalId gets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
func (m *TiIndicator) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFileCompileDateTime gets the fileCompileDateTime property value. 
func (m *TiIndicator) GetFileCompileDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fileCompileDateTime
    }
}
// GetFileCreatedDateTime gets the fileCreatedDateTime property value. 
func (m *TiIndicator) GetFileCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fileCreatedDateTime
    }
}
// GetFileHashType gets the fileHashType property value. 
func (m *TiIndicator) GetFileHashType()(*FileHashType) {
    if m == nil {
        return nil
    } else {
        return m.fileHashType
    }
}
// GetFileHashValue gets the fileHashValue property value. 
func (m *TiIndicator) GetFileHashValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileHashValue
    }
}
// GetFileMutexName gets the fileMutexName property value. 
func (m *TiIndicator) GetFileMutexName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileMutexName
    }
}
// GetFileName gets the fileName property value. 
func (m *TiIndicator) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetFilePacker gets the filePacker property value. 
func (m *TiIndicator) GetFilePacker()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePacker
    }
}
// GetFilePath gets the filePath property value. 
func (m *TiIndicator) GetFilePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePath
    }
}
// GetFileSize gets the fileSize property value. 
func (m *TiIndicator) GetFileSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileSize
    }
}
// GetFileType gets the fileType property value. 
func (m *TiIndicator) GetFileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileType
    }
}
// GetIngestedDateTime gets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) GetIngestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.ingestedDateTime
    }
}
// GetIsActive gets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
func (m *TiIndicator) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetKillChain gets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
func (m *TiIndicator) GetKillChain()([]string) {
    if m == nil {
        return nil
    } else {
        return m.killChain
    }
}
// GetKnownFalsePositives gets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
func (m *TiIndicator) GetKnownFalsePositives()(*string) {
    if m == nil {
        return nil
    } else {
        return m.knownFalsePositives
    }
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// GetMalwareFamilyNames gets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
func (m *TiIndicator) GetMalwareFamilyNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.malwareFamilyNames
    }
}
// GetNetworkCidrBlock gets the networkCidrBlock property value. 
func (m *TiIndicator) GetNetworkCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkCidrBlock
    }
}
// GetNetworkDestinationAsn gets the networkDestinationAsn property value. 
func (m *TiIndicator) GetNetworkDestinationAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationAsn
    }
}
// GetNetworkDestinationCidrBlock gets the networkDestinationCidrBlock property value. 
func (m *TiIndicator) GetNetworkDestinationCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationCidrBlock
    }
}
// GetNetworkDestinationIPv4 gets the networkDestinationIPv4 property value. 
func (m *TiIndicator) GetNetworkDestinationIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationIPv4
    }
}
// GetNetworkDestinationIPv6 gets the networkDestinationIPv6 property value. 
func (m *TiIndicator) GetNetworkDestinationIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationIPv6
    }
}
// GetNetworkDestinationPort gets the networkDestinationPort property value. 
func (m *TiIndicator) GetNetworkDestinationPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationPort
    }
}
// GetNetworkIPv4 gets the networkIPv4 property value. 
func (m *TiIndicator) GetNetworkIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkIPv4
    }
}
// GetNetworkIPv6 gets the networkIPv6 property value. 
func (m *TiIndicator) GetNetworkIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkIPv6
    }
}
// GetNetworkPort gets the networkPort property value. 
func (m *TiIndicator) GetNetworkPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkPort
    }
}
// GetNetworkProtocol gets the networkProtocol property value. 
func (m *TiIndicator) GetNetworkProtocol()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkProtocol
    }
}
// GetNetworkSourceAsn gets the networkSourceAsn property value. 
func (m *TiIndicator) GetNetworkSourceAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceAsn
    }
}
// GetNetworkSourceCidrBlock gets the networkSourceCidrBlock property value. 
func (m *TiIndicator) GetNetworkSourceCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceCidrBlock
    }
}
// GetNetworkSourceIPv4 gets the networkSourceIPv4 property value. 
func (m *TiIndicator) GetNetworkSourceIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceIPv4
    }
}
// GetNetworkSourceIPv6 gets the networkSourceIPv6 property value. 
func (m *TiIndicator) GetNetworkSourceIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceIPv6
    }
}
// GetNetworkSourcePort gets the networkSourcePort property value. 
func (m *TiIndicator) GetNetworkSourcePort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkSourcePort
    }
}
// GetPassiveOnly gets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
func (m *TiIndicator) GetPassiveOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passiveOnly
    }
}
// GetSeverity gets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
func (m *TiIndicator) GetSeverity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// GetTags gets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
func (m *TiIndicator) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetTargetProduct gets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
func (m *TiIndicator) GetTargetProduct()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetProduct
    }
}
// GetThreatType gets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
func (m *TiIndicator) GetThreatType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatType
    }
}
// GetTlpLevel gets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
func (m *TiIndicator) GetTlpLevel()(*TlpLevel) {
    if m == nil {
        return nil
    } else {
        return m.tlpLevel
    }
}
// GetUrl gets the url property value. 
func (m *TiIndicator) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetUserAgent gets the userAgent property value. 
func (m *TiIndicator) GetUserAgent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userAgent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TiIndicator) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTiAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*TiAction))
        }
        return nil
    }
    res["activityGroupNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetActivityGroupNames(res)
        }
        return nil
    }
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["diamondModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiamondModel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiamondModel(val.(*DiamondModel))
        }
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["emailEncoding"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailEncoding(val)
        }
        return nil
    }
    res["emailLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailLanguage(val)
        }
        return nil
    }
    res["emailRecipient"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailRecipient(val)
        }
        return nil
    }
    res["emailSenderAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSenderAddress(val)
        }
        return nil
    }
    res["emailSenderName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSenderName(val)
        }
        return nil
    }
    res["emailSourceDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSourceDomain(val)
        }
        return nil
    }
    res["emailSourceIpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSourceIpAddress(val)
        }
        return nil
    }
    res["emailSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSubject(val)
        }
        return nil
    }
    res["emailXMailer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailXMailer(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["fileCompileDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileCompileDateTime(val)
        }
        return nil
    }
    res["fileCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileCreatedDateTime(val)
        }
        return nil
    }
    res["fileHashType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileHashType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHashType(val.(*FileHashType))
        }
        return nil
    }
    res["fileHashValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHashValue(val)
        }
        return nil
    }
    res["fileMutexName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileMutexName(val)
        }
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["filePacker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePacker(val)
        }
        return nil
    }
    res["filePath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
        }
        return nil
    }
    res["fileSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["fileType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileType(val)
        }
        return nil
    }
    res["ingestedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIngestedDateTime(val)
        }
        return nil
    }
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["killChain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKillChain(res)
        }
        return nil
    }
    res["knownFalsePositives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKnownFalsePositives(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["malwareFamilyNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMalwareFamilyNames(res)
        }
        return nil
    }
    res["networkCidrBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkCidrBlock(val)
        }
        return nil
    }
    res["networkDestinationAsn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationAsn(val)
        }
        return nil
    }
    res["networkDestinationCidrBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationCidrBlock(val)
        }
        return nil
    }
    res["networkDestinationIPv4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationIPv4(val)
        }
        return nil
    }
    res["networkDestinationIPv6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationIPv6(val)
        }
        return nil
    }
    res["networkDestinationPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkDestinationPort(val)
        }
        return nil
    }
    res["networkIPv4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkIPv4(val)
        }
        return nil
    }
    res["networkIPv6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkIPv6(val)
        }
        return nil
    }
    res["networkPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkPort(val)
        }
        return nil
    }
    res["networkProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProtocol(val)
        }
        return nil
    }
    res["networkSourceAsn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceAsn(val)
        }
        return nil
    }
    res["networkSourceCidrBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceCidrBlock(val)
        }
        return nil
    }
    res["networkSourceIPv4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceIPv4(val)
        }
        return nil
    }
    res["networkSourceIPv6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourceIPv6(val)
        }
        return nil
    }
    res["networkSourcePort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSourcePort(val)
        }
        return nil
    }
    res["passiveOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassiveOnly(val)
        }
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val)
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["targetProduct"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetProduct(val)
        }
        return nil
    }
    res["threatType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatType(val)
        }
        return nil
    }
    res["tlpLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTlpLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlpLevel(val.(*TlpLevel))
        }
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["userAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *TiIndicator) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TiIndicator) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.action = value
    }
}
// SetActivityGroupNames sets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
func (m *TiIndicator) SetActivityGroupNames(value []string)() {
    if m != nil {
        m.activityGroupNames = value
    }
}
// SetAdditionalInformation sets the additionalInformation property value. A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
func (m *TiIndicator) SetAdditionalInformation(value *string)() {
    if m != nil {
        m.additionalInformation = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
func (m *TiIndicator) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetConfidence sets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
func (m *TiIndicator) SetConfidence(value *int32)() {
    if m != nil {
        m.confidence = value
    }
}
// SetDescription sets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
func (m *TiIndicator) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDiamondModel sets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
func (m *TiIndicator) SetDiamondModel(value *DiamondModel)() {
    if m != nil {
        m.diamondModel = value
    }
}
// SetDomainName sets the domainName property value. 
func (m *TiIndicator) SetDomainName(value *string)() {
    if m != nil {
        m.domainName = value
    }
}
// SetEmailEncoding sets the emailEncoding property value. 
func (m *TiIndicator) SetEmailEncoding(value *string)() {
    if m != nil {
        m.emailEncoding = value
    }
}
// SetEmailLanguage sets the emailLanguage property value. 
func (m *TiIndicator) SetEmailLanguage(value *string)() {
    if m != nil {
        m.emailLanguage = value
    }
}
// SetEmailRecipient sets the emailRecipient property value. 
func (m *TiIndicator) SetEmailRecipient(value *string)() {
    if m != nil {
        m.emailRecipient = value
    }
}
// SetEmailSenderAddress sets the emailSenderAddress property value. 
func (m *TiIndicator) SetEmailSenderAddress(value *string)() {
    if m != nil {
        m.emailSenderAddress = value
    }
}
// SetEmailSenderName sets the emailSenderName property value. 
func (m *TiIndicator) SetEmailSenderName(value *string)() {
    if m != nil {
        m.emailSenderName = value
    }
}
// SetEmailSourceDomain sets the emailSourceDomain property value. 
func (m *TiIndicator) SetEmailSourceDomain(value *string)() {
    if m != nil {
        m.emailSourceDomain = value
    }
}
// SetEmailSourceIpAddress sets the emailSourceIpAddress property value. 
func (m *TiIndicator) SetEmailSourceIpAddress(value *string)() {
    if m != nil {
        m.emailSourceIpAddress = value
    }
}
// SetEmailSubject sets the emailSubject property value. 
func (m *TiIndicator) SetEmailSubject(value *string)() {
    if m != nil {
        m.emailSubject = value
    }
}
// SetEmailXMailer sets the emailXMailer property value. 
func (m *TiIndicator) SetEmailXMailer(value *string)() {
    if m != nil {
        m.emailXMailer = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
func (m *TiIndicator) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetExternalId sets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
func (m *TiIndicator) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetFileCompileDateTime sets the fileCompileDateTime property value. 
func (m *TiIndicator) SetFileCompileDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.fileCompileDateTime = value
    }
}
// SetFileCreatedDateTime sets the fileCreatedDateTime property value. 
func (m *TiIndicator) SetFileCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.fileCreatedDateTime = value
    }
}
// SetFileHashType sets the fileHashType property value. 
func (m *TiIndicator) SetFileHashType(value *FileHashType)() {
    if m != nil {
        m.fileHashType = value
    }
}
// SetFileHashValue sets the fileHashValue property value. 
func (m *TiIndicator) SetFileHashValue(value *string)() {
    if m != nil {
        m.fileHashValue = value
    }
}
// SetFileMutexName sets the fileMutexName property value. 
func (m *TiIndicator) SetFileMutexName(value *string)() {
    if m != nil {
        m.fileMutexName = value
    }
}
// SetFileName sets the fileName property value. 
func (m *TiIndicator) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetFilePacker sets the filePacker property value. 
func (m *TiIndicator) SetFilePacker(value *string)() {
    if m != nil {
        m.filePacker = value
    }
}
// SetFilePath sets the filePath property value. 
func (m *TiIndicator) SetFilePath(value *string)() {
    if m != nil {
        m.filePath = value
    }
}
// SetFileSize sets the fileSize property value. 
func (m *TiIndicator) SetFileSize(value *int64)() {
    if m != nil {
        m.fileSize = value
    }
}
// SetFileType sets the fileType property value. 
func (m *TiIndicator) SetFileType(value *string)() {
    if m != nil {
        m.fileType = value
    }
}
// SetIngestedDateTime sets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) SetIngestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.ingestedDateTime = value
    }
}
// SetIsActive sets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
func (m *TiIndicator) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetKillChain sets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
func (m *TiIndicator) SetKillChain(value []string)() {
    if m != nil {
        m.killChain = value
    }
}
// SetKnownFalsePositives sets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
func (m *TiIndicator) SetKnownFalsePositives(value *string)() {
    if m != nil {
        m.knownFalsePositives = value
    }
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastReportedDateTime = value
    }
}
// SetMalwareFamilyNames sets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
func (m *TiIndicator) SetMalwareFamilyNames(value []string)() {
    if m != nil {
        m.malwareFamilyNames = value
    }
}
// SetNetworkCidrBlock sets the networkCidrBlock property value. 
func (m *TiIndicator) SetNetworkCidrBlock(value *string)() {
    if m != nil {
        m.networkCidrBlock = value
    }
}
// SetNetworkDestinationAsn sets the networkDestinationAsn property value. 
func (m *TiIndicator) SetNetworkDestinationAsn(value *int64)() {
    if m != nil {
        m.networkDestinationAsn = value
    }
}
// SetNetworkDestinationCidrBlock sets the networkDestinationCidrBlock property value. 
func (m *TiIndicator) SetNetworkDestinationCidrBlock(value *string)() {
    if m != nil {
        m.networkDestinationCidrBlock = value
    }
}
// SetNetworkDestinationIPv4 sets the networkDestinationIPv4 property value. 
func (m *TiIndicator) SetNetworkDestinationIPv4(value *string)() {
    if m != nil {
        m.networkDestinationIPv4 = value
    }
}
// SetNetworkDestinationIPv6 sets the networkDestinationIPv6 property value. 
func (m *TiIndicator) SetNetworkDestinationIPv6(value *string)() {
    if m != nil {
        m.networkDestinationIPv6 = value
    }
}
// SetNetworkDestinationPort sets the networkDestinationPort property value. 
func (m *TiIndicator) SetNetworkDestinationPort(value *int32)() {
    if m != nil {
        m.networkDestinationPort = value
    }
}
// SetNetworkIPv4 sets the networkIPv4 property value. 
func (m *TiIndicator) SetNetworkIPv4(value *string)() {
    if m != nil {
        m.networkIPv4 = value
    }
}
// SetNetworkIPv6 sets the networkIPv6 property value. 
func (m *TiIndicator) SetNetworkIPv6(value *string)() {
    if m != nil {
        m.networkIPv6 = value
    }
}
// SetNetworkPort sets the networkPort property value. 
func (m *TiIndicator) SetNetworkPort(value *int32)() {
    if m != nil {
        m.networkPort = value
    }
}
// SetNetworkProtocol sets the networkProtocol property value. 
func (m *TiIndicator) SetNetworkProtocol(value *int32)() {
    if m != nil {
        m.networkProtocol = value
    }
}
// SetNetworkSourceAsn sets the networkSourceAsn property value. 
func (m *TiIndicator) SetNetworkSourceAsn(value *int64)() {
    if m != nil {
        m.networkSourceAsn = value
    }
}
// SetNetworkSourceCidrBlock sets the networkSourceCidrBlock property value. 
func (m *TiIndicator) SetNetworkSourceCidrBlock(value *string)() {
    if m != nil {
        m.networkSourceCidrBlock = value
    }
}
// SetNetworkSourceIPv4 sets the networkSourceIPv4 property value. 
func (m *TiIndicator) SetNetworkSourceIPv4(value *string)() {
    if m != nil {
        m.networkSourceIPv4 = value
    }
}
// SetNetworkSourceIPv6 sets the networkSourceIPv6 property value. 
func (m *TiIndicator) SetNetworkSourceIPv6(value *string)() {
    if m != nil {
        m.networkSourceIPv6 = value
    }
}
// SetNetworkSourcePort sets the networkSourcePort property value. 
func (m *TiIndicator) SetNetworkSourcePort(value *int32)() {
    if m != nil {
        m.networkSourcePort = value
    }
}
// SetPassiveOnly sets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
func (m *TiIndicator) SetPassiveOnly(value *bool)() {
    if m != nil {
        m.passiveOnly = value
    }
}
// SetSeverity sets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
func (m *TiIndicator) SetSeverity(value *int32)() {
    if m != nil {
        m.severity = value
    }
}
// SetTags sets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
func (m *TiIndicator) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetTargetProduct sets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
func (m *TiIndicator) SetTargetProduct(value *string)() {
    if m != nil {
        m.targetProduct = value
    }
}
// SetThreatType sets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
func (m *TiIndicator) SetThreatType(value *string)() {
    if m != nil {
        m.threatType = value
    }
}
// SetTlpLevel sets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
func (m *TiIndicator) SetTlpLevel(value *TlpLevel)() {
    if m != nil {
        m.tlpLevel = value
    }
}
// SetUrl sets the url property value. 
func (m *TiIndicator) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
// SetUserAgent sets the userAgent property value. 
func (m *TiIndicator) SetUserAgent(value *string)() {
    if m != nil {
        m.userAgent = value
    }
}
