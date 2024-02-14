package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DetectionAction struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDetectionAction instantiates a new DetectionAction and sets the default values.
func NewDetectionAction()(*DetectionAction) {
    m := &DetectionAction{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDetectionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDetectionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetectionAction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DetectionAction) GetAdditionalData()(map[string]any) {
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
// GetAlertTemplate gets the alertTemplate property value. The alertTemplate property
// returns a AlertTemplateable when successful
func (m *DetectionAction) GetAlertTemplate()(AlertTemplateable) {
    val, err := m.GetBackingStore().Get("alertTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AlertTemplateable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DetectionAction) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DetectionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlertTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertTemplate(val.(AlertTemplateable))
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
    res["organizationalScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalScope(val.(OrganizationalScopeable))
        }
        return nil
    }
    res["responseActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResponseActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResponseActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResponseActionable)
                }
            }
            m.SetResponseActions(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DetectionAction) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationalScope gets the organizationalScope property value. Groups to which the custom detection rule applies.
// returns a OrganizationalScopeable when successful
func (m *DetectionAction) GetOrganizationalScope()(OrganizationalScopeable) {
    val, err := m.GetBackingStore().Get("organizationalScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OrganizationalScopeable)
    }
    return nil
}
// GetResponseActions gets the responseActions property value. Actions taken on impacted assets as set in the custom detection rule.
// returns a []ResponseActionable when successful
func (m *DetectionAction) GetResponseActions()([]ResponseActionable) {
    val, err := m.GetBackingStore().Get("responseActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ResponseActionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DetectionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("alertTemplate", m.GetAlertTemplate())
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
        err := writer.WriteObjectValue("organizationalScope", m.GetOrganizationalScope())
        if err != nil {
            return err
        }
    }
    if m.GetResponseActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResponseActions()))
        for i, v := range m.GetResponseActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("responseActions", cast)
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
func (m *DetectionAction) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertTemplate sets the alertTemplate property value. The alertTemplate property
func (m *DetectionAction) SetAlertTemplate(value AlertTemplateable)() {
    err := m.GetBackingStore().Set("alertTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DetectionAction) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DetectionAction) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalScope sets the organizationalScope property value. Groups to which the custom detection rule applies.
func (m *DetectionAction) SetOrganizationalScope(value OrganizationalScopeable)() {
    err := m.GetBackingStore().Set("organizationalScope", value)
    if err != nil {
        panic(err)
    }
}
// SetResponseActions sets the responseActions property value. Actions taken on impacted assets as set in the custom detection rule.
func (m *DetectionAction) SetResponseActions(value []ResponseActionable)() {
    err := m.GetBackingStore().Set("responseActions", value)
    if err != nil {
        panic(err)
    }
}
type DetectionActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertTemplate()(AlertTemplateable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetOrganizationalScope()(OrganizationalScopeable)
    GetResponseActions()([]ResponseActionable)
    SetAlertTemplate(value AlertTemplateable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetOrganizationalScope(value OrganizationalScopeable)()
    SetResponseActions(value []ResponseActionable)()
}
