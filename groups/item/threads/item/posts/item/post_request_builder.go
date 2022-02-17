package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i156960fb543e0ce4bf625d42526a73c4f7a0ec9672c001acf9d49cf44b8c2ff5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/mentions"
    i248c2e72a3592bf28dfbc2f71844da6e147b655d48ef2eb3f3ea62f42d90d966 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties"
    i43d5bf7d266db50b50e04a682edc7fe514bafa2465355a94a007e6320d99d617 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto"
    i752e70d2b7fc2704d4eef3b8d3f19794b3d3f44258045bb54edd2e65892d05bd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/reply"
    i89db43ff66f447209002f5665fd5f78ed25f827659cff39d193e2f33ae220e2e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties"
    i9df977ca780dc0d140675ce87069d982d3151d901c7fd2a7e5823b3934fa25f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/attachments"
    ic609b04e6ef2ba4b655b061ece719f39685820230e434a7b13efd1a16d666761 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/extensions"
    iea164510b0c521876f9bf52ef4f4f7bcfb479964d6a7a35ad5ca843d1b5f2808 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/forward"
    i0f1e494ef605c5c995db7dbaa4086c416a19342429243811dc9e2320600ae90e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/extensions/item"
    i31e115848de504a674c4e13d6beadc9b2d83e879a5fdc82f93f5ddfb706ac343 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/attachments/item"
    i4120f55900fac33f51a526e959eeabc979bb8d2ba02378a3af4bf2b6fb858a02 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties/item"
    i767b4645e14dda89ea8a901343958f824c9cc0eb861a5ea7515de33e1d921593 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/mentions/item"
    ie83abea890f828e4525d3a6f7a17ab8967c2b4d3e746928994f05c75ce15a262 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties/item"
)

// PostRequestBuilder builds and executes requests for operations under \groups\{group-id}\threads\{conversationThread-id}\posts\{post-id}
type PostRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PostRequestBuilderDeleteOptions options for Delete
type PostRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PostRequestBuilderGetOptions options for Get
type PostRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PostRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PostRequestBuilderGetQueryParameters read-only. Nullable.
type PostRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PostRequestBuilderPatchOptions options for Patch
type PostRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Post;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PostRequestBuilder) Attachments()(*i9df977ca780dc0d140675ce87069d982d3151d901c7fd2a7e5823b3934fa25f4.AttachmentsRequestBuilder) {
    return i9df977ca780dc0d140675ce87069d982d3151d901c7fd2a7e5823b3934fa25f4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.attachments.item collection
func (m *PostRequestBuilder) AttachmentsById(id string)(*i31e115848de504a674c4e13d6beadc9b2d83e879a5fdc82f93f5ddfb706ac343.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i31e115848de504a674c4e13d6beadc9b2d83e879a5fdc82f93f5ddfb706ac343.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostRequestBuilderInternal instantiates a new PostRequestBuilder and sets the default values.
func NewPostRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostRequestBuilder) {
    m := &PostRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPostRequestBuilder instantiates a new PostRequestBuilder and sets the default values.
func NewPostRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable.
func (m *PostRequestBuilder) CreateDeleteRequestInformation(options *PostRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable.
func (m *PostRequestBuilder) CreateGetRequestInformation(options *PostRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only. Nullable.
func (m *PostRequestBuilder) CreatePatchRequestInformation(options *PostRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only. Nullable.
func (m *PostRequestBuilder) Delete(options *PostRequestBuilderDeleteOptions)(error) {
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
func (m *PostRequestBuilder) Extensions()(*ic609b04e6ef2ba4b655b061ece719f39685820230e434a7b13efd1a16d666761.ExtensionsRequestBuilder) {
    return ic609b04e6ef2ba4b655b061ece719f39685820230e434a7b13efd1a16d666761.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.extensions.item collection
func (m *PostRequestBuilder) ExtensionsById(id string)(*i0f1e494ef605c5c995db7dbaa4086c416a19342429243811dc9e2320600ae90e.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i0f1e494ef605c5c995db7dbaa4086c416a19342429243811dc9e2320600ae90e.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostRequestBuilder) Forward()(*iea164510b0c521876f9bf52ef4f4f7bcfb479964d6a7a35ad5ca843d1b5f2808.ForwardRequestBuilder) {
    return iea164510b0c521876f9bf52ef4f4f7bcfb479964d6a7a35ad5ca843d1b5f2808.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable.
func (m *PostRequestBuilder) Get(options *PostRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Post, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPost() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Post), nil
}
func (m *PostRequestBuilder) InReplyTo()(*i43d5bf7d266db50b50e04a682edc7fe514bafa2465355a94a007e6320d99d617.InReplyToRequestBuilder) {
    return i43d5bf7d266db50b50e04a682edc7fe514bafa2465355a94a007e6320d99d617.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) Mentions()(*i156960fb543e0ce4bf625d42526a73c4f7a0ec9672c001acf9d49cf44b8c2ff5.MentionsRequestBuilder) {
    return i156960fb543e0ce4bf625d42526a73c4f7a0ec9672c001acf9d49cf44b8c2ff5.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.mentions.item collection
func (m *PostRequestBuilder) MentionsById(id string)(*i767b4645e14dda89ea8a901343958f824c9cc0eb861a5ea7515de33e1d921593.MentionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i767b4645e14dda89ea8a901343958f824c9cc0eb861a5ea7515de33e1d921593.NewMentionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostRequestBuilder) MultiValueExtendedProperties()(*i89db43ff66f447209002f5665fd5f78ed25f827659cff39d193e2f33ae220e2e.MultiValueExtendedPropertiesRequestBuilder) {
    return i89db43ff66f447209002f5665fd5f78ed25f827659cff39d193e2f33ae220e2e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4120f55900fac33f51a526e959eeabc979bb8d2ba02378a3af4bf2b6fb858a02.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4120f55900fac33f51a526e959eeabc979bb8d2ba02378a3af4bf2b6fb858a02.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch read-only. Nullable.
func (m *PostRequestBuilder) Patch(options *PostRequestBuilderPatchOptions)(error) {
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
func (m *PostRequestBuilder) Reply()(*i752e70d2b7fc2704d4eef3b8d3f19794b3d3f44258045bb54edd2e65892d05bd.ReplyRequestBuilder) {
    return i752e70d2b7fc2704d4eef3b8d3f19794b3d3f44258045bb54edd2e65892d05bd.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) SingleValueExtendedProperties()(*i248c2e72a3592bf28dfbc2f71844da6e147b655d48ef2eb3f3ea62f42d90d966.SingleValueExtendedPropertiesRequestBuilder) {
    return i248c2e72a3592bf28dfbc2f71844da6e147b655d48ef2eb3f3ea62f42d90d966.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie83abea890f828e4525d3a6f7a17ab8967c2b4d3e746928994f05c75ce15a262.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ie83abea890f828e4525d3a6f7a17ab8967c2b4d3e746928994f05c75ce15a262.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
