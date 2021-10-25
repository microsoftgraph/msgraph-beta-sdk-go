package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0eb0ef44081f6129ed82f84b95a9d8901c5c53db2ff78cf336da3540e51902a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/copytodefaultcontentlocation"
    i24fd2f2396a7ad9f1c120d6ef5eea2bb0f52a3b0ab9176b7b84abea77552d44d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes"
    i2c3f11bea353e2053360132c240fdb0421d0e64bd89c28e21a2a5f54d072902c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/associatewithhubsites"
    i7971ec5b6746bc19a515d87f0c804fe6b9801df230d1bd98528a5cc26b925b7e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/unpublish"
    i99732ebd53d03d959a23a85d7cef79aa48fc15faa7863489113f7954f63618a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base"
    iafaa9e8f8d30224a4b19bafd6f27a6391457da364f2742173947aa5f8e09182a "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/publish"
    ibc27ed7d114db34133b5c7801610abab919a9cc58319247f2378aacf6c52c697 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions"
    id9af0414991da05e0a4cf709357ff74c7cf2e09da2b0b24dd95a2a0359adf307 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columns"
    ifb49fb4626206326449bcf75747e5bc8277181f4515cf8f4cbc7f5aa84a04e4b "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks"
    ifdcf97bbb94f49d415a91428a9d320151872eecd6c1f08d3669c845444f4ec54 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/ispublished"
    i27cf64cc6580d9aab73576050bd5ff42d2a2ca98593e1793629372ac762dd1ec "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columns/item"
    i800508a07b27c3d6db01cd892d2b2ecaff16c5d270423c5b1cc9e33635b0bf53 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks/item"
)

type ContentTypeRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ContentTypeRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i2c3f11bea353e2053360132c240fdb0421d0e64bd89c28e21a2a5f54d072902c.AssociateWithHubSitesRequestBuilder) {
    return i2c3f11bea353e2053360132c240fdb0421d0e64bd89c28e21a2a5f54d072902c.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i99732ebd53d03d959a23a85d7cef79aa48fc15faa7863489113f7954f63618a3.BaseRequestBuilder) {
    return i99732ebd53d03d959a23a85d7cef79aa48fc15faa7863489113f7954f63618a3.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*i24fd2f2396a7ad9f1c120d6ef5eea2bb0f52a3b0ab9176b7b84abea77552d44d.BaseTypesRequestBuilder) {
    return i24fd2f2396a7ad9f1c120d6ef5eea2bb0f52a3b0ab9176b7b84abea77552d44d.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*ifb49fb4626206326449bcf75747e5bc8277181f4515cf8f4cbc7f5aa84a04e4b.ColumnLinksRequestBuilder) {
    return ifb49fb4626206326449bcf75747e5bc8277181f4515cf8f4cbc7f5aa84a04e4b.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*i800508a07b27c3d6db01cd892d2b2ecaff16c5d270423c5b1cc9e33635b0bf53.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i800508a07b27c3d6db01cd892d2b2ecaff16c5d270423c5b1cc9e33635b0bf53.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*ibc27ed7d114db34133b5c7801610abab919a9cc58319247f2378aacf6c52c697.ColumnPositionsRequestBuilder) {
    return ibc27ed7d114db34133b5c7801610abab919a9cc58319247f2378aacf6c52c697.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*id9af0414991da05e0a4cf709357ff74c7cf2e09da2b0b24dd95a2a0359adf307.ColumnsRequestBuilder) {
    return id9af0414991da05e0a4cf709357ff74c7cf2e09da2b0b24dd95a2a0359adf307.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i27cf64cc6580d9aab73576050bd5ff42d2a2ca98593e1793629372ac762dd1ec.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i27cf64cc6580d9aab73576050bd5ff42d2a2ca98593e1793629372ac762dd1ec.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i0eb0ef44081f6129ed82f84b95a9d8901c5c53db2ff78cf336da3540e51902a2.CopyToDefaultContentLocationRequestBuilder) {
    return i0eb0ef44081f6129ed82f84b95a9d8901c5c53db2ff78cf336da3540e51902a2.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(q func (value *ContentTypeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ContentTypeRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ContentTypeRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Get(q func (value *ContentTypeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentType() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType), nil
}
func (m *ContentTypeRequestBuilder) IsPublished()(*ifdcf97bbb94f49d415a91428a9d320151872eecd6c1f08d3669c845444f4ec54.IsPublishedRequestBuilder) {
    return ifdcf97bbb94f49d415a91428a9d320151872eecd6c1f08d3669c845444f4ec54.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*iafaa9e8f8d30224a4b19bafd6f27a6391457da364f2742173947aa5f8e09182a.PublishRequestBuilder) {
    return iafaa9e8f8d30224a4b19bafd6f27a6391457da364f2742173947aa5f8e09182a.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i7971ec5b6746bc19a515d87f0c804fe6b9801df230d1bd98528a5cc26b925b7e.UnpublishRequestBuilder) {
    return i7971ec5b6746bc19a515d87f0c804fe6b9801df230d1bd98528a5cc26b925b7e.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
