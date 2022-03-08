package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActionSetable 
type ItemActionSetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetComment()(CommentActionable)
    GetCreate()(CreateActionable)
    GetDelete()(DeleteActionable)
    GetEdit()(EditActionable)
    GetMention()(MentionActionable)
    GetMove()(MoveActionable)
    GetRename()(RenameActionable)
    GetRestore()(RestoreActionable)
    GetShare()(ShareActionable)
    GetVersion()(VersionActionable)
    SetComment(value CommentActionable)()
    SetCreate(value CreateActionable)()
    SetDelete(value DeleteActionable)()
    SetEdit(value EditActionable)()
    SetMention(value MentionActionable)()
    SetMove(value MoveActionable)()
    SetRename(value RenameActionable)()
    SetRestore(value RestoreActionable)()
    SetShare(value ShareActionable)()
    SetVersion(value VersionActionable)()
}
