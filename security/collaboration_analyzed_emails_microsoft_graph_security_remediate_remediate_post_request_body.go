package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

type CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody instantiates a new CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody and sets the default values.
func NewCollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody()(*CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) {
    m := &CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody(), nil
}
// GetAction gets the action property value. The action property
// returns a *RemediationAction when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetAction()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationAction) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationAction)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAnalyzedEmails gets the analyzedEmails property value. The analyzedEmails property
// returns a []AnalyzedEmailable when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetAnalyzedEmails()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable) {
    val, err := m.GetBackingStore().Get("analyzedEmails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable)
    }
    return nil
}
// GetApproverUpn gets the approverUpn property value. The approverUpn property
// returns a *string when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetApproverUpn()(*string) {
    val, err := m.GetBackingStore().Get("approverUpn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationAction))
        }
        return nil
    }
    res["analyzedEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateAnalyzedEmailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable)
                }
            }
            m.SetAnalyzedEmails(res)
        }
        return nil
    }
    res["approverUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproverUpn(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["remediateSendersCopy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediateSendersCopy(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseRemediationSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationSeverity))
        }
        return nil
    }
    return res
}
// GetRemediateSendersCopy gets the remediateSendersCopy property value. The remediateSendersCopy property
// returns a *bool when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetRemediateSendersCopy()(*bool) {
    val, err := m.GetBackingStore().Get("remediateSendersCopy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSeverity gets the severity property value. The severity property
// returns a *RemediationSeverity when successful
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) GetSeverity()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationSeverity) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationSeverity)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnalyzedEmails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAnalyzedEmails()))
        for i, v := range m.GetAnalyzedEmails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("analyzedEmails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("approverUpn", m.GetApproverUpn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("remediateSendersCopy", m.GetRemediateSendersCopy())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetAction(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationAction)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAnalyzedEmails sets the analyzedEmails property value. The analyzedEmails property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetAnalyzedEmails(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable)() {
    err := m.GetBackingStore().Set("analyzedEmails", value)
    if err != nil {
        panic(err)
    }
}
// SetApproverUpn sets the approverUpn property value. The approverUpn property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetApproverUpn(value *string)() {
    err := m.GetBackingStore().Set("approverUpn", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. The description property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediateSendersCopy sets the remediateSendersCopy property value. The remediateSendersCopy property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetRemediateSendersCopy(value *bool)() {
    err := m.GetBackingStore().Set("remediateSendersCopy", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The severity property
func (m *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBody) SetSeverity(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationSeverity)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
type CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRemediatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationAction)
    GetAnalyzedEmails()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable)
    GetApproverUpn()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetRemediateSendersCopy()(*bool)
    GetSeverity()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationSeverity)
    SetAction(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationAction)()
    SetAnalyzedEmails(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable)()
    SetApproverUpn(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetRemediateSendersCopy(value *bool)()
    SetSeverity(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RemediationSeverity)()
}
