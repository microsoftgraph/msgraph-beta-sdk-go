package contacts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i078866701ebd014ef41908c3e576ecccdc54753b23645955da30d52b28cf8773 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/getuserownedobjects"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i263e3a49923c8b3d5d4677d4bffaef4520d3839dbb9e191c0ead10a25a4595b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/delta"
    i8fa8f6aa0a9b25f3a77255c9f362eca9df147df0488d2dc3b9c1a7a0c9afe5f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/getbyids"
    idc9921cc70b53e0e0915a24e8dbffa0e67b4c8c3233929cbec9702d7a8422ce4 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/count"
    ife4ab247d2377531ea11bb4c8a0ab1e2b8f2156beb597f561bf1f7d9bd436500 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/validateproperties"
)

// ContactsRequestBuilder provides operations to manage the collection of orgContact entities.
type ContactsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContactsRequestBuilderGetOptions options for Get
type ContactsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContactsRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ContactsRequestBuilderGetQueryParameters get entities from contacts
type ContactsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ContactsRequestBuilderPostOptions options for Post
type ContactsRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewContactsRequestBuilderInternal instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactsRequestBuilder) {
    m := &ContactsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactsRequestBuilder instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ContactsRequestBuilder) Count()(*idc9921cc70b53e0e0915a24e8dbffa0e67b4c8c3233929cbec9702d7a8422ce4.CountRequestBuilder) {
    return idc9921cc70b53e0e0915a24e8dbffa0e67b4c8c3233929cbec9702d7a8422ce4.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from contacts
func (m *ContactsRequestBuilder) CreateGetRequestInformation(options *ContactsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation add new entity to contacts
func (m *ContactsRequestBuilder) CreatePostRequestInformation(options *ContactsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *ContactsRequestBuilder) Delta()(*i263e3a49923c8b3d5d4677d4bffaef4520d3839dbb9e191c0ead10a25a4595b3.DeltaRequestBuilder) {
    return i263e3a49923c8b3d5d4677d4bffaef4520d3839dbb9e191c0ead10a25a4595b3.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from contacts
func (m *ContactsRequestBuilder) Get(options *ContactsRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrgContactCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactCollectionResponseable), nil
}
// GetByIds the getByIds property
func (m *ContactsRequestBuilder) GetByIds()(*i8fa8f6aa0a9b25f3a77255c9f362eca9df147df0488d2dc3b9c1a7a0c9afe5f4.GetByIdsRequestBuilder) {
    return i8fa8f6aa0a9b25f3a77255c9f362eca9df147df0488d2dc3b9c1a7a0c9afe5f4.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *ContactsRequestBuilder) GetUserOwnedObjects()(*i078866701ebd014ef41908c3e576ecccdc54753b23645955da30d52b28cf8773.GetUserOwnedObjectsRequestBuilder) {
    return i078866701ebd014ef41908c3e576ecccdc54753b23645955da30d52b28cf8773.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to contacts
func (m *ContactsRequestBuilder) Post(options *ContactsRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrgContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable), nil
}
// ValidateProperties the validateProperties property
func (m *ContactsRequestBuilder) ValidateProperties()(*ife4ab247d2377531ea11bb4c8a0ab1e2b8f2156beb597f561bf1f7d9bd436500.ValidatePropertiesRequestBuilder) {
    return ife4ab247d2377531ea11bb4c8a0ab1e2b8f2156beb597f561bf1f7d9bd436500.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
