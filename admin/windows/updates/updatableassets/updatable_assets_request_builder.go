package updatableassets

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/enrollassetsbyid"
    i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/unenrollassets"
    i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/enrollassets"
    i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/unenrollassetsbyid"
    ibf8eb5ac9b044029fa9545b3e41a68440c40e3423a43368a9196b474051989b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/count"
)

// UpdatableAssetsRequestBuilder provides operations to manage the updatableAssets property of the microsoft.graph.windowsUpdates.updates entity.
type UpdatableAssetsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UpdatableAssetsRequestBuilderGetQueryParameters assets registered with the deployment service that can receive updates. Read-only.
type UpdatableAssetsRequestBuilderGetQueryParameters struct {
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
// UpdatableAssetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UpdatableAssetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UpdatableAssetsRequestBuilderGetQueryParameters
}
// UpdatableAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UpdatableAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUpdatableAssetsRequestBuilderInternal instantiates a new UpdatableAssetsRequestBuilder and sets the default values.
func NewUpdatableAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdatableAssetsRequestBuilder) {
    m := &UpdatableAssetsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/updatableAssets{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatableAssetsRequestBuilder instantiates a new UpdatableAssetsRequestBuilder and sets the default values.
func NewUpdatableAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdatableAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatableAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *UpdatableAssetsRequestBuilder) Count()(*ibf8eb5ac9b044029fa9545b3e41a68440c40e3423a43368a9196b474051989b5.CountRequestBuilder) {
    return ibf8eb5ac9b044029fa9545b3e41a68440c40e3423a43368a9196b474051989b5.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformationWithRequestConfiguration assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UpdatableAssetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to updatableAssets for admin
func (m *UpdatableAssetsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to updatableAssets for admin
func (m *UpdatableAssetsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *UpdatableAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EnrollAssets the enrollAssets property
func (m *UpdatableAssetsRequestBuilder) EnrollAssets()(*i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e.EnrollAssetsRequestBuilder) {
    return i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e.NewEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollAssetsById the enrollAssetsById property
func (m *UpdatableAssetsRequestBuilder) EnrollAssetsById()(*i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7.EnrollAssetsByIdRequestBuilder) {
    return i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7.NewEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) GetWithResponseHandler(requestConfiguration *UpdatableAssetsRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) GetWithResponseHandler(requestConfiguration *UpdatableAssetsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable), nil
}
// PostWithResponseHandler create new navigation property to updatableAssets for admin
func (m *UpdatableAssetsRequestBuilder) PostWithResponseHandler(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *UpdatableAssetsRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler create new navigation property to updatableAssets for admin
func (m *UpdatableAssetsRequestBuilder) PostWithResponseHandler(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *UpdatableAssetsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// UnenrollAssets the unenrollAssets property
func (m *UpdatableAssetsRequestBuilder) UnenrollAssets()(*i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746.UnenrollAssetsRequestBuilder) {
    return i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746.NewUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnenrollAssetsById the unenrollAssetsById property
func (m *UpdatableAssetsRequestBuilder) UnenrollAssetsById()(*i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796.UnenrollAssetsByIdRequestBuilder) {
    return i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796.NewUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
