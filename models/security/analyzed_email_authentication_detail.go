package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AnalyzedEmailAuthenticationDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAnalyzedEmailAuthenticationDetail instantiates a new AnalyzedEmailAuthenticationDetail and sets the default values.
func NewAnalyzedEmailAuthenticationDetail()(*AnalyzedEmailAuthenticationDetail) {
    m := &AnalyzedEmailAuthenticationDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAnalyzedEmailAuthenticationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnalyzedEmailAuthenticationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnalyzedEmailAuthenticationDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AnalyzedEmailAuthenticationDetail) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AnalyzedEmailAuthenticationDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompositeAuthentication gets the compositeAuthentication property value. A value used by Microsoft 365 to combine email authentication such as SPF, DKIM, and DMARC, to determine whether the message is authentic.
// returns a *string when successful
func (m *AnalyzedEmailAuthenticationDetail) GetCompositeAuthentication()(*string) {
    val, err := m.GetBackingStore().Get("compositeAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDkim gets the dkim property value. DomainKeys identified mail (DKIM). Indicates whether it was pass/fail/soft fail.
// returns a *string when successful
func (m *AnalyzedEmailAuthenticationDetail) GetDkim()(*string) {
    val, err := m.GetBackingStore().Get("dkim")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDmarc gets the dmarc property value. Domain-based Message Authentication. Indicates whether it was pass/fail/soft fail.
// returns a *string when successful
func (m *AnalyzedEmailAuthenticationDetail) GetDmarc()(*string) {
    val, err := m.GetBackingStore().Get("dmarc")
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
func (m *AnalyzedEmailAuthenticationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compositeAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompositeAuthentication(val)
        }
        return nil
    }
    res["dkim"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDkim(val)
        }
        return nil
    }
    res["dmarc"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDmarc(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["senderPolicyFramework"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderPolicyFramework(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AnalyzedEmailAuthenticationDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSenderPolicyFramework gets the senderPolicyFramework property value. Sender Policy Framework (SPF). Indicates whether it was pass/fail/soft fail.
// returns a *string when successful
func (m *AnalyzedEmailAuthenticationDetail) GetSenderPolicyFramework()(*string) {
    val, err := m.GetBackingStore().Get("senderPolicyFramework")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AnalyzedEmailAuthenticationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("compositeAuthentication", m.GetCompositeAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dkim", m.GetDkim())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dmarc", m.GetDmarc())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("senderPolicyFramework", m.GetSenderPolicyFramework())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnalyzedEmailAuthenticationDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AnalyzedEmailAuthenticationDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompositeAuthentication sets the compositeAuthentication property value. A value used by Microsoft 365 to combine email authentication such as SPF, DKIM, and DMARC, to determine whether the message is authentic.
func (m *AnalyzedEmailAuthenticationDetail) SetCompositeAuthentication(value *string)() {
    err := m.GetBackingStore().Set("compositeAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetDkim sets the dkim property value. DomainKeys identified mail (DKIM). Indicates whether it was pass/fail/soft fail.
func (m *AnalyzedEmailAuthenticationDetail) SetDkim(value *string)() {
    err := m.GetBackingStore().Set("dkim", value)
    if err != nil {
        panic(err)
    }
}
// SetDmarc sets the dmarc property value. Domain-based Message Authentication. Indicates whether it was pass/fail/soft fail.
func (m *AnalyzedEmailAuthenticationDetail) SetDmarc(value *string)() {
    err := m.GetBackingStore().Set("dmarc", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AnalyzedEmailAuthenticationDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderPolicyFramework sets the senderPolicyFramework property value. Sender Policy Framework (SPF). Indicates whether it was pass/fail/soft fail.
func (m *AnalyzedEmailAuthenticationDetail) SetSenderPolicyFramework(value *string)() {
    err := m.GetBackingStore().Set("senderPolicyFramework", value)
    if err != nil {
        panic(err)
    }
}
type AnalyzedEmailAuthenticationDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompositeAuthentication()(*string)
    GetDkim()(*string)
    GetDmarc()(*string)
    GetOdataType()(*string)
    GetSenderPolicyFramework()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompositeAuthentication(value *string)()
    SetDkim(value *string)()
    SetDmarc(value *string)()
    SetOdataType(value *string)()
    SetSenderPolicyFramework(value *string)()
}
