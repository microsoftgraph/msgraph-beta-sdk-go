package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InboundSharedUserProfilesItemExportPersonalDataRequestBuilder provides operations to call the exportPersonalData method.
type InboundSharedUserProfilesItemExportPersonalDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InboundSharedUserProfilesItemExportPersonalDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundSharedUserProfilesItemExportPersonalDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilderInternal instantiates a new ExportPersonalDataRequestBuilder and sets the default values.
func NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) {
    m := &InboundSharedUserProfilesItemExportPersonalDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/inboundSharedUserProfiles/{inboundSharedUserProfile%2DuserId}/exportPersonalData", pathParameters),
    }
    return m
}
// NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilder instantiates a new ExportPersonalDataRequestBuilder and sets the default values.
func NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a request to export the personal data for an inboundSharedUserProfile.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/inboundshareduserprofile-exportpersonaldata?view=graph-rest-1.0
func (m *InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) Post(ctx context.Context, body InboundSharedUserProfilesItemExportPersonalDataPostRequestBodyable, requestConfiguration *InboundSharedUserProfilesItemExportPersonalDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation create a request to export the personal data for an inboundSharedUserProfile.
func (m *InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) ToPostRequestInformation(ctx context.Context, body InboundSharedUserProfilesItemExportPersonalDataPostRequestBodyable, requestConfiguration *InboundSharedUserProfilesItemExportPersonalDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) WithUrl(rawUrl string)(*InboundSharedUserProfilesItemExportPersonalDataRequestBuilder) {
    return NewInboundSharedUserProfilesItemExportPersonalDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
