package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
type ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters the to [term] of the relation. The term to which the relationship is defined.
type ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters
}
// NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal instantiates a new ToTermRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) {
    m := &ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}/toTerm{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder instantiates a new ToTermRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the to [term] of the relation. The term to which the relationship is defined.
func (m *ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
// ToGetRequestInformation the to [term] of the relation. The term to which the relationship is defined.
func (m *ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}