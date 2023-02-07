package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder provides operations to call the getShiftWorkCloudPcAccessState method.
type ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderInternal instantiates a new MicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder) {
    m := &ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/microsoft.graph.getShiftWorkCloudPcAccessState()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder instantiates a new MicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getShiftWorkCloudPcAccessState
func (m *ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftWorkCloudPcAccessState, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendEnum(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseShiftWorkCloudPcAccessState, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftWorkCloudPcAccessState), nil
}
// ToGetRequestInformation invoke function getShiftWorkCloudPcAccessState
func (m *ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
