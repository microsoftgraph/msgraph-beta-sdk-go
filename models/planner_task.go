package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerTask struct {
    PlannerDelta
}
// NewPlannerTask instantiates a new PlannerTask and sets the default values.
func NewPlannerTask()(*PlannerTask) {
    m := &PlannerTask{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// CreatePlannerTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.businessScenarioTask":
                        return NewBusinessScenarioTask(), nil
                }
            }
        }
    }
    return NewPlannerTask(), nil
}
// GetActiveChecklistItemCount gets the activeChecklistItemCount property value. The number of checklist items with value set to false, representing incomplete items.
// returns a *int32 when successful
func (m *PlannerTask) GetActiveChecklistItemCount()(*int32) {
    val, err := m.GetBackingStore().Get("activeChecklistItemCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAppliedCategories gets the appliedCategories property value. The categories to which the task is applied. See plannerAppliedCategories resource type for possible values.
// returns a PlannerAppliedCategoriesable when successful
func (m *PlannerTask) GetAppliedCategories()(PlannerAppliedCategoriesable) {
    val, err := m.GetBackingStore().Get("appliedCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerAppliedCategoriesable)
    }
    return nil
}
// GetArchivalInfo gets the archivalInfo property value. Read-only. Nullable. Contains information about who archived or unarchived the task and why.
// returns a PlannerArchivalInfoable when successful
func (m *PlannerTask) GetArchivalInfo()(PlannerArchivalInfoable) {
    val, err := m.GetBackingStore().Get("archivalInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerArchivalInfoable)
    }
    return nil
}
// GetAssignedToTaskBoardFormat gets the assignedToTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
// returns a PlannerAssignedToTaskBoardTaskFormatable when successful
func (m *PlannerTask) GetAssignedToTaskBoardFormat()(PlannerAssignedToTaskBoardTaskFormatable) {
    val, err := m.GetBackingStore().Get("assignedToTaskBoardFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerAssignedToTaskBoardTaskFormatable)
    }
    return nil
}
// GetAssigneePriority gets the assigneePriority property value. A hint that is used to order items of this type in a list view. For more information, see Using order hints in planner.
// returns a *string when successful
func (m *PlannerTask) GetAssigneePriority()(*string) {
    val, err := m.GetBackingStore().Get("assigneePriority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignments gets the assignments property value. The set of assignees the task is assigned to.
// returns a PlannerAssignmentsable when successful
func (m *PlannerTask) GetAssignments()(PlannerAssignmentsable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerAssignmentsable)
    }
    return nil
}
// GetBucketId gets the bucketId property value. Bucket ID to which the task belongs. The bucket needs to be in the same plan as the task. The value of the bucketId property is 28 characters long and case-sensitive. Format validation is done on the service.
// returns a *string when successful
func (m *PlannerTask) GetBucketId()(*string) {
    val, err := m.GetBackingStore().Get("bucketId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBucketTaskBoardFormat gets the bucketTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
// returns a PlannerBucketTaskBoardTaskFormatable when successful
func (m *PlannerTask) GetBucketTaskBoardFormat()(PlannerBucketTaskBoardTaskFormatable) {
    val, err := m.GetBackingStore().Get("bucketTaskBoardFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerBucketTaskBoardTaskFormatable)
    }
    return nil
}
// GetChecklistItemCount gets the checklistItemCount property value. The number of checklist items that are present on the task.
// returns a *int32 when successful
func (m *PlannerTask) GetChecklistItemCount()(*int32) {
    val, err := m.GetBackingStore().Get("checklistItemCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCompletedBy gets the completedBy property value. The identity of the user that completed the task.
// returns a IdentitySetable when successful
func (m *PlannerTask) GetCompletedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("completedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCompletedDateTime gets the completedDateTime property value. Read-only. The date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *PlannerTask) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetConversationThreadId gets the conversationThreadId property value. The thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
// returns a *string when successful
func (m *PlannerTask) GetConversationThreadId()(*string) {
    val, err := m.GetBackingStore().Get("conversationThreadId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The identity of the user who created the task.
// returns a IdentitySetable when successful
func (m *PlannerTask) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Read-only. The date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *PlannerTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreationSource gets the creationSource property value. Information about the origin of the task.
// returns a PlannerTaskCreationable when successful
func (m *PlannerTask) GetCreationSource()(PlannerTaskCreationable) {
    val, err := m.GetBackingStore().Get("creationSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskCreationable)
    }
    return nil
}
// GetDetails gets the details property value. Read-only. Nullable. More details about the task.
// returns a PlannerTaskDetailsable when successful
func (m *PlannerTask) GetDetails()(PlannerTaskDetailsable) {
    val, err := m.GetBackingStore().Get("details")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskDetailsable)
    }
    return nil
}
// GetDueDateTime gets the dueDateTime property value. The date and time at which the task is due. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *PlannerTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("dueDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["activeChecklistItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveChecklistItemCount(val)
        }
        return nil
    }
    res["appliedCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerAppliedCategoriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedCategories(val.(PlannerAppliedCategoriesable))
        }
        return nil
    }
    res["archivalInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerArchivalInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchivalInfo(val.(PlannerArchivalInfoable))
        }
        return nil
    }
    res["assignedToTaskBoardFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedToTaskBoardFormat(val.(PlannerAssignedToTaskBoardTaskFormatable))
        }
        return nil
    }
    res["assigneePriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneePriority(val)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignments(val.(PlannerAssignmentsable))
        }
        return nil
    }
    res["bucketId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketId(val)
        }
        return nil
    }
    res["bucketTaskBoardFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerBucketTaskBoardTaskFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketTaskBoardFormat(val.(PlannerBucketTaskBoardTaskFormatable))
        }
        return nil
    }
    res["checklistItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklistItemCount(val)
        }
        return nil
    }
    res["completedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["conversationThreadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationThreadId(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["creationSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskCreationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSource(val.(PlannerTaskCreationable))
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(PlannerTaskDetailsable))
        }
        return nil
    }
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
        }
        return nil
    }
    res["hasDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDescription(val)
        }
        return nil
    }
    res["isArchived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["isOnMyDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnMyDay(val)
        }
        return nil
    }
    res["isOnMyDayLastModifiedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnMyDayLastModifiedDate(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["orderHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHint(val)
        }
        return nil
    }
    res["percentComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentComplete(val)
        }
        return nil
    }
    res["planId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanId(val)
        }
        return nil
    }
    res["previewType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewType(val.(*PlannerPreviewType))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["progressTaskBoardFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerProgressTaskBoardTaskFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgressTaskBoardFormat(val.(PlannerProgressTaskBoardTaskFormatable))
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PlannerTaskRecurrenceable))
        }
        return nil
    }
    res["referenceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceCount(val)
        }
        return nil
    }
    res["specifiedCompletionRequirements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerTaskCompletionRequirements)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecifiedCompletionRequirements(val.(*PlannerTaskCompletionRequirements))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetHasDescription gets the hasDescription property value. Read-only. This value is true if the details object of the task has a nonempty description. Otherwise,false.
// returns a *bool when successful
func (m *PlannerTask) GetHasDescription()(*bool) {
    val, err := m.GetBackingStore().Get("hasDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsArchived gets the isArchived property value. Read-only. If set to true, the task is archived. An archived task is read-only.
// returns a *bool when successful
func (m *PlannerTask) GetIsArchived()(*bool) {
    val, err := m.GetBackingStore().Get("isArchived")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsOnMyDay gets the isOnMyDay property value. Indicates whether to show this task in the MyDay view. If true, it shows the task.
// returns a *bool when successful
func (m *PlannerTask) GetIsOnMyDay()(*bool) {
    val, err := m.GetBackingStore().Get("isOnMyDay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsOnMyDayLastModifiedDate gets the isOnMyDayLastModifiedDate property value. Read-only. The date on which task is added to or removed from MyDay.
// returns a *DateOnly when successful
func (m *PlannerTask) GetIsOnMyDayLastModifiedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("isOnMyDayLastModifiedDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
// returns a IdentitySetable when successful
func (m *PlannerTask) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *PlannerTask) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOrderHint gets the orderHint property value. The hint used to order items of this type in a list view. For more information, see Using order hints in plannern.
// returns a *string when successful
func (m *PlannerTask) GetOrderHint()(*string) {
    val, err := m.GetBackingStore().Get("orderHint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPercentComplete gets the percentComplete property value. The percentage of task completion. When set to 100, the task is completed.
// returns a *int32 when successful
func (m *PlannerTask) GetPercentComplete()(*int32) {
    val, err := m.GetBackingStore().Get("percentComplete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPlanId gets the planId property value. Plan ID to which the task belongs.
// returns a *string when successful
func (m *PlannerTask) GetPlanId()(*string) {
    val, err := m.GetBackingStore().Get("planId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreviewType gets the previewType property value. The type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference.
// returns a *PlannerPreviewType when successful
func (m *PlannerTask) GetPreviewType()(*PlannerPreviewType) {
    val, err := m.GetBackingStore().Get("previewType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PlannerPreviewType)
    }
    return nil
}
// GetPriority gets the priority property value. The priority of the task. Valid values are between 0 and 10, inclusive. Larger values indicate lower priority. For example, 0 has the highest priority and 10 has the lowest priority. Currently, planner interprets values 0 and 1 as 'urgent', 2 and 3 and 4 as 'important', 5, 6, and 7 as 'medium', and 8, 9, and 10 as 'low'. Currently, planner sets the value 1 for 'urgent', 3 for 'important', 5 for 'medium', and 9 for 'low'.
// returns a *int32 when successful
func (m *PlannerTask) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetProgressTaskBoardFormat gets the progressTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
// returns a PlannerProgressTaskBoardTaskFormatable when successful
func (m *PlannerTask) GetProgressTaskBoardFormat()(PlannerProgressTaskBoardTaskFormatable) {
    val, err := m.GetBackingStore().Get("progressTaskBoardFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerProgressTaskBoardTaskFormatable)
    }
    return nil
}
// GetRecurrence gets the recurrence property value. Defines active or inactive recurrence for the task. null when the recurrence has never been defined for the task.
// returns a PlannerTaskRecurrenceable when successful
func (m *PlannerTask) GetRecurrence()(PlannerTaskRecurrenceable) {
    val, err := m.GetBackingStore().Get("recurrence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskRecurrenceable)
    }
    return nil
}
// GetReferenceCount gets the referenceCount property value. Number of external references that exist on the task.
// returns a *int32 when successful
func (m *PlannerTask) GetReferenceCount()(*int32) {
    val, err := m.GetBackingStore().Get("referenceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSpecifiedCompletionRequirements gets the specifiedCompletionRequirements property value. Indicates all the requirements specified on the plannerTask. Possible values are: none, checklistCompletion, unknownFutureValue, formCompletion, approvalCompletion. Read-only. Use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: formCompletion, approvalCompletion. The plannerTaskCompletionRequirementDetails in plannerTaskDetails has details of the requirements specified, if any.
// returns a *PlannerTaskCompletionRequirements when successful
func (m *PlannerTask) GetSpecifiedCompletionRequirements()(*PlannerTaskCompletionRequirements) {
    val, err := m.GetBackingStore().Get("specifiedCompletionRequirements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PlannerTaskCompletionRequirements)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *PlannerTask) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTitle gets the title property value. Title of the task.
// returns a *string when successful
func (m *PlannerTask) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeChecklistItemCount", m.GetActiveChecklistItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appliedCategories", m.GetAppliedCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("archivalInfo", m.GetArchivalInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignedToTaskBoardFormat", m.GetAssignedToTaskBoardFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assigneePriority", m.GetAssigneePriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignments", m.GetAssignments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bucketId", m.GetBucketId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bucketTaskBoardFormat", m.GetBucketTaskBoardFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("checklistItemCount", m.GetChecklistItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedBy", m.GetCompletedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conversationThreadId", m.GetConversationThreadId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("creationSource", m.GetCreationSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasDescription", m.GetHasDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOnMyDay", m.GetIsOnMyDay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("isOnMyDayLastModifiedDate", m.GetIsOnMyDayLastModifiedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planId", m.GetPlanId())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        cast := (*m.GetPreviewType()).String()
        err = writer.WriteStringValue("previewType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("progressTaskBoardFormat", m.GetProgressTaskBoardFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("referenceCount", m.GetReferenceCount())
        if err != nil {
            return err
        }
    }
    if m.GetSpecifiedCompletionRequirements() != nil {
        cast := (*m.GetSpecifiedCompletionRequirements()).String()
        err = writer.WriteStringValue("specifiedCompletionRequirements", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveChecklistItemCount sets the activeChecklistItemCount property value. The number of checklist items with value set to false, representing incomplete items.
func (m *PlannerTask) SetActiveChecklistItemCount(value *int32)() {
    err := m.GetBackingStore().Set("activeChecklistItemCount", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedCategories sets the appliedCategories property value. The categories to which the task is applied. See plannerAppliedCategories resource type for possible values.
func (m *PlannerTask) SetAppliedCategories(value PlannerAppliedCategoriesable)() {
    err := m.GetBackingStore().Set("appliedCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetArchivalInfo sets the archivalInfo property value. Read-only. Nullable. Contains information about who archived or unarchived the task and why.
func (m *PlannerTask) SetArchivalInfo(value PlannerArchivalInfoable)() {
    err := m.GetBackingStore().Set("archivalInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedToTaskBoardFormat sets the assignedToTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
func (m *PlannerTask) SetAssignedToTaskBoardFormat(value PlannerAssignedToTaskBoardTaskFormatable)() {
    err := m.GetBackingStore().Set("assignedToTaskBoardFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneePriority sets the assigneePriority property value. A hint that is used to order items of this type in a list view. For more information, see Using order hints in planner.
func (m *PlannerTask) SetAssigneePriority(value *string)() {
    err := m.GetBackingStore().Set("assigneePriority", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. The set of assignees the task is assigned to.
func (m *PlannerTask) SetAssignments(value PlannerAssignmentsable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetBucketId sets the bucketId property value. Bucket ID to which the task belongs. The bucket needs to be in the same plan as the task. The value of the bucketId property is 28 characters long and case-sensitive. Format validation is done on the service.
func (m *PlannerTask) SetBucketId(value *string)() {
    err := m.GetBackingStore().Set("bucketId", value)
    if err != nil {
        panic(err)
    }
}
// SetBucketTaskBoardFormat sets the bucketTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
func (m *PlannerTask) SetBucketTaskBoardFormat(value PlannerBucketTaskBoardTaskFormatable)() {
    err := m.GetBackingStore().Set("bucketTaskBoardFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetChecklistItemCount sets the checklistItemCount property value. The number of checklist items that are present on the task.
func (m *PlannerTask) SetChecklistItemCount(value *int32)() {
    err := m.GetBackingStore().Set("checklistItemCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedBy sets the completedBy property value. The identity of the user that completed the task.
func (m *PlannerTask) SetCompletedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("completedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedDateTime sets the completedDateTime property value. Read-only. The date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetConversationThreadId sets the conversationThreadId property value. The thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
func (m *PlannerTask) SetConversationThreadId(value *string)() {
    err := m.GetBackingStore().Set("conversationThreadId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The identity of the user who created the task.
func (m *PlannerTask) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Read-only. The date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationSource sets the creationSource property value. Information about the origin of the task.
func (m *PlannerTask) SetCreationSource(value PlannerTaskCreationable)() {
    err := m.GetBackingStore().Set("creationSource", value)
    if err != nil {
        panic(err)
    }
}
// SetDetails sets the details property value. Read-only. Nullable. More details about the task.
func (m *PlannerTask) SetDetails(value PlannerTaskDetailsable)() {
    err := m.GetBackingStore().Set("details", value)
    if err != nil {
        panic(err)
    }
}
// SetDueDateTime sets the dueDateTime property value. The date and time at which the task is due. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("dueDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetHasDescription sets the hasDescription property value. Read-only. This value is true if the details object of the task has a nonempty description. Otherwise,false.
func (m *PlannerTask) SetHasDescription(value *bool)() {
    err := m.GetBackingStore().Set("hasDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetIsArchived sets the isArchived property value. Read-only. If set to true, the task is archived. An archived task is read-only.
func (m *PlannerTask) SetIsArchived(value *bool)() {
    err := m.GetBackingStore().Set("isArchived", value)
    if err != nil {
        panic(err)
    }
}
// SetIsOnMyDay sets the isOnMyDay property value. Indicates whether to show this task in the MyDay view. If true, it shows the task.
func (m *PlannerTask) SetIsOnMyDay(value *bool)() {
    err := m.GetBackingStore().Set("isOnMyDay", value)
    if err != nil {
        panic(err)
    }
}
// SetIsOnMyDayLastModifiedDate sets the isOnMyDayLastModifiedDate property value. Read-only. The date on which task is added to or removed from MyDay.
func (m *PlannerTask) SetIsOnMyDayLastModifiedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("isOnMyDayLastModifiedDate", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *PlannerTask) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PlannerTask) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOrderHint sets the orderHint property value. The hint used to order items of this type in a list view. For more information, see Using order hints in plannern.
func (m *PlannerTask) SetOrderHint(value *string)() {
    err := m.GetBackingStore().Set("orderHint", value)
    if err != nil {
        panic(err)
    }
}
// SetPercentComplete sets the percentComplete property value. The percentage of task completion. When set to 100, the task is completed.
func (m *PlannerTask) SetPercentComplete(value *int32)() {
    err := m.GetBackingStore().Set("percentComplete", value)
    if err != nil {
        panic(err)
    }
}
// SetPlanId sets the planId property value. Plan ID to which the task belongs.
func (m *PlannerTask) SetPlanId(value *string)() {
    err := m.GetBackingStore().Set("planId", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviewType sets the previewType property value. The type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference.
func (m *PlannerTask) SetPreviewType(value *PlannerPreviewType)() {
    err := m.GetBackingStore().Set("previewType", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority of the task. Valid values are between 0 and 10, inclusive. Larger values indicate lower priority. For example, 0 has the highest priority and 10 has the lowest priority. Currently, planner interprets values 0 and 1 as 'urgent', 2 and 3 and 4 as 'important', 5, 6, and 7 as 'medium', and 8, 9, and 10 as 'low'. Currently, planner sets the value 1 for 'urgent', 3 for 'important', 5 for 'medium', and 9 for 'low'.
func (m *PlannerTask) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetProgressTaskBoardFormat sets the progressTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
func (m *PlannerTask) SetProgressTaskBoardFormat(value PlannerProgressTaskBoardTaskFormatable)() {
    err := m.GetBackingStore().Set("progressTaskBoardFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrence sets the recurrence property value. Defines active or inactive recurrence for the task. null when the recurrence has never been defined for the task.
func (m *PlannerTask) SetRecurrence(value PlannerTaskRecurrenceable)() {
    err := m.GetBackingStore().Set("recurrence", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceCount sets the referenceCount property value. Number of external references that exist on the task.
func (m *PlannerTask) SetReferenceCount(value *int32)() {
    err := m.GetBackingStore().Set("referenceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSpecifiedCompletionRequirements sets the specifiedCompletionRequirements property value. Indicates all the requirements specified on the plannerTask. Possible values are: none, checklistCompletion, unknownFutureValue, formCompletion, approvalCompletion. Read-only. Use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: formCompletion, approvalCompletion. The plannerTaskCompletionRequirementDetails in plannerTaskDetails has details of the requirements specified, if any.
func (m *PlannerTask) SetSpecifiedCompletionRequirements(value *PlannerTaskCompletionRequirements)() {
    err := m.GetBackingStore().Set("specifiedCompletionRequirements", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Title of the task.
func (m *PlannerTask) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
type PlannerTaskable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetActiveChecklistItemCount()(*int32)
    GetAppliedCategories()(PlannerAppliedCategoriesable)
    GetArchivalInfo()(PlannerArchivalInfoable)
    GetAssignedToTaskBoardFormat()(PlannerAssignedToTaskBoardTaskFormatable)
    GetAssigneePriority()(*string)
    GetAssignments()(PlannerAssignmentsable)
    GetBucketId()(*string)
    GetBucketTaskBoardFormat()(PlannerBucketTaskBoardTaskFormatable)
    GetChecklistItemCount()(*int32)
    GetCompletedBy()(IdentitySetable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConversationThreadId()(*string)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreationSource()(PlannerTaskCreationable)
    GetDetails()(PlannerTaskDetailsable)
    GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHasDescription()(*bool)
    GetIsArchived()(*bool)
    GetIsOnMyDay()(*bool)
    GetIsOnMyDayLastModifiedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOrderHint()(*string)
    GetPercentComplete()(*int32)
    GetPlanId()(*string)
    GetPreviewType()(*PlannerPreviewType)
    GetPriority()(*int32)
    GetProgressTaskBoardFormat()(PlannerProgressTaskBoardTaskFormatable)
    GetRecurrence()(PlannerTaskRecurrenceable)
    GetReferenceCount()(*int32)
    GetSpecifiedCompletionRequirements()(*PlannerTaskCompletionRequirements)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTitle()(*string)
    SetActiveChecklistItemCount(value *int32)()
    SetAppliedCategories(value PlannerAppliedCategoriesable)()
    SetArchivalInfo(value PlannerArchivalInfoable)()
    SetAssignedToTaskBoardFormat(value PlannerAssignedToTaskBoardTaskFormatable)()
    SetAssigneePriority(value *string)()
    SetAssignments(value PlannerAssignmentsable)()
    SetBucketId(value *string)()
    SetBucketTaskBoardFormat(value PlannerBucketTaskBoardTaskFormatable)()
    SetChecklistItemCount(value *int32)()
    SetCompletedBy(value IdentitySetable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConversationThreadId(value *string)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreationSource(value PlannerTaskCreationable)()
    SetDetails(value PlannerTaskDetailsable)()
    SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHasDescription(value *bool)()
    SetIsArchived(value *bool)()
    SetIsOnMyDay(value *bool)()
    SetIsOnMyDayLastModifiedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOrderHint(value *string)()
    SetPercentComplete(value *int32)()
    SetPlanId(value *string)()
    SetPreviewType(value *PlannerPreviewType)()
    SetPriority(value *int32)()
    SetProgressTaskBoardFormat(value PlannerProgressTaskBoardTaskFormatable)()
    SetRecurrence(value PlannerTaskRecurrenceable)()
    SetReferenceCount(value *int32)()
    SetSpecifiedCompletionRequirements(value *PlannerTaskCompletionRequirements)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTitle(value *string)()
}
