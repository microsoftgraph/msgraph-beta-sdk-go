package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBundlesDriveItemItemRequestBuilder provides operations to manage the bundles property of the microsoft.graph.drive entity.
type ItemBundlesDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBundlesDriveItemItemRequestBuilderGetQueryParameters collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
type ItemBundlesDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemBundlesDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBundlesDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemBundlesDriveItemItemRequestBuilderGetQueryParameters
}
// NewItemBundlesDriveItemItemRequestBuilderInternal instantiates a new ItemBundlesDriveItemItemRequestBuilder and sets the default values.
func NewItemBundlesDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBundlesDriveItemItemRequestBuilder) {
    m := &ItemBundlesDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/bundles/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemBundlesDriveItemItemRequestBuilder instantiates a new ItemBundlesDriveItemItemRequestBuilder and sets the default values.
func NewItemBundlesDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBundlesDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBundlesDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the drive entity.
// returns a *ItemBundlesItemContentRequestBuilder when successful
func (m *ItemBundlesDriveItemItemRequestBuilder) Content()(*ItemBundlesItemContentRequestBuilder) {
    return NewItemBundlesItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBundlesDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBundlesDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ToGetRequestInformation collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
// returns a *RequestInformation when successful
func (m *ItemBundlesDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBundlesDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemBundlesDriveItemItemRequestBuilder when successful
func (m *ItemBundlesDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*ItemBundlesDriveItemItemRequestBuilder) {
    return NewItemBundlesDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
