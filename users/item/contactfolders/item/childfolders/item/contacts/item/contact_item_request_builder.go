package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2c8c2f63774fc0ba523f7e67cc447bd31db1c15aeb803e68d58f75c6e562eac1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/extensions"
    i7ce5b3690847a71aa4f8f5b2b320a455d7c06f6f6c5857267261f05587ad1add "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/photo"
    ibba8ab63ffb942d9e28c2cfbc61662e6c97a6302ea449032869a7ab6360947b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/singlevalueextendedproperties"
    ida4aa1f41f643c3c2533cf0b676f13e2be2927f3605c609fbf0ee141acb79e94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/multivalueextendedproperties"
    i2096c30251c32d3172ef5d51ec2184291c6cdf193fb06068fc2067655f410230 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/singlevalueextendedproperties/item"
    i88edbc5825d3d24d04b9c9a16e5e4766a8404a856ac57312dea4034bbabedc63 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/multivalueextendedproperties/item"
    ic1df9a41436e08354d28b930a52dc36f8ea1ed80be66b72634fd708f61a3f082 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item/extensions/item"
)

// ContactItemRequestBuilder provides operations to manage the contacts property of the microsoft.graph.contactFolder entity.
type ContactItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContactItemRequestBuilderGetQueryParameters the contacts in the folder. Navigation property. Read-only. Nullable.
type ContactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContactItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContactItemRequestBuilderGetQueryParameters
}
// ContactItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContactItemRequestBuilderInternal instantiates a new ContactItemRequestBuilder and sets the default values.
func NewContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactItemRequestBuilder) {
    m := &ContactItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/childFolders/{contactFolder%2Did1}/contacts/{contact%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactItemRequestBuilder instantiates a new ContactItemRequestBuilder and sets the default values.
func NewContactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property contacts for users
func (m *ContactItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ContactItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contacts in users
func (m *ContactItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, requestConfiguration *ContactItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property contacts for users
func (m *ContactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) Extensions()(*i2c8c2f63774fc0ba523f7e67cc447bd31db1c15aeb803e68d58f75c6e562eac1.ExtensionsRequestBuilder) {
    return i2c8c2f63774fc0ba523f7e67cc447bd31db1c15aeb803e68d58f75c6e562eac1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) ExtensionsById(id string)(*ic1df9a41436e08354d28b930a52dc36f8ea1ed80be66b72634fd708f61a3f082.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic1df9a41436e08354d28b930a52dc36f8ea1ed80be66b72634fd708f61a3f082.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) MultiValueExtendedProperties()(*ida4aa1f41f643c3c2533cf0b676f13e2be2927f3605c609fbf0ee141acb79e94.MultiValueExtendedPropertiesRequestBuilder) {
    return ida4aa1f41f643c3c2533cf0b676f13e2be2927f3605c609fbf0ee141acb79e94.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i88edbc5825d3d24d04b9c9a16e5e4766a8404a856ac57312dea4034bbabedc63.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i88edbc5825d3d24d04b9c9a16e5e4766a8404a856ac57312dea4034bbabedc63.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property contacts in users
func (m *ContactItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, requestConfiguration *ContactItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable), nil
}
// Photo provides operations to manage the photo property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) Photo()(*i7ce5b3690847a71aa4f8f5b2b320a455d7c06f6f6c5857267261f05587ad1add.PhotoRequestBuilder) {
    return i7ce5b3690847a71aa4f8f5b2b320a455d7c06f6f6c5857267261f05587ad1add.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) SingleValueExtendedProperties()(*ibba8ab63ffb942d9e28c2cfbc61662e6c97a6302ea449032869a7ab6360947b5.SingleValueExtendedPropertiesRequestBuilder) {
    return ibba8ab63ffb942d9e28c2cfbc61662e6c97a6302ea449032869a7ab6360947b5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.contact entity.
func (m *ContactItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2096c30251c32d3172ef5d51ec2184291c6cdf193fb06068fc2067655f410230.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i2096c30251c32d3172ef5d51ec2184291c6cdf193fb06068fc2067655f410230.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
