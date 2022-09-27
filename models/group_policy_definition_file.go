package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinitionFile 
type GroupPolicyDefinitionFile struct {
    Entity
    // The group policy definitions associated with the file.
    definitions []GroupPolicyDefinitionable
    // The localized description of the policy settings in the ADMX file. The default value is empty.
    description *string
    // The localized friendly name of the ADMX file.
    displayName *string
    // The file name of the ADMX file without the path. For example: edge.admx
    fileName *string
    // The supported language codes for the ADMX file.
    languageCodes []string
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Type of Group Policy File or Definition.
    policyType *GroupPolicyType
    // The revision version associated with the file.
    revision *string
    // Specifies the URI used to identify the namespace within the ADMX file.
    targetNamespace *string
    // Specifies the logical name that refers to the namespace within the ADMX file.
    targetPrefix *string
}
// NewGroupPolicyDefinitionFile instantiates a new groupPolicyDefinitionFile and sets the default values.
func NewGroupPolicyDefinitionFile()(*GroupPolicyDefinitionFile) {
    m := &GroupPolicyDefinitionFile{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyDefinitionFile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyDefinitionFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyDefinitionFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.groupPolicyUploadedDefinitionFile":
                        return NewGroupPolicyUploadedDefinitionFile(), nil
                }
            }
        }
    }
    return NewGroupPolicyDefinitionFile(), nil
}
// GetDefinitions gets the definitions property value. The group policy definitions associated with the file.
func (m *GroupPolicyDefinitionFile) GetDefinitions()([]GroupPolicyDefinitionable) {
    return m.definitions
}
// GetDescription gets the description property value. The localized description of the policy settings in the ADMX file. The default value is empty.
func (m *GroupPolicyDefinitionFile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The localized friendly name of the ADMX file.
func (m *GroupPolicyDefinitionFile) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyDefinitionFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyDefinitionFromDiscriminatorValue , m.SetDefinitions)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["languageCodes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetLanguageCodes)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["policyType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseGroupPolicyType , m.SetPolicyType)
    res["revision"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRevision)
    res["targetNamespace"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetNamespace)
    res["targetPrefix"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetPrefix)
    return res
}
// GetFileName gets the fileName property value. The file name of the ADMX file without the path. For example: edge.admx
func (m *GroupPolicyDefinitionFile) GetFileName()(*string) {
    return m.fileName
}
// GetLanguageCodes gets the languageCodes property value. The supported language codes for the ADMX file.
func (m *GroupPolicyDefinitionFile) GetLanguageCodes()([]string) {
    return m.languageCodes
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPolicyType gets the policyType property value. Type of Group Policy File or Definition.
func (m *GroupPolicyDefinitionFile) GetPolicyType()(*GroupPolicyType) {
    return m.policyType
}
// GetRevision gets the revision property value. The revision version associated with the file.
func (m *GroupPolicyDefinitionFile) GetRevision()(*string) {
    return m.revision
}
// GetTargetNamespace gets the targetNamespace property value. Specifies the URI used to identify the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) GetTargetNamespace()(*string) {
    return m.targetNamespace
}
// GetTargetPrefix gets the targetPrefix property value. Specifies the logical name that refers to the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) GetTargetPrefix()(*string) {
    return m.targetPrefix
}
// Serialize serializes information the current object
func (m *GroupPolicyDefinitionFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDefinitions())
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
func (m *GroupPolicyDefinitionFile) SetDefinitions(value []GroupPolicyDefinitionable)() {
    m.definitions = value
}
// SetDescription sets the description property value. The localized description of the policy settings in the ADMX file. The default value is empty.
func (m *GroupPolicyDefinitionFile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The localized friendly name of the ADMX file.
func (m *GroupPolicyDefinitionFile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFileName sets the fileName property value. The file name of the ADMX file without the path. For example: edge.admx
func (m *GroupPolicyDefinitionFile) SetFileName(value *string)() {
    m.fileName = value
}
// SetLanguageCodes sets the languageCodes property value. The supported language codes for the ADMX file.
func (m *GroupPolicyDefinitionFile) SetLanguageCodes(value []string)() {
    m.languageCodes = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPolicyType sets the policyType property value. Type of Group Policy File or Definition.
func (m *GroupPolicyDefinitionFile) SetPolicyType(value *GroupPolicyType)() {
    m.policyType = value
}
// SetRevision sets the revision property value. The revision version associated with the file.
func (m *GroupPolicyDefinitionFile) SetRevision(value *string)() {
    m.revision = value
}
// SetTargetNamespace sets the targetNamespace property value. Specifies the URI used to identify the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) SetTargetNamespace(value *string)() {
    m.targetNamespace = value
}
// SetTargetPrefix sets the targetPrefix property value. Specifies the logical name that refers to the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) SetTargetPrefix(value *string)() {
    m.targetPrefix = value
}
