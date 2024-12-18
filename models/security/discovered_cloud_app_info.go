package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type DiscoveredCloudAppInfo struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDiscoveredCloudAppInfo instantiates a new DiscoveredCloudAppInfo and sets the default values.
func NewDiscoveredCloudAppInfo()(*DiscoveredCloudAppInfo) {
    m := &DiscoveredCloudAppInfo{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDiscoveredCloudAppInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDiscoveredCloudAppInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDiscoveredCloudAppInfo(), nil
}
// GetCsaStarLevel gets the csaStarLevel property value. The csaStarLevel property
// returns a *AppInfoCsaStarLevel when successful
func (m *DiscoveredCloudAppInfo) GetCsaStarLevel()(*AppInfoCsaStarLevel) {
    val, err := m.GetBackingStore().Get("csaStarLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoCsaStarLevel)
    }
    return nil
}
// GetDataAtRestEncryptionMethod gets the dataAtRestEncryptionMethod property value. The dataAtRestEncryptionMethod property
// returns a *AppInfoDataAtRestEncryptionMethod when successful
func (m *DiscoveredCloudAppInfo) GetDataAtRestEncryptionMethod()(*AppInfoDataAtRestEncryptionMethod) {
    val, err := m.GetBackingStore().Get("dataAtRestEncryptionMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoDataAtRestEncryptionMethod)
    }
    return nil
}
// GetDataCenter gets the dataCenter property value. Indicates the countries or regions in which your data center resides.
// returns a *string when successful
func (m *DiscoveredCloudAppInfo) GetDataCenter()(*string) {
    val, err := m.GetBackingStore().Get("dataCenter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDataRetentionPolicy gets the dataRetentionPolicy property value. The dataRetentionPolicy property
// returns a *AppInfoDataRetentionPolicy when successful
func (m *DiscoveredCloudAppInfo) GetDataRetentionPolicy()(*AppInfoDataRetentionPolicy) {
    val, err := m.GetBackingStore().Get("dataRetentionPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoDataRetentionPolicy)
    }
    return nil
}
// GetDataTypes gets the dataTypes property value. The dataTypes property
// returns a *AppInfoUploadedDataTypes when successful
func (m *DiscoveredCloudAppInfo) GetDataTypes()(*AppInfoUploadedDataTypes) {
    val, err := m.GetBackingStore().Get("dataTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoUploadedDataTypes)
    }
    return nil
}
// GetDomainRegistrationDateTime gets the domainRegistrationDateTime property value. Indicates the date when the app domain was registered.
// returns a *Time when successful
func (m *DiscoveredCloudAppInfo) GetDomainRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("domainRegistrationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEncryptionProtocol gets the encryptionProtocol property value. The encryptionProtocol property
// returns a *AppInfoEncryptionProtocol when successful
func (m *DiscoveredCloudAppInfo) GetEncryptionProtocol()(*AppInfoEncryptionProtocol) {
    val, err := m.GetBackingStore().Get("encryptionProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoEncryptionProtocol)
    }
    return nil
}
// GetFedRampLevel gets the fedRampLevel property value. The fedRampLevel property
// returns a *AppInfoFedRampLevel when successful
func (m *DiscoveredCloudAppInfo) GetFedRampLevel()(*AppInfoFedRampLevel) {
    val, err := m.GetBackingStore().Get("fedRampLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoFedRampLevel)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DiscoveredCloudAppInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["csaStarLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoCsaStarLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCsaStarLevel(val.(*AppInfoCsaStarLevel))
        }
        return nil
    }
    res["dataAtRestEncryptionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoDataAtRestEncryptionMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataAtRestEncryptionMethod(val.(*AppInfoDataAtRestEncryptionMethod))
        }
        return nil
    }
    res["dataCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataCenter(val)
        }
        return nil
    }
    res["dataRetentionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoDataRetentionPolicy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataRetentionPolicy(val.(*AppInfoDataRetentionPolicy))
        }
        return nil
    }
    res["dataTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoUploadedDataTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataTypes(val.(*AppInfoUploadedDataTypes))
        }
        return nil
    }
    res["domainRegistrationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainRegistrationDateTime(val)
        }
        return nil
    }
    res["encryptionProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoEncryptionProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionProtocol(val.(*AppInfoEncryptionProtocol))
        }
        return nil
    }
    res["fedRampLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoFedRampLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFedRampLevel(val.(*AppInfoFedRampLevel))
        }
        return nil
    }
    res["founded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFounded(val)
        }
        return nil
    }
    res["gdprReadinessStatement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGdprReadinessStatement(val)
        }
        return nil
    }
    res["headquarters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeadquarters(val)
        }
        return nil
    }
    res["holding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoHolding)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHolding(val.(*AppInfoHolding))
        }
        return nil
    }
    res["hostingCompany"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostingCompany(val)
        }
        return nil
    }
    res["isAdminAuditTrail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAdminAuditTrail(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isCobitCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCobitCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isCoppaCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCoppaCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isDataAuditTrail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDataAuditTrail(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isDataClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDataClassification(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isDataOwnership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDataOwnership(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isDisasterRecoveryPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDisasterRecoveryPlan(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isDmca"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDmca(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isFerpaCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFerpaCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isFfiecCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFfiecCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isFileSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFileSharing(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isFinraCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFinraCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isFismaCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFismaCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGaapCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGaapCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprDataProtectionImpactAssessment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprDataProtectionImpactAssessment(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprDataProtectionOfficer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprDataProtectionOfficer(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprDataProtectionSecureCrossBorderDataTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprDataProtectionSecureCrossBorderDataTransfer(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprImpactAssessment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprImpactAssessment(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprLawfulBasisForProcessing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprLawfulBasisForProcessing(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprReportDataBreaches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprReportDataBreaches(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightsRelatedToAutomatedDecisionMaking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightsRelatedToAutomatedDecisionMaking(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToAccess(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToBeInformed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToBeInformed(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToDataPortablility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToDataPortablility(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToErasure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToErasure(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToObject(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToRectification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToRectification(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprRightToRestrictionOfProcessing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprRightToRestrictionOfProcessing(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGdprSecureCrossBorderDataControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGdprSecureCrossBorderDataControl(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isGlbaCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlbaCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHipaaCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHipaaCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHitrustCsfCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHitrustCsfCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHttpSecurityHeadersContentSecurityPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpSecurityHeadersContentSecurityPolicy(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHttpSecurityHeadersStrictTransportSecurity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpSecurityHeadersStrictTransportSecurity(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHttpSecurityHeadersXContentTypeOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpSecurityHeadersXContentTypeOptions(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHttpSecurityHeadersXFrameOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpSecurityHeadersXFrameOptions(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isHttpSecurityHeadersXXssProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpSecurityHeadersXXssProtection(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isIpAddressRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsIpAddressRestriction(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isIsae3402Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsIsae3402Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isIso27001Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsIso27001Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isIso27017Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsIso27017Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isIso27018Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsIso27018Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isItarCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsItarCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isMultiFactorAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMultiFactorAuthentication(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPasswordPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordPolicy(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPasswordPolicyChangePasswordPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordPolicyChangePasswordPeriod(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPasswordPolicyCharacterCombination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordPolicyCharacterCombination(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPasswordPolicyPasswordHistoryAndReuse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordPolicyPasswordHistoryAndReuse(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPasswordPolicyPasswordLengthLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordPolicyPasswordLengthLimit(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPasswordPolicyPersonalInformationUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordPolicyPersonalInformationUse(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPenetrationTesting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPenetrationTesting(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isPrivacyShieldCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivacyShieldCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isRememberPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRememberPassword(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isRequiresUserAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequiresUserAuthentication(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSoc1Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSoc1Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSoc2Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSoc2Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSoc3Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSoc3Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSoxCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSoxCompliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSp80053Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSp80053Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSsae16Compliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsae16Compliant(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isSupportsSaml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupportsSaml(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isTrustedCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTrustedCertificate(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isUserAuditTrail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUserAuditTrail(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isUserCanUploadData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUserCanUploadData(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isUserRolesSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUserRolesSupport(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["isValidCertificateName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppInfoState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidCertificateName(val.(*CloudAppInfoState))
        }
        return nil
    }
    res["latestBreachDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestBreachDateTime(val)
        }
        return nil
    }
    res["logonUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonUrls(val)
        }
        return nil
    }
    res["pciDssVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppInfoPciDssVersion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPciDssVersion(val.(*AppInfoPciDssVersion))
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val)
        }
        return nil
    }
    return res
}
// GetFounded gets the founded property value. Indicates the year that the specific app vendor was established.
// returns a *int32 when successful
func (m *DiscoveredCloudAppInfo) GetFounded()(*int32) {
    val, err := m.GetBackingStore().Get("founded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetGdprReadinessStatement gets the gdprReadinessStatement property value. Indicates the GDPR readiness of the app in relation to policies app provides to safeguard personal user data.
// returns a *string when successful
func (m *DiscoveredCloudAppInfo) GetGdprReadinessStatement()(*string) {
    val, err := m.GetBackingStore().Get("gdprReadinessStatement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHeadquarters gets the headquarters property value. Indicates the location of the headquarters of the app.
// returns a *string when successful
func (m *DiscoveredCloudAppInfo) GetHeadquarters()(*string) {
    val, err := m.GetBackingStore().Get("headquarters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHolding gets the holding property value. The holding property
// returns a *AppInfoHolding when successful
func (m *DiscoveredCloudAppInfo) GetHolding()(*AppInfoHolding) {
    val, err := m.GetBackingStore().Get("holding")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoHolding)
    }
    return nil
}
// GetHostingCompany gets the hostingCompany property value. Indicates the company name that provides hosting services for the app.
// returns a *string when successful
func (m *DiscoveredCloudAppInfo) GetHostingCompany()(*string) {
    val, err := m.GetBackingStore().Get("hostingCompany")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsAdminAuditTrail gets the isAdminAuditTrail property value. The isAdminAuditTrail property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsAdminAuditTrail()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isAdminAuditTrail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsCobitCompliant gets the isCobitCompliant property value. The isCobitCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsCobitCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isCobitCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsCoppaCompliant gets the isCoppaCompliant property value. The isCoppaCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsCoppaCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isCoppaCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsDataAuditTrail gets the isDataAuditTrail property value. The isDataAuditTrail property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsDataAuditTrail()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isDataAuditTrail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsDataClassification gets the isDataClassification property value. The isDataClassification property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsDataClassification()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isDataClassification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsDataOwnership gets the isDataOwnership property value. The isDataOwnership property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsDataOwnership()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isDataOwnership")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsDisasterRecoveryPlan gets the isDisasterRecoveryPlan property value. The isDisasterRecoveryPlan property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsDisasterRecoveryPlan()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isDisasterRecoveryPlan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsDmca gets the isDmca property value. The isDmca property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsDmca()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isDmca")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsFerpaCompliant gets the isFerpaCompliant property value. The isFerpaCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsFerpaCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isFerpaCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsFfiecCompliant gets the isFfiecCompliant property value. The isFfiecCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsFfiecCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isFfiecCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsFileSharing gets the isFileSharing property value. The isFileSharing property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsFileSharing()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isFileSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsFinraCompliant gets the isFinraCompliant property value. The isFinraCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsFinraCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isFinraCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsFismaCompliant gets the isFismaCompliant property value. The isFismaCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsFismaCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isFismaCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGaapCompliant gets the isGaapCompliant property value. The isGaapCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGaapCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGaapCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprDataProtectionImpactAssessment gets the isGdprDataProtectionImpactAssessment property value. The isGdprDataProtectionImpactAssessment property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprDataProtectionImpactAssessment()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprDataProtectionImpactAssessment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprDataProtectionOfficer gets the isGdprDataProtectionOfficer property value. The isGdprDataProtectionOfficer property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprDataProtectionOfficer()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprDataProtectionOfficer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprDataProtectionSecureCrossBorderDataTransfer gets the isGdprDataProtectionSecureCrossBorderDataTransfer property value. The isGdprDataProtectionSecureCrossBorderDataTransfer property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprDataProtectionSecureCrossBorderDataTransfer()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprDataProtectionSecureCrossBorderDataTransfer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprImpactAssessment gets the isGdprImpactAssessment property value. The isGdprImpactAssessment property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprImpactAssessment()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprImpactAssessment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprLawfulBasisForProcessing gets the isGdprLawfulBasisForProcessing property value. The isGdprLawfulBasisForProcessing property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprLawfulBasisForProcessing()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprLawfulBasisForProcessing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprReportDataBreaches gets the isGdprReportDataBreaches property value. The isGdprReportDataBreaches property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprReportDataBreaches()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprReportDataBreaches")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightsRelatedToAutomatedDecisionMaking gets the isGdprRightsRelatedToAutomatedDecisionMaking property value. The isGdprRightsRelatedToAutomatedDecisionMaking property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightsRelatedToAutomatedDecisionMaking()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightsRelatedToAutomatedDecisionMaking")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToAccess gets the isGdprRightToAccess property value. The isGdprRightToAccess property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToAccess()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToBeInformed gets the isGdprRightToBeInformed property value. The isGdprRightToBeInformed property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToBeInformed()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToBeInformed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToDataPortablility gets the isGdprRightToDataPortablility property value. The isGdprRightToDataPortablility property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToDataPortablility()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToDataPortablility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToErasure gets the isGdprRightToErasure property value. The isGdprRightToErasure property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToErasure()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToErasure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToObject gets the isGdprRightToObject property value. The isGdprRightToObject property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToObject()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToObject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToRectification gets the isGdprRightToRectification property value. The isGdprRightToRectification property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToRectification()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToRectification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprRightToRestrictionOfProcessing gets the isGdprRightToRestrictionOfProcessing property value. The isGdprRightToRestrictionOfProcessing property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprRightToRestrictionOfProcessing()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprRightToRestrictionOfProcessing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGdprSecureCrossBorderDataControl gets the isGdprSecureCrossBorderDataControl property value. The isGdprSecureCrossBorderDataControl property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGdprSecureCrossBorderDataControl()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGdprSecureCrossBorderDataControl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsGlbaCompliant gets the isGlbaCompliant property value. The isGlbaCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsGlbaCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isGlbaCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHipaaCompliant gets the isHipaaCompliant property value. The isHipaaCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHipaaCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHipaaCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHitrustCsfCompliant gets the isHitrustCsfCompliant property value. The isHitrustCsfCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHitrustCsfCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHitrustCsfCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHttpSecurityHeadersContentSecurityPolicy gets the isHttpSecurityHeadersContentSecurityPolicy property value. The isHttpSecurityHeadersContentSecurityPolicy property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHttpSecurityHeadersContentSecurityPolicy()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHttpSecurityHeadersContentSecurityPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHttpSecurityHeadersStrictTransportSecurity gets the isHttpSecurityHeadersStrictTransportSecurity property value. The isHttpSecurityHeadersStrictTransportSecurity property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHttpSecurityHeadersStrictTransportSecurity()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHttpSecurityHeadersStrictTransportSecurity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHttpSecurityHeadersXContentTypeOptions gets the isHttpSecurityHeadersXContentTypeOptions property value. The isHttpSecurityHeadersXContentTypeOptions property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHttpSecurityHeadersXContentTypeOptions()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHttpSecurityHeadersXContentTypeOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHttpSecurityHeadersXFrameOptions gets the isHttpSecurityHeadersXFrameOptions property value. The isHttpSecurityHeadersXFrameOptions property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHttpSecurityHeadersXFrameOptions()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHttpSecurityHeadersXFrameOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsHttpSecurityHeadersXXssProtection gets the isHttpSecurityHeadersXXssProtection property value. The isHttpSecurityHeadersXXssProtection property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsHttpSecurityHeadersXXssProtection()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isHttpSecurityHeadersXXssProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsIpAddressRestriction gets the isIpAddressRestriction property value. The isIpAddressRestriction property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsIpAddressRestriction()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isIpAddressRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsIsae3402Compliant gets the isIsae3402Compliant property value. The isIsae3402Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsIsae3402Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isIsae3402Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsIso27001Compliant gets the isIso27001Compliant property value. The isIso27001Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsIso27001Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isIso27001Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsIso27017Compliant gets the isIso27017Compliant property value. The isIso27017Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsIso27017Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isIso27017Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsIso27018Compliant gets the isIso27018Compliant property value. The isIso27018Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsIso27018Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isIso27018Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsItarCompliant gets the isItarCompliant property value. The isItarCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsItarCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isItarCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsMultiFactorAuthentication gets the isMultiFactorAuthentication property value. The isMultiFactorAuthentication property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsMultiFactorAuthentication()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isMultiFactorAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPasswordPolicy gets the isPasswordPolicy property value. The isPasswordPolicy property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPasswordPolicy()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPasswordPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPasswordPolicyChangePasswordPeriod gets the isPasswordPolicyChangePasswordPeriod property value. The isPasswordPolicyChangePasswordPeriod property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPasswordPolicyChangePasswordPeriod()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPasswordPolicyChangePasswordPeriod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPasswordPolicyCharacterCombination gets the isPasswordPolicyCharacterCombination property value. The isPasswordPolicyCharacterCombination property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPasswordPolicyCharacterCombination()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPasswordPolicyCharacterCombination")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPasswordPolicyPasswordHistoryAndReuse gets the isPasswordPolicyPasswordHistoryAndReuse property value. The isPasswordPolicyPasswordHistoryAndReuse property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPasswordPolicyPasswordHistoryAndReuse()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPasswordPolicyPasswordHistoryAndReuse")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPasswordPolicyPasswordLengthLimit gets the isPasswordPolicyPasswordLengthLimit property value. The isPasswordPolicyPasswordLengthLimit property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPasswordPolicyPasswordLengthLimit()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPasswordPolicyPasswordLengthLimit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPasswordPolicyPersonalInformationUse gets the isPasswordPolicyPersonalInformationUse property value. The isPasswordPolicyPersonalInformationUse property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPasswordPolicyPersonalInformationUse()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPasswordPolicyPersonalInformationUse")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPenetrationTesting gets the isPenetrationTesting property value. The isPenetrationTesting property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPenetrationTesting()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPenetrationTesting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsPrivacyShieldCompliant gets the isPrivacyShieldCompliant property value. The isPrivacyShieldCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsPrivacyShieldCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isPrivacyShieldCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsRememberPassword gets the isRememberPassword property value. The isRememberPassword property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsRememberPassword()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isRememberPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsRequiresUserAuthentication gets the isRequiresUserAuthentication property value. The isRequiresUserAuthentication property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsRequiresUserAuthentication()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isRequiresUserAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSoc1Compliant gets the isSoc1Compliant property value. The isSoc1Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSoc1Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSoc1Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSoc2Compliant gets the isSoc2Compliant property value. The isSoc2Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSoc2Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSoc2Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSoc3Compliant gets the isSoc3Compliant property value. The isSoc3Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSoc3Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSoc3Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSoxCompliant gets the isSoxCompliant property value. The isSoxCompliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSoxCompliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSoxCompliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSp80053Compliant gets the isSp80053Compliant property value. The isSp80053Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSp80053Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSp80053Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSsae16Compliant gets the isSsae16Compliant property value. The isSsae16Compliant property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSsae16Compliant()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSsae16Compliant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsSupportsSaml gets the isSupportsSaml property value. The isSupportsSaml property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsSupportsSaml()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isSupportsSaml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsTrustedCertificate gets the isTrustedCertificate property value. The isTrustedCertificate property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsTrustedCertificate()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isTrustedCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsUserAuditTrail gets the isUserAuditTrail property value. The isUserAuditTrail property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsUserAuditTrail()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isUserAuditTrail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsUserCanUploadData gets the isUserCanUploadData property value. The isUserCanUploadData property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsUserCanUploadData()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isUserCanUploadData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsUserRolesSupport gets the isUserRolesSupport property value. The isUserRolesSupport property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsUserRolesSupport()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isUserRolesSupport")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetIsValidCertificateName gets the isValidCertificateName property value. The isValidCertificateName property
// returns a *CloudAppInfoState when successful
func (m *DiscoveredCloudAppInfo) GetIsValidCertificateName()(*CloudAppInfoState) {
    val, err := m.GetBackingStore().Get("isValidCertificateName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudAppInfoState)
    }
    return nil
}
// GetLatestBreachDateTime gets the latestBreachDateTime property value. Indicates the last date of the data breach for the company.
// returns a *Time when successful
func (m *DiscoveredCloudAppInfo) GetLatestBreachDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("latestBreachDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLogonUrls gets the logonUrls property value. Indicates the URL that users can use to sign into the app.
// returns a *string when successful
func (m *DiscoveredCloudAppInfo) GetLogonUrls()(*string) {
    val, err := m.GetBackingStore().Get("logonUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPciDssVersion gets the pciDssVersion property value. The pciDssVersion property
// returns a *AppInfoPciDssVersion when successful
func (m *DiscoveredCloudAppInfo) GetPciDssVersion()(*AppInfoPciDssVersion) {
    val, err := m.GetBackingStore().Get("pciDssVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppInfoPciDssVersion)
    }
    return nil
}
// GetVendorEscaped gets the vendor property value. Indicates the app vendor.
// returns a *string when successful
func (m *DiscoveredCloudAppInfo) GetVendorEscaped()(*string) {
    val, err := m.GetBackingStore().Get("vendorEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DiscoveredCloudAppInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCsaStarLevel() != nil {
        cast := (*m.GetCsaStarLevel()).String()
        err = writer.WriteStringValue("csaStarLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDataAtRestEncryptionMethod() != nil {
        cast := (*m.GetDataAtRestEncryptionMethod()).String()
        err = writer.WriteStringValue("dataAtRestEncryptionMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dataCenter", m.GetDataCenter())
        if err != nil {
            return err
        }
    }
    if m.GetDataRetentionPolicy() != nil {
        cast := (*m.GetDataRetentionPolicy()).String()
        err = writer.WriteStringValue("dataRetentionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDataTypes() != nil {
        cast := (*m.GetDataTypes()).String()
        err = writer.WriteStringValue("dataTypes", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("domainRegistrationDateTime", m.GetDomainRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionProtocol() != nil {
        cast := (*m.GetEncryptionProtocol()).String()
        err = writer.WriteStringValue("encryptionProtocol", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFedRampLevel() != nil {
        cast := (*m.GetFedRampLevel()).String()
        err = writer.WriteStringValue("fedRampLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("founded", m.GetFounded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("gdprReadinessStatement", m.GetGdprReadinessStatement())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headquarters", m.GetHeadquarters())
        if err != nil {
            return err
        }
    }
    if m.GetHolding() != nil {
        cast := (*m.GetHolding()).String()
        err = writer.WriteStringValue("holding", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hostingCompany", m.GetHostingCompany())
        if err != nil {
            return err
        }
    }
    if m.GetIsAdminAuditTrail() != nil {
        cast := (*m.GetIsAdminAuditTrail()).String()
        err = writer.WriteStringValue("isAdminAuditTrail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsCobitCompliant() != nil {
        cast := (*m.GetIsCobitCompliant()).String()
        err = writer.WriteStringValue("isCobitCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsCoppaCompliant() != nil {
        cast := (*m.GetIsCoppaCompliant()).String()
        err = writer.WriteStringValue("isCoppaCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsDataAuditTrail() != nil {
        cast := (*m.GetIsDataAuditTrail()).String()
        err = writer.WriteStringValue("isDataAuditTrail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsDataClassification() != nil {
        cast := (*m.GetIsDataClassification()).String()
        err = writer.WriteStringValue("isDataClassification", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsDataOwnership() != nil {
        cast := (*m.GetIsDataOwnership()).String()
        err = writer.WriteStringValue("isDataOwnership", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsDisasterRecoveryPlan() != nil {
        cast := (*m.GetIsDisasterRecoveryPlan()).String()
        err = writer.WriteStringValue("isDisasterRecoveryPlan", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsDmca() != nil {
        cast := (*m.GetIsDmca()).String()
        err = writer.WriteStringValue("isDmca", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsFerpaCompliant() != nil {
        cast := (*m.GetIsFerpaCompliant()).String()
        err = writer.WriteStringValue("isFerpaCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsFfiecCompliant() != nil {
        cast := (*m.GetIsFfiecCompliant()).String()
        err = writer.WriteStringValue("isFfiecCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsFileSharing() != nil {
        cast := (*m.GetIsFileSharing()).String()
        err = writer.WriteStringValue("isFileSharing", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsFinraCompliant() != nil {
        cast := (*m.GetIsFinraCompliant()).String()
        err = writer.WriteStringValue("isFinraCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsFismaCompliant() != nil {
        cast := (*m.GetIsFismaCompliant()).String()
        err = writer.WriteStringValue("isFismaCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGaapCompliant() != nil {
        cast := (*m.GetIsGaapCompliant()).String()
        err = writer.WriteStringValue("isGaapCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprDataProtectionImpactAssessment() != nil {
        cast := (*m.GetIsGdprDataProtectionImpactAssessment()).String()
        err = writer.WriteStringValue("isGdprDataProtectionImpactAssessment", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprDataProtectionOfficer() != nil {
        cast := (*m.GetIsGdprDataProtectionOfficer()).String()
        err = writer.WriteStringValue("isGdprDataProtectionOfficer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprDataProtectionSecureCrossBorderDataTransfer() != nil {
        cast := (*m.GetIsGdprDataProtectionSecureCrossBorderDataTransfer()).String()
        err = writer.WriteStringValue("isGdprDataProtectionSecureCrossBorderDataTransfer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprImpactAssessment() != nil {
        cast := (*m.GetIsGdprImpactAssessment()).String()
        err = writer.WriteStringValue("isGdprImpactAssessment", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprLawfulBasisForProcessing() != nil {
        cast := (*m.GetIsGdprLawfulBasisForProcessing()).String()
        err = writer.WriteStringValue("isGdprLawfulBasisForProcessing", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprReportDataBreaches() != nil {
        cast := (*m.GetIsGdprReportDataBreaches()).String()
        err = writer.WriteStringValue("isGdprReportDataBreaches", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightsRelatedToAutomatedDecisionMaking() != nil {
        cast := (*m.GetIsGdprRightsRelatedToAutomatedDecisionMaking()).String()
        err = writer.WriteStringValue("isGdprRightsRelatedToAutomatedDecisionMaking", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToAccess() != nil {
        cast := (*m.GetIsGdprRightToAccess()).String()
        err = writer.WriteStringValue("isGdprRightToAccess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToBeInformed() != nil {
        cast := (*m.GetIsGdprRightToBeInformed()).String()
        err = writer.WriteStringValue("isGdprRightToBeInformed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToDataPortablility() != nil {
        cast := (*m.GetIsGdprRightToDataPortablility()).String()
        err = writer.WriteStringValue("isGdprRightToDataPortablility", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToErasure() != nil {
        cast := (*m.GetIsGdprRightToErasure()).String()
        err = writer.WriteStringValue("isGdprRightToErasure", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToObject() != nil {
        cast := (*m.GetIsGdprRightToObject()).String()
        err = writer.WriteStringValue("isGdprRightToObject", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToRectification() != nil {
        cast := (*m.GetIsGdprRightToRectification()).String()
        err = writer.WriteStringValue("isGdprRightToRectification", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprRightToRestrictionOfProcessing() != nil {
        cast := (*m.GetIsGdprRightToRestrictionOfProcessing()).String()
        err = writer.WriteStringValue("isGdprRightToRestrictionOfProcessing", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGdprSecureCrossBorderDataControl() != nil {
        cast := (*m.GetIsGdprSecureCrossBorderDataControl()).String()
        err = writer.WriteStringValue("isGdprSecureCrossBorderDataControl", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsGlbaCompliant() != nil {
        cast := (*m.GetIsGlbaCompliant()).String()
        err = writer.WriteStringValue("isGlbaCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHipaaCompliant() != nil {
        cast := (*m.GetIsHipaaCompliant()).String()
        err = writer.WriteStringValue("isHipaaCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHitrustCsfCompliant() != nil {
        cast := (*m.GetIsHitrustCsfCompliant()).String()
        err = writer.WriteStringValue("isHitrustCsfCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHttpSecurityHeadersContentSecurityPolicy() != nil {
        cast := (*m.GetIsHttpSecurityHeadersContentSecurityPolicy()).String()
        err = writer.WriteStringValue("isHttpSecurityHeadersContentSecurityPolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHttpSecurityHeadersStrictTransportSecurity() != nil {
        cast := (*m.GetIsHttpSecurityHeadersStrictTransportSecurity()).String()
        err = writer.WriteStringValue("isHttpSecurityHeadersStrictTransportSecurity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHttpSecurityHeadersXContentTypeOptions() != nil {
        cast := (*m.GetIsHttpSecurityHeadersXContentTypeOptions()).String()
        err = writer.WriteStringValue("isHttpSecurityHeadersXContentTypeOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHttpSecurityHeadersXFrameOptions() != nil {
        cast := (*m.GetIsHttpSecurityHeadersXFrameOptions()).String()
        err = writer.WriteStringValue("isHttpSecurityHeadersXFrameOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsHttpSecurityHeadersXXssProtection() != nil {
        cast := (*m.GetIsHttpSecurityHeadersXXssProtection()).String()
        err = writer.WriteStringValue("isHttpSecurityHeadersXXssProtection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsIpAddressRestriction() != nil {
        cast := (*m.GetIsIpAddressRestriction()).String()
        err = writer.WriteStringValue("isIpAddressRestriction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsIsae3402Compliant() != nil {
        cast := (*m.GetIsIsae3402Compliant()).String()
        err = writer.WriteStringValue("isIsae3402Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsIso27001Compliant() != nil {
        cast := (*m.GetIsIso27001Compliant()).String()
        err = writer.WriteStringValue("isIso27001Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsIso27017Compliant() != nil {
        cast := (*m.GetIsIso27017Compliant()).String()
        err = writer.WriteStringValue("isIso27017Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsIso27018Compliant() != nil {
        cast := (*m.GetIsIso27018Compliant()).String()
        err = writer.WriteStringValue("isIso27018Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsItarCompliant() != nil {
        cast := (*m.GetIsItarCompliant()).String()
        err = writer.WriteStringValue("isItarCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsMultiFactorAuthentication() != nil {
        cast := (*m.GetIsMultiFactorAuthentication()).String()
        err = writer.WriteStringValue("isMultiFactorAuthentication", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPasswordPolicy() != nil {
        cast := (*m.GetIsPasswordPolicy()).String()
        err = writer.WriteStringValue("isPasswordPolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPasswordPolicyChangePasswordPeriod() != nil {
        cast := (*m.GetIsPasswordPolicyChangePasswordPeriod()).String()
        err = writer.WriteStringValue("isPasswordPolicyChangePasswordPeriod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPasswordPolicyCharacterCombination() != nil {
        cast := (*m.GetIsPasswordPolicyCharacterCombination()).String()
        err = writer.WriteStringValue("isPasswordPolicyCharacterCombination", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPasswordPolicyPasswordHistoryAndReuse() != nil {
        cast := (*m.GetIsPasswordPolicyPasswordHistoryAndReuse()).String()
        err = writer.WriteStringValue("isPasswordPolicyPasswordHistoryAndReuse", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPasswordPolicyPasswordLengthLimit() != nil {
        cast := (*m.GetIsPasswordPolicyPasswordLengthLimit()).String()
        err = writer.WriteStringValue("isPasswordPolicyPasswordLengthLimit", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPasswordPolicyPersonalInformationUse() != nil {
        cast := (*m.GetIsPasswordPolicyPersonalInformationUse()).String()
        err = writer.WriteStringValue("isPasswordPolicyPersonalInformationUse", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPenetrationTesting() != nil {
        cast := (*m.GetIsPenetrationTesting()).String()
        err = writer.WriteStringValue("isPenetrationTesting", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsPrivacyShieldCompliant() != nil {
        cast := (*m.GetIsPrivacyShieldCompliant()).String()
        err = writer.WriteStringValue("isPrivacyShieldCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsRememberPassword() != nil {
        cast := (*m.GetIsRememberPassword()).String()
        err = writer.WriteStringValue("isRememberPassword", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsRequiresUserAuthentication() != nil {
        cast := (*m.GetIsRequiresUserAuthentication()).String()
        err = writer.WriteStringValue("isRequiresUserAuthentication", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSoc1Compliant() != nil {
        cast := (*m.GetIsSoc1Compliant()).String()
        err = writer.WriteStringValue("isSoc1Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSoc2Compliant() != nil {
        cast := (*m.GetIsSoc2Compliant()).String()
        err = writer.WriteStringValue("isSoc2Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSoc3Compliant() != nil {
        cast := (*m.GetIsSoc3Compliant()).String()
        err = writer.WriteStringValue("isSoc3Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSoxCompliant() != nil {
        cast := (*m.GetIsSoxCompliant()).String()
        err = writer.WriteStringValue("isSoxCompliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSp80053Compliant() != nil {
        cast := (*m.GetIsSp80053Compliant()).String()
        err = writer.WriteStringValue("isSp80053Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSsae16Compliant() != nil {
        cast := (*m.GetIsSsae16Compliant()).String()
        err = writer.WriteStringValue("isSsae16Compliant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsSupportsSaml() != nil {
        cast := (*m.GetIsSupportsSaml()).String()
        err = writer.WriteStringValue("isSupportsSaml", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsTrustedCertificate() != nil {
        cast := (*m.GetIsTrustedCertificate()).String()
        err = writer.WriteStringValue("isTrustedCertificate", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsUserAuditTrail() != nil {
        cast := (*m.GetIsUserAuditTrail()).String()
        err = writer.WriteStringValue("isUserAuditTrail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsUserCanUploadData() != nil {
        cast := (*m.GetIsUserCanUploadData()).String()
        err = writer.WriteStringValue("isUserCanUploadData", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsUserRolesSupport() != nil {
        cast := (*m.GetIsUserRolesSupport()).String()
        err = writer.WriteStringValue("isUserRolesSupport", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIsValidCertificateName() != nil {
        cast := (*m.GetIsValidCertificateName()).String()
        err = writer.WriteStringValue("isValidCertificateName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("latestBreachDateTime", m.GetLatestBreachDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("logonUrls", m.GetLogonUrls())
        if err != nil {
            return err
        }
    }
    if m.GetPciDssVersion() != nil {
        cast := (*m.GetPciDssVersion()).String()
        err = writer.WriteStringValue("pciDssVersion", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendor", m.GetVendorEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCsaStarLevel sets the csaStarLevel property value. The csaStarLevel property
func (m *DiscoveredCloudAppInfo) SetCsaStarLevel(value *AppInfoCsaStarLevel)() {
    err := m.GetBackingStore().Set("csaStarLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetDataAtRestEncryptionMethod sets the dataAtRestEncryptionMethod property value. The dataAtRestEncryptionMethod property
func (m *DiscoveredCloudAppInfo) SetDataAtRestEncryptionMethod(value *AppInfoDataAtRestEncryptionMethod)() {
    err := m.GetBackingStore().Set("dataAtRestEncryptionMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetDataCenter sets the dataCenter property value. Indicates the countries or regions in which your data center resides.
func (m *DiscoveredCloudAppInfo) SetDataCenter(value *string)() {
    err := m.GetBackingStore().Set("dataCenter", value)
    if err != nil {
        panic(err)
    }
}
// SetDataRetentionPolicy sets the dataRetentionPolicy property value. The dataRetentionPolicy property
func (m *DiscoveredCloudAppInfo) SetDataRetentionPolicy(value *AppInfoDataRetentionPolicy)() {
    err := m.GetBackingStore().Set("dataRetentionPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetDataTypes sets the dataTypes property value. The dataTypes property
func (m *DiscoveredCloudAppInfo) SetDataTypes(value *AppInfoUploadedDataTypes)() {
    err := m.GetBackingStore().Set("dataTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainRegistrationDateTime sets the domainRegistrationDateTime property value. Indicates the date when the app domain was registered.
func (m *DiscoveredCloudAppInfo) SetDomainRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("domainRegistrationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionProtocol sets the encryptionProtocol property value. The encryptionProtocol property
func (m *DiscoveredCloudAppInfo) SetEncryptionProtocol(value *AppInfoEncryptionProtocol)() {
    err := m.GetBackingStore().Set("encryptionProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetFedRampLevel sets the fedRampLevel property value. The fedRampLevel property
func (m *DiscoveredCloudAppInfo) SetFedRampLevel(value *AppInfoFedRampLevel)() {
    err := m.GetBackingStore().Set("fedRampLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetFounded sets the founded property value. Indicates the year that the specific app vendor was established.
func (m *DiscoveredCloudAppInfo) SetFounded(value *int32)() {
    err := m.GetBackingStore().Set("founded", value)
    if err != nil {
        panic(err)
    }
}
// SetGdprReadinessStatement sets the gdprReadinessStatement property value. Indicates the GDPR readiness of the app in relation to policies app provides to safeguard personal user data.
func (m *DiscoveredCloudAppInfo) SetGdprReadinessStatement(value *string)() {
    err := m.GetBackingStore().Set("gdprReadinessStatement", value)
    if err != nil {
        panic(err)
    }
}
// SetHeadquarters sets the headquarters property value. Indicates the location of the headquarters of the app.
func (m *DiscoveredCloudAppInfo) SetHeadquarters(value *string)() {
    err := m.GetBackingStore().Set("headquarters", value)
    if err != nil {
        panic(err)
    }
}
// SetHolding sets the holding property value. The holding property
func (m *DiscoveredCloudAppInfo) SetHolding(value *AppInfoHolding)() {
    err := m.GetBackingStore().Set("holding", value)
    if err != nil {
        panic(err)
    }
}
// SetHostingCompany sets the hostingCompany property value. Indicates the company name that provides hosting services for the app.
func (m *DiscoveredCloudAppInfo) SetHostingCompany(value *string)() {
    err := m.GetBackingStore().Set("hostingCompany", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAdminAuditTrail sets the isAdminAuditTrail property value. The isAdminAuditTrail property
func (m *DiscoveredCloudAppInfo) SetIsAdminAuditTrail(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isAdminAuditTrail", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCobitCompliant sets the isCobitCompliant property value. The isCobitCompliant property
func (m *DiscoveredCloudAppInfo) SetIsCobitCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isCobitCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCoppaCompliant sets the isCoppaCompliant property value. The isCoppaCompliant property
func (m *DiscoveredCloudAppInfo) SetIsCoppaCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isCoppaCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDataAuditTrail sets the isDataAuditTrail property value. The isDataAuditTrail property
func (m *DiscoveredCloudAppInfo) SetIsDataAuditTrail(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isDataAuditTrail", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDataClassification sets the isDataClassification property value. The isDataClassification property
func (m *DiscoveredCloudAppInfo) SetIsDataClassification(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isDataClassification", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDataOwnership sets the isDataOwnership property value. The isDataOwnership property
func (m *DiscoveredCloudAppInfo) SetIsDataOwnership(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isDataOwnership", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDisasterRecoveryPlan sets the isDisasterRecoveryPlan property value. The isDisasterRecoveryPlan property
func (m *DiscoveredCloudAppInfo) SetIsDisasterRecoveryPlan(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isDisasterRecoveryPlan", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDmca sets the isDmca property value. The isDmca property
func (m *DiscoveredCloudAppInfo) SetIsDmca(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isDmca", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFerpaCompliant sets the isFerpaCompliant property value. The isFerpaCompliant property
func (m *DiscoveredCloudAppInfo) SetIsFerpaCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isFerpaCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFfiecCompliant sets the isFfiecCompliant property value. The isFfiecCompliant property
func (m *DiscoveredCloudAppInfo) SetIsFfiecCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isFfiecCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFileSharing sets the isFileSharing property value. The isFileSharing property
func (m *DiscoveredCloudAppInfo) SetIsFileSharing(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isFileSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFinraCompliant sets the isFinraCompliant property value. The isFinraCompliant property
func (m *DiscoveredCloudAppInfo) SetIsFinraCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isFinraCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFismaCompliant sets the isFismaCompliant property value. The isFismaCompliant property
func (m *DiscoveredCloudAppInfo) SetIsFismaCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isFismaCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGaapCompliant sets the isGaapCompliant property value. The isGaapCompliant property
func (m *DiscoveredCloudAppInfo) SetIsGaapCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGaapCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprDataProtectionImpactAssessment sets the isGdprDataProtectionImpactAssessment property value. The isGdprDataProtectionImpactAssessment property
func (m *DiscoveredCloudAppInfo) SetIsGdprDataProtectionImpactAssessment(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprDataProtectionImpactAssessment", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprDataProtectionOfficer sets the isGdprDataProtectionOfficer property value. The isGdprDataProtectionOfficer property
func (m *DiscoveredCloudAppInfo) SetIsGdprDataProtectionOfficer(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprDataProtectionOfficer", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprDataProtectionSecureCrossBorderDataTransfer sets the isGdprDataProtectionSecureCrossBorderDataTransfer property value. The isGdprDataProtectionSecureCrossBorderDataTransfer property
func (m *DiscoveredCloudAppInfo) SetIsGdprDataProtectionSecureCrossBorderDataTransfer(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprDataProtectionSecureCrossBorderDataTransfer", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprImpactAssessment sets the isGdprImpactAssessment property value. The isGdprImpactAssessment property
func (m *DiscoveredCloudAppInfo) SetIsGdprImpactAssessment(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprImpactAssessment", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprLawfulBasisForProcessing sets the isGdprLawfulBasisForProcessing property value. The isGdprLawfulBasisForProcessing property
func (m *DiscoveredCloudAppInfo) SetIsGdprLawfulBasisForProcessing(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprLawfulBasisForProcessing", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprReportDataBreaches sets the isGdprReportDataBreaches property value. The isGdprReportDataBreaches property
func (m *DiscoveredCloudAppInfo) SetIsGdprReportDataBreaches(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprReportDataBreaches", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightsRelatedToAutomatedDecisionMaking sets the isGdprRightsRelatedToAutomatedDecisionMaking property value. The isGdprRightsRelatedToAutomatedDecisionMaking property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightsRelatedToAutomatedDecisionMaking(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightsRelatedToAutomatedDecisionMaking", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToAccess sets the isGdprRightToAccess property value. The isGdprRightToAccess property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToAccess(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToBeInformed sets the isGdprRightToBeInformed property value. The isGdprRightToBeInformed property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToBeInformed(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToBeInformed", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToDataPortablility sets the isGdprRightToDataPortablility property value. The isGdprRightToDataPortablility property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToDataPortablility(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToDataPortablility", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToErasure sets the isGdprRightToErasure property value. The isGdprRightToErasure property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToErasure(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToErasure", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToObject sets the isGdprRightToObject property value. The isGdprRightToObject property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToObject(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToObject", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToRectification sets the isGdprRightToRectification property value. The isGdprRightToRectification property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToRectification(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToRectification", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprRightToRestrictionOfProcessing sets the isGdprRightToRestrictionOfProcessing property value. The isGdprRightToRestrictionOfProcessing property
func (m *DiscoveredCloudAppInfo) SetIsGdprRightToRestrictionOfProcessing(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprRightToRestrictionOfProcessing", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGdprSecureCrossBorderDataControl sets the isGdprSecureCrossBorderDataControl property value. The isGdprSecureCrossBorderDataControl property
func (m *DiscoveredCloudAppInfo) SetIsGdprSecureCrossBorderDataControl(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGdprSecureCrossBorderDataControl", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGlbaCompliant sets the isGlbaCompliant property value. The isGlbaCompliant property
func (m *DiscoveredCloudAppInfo) SetIsGlbaCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isGlbaCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHipaaCompliant sets the isHipaaCompliant property value. The isHipaaCompliant property
func (m *DiscoveredCloudAppInfo) SetIsHipaaCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHipaaCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHitrustCsfCompliant sets the isHitrustCsfCompliant property value. The isHitrustCsfCompliant property
func (m *DiscoveredCloudAppInfo) SetIsHitrustCsfCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHitrustCsfCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHttpSecurityHeadersContentSecurityPolicy sets the isHttpSecurityHeadersContentSecurityPolicy property value. The isHttpSecurityHeadersContentSecurityPolicy property
func (m *DiscoveredCloudAppInfo) SetIsHttpSecurityHeadersContentSecurityPolicy(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHttpSecurityHeadersContentSecurityPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHttpSecurityHeadersStrictTransportSecurity sets the isHttpSecurityHeadersStrictTransportSecurity property value. The isHttpSecurityHeadersStrictTransportSecurity property
func (m *DiscoveredCloudAppInfo) SetIsHttpSecurityHeadersStrictTransportSecurity(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHttpSecurityHeadersStrictTransportSecurity", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHttpSecurityHeadersXContentTypeOptions sets the isHttpSecurityHeadersXContentTypeOptions property value. The isHttpSecurityHeadersXContentTypeOptions property
func (m *DiscoveredCloudAppInfo) SetIsHttpSecurityHeadersXContentTypeOptions(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHttpSecurityHeadersXContentTypeOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHttpSecurityHeadersXFrameOptions sets the isHttpSecurityHeadersXFrameOptions property value. The isHttpSecurityHeadersXFrameOptions property
func (m *DiscoveredCloudAppInfo) SetIsHttpSecurityHeadersXFrameOptions(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHttpSecurityHeadersXFrameOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHttpSecurityHeadersXXssProtection sets the isHttpSecurityHeadersXXssProtection property value. The isHttpSecurityHeadersXXssProtection property
func (m *DiscoveredCloudAppInfo) SetIsHttpSecurityHeadersXXssProtection(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isHttpSecurityHeadersXXssProtection", value)
    if err != nil {
        panic(err)
    }
}
// SetIsIpAddressRestriction sets the isIpAddressRestriction property value. The isIpAddressRestriction property
func (m *DiscoveredCloudAppInfo) SetIsIpAddressRestriction(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isIpAddressRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetIsIsae3402Compliant sets the isIsae3402Compliant property value. The isIsae3402Compliant property
func (m *DiscoveredCloudAppInfo) SetIsIsae3402Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isIsae3402Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsIso27001Compliant sets the isIso27001Compliant property value. The isIso27001Compliant property
func (m *DiscoveredCloudAppInfo) SetIsIso27001Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isIso27001Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsIso27017Compliant sets the isIso27017Compliant property value. The isIso27017Compliant property
func (m *DiscoveredCloudAppInfo) SetIsIso27017Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isIso27017Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsIso27018Compliant sets the isIso27018Compliant property value. The isIso27018Compliant property
func (m *DiscoveredCloudAppInfo) SetIsIso27018Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isIso27018Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsItarCompliant sets the isItarCompliant property value. The isItarCompliant property
func (m *DiscoveredCloudAppInfo) SetIsItarCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isItarCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMultiFactorAuthentication sets the isMultiFactorAuthentication property value. The isMultiFactorAuthentication property
func (m *DiscoveredCloudAppInfo) SetIsMultiFactorAuthentication(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isMultiFactorAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordPolicy sets the isPasswordPolicy property value. The isPasswordPolicy property
func (m *DiscoveredCloudAppInfo) SetIsPasswordPolicy(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPasswordPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordPolicyChangePasswordPeriod sets the isPasswordPolicyChangePasswordPeriod property value. The isPasswordPolicyChangePasswordPeriod property
func (m *DiscoveredCloudAppInfo) SetIsPasswordPolicyChangePasswordPeriod(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPasswordPolicyChangePasswordPeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordPolicyCharacterCombination sets the isPasswordPolicyCharacterCombination property value. The isPasswordPolicyCharacterCombination property
func (m *DiscoveredCloudAppInfo) SetIsPasswordPolicyCharacterCombination(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPasswordPolicyCharacterCombination", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordPolicyPasswordHistoryAndReuse sets the isPasswordPolicyPasswordHistoryAndReuse property value. The isPasswordPolicyPasswordHistoryAndReuse property
func (m *DiscoveredCloudAppInfo) SetIsPasswordPolicyPasswordHistoryAndReuse(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPasswordPolicyPasswordHistoryAndReuse", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordPolicyPasswordLengthLimit sets the isPasswordPolicyPasswordLengthLimit property value. The isPasswordPolicyPasswordLengthLimit property
func (m *DiscoveredCloudAppInfo) SetIsPasswordPolicyPasswordLengthLimit(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPasswordPolicyPasswordLengthLimit", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPasswordPolicyPersonalInformationUse sets the isPasswordPolicyPersonalInformationUse property value. The isPasswordPolicyPersonalInformationUse property
func (m *DiscoveredCloudAppInfo) SetIsPasswordPolicyPersonalInformationUse(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPasswordPolicyPersonalInformationUse", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPenetrationTesting sets the isPenetrationTesting property value. The isPenetrationTesting property
func (m *DiscoveredCloudAppInfo) SetIsPenetrationTesting(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPenetrationTesting", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPrivacyShieldCompliant sets the isPrivacyShieldCompliant property value. The isPrivacyShieldCompliant property
func (m *DiscoveredCloudAppInfo) SetIsPrivacyShieldCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isPrivacyShieldCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRememberPassword sets the isRememberPassword property value. The isRememberPassword property
func (m *DiscoveredCloudAppInfo) SetIsRememberPassword(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isRememberPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRequiresUserAuthentication sets the isRequiresUserAuthentication property value. The isRequiresUserAuthentication property
func (m *DiscoveredCloudAppInfo) SetIsRequiresUserAuthentication(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isRequiresUserAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSoc1Compliant sets the isSoc1Compliant property value. The isSoc1Compliant property
func (m *DiscoveredCloudAppInfo) SetIsSoc1Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSoc1Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSoc2Compliant sets the isSoc2Compliant property value. The isSoc2Compliant property
func (m *DiscoveredCloudAppInfo) SetIsSoc2Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSoc2Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSoc3Compliant sets the isSoc3Compliant property value. The isSoc3Compliant property
func (m *DiscoveredCloudAppInfo) SetIsSoc3Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSoc3Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSoxCompliant sets the isSoxCompliant property value. The isSoxCompliant property
func (m *DiscoveredCloudAppInfo) SetIsSoxCompliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSoxCompliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSp80053Compliant sets the isSp80053Compliant property value. The isSp80053Compliant property
func (m *DiscoveredCloudAppInfo) SetIsSp80053Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSp80053Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSsae16Compliant sets the isSsae16Compliant property value. The isSsae16Compliant property
func (m *DiscoveredCloudAppInfo) SetIsSsae16Compliant(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSsae16Compliant", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSupportsSaml sets the isSupportsSaml property value. The isSupportsSaml property
func (m *DiscoveredCloudAppInfo) SetIsSupportsSaml(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isSupportsSaml", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTrustedCertificate sets the isTrustedCertificate property value. The isTrustedCertificate property
func (m *DiscoveredCloudAppInfo) SetIsTrustedCertificate(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isTrustedCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetIsUserAuditTrail sets the isUserAuditTrail property value. The isUserAuditTrail property
func (m *DiscoveredCloudAppInfo) SetIsUserAuditTrail(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isUserAuditTrail", value)
    if err != nil {
        panic(err)
    }
}
// SetIsUserCanUploadData sets the isUserCanUploadData property value. The isUserCanUploadData property
func (m *DiscoveredCloudAppInfo) SetIsUserCanUploadData(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isUserCanUploadData", value)
    if err != nil {
        panic(err)
    }
}
// SetIsUserRolesSupport sets the isUserRolesSupport property value. The isUserRolesSupport property
func (m *DiscoveredCloudAppInfo) SetIsUserRolesSupport(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isUserRolesSupport", value)
    if err != nil {
        panic(err)
    }
}
// SetIsValidCertificateName sets the isValidCertificateName property value. The isValidCertificateName property
func (m *DiscoveredCloudAppInfo) SetIsValidCertificateName(value *CloudAppInfoState)() {
    err := m.GetBackingStore().Set("isValidCertificateName", value)
    if err != nil {
        panic(err)
    }
}
// SetLatestBreachDateTime sets the latestBreachDateTime property value. Indicates the last date of the data breach for the company.
func (m *DiscoveredCloudAppInfo) SetLatestBreachDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("latestBreachDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLogonUrls sets the logonUrls property value. Indicates the URL that users can use to sign into the app.
func (m *DiscoveredCloudAppInfo) SetLogonUrls(value *string)() {
    err := m.GetBackingStore().Set("logonUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetPciDssVersion sets the pciDssVersion property value. The pciDssVersion property
func (m *DiscoveredCloudAppInfo) SetPciDssVersion(value *AppInfoPciDssVersion)() {
    err := m.GetBackingStore().Set("pciDssVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorEscaped sets the vendor property value. Indicates the app vendor.
func (m *DiscoveredCloudAppInfo) SetVendorEscaped(value *string)() {
    err := m.GetBackingStore().Set("vendorEscaped", value)
    if err != nil {
        panic(err)
    }
}
type DiscoveredCloudAppInfoable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCsaStarLevel()(*AppInfoCsaStarLevel)
    GetDataAtRestEncryptionMethod()(*AppInfoDataAtRestEncryptionMethod)
    GetDataCenter()(*string)
    GetDataRetentionPolicy()(*AppInfoDataRetentionPolicy)
    GetDataTypes()(*AppInfoUploadedDataTypes)
    GetDomainRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEncryptionProtocol()(*AppInfoEncryptionProtocol)
    GetFedRampLevel()(*AppInfoFedRampLevel)
    GetFounded()(*int32)
    GetGdprReadinessStatement()(*string)
    GetHeadquarters()(*string)
    GetHolding()(*AppInfoHolding)
    GetHostingCompany()(*string)
    GetIsAdminAuditTrail()(*CloudAppInfoState)
    GetIsCobitCompliant()(*CloudAppInfoState)
    GetIsCoppaCompliant()(*CloudAppInfoState)
    GetIsDataAuditTrail()(*CloudAppInfoState)
    GetIsDataClassification()(*CloudAppInfoState)
    GetIsDataOwnership()(*CloudAppInfoState)
    GetIsDisasterRecoveryPlan()(*CloudAppInfoState)
    GetIsDmca()(*CloudAppInfoState)
    GetIsFerpaCompliant()(*CloudAppInfoState)
    GetIsFfiecCompliant()(*CloudAppInfoState)
    GetIsFileSharing()(*CloudAppInfoState)
    GetIsFinraCompliant()(*CloudAppInfoState)
    GetIsFismaCompliant()(*CloudAppInfoState)
    GetIsGaapCompliant()(*CloudAppInfoState)
    GetIsGdprDataProtectionImpactAssessment()(*CloudAppInfoState)
    GetIsGdprDataProtectionOfficer()(*CloudAppInfoState)
    GetIsGdprDataProtectionSecureCrossBorderDataTransfer()(*CloudAppInfoState)
    GetIsGdprImpactAssessment()(*CloudAppInfoState)
    GetIsGdprLawfulBasisForProcessing()(*CloudAppInfoState)
    GetIsGdprReportDataBreaches()(*CloudAppInfoState)
    GetIsGdprRightsRelatedToAutomatedDecisionMaking()(*CloudAppInfoState)
    GetIsGdprRightToAccess()(*CloudAppInfoState)
    GetIsGdprRightToBeInformed()(*CloudAppInfoState)
    GetIsGdprRightToDataPortablility()(*CloudAppInfoState)
    GetIsGdprRightToErasure()(*CloudAppInfoState)
    GetIsGdprRightToObject()(*CloudAppInfoState)
    GetIsGdprRightToRectification()(*CloudAppInfoState)
    GetIsGdprRightToRestrictionOfProcessing()(*CloudAppInfoState)
    GetIsGdprSecureCrossBorderDataControl()(*CloudAppInfoState)
    GetIsGlbaCompliant()(*CloudAppInfoState)
    GetIsHipaaCompliant()(*CloudAppInfoState)
    GetIsHitrustCsfCompliant()(*CloudAppInfoState)
    GetIsHttpSecurityHeadersContentSecurityPolicy()(*CloudAppInfoState)
    GetIsHttpSecurityHeadersStrictTransportSecurity()(*CloudAppInfoState)
    GetIsHttpSecurityHeadersXContentTypeOptions()(*CloudAppInfoState)
    GetIsHttpSecurityHeadersXFrameOptions()(*CloudAppInfoState)
    GetIsHttpSecurityHeadersXXssProtection()(*CloudAppInfoState)
    GetIsIpAddressRestriction()(*CloudAppInfoState)
    GetIsIsae3402Compliant()(*CloudAppInfoState)
    GetIsIso27001Compliant()(*CloudAppInfoState)
    GetIsIso27017Compliant()(*CloudAppInfoState)
    GetIsIso27018Compliant()(*CloudAppInfoState)
    GetIsItarCompliant()(*CloudAppInfoState)
    GetIsMultiFactorAuthentication()(*CloudAppInfoState)
    GetIsPasswordPolicy()(*CloudAppInfoState)
    GetIsPasswordPolicyChangePasswordPeriod()(*CloudAppInfoState)
    GetIsPasswordPolicyCharacterCombination()(*CloudAppInfoState)
    GetIsPasswordPolicyPasswordHistoryAndReuse()(*CloudAppInfoState)
    GetIsPasswordPolicyPasswordLengthLimit()(*CloudAppInfoState)
    GetIsPasswordPolicyPersonalInformationUse()(*CloudAppInfoState)
    GetIsPenetrationTesting()(*CloudAppInfoState)
    GetIsPrivacyShieldCompliant()(*CloudAppInfoState)
    GetIsRememberPassword()(*CloudAppInfoState)
    GetIsRequiresUserAuthentication()(*CloudAppInfoState)
    GetIsSoc1Compliant()(*CloudAppInfoState)
    GetIsSoc2Compliant()(*CloudAppInfoState)
    GetIsSoc3Compliant()(*CloudAppInfoState)
    GetIsSoxCompliant()(*CloudAppInfoState)
    GetIsSp80053Compliant()(*CloudAppInfoState)
    GetIsSsae16Compliant()(*CloudAppInfoState)
    GetIsSupportsSaml()(*CloudAppInfoState)
    GetIsTrustedCertificate()(*CloudAppInfoState)
    GetIsUserAuditTrail()(*CloudAppInfoState)
    GetIsUserCanUploadData()(*CloudAppInfoState)
    GetIsUserRolesSupport()(*CloudAppInfoState)
    GetIsValidCertificateName()(*CloudAppInfoState)
    GetLatestBreachDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogonUrls()(*string)
    GetPciDssVersion()(*AppInfoPciDssVersion)
    GetVendorEscaped()(*string)
    SetCsaStarLevel(value *AppInfoCsaStarLevel)()
    SetDataAtRestEncryptionMethod(value *AppInfoDataAtRestEncryptionMethod)()
    SetDataCenter(value *string)()
    SetDataRetentionPolicy(value *AppInfoDataRetentionPolicy)()
    SetDataTypes(value *AppInfoUploadedDataTypes)()
    SetDomainRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEncryptionProtocol(value *AppInfoEncryptionProtocol)()
    SetFedRampLevel(value *AppInfoFedRampLevel)()
    SetFounded(value *int32)()
    SetGdprReadinessStatement(value *string)()
    SetHeadquarters(value *string)()
    SetHolding(value *AppInfoHolding)()
    SetHostingCompany(value *string)()
    SetIsAdminAuditTrail(value *CloudAppInfoState)()
    SetIsCobitCompliant(value *CloudAppInfoState)()
    SetIsCoppaCompliant(value *CloudAppInfoState)()
    SetIsDataAuditTrail(value *CloudAppInfoState)()
    SetIsDataClassification(value *CloudAppInfoState)()
    SetIsDataOwnership(value *CloudAppInfoState)()
    SetIsDisasterRecoveryPlan(value *CloudAppInfoState)()
    SetIsDmca(value *CloudAppInfoState)()
    SetIsFerpaCompliant(value *CloudAppInfoState)()
    SetIsFfiecCompliant(value *CloudAppInfoState)()
    SetIsFileSharing(value *CloudAppInfoState)()
    SetIsFinraCompliant(value *CloudAppInfoState)()
    SetIsFismaCompliant(value *CloudAppInfoState)()
    SetIsGaapCompliant(value *CloudAppInfoState)()
    SetIsGdprDataProtectionImpactAssessment(value *CloudAppInfoState)()
    SetIsGdprDataProtectionOfficer(value *CloudAppInfoState)()
    SetIsGdprDataProtectionSecureCrossBorderDataTransfer(value *CloudAppInfoState)()
    SetIsGdprImpactAssessment(value *CloudAppInfoState)()
    SetIsGdprLawfulBasisForProcessing(value *CloudAppInfoState)()
    SetIsGdprReportDataBreaches(value *CloudAppInfoState)()
    SetIsGdprRightsRelatedToAutomatedDecisionMaking(value *CloudAppInfoState)()
    SetIsGdprRightToAccess(value *CloudAppInfoState)()
    SetIsGdprRightToBeInformed(value *CloudAppInfoState)()
    SetIsGdprRightToDataPortablility(value *CloudAppInfoState)()
    SetIsGdprRightToErasure(value *CloudAppInfoState)()
    SetIsGdprRightToObject(value *CloudAppInfoState)()
    SetIsGdprRightToRectification(value *CloudAppInfoState)()
    SetIsGdprRightToRestrictionOfProcessing(value *CloudAppInfoState)()
    SetIsGdprSecureCrossBorderDataControl(value *CloudAppInfoState)()
    SetIsGlbaCompliant(value *CloudAppInfoState)()
    SetIsHipaaCompliant(value *CloudAppInfoState)()
    SetIsHitrustCsfCompliant(value *CloudAppInfoState)()
    SetIsHttpSecurityHeadersContentSecurityPolicy(value *CloudAppInfoState)()
    SetIsHttpSecurityHeadersStrictTransportSecurity(value *CloudAppInfoState)()
    SetIsHttpSecurityHeadersXContentTypeOptions(value *CloudAppInfoState)()
    SetIsHttpSecurityHeadersXFrameOptions(value *CloudAppInfoState)()
    SetIsHttpSecurityHeadersXXssProtection(value *CloudAppInfoState)()
    SetIsIpAddressRestriction(value *CloudAppInfoState)()
    SetIsIsae3402Compliant(value *CloudAppInfoState)()
    SetIsIso27001Compliant(value *CloudAppInfoState)()
    SetIsIso27017Compliant(value *CloudAppInfoState)()
    SetIsIso27018Compliant(value *CloudAppInfoState)()
    SetIsItarCompliant(value *CloudAppInfoState)()
    SetIsMultiFactorAuthentication(value *CloudAppInfoState)()
    SetIsPasswordPolicy(value *CloudAppInfoState)()
    SetIsPasswordPolicyChangePasswordPeriod(value *CloudAppInfoState)()
    SetIsPasswordPolicyCharacterCombination(value *CloudAppInfoState)()
    SetIsPasswordPolicyPasswordHistoryAndReuse(value *CloudAppInfoState)()
    SetIsPasswordPolicyPasswordLengthLimit(value *CloudAppInfoState)()
    SetIsPasswordPolicyPersonalInformationUse(value *CloudAppInfoState)()
    SetIsPenetrationTesting(value *CloudAppInfoState)()
    SetIsPrivacyShieldCompliant(value *CloudAppInfoState)()
    SetIsRememberPassword(value *CloudAppInfoState)()
    SetIsRequiresUserAuthentication(value *CloudAppInfoState)()
    SetIsSoc1Compliant(value *CloudAppInfoState)()
    SetIsSoc2Compliant(value *CloudAppInfoState)()
    SetIsSoc3Compliant(value *CloudAppInfoState)()
    SetIsSoxCompliant(value *CloudAppInfoState)()
    SetIsSp80053Compliant(value *CloudAppInfoState)()
    SetIsSsae16Compliant(value *CloudAppInfoState)()
    SetIsSupportsSaml(value *CloudAppInfoState)()
    SetIsTrustedCertificate(value *CloudAppInfoState)()
    SetIsUserAuditTrail(value *CloudAppInfoState)()
    SetIsUserCanUploadData(value *CloudAppInfoState)()
    SetIsUserRolesSupport(value *CloudAppInfoState)()
    SetIsValidCertificateName(value *CloudAppInfoState)()
    SetLatestBreachDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogonUrls(value *string)()
    SetPciDssVersion(value *AppInfoPciDssVersion)()
    SetVendorEscaped(value *string)()
}
