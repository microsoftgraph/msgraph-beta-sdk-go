package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileSecurityProfile 
type FileSecurityProfile struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewFileSecurityProfile instantiates a new fileSecurityProfile and sets the default values.
func NewFileSecurityProfile()(*FileSecurityProfile) {
    m := &FileSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFileSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileSecurityProfile(), nil
}
// GetActivityGroupNames gets the activityGroupNames property value. The activityGroupNames property
func (m *FileSecurityProfile) GetActivityGroupNames()([]string) {
    val, err := m.GetBackingStore().Get("activityGroupNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *FileSecurityProfile) GetAzureSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("azureSubscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *FileSecurityProfile) GetAzureTenantId()(*string) {
    val, err := m.GetBackingStore().Get("azureTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateThumbprint gets the certificateThumbprint property value. The certificateThumbprint property
func (m *FileSecurityProfile) GetCertificateThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("certificateThumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExtensions gets the extensions property value. The extensions property
func (m *FileSecurityProfile) GetExtensions()([]string) {
    val, err := m.GetBackingStore().Get("extensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["azureSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
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
    res["certificateThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateThumbprint(val)
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExtensions(res)
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
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["hashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileHashable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileHashable)
                }
            }
            m.SetHashes(res)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["malwareStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMalwareStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MalwareStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MalwareStateable)
                }
            }
            m.SetMalwareStates(res)
        }
        return nil
    }
    res["names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetNames(res)
        }
        return nil
    }
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
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
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    res["vulnerabilityStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVulnerabilityStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VulnerabilityStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VulnerabilityStateable)
                }
            }
            m.SetVulnerabilityStates(res)
        }
        return nil
    }
    return res
}
// GetFileType gets the fileType property value. The fileType property
func (m *FileSecurityProfile) GetFileType()(*string) {
    val, err := m.GetBackingStore().Get("fileType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *FileSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetHashes gets the hashes property value. The hashes property
func (m *FileSecurityProfile) GetHashes()([]FileHashable) {
    val, err := m.GetBackingStore().Get("hashes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FileHashable)
    }
    return nil
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *FileSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMalwareStates gets the malwareStates property value. The malwareStates property
func (m *FileSecurityProfile) GetMalwareStates()([]MalwareStateable) {
    val, err := m.GetBackingStore().Get("malwareStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MalwareStateable)
    }
    return nil
}
// GetNames gets the names property value. The names property
func (m *FileSecurityProfile) GetNames()([]string) {
    val, err := m.GetBackingStore().Get("names")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *FileSecurityProfile) GetRiskScore()(*string) {
    val, err := m.GetBackingStore().Get("riskScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSize gets the size property value. The size property
func (m *FileSecurityProfile) GetSize()(*int64) {
    val, err := m.GetBackingStore().Get("size")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTags gets the tags property value. The tags property
func (m *FileSecurityProfile) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *FileSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    val, err := m.GetBackingStore().Get("vendorInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecurityVendorInformationable)
    }
    return nil
}
// GetVulnerabilityStates gets the vulnerabilityStates property value. The vulnerabilityStates property
func (m *FileSecurityProfile) GetVulnerabilityStates()([]VulnerabilityStateable) {
    val, err := m.GetBackingStore().Get("vulnerabilityStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VulnerabilityStateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FileSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityGroupNames() != nil {
        err = writer.WriteCollectionOfStringValues("activityGroupNames", m.GetActivityGroupNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
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
        err = writer.WriteStringValue("certificateThumbprint", m.GetCertificateThumbprint())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        err = writer.WriteCollectionOfStringValues("extensions", m.GetExtensions())
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
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetHashes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHashes()))
        for i, v := range m.GetHashes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hashes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMalwareStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMalwareStates()))
        for i, v := range m.GetMalwareStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("malwareStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        err = writer.WriteCollectionOfStringValues("names", m.GetNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
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
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilityStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVulnerabilityStates()))
        for i, v := range m.GetVulnerabilityStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("vulnerabilityStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityGroupNames sets the activityGroupNames property value. The activityGroupNames property
func (m *FileSecurityProfile) SetActivityGroupNames(value []string)() {
    err := m.GetBackingStore().Set("activityGroupNames", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *FileSecurityProfile) SetAzureSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("azureSubscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *FileSecurityProfile) SetAzureTenantId(value *string)() {
    err := m.GetBackingStore().Set("azureTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateThumbprint sets the certificateThumbprint property value. The certificateThumbprint property
func (m *FileSecurityProfile) SetCertificateThumbprint(value *string)() {
    err := m.GetBackingStore().Set("certificateThumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetExtensions sets the extensions property value. The extensions property
func (m *FileSecurityProfile) SetExtensions(value []string)() {
    err := m.GetBackingStore().Set("extensions", value)
    if err != nil {
        panic(err)
    }
}
// SetFileType sets the fileType property value. The fileType property
func (m *FileSecurityProfile) SetFileType(value *string)() {
    err := m.GetBackingStore().Set("fileType", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *FileSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetHashes sets the hashes property value. The hashes property
func (m *FileSecurityProfile) SetHashes(value []FileHashable)() {
    err := m.GetBackingStore().Set("hashes", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *FileSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMalwareStates sets the malwareStates property value. The malwareStates property
func (m *FileSecurityProfile) SetMalwareStates(value []MalwareStateable)() {
    err := m.GetBackingStore().Set("malwareStates", value)
    if err != nil {
        panic(err)
    }
}
// SetNames sets the names property value. The names property
func (m *FileSecurityProfile) SetNames(value []string)() {
    err := m.GetBackingStore().Set("names", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *FileSecurityProfile) SetRiskScore(value *string)() {
    err := m.GetBackingStore().Set("riskScore", value)
    if err != nil {
        panic(err)
    }
}
// SetSize sets the size property value. The size property
func (m *FileSecurityProfile) SetSize(value *int64)() {
    err := m.GetBackingStore().Set("size", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. The tags property
func (m *FileSecurityProfile) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *FileSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    err := m.GetBackingStore().Set("vendorInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetVulnerabilityStates sets the vulnerabilityStates property value. The vulnerabilityStates property
func (m *FileSecurityProfile) SetVulnerabilityStates(value []VulnerabilityStateable)() {
    err := m.GetBackingStore().Set("vulnerabilityStates", value)
    if err != nil {
        panic(err)
    }
}
// FileSecurityProfileable 
type FileSecurityProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityGroupNames()([]string)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCertificateThumbprint()(*string)
    GetExtensions()([]string)
    GetFileType()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHashes()([]FileHashable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareStates()([]MalwareStateable)
    GetNames()([]string)
    GetRiskScore()(*string)
    GetSize()(*int64)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    GetVulnerabilityStates()([]VulnerabilityStateable)
    SetActivityGroupNames(value []string)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCertificateThumbprint(value *string)()
    SetExtensions(value []string)()
    SetFileType(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHashes(value []FileHashable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareStates(value []MalwareStateable)()
    SetNames(value []string)()
    SetRiskScore(value *string)()
    SetSize(value *int64)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
    SetVulnerabilityStates(value []VulnerabilityStateable)()
}
