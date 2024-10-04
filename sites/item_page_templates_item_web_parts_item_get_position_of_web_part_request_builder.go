package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
type ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderInternal instantiates a new ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder) {
    m := &ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pageTemplates/{pageTemplate%2Did}/webParts/{webPart%2Did}/getPositionOfWebPart", pathParameters),
    }
    return m
}
// NewItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder instantiates a new ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPositionOfWebPart
// returns a WebPartPositionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartPositionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWebPartPositionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartPositionable), nil
}
// ToPostRequestInformation invoke action getPositionOfWebPart
// returns a *RequestInformation when successful
func (m *ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder when successful
func (m *ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder) WithUrl(rawUrl string)(*ItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder) {
    return NewItemPageTemplatesItemWebPartsItemGetPositionOfWebPartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
