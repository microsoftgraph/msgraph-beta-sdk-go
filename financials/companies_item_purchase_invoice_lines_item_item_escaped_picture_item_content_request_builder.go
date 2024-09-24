package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder provides operations to manage the media for the financials entity.
type CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderInternal instantiates a new CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) {
    m := &CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoiceLines/{purchaseInvoiceLine%2Did}/item/picture/{picture%2Did}/content", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder instantiates a new CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete content for the navigation property picture in financials
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get content for the navigation property picture from financials
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put update content for the navigation property picture in financials
// returns a Pictureable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Pictureable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePictureFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Pictureable), nil
}
// ToDeleteRequestInformation delete content for the navigation property picture in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get content for the navigation property picture from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update content for the navigation property picture in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder when successful
func (m *CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder) {
    return NewCompaniesItemPurchaseInvoiceLinesItemItem_EscapedPictureItemContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
