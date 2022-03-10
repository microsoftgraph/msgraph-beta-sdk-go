package inreplyto

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i349c66809ee5ffcd3487e1cb6aa7d01312016b32297c66e02811e6519d8a8765 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/forward"
    i7c88f1f031be9c8d7817e9ac34c6ca165e57a6a15dae4048080ea6d3b696d519 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i8021dc8cc8fa143d8c2d23f4df5c99bb0860539eba9975ce07394acf6a61b35c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/attachments"
    i8e7a729a1f87a2a3c2d64eddff0b01a17f98b024619490445d0f9049e1002f59 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    ib1720b7cd9ad83a3abddc61dacc0f8e2918e8cdfa059bba98727bc89bc9e3eb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/reply"
    ic7cd8f5d1b749a6f3179065b45fe9c9b448775ffabce5e24d80f330770759626 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/extensions"
    id4342d4ceb0237ffc99bcf5d622af1d89b6857a701b86a3a6c8e0cad08b32b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/mentions"
    i0208c190bae332c0d33a3314acc48f4169d0163b7f68eb7f40f5b625a79c84b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
    i035365952aea96ad38a7d11d674014afa77467d9ba7a5f857efa06ed5ce2368e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/extensions/item"
    i073fc856bbf5ea79a432559129799046dbd9348990766d471e50d7b6f3e05cec "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/mentions/item"
    i75a7b92d2a188e8d9cc2e82addf05c5113eec371862830e85ef03fe4a56a8962 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
    ic9fbcaff777b55b4cebc2a586c9f536de98fdd200ab17ab910893de2ac5a6d62 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto/attachments/item"
)

// InReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type InReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InReplyToRequestBuilderDeleteOptions options for Delete
type InReplyToRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InReplyToRequestBuilderGetOptions options for Get
type InReplyToRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InReplyToRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InReplyToRequestBuilderGetQueryParameters read-only. Supports $expand.
type InReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InReplyToRequestBuilderPatchOptions options for Patch
type InReplyToRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Postable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *InReplyToRequestBuilder) Attachments()(*i8021dc8cc8fa143d8c2d23f4df5c99bb0860539eba9975ce07394acf6a61b35c.AttachmentsRequestBuilder) {
    return i8021dc8cc8fa143d8c2d23f4df5c99bb0860539eba9975ce07394acf6a61b35c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*ic9fbcaff777b55b4cebc2a586c9f536de98fdd200ab17ab910893de2ac5a6d62.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ic9fbcaff777b55b4cebc2a586c9f536de98fdd200ab17ab910893de2ac5a6d62.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/threads/{conversationThread_id}/posts/{post_id}/inReplyTo{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property inReplyTo for groups
func (m *InReplyToRequestBuilder) CreateDeleteRequestInformation(options *InReplyToRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Supports $expand.
func (m *InReplyToRequestBuilder) CreateGetRequestInformation(options *InReplyToRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property inReplyTo in groups
func (m *InReplyToRequestBuilder) CreatePatchRequestInformation(options *InReplyToRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property inReplyTo for groups
func (m *InReplyToRequestBuilder) Delete(options *InReplyToRequestBuilderDeleteOptions)(error) {
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
func (m *InReplyToRequestBuilder) Extensions()(*ic7cd8f5d1b749a6f3179065b45fe9c9b448775ffabce5e24d80f330770759626.ExtensionsRequestBuilder) {
    return ic7cd8f5d1b749a6f3179065b45fe9c9b448775ffabce5e24d80f330770759626.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*i035365952aea96ad38a7d11d674014afa77467d9ba7a5f857efa06ed5ce2368e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i035365952aea96ad38a7d11d674014afa77467d9ba7a5f857efa06ed5ce2368e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InReplyToRequestBuilder) Forward()(*i349c66809ee5ffcd3487e1cb6aa7d01312016b32297c66e02811e6519d8a8765.ForwardRequestBuilder) {
    return i349c66809ee5ffcd3487e1cb6aa7d01312016b32297c66e02811e6519d8a8765.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Supports $expand.
func (m *InReplyToRequestBuilder) Get(options *InReplyToRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePostFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Postable), nil
}
func (m *InReplyToRequestBuilder) Mentions()(*id4342d4ceb0237ffc99bcf5d622af1d89b6857a701b86a3a6c8e0cad08b32b4f.MentionsRequestBuilder) {
    return id4342d4ceb0237ffc99bcf5d622af1d89b6857a701b86a3a6c8e0cad08b32b4f.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.inReplyTo.mentions.item collection
func (m *InReplyToRequestBuilder) MentionsById(id string)(*i073fc856bbf5ea79a432559129799046dbd9348990766d471e50d7b6f3e05cec.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i073fc856bbf5ea79a432559129799046dbd9348990766d471e50d7b6f3e05cec.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*i7c88f1f031be9c8d7817e9ac34c6ca165e57a6a15dae4048080ea6d3b696d519.MultiValueExtendedPropertiesRequestBuilder) {
    return i7c88f1f031be9c8d7817e9ac34c6ca165e57a6a15dae4048080ea6d3b696d519.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i75a7b92d2a188e8d9cc2e82addf05c5113eec371862830e85ef03fe4a56a8962.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i75a7b92d2a188e8d9cc2e82addf05c5113eec371862830e85ef03fe4a56a8962.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property inReplyTo in groups
func (m *InReplyToRequestBuilder) Patch(options *InReplyToRequestBuilderPatchOptions)(error) {
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
func (m *InReplyToRequestBuilder) Reply()(*ib1720b7cd9ad83a3abddc61dacc0f8e2918e8cdfa059bba98727bc89bc9e3eb7.ReplyRequestBuilder) {
    return ib1720b7cd9ad83a3abddc61dacc0f8e2918e8cdfa059bba98727bc89bc9e3eb7.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*i8e7a729a1f87a2a3c2d64eddff0b01a17f98b024619490445d0f9049e1002f59.SingleValueExtendedPropertiesRequestBuilder) {
    return i8e7a729a1f87a2a3c2d64eddff0b01a17f98b024619490445d0f9049e1002f59.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0208c190bae332c0d33a3314acc48f4169d0163b7f68eb7f40f5b625a79c84b3.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i0208c190bae332c0d33a3314acc48f4169d0163b7f68eb7f40f5b625a79c84b3.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
