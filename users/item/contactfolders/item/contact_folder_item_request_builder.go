package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties"
    ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/multivalueextendedproperties"
    id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders"
    if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts"
    i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/multivalueextendedproperties/item"
    i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties/item"
    ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item"
    icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item"
)

// ContactFolderItemRequestBuilder provides operations to manage the contactFolders property of the microsoft.graph.user entity.
type ContactFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContactFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContactFolderItemRequestBuilderGetQueryParameters the user's contacts folders. Read-only. Nullable.
type ContactFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContactFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContactFolderItemRequestBuilderGetQueryParameters
}
// ContactFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChildFolders provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) ChildFolders()(*id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7.ChildFoldersRequestBuilder) {
    return id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) ChildFoldersById(id string)(*ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39.ContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did1"] = id
    }
    return ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39.NewContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContactFolderItemRequestBuilderInternal instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    m := &ContactFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactFolderItemRequestBuilder instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) Contacts()(*if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd.ContactsRequestBuilder) {
    return if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById provides operations to manage the contacts property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) ContactsById(id string)(*icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b.ContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b.NewContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contactFolders for users
func (m *ContactFolderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ContactFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contactFolders in users
func (m *ContactFolderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, requestConfiguration *ContactFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contactFolders for users
func (m *ContactFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContactFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedProperties()(*ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e.MultiValueExtendedPropertiesRequestBuilder) {
    return ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property contactFolders in users
func (m *ContactFolderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, requestConfiguration *ContactFolderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContactFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactFolderable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedProperties()(*i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6.SingleValueExtendedPropertiesRequestBuilder) {
    return i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.contactFolder entity.
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
