package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicator provides operations to manage the collection of accessReview entities.
type TiIndicator struct {
    Entity
    // The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
    action *TiAction
    // The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
    activityGroupNames []string
    // A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
    additionalInformation *string
    // Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
    azureTenantId *string
    // An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
    confidence *int32
    // Brief description (100 characters or less) of the threat represented by the indicator. Required.
    description *string
    // The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
    diamondModel *DiamondModel
    // The domainName property
    domainName *string
    // The emailEncoding property
    emailEncoding *string
    // The emailLanguage property
    emailLanguage *string
    // The emailRecipient property
    emailRecipient *string
    // The emailSenderAddress property
    emailSenderAddress *string
    // The emailSenderName property
    emailSenderName *string
    // The emailSourceDomain property
    emailSourceDomain *string
    // The emailSourceIpAddress property
    emailSourceIpAddress *string
    // The emailSubject property
    emailSubject *string
    // The emailXMailer property
    emailXMailer *string
    // DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
    externalId *string
    // The fileCompileDateTime property
    fileCompileDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fileCreatedDateTime property
    fileCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fileHashType property
    fileHashType *FileHashType
    // The fileHashValue property
    fileHashValue *string
    // The fileMutexName property
    fileMutexName *string
    // The fileName property
    fileName *string
    // The filePacker property
    filePacker *string
    // The filePath property
    filePath *string
    // The fileSize property
    fileSize *int64
    // The fileType property
    fileType *string
    // Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    ingestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
    isActive *bool
    // A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
    killChain []string
    // Scenarios in which the indicator may cause false positives. This should be human-readable text.
    knownFalsePositives *string
    // The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
    malwareFamilyNames []string
    // The networkCidrBlock property
    networkCidrBlock *string
    // The networkDestinationAsn property
    networkDestinationAsn *int64
    // The networkDestinationCidrBlock property
    networkDestinationCidrBlock *string
    // The networkDestinationIPv4 property
    networkDestinationIPv4 *string
    // The networkDestinationIPv6 property
    networkDestinationIPv6 *string
    // The networkDestinationPort property
    networkDestinationPort *int32
    // The networkIPv4 property
    networkIPv4 *string
    // The networkIPv6 property
    networkIPv6 *string
    // The networkPort property
    networkPort *int32
    // The networkProtocol property
    networkProtocol *int32
    // The networkSourceAsn property
    networkSourceAsn *int64
    // The networkSourceCidrBlock property
    networkSourceCidrBlock *string
    // The networkSourceIPv4 property
    networkSourceIPv4 *string
    // The networkSourceIPv6 property
    networkSourceIPv6 *string
    // The networkSourcePort property
    networkSourcePort *int32
    // Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
    passiveOnly *bool
    // An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
    severity *int32
    // A JSON array of strings that stores arbitrary tags/keywords.
    tags []string
    // A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
    targetProduct *string
    // Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
    threatType *string
    // Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
    tlpLevel *TlpLevel
    // The url property
    url *string
    // The userAgent property
    userAgent *string
}
// NewTiIndicator instantiates a new tiIndicator and sets the default values.
func NewTiIndicator()(*TiIndicator) {
    m := &TiIndicator{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.tiIndicator";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTiIndicatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicator(), nil
}
// GetAction gets the action property value. The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are: unknown, allow, block, alert. Required.
func (m *TiIndicator) GetAction()(*TiAction) {
    return m.action
}
// GetActivityGroupNames gets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
func (m *TiIndicator) GetActivityGroupNames()([]string) {
    return m.activityGroupNames
}
// GetAdditionalInformation gets the additionalInformation property value. A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
func (m *TiIndicator) GetAdditionalInformation()(*string) {
    return m.additionalInformation
}
// GetAzureTenantId gets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
func (m *TiIndicator) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetConfidence gets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
func (m *TiIndicator) GetConfidence()(*int32) {
    return m.confidence
}
// GetDescription gets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
func (m *TiIndicator) GetDescription()(*string) {
    return m.description
}
// GetDiamondModel gets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
func (m *TiIndicator) GetDiamondModel()(*DiamondModel) {
    return m.diamondModel
}
// GetDomainName gets the domainName property value. The domainName property
func (m *TiIndicator) GetDomainName()(*string) {
    return m.domainName
}
// GetEmailEncoding gets the emailEncoding property value. The emailEncoding property
func (m *TiIndicator) GetEmailEncoding()(*string) {
    return m.emailEncoding
}
// GetEmailLanguage gets the emailLanguage property value. The emailLanguage property
func (m *TiIndicator) GetEmailLanguage()(*string) {
    return m.emailLanguage
}
// GetEmailRecipient gets the emailRecipient property value. The emailRecipient property
func (m *TiIndicator) GetEmailRecipient()(*string) {
    return m.emailRecipient
}
// GetEmailSenderAddress gets the emailSenderAddress property value. The emailSenderAddress property
func (m *TiIndicator) GetEmailSenderAddress()(*string) {
    return m.emailSenderAddress
}
// GetEmailSenderName gets the emailSenderName property value. The emailSenderName property
func (m *TiIndicator) GetEmailSenderName()(*string) {
    return m.emailSenderName
}
// GetEmailSourceDomain gets the emailSourceDomain property value. The emailSourceDomain property
func (m *TiIndicator) GetEmailSourceDomain()(*string) {
    return m.emailSourceDomain
}
// GetEmailSourceIpAddress gets the emailSourceIpAddress property value. The emailSourceIpAddress property
func (m *TiIndicator) GetEmailSourceIpAddress()(*string) {
    return m.emailSourceIpAddress
}
// GetEmailSubject gets the emailSubject property value. The emailSubject property
func (m *TiIndicator) GetEmailSubject()(*string) {
    return m.emailSubject
}
// GetEmailXMailer gets the emailXMailer property value. The emailXMailer property
func (m *TiIndicator) GetEmailXMailer()(*string) {
    return m.emailXMailer
}
// GetExpirationDateTime gets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
func (m *TiIndicator) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetExternalId gets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
func (m *TiIndicator) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TiIndicator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTiAction , m.SetAction)
    res["activityGroupNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetActivityGroupNames)
    res["additionalInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAdditionalInformation)
    res["azureTenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureTenantId)
    res["confidence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConfidence)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["diamondModel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDiamondModel , m.SetDiamondModel)
    res["domainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDomainName)
    res["emailEncoding"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailEncoding)
    res["emailLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailLanguage)
    res["emailRecipient"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailRecipient)
    res["emailSenderAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailSenderAddress)
    res["emailSenderName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailSenderName)
    res["emailSourceDomain"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailSourceDomain)
    res["emailSourceIpAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailSourceIpAddress)
    res["emailSubject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailSubject)
    res["emailXMailer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailXMailer)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["externalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalId)
    res["fileCompileDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetFileCompileDateTime)
    res["fileCreatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetFileCreatedDateTime)
    res["fileHashType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseFileHashType , m.SetFileHashType)
    res["fileHashValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileHashValue)
    res["fileMutexName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileMutexName)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["filePacker"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFilePacker)
    res["filePath"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFilePath)
    res["fileSize"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetFileSize)
    res["fileType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileType)
    res["ingestedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetIngestedDateTime)
    res["isActive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsActive)
    res["killChain"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetKillChain)
    res["knownFalsePositives"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetKnownFalsePositives)
    res["lastReportedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastReportedDateTime)
    res["malwareFamilyNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetMalwareFamilyNames)
    res["networkCidrBlock"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkCidrBlock)
    res["networkDestinationAsn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetNetworkDestinationAsn)
    res["networkDestinationCidrBlock"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkDestinationCidrBlock)
    res["networkDestinationIPv4"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkDestinationIPv4)
    res["networkDestinationIPv6"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkDestinationIPv6)
    res["networkDestinationPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNetworkDestinationPort)
    res["networkIPv4"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkIPv4)
    res["networkIPv6"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkIPv6)
    res["networkPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNetworkPort)
    res["networkProtocol"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNetworkProtocol)
    res["networkSourceAsn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetNetworkSourceAsn)
    res["networkSourceCidrBlock"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkSourceCidrBlock)
    res["networkSourceIPv4"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkSourceIPv4)
    res["networkSourceIPv6"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkSourceIPv6)
    res["networkSourcePort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNetworkSourcePort)
    res["passiveOnly"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPassiveOnly)
    res["severity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSeverity)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTags)
    res["targetProduct"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetProduct)
    res["threatType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetThreatType)
    res["tlpLevel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTlpLevel , m.SetTlpLevel)
    res["url"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUrl)
    res["userAgent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserAgent)
    return res
}
// GetFileCompileDateTime gets the fileCompileDateTime property value. The fileCompileDateTime property
func (m *TiIndicator) GetFileCompileDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.fileCompileDateTime
}
// GetFileCreatedDateTime gets the fileCreatedDateTime property value. The fileCreatedDateTime property
func (m *TiIndicator) GetFileCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.fileCreatedDateTime
}
// GetFileHashType gets the fileHashType property value. The fileHashType property
func (m *TiIndicator) GetFileHashType()(*FileHashType) {
    return m.fileHashType
}
// GetFileHashValue gets the fileHashValue property value. The fileHashValue property
func (m *TiIndicator) GetFileHashValue()(*string) {
    return m.fileHashValue
}
// GetFileMutexName gets the fileMutexName property value. The fileMutexName property
func (m *TiIndicator) GetFileMutexName()(*string) {
    return m.fileMutexName
}
// GetFileName gets the fileName property value. The fileName property
func (m *TiIndicator) GetFileName()(*string) {
    return m.fileName
}
// GetFilePacker gets the filePacker property value. The filePacker property
func (m *TiIndicator) GetFilePacker()(*string) {
    return m.filePacker
}
// GetFilePath gets the filePath property value. The filePath property
func (m *TiIndicator) GetFilePath()(*string) {
    return m.filePath
}
// GetFileSize gets the fileSize property value. The fileSize property
func (m *TiIndicator) GetFileSize()(*int64) {
    return m.fileSize
}
// GetFileType gets the fileType property value. The fileType property
func (m *TiIndicator) GetFileType()(*string) {
    return m.fileType
}
// GetIngestedDateTime gets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) GetIngestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.ingestedDateTime
}
// GetIsActive gets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
func (m *TiIndicator) GetIsActive()(*bool) {
    return m.isActive
}
// GetKillChain gets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
func (m *TiIndicator) GetKillChain()([]string) {
    return m.killChain
}
// GetKnownFalsePositives gets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
func (m *TiIndicator) GetKnownFalsePositives()(*string) {
    return m.knownFalsePositives
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetMalwareFamilyNames gets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
func (m *TiIndicator) GetMalwareFamilyNames()([]string) {
    return m.malwareFamilyNames
}
// GetNetworkCidrBlock gets the networkCidrBlock property value. The networkCidrBlock property
func (m *TiIndicator) GetNetworkCidrBlock()(*string) {
    return m.networkCidrBlock
}
// GetNetworkDestinationAsn gets the networkDestinationAsn property value. The networkDestinationAsn property
func (m *TiIndicator) GetNetworkDestinationAsn()(*int64) {
    return m.networkDestinationAsn
}
// GetNetworkDestinationCidrBlock gets the networkDestinationCidrBlock property value. The networkDestinationCidrBlock property
func (m *TiIndicator) GetNetworkDestinationCidrBlock()(*string) {
    return m.networkDestinationCidrBlock
}
// GetNetworkDestinationIPv4 gets the networkDestinationIPv4 property value. The networkDestinationIPv4 property
func (m *TiIndicator) GetNetworkDestinationIPv4()(*string) {
    return m.networkDestinationIPv4
}
// GetNetworkDestinationIPv6 gets the networkDestinationIPv6 property value. The networkDestinationIPv6 property
func (m *TiIndicator) GetNetworkDestinationIPv6()(*string) {
    return m.networkDestinationIPv6
}
// GetNetworkDestinationPort gets the networkDestinationPort property value. The networkDestinationPort property
func (m *TiIndicator) GetNetworkDestinationPort()(*int32) {
    return m.networkDestinationPort
}
// GetNetworkIPv4 gets the networkIPv4 property value. The networkIPv4 property
func (m *TiIndicator) GetNetworkIPv4()(*string) {
    return m.networkIPv4
}
// GetNetworkIPv6 gets the networkIPv6 property value. The networkIPv6 property
func (m *TiIndicator) GetNetworkIPv6()(*string) {
    return m.networkIPv6
}
// GetNetworkPort gets the networkPort property value. The networkPort property
func (m *TiIndicator) GetNetworkPort()(*int32) {
    return m.networkPort
}
// GetNetworkProtocol gets the networkProtocol property value. The networkProtocol property
func (m *TiIndicator) GetNetworkProtocol()(*int32) {
    return m.networkProtocol
}
// GetNetworkSourceAsn gets the networkSourceAsn property value. The networkSourceAsn property
func (m *TiIndicator) GetNetworkSourceAsn()(*int64) {
    return m.networkSourceAsn
}
// GetNetworkSourceCidrBlock gets the networkSourceCidrBlock property value. The networkSourceCidrBlock property
func (m *TiIndicator) GetNetworkSourceCidrBlock()(*string) {
    return m.networkSourceCidrBlock
}
// GetNetworkSourceIPv4 gets the networkSourceIPv4 property value. The networkSourceIPv4 property
func (m *TiIndicator) GetNetworkSourceIPv4()(*string) {
    return m.networkSourceIPv4
}
// GetNetworkSourceIPv6 gets the networkSourceIPv6 property value. The networkSourceIPv6 property
func (m *TiIndicator) GetNetworkSourceIPv6()(*string) {
    return m.networkSourceIPv6
}
// GetNetworkSourcePort gets the networkSourcePort property value. The networkSourcePort property
func (m *TiIndicator) GetNetworkSourcePort()(*int32) {
    return m.networkSourcePort
}
// GetPassiveOnly gets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
func (m *TiIndicator) GetPassiveOnly()(*bool) {
    return m.passiveOnly
}
// GetSeverity gets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
func (m *TiIndicator) GetSeverity()(*int32) {
    return m.severity
}
// GetTags gets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
func (m *TiIndicator) GetTags()([]string) {
    return m.tags
}
// GetTargetProduct gets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
func (m *TiIndicator) GetTargetProduct()(*string) {
    return m.targetProduct
}
// GetThreatType gets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
func (m *TiIndicator) GetThreatType()(*string) {
    return m.threatType
}
// GetTlpLevel gets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
func (m *TiIndicator) GetTlpLevel()(*TlpLevel) {
    return m.tlpLevel
}
// GetUrl gets the url property value. The url property
func (m *TiIndicator) GetUrl()(*string) {
    return m.url
}
// GetUserAgent gets the userAgent property value. The userAgent property
func (m *TiIndicator) GetUserAgent()(*string) {
    return m.userAgent
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
    m.action = value
}
// SetActivityGroupNames sets the activityGroupNames property value. The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat indicator.
func (m *TiIndicator) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// SetAdditionalInformation sets the additionalInformation property value. A catchall area into which extra data from the indicator not covered by the other tiIndicator properties may be placed. Data placed into additionalInformation will typically not be utilized by the targetProduct security tool.
func (m *TiIndicator) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// SetAzureTenantId sets the azureTenantId property value. Stamped by the system when the indicator is ingested. The Azure Active Directory tenant id of submitting client. Required.
func (m *TiIndicator) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetConfidence sets the confidence property value. An integer representing the confidence the data within the indicator accurately identifies malicious behavior. Acceptable values are 0 – 100 with 100 being the highest.
func (m *TiIndicator) SetConfidence(value *int32)() {
    m.confidence = value
}
// SetDescription sets the description property value. Brief description (100 characters or less) of the threat represented by the indicator. Required.
func (m *TiIndicator) SetDescription(value *string)() {
    m.description = value
}
// SetDiamondModel sets the diamondModel property value. The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability, infrastructure, victim.
func (m *TiIndicator) SetDiamondModel(value *DiamondModel)() {
    m.diamondModel = value
}
// SetDomainName sets the domainName property value. The domainName property
func (m *TiIndicator) SetDomainName(value *string)() {
    m.domainName = value
}
// SetEmailEncoding sets the emailEncoding property value. The emailEncoding property
func (m *TiIndicator) SetEmailEncoding(value *string)() {
    m.emailEncoding = value
}
// SetEmailLanguage sets the emailLanguage property value. The emailLanguage property
func (m *TiIndicator) SetEmailLanguage(value *string)() {
    m.emailLanguage = value
}
// SetEmailRecipient sets the emailRecipient property value. The emailRecipient property
func (m *TiIndicator) SetEmailRecipient(value *string)() {
    m.emailRecipient = value
}
// SetEmailSenderAddress sets the emailSenderAddress property value. The emailSenderAddress property
func (m *TiIndicator) SetEmailSenderAddress(value *string)() {
    m.emailSenderAddress = value
}
// SetEmailSenderName sets the emailSenderName property value. The emailSenderName property
func (m *TiIndicator) SetEmailSenderName(value *string)() {
    m.emailSenderName = value
}
// SetEmailSourceDomain sets the emailSourceDomain property value. The emailSourceDomain property
func (m *TiIndicator) SetEmailSourceDomain(value *string)() {
    m.emailSourceDomain = value
}
// SetEmailSourceIpAddress sets the emailSourceIpAddress property value. The emailSourceIpAddress property
func (m *TiIndicator) SetEmailSourceIpAddress(value *string)() {
    m.emailSourceIpAddress = value
}
// SetEmailSubject sets the emailSubject property value. The emailSubject property
func (m *TiIndicator) SetEmailSubject(value *string)() {
    m.emailSubject = value
}
// SetEmailXMailer sets the emailXMailer property value. The emailXMailer property
func (m *TiIndicator) SetEmailXMailer(value *string)() {
    m.emailXMailer = value
}
// SetExpirationDateTime sets the expirationDateTime property value. DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
func (m *TiIndicator) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetExternalId sets the externalId property value. An identification number that ties the indicator back to the indicator provider’s system (e.g. a foreign key).
func (m *TiIndicator) SetExternalId(value *string)() {
    m.externalId = value
}
// SetFileCompileDateTime sets the fileCompileDateTime property value. The fileCompileDateTime property
func (m *TiIndicator) SetFileCompileDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fileCompileDateTime = value
}
// SetFileCreatedDateTime sets the fileCreatedDateTime property value. The fileCreatedDateTime property
func (m *TiIndicator) SetFileCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fileCreatedDateTime = value
}
// SetFileHashType sets the fileHashType property value. The fileHashType property
func (m *TiIndicator) SetFileHashType(value *FileHashType)() {
    m.fileHashType = value
}
// SetFileHashValue sets the fileHashValue property value. The fileHashValue property
func (m *TiIndicator) SetFileHashValue(value *string)() {
    m.fileHashValue = value
}
// SetFileMutexName sets the fileMutexName property value. The fileMutexName property
func (m *TiIndicator) SetFileMutexName(value *string)() {
    m.fileMutexName = value
}
// SetFileName sets the fileName property value. The fileName property
func (m *TiIndicator) SetFileName(value *string)() {
    m.fileName = value
}
// SetFilePacker sets the filePacker property value. The filePacker property
func (m *TiIndicator) SetFilePacker(value *string)() {
    m.filePacker = value
}
// SetFilePath sets the filePath property value. The filePath property
func (m *TiIndicator) SetFilePath(value *string)() {
    m.filePath = value
}
// SetFileSize sets the fileSize property value. The fileSize property
func (m *TiIndicator) SetFileSize(value *int64)() {
    m.fileSize = value
}
// SetFileType sets the fileType property value. The fileType property
func (m *TiIndicator) SetFileType(value *string)() {
    m.fileType = value
}
// SetIngestedDateTime sets the ingestedDateTime property value. Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) SetIngestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.ingestedDateTime = value
}
// SetIsActive sets the isActive property value. Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
func (m *TiIndicator) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetKillChain sets the killChain property value. A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See ‘killChain values’ below for exact values.
func (m *TiIndicator) SetKillChain(value []string)() {
    m.killChain = value
}
// SetKnownFalsePositives sets the knownFalsePositives property value. Scenarios in which the indicator may cause false positives. This should be human-readable text.
func (m *TiIndicator) SetKnownFalsePositives(value *string)() {
    m.knownFalsePositives = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *TiIndicator) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetMalwareFamilyNames sets the malwareFamilyNames property value. The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family name if at all possible which can be found via the Windows Defender Security Intelligence threat encyclopedia.
func (m *TiIndicator) SetMalwareFamilyNames(value []string)() {
    m.malwareFamilyNames = value
}
// SetNetworkCidrBlock sets the networkCidrBlock property value. The networkCidrBlock property
func (m *TiIndicator) SetNetworkCidrBlock(value *string)() {
    m.networkCidrBlock = value
}
// SetNetworkDestinationAsn sets the networkDestinationAsn property value. The networkDestinationAsn property
func (m *TiIndicator) SetNetworkDestinationAsn(value *int64)() {
    m.networkDestinationAsn = value
}
// SetNetworkDestinationCidrBlock sets the networkDestinationCidrBlock property value. The networkDestinationCidrBlock property
func (m *TiIndicator) SetNetworkDestinationCidrBlock(value *string)() {
    m.networkDestinationCidrBlock = value
}
// SetNetworkDestinationIPv4 sets the networkDestinationIPv4 property value. The networkDestinationIPv4 property
func (m *TiIndicator) SetNetworkDestinationIPv4(value *string)() {
    m.networkDestinationIPv4 = value
}
// SetNetworkDestinationIPv6 sets the networkDestinationIPv6 property value. The networkDestinationIPv6 property
func (m *TiIndicator) SetNetworkDestinationIPv6(value *string)() {
    m.networkDestinationIPv6 = value
}
// SetNetworkDestinationPort sets the networkDestinationPort property value. The networkDestinationPort property
func (m *TiIndicator) SetNetworkDestinationPort(value *int32)() {
    m.networkDestinationPort = value
}
// SetNetworkIPv4 sets the networkIPv4 property value. The networkIPv4 property
func (m *TiIndicator) SetNetworkIPv4(value *string)() {
    m.networkIPv4 = value
}
// SetNetworkIPv6 sets the networkIPv6 property value. The networkIPv6 property
func (m *TiIndicator) SetNetworkIPv6(value *string)() {
    m.networkIPv6 = value
}
// SetNetworkPort sets the networkPort property value. The networkPort property
func (m *TiIndicator) SetNetworkPort(value *int32)() {
    m.networkPort = value
}
// SetNetworkProtocol sets the networkProtocol property value. The networkProtocol property
func (m *TiIndicator) SetNetworkProtocol(value *int32)() {
    m.networkProtocol = value
}
// SetNetworkSourceAsn sets the networkSourceAsn property value. The networkSourceAsn property
func (m *TiIndicator) SetNetworkSourceAsn(value *int64)() {
    m.networkSourceAsn = value
}
// SetNetworkSourceCidrBlock sets the networkSourceCidrBlock property value. The networkSourceCidrBlock property
func (m *TiIndicator) SetNetworkSourceCidrBlock(value *string)() {
    m.networkSourceCidrBlock = value
}
// SetNetworkSourceIPv4 sets the networkSourceIPv4 property value. The networkSourceIPv4 property
func (m *TiIndicator) SetNetworkSourceIPv4(value *string)() {
    m.networkSourceIPv4 = value
}
// SetNetworkSourceIPv6 sets the networkSourceIPv6 property value. The networkSourceIPv6 property
func (m *TiIndicator) SetNetworkSourceIPv6(value *string)() {
    m.networkSourceIPv6 = value
}
// SetNetworkSourcePort sets the networkSourcePort property value. The networkSourcePort property
func (m *TiIndicator) SetNetworkSourcePort(value *int32)() {
    m.networkSourcePort = value
}
// SetPassiveOnly sets the passiveOnly property value. Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security tools will not notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by security products where they will simply log that a match occurred but will not perform the action. Default value is false.
func (m *TiIndicator) SetPassiveOnly(value *bool)() {
    m.passiveOnly = value
}
// SetSeverity sets the severity property value. An integer representing the severity of the malicious behavior identified by the data within the indicator. Acceptable values are 0 – 5 where 5 is the most severe and zero is not severe at all. Default value is 3.
func (m *TiIndicator) SetSeverity(value *int32)() {
    m.severity = value
}
// SetTags sets the tags property value. A JSON array of strings that stores arbitrary tags/keywords.
func (m *TiIndicator) SetTags(value []string)() {
    m.tags = value
}
// SetTargetProduct sets the targetProduct property value. A string value representing a single security product to which the indicator should be applied. Acceptable values are: Azure Sentinel, Microsoft Defender ATP. Required
func (m *TiIndicator) SetTargetProduct(value *string)() {
    m.targetProduct = value
}
// SetThreatType sets the threatType property value. Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS, MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
func (m *TiIndicator) SetThreatType(value *string)() {
    m.threatType = value
}
// SetTlpLevel sets the tlpLevel property value. Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
func (m *TiIndicator) SetTlpLevel(value *TlpLevel)() {
    m.tlpLevel = value
}
// SetUrl sets the url property value. The url property
func (m *TiIndicator) SetUrl(value *string)() {
    m.url = value
}
// SetUserAgent sets the userAgent property value. The userAgent property
func (m *TiIndicator) SetUserAgent(value *string)() {
    m.userAgent = value
}
