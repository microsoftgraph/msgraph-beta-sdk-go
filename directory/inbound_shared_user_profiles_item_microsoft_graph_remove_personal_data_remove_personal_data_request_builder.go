package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder provides operations to call the removePersonalData method.
type InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderInternal instantiates a new RemovePersonalDataRequestBuilder and sets the default values.
func NewInboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder) {
    m := &InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/inboundSharedUserProfiles/{inboundSharedUserProfile%2DuserId}/microsoft.graph.removePersonalData";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewInboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder instantiates a new RemovePersonalDataRequestBuilder and sets the default values.
func NewInboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removePersonalData
func (m *InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder) Post(ctx context.Context, requestConfiguration *InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action removePersonalData
func (m *InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *InboundSharedUserProfilesItemMicrosoftGraphRemovePersonalDataRemovePersonalDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
