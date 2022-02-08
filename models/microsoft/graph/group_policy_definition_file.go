package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyDefinitionFile 
type GroupPolicyDefinitionFile struct {
    Entity
    // The group policy definitions associated with the file.
    definitions []GroupPolicyDefinition;
    // The localized description of the policy settings in the ADMX file. The default value is empty.
    description *string;
    // The localized friendly name of the ADMX file.
    displayName *string;
    // The file name of the ADMX file without the path. For example: edge.admx
    fileName *string;
    // The supported language codes for the ADMX file.
    languageCodes []string;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
    policyType *GroupPolicyType;
    // The revision version associated with the file.
    revision *string;
    // Specifies the URI used to identify the namespace within the ADMX file.
    targetNamespace *string;
    // Specifies the logical name that refers to the namespace within the ADMX file.
    targetPrefix *string;
}
// NewGroupPolicyDefinitionFile instantiates a new groupPolicyDefinitionFile and sets the default values.
func NewGroupPolicyDefinitionFile()(*GroupPolicyDefinitionFile) {
    m := &GroupPolicyDefinitionFile{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefinitions gets the definitions property value. The group policy definitions associated with the file.
func (m *GroupPolicyDefinitionFile) GetDefinitions()([]GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
// GetDescription gets the description property value. The localized description of the policy settings in the ADMX file. The default value is empty.
func (m *GroupPolicyDefinitionFile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The localized friendly name of the ADMX file.
func (m *GroupPolicyDefinitionFile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFileName gets the fileName property value. The file name of the ADMX file without the path. For example: edge.admx
func (m *GroupPolicyDefinitionFile) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetLanguageCodes gets the languageCodes property value. The supported language codes for the ADMX file.
func (m *GroupPolicyDefinitionFile) GetLanguageCodes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.languageCodes
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPolicyType gets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
func (m *GroupPolicyDefinitionFile) GetPolicyType()(*GroupPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
// GetRevision gets the revision property value. The revision version associated with the file.
func (m *GroupPolicyDefinitionFile) GetRevision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.revision
    }
}
// GetTargetNamespace gets the targetNamespace property value. Specifies the URI used to identify the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) GetTargetNamespace()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetNamespace
    }
}
// GetTargetPrefix gets the targetPrefix property value. Specifies the logical name that refers to the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) GetTargetPrefix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetPrefix
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyDefinitionFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyDefinition))
            }
            m.SetDefinitions(res)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["languageCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLanguageCodes(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["policyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyType(val.(*GroupPolicyType))
        }
        return nil
    }
    res["revision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevision(val)
        }
        return nil
    }
    res["targetNamespace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetNamespace(val)
        }
        return nil
    }
    res["targetPrefix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPrefix(val)
        }
        return nil
    }
    return res
}
func (m *GroupPolicyDefinitionFile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyDefinitionFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
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
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetLanguageCodes() != nil {
        err = writer.WriteCollectionOfStringValues("languageCodes", m.GetLanguageCodes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := (*m.GetPolicyType()).String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("revision", m.GetRevision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetNamespace", m.GetTargetNamespace())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetPrefix", m.GetTargetPrefix())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinitions sets the definitions property value. The group policy definitions associated with the file.
func (m *GroupPolicyDefinitionFile) SetDefinitions(value []GroupPolicyDefinition)() {
    if m != nil {
        m.definitions = value
    }
}
// SetDescription sets the description property value. The localized description of the policy settings in the ADMX file. The default value is empty.
func (m *GroupPolicyDefinitionFile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The localized friendly name of the ADMX file.
func (m *GroupPolicyDefinitionFile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFileName sets the fileName property value. The file name of the ADMX file without the path. For example: edge.admx
func (m *GroupPolicyDefinitionFile) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetLanguageCodes sets the languageCodes property value. The supported language codes for the ADMX file.
func (m *GroupPolicyDefinitionFile) SetLanguageCodes(value []string)() {
    if m != nil {
        m.languageCodes = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPolicyType sets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
func (m *GroupPolicyDefinitionFile) SetPolicyType(value *GroupPolicyType)() {
    if m != nil {
        m.policyType = value
    }
}
// SetRevision sets the revision property value. The revision version associated with the file.
func (m *GroupPolicyDefinitionFile) SetRevision(value *string)() {
    if m != nil {
        m.revision = value
    }
}
// SetTargetNamespace sets the targetNamespace property value. Specifies the URI used to identify the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) SetTargetNamespace(value *string)() {
    if m != nil {
        m.targetNamespace = value
    }
}
// SetTargetPrefix sets the targetPrefix property value. Specifies the logical name that refers to the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) SetTargetPrefix(value *string)() {
    if m != nil {
        m.targetPrefix = value
    }
}
