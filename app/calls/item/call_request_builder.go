package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/reject"
    i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/audioroutinggroups"
    i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/unmute"
    i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants"
    i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/cancelmediaprocessing"
    i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/playprompt"
    i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/updaterecordingstatus"
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
    i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/operations/item"
    i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/audioroutinggroups/item"
    idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants/item"
)

// CallRequestBuilder builds and executes requests for operations under \app\calls\{call-id}
type CallRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CallRequestBuilderDeleteOptions options for Delete
type CallRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CallRequestBuilderGetOptions options for Get
type CallRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CallRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CallRequestBuilderGetQueryParameters get calls from app
type CallRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CallRequestBuilderPatchOptions options for Patch
type CallRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Call;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CallRequestBuilder) Answer()(*iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37.AnswerRequestBuilder) {
    return iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) AudioRoutingGroups()(*i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865.AudioRoutingGroupsRequestBuilder) {
    return i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865.NewAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.audioRoutingGroups.item collection
func (m *CallRequestBuilder) AudioRoutingGroupsById(id string)(*i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37.AudioRoutingGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup_id"] = id
    }
    return i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37.NewAudioRoutingGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallRequestBuilder) CancelMediaProcessing()(*i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c.CancelMediaProcessingRequestBuilder) {
    return i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) ChangeScreenSharingRole()(*icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2.ChangeScreenSharingRoleRequestBuilder) {
    return icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallRequestBuilderInternal instantiates a new CallRequestBuilder and sets the default values.
func NewCallRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallRequestBuilder) {
    m := &CallRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallRequestBuilder instantiates a new CallRequestBuilder and sets the default values.
func NewCallRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calls for app
func (m *CallRequestBuilder) CreateDeleteRequestInformation(options *CallRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CallRequestBuilder) CreateGetRequestInformation(options *CallRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CallRequestBuilder) CreatePatchRequestInformation(options *CallRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CallRequestBuilder) Delete(options *CallRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get calls from app
func (m *CallRequestBuilder) Get(options *CallRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Call, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCall() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Call), nil
}
func (m *CallRequestBuilder) KeepAlive()(*iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8.KeepAliveRequestBuilder) {
    return iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Mute()(*i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0.MuteRequestBuilder) {
    return i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Operations()(*i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1.OperationsRequestBuilder) {
    return i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.operations.item collection
func (m *CallRequestBuilder) OperationsById(id string)(*i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a.CommsOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation_id"] = id
    }
    return i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a.NewCommsOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallRequestBuilder) Participants()(*i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6.ParticipantsRequestBuilder) {
    return i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6.NewParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.participants.item collection
func (m *CallRequestBuilder) ParticipantsById(id string)(*idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a.ParticipantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant_id"] = id
    }
    return idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a.NewParticipantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in app
func (m *CallRequestBuilder) Patch(options *CallRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CallRequestBuilder) PlayPrompt()(*i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287.PlayPromptRequestBuilder) {
    return i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Record()(*i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6.RecordRequestBuilder) {
    return i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6.NewRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) RecordResponse()(*i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf.RecordResponseRequestBuilder) {
    return i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Redirect()(*i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662.RedirectRequestBuilder) {
    return i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Reject()(*i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d.RejectRequestBuilder) {
    return i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) SubscribeToTone()(*i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2.SubscribeToToneRequestBuilder) {
    return i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Transfer()(*if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e.TransferRequestBuilder) {
    return if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Unmute()(*i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee.UnmuteRequestBuilder) {
    return i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) UpdateRecordingStatus()(*i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce.UpdateRecordingStatusRequestBuilder) {
    return i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
