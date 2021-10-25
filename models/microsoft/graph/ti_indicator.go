package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TiIndicator struct {
    Entity
    action *TiAction;
    activityGroupNames []string;
    additionalInformation *string;
    azureTenantId *string;
    confidence *int32;
    description *string;
    diamondModel *DiamondModel;
    domainName *string;
    emailEncoding *string;
    emailLanguage *string;
    emailRecipient *string;
    emailSenderAddress *string;
    emailSenderName *string;
    emailSourceDomain *string;
    emailSourceIpAddress *string;
    emailSubject *string;
    emailXMailer *string;
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    externalId *string;
    fileCompileDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    fileCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    fileHashType *FileHashType;
    fileHashValue *string;
    fileMutexName *string;
    fileName *string;
    filePacker *string;
    filePath *string;
    fileSize *int64;
    fileType *string;
    ingestedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    isActive *bool;
    killChain []string;
    knownFalsePositives *string;
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    malwareFamilyNames []string;
    networkCidrBlock *string;
    networkDestinationAsn *int64;
    networkDestinationCidrBlock *string;
    networkDestinationIPv4 *string;
    networkDestinationIPv6 *string;
    networkDestinationPort *int32;
    networkIPv4 *string;
    networkIPv6 *string;
    networkPort *int32;
    networkProtocol *int32;
    networkSourceAsn *int64;
    networkSourceCidrBlock *string;
    networkSourceIPv4 *string;
    networkSourceIPv6 *string;
    networkSourcePort *int32;
    passiveOnly *bool;
    severity *int32;
    tags []string;
    targetProduct *string;
    threatType *string;
    tlpLevel *TlpLevel;
    url *string;
    userAgent *string;
}
func NewTiIndicator()(*TiIndicator) {
    m := &TiIndicator{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TiIndicator) GetAction()(*TiAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *TiIndicator) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
func (m *TiIndicator) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
func (m *TiIndicator) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
func (m *TiIndicator) GetConfidence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
func (m *TiIndicator) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *TiIndicator) GetDiamondModel()(*DiamondModel) {
    if m == nil {
        return nil
    } else {
        return m.diamondModel
    }
}
func (m *TiIndicator) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
func (m *TiIndicator) GetEmailEncoding()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailEncoding
    }
}
func (m *TiIndicator) GetEmailLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailLanguage
    }
}
func (m *TiIndicator) GetEmailRecipient()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailRecipient
    }
}
func (m *TiIndicator) GetEmailSenderAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSenderAddress
    }
}
func (m *TiIndicator) GetEmailSenderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSenderName
    }
}
func (m *TiIndicator) GetEmailSourceDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSourceDomain
    }
}
func (m *TiIndicator) GetEmailSourceIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSourceIpAddress
    }
}
func (m *TiIndicator) GetEmailSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailSubject
    }
}
func (m *TiIndicator) GetEmailXMailer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailXMailer
    }
}
func (m *TiIndicator) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
func (m *TiIndicator) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *TiIndicator) GetFileCompileDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fileCompileDateTime
    }
}
func (m *TiIndicator) GetFileCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fileCreatedDateTime
    }
}
func (m *TiIndicator) GetFileHashType()(*FileHashType) {
    if m == nil {
        return nil
    } else {
        return m.fileHashType
    }
}
func (m *TiIndicator) GetFileHashValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileHashValue
    }
}
func (m *TiIndicator) GetFileMutexName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileMutexName
    }
}
func (m *TiIndicator) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
func (m *TiIndicator) GetFilePacker()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePacker
    }
}
func (m *TiIndicator) GetFilePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePath
    }
}
func (m *TiIndicator) GetFileSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileSize
    }
}
func (m *TiIndicator) GetFileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileType
    }
}
func (m *TiIndicator) GetIngestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.ingestedDateTime
    }
}
func (m *TiIndicator) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
func (m *TiIndicator) GetKillChain()([]string) {
    if m == nil {
        return nil
    } else {
        return m.killChain
    }
}
func (m *TiIndicator) GetKnownFalsePositives()(*string) {
    if m == nil {
        return nil
    } else {
        return m.knownFalsePositives
    }
}
func (m *TiIndicator) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
func (m *TiIndicator) GetMalwareFamilyNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.malwareFamilyNames
    }
}
func (m *TiIndicator) GetNetworkCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkCidrBlock
    }
}
func (m *TiIndicator) GetNetworkDestinationAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationAsn
    }
}
func (m *TiIndicator) GetNetworkDestinationCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationCidrBlock
    }
}
func (m *TiIndicator) GetNetworkDestinationIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationIPv4
    }
}
func (m *TiIndicator) GetNetworkDestinationIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationIPv6
    }
}
func (m *TiIndicator) GetNetworkDestinationPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkDestinationPort
    }
}
func (m *TiIndicator) GetNetworkIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkIPv4
    }
}
func (m *TiIndicator) GetNetworkIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkIPv6
    }
}
func (m *TiIndicator) GetNetworkPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkPort
    }
}
func (m *TiIndicator) GetNetworkProtocol()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkProtocol
    }
}
func (m *TiIndicator) GetNetworkSourceAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceAsn
    }
}
func (m *TiIndicator) GetNetworkSourceCidrBlock()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceCidrBlock
    }
}
func (m *TiIndicator) GetNetworkSourceIPv4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceIPv4
    }
}
func (m *TiIndicator) GetNetworkSourceIPv6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkSourceIPv6
    }
}
func (m *TiIndicator) GetNetworkSourcePort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkSourcePort
    }
}
func (m *TiIndicator) GetPassiveOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passiveOnly
    }
}
func (m *TiIndicator) GetSeverity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
func (m *TiIndicator) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
func (m *TiIndicator) GetTargetProduct()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetProduct
    }
}
func (m *TiIndicator) GetThreatType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatType
    }
}
func (m *TiIndicator) GetTlpLevel()(*TlpLevel) {
    if m == nil {
        return nil
    } else {
        return m.tlpLevel
    }
}
func (m *TiIndicator) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *TiIndicator) GetUserAgent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userAgent
    }
}
func (m *TiIndicator) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTiAction)
        if err != nil {
            return err
        }
        cast := val.(TiAction)
        m.SetAction(&cast)
        return nil
    }
    res["activityGroupNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetActivityGroupNames(res)
        return nil
    }
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalInformation(val)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfidence(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["diamondModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiamondModel)
        if err != nil {
            return err
        }
        cast := val.(DiamondModel)
        m.SetDiamondModel(&cast)
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDomainName(val)
        return nil
    }
    res["emailEncoding"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailEncoding(val)
        return nil
    }
    res["emailLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailLanguage(val)
        return nil
    }
    res["emailRecipient"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailRecipient(val)
        return nil
    }
    res["emailSenderAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailSenderAddress(val)
        return nil
    }
    res["emailSenderName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailSenderName(val)
        return nil
    }
    res["emailSourceDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailSourceDomain(val)
        return nil
    }
    res["emailSourceIpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailSourceIpAddress(val)
        return nil
    }
    res["emailSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailSubject(val)
        return nil
    }
    res["emailXMailer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailXMailer(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["fileCompileDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetFileCompileDateTime(val)
        return nil
    }
    res["fileCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetFileCreatedDateTime(val)
        return nil
    }
    res["fileHashType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileHashType)
        if err != nil {
            return err
        }
        cast := val.(FileHashType)
        m.SetFileHashType(&cast)
        return nil
    }
    res["fileHashValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileHashValue(val)
        return nil
    }
    res["fileMutexName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileMutexName(val)
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["filePacker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFilePacker(val)
        return nil
    }
    res["filePath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFilePath(val)
        return nil
    }
    res["fileSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFileSize(val)
        return nil
    }
    res["fileType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileType(val)
        return nil
    }
    res["ingestedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetIngestedDateTime(val)
        return nil
    }
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsActive(val)
        return nil
    }
    res["killChain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetKillChain(res)
        return nil
    }
    res["knownFalsePositives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKnownFalsePositives(val)
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastReportedDateTime(val)
        return nil
    }
    res["malwareFamilyNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMalwareFamilyNames(res)
        return nil
    }
    res["networkCidrBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkCidrBlock(val)
        return nil
    }
    res["networkDestinationAsn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNetworkDestinationAsn(val)
        return nil
    }
    res["networkDestinationCidrBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkDestinationCidrBlock(val)
        return nil
    }
    res["networkDestinationIPv4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkDestinationIPv4(val)
        return nil
    }
    res["networkDestinationIPv6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkDestinationIPv6(val)
        return nil
    }
    res["networkDestinationPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNetworkDestinationPort(val)
        return nil
    }
    res["networkIPv4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkIPv4(val)
        return nil
    }
    res["networkIPv6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkIPv6(val)
        return nil
    }
    res["networkPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNetworkPort(val)
        return nil
    }
    res["networkProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNetworkProtocol(val)
        return nil
    }
    res["networkSourceAsn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNetworkSourceAsn(val)
        return nil
    }
    res["networkSourceCidrBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkSourceCidrBlock(val)
        return nil
    }
    res["networkSourceIPv4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkSourceIPv4(val)
        return nil
    }
    res["networkSourceIPv6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkSourceIPv6(val)
        return nil
    }
    res["networkSourcePort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNetworkSourcePort(val)
        return nil
    }
    res["passiveOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPassiveOnly(val)
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSeverity(val)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTags(res)
        return nil
    }
    res["targetProduct"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetProduct(val)
        return nil
    }
    res["threatType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThreatType(val)
        return nil
    }
    res["tlpLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTlpLevel)
        if err != nil {
            return err
        }
        cast := val.(TlpLevel)
        m.SetTlpLevel(&cast)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    res["userAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserAgent(val)
        return nil
    }
    return res
}
func (m *TiIndicator) IsNil()(bool) {
    return m == nil
}
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
func (m *TiIndicator) SetAction(value *TiAction)() {
    m.action = value
}
func (m *TiIndicator) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
func (m *TiIndicator) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
func (m *TiIndicator) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
func (m *TiIndicator) SetConfidence(value *int32)() {
    m.confidence = value
}
func (m *TiIndicator) SetDescription(value *string)() {
    m.description = value
}
func (m *TiIndicator) SetDiamondModel(value *DiamondModel)() {
    m.diamondModel = value
}
func (m *TiIndicator) SetDomainName(value *string)() {
    m.domainName = value
}
func (m *TiIndicator) SetEmailEncoding(value *string)() {
    m.emailEncoding = value
}
func (m *TiIndicator) SetEmailLanguage(value *string)() {
    m.emailLanguage = value
}
func (m *TiIndicator) SetEmailRecipient(value *string)() {
    m.emailRecipient = value
}
func (m *TiIndicator) SetEmailSenderAddress(value *string)() {
    m.emailSenderAddress = value
}
func (m *TiIndicator) SetEmailSenderName(value *string)() {
    m.emailSenderName = value
}
func (m *TiIndicator) SetEmailSourceDomain(value *string)() {
    m.emailSourceDomain = value
}
func (m *TiIndicator) SetEmailSourceIpAddress(value *string)() {
    m.emailSourceIpAddress = value
}
func (m *TiIndicator) SetEmailSubject(value *string)() {
    m.emailSubject = value
}
func (m *TiIndicator) SetEmailXMailer(value *string)() {
    m.emailXMailer = value
}
func (m *TiIndicator) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
func (m *TiIndicator) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *TiIndicator) SetFileCompileDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fileCompileDateTime = value
}
func (m *TiIndicator) SetFileCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fileCreatedDateTime = value
}
func (m *TiIndicator) SetFileHashType(value *FileHashType)() {
    m.fileHashType = value
}
func (m *TiIndicator) SetFileHashValue(value *string)() {
    m.fileHashValue = value
}
func (m *TiIndicator) SetFileMutexName(value *string)() {
    m.fileMutexName = value
}
func (m *TiIndicator) SetFileName(value *string)() {
    m.fileName = value
}
func (m *TiIndicator) SetFilePacker(value *string)() {
    m.filePacker = value
}
func (m *TiIndicator) SetFilePath(value *string)() {
    m.filePath = value
}
func (m *TiIndicator) SetFileSize(value *int64)() {
    m.fileSize = value
}
func (m *TiIndicator) SetFileType(value *string)() {
    m.fileType = value
}
func (m *TiIndicator) SetIngestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.ingestedDateTime = value
}
func (m *TiIndicator) SetIsActive(value *bool)() {
    m.isActive = value
}
func (m *TiIndicator) SetKillChain(value []string)() {
    m.killChain = value
}
func (m *TiIndicator) SetKnownFalsePositives(value *string)() {
    m.knownFalsePositives = value
}
func (m *TiIndicator) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
func (m *TiIndicator) SetMalwareFamilyNames(value []string)() {
    m.malwareFamilyNames = value
}
func (m *TiIndicator) SetNetworkCidrBlock(value *string)() {
    m.networkCidrBlock = value
}
func (m *TiIndicator) SetNetworkDestinationAsn(value *int64)() {
    m.networkDestinationAsn = value
}
func (m *TiIndicator) SetNetworkDestinationCidrBlock(value *string)() {
    m.networkDestinationCidrBlock = value
}
func (m *TiIndicator) SetNetworkDestinationIPv4(value *string)() {
    m.networkDestinationIPv4 = value
}
func (m *TiIndicator) SetNetworkDestinationIPv6(value *string)() {
    m.networkDestinationIPv6 = value
}
func (m *TiIndicator) SetNetworkDestinationPort(value *int32)() {
    m.networkDestinationPort = value
}
func (m *TiIndicator) SetNetworkIPv4(value *string)() {
    m.networkIPv4 = value
}
func (m *TiIndicator) SetNetworkIPv6(value *string)() {
    m.networkIPv6 = value
}
func (m *TiIndicator) SetNetworkPort(value *int32)() {
    m.networkPort = value
}
func (m *TiIndicator) SetNetworkProtocol(value *int32)() {
    m.networkProtocol = value
}
func (m *TiIndicator) SetNetworkSourceAsn(value *int64)() {
    m.networkSourceAsn = value
}
func (m *TiIndicator) SetNetworkSourceCidrBlock(value *string)() {
    m.networkSourceCidrBlock = value
}
func (m *TiIndicator) SetNetworkSourceIPv4(value *string)() {
    m.networkSourceIPv4 = value
}
func (m *TiIndicator) SetNetworkSourceIPv6(value *string)() {
    m.networkSourceIPv6 = value
}
func (m *TiIndicator) SetNetworkSourcePort(value *int32)() {
    m.networkSourcePort = value
}
func (m *TiIndicator) SetPassiveOnly(value *bool)() {
    m.passiveOnly = value
}
func (m *TiIndicator) SetSeverity(value *int32)() {
    m.severity = value
}
func (m *TiIndicator) SetTags(value []string)() {
    m.tags = value
}
func (m *TiIndicator) SetTargetProduct(value *string)() {
    m.targetProduct = value
}
func (m *TiIndicator) SetThreatType(value *string)() {
    m.threatType = value
}
func (m *TiIndicator) SetTlpLevel(value *TlpLevel)() {
    m.tlpLevel = value
}
func (m *TiIndicator) SetUrl(value *string)() {
    m.url = value
}
func (m *TiIndicator) SetUserAgent(value *string)() {
    m.userAgent = value
}
