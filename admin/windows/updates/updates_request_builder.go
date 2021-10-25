package updates

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/catalog"
    i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments"
    id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets"
    i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item"
    ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item"
)

type UpdatesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type UpdatesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *UpdatesRequestBuilder) Catalog()(*i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080.CatalogRequestBuilder) {
    return i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080.NewCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewUpdatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatesRequestBuilder) {
    m := &UpdatesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/admin/windows/updates{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewUpdatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UpdatesRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UpdatesRequestBuilder) CreateGetRequestInformation(q func (value *UpdatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UpdatesRequestBuilderGetQueryParameters)
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
func (m *UpdatesRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Updates, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UpdatesRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *UpdatesRequestBuilder) Deployments()(*i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669.DeploymentsRequestBuilder) {
    return i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669.NewDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatesRequestBuilder) DeploymentsById(id string)(*ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b.DeploymentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deployment_id"] = id
    }
    return ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b.NewDeploymentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UpdatesRequestBuilder) Get(q func (value *UpdatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Updates, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUpdates() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Updates), nil
}
func (m *UpdatesRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Updates, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *UpdatesRequestBuilder) UpdatableAssets()(*id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863.UpdatableAssetsRequestBuilder) {
    return id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863.NewUpdatableAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatesRequestBuilder) UpdatableAssetsById(id string)(*i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b.UpdatableAssetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset_id"] = id
    }
    return i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b.NewUpdatableAssetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
