package channels

// Channelsetup used to initialie channel configuration
func ChannelSetup(config Config) *Channel {

	return &Channel{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Authenticate:     config.Authenticate,
	}

}

// get all channel list
func (channel Channel) ListChannel(limit, offset int, filter Filter, activestatus bool) (channelList []tblchannel, channelcount int, err error) {

	if channel.AuthEnable && !channel.AuthFlg {

		return []tblchannel{}, 0, ErrorAuth
	}

	if channel.PermissionEnable && !channel.PermissionFlg {

		return []tblchannel{}, 0, ErrorPermission

	}

	channellist, _, _ := CH.Channellist(limit, offset, filter, activestatus, channel.DB)

	var chnallist []tblchannel

	for _, val := range channellist {

		val.SlugName = val.ChannelDescription

		val.ChannelDescription = TruncateDescription(val.ChannelDescription, 130)

		// entrcount, _ := CH.ChannelEntryList(&[]TblChannelEntries{}, 0, 0, val.Id, EntriesFilter{}, false, roleid, false, channel.DBString)

		// val.EntriesCount = int(entrcount)

		chnallist = append(chnallist, val)

	}

	_, chcount, _ := CH.Channellist(0, 0, filter, activestatus, channel.DB)

	return chnallist, int(chcount), nil
}

/*Get channel by name*/
func (channel Channel) GetchannelByName(channelname string) (channels tblchannel, err error) {

	if channel.AuthEnable && !channel.AuthFlg {

		return tblchannel{}, ErrorAuth
	}

	if channel.PermissionEnable && !channel.PermissionFlg {

		return tblchannel{}, ErrorPermission

	}

	var channellist tblchannel

	err1 := CH.GetChannelByChannelName(&channellist, channelname, channel.DB)

	if err1 != nil {

		return tblchannel{}, err1
	}

	return channellist, nil

}
