package widget

type Label struct{
	Widget
}

func (this *Label) SetContent(content string) {
	this.content = content+"\n"
}