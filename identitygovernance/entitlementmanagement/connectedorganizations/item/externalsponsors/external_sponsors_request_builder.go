package externalsponsors

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c739a90bae49bb1f0e0b654af3860314123c667b24c636023d0960b10dd3b7c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/getuserownedobjects"
    i49612de40d9c84a64b500b2b165a45e2f0021cdf66632510c14f5180222d1354 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/validateproperties"
    i66ed94612eb80575e321bc56844adac51dd5d430875ab95fe35a0053af1d12cc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/ref"
    ia5d241124c155a82f1bc6c8974c36dc96711ea7783bcefab7e9d2a8a28d882df "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/count"
    if60b0c836a37d6e059e94d6d3b736c64c7ded70e5aa6f99931a1860e177cc89c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/getbyids"
)

// ExternalSponsorsRequestBuilder provides operations to manage the externalSponsors property of the microsoft.graph.connectedOrganization entity.
type ExternalSponsorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExternalSponsorsRequestBuilderGetQueryParameters get externalSponsors from identityGovernance
type ExternalSponsorsRequestBuilderGetQueryParameters struct {
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
// ExternalSponsorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalSponsorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExternalSponsorsRequestBuilderGetQueryParameters
}
// ExternalSponsorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalSponsorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExternalSponsorsRequestBuilderInternal instantiates a new ExternalSponsorsRequestBuilder and sets the default values.
func NewExternalSponsorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalSponsorsRequestBuilder) {
    m := &ExternalSponsorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalSponsorsRequestBuilder instantiates a new ExternalSponsorsRequestBuilder and sets the default values.
func NewExternalSponsorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalSponsorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalSponsorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ExternalSponsorsRequestBuilder) Count()(*ia5d241124c155a82f1bc6c8974c36dc96711ea7783bcefab7e9d2a8a28d882df.CountRequestBuilder) {
    return ia5d241124c155a82f1bc6c8974c36dc96711ea7783bcefab7e9d2a8a28d882df.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get externalSponsors from identityGovernance
func (m *ExternalSponsorsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get externalSponsors from identityGovernance
func (m *ExternalSponsorsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ExternalSponsorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to externalSponsors for identityGovernance
func (m *ExternalSponsorsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to externalSponsors for identityGovernance
func (m *ExternalSponsorsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *ExternalSponsorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get externalSponsors from identityGovernance
func (m *ExternalSponsorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExternalSponsorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// GetByIds the getByIds property
func (m *ExternalSponsorsRequestBuilder) GetByIds()(*if60b0c836a37d6e059e94d6d3b736c64c7ded70e5aa6f99931a1860e177cc89c.GetByIdsRequestBuilder) {
    return if60b0c836a37d6e059e94d6d3b736c64c7ded70e5aa6f99931a1860e177cc89c.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *ExternalSponsorsRequestBuilder) GetUserOwnedObjects()(*i0c739a90bae49bb1f0e0b654af3860314123c667b24c636023d0960b10dd3b7c.GetUserOwnedObjectsRequestBuilder) {
    return i0c739a90bae49bb1f0e0b654af3860314123c667b24c636023d0960b10dd3b7c.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to externalSponsors for identityGovernance
func (m *ExternalSponsorsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *ExternalSponsorsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Ref the Ref property
func (m *ExternalSponsorsRequestBuilder) Ref()(*i66ed94612eb80575e321bc56844adac51dd5d430875ab95fe35a0053af1d12cc.RefRequestBuilder) {
    return i66ed94612eb80575e321bc56844adac51dd5d430875ab95fe35a0053af1d12cc.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *ExternalSponsorsRequestBuilder) ValidateProperties()(*i49612de40d9c84a64b500b2b165a45e2f0021cdf66632510c14f5180222d1354.ValidatePropertiesRequestBuilder) {
    return i49612de40d9c84a64b500b2b165a45e2f0021cdf66632510c14f5180222d1354.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
