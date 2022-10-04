package appliesto

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2faedfd6563232ad03511594aaa725b4aa98d0f39a974fb030b17a46e225f194 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/getbyids"
    i39c2c50639f8c9a9a5cfc3426436150d193d7a9ce9c827dc564a7499b69c48fd "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/validateproperties"
    i80cd825752a6c36726e34a08801539564e2e4948894a9fdc9512ef76fa5bd4b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/ref"
    i9385584faf3cda07d9406c6ada0bc2dc1333495a084ce94fb82d7c1086ded9f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/getuserownedobjects"
    ic54e399544e5bee9521a9e228a9cbd75d26d094bf5295b258e44df7102354067 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/count"
)

// AppliesToRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.featureRolloutPolicy entity.
type AppliesToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppliesToRequestBuilderGetQueryParameters nullable. Specifies a list of directoryObjects that feature is enabled for.
type AppliesToRequestBuilderGetQueryParameters struct {
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
// AppliesToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppliesToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppliesToRequestBuilderGetQueryParameters
}
// AppliesToRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppliesToRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppliesToRequestBuilderInternal instantiates a new AppliesToRequestBuilder and sets the default values.
func NewAppliesToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppliesToRequestBuilder) {
    m := &AppliesToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppliesToRequestBuilder instantiates a new AppliesToRequestBuilder and sets the default values.
func NewAppliesToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppliesToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppliesToRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *AppliesToRequestBuilder) Count()(*ic54e399544e5bee9521a9e228a9cbd75d26d094bf5295b258e44df7102354067.CountRequestBuilder) {
    return ic54e399544e5bee9521a9e228a9cbd75d26d094bf5295b258e44df7102354067.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *AppliesToRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AppliesToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to appliesTo for directory
func (m *AppliesToRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *AppliesToRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *AppliesToRequestBuilder) Get(ctx context.Context, requestConfiguration *AppliesToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
func (m *AppliesToRequestBuilder) GetByIds()(*i2faedfd6563232ad03511594aaa725b4aa98d0f39a974fb030b17a46e225f194.GetByIdsRequestBuilder) {
    return i2faedfd6563232ad03511594aaa725b4aa98d0f39a974fb030b17a46e225f194.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *AppliesToRequestBuilder) GetUserOwnedObjects()(*i9385584faf3cda07d9406c6ada0bc2dc1333495a084ce94fb82d7c1086ded9f7.GetUserOwnedObjectsRequestBuilder) {
    return i9385584faf3cda07d9406c6ada0bc2dc1333495a084ce94fb82d7c1086ded9f7.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to appliesTo for directory
func (m *AppliesToRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *AppliesToRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
func (m *AppliesToRequestBuilder) Ref()(*i80cd825752a6c36726e34a08801539564e2e4948894a9fdc9512ef76fa5bd4b5.RefRequestBuilder) {
    return i80cd825752a6c36726e34a08801539564e2e4948894a9fdc9512ef76fa5bd4b5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *AppliesToRequestBuilder) ValidateProperties()(*i39c2c50639f8c9a9a5cfc3426436150d193d7a9ce9c827dc564a7499b69c48fd.ValidatePropertiesRequestBuilder) {
    return i39c2c50639f8c9a9a5cfc3426436150d193d7a9ce9c827dc564a7499b69c48fd.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
