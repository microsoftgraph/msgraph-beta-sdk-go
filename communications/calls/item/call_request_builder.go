package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i08e66c5d7bd1abccfebd8fb93bd3824b29263d4a78a8e0619b6492ed8f3fad49 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/playprompt"
    i0c63fde205100ca16ab7a423f47cbdbde7db96e5bfd78a3f404e33371a9a304c "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/record"
    i18782668e274df0b14fede1b2d0b04b9dfad365a601b91cb1ae1236333912f77 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/recordresponse"
    i2ad07b1a9d626a23f6fb4c6980016bdfa91427506827a8183fe0ab85ccea7297 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/mute"
    i2e0863e08fc93d954484ce570f9308574c3c610b9cbf1e2c9d07876c640463ac "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/updaterecordingstatus"
    i3af510b639a3bfd33b6d77bb52f8d92c7f9e25367174bdfa6142903edc74995c "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/keepalive"
    i4cae3b1a69ba2254d74646c3ba02c22835c51e378675f8f86fa207a4bf13fbb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/unmute"
    i4ffb8669c93d0b11c0238e80369824011ea5e33bda9395560f9abe880eb8acc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/transfer"
    i51c2a09e4f41fd4c07a21a333066d7d9b5683f48919e40d555e43eeedc2d2b0f "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/redirect"
    i59f1af4a71937287f9a647acee47040de638eb6d16981d2dc3f1c22d63d3d210 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/subscribetotone"
    i6a96c1601820e25e2c11e319760d663fd67b6cdb677b5d599e7f4560698143e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/changescreensharingrole"
    i7b15e8c80b7aa50a0151a753f97ef7289d17791928250d62926b97f749d6c305 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/answer"
    i8659ea4008fd7459c7570b3b977a99568235e5600beaac36764abf7a09a1c6a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/operations"
    ic54ac4adcd492f82a003801c0ce35301e9f04c5676f80853066311daf9c569da "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/audioroutinggroups"
    id1e985c21267b30b9212f7e3d69af1960754fba433a529ab9f2819e92d7804e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/cancelmediaprocessing"
    idaf0ca56e2f68098afebd431de7c548ac38196c36423e844dd58942eb2a4ee85 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/reject"
    if78246d1256eb5bc7db2b2c1dcd3d9daaa6de508f121f6f0bb97fcc022fe09f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/participants"
    i151a479d1626e685450cad3f10b2ae807a5841c52441c261bfec932663c91e9f "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/audioroutinggroups/item"
    ie3ddf2285119a257859e618e7d13130592dc1f6e886e271f46d05000453bffd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/participants/item"
    ifc54b1625a5ef155b182402b189053da9800796fb7507ed629f84919426791d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/operations/item"
)

// CallRequestBuilder builds and executes requests for operations under \communications\calls\{call-id}
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
// CallRequestBuilderGetQueryParameters get calls from communications
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
func (m *CallRequestBuilder) Answer()(*i7b15e8c80b7aa50a0151a753f97ef7289d17791928250d62926b97f749d6c305.AnswerRequestBuilder) {
    return i7b15e8c80b7aa50a0151a753f97ef7289d17791928250d62926b97f749d6c305.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) AudioRoutingGroups()(*ic54ac4adcd492f82a003801c0ce35301e9f04c5676f80853066311daf9c569da.AudioRoutingGroupsRequestBuilder) {
    return ic54ac4adcd492f82a003801c0ce35301e9f04c5676f80853066311daf9c569da.NewAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.audioRoutingGroups.item collection
func (m *CallRequestBuilder) AudioRoutingGroupsById(id string)(*i151a479d1626e685450cad3f10b2ae807a5841c52441c261bfec932663c91e9f.AudioRoutingGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup_id"] = id
    }
    return i151a479d1626e685450cad3f10b2ae807a5841c52441c261bfec932663c91e9f.NewAudioRoutingGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallRequestBuilder) CancelMediaProcessing()(*id1e985c21267b30b9212f7e3d69af1960754fba433a529ab9f2819e92d7804e6.CancelMediaProcessingRequestBuilder) {
    return id1e985c21267b30b9212f7e3d69af1960754fba433a529ab9f2819e92d7804e6.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) ChangeScreenSharingRole()(*i6a96c1601820e25e2c11e319760d663fd67b6cdb677b5d599e7f4560698143e1.ChangeScreenSharingRoleRequestBuilder) {
    return i6a96c1601820e25e2c11e319760d663fd67b6cdb677b5d599e7f4560698143e1.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallRequestBuilderInternal instantiates a new CallRequestBuilder and sets the default values.
func NewCallRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallRequestBuilder) {
    m := &CallRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property calls for communications
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
// CreateGetRequestInformation get calls from communications
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
// CreatePatchRequestInformation update the navigation property calls in communications
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
// Delete delete navigation property calls for communications
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
// Get get calls from communications
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
func (m *CallRequestBuilder) KeepAlive()(*i3af510b639a3bfd33b6d77bb52f8d92c7f9e25367174bdfa6142903edc74995c.KeepAliveRequestBuilder) {
    return i3af510b639a3bfd33b6d77bb52f8d92c7f9e25367174bdfa6142903edc74995c.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Mute()(*i2ad07b1a9d626a23f6fb4c6980016bdfa91427506827a8183fe0ab85ccea7297.MuteRequestBuilder) {
    return i2ad07b1a9d626a23f6fb4c6980016bdfa91427506827a8183fe0ab85ccea7297.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Operations()(*i8659ea4008fd7459c7570b3b977a99568235e5600beaac36764abf7a09a1c6a2.OperationsRequestBuilder) {
    return i8659ea4008fd7459c7570b3b977a99568235e5600beaac36764abf7a09a1c6a2.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.operations.item collection
func (m *CallRequestBuilder) OperationsById(id string)(*ifc54b1625a5ef155b182402b189053da9800796fb7507ed629f84919426791d1.CommsOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation_id"] = id
    }
    return ifc54b1625a5ef155b182402b189053da9800796fb7507ed629f84919426791d1.NewCommsOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallRequestBuilder) Participants()(*if78246d1256eb5bc7db2b2c1dcd3d9daaa6de508f121f6f0bb97fcc022fe09f5.ParticipantsRequestBuilder) {
    return if78246d1256eb5bc7db2b2c1dcd3d9daaa6de508f121f6f0bb97fcc022fe09f5.NewParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.participants.item collection
func (m *CallRequestBuilder) ParticipantsById(id string)(*ie3ddf2285119a257859e618e7d13130592dc1f6e886e271f46d05000453bffd2.ParticipantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant_id"] = id
    }
    return ie3ddf2285119a257859e618e7d13130592dc1f6e886e271f46d05000453bffd2.NewParticipantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in communications
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
func (m *CallRequestBuilder) PlayPrompt()(*i08e66c5d7bd1abccfebd8fb93bd3824b29263d4a78a8e0619b6492ed8f3fad49.PlayPromptRequestBuilder) {
    return i08e66c5d7bd1abccfebd8fb93bd3824b29263d4a78a8e0619b6492ed8f3fad49.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Record()(*i0c63fde205100ca16ab7a423f47cbdbde7db96e5bfd78a3f404e33371a9a304c.RecordRequestBuilder) {
    return i0c63fde205100ca16ab7a423f47cbdbde7db96e5bfd78a3f404e33371a9a304c.NewRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) RecordResponse()(*i18782668e274df0b14fede1b2d0b04b9dfad365a601b91cb1ae1236333912f77.RecordResponseRequestBuilder) {
    return i18782668e274df0b14fede1b2d0b04b9dfad365a601b91cb1ae1236333912f77.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Redirect()(*i51c2a09e4f41fd4c07a21a333066d7d9b5683f48919e40d555e43eeedc2d2b0f.RedirectRequestBuilder) {
    return i51c2a09e4f41fd4c07a21a333066d7d9b5683f48919e40d555e43eeedc2d2b0f.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Reject()(*idaf0ca56e2f68098afebd431de7c548ac38196c36423e844dd58942eb2a4ee85.RejectRequestBuilder) {
    return idaf0ca56e2f68098afebd431de7c548ac38196c36423e844dd58942eb2a4ee85.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) SubscribeToTone()(*i59f1af4a71937287f9a647acee47040de638eb6d16981d2dc3f1c22d63d3d210.SubscribeToToneRequestBuilder) {
    return i59f1af4a71937287f9a647acee47040de638eb6d16981d2dc3f1c22d63d3d210.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Transfer()(*i4ffb8669c93d0b11c0238e80369824011ea5e33bda9395560f9abe880eb8acc0.TransferRequestBuilder) {
    return i4ffb8669c93d0b11c0238e80369824011ea5e33bda9395560f9abe880eb8acc0.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) Unmute()(*i4cae3b1a69ba2254d74646c3ba02c22835c51e378675f8f86fa207a4bf13fbb9.UnmuteRequestBuilder) {
    return i4cae3b1a69ba2254d74646c3ba02c22835c51e378675f8f86fa207a4bf13fbb9.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallRequestBuilder) UpdateRecordingStatus()(*i2e0863e08fc93d954484ce570f9308574c3c610b9cbf1e2c9d07876c640463ac.UpdateRecordingStatusRequestBuilder) {
    return i2e0863e08fc93d954484ce570f9308574c3c610b9cbf1e2c9d07876c640463ac.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
