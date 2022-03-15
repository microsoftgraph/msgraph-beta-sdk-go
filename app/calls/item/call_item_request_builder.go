package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i032e80f26f63621eac368a0e78e5d567ee153d1bcbbb457d6b581d74b6b0dc6a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/contentsharingsessions"
    i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/reject"
    i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/audioroutinggroups"
    i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/unmute"
    i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants"
    i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/cancelmediaprocessing"
    i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/playprompt"
    i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/updaterecordingstatus"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/recordresponse"
    i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/operations"
    i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/record"
    i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/subscribetotone"
    i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/redirect"
    i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/mute"
    iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/keepalive"
    icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/changescreensharingrole"
    iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/answer"
    if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/transfer"
    i0e77db2894bae546ce3028844c9be92c1f453524b9969cb6a1c5fbc30a81428b "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/contentsharingsessions/item"
    i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/operations/item"
    i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/audioroutinggroups/item"
    idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants/item"
)

// CallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.commsApplication entity.
type CallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CallItemRequestBuilderDeleteOptions options for Delete
type CallItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CallItemRequestBuilderGetOptions options for Get
type CallItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CallItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CallItemRequestBuilderGetQueryParameters get calls from app
type CallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CallItemRequestBuilderPatchOptions options for Patch
type CallItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Callable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CallItemRequestBuilder) Answer()(*iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37.AnswerRequestBuilder) {
    return iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) AudioRoutingGroups()(*i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865.AudioRoutingGroupsRequestBuilder) {
    return i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865.NewAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.audioRoutingGroups.item collection
func (m *CallItemRequestBuilder) AudioRoutingGroupsById(id string)(*i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37.AudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup_id"] = id
    }
    return i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37.NewAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallItemRequestBuilder) CancelMediaProcessing()(*i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c.CancelMediaProcessingRequestBuilder) {
    return i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) ChangeScreenSharingRole()(*icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2.ChangeScreenSharingRoleRequestBuilder) {
    return icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallItemRequestBuilder) {
    m := &CallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CallItemRequestBuilder) ContentSharingSessions()(*i032e80f26f63621eac368a0e78e5d567ee153d1bcbbb457d6b581d74b6b0dc6a.ContentSharingSessionsRequestBuilder) {
    return i032e80f26f63621eac368a0e78e5d567ee153d1bcbbb457d6b581d74b6b0dc6a.NewContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentSharingSessionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.contentSharingSessions.item collection
func (m *CallItemRequestBuilder) ContentSharingSessionsById(id string)(*i0e77db2894bae546ce3028844c9be92c1f453524b9969cb6a1c5fbc30a81428b.ContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession_id"] = id
    }
    return i0e77db2894bae546ce3028844c9be92c1f453524b9969cb6a1c5fbc30a81428b.NewContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property calls for app
func (m *CallItemRequestBuilder) CreateDeleteRequestInformation(options *CallItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get calls from app
func (m *CallItemRequestBuilder) CreateGetRequestInformation(options *CallItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property calls in app
func (m *CallItemRequestBuilder) CreatePatchRequestInformation(options *CallItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property calls for app
func (m *CallItemRequestBuilder) Delete(options *CallItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get calls from app
func (m *CallItemRequestBuilder) Get(options *CallItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Callable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCallFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Callable), nil
}
func (m *CallItemRequestBuilder) KeepAlive()(*iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8.KeepAliveRequestBuilder) {
    return iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Mute()(*i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0.MuteRequestBuilder) {
    return i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Operations()(*i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1.OperationsRequestBuilder) {
    return i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.operations.item collection
func (m *CallItemRequestBuilder) OperationsById(id string)(*i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a.CommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation_id"] = id
    }
    return i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a.NewCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Participants()(*i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6.ParticipantsRequestBuilder) {
    return i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6.NewParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.participants.item collection
func (m *CallItemRequestBuilder) ParticipantsById(id string)(*idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a.ParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant_id"] = id
    }
    return idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a.NewParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in app
func (m *CallItemRequestBuilder) Patch(options *CallItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *CallItemRequestBuilder) PlayPrompt()(*i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287.PlayPromptRequestBuilder) {
    return i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Record()(*i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6.RecordRequestBuilder) {
    return i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6.NewRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) RecordResponse()(*i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf.RecordResponseRequestBuilder) {
    return i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Redirect()(*i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662.RedirectRequestBuilder) {
    return i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Reject()(*i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d.RejectRequestBuilder) {
    return i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) SubscribeToTone()(*i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2.SubscribeToToneRequestBuilder) {
    return i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Transfer()(*if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e.TransferRequestBuilder) {
    return if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Unmute()(*i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee.UnmuteRequestBuilder) {
    return i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) UpdateRecordingStatus()(*i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce.UpdateRecordingStatusRequestBuilder) {
    return i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
