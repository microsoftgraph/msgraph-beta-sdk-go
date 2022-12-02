package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder provides operations to count the resources in the collection.
type OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "text/plain"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the number of the resource
func (m *OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorGroupsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
