package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new tiIndicator and sets the default values.
func NewTiIndicator()(*TiIndicator) {
    m := &TiIndicator{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the action property value. The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
func (m *TiIndicator) GetAction()(*TiAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// Gets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
func (m *TiIndicator) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
// Gets the additionalInformation property value. A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
func (m *TiIndicator) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// Gets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
func (m *TiIndicator) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
func (m *TiIndicator) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// Gets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
func (m *TiIndicator) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
func (m *TiIndicator) GetDiamondModel()(*DiamondModel) {
    if m == nil {
        return nil
    } else {
        return m.diamondModel
    }
}
// Gets the domainName property value. 
func (m *TiIndicator) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// Gets the emailEncoding property value. 
func (m *TiIndicator) GetEmailEncoding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailEncoding
    }
}
// Gets the emailLanguage property value. 
func (m *TiIndicator) GetEmailLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailLanguage
    }
}
// Gets the emailRecipient property value. 
func (m *TiIndicator) GetEmailRecipient()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailRecipient
    }
}
// Gets the emailSenderAddress property value. 
func (m *TiIndicator) GetEmailSenderAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSenderAddress
    }
}
// Gets the emailSenderName property value. 
func (m *TiIndicator) GetEmailSenderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSenderName
    }
}
// Gets the emailSourceDomain property value. 
func (m *TiIndicator) GetEmailSourceDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSourceDomain
    }
}
// Gets the emailSourceIpAddress property value. 
func (m *TiIndicator) GetEmailSourceIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSourceIpAddress
    }
}
// Gets the emailSubject property value. 
func (m *TiIndicator) GetEmailSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSubject
    }
}
// Gets the emailXMailer property value. 
func (m *TiIndicator) GetEmailXMailer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailXMailer
    }
}
// Gets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
func (m *TiIndicator) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
func (m *TiIndicator) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the fileCompileDateTime property value. 
func (m *TiIndicator) GetFileCompileDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fileCompileDateTime
    }
}
// Gets the fileCreatedDateTime property value. 
func (m *TiIndicator) GetFileCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fileCreatedDateTime
    }
}
// Gets the fileHashType property value. 
func (m *TiIndicator) GetFileHashType()(*FileHashType) {
    if m == nil {
        return nil
    } else {
        return m.fileHashType
    }
}
// Gets the fileHashValue property value. 
func (m *TiIndicator) GetFileHashValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileHashValue
    }
}
// Gets the fileMutexName property value. 
func (m *TiIndicator) GetFileMutexName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileMutexName
    }
}
// Gets the fileName property value. 
func (m *TiIndicator) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Gets the filePacker property value. 
func (m *TiIndicator) GetFilePacker()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePacker
    }
}
// Gets the filePath property value. 
func (m *TiIndicator) GetFilePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePath
    }
}
// Gets the fileSize property value. 
func (m *TiIndicator) GetFileSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileSize
    }
}
// Gets the fileType property value. 
func (m *TiIndicator) GetFileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileType
    }
}
// Gets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) GetIngestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.ingestedDateTime
    }
}
// Gets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
func (m *TiIndicator) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// Gets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
func (m *TiIndicator) GetKillChain()([]string) {
    if m == nil {
        return nil
    } else {
        return m.killChain
    }
}
// Gets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
func (m *TiIndicator) GetKnownFalsePositives()(*string) {
    if m == nil {
        return nil
    } else {
        return m.knownFalsePositives
    }
}
// Gets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// Gets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
func (m *TiIndicator) GetMalwareFamilyNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.malwareFamilyNames
    }
}
// Gets the networkCidrBlock property value. 
func (m *TiIndicator) GetNetworkCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkCidrBlock
    }
}
// Gets the networkDestinationAsn property value. 
func (m *TiIndicator) GetNetworkDestinationAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationAsn
    }
}
// Gets the networkDestinationCidrBlock property value. 
func (m *TiIndicator) GetNetworkDestinationCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationCidrBlock
    }
}
// Gets the networkDestinationIPv4 property value. 
func (m *TiIndicator) GetNetworkDestinationIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationIPv4
    }
}
// Gets the networkDestinationIPv6 property value. 
func (m *TiIndicator) GetNetworkDestinationIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationIPv6
    }
}
// Gets the networkDestinationPort property value. 
func (m *TiIndicator) GetNetworkDestinationPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationPort
    }
}
// Gets the networkIPv4 property value. 
func (m *TiIndicator) GetNetworkIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkIPv4
    }
}
// Gets the networkIPv6 property value. 
func (m *TiIndicator) GetNetworkIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkIPv6
    }
}
// Gets the networkPort property value. 
func (m *TiIndicator) GetNetworkPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkPort
    }
}
// Gets the networkProtocol property value. 
func (m *TiIndicator) GetNetworkProtocol()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkProtocol
    }
}
// Gets the networkSourceAsn property value. 
func (m *TiIndicator) GetNetworkSourceAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceAsn
    }
}
// Gets the networkSourceCidrBlock property value. 
func (m *TiIndicator) GetNetworkSourceCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceCidrBlock
    }
}
// Gets the networkSourceIPv4 property value. 
func (m *TiIndicator) GetNetworkSourceIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceIPv4
    }
}
// Gets the networkSourceIPv6 property value. 
func (m *TiIndicator) GetNetworkSourceIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceIPv6
    }
}
// Gets the networkSourcePort property value. 
func (m *TiIndicator) GetNetworkSourcePort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkSourcePort
    }
}
// Gets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
func (m *TiIndicator) GetPassiveOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passiveOnly
    }
}
// Gets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
func (m *TiIndicator) GetSeverity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// Gets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
func (m *TiIndicator) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
func (m *TiIndicator) GetTargetProduct()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetProduct
    }
}
// Gets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
func (m *TiIndicator) GetThreatType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatType
    }
}
// Gets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
func (m *TiIndicator) GetTlpLevel()(*TlpLevel) {
    if m == nil {
        return nil
    } else {
        return m.tlpLevel
    }
}
// Gets the url property value. 
func (m *TiIndicator) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// Gets the userAgent property value. 
func (m *TiIndicator) GetUserAgent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userAgent
    }
}
// The deserialization information for the current model
func (m *TiIndicator) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTiAction)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TiAction)
            m.SetAction(&cast)
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
            cast := val.(DiamondModel)
            m.SetDiamondModel(&cast)
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
            cast := val.(FileHashType)
            m.SetFileHashType(&cast)
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
            cast := val.(TlpLevel)
            m.SetTlpLevel(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TiIndicator) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
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
        cast := m.GetDiamondModel().String()
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
        cast := m.GetFileHashType().String()
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
    {
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
    {
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
    {
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
        cast := m.GetTlpLevel().String()
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
// Sets the action property value. The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
// Parameters:
//  - value : Value to set for the action property.
func (m *TiIndicator) SetAction(value *TiAction)() {
    m.action = value
}
// Sets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
// Parameters:
//  - value : Value to set for the activityGroupNames property.
func (m *TiIndicator) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// Sets the additionalInformation property value. A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
// Parameters:
//  - value : Value to set for the additionalInformation property.
func (m *TiIndicator) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// Sets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *TiIndicator) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
// Parameters:
//  - value : Value to set for the confidence property.
func (m *TiIndicator) SetConfidence(value *int32)() {
    m.confidence = value
}
// Sets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
// Parameters:
//  - value : Value to set for the description property.
func (m *TiIndicator) SetDescription(value *string)() {
    m.description = value
}
// Sets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
// Parameters:
//  - value : Value to set for the diamondModel property.
func (m *TiIndicator) SetDiamondModel(value *DiamondModel)() {
    m.diamondModel = value
}
// Sets the domainName property value. 
// Parameters:
//  - value : Value to set for the domainName property.
func (m *TiIndicator) SetDomainName(value *string)() {
    m.domainName = value
}
// Sets the emailEncoding property value. 
// Parameters:
//  - value : Value to set for the emailEncoding property.
func (m *TiIndicator) SetEmailEncoding(value *string)() {
    m.emailEncoding = value
}
// Sets the emailLanguage property value. 
// Parameters:
//  - value : Value to set for the emailLanguage property.
func (m *TiIndicator) SetEmailLanguage(value *string)() {
    m.emailLanguage = value
}
// Sets the emailRecipient property value. 
// Parameters:
//  - value : Value to set for the emailRecipient property.
func (m *TiIndicator) SetEmailRecipient(value *string)() {
    m.emailRecipient = value
}
// Sets the emailSenderAddress property value. 
// Parameters:
//  - value : Value to set for the emailSenderAddress property.
func (m *TiIndicator) SetEmailSenderAddress(value *string)() {
    m.emailSenderAddress = value
}
// Sets the emailSenderName property value. 
// Parameters:
//  - value : Value to set for the emailSenderName property.
func (m *TiIndicator) SetEmailSenderName(value *string)() {
    m.emailSenderName = value
}
// Sets the emailSourceDomain property value. 
// Parameters:
//  - value : Value to set for the emailSourceDomain property.
func (m *TiIndicator) SetEmailSourceDomain(value *string)() {
    m.emailSourceDomain = value
}
// Sets the emailSourceIpAddress property value. 
// Parameters:
//  - value : Value to set for the emailSourceIpAddress property.
func (m *TiIndicator) SetEmailSourceIpAddress(value *string)() {
    m.emailSourceIpAddress = value
}
// Sets the emailSubject property value. 
// Parameters:
//  - value : Value to set for the emailSubject property.
func (m *TiIndicator) SetEmailSubject(value *string)() {
    m.emailSubject = value
}
// Sets the emailXMailer property value. 
// Parameters:
//  - value : Value to set for the emailXMailer property.
func (m *TiIndicator) SetEmailXMailer(value *string)() {
    m.emailXMailer = value
}
// Sets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *TiIndicator) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
// Parameters:
//  - value : Value to set for the externalId property.
func (m *TiIndicator) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the fileCompileDateTime property value. 
// Parameters:
//  - value : Value to set for the fileCompileDateTime property.
func (m *TiIndicator) SetFileCompileDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fileCompileDateTime = value
}
// Sets the fileCreatedDateTime property value. 
// Parameters:
//  - value : Value to set for the fileCreatedDateTime property.
func (m *TiIndicator) SetFileCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fileCreatedDateTime = value
}
// Sets the fileHashType property value. 
// Parameters:
//  - value : Value to set for the fileHashType property.
func (m *TiIndicator) SetFileHashType(value *FileHashType)() {
    m.fileHashType = value
}
// Sets the fileHashValue property value. 
// Parameters:
//  - value : Value to set for the fileHashValue property.
func (m *TiIndicator) SetFileHashValue(value *string)() {
    m.fileHashValue = value
}
// Sets the fileMutexName property value. 
// Parameters:
//  - value : Value to set for the fileMutexName property.
func (m *TiIndicator) SetFileMutexName(value *string)() {
    m.fileMutexName = value
}
// Sets the fileName property value. 
// Parameters:
//  - value : Value to set for the fileName property.
func (m *TiIndicator) SetFileName(value *string)() {
    m.fileName = value
}
// Sets the filePacker property value. 
// Parameters:
//  - value : Value to set for the filePacker property.
func (m *TiIndicator) SetFilePacker(value *string)() {
    m.filePacker = value
}
// Sets the filePath property value. 
// Parameters:
//  - value : Value to set for the filePath property.
func (m *TiIndicator) SetFilePath(value *string)() {
    m.filePath = value
}
// Sets the fileSize property value. 
// Parameters:
//  - value : Value to set for the fileSize property.
func (m *TiIndicator) SetFileSize(value *int64)() {
    m.fileSize = value
}
// Sets the fileType property value. 
// Parameters:
//  - value : Value to set for the fileType property.
func (m *TiIndicator) SetFileType(value *string)() {
    m.fileType = value
}
// Sets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the ingestedDateTime property.
func (m *TiIndicator) SetIngestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.ingestedDateTime = value
}
// Sets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
// Parameters:
//  - value : Value to set for the isActive property.
func (m *TiIndicator) SetIsActive(value *bool)() {
    m.isActive = value
}
// Sets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
// Parameters:
//  - value : Value to set for the killChain property.
func (m *TiIndicator) SetKillChain(value []string)() {
    m.killChain = value
}
// Sets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
// Parameters:
//  - value : Value to set for the knownFalsePositives property.
func (m *TiIndicator) SetKnownFalsePositives(value *string)() {
    m.knownFalsePositives = value
}
// Sets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastReportedDateTime property.
func (m *TiIndicator) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// Sets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
// Parameters:
//  - value : Value to set for the malwareFamilyNames property.
func (m *TiIndicator) SetMalwareFamilyNames(value []string)() {
    m.malwareFamilyNames = value
}
// Sets the networkCidrBlock property value. 
// Parameters:
//  - value : Value to set for the networkCidrBlock property.
func (m *TiIndicator) SetNetworkCidrBlock(value *string)() {
    m.networkCidrBlock = value
}
// Sets the networkDestinationAsn property value. 
// Parameters:
//  - value : Value to set for the networkDestinationAsn property.
func (m *TiIndicator) SetNetworkDestinationAsn(value *int64)() {
    m.networkDestinationAsn = value
}
// Sets the networkDestinationCidrBlock property value. 
// Parameters:
//  - value : Value to set for the networkDestinationCidrBlock property.
func (m *TiIndicator) SetNetworkDestinationCidrBlock(value *string)() {
    m.networkDestinationCidrBlock = value
}
// Sets the networkDestinationIPv4 property value. 
// Parameters:
//  - value : Value to set for the networkDestinationIPv4 property.
func (m *TiIndicator) SetNetworkDestinationIPv4(value *string)() {
    m.networkDestinationIPv4 = value
}
// Sets the networkDestinationIPv6 property value. 
// Parameters:
//  - value : Value to set for the networkDestinationIPv6 property.
func (m *TiIndicator) SetNetworkDestinationIPv6(value *string)() {
    m.networkDestinationIPv6 = value
}
// Sets the networkDestinationPort property value. 
// Parameters:
//  - value : Value to set for the networkDestinationPort property.
func (m *TiIndicator) SetNetworkDestinationPort(value *int32)() {
    m.networkDestinationPort = value
}
// Sets the networkIPv4 property value. 
// Parameters:
//  - value : Value to set for the networkIPv4 property.
func (m *TiIndicator) SetNetworkIPv4(value *string)() {
    m.networkIPv4 = value
}
// Sets the networkIPv6 property value. 
// Parameters:
//  - value : Value to set for the networkIPv6 property.
func (m *TiIndicator) SetNetworkIPv6(value *string)() {
    m.networkIPv6 = value
}
// Sets the networkPort property value. 
// Parameters:
//  - value : Value to set for the networkPort property.
func (m *TiIndicator) SetNetworkPort(value *int32)() {
    m.networkPort = value
}
// Sets the networkProtocol property value. 
// Parameters:
//  - value : Value to set for the networkProtocol property.
func (m *TiIndicator) SetNetworkProtocol(value *int32)() {
    m.networkProtocol = value
}
// Sets the networkSourceAsn property value. 
// Parameters:
//  - value : Value to set for the networkSourceAsn property.
func (m *TiIndicator) SetNetworkSourceAsn(value *int64)() {
    m.networkSourceAsn = value
}
// Sets the networkSourceCidrBlock property value. 
// Parameters:
//  - value : Value to set for the networkSourceCidrBlock property.
func (m *TiIndicator) SetNetworkSourceCidrBlock(value *string)() {
    m.networkSourceCidrBlock = value
}
// Sets the networkSourceIPv4 property value. 
// Parameters:
//  - value : Value to set for the networkSourceIPv4 property.
func (m *TiIndicator) SetNetworkSourceIPv4(value *string)() {
    m.networkSourceIPv4 = value
}
// Sets the networkSourceIPv6 property value. 
// Parameters:
//  - value : Value to set for the networkSourceIPv6 property.
func (m *TiIndicator) SetNetworkSourceIPv6(value *string)() {
    m.networkSourceIPv6 = value
}
// Sets the networkSourcePort property value. 
// Parameters:
//  - value : Value to set for the networkSourcePort property.
func (m *TiIndicator) SetNetworkSourcePort(value *int32)() {
    m.networkSourcePort = value
}
// Sets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
// Parameters:
//  - value : Value to set for the passiveOnly property.
func (m *TiIndicator) SetPassiveOnly(value *bool)() {
    m.passiveOnly = value
}
// Sets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
// Parameters:
//  - value : Value to set for the severity property.
func (m *TiIndicator) SetSeverity(value *int32)() {
    m.severity = value
}
// Sets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
// Parameters:
//  - value : Value to set for the tags property.
func (m *TiIndicator) SetTags(value []string)() {
    m.tags = value
}
// Sets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
// Parameters:
//  - value : Value to set for the targetProduct property.
func (m *TiIndicator) SetTargetProduct(value *string)() {
    m.targetProduct = value
}
// Sets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
// Parameters:
//  - value : Value to set for the threatType property.
func (m *TiIndicator) SetThreatType(value *string)() {
    m.threatType = value
}
// Sets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
// Parameters:
//  - value : Value to set for the tlpLevel property.
func (m *TiIndicator) SetTlpLevel(value *TlpLevel)() {
    m.tlpLevel = value
}
// Sets the url property value. 
// Parameters:
//  - value : Value to set for the url property.
func (m *TiIndicator) SetUrl(value *string)() {
    m.url = value
}
// Sets the userAgent property value. 
// Parameters:
//  - value : Value to set for the userAgent property.
func (m *TiIndicator) SetUserAgent(value *string)() {
    m.userAgent = value
}
