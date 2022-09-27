package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAllowOrBlockListAction 
type TenantAllowOrBlockListAction struct {
    // Specifies whether the tenant allow block list is an allow or block. The possible values are: allow, block, and unkownFutureValue.
    action *TenantAllowBlockListAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies when the tenant allow-block-list expires in date time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the note added to the tenant allow block list entry in the format of string.
    note *string
    // The OdataType property
    odataType *string
    // Contains the result of the submission that lead to the tenant allow-block-list entry creation.
    results []TenantAllowBlockListEntryResultable
}
// NewTenantAllowOrBlockListAction instantiates a new tenantAllowOrBlockListAction and sets the default values.
func NewTenantAllowOrBlockListAction()(*TenantAllowOrBlockListAction) {
    m := &TenantAllowOrBlockListAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.security.tenantAllowOrBlockListAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantAllowOrBlockListActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAllowOrBlockListActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAllowOrBlockListAction(), nil
}
// GetAction gets the action property value. Specifies whether the tenant allow block list is an allow or block. The possible values are: allow, block, and unkownFutureValue.
func (m *TenantAllowOrBlockListAction) GetAction()(*TenantAllowBlockListAction) {
    return m.action
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowOrBlockListAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. Specifies when the tenant allow-block-list expires in date time.
func (m *TenantAllowOrBlockListAction) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAllowOrBlockListAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTenantAllowBlockListAction , m.SetAction)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["note"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNote)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["results"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantAllowBlockListEntryResultFromDiscriminatorValue , m.SetResults)
    return res
}
// GetNote gets the note property value. Specifies the note added to the tenant allow block list entry in the format of string.
func (m *TenantAllowOrBlockListAction) GetNote()(*string) {
    return m.note
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantAllowOrBlockListAction) GetOdataType()(*string) {
    return m.odataType
}
// GetResults gets the results property value. Contains the result of the submission that lead to the tenant allow-block-list entry creation.
func (m *TenantAllowOrBlockListAction) GetResults()([]TenantAllowBlockListEntryResultable) {
    return m.results
}
// Serialize serializes information the current object
func (m *TenantAllowOrBlockListAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note", m.GetNote())
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
    if m.GetResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResults())
        err := writer.WriteCollectionOfObjectValues("results", cast)
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
// SetAction sets the action property value. Specifies whether the tenant allow block list is an allow or block. The possible values are: allow, block, and unkownFutureValue.
func (m *TenantAllowOrBlockListAction) SetAction(value *TenantAllowBlockListAction)() {
    m.action = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowOrBlockListAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Specifies when the tenant allow-block-list expires in date time.
func (m *TenantAllowOrBlockListAction) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetNote sets the note property value. Specifies the note added to the tenant allow block list entry in the format of string.
func (m *TenantAllowOrBlockListAction) SetNote(value *string)() {
    m.note = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantAllowOrBlockListAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResults sets the results property value. Contains the result of the submission that lead to the tenant allow-block-list entry creation.
func (m *TenantAllowOrBlockListAction) SetResults(value []TenantAllowBlockListEntryResultable)() {
    m.results = value
}
