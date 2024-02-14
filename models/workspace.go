package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Workspace struct {
    Place
}
// NewWorkspace instantiates a new Workspace and sets the default values.
func NewWorkspace()(*Workspace) {
    m := &Workspace{
        Place: *NewPlace(),
    }
    odataTypeValue := "#microsoft.graph.workspace"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWorkspaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkspaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkspace(), nil
}
// GetBuilding gets the building property value. Specifies the building name or building number that the workspace is in.
// returns a *string when successful
func (m *Workspace) GetBuilding()(*string) {
    val, err := m.GetBackingStore().Get("building")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCapacity gets the capacity property value. Specifies the capacity of the workspace.
// returns a *int32 when successful
func (m *Workspace) GetCapacity()(*int32) {
    val, err := m.GetBackingStore().Get("capacity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEmailAddress gets the emailAddress property value. Email address of the workspace.
// returns a *string when successful
func (m *Workspace) GetEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("emailAddress")
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
func (m *Workspace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Place.GetFieldDeserializers()
    res["building"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuilding(val)
        }
        return nil
    }
    res["capacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacity(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["floorLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorLabel(val)
        }
        return nil
    }
    res["floorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorNumber(val)
        }
        return nil
    }
    res["isWheelChairAccessible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsWheelChairAccessible(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["nickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickname(val)
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
    return res
}
// GetFloorLabel gets the floorLabel property value. Specifies a descriptive label for the floor, for example, P.
// returns a *string when successful
func (m *Workspace) GetFloorLabel()(*string) {
    val, err := m.GetBackingStore().Get("floorLabel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFloorNumber gets the floorNumber property value. Specifies the floor number that the workspace is on.
// returns a *int32 when successful
func (m *Workspace) GetFloorNumber()(*int32) {
    val, err := m.GetBackingStore().Get("floorNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIsWheelChairAccessible gets the isWheelChairAccessible property value. Specifies whether the workspace is wheelchair accessible.
// returns a *bool when successful
func (m *Workspace) GetIsWheelChairAccessible()(*bool) {
    val, err := m.GetBackingStore().Get("isWheelChairAccessible")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLabel gets the label property value. Specifies a descriptive label for the workspace, for example, a number or name.
// returns a *string when successful
func (m *Workspace) GetLabel()(*string) {
    val, err := m.GetBackingStore().Get("label")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNickname gets the nickname property value. Specifies a nickname for the workspace, for example, 'quiet workspace'.
// returns a *string when successful
func (m *Workspace) GetNickname()(*string) {
    val, err := m.GetBackingStore().Get("nickname")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTags gets the tags property value. Specifies other features of the workspace; for example, the type of view or furniture type.
// returns a []string when successful
func (m *Workspace) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Workspace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Place.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("building", m.GetBuilding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("capacity", m.GetCapacity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("floorLabel", m.GetFloorLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("floorNumber", m.GetFloorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isWheelChairAccessible", m.GetIsWheelChairAccessible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickname", m.GetNickname())
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
    return nil
}
// SetBuilding sets the building property value. Specifies the building name or building number that the workspace is in.
func (m *Workspace) SetBuilding(value *string)() {
    err := m.GetBackingStore().Set("building", value)
    if err != nil {
        panic(err)
    }
}
// SetCapacity sets the capacity property value. Specifies the capacity of the workspace.
func (m *Workspace) SetCapacity(value *int32)() {
    err := m.GetBackingStore().Set("capacity", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddress sets the emailAddress property value. Email address of the workspace.
func (m *Workspace) SetEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("emailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetFloorLabel sets the floorLabel property value. Specifies a descriptive label for the floor, for example, P.
func (m *Workspace) SetFloorLabel(value *string)() {
    err := m.GetBackingStore().Set("floorLabel", value)
    if err != nil {
        panic(err)
    }
}
// SetFloorNumber sets the floorNumber property value. Specifies the floor number that the workspace is on.
func (m *Workspace) SetFloorNumber(value *int32)() {
    err := m.GetBackingStore().Set("floorNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetIsWheelChairAccessible sets the isWheelChairAccessible property value. Specifies whether the workspace is wheelchair accessible.
func (m *Workspace) SetIsWheelChairAccessible(value *bool)() {
    err := m.GetBackingStore().Set("isWheelChairAccessible", value)
    if err != nil {
        panic(err)
    }
}
// SetLabel sets the label property value. Specifies a descriptive label for the workspace, for example, a number or name.
func (m *Workspace) SetLabel(value *string)() {
    err := m.GetBackingStore().Set("label", value)
    if err != nil {
        panic(err)
    }
}
// SetNickname sets the nickname property value. Specifies a nickname for the workspace, for example, 'quiet workspace'.
func (m *Workspace) SetNickname(value *string)() {
    err := m.GetBackingStore().Set("nickname", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. Specifies other features of the workspace; for example, the type of view or furniture type.
func (m *Workspace) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
type Workspaceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Placeable
    GetBuilding()(*string)
    GetCapacity()(*int32)
    GetEmailAddress()(*string)
    GetFloorLabel()(*string)
    GetFloorNumber()(*int32)
    GetIsWheelChairAccessible()(*bool)
    GetLabel()(*string)
    GetNickname()(*string)
    GetTags()([]string)
    SetBuilding(value *string)()
    SetCapacity(value *int32)()
    SetEmailAddress(value *string)()
    SetFloorLabel(value *string)()
    SetFloorNumber(value *int32)()
    SetIsWheelChairAccessible(value *bool)()
    SetLabel(value *string)()
    SetNickname(value *string)()
    SetTags(value []string)()
}
