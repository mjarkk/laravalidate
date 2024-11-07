package laravalidate

type templateVariableT struct {
	from int
	to   int
}

type messageTemplateParserT struct {
	msg       []byte
	idx       int
	variables []templateVariableT
}

func parseMsgTemplate(msg []byte) []templateVariableT {
	if len(msg) == 0 {
		return []templateVariableT{}
	}

	parser := &messageTemplateParserT{
		msg:       msg,
		idx:       0,
		variables: []templateVariableT{},
	}

	parser.MightParseNextVariable()

	return parser.variables
}

func (p *messageTemplateParserT) Next() bool {
	return p.idx < len(p.msg)
}

func (p *messageTemplateParserT) C() byte {
	return p.msg[p.idx]
}

func (p *messageTemplateParserT) MightParseVariable() {
	startIdx := p.idx
	for p.Next() {
		c := p.C()
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_' {
			p.idx++
			continue
		}
		break
	}
	endIdx := p.idx
	if startIdx != endIdx {
		p.variables = append(p.variables, templateVariableT{from: startIdx - 1, to: endIdx})
	}

	p.SearchForNextVariable()
}

func (p *messageTemplateParserT) MightParseNextVariable() {
	for p.Next() {
		c := p.C()
		p.idx++
		switch c {
		case ':':
			p.MightParseVariable()
		case ' ', '\t', '\n', '\r':
			// Do nothing
		default:
			p.SearchForNextVariable()
			return
		}
	}
}

func (p *messageTemplateParserT) SearchForNextVariable() {
	for p.Next() {
		c := p.C()
		p.idx++
		switch c {
		case ' ', '\t', '\n', '\r':
			p.MightParseNextVariable()
			return
		}
	}
}
