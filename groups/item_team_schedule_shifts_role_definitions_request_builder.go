package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder provides operations to manage the shiftsRoleDefinitions property of the microsoft.graph.schedule entity.
type ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetQueryParameters the definitions of the roles in the schedule.
type ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetQueryParameters
}
// ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByShiftsRoleDefinitionId provides operations to manage the shiftsRoleDefinitions property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) ByShiftsRoleDefinitionId(shiftsRoleDefinitionId string)(*ItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if shiftsRoleDefinitionId != "" {
        urlTplParams["shiftsRoleDefinition%2Did"] = shiftsRoleDefinitionId
    }
    return NewItemTeamScheduleShiftsRoleDefinitionsShiftsRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilderInternal instantiates a new ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder and sets the default values.
func NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) {
    m := &ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/schedule/shiftsRoleDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilder instantiates a new ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder and sets the default values.
func NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamScheduleShiftsRoleDefinitionsCountRequestBuilder when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) Count()(*ItemTeamScheduleShiftsRoleDefinitionsCountRequestBuilder) {
    return NewItemTeamScheduleShiftsRoleDefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the definitions of the roles in the schedule.
// returns a ShiftsRoleDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateShiftsRoleDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionCollectionResponseable), nil
}
// Post create new navigation property to shiftsRoleDefinitions for groups
// returns a ShiftsRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateShiftsRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable), nil
}
// ToGetRequestInformation the definitions of the roles in the schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to shiftsRoleDefinitions for groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftsRoleDefinitionable, requestConfiguration *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/groups/{group%2Did}/team/schedule/shiftsRoleDefinitions", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder when successful
func (m *ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamScheduleShiftsRoleDefinitionsRequestBuilder) {
    return NewItemTeamScheduleShiftsRoleDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
