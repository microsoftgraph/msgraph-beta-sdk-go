package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FileSecurityProfile 
type FileSecurityProfile struct {
    Entity
    // 
    activityGroupNames []string;
    // 
    azureSubscriptionId *string;
    // 
    azureTenantId *string;
    // 
    certificateThumbprint *string;
    // 
    extensions []string;
    // 
    fileType *string;
    // 
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    hashes []FileHash;
    // 
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    malwareStates []MalwareState;
    // 
    names []string;
    // 
    riskScore *string;
    // 
    size *int64;
    // 
    tags []string;
    // 
    vendorInformation *SecurityVendorInformation;
    // 
    vulnerabilityStates []VulnerabilityState;
}
// NewFileSecurityProfile instantiates a new fileSecurityProfile and sets the default values.
func NewFileSecurityProfile()(*FileSecurityProfile) {
    m := &FileSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivityGroupNames gets the activityGroupNames property value. 
func (m *FileSecurityProfile) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. 
func (m *FileSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// GetAzureTenantId gets the azureTenantId property value. 
func (m *FileSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetCertificateThumbprint gets the certificateThumbprint property value. 
func (m *FileSecurityProfile) GetCertificateThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateThumbprint
    }
}
// GetExtensions gets the extensions property value. 
func (m *FileSecurityProfile) GetExtensions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFileType gets the fileType property value. 
func (m *FileSecurityProfile) GetFileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileType
    }
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. 
func (m *FileSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// GetHashes gets the hashes property value. 
func (m *FileSecurityProfile) GetHashes()([]FileHash) {
    if m == nil {
        return nil
    } else {
        return m.hashes
    }
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. 
func (m *FileSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// GetMalwareStates gets the malwareStates property value. 
func (m *FileSecurityProfile) GetMalwareStates()([]MalwareState) {
    if m == nil {
        return nil
    } else {
        return m.malwareStates
    }
}
// GetNames gets the names property value. 
func (m *FileSecurityProfile) GetNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// GetRiskScore gets the riskScore property value. 
func (m *FileSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// GetSize gets the size property value. 
func (m *FileSecurityProfile) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetTags gets the tags property value. 
func (m *FileSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetVendorInformation gets the vendorInformation property value. 
func (m *FileSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// GetVulnerabilityStates gets the vulnerabilityStates property value. 
func (m *FileSecurityProfile) GetVulnerabilityStates()([]VulnerabilityState) {
    if m == nil {
        return nil
    } else {
        return m.vulnerabilityStates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
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
    res["certificateThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateThumbprint(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExtensions(res)
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
    res["firstSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["hashes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileHash() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileHash, len(val))
            for i, v := range val {
                res[i] = *(v.(*FileHash))
            }
            m.SetHashes(res)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["malwareStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMalwareState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MalwareState, len(val))
            for i, v := range val {
                res[i] = *(v.(*MalwareState))
            }
            m.SetMalwareStates(res)
        }
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNames(res)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
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
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(*SecurityVendorInformation))
        }
        return nil
    }
    res["vulnerabilityStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVulnerabilityState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VulnerabilityState, len(val))
            for i, v := range val {
                res[i] = *(v.(*VulnerabilityState))
            }
            m.SetVulnerabilityStates(res)
        }
        return nil
    }
    return res
}
func (m *FileSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FileSecurityProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHashes()))
        for i, v := range m.GetHashes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMalwareStates()))
        for i, v := range m.GetMalwareStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVulnerabilityStates()))
        for i, v := range m.GetVulnerabilityStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("vulnerabilityStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityGroupNames sets the activityGroupNames property value. 
func (m *FileSecurityProfile) SetActivityGroupNames(value []string)() {
    if m != nil {
        m.activityGroupNames = value
    }
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. 
func (m *FileSecurityProfile) SetAzureSubscriptionId(value *string)() {
    if m != nil {
        m.azureSubscriptionId = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. 
func (m *FileSecurityProfile) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetCertificateThumbprint sets the certificateThumbprint property value. 
func (m *FileSecurityProfile) SetCertificateThumbprint(value *string)() {
    if m != nil {
        m.certificateThumbprint = value
    }
}
// SetExtensions sets the extensions property value. 
func (m *FileSecurityProfile) SetExtensions(value []string)() {
    if m != nil {
        m.extensions = value
    }
}
// SetFileType sets the fileType property value. 
func (m *FileSecurityProfile) SetFileType(value *string)() {
    if m != nil {
        m.fileType = value
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. 
func (m *FileSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.firstSeenDateTime = value
    }
}
// SetHashes sets the hashes property value. 
func (m *FileSecurityProfile) SetHashes(value []FileHash)() {
    if m != nil {
        m.hashes = value
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. 
func (m *FileSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSeenDateTime = value
    }
}
// SetMalwareStates sets the malwareStates property value. 
func (m *FileSecurityProfile) SetMalwareStates(value []MalwareState)() {
    if m != nil {
        m.malwareStates = value
    }
}
// SetNames sets the names property value. 
func (m *FileSecurityProfile) SetNames(value []string)() {
    if m != nil {
        m.names = value
    }
}
// SetRiskScore sets the riskScore property value. 
func (m *FileSecurityProfile) SetRiskScore(value *string)() {
    if m != nil {
        m.riskScore = value
    }
}
// SetSize sets the size property value. 
func (m *FileSecurityProfile) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
// SetTags sets the tags property value. 
func (m *FileSecurityProfile) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetVendorInformation sets the vendorInformation property value. 
func (m *FileSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    if m != nil {
        m.vendorInformation = value
    }
}
// SetVulnerabilityStates sets the vulnerabilityStates property value. 
func (m *FileSecurityProfile) SetVulnerabilityStates(value []VulnerabilityState)() {
    if m != nil {
        m.vulnerabilityStates = value
    }
}
