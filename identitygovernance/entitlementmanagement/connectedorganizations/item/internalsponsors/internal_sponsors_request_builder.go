package internalsponsors

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11bc5392c3a4232164d7ff0d40300c2bc9f854340d2ff86e6158e097a3da46bc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/ref"
    i1cda344c11a962065b879a971bf1f427333deb383b485a220fe89eaf9824d415 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/getbyids"
    i6234282d2060bd663e404a1e137e2cd1e63cc8e34e34c07a80a0194fc6de64c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/validateproperties"
    ia3ddcaee7edd6e7d20f7a8033a83e2a9f7b04b044d99d9614447820763215910 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/getuserownedobjects"
    iffcb865fe21325304d423c2cb93a3f8fd210a3c0f56f316c7c090898ceaac7dc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/count"
)

// InternalSponsorsRequestBuilder provides operations to manage the internalSponsors property of the microsoft.graph.connectedOrganization entity.
type InternalSponsorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InternalSponsorsRequestBuilderGetQueryParameters get internalSponsors from identityGovernance
type InternalSponsorsRequestBuilderGetQueryParameters struct {
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
// InternalSponsorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InternalSponsorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InternalSponsorsRequestBuilderGetQueryParameters
}
// InternalSponsorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InternalSponsorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInternalSponsorsRequestBuilderInternal instantiates a new InternalSponsorsRequestBuilder and sets the default values.
func NewInternalSponsorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InternalSponsorsRequestBuilder) {
    m := &InternalSponsorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInternalSponsorsRequestBuilder instantiates a new InternalSponsorsRequestBuilder and sets the default values.
func NewInternalSponsorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InternalSponsorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInternalSponsorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *InternalSponsorsRequestBuilder) Count()(*iffcb865fe21325304d423c2cb93a3f8fd210a3c0f56f316c7c090898ceaac7dc.CountRequestBuilder) {
    return iffcb865fe21325304d423c2cb93a3f8fd210a3c0f56f316c7c090898ceaac7dc.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get internalSponsors from identityGovernance
func (m *InternalSponsorsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get internalSponsors from identityGovernance
func (m *InternalSponsorsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *InternalSponsorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to internalSponsors for identityGovernance
func (m *InternalSponsorsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to internalSponsors for identityGovernance
func (m *InternalSponsorsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *InternalSponsorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get internalSponsors from identityGovernance
func (m *InternalSponsorsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetByIds the getByIds property
func (m *InternalSponsorsRequestBuilder) GetByIds()(*i1cda344c11a962065b879a971bf1f427333deb383b485a220fe89eaf9824d415.GetByIdsRequestBuilder) {
    return i1cda344c11a962065b879a971bf1f427333deb383b485a220fe89eaf9824d415.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *InternalSponsorsRequestBuilder) GetUserOwnedObjects()(*ia3ddcaee7edd6e7d20f7a8033a83e2a9f7b04b044d99d9614447820763215910.GetUserOwnedObjectsRequestBuilder) {
    return ia3ddcaee7edd6e7d20f7a8033a83e2a9f7b04b044d99d9614447820763215910.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get internalSponsors from identityGovernance
func (m *InternalSponsorsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *InternalSponsorsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Post create new navigation property to internalSponsors for identityGovernance
func (m *InternalSponsorsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to internalSponsors for identityGovernance
func (m *InternalSponsorsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *InternalSponsorsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Ref the ref property
func (m *InternalSponsorsRequestBuilder) Ref()(*i11bc5392c3a4232164d7ff0d40300c2bc9f854340d2ff86e6158e097a3da46bc.RefRequestBuilder) {
    return i11bc5392c3a4232164d7ff0d40300c2bc9f854340d2ff86e6158e097a3da46bc.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *InternalSponsorsRequestBuilder) ValidateProperties()(*i6234282d2060bd663e404a1e137e2cd1e63cc8e34e34c07a80a0194fc6de64c7.ValidatePropertiesRequestBuilder) {
    return i6234282d2060bd663e404a1e137e2cd1e63cc8e34e34c07a80a0194fc6de64c7.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
