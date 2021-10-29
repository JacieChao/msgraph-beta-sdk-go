package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/operations"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/schema"
    ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups"
    ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/items"
    i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups/item"
    i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/operations/item"
    ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/items/item"
)

type ExternalConnectionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ExternalConnectionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewExternalConnectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionRequestBuilder) {
    m := &ExternalConnectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/connections/{externalConnection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewExternalConnectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExternalConnectionRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ExternalConnectionRequestBuilder) CreateGetRequestInformation(q func (value *ExternalConnectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ExternalConnectionRequestBuilderGetQueryParameters)
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
func (m *ExternalConnectionRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ExternalConnectionRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ExternalConnectionRequestBuilder) Get(q func (value *ExternalConnectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExternalConnection() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection), nil
}
func (m *ExternalConnectionRequestBuilder) Groups()(*ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0.GroupsRequestBuilder) {
    return ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) GroupsById(id string)(*i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7.ExternalGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup_id"] = id
    }
    return i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7.NewExternalGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Items()(*ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df.ItemsRequestBuilder) {
    return ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) ItemsById(id string)(*ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70.ExternalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem_id"] = id
    }
    return ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70.NewExternalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Operations()(*i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1.OperationsRequestBuilder) {
    return i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) OperationsById(id string)(*i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd.ConnectionOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation_id"] = id
    }
    return i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd.NewConnectionOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ExternalConnectionRequestBuilder) Schema()(*ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd.SchemaRequestBuilder) {
    return ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}