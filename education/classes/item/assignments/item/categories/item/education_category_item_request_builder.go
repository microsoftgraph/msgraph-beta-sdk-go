package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0097df2d4a822db8beacbc7b0687532ca1e19dc24c8ce5066863f8683c2d67f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/categories/item/ref"
)

// EducationCategoryItemRequestBuilder builds and executes requests for operations under \education\classes\{educationClass-id}\assignments\{educationAssignment-id}\categories\{educationCategory-id}
type EducationCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEducationCategoryItemRequestBuilderInternal instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewEducationCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationCategoryItemRequestBuilder) {
    m := &EducationCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/categories/{educationCategory%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationCategoryItemRequestBuilder instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewEducationCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the Ref property
func (m *EducationCategoryItemRequestBuilder) Ref()(*i0097df2d4a822db8beacbc7b0687532ca1e19dc24c8ce5066863f8683c2d67f4.RefRequestBuilder) {
    return i0097df2d4a822db8beacbc7b0687532ca1e19dc24c8ce5066863f8683c2d67f4.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
