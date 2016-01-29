package widget

type Button struct {
	Widget
}

func (this *Button) SetContent(label string) {
	this.content += "╭"
	for _,_ = range label { this.content += "─" }
	this.content += "╮\n"

	this.content += "│"
	this.content += label
	this.content += "│\n"

	this.content += "╰"
	for _,_ = range label { this.content += "─" }
	this.content += "╯\n"
}

func (this *Widget) Draw() string {
	var out string
	
	//DEFAULT COLOR
	if this.foreground == 0 { this.foreground = DEFAULT }
	if this.background == 0 { this.background = DEFAULT }
	
	//DEFINE POSITION
	out = "\033["+strconv.Itoa(int(this.canvas.X))+";"+strconv.Itoa(int(this.canvas.Y))+"H"
	
	//DEFINE STYLE
	out += "\033["
	out += strconv.Itoa(this.style)+";"
	out += "3"+strconv.Itoa(this.foreground-1)+";"
	out += "4"+strconv.Itoa(this.background-1)+"m"
	
	//SET CONTENT
	out += this.content
	
	//UNDEFINE STYLE
	out += "\033[0m"
	
	return out
}