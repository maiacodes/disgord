package disgord

// Code generated - This file has been automatically generated by generate/restbuilders/main.go - DO NOT EDIT.
// Warning: This file is overwritten by the "go generate" command
// This file holds all the basic RESTBuilder methods a builder is expected to.

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *guildAuditLogsBuilder) IgnoreCache() *guildAuditLogsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *guildAuditLogsBuilder) CancelOnRatelimit() *guildAuditLogsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *guildAuditLogsBuilder) URLParam(name string, v interface{}) *guildAuditLogsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *guildAuditLogsBuilder) Set(name string, v interface{}) *guildAuditLogsBuilder {
	b.r.body[name] = v
	return b
}

func (b *guildAuditLogsBuilder) SetUserID(userID Snowflake) *guildAuditLogsBuilder {
	b.r.addPrereq(userID.Empty(), "userID can not be 0")
	b.r.param("user_id", userID)
	return b
}

func (b *guildAuditLogsBuilder) SetActionType(actionType uint) *guildAuditLogsBuilder {
	b.r.param("action_type", actionType)
	return b
}

func (b *guildAuditLogsBuilder) SetBefore(before Snowflake) *guildAuditLogsBuilder {
	b.r.addPrereq(before.Empty(), "before can not be 0")
	b.r.param("before", before)
	return b
}

func (b *guildAuditLogsBuilder) SetLimit(limit int) *guildAuditLogsBuilder {
	b.r.param("limit", limit)
	return b
}

func (b *guildAuditLogsBuilder) Execute() (log *AuditLog, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*AuditLog), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateChannelBuilder) IgnoreCache() *updateChannelBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateChannelBuilder) CancelOnRatelimit() *updateChannelBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateChannelBuilder) URLParam(name string, v interface{}) *updateChannelBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateChannelBuilder) Set(name string, v interface{}) *updateChannelBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateChannelBuilder) SetParentID(parentID Snowflake) *updateChannelBuilder {
	b.r.addPrereq(parentID.Empty(), "parentID can not be 0")
	b.r.param("parent_id", parentID)
	return b
}

func (b *updateChannelBuilder) SetPermissionOverwrites(permissionOverwrites []PermissionOverwrite) *updateChannelBuilder {
	b.r.param("permission_overwrites", permissionOverwrites)
	return b
}

func (b *updateChannelBuilder) SetUserLimit(userLimit uint) *updateChannelBuilder {
	b.r.param("user_limit", userLimit)
	return b
}

func (b *updateChannelBuilder) SetBitrate(bitrate uint) *updateChannelBuilder {
	b.r.param("bitrate", bitrate)
	return b
}

func (b *updateChannelBuilder) SetRateLimitPerUser(rateLimitPerUser uint) *updateChannelBuilder {
	b.r.param("rate_limit_per_user", rateLimitPerUser)
	return b
}

func (b *updateChannelBuilder) SetNsfw(nsfw bool) *updateChannelBuilder {
	b.r.param("nsfw", nsfw)
	return b
}

func (b *updateChannelBuilder) SetTopic(topic string) *updateChannelBuilder {
	b.r.param("topic", topic)
	return b
}

func (b *updateChannelBuilder) SetPosition(position int) *updateChannelBuilder {
	b.r.param("position", position)
	return b
}

func (b *updateChannelBuilder) SetName(name string) *updateChannelBuilder {
	b.r.param("name", name)
	return b
}

func (b *updateChannelBuilder) Execute() (channel *Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Channel), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createGuildEmojiBuilder) IgnoreCache() *createGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createGuildEmojiBuilder) CancelOnRatelimit() *createGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createGuildEmojiBuilder) URLParam(name string, v interface{}) *createGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createGuildEmojiBuilder) Set(name string, v interface{}) *createGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

func (b *createGuildEmojiBuilder) SetRoles(roles []Snowflake) *createGuildEmojiBuilder {
	b.r.param("roles", roles)
	return b
}

func (b *createGuildEmojiBuilder) Execute() (emoji *Emoji, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Emoji), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildEmojiBuilder) IgnoreCache() *updateGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildEmojiBuilder) CancelOnRatelimit() *updateGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildEmojiBuilder) URLParam(name string, v interface{}) *updateGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildEmojiBuilder) Set(name string, v interface{}) *updateGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateGuildEmojiBuilder) SetName(name string) *updateGuildEmojiBuilder {
	b.r.param("name", name)
	return b
}

func (b *updateGuildEmojiBuilder) SetRoles(roles []Snowflake) *updateGuildEmojiBuilder {
	b.r.param("roles", roles)
	return b
}

func (b *updateGuildEmojiBuilder) Execute() (emoji *Emoji, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Emoji), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildBuilder) IgnoreCache() *updateGuildBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildBuilder) CancelOnRatelimit() *updateGuildBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildBuilder) URLParam(name string, v interface{}) *updateGuildBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildBuilder) Set(name string, v interface{}) *updateGuildBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateGuildBuilder) SetName(name string) *updateGuildBuilder {
	b.r.param("name", name)
	return b
}

func (b *updateGuildBuilder) SetRegion(region string) *updateGuildBuilder {
	b.r.param("region", region)
	return b
}

func (b *updateGuildBuilder) SetVerificationLevel(verificationLevel int) *updateGuildBuilder {
	b.r.param("verification_level", verificationLevel)
	return b
}

func (b *updateGuildBuilder) SetDefaultMessageNotifications(defaultMessageNotifications DefaultMessageNotificationLvl) *updateGuildBuilder {
	b.r.param("default_message_notifications", defaultMessageNotifications)
	return b
}

func (b *updateGuildBuilder) SetExplicitContentFilter(explicitContentFilter ExplicitContentFilterLvl) *updateGuildBuilder {
	b.r.param("explicit_content_filter", explicitContentFilter)
	return b
}

func (b *updateGuildBuilder) SetAfkChannelID(afkChannelID Snowflake) *updateGuildBuilder {
	b.r.addPrereq(afkChannelID.Empty(), "afkChannelID can not be 0")
	b.r.param("afk_channel_id", afkChannelID)
	return b
}

func (b *updateGuildBuilder) SetAfkTimeout(afkTimeout int) *updateGuildBuilder {
	b.r.param("afk_timeout", afkTimeout)
	return b
}

func (b *updateGuildBuilder) SetIcon(icon string) *updateGuildBuilder {
	b.r.param("icon", icon)
	return b
}

func (b *updateGuildBuilder) SetOwnerID(ownerID Snowflake) *updateGuildBuilder {
	b.r.addPrereq(ownerID.Empty(), "ownerID can not be 0")
	b.r.param("owner_id", ownerID)
	return b
}

func (b *updateGuildBuilder) SetSplash(splash string) *updateGuildBuilder {
	b.r.param("splash", splash)
	return b
}

func (b *updateGuildBuilder) SetSystemChannelID(systemChannelID Snowflake) *updateGuildBuilder {
	b.r.addPrereq(systemChannelID.Empty(), "systemChannelID can not be 0")
	b.r.param("system_channel_id", systemChannelID)
	return b
}

func (b *updateGuildBuilder) Execute() (guild *Guild, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Guild), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildEmbedBuilder) IgnoreCache() *updateGuildEmbedBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildEmbedBuilder) CancelOnRatelimit() *updateGuildEmbedBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildEmbedBuilder) URLParam(name string, v interface{}) *updateGuildEmbedBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildEmbedBuilder) Set(name string, v interface{}) *updateGuildEmbedBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateGuildEmbedBuilder) SetEnabled(enabled bool) *updateGuildEmbedBuilder {
	b.r.param("enabled", enabled)
	return b
}

func (b *updateGuildEmbedBuilder) SetChannelID(channelID Snowflake) *updateGuildEmbedBuilder {
	b.r.addPrereq(channelID.Empty(), "channelID can not be 0")
	b.r.param("channel_id", channelID)
	return b
}

func (b *updateGuildEmbedBuilder) Execute() (embed *GuildEmbed, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*GuildEmbed), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildMemberBuilder) IgnoreCache() *updateGuildMemberBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildMemberBuilder) CancelOnRatelimit() *updateGuildMemberBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildMemberBuilder) URLParam(name string, v interface{}) *updateGuildMemberBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildMemberBuilder) Set(name string, v interface{}) *updateGuildMemberBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateGuildMemberBuilder) SetNick(nick string) *updateGuildMemberBuilder {
	b.r.param("nick", nick)
	return b
}

func (b *updateGuildMemberBuilder) SetRoles(roles []Snowflake) *updateGuildMemberBuilder {
	b.r.param("roles", roles)
	return b
}

func (b *updateGuildMemberBuilder) SetMute(mute bool) *updateGuildMemberBuilder {
	b.r.param("mute", mute)
	return b
}

func (b *updateGuildMemberBuilder) SetDeaf(deaf bool) *updateGuildMemberBuilder {
	b.r.param("deaf", deaf)
	return b
}

func (b *updateGuildMemberBuilder) SetChannelID(channelID Snowflake) *updateGuildMemberBuilder {
	b.r.addPrereq(channelID.Empty(), "channelID can not be 0")
	b.r.param("channel_id", channelID)
	return b
}

func (b *updateGuildMemberBuilder) Execute() (err error) {
	_, err = b.r.execute()
	return
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateMessageBuilder) IgnoreCache() *updateMessageBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateMessageBuilder) CancelOnRatelimit() *updateMessageBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateMessageBuilder) URLParam(name string, v interface{}) *updateMessageBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateMessageBuilder) Set(name string, v interface{}) *updateMessageBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateMessageBuilder) SetContent(content string) *updateMessageBuilder {
	b.r.param("content", content)
	return b
}

func (b *updateMessageBuilder) SetEmbed(embed *Embed) *updateMessageBuilder {
	b.r.param("embed", embed)
	return b
}

func (b *updateMessageBuilder) Execute() (message *Message, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Message), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *basicBuilder) IgnoreCache() *basicBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *basicBuilder) CancelOnRatelimit() *basicBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *basicBuilder) URLParam(name string, v interface{}) *basicBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *basicBuilder) Set(name string, v interface{}) *basicBuilder {
	b.r.body[name] = v
	return b
}

func (b *basicBuilder) Execute() (err error) {
	_, err = b.r.execute()
	return
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildRoleBuilder) IgnoreCache() *updateGuildRoleBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildRoleBuilder) CancelOnRatelimit() *updateGuildRoleBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildRoleBuilder) URLParam(name string, v interface{}) *updateGuildRoleBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildRoleBuilder) Set(name string, v interface{}) *updateGuildRoleBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateGuildRoleBuilder) SetName(name string) *updateGuildRoleBuilder {
	b.r.param("name", name)
	return b
}

func (b *updateGuildRoleBuilder) SetPermissions(permissions PermissionBits) *updateGuildRoleBuilder {
	b.r.param("permissions", permissions)
	return b
}

func (b *updateGuildRoleBuilder) SetColor(color uint) *updateGuildRoleBuilder {
	b.r.param("color", color)
	return b
}

func (b *updateGuildRoleBuilder) SetHoist(hoist bool) *updateGuildRoleBuilder {
	b.r.param("hoist", hoist)
	return b
}

func (b *updateGuildRoleBuilder) SetMentionable(mentionable bool) *updateGuildRoleBuilder {
	b.r.param("mentionable", mentionable)
	return b
}

func (b *updateGuildRoleBuilder) Execute() (role *Role, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Role), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createDMBuilder) IgnoreCache() *createDMBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createDMBuilder) CancelOnRatelimit() *createDMBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createDMBuilder) URLParam(name string, v interface{}) *createDMBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createDMBuilder) Set(name string, v interface{}) *createDMBuilder {
	b.r.body[name] = v
	return b
}

func (b *createDMBuilder) Execute() (channel *Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Channel), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createGroupDMBuilder) IgnoreCache() *createGroupDMBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createGroupDMBuilder) CancelOnRatelimit() *createGroupDMBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createGroupDMBuilder) URLParam(name string, v interface{}) *createGroupDMBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createGroupDMBuilder) Set(name string, v interface{}) *createGroupDMBuilder {
	b.r.body[name] = v
	return b
}

func (b *createGroupDMBuilder) Execute() (channel *Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Channel), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getCurrentUserGuildsBuilder) IgnoreCache() *getCurrentUserGuildsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getCurrentUserGuildsBuilder) CancelOnRatelimit() *getCurrentUserGuildsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getCurrentUserGuildsBuilder) URLParam(name string, v interface{}) *getCurrentUserGuildsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getCurrentUserGuildsBuilder) Set(name string, v interface{}) *getCurrentUserGuildsBuilder {
	b.r.body[name] = v
	return b
}

func (b *getCurrentUserGuildsBuilder) SetBefore(before Snowflake) *getCurrentUserGuildsBuilder {
	b.r.addPrereq(before.Empty(), "before can not be 0")
	b.r.param("before", before)
	return b
}

func (b *getCurrentUserGuildsBuilder) SetAfter(after Snowflake) *getCurrentUserGuildsBuilder {
	b.r.addPrereq(after.Empty(), "after can not be 0")
	b.r.param("after", after)
	return b
}

func (b *getCurrentUserGuildsBuilder) SetLimit(limit int) *getCurrentUserGuildsBuilder {
	b.r.param("limit", limit)
	return b
}

func (b *getCurrentUserGuildsBuilder) Execute() (guilds []*Guild, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	tmp := v.(*[]*Guild)
	return *tmp, nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getUserBuilder) IgnoreCache() *getUserBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getUserBuilder) CancelOnRatelimit() *getUserBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getUserBuilder) URLParam(name string, v interface{}) *getUserBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getUserBuilder) Set(name string, v interface{}) *getUserBuilder {
	b.r.body[name] = v
	return b
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getUserConnectionsBuilder) IgnoreCache() *getUserConnectionsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getUserConnectionsBuilder) CancelOnRatelimit() *getUserConnectionsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getUserConnectionsBuilder) URLParam(name string, v interface{}) *getUserConnectionsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getUserConnectionsBuilder) Set(name string, v interface{}) *getUserConnectionsBuilder {
	b.r.body[name] = v
	return b
}

func (b *getUserConnectionsBuilder) Execute() (cons []*UserConnection, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	tmp := v.(*[]*UserConnection)
	return *tmp, nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *getUserDMsBuilder) IgnoreCache() *getUserDMsBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *getUserDMsBuilder) CancelOnRatelimit() *getUserDMsBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *getUserDMsBuilder) URLParam(name string, v interface{}) *getUserDMsBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *getUserDMsBuilder) Set(name string, v interface{}) *getUserDMsBuilder {
	b.r.body[name] = v
	return b
}

func (b *getUserDMsBuilder) Execute() (channels []*Channel, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	tmp := v.(*[]*Channel)
	return *tmp, nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateCurrentUserBuilder) IgnoreCache() *updateCurrentUserBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateCurrentUserBuilder) CancelOnRatelimit() *updateCurrentUserBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateCurrentUserBuilder) URLParam(name string, v interface{}) *updateCurrentUserBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateCurrentUserBuilder) Set(name string, v interface{}) *updateCurrentUserBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateCurrentUserBuilder) SetUsername(username string) *updateCurrentUserBuilder {
	b.r.param("username", username)
	return b
}

func (b *updateCurrentUserBuilder) SetAvatar(avatar string) *updateCurrentUserBuilder {
	b.r.param("avatar", avatar)
	return b
}

func (b *updateCurrentUserBuilder) Execute() (user *User, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*User), nil
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateWebhookBuilder) IgnoreCache() *updateWebhookBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateWebhookBuilder) CancelOnRatelimit() *updateWebhookBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateWebhookBuilder) URLParam(name string, v interface{}) *updateWebhookBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateWebhookBuilder) Set(name string, v interface{}) *updateWebhookBuilder {
	b.r.body[name] = v
	return b
}

func (b *updateWebhookBuilder) SetName(name string) *updateWebhookBuilder {
	b.r.param("name", name)
	return b
}

func (b *updateWebhookBuilder) SetAvatar(avatar string) *updateWebhookBuilder {
	b.r.param("avatar", avatar)
	return b
}

func (b *updateWebhookBuilder) SetChannelID(channelID Snowflake) *updateWebhookBuilder {
	b.r.addPrereq(channelID.Empty(), "channelID can not be 0")
	b.r.param("channel_id", channelID)
	return b
}

func (b *updateWebhookBuilder) Execute() (webhook *Webhook, err error) {
	var v interface{}
	if v, err = b.r.execute(); err != nil {
		return nil, err
	}
	return v.(*Webhook), nil
}
