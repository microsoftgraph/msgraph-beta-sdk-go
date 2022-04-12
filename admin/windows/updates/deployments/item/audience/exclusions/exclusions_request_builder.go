package exclusions

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09062f21d18bbad1807519431d48aac001991e078b5e4967de26f777128549c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/count"
    i5c9170b81c201b5105566ee9502169d272d67357d920434087742a266367025d "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/unenrollassets"
    i7a5ce3b3753fea9d5a62c39633c51d83d1cd2403b75643a52562b7a607be2f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/enrollassetsbyid"
    ic1f0bd7dc659fdc462a2c99f69ef901e44c18a5ef5ab6d0830bbfd7789e3667d "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/enrollassets"
    ie0c0fd2f52c8949442d7606cb6f684535237f75ae9a8505ddead286fdb640838 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/unenrollassetsbyid"
)

// ExclusionsRequestBuilder provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type ExclusionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExclusionsRequestBuilderGetOptions options for Get
type ExclusionsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExclusionsRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ExclusionsRequestBuilderGetQueryParameters specifies the assets to exclude from the audience.
type ExclusionsRequestBuilderGetQueryParameters struct {
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
// ExclusionsRequestBuilderPostOptions options for Post
type ExclusionsRequestBuilderPostOptions struct {
    // 
    Body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewExclusionsRequestBuilderInternal instantiates a new ExclusionsRequestBuilder and sets the default values.
func NewExclusionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExclusionsRequestBuilder) {
    m := &ExclusionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/exclusions{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExclusionsRequestBuilder instantiates a new ExclusionsRequestBuilder and sets the default values.
func NewExclusionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExclusionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExclusionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ExclusionsRequestBuilder) Count()(*i09062f21d18bbad1807519431d48aac001991e078b5e4967de26f777128549c4.CountRequestBuilder) {
    return i09062f21d18bbad1807519431d48aac001991e078b5e4967de26f777128549c4.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation specifies the assets to exclude from the audience.
func (m *ExclusionsRequestBuilder) CreateGetRequestInformation(options *ExclusionsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to exclusions for admin
func (m *ExclusionsRequestBuilder) CreatePostRequestInformation(options *ExclusionsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// EnrollAssets the enrollAssets property
func (m *ExclusionsRequestBuilder) EnrollAssets()(*ic1f0bd7dc659fdc462a2c99f69ef901e44c18a5ef5ab6d0830bbfd7789e3667d.EnrollAssetsRequestBuilder) {
    return ic1f0bd7dc659fdc462a2c99f69ef901e44c18a5ef5ab6d0830bbfd7789e3667d.NewEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollAssetsById the enrollAssetsById property
func (m *ExclusionsRequestBuilder) EnrollAssetsById()(*i7a5ce3b3753fea9d5a62c39633c51d83d1cd2403b75643a52562b7a607be2f57.EnrollAssetsByIdRequestBuilder) {
    return i7a5ce3b3753fea9d5a62c39633c51d83d1cd2403b75643a52562b7a607be2f57.NewEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get specifies the assets to exclude from the audience.
func (m *ExclusionsRequestBuilder) Get(options *ExclusionsRequestBuilderGetOptions)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetCollectionResponseable), nil
}
// Post create new navigation property to exclusions for admin
func (m *ExclusionsRequestBuilder) Post(options *ExclusionsRequestBuilderPostOptions)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// UnenrollAssets the unenrollAssets property
func (m *ExclusionsRequestBuilder) UnenrollAssets()(*i5c9170b81c201b5105566ee9502169d272d67357d920434087742a266367025d.UnenrollAssetsRequestBuilder) {
    return i5c9170b81c201b5105566ee9502169d272d67357d920434087742a266367025d.NewUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnenrollAssetsById the unenrollAssetsById property
func (m *ExclusionsRequestBuilder) UnenrollAssetsById()(*ie0c0fd2f52c8949442d7606cb6f684535237f75ae9a8505ddead286fdb640838.UnenrollAssetsByIdRequestBuilder) {
    return ie0c0fd2f52c8949442d7606cb6f684535237f75ae9a8505ddead286fdb640838.NewUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
