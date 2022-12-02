package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder provides operations to call the getCloudPcLaunchInfo method.
type MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal instantiates a new GetCloudPcLaunchInfoRequestBuilder and sets the default values.
func NewMeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    m := &MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/cloudPCs/{cloudPC%2Did}/microsoft.graph.getCloudPcLaunchInfo()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder instantiates a new GetCloudPcLaunchInfoRequestBuilder and sets the default values.
func NewMeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getCloudPcLaunchInfo
func (m *MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getCloudPcLaunchInfo
func (m *MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcLaunchInfoable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcLaunchInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcLaunchInfoable), nil
}
