package api

type Message struct {
	Type   string
	Text   string
	IsHTML bool
}

func GetSystemMessages() []Message {
	msgs := []Message{}

	// Type can be: success, notice, warning, error
	msgs = append(msgs, Message{Type: "warning", IsHTML: true, Text: "<b>管理员消息</b>： 同一个人的项目是有优先级的，大家按照从上到下的顺序点哈，比如列表里显示一个人有1、2、3、4共四个项目，如果你只想赞两个，那就从上面开始点1、2，不要点下面的3、4。否则将来有可能将此行为按照无效点赞处理。"})
	return msgs
}
