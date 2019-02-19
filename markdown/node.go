package main


type tree struct {

}

type node struct {
	syntax       string
	attr      map[string][]string
	value     string
	children  []*node
}

type block struct {
	node
}

type

const (
	tagArticle    = "article"
	tagH1         = "h1"
	tagH2         = "h2"
	tagH3         = "h3"
	tagH4         = "h4"
	tagH5         = "h5"
	tagH6         = "h6"
	tagP          = "p"
	tagStrong     = "strong"
	tagEm         = "em"
	tagBlockquote = "blockquote"
	tagUl         = "ul"
	tagOl         = "ol"
	tagLi         = "li"
	tagBr         = "br"
	tagA          = "a"
	tagImg        = "img"
	tagPre        = "pre"
	tagCode       = "Code"
)

func newRoot() *node {
	return &node{
		tag: tagArticle,
	}
}

func newNode(raw string) *node{
	node:= &node{}
	node.parse(text)
	return node
}

func (n *node) parse(text) {
	lines := strings.Split(input, "\n")
	for i
}


func (n *node) addChild(raw){
	n.children = append(n.children, newNode(raw))
}


func renderLine(line string, index int, flag *string) (output string) {

	if text, err := testAndTrimPrefix(line, "###### "); err != nil {
		return withTag(text, tagH6 )
	}
	if text, err := testAndTrimPrefix(line, "##### "); err != nil {
		return withTag(text, tagH5)
	}
	if text, err := testAndTrimPrefix(line, "#### "); err != nil {
		return withTag(text, tagH4)
	}
	if text, err := testAndTrimPrefix(line, "### "); err != nil {
		return withTag(text, tagH3)
	}
	if text, err := testAndTrimPrefix(line, "## "); err != nil {
		return withTag(text, tagH2)
	}
	if text, err := testAndTrimPrefix(line, "# "); err != nil {
		return withTag(text, tagH1)
	}
	if text, err := testAndTrimPrefix(line, "< "); err != nil {
		return withTag(text, tagBlockquote)
	}
	return withTag(line, tagP)
}

func withTag(text, tag string, flag) (output string) {

	return fmt.Sprintf("<%s>%s</%s>", tag, text, tag)
}

func testAndTrimPrefix(input, prefix string) (output string, err error) {
	if strings.HasPrefix(input, prefix) {
		return strings.TrimPrefix(input, prefix), nil
	}
	return "", errors.New("prefix not found")
}

// </br>
// if strings.TrimSpace(n.children) == "" {
// 	n.tag = "br"
// }
// >blockquote 区块引用
// ``` code

// +-表示无序列表
// 三个以上___---***分割线
// [超链接](/dsada)。
// ![图片](a.jpg)
// `code`
// **Strong**一星斜体,二星粗体,三星斜粗
// 段落

// func parseFromPrefixAndSuffix(prefix, suffixtag string) {
// 	if strings.HasPrefix(n.raw, prefix) && strings.HasSuffix(n.raw, prefix) {

// 	}

// // replace strong
// regStrong := regexp.MustCompile(`(?<=\*\*).*(?=\*\*)`)
// raw = regStrong.ReplaceAllStringFunc(raw, func(s string) string {
// 	return "<strong>" + s + "<strong>"
// })

// // replace em
// regEm := regexp.MustCompile(`(?<=\*).*(?=\*)`)
// raw = regEm.ReplaceAllStringFunc(raw, func(s string) string {
// 	return "<em>" + s + "<em>"
// })

// replace link


// type attribute map[string][]string

// func (a attribute) toString() string {
// 	s := ""
// 	for key, values := range a {
// 		valueStr := strings.Join(values, " ")
// 		pair := fmt.Sprintf(" %s=\"%s\"", key, valueStr)
// 		s += pair
// 	}
// 	return s
// }

// func (n *node) toHTML() string {
// 	attrStr := n.attr.toString()
// 	innerHTML := ""
// 	for _, child := range n.children {
// 		innerHTML += child.toHTML()
// 	}
// 	return fmt.Sprintf("<%s%s>%s</%s>", n.tag, attrStr, innerHTML, n.tag)
// }
