package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder provides operations to manage the activities property of the microsoft.graph.driveItem entity.
type DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetQueryParameters the list of recent activities that took place on this item.
type DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetQueryParameters
}
// NewDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderInternal instantiates a new ItemActivityOLDItemRequestBuilder and sets the default values.
func NewDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) {
    m := &DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/activities/{itemActivityOLD%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder instantiates a new ItemActivityOLDItemRequestBuilder and sets the default values.
func NewDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the list of recent activities that took place on this item.
func (m *DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the list of recent activities that took place on this item.
func (m *DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityOLDable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemActivityOLDFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityOLDable), nil
}
