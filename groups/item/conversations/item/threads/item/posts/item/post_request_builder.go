package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i21b79d07c761682bb673240e4eb8319de00acb9fd95a4d62fc1438cdd1fe360f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto"
    i2350ee4e772f47bc96204e99012446a5ba284b89945297f51592795ca9939bcf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/mentions"
    i42c55eaf6d078713fbbf1cd7fc803f162bdd71d984883f73407f543b856e9052 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/forward"
    i48af31cfc155d6f278f3210ab213fd91cfb71d72f366254965a778c7a6645514 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties"
    i50b4455e2e8814d7b588669dc16eb6c4f3bf3e20f9e6968c681e1d8788e305a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/extensions"
    i5134b09571b88f3636966a5bd43caa67e98e4f65492bf4fc6a10b66b4c5e832e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/reply"
    i617649b82ab88335163983d60fb0153265183ca3a907aeea9c6b7f63f9926ad1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/attachments"
    icec1e01a6e690d90b8c752452fa1e7627a0b210a79ab2d52c656b39a31112cd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties"
    i02a31a2229ed9f6a9aec46303f3897896f086cb210422444f2ecd09c6a9e7a90 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/mentions/item"
    i04defb5caad166768a587cdc435f474c59f112ad68c86aeafb64d374def61578 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/extensions/item"
    i1dc3962d12cbe8ce1c60143acea51d6c5eb489bde18566f4c5563393f6915495 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/attachments/item"
    i2a79bd05857684e1a0bb06e60c1a996c3806ae20bac339617a463b711ddcd46e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties/item"
    id4c8d1b625a33def3fad99066f1b541c48ec7d4d641776001d150917f5ce13d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties/item"
)

// PostRequestBuilder builds and executes requests for operations under \groups\{group-id}\conversations\{conversation-id}\threads\{conversationThread-id}\posts\{post-id}
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
func (m *PostRequestBuilder) Attachments()(*i617649b82ab88335163983d60fb0153265183ca3a907aeea9c6b7f63f9926ad1.AttachmentsRequestBuilder) {
    return i617649b82ab88335163983d60fb0153265183ca3a907aeea9c6b7f63f9926ad1.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.attachments.item collection
func (m *PostRequestBuilder) AttachmentsById(id string)(*i1dc3962d12cbe8ce1c60143acea51d6c5eb489bde18566f4c5563393f6915495.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i1dc3962d12cbe8ce1c60143acea51d6c5eb489bde18566f4c5563393f6915495.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostRequestBuilderInternal instantiates a new PostRequestBuilder and sets the default values.
func NewPostRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostRequestBuilder) {
    m := &PostRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/conversations/{conversation_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
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
func (m *PostRequestBuilder) Extensions()(*i50b4455e2e8814d7b588669dc16eb6c4f3bf3e20f9e6968c681e1d8788e305a4.ExtensionsRequestBuilder) {
    return i50b4455e2e8814d7b588669dc16eb6c4f3bf3e20f9e6968c681e1d8788e305a4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.extensions.item collection
func (m *PostRequestBuilder) ExtensionsById(id string)(*i04defb5caad166768a587cdc435f474c59f112ad68c86aeafb64d374def61578.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i04defb5caad166768a587cdc435f474c59f112ad68c86aeafb64d374def61578.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostRequestBuilder) Forward()(*i42c55eaf6d078713fbbf1cd7fc803f162bdd71d984883f73407f543b856e9052.ForwardRequestBuilder) {
    return i42c55eaf6d078713fbbf1cd7fc803f162bdd71d984883f73407f543b856e9052.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *PostRequestBuilder) InReplyTo()(*i21b79d07c761682bb673240e4eb8319de00acb9fd95a4d62fc1438cdd1fe360f.InReplyToRequestBuilder) {
    return i21b79d07c761682bb673240e4eb8319de00acb9fd95a4d62fc1438cdd1fe360f.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) Mentions()(*i2350ee4e772f47bc96204e99012446a5ba284b89945297f51592795ca9939bcf.MentionsRequestBuilder) {
    return i2350ee4e772f47bc96204e99012446a5ba284b89945297f51592795ca9939bcf.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.mentions.item collection
func (m *PostRequestBuilder) MentionsById(id string)(*i02a31a2229ed9f6a9aec46303f3897896f086cb210422444f2ecd09c6a9e7a90.MentionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i02a31a2229ed9f6a9aec46303f3897896f086cb210422444f2ecd09c6a9e7a90.NewMentionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostRequestBuilder) MultiValueExtendedProperties()(*i48af31cfc155d6f278f3210ab213fd91cfb71d72f366254965a778c7a6645514.MultiValueExtendedPropertiesRequestBuilder) {
    return i48af31cfc155d6f278f3210ab213fd91cfb71d72f366254965a778c7a6645514.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2a79bd05857684e1a0bb06e60c1a996c3806ae20bac339617a463b711ddcd46e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2a79bd05857684e1a0bb06e60c1a996c3806ae20bac339617a463b711ddcd46e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *PostRequestBuilder) Reply()(*i5134b09571b88f3636966a5bd43caa67e98e4f65492bf4fc6a10b66b4c5e832e.ReplyRequestBuilder) {
    return i5134b09571b88f3636966a5bd43caa67e98e4f65492bf4fc6a10b66b4c5e832e.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) SingleValueExtendedProperties()(*icec1e01a6e690d90b8c752452fa1e7627a0b210a79ab2d52c656b39a31112cd3.SingleValueExtendedPropertiesRequestBuilder) {
    return icec1e01a6e690d90b8c752452fa1e7627a0b210a79ab2d52c656b39a31112cd3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id4c8d1b625a33def3fad99066f1b541c48ec7d4d641776001d150917f5ce13d4.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id4c8d1b625a33def3fad99066f1b541c48ec7d4d641776001d150917f5ce13d4.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
