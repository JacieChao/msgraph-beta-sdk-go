package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder provides operations to call the getAllRetainedMessages method.
type ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetQueryParameters invoke function getAllRetainedMessages
type ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetQueryParameters
}
// NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderInternal instantiates a new GetAllRetainedMessagesRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) {
    m := &ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/getAllRetainedMessages(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder instantiates a new GetAllRetainedMessagesRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAllRetainedMessages
// Deprecated: This method is obsolete. Use GetAsGetAllRetainedMessagesGetResponse instead.
func (m *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetRequestConfiguration)(ItemTeamDefinitionChannelsGetAllRetainedMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionChannelsGetAllRetainedMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionChannelsGetAllRetainedMessagesResponseable), nil
}
// GetAsGetAllRetainedMessagesGetResponse invoke function getAllRetainedMessages
func (m *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) GetAsGetAllRetainedMessagesGetResponse(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetRequestConfiguration)(ItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponseable), nil
}
// ToGetRequestInformation invoke function getAllRetainedMessages
func (m *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) {
    return NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
