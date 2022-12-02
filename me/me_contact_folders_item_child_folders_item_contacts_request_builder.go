package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeContactFoldersItemChildFoldersItemContactsRequestBuilder provides operations to manage the contacts property of the microsoft.graph.contactFolder entity.
type MeContactFoldersItemChildFoldersItemContactsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetQueryParameters get all the contacts in the signed-in user's mailbox (.../me/contacts), or from the specified contact folder.
type MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetQueryParameters
}
// MeContactFoldersItemChildFoldersItemContactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeContactFoldersItemChildFoldersItemContactsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeContactFoldersItemChildFoldersItemContactsRequestBuilderInternal instantiates a new ContactsRequestBuilder and sets the default values.
func NewMeContactFoldersItemChildFoldersItemContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeContactFoldersItemChildFoldersItemContactsRequestBuilder) {
    m := &MeContactFoldersItemChildFoldersItemContactsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder%2Did}/childFolders/{contactFolder%2Did1}/contacts{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeContactFoldersItemChildFoldersItemContactsRequestBuilder instantiates a new ContactsRequestBuilder and sets the default values.
func NewMeContactFoldersItemChildFoldersItemContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeContactFoldersItemChildFoldersItemContactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeContactFoldersItemChildFoldersItemContactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MeContactFoldersItemChildFoldersItemContactsRequestBuilder) Count()(*MeContactFoldersItemChildFoldersItemContactsCountRequestBuilder) {
    return NewMeContactFoldersItemChildFoldersItemContactsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get all the contacts in the signed-in user's mailbox (.../me/contacts), or from the specified contact folder.
func (m *MeContactFoldersItemChildFoldersItemContactsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add a contact to the root Contacts folder or to the `contacts` endpoint of another contact folder.
func (m *MeContactFoldersItemChildFoldersItemContactsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, requestConfiguration *MeContactFoldersItemChildFoldersItemContactsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *MeContactFoldersItemChildFoldersItemContactsRequestBuilder) Delta()(*MeContactFoldersItemChildFoldersItemContactsDeltaRequestBuilder) {
    return NewMeContactFoldersItemChildFoldersItemContactsDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get all the contacts in the signed-in user's mailbox (.../me/contacts), or from the specified contact folder.
func (m *MeContactFoldersItemChildFoldersItemContactsRequestBuilder) Get(ctx context.Context, requestConfiguration *MeContactFoldersItemChildFoldersItemContactsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContactCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContactCollectionResponseable), nil
}
// Post add a contact to the root Contacts folder or to the `contacts` endpoint of another contact folder.
func (m *MeContactFoldersItemChildFoldersItemContactsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, requestConfiguration *MeContactFoldersItemChildFoldersItemContactsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Contactable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
