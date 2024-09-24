package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
type ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal instantiates a new ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) {
    m := &ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pageTemplates/{pageTemplate%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts/{webPart%2Did}/getPositionOfWebPart", pathParameters),
    }
    return m
}
// NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder instantiates a new ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPositionOfWebPart
// returns a WebPartPositionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartPositionable, error) {
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
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder) {
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemGetPositionOfWebPartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
