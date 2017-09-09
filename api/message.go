package api

type Message struct {
	Type   string
	Text   string
	IsHTML bool
}

func GetSystemMessages(user string) []Message {
	msgs := []Message{}

	if user != "" {
		objUser := GetUser(user)
		if objUser.IsFlagged {
			msgs = append(msgs, Message{Type: "error", IsHTML: true, Text: `该账户所使用的点赞账号已经被GitHub Flagged，所有点赞清零。请联系GitStar官方QQ群获取解决方案：<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=3be0e33b7f8fd6da953be4d06c9c1d81b6ce2d93e126898087681e43d32ec0d9"><img border="0" src="//pub.idqqimg.com/wpa/images/group.png" alt="GitHub Star项目点赞" title="GitHub Star项目点赞"></a> (群号：646373152)。`})
		}
	}

	// Type can be: success, notice, warning, error
	msgs = append(msgs, Message{Type: "success", IsHTML: true, Text: `<b>置顶</b>：为了方便大家@不回赞行为，更好地服务广大GitStar平台用户，本平台建立了GitStar官方QQ群：<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=3be0e33b7f8fd6da953be4d06c9c1d81b6ce2d93e126898087681e43d32ec0d9"><img border="0" src="//pub.idqqimg.com/wpa/images/group.png" alt="GitHub Star项目点赞" title="GitHub Star项目点赞"></a> (群号：646373152)，原则上所有平台用户都需要在一周内完成加群操作，逾期不加入的用户可能会被封号处理。 —管理员团队`})
	// msgs = append(msgs, Message{Type: "warning", IsHTML: true, Text: "<b>管理员消息</b>： 同一个人的项目是有优先级的，大家按照从上到下的顺序点哈，比如列表里显示一个人有1、2、3、4共四个项目，如果你只想赞两个，那就从上面开始点1、2，不要点下面的3、4。否则将来有可能将此行为按照无效点赞处理。"})
	return msgs
}
