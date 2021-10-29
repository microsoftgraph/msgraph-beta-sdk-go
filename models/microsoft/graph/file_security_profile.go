package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new fileSecurityProfile and sets the default values.
func NewFileSecurityProfile()(*FileSecurityProfile) {
    m := &FileSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activityGroupNames property value. 
func (m *FileSecurityProfile) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
// Gets the azureSubscriptionId property value. 
func (m *FileSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// Gets the azureTenantId property value. 
func (m *FileSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the certificateThumbprint property value. 
func (m *FileSecurityProfile) GetCertificateThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateThumbprint
    }
}
// Gets the extensions property value. 
func (m *FileSecurityProfile) GetExtensions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the fileType property value. 
func (m *FileSecurityProfile) GetFileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileType
    }
}
// Gets the firstSeenDateTime property value. 
func (m *FileSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// Gets the hashes property value. 
func (m *FileSecurityProfile) GetHashes()([]FileHash) {
    if m == nil {
        return nil
    } else {
        return m.hashes
    }
}
// Gets the lastSeenDateTime property value. 
func (m *FileSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// Gets the malwareStates property value. 
func (m *FileSecurityProfile) GetMalwareStates()([]MalwareState) {
    if m == nil {
        return nil
    } else {
        return m.malwareStates
    }
}
// Gets the names property value. 
func (m *FileSecurityProfile) GetNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// Gets the riskScore property value. 
func (m *FileSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the size property value. 
func (m *FileSecurityProfile) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Gets the tags property value. 
func (m *FileSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the vendorInformation property value. 
func (m *FileSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// Gets the vulnerabilityStates property value. 
func (m *FileSecurityProfile) GetVulnerabilityStates()([]VulnerabilityState) {
    if m == nil {
        return nil
    } else {
        return m.vulnerabilityStates
    }
}
// The deserialization information for the current model
func (m *FileSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureSubscriptionId(val)
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
    res["certificateThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateThumbprint(val)
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExtensions(res)
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
    res["firstSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetFirstSeenDateTime(val)
        return nil
    }
    res["hashes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileHash() })
        if err != nil {
            return err
        }
        res := make([]FileHash, len(val))
        for i, v := range val {
            res[i] = *(v.(*FileHash))
        }
        m.SetHashes(res)
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSeenDateTime(val)
        return nil
    }
    res["malwareStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMalwareState() })
        if err != nil {
            return err
        }
        res := make([]MalwareState, len(val))
        for i, v := range val {
            res[i] = *(v.(*MalwareState))
        }
        m.SetMalwareStates(res)
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetNames(res)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
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
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    res["vulnerabilityStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVulnerabilityState() })
        if err != nil {
            return err
        }
        res := make([]VulnerabilityState, len(val))
        for i, v := range val {
            res[i] = *(v.(*VulnerabilityState))
        }
        m.SetVulnerabilityStates(res)
        return nil
    }
    return res
}
func (m *FileSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *FileSecurityProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
// Sets the activityGroupNames property value. 
// Parameters:
//  - value : Value to set for the activityGroupNames property.
func (m *FileSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// Sets the azureSubscriptionId property value. 
// Parameters:
//  - value : Value to set for the azureSubscriptionId property.
func (m *FileSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// Sets the azureTenantId property value. 
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *FileSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the certificateThumbprint property value. 
// Parameters:
//  - value : Value to set for the certificateThumbprint property.
func (m *FileSecurityProfile) SetCertificateThumbprint(value *string)() {
    m.certificateThumbprint = value
}
// Sets the extensions property value. 
// Parameters:
//  - value : Value to set for the extensions property.
func (m *FileSecurityProfile) SetExtensions(value []string)() {
    m.extensions = value
}
// Sets the fileType property value. 
// Parameters:
//  - value : Value to set for the fileType property.
func (m *FileSecurityProfile) SetFileType(value *string)() {
    m.fileType = value
}
// Sets the firstSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the firstSeenDateTime property.
func (m *FileSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// Sets the hashes property value. 
// Parameters:
//  - value : Value to set for the hashes property.
func (m *FileSecurityProfile) SetHashes(value []FileHash)() {
    m.hashes = value
}
// Sets the lastSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the lastSeenDateTime property.
func (m *FileSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// Sets the malwareStates property value. 
// Parameters:
//  - value : Value to set for the malwareStates property.
func (m *FileSecurityProfile) SetMalwareStates(value []MalwareState)() {
    m.malwareStates = value
}
// Sets the names property value. 
// Parameters:
//  - value : Value to set for the names property.
func (m *FileSecurityProfile) SetNames(value []string)() {
    m.names = value
}
// Sets the riskScore property value. 
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *FileSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the size property value. 
// Parameters:
//  - value : Value to set for the size property.
func (m *FileSecurityProfile) SetSize(value *int64)() {
    m.size = value
}
// Sets the tags property value. 
// Parameters:
//  - value : Value to set for the tags property.
func (m *FileSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// Sets the vendorInformation property value. 
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *FileSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
// Sets the vulnerabilityStates property value. 
// Parameters:
//  - value : Value to set for the vulnerabilityStates property.
func (m *FileSecurityProfile) SetVulnerabilityStates(value []VulnerabilityState)() {
    m.vulnerabilityStates = value
}
