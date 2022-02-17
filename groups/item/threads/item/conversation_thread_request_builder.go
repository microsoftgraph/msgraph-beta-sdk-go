package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9113620de43a14942e3e5afbc6172c103505bf5149936f5c3293089e94efd9f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/reply"
    ic469e64dd7cbcacacb9f4b3c0fed127ca1a45f6233013122693848cc495dbad3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts"
    i4e2619f0471e4cce8101739f21c83490aef6c29642b652a22192dbdd5edd8b10 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item"
)

// ConversationThreadRequestBuilder builds and executes requests for operations under \groups\{group-id}\threads\{conversationThread-id}
type ConversationThreadRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ConversationThreadRequestBuilderDeleteOptions options for Delete
type ConversationThreadRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConversationThreadRequestBuilderGetOptions options for Get
type ConversationThreadRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConversationThreadRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConversationThreadRequestBuilderGetQueryParameters the group's conversation threads. Nullable.
type ConversationThreadRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// ConversationThreadRequestBuilderPatchOptions options for Patch
type ConversationThreadRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConversationThread;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewConversationThreadRequestBuilderInternal instantiates a new ConversationThreadRequestBuilder and sets the default values.
func NewConversationThreadRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConversationThreadRequestBuilder) {
    m := &ConversationThreadRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/threads/{conversationThread_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConversationThreadRequestBuilder instantiates a new ConversationThreadRequestBuilder and sets the default values.
func NewConversationThreadRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConversationThreadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConversationThreadRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the group's conversation threads. Nullable.
func (m *ConversationThreadRequestBuilder) CreateDeleteRequestInformation(options *ConversationThreadRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the group's conversation threads. Nullable.
func (m *ConversationThreadRequestBuilder) CreateGetRequestInformation(options *ConversationThreadRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the group's conversation threads. Nullable.
func (m *ConversationThreadRequestBuilder) CreatePatchRequestInformation(options *ConversationThreadRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the group's conversation threads. Nullable.
func (m *ConversationThreadRequestBuilder) Delete(options *ConversationThreadRequestBuilderDeleteOptions)(error) {
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
// Get the group's conversation threads. Nullable.
func (m *ConversationThreadRequestBuilder) Get(options *ConversationThreadRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConversationThread, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewConversationThread() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConversationThread), nil
}
// Patch the group's conversation threads. Nullable.
func (m *ConversationThreadRequestBuilder) Patch(options *ConversationThreadRequestBuilderPatchOptions)(error) {
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
func (m *ConversationThreadRequestBuilder) Posts()(*ic469e64dd7cbcacacb9f4b3c0fed127ca1a45f6233013122693848cc495dbad3.PostsRequestBuilder) {
    return ic469e64dd7cbcacacb9f4b3c0fed127ca1a45f6233013122693848cc495dbad3.NewPostsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PostsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item collection
func (m *ConversationThreadRequestBuilder) PostsById(id string)(*i4e2619f0471e4cce8101739f21c83490aef6c29642b652a22192dbdd5edd8b10.PostRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["post_id"] = id
    }
    return i4e2619f0471e4cce8101739f21c83490aef6c29642b652a22192dbdd5edd8b10.NewPostRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ConversationThreadRequestBuilder) Reply()(*i9113620de43a14942e3e5afbc6172c103505bf5149936f5c3293089e94efd9f5.ReplyRequestBuilder) {
    return i9113620de43a14942e3e5afbc6172c103505bf5149936f5c3293089e94efd9f5.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
