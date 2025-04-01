package entity

type Topic struct {
	ID          uint   `bson:"id"`
	TextExplain string `bson:"text_explain"`
	Desc        string `bson:"desc"`
}

type GraphTopic struct {
	UID         string `json:"uid,omitempty"`
	ID          uint   `json:"id,omitempty"`
	TextExplain string `json:"text_explain,omitempty"`
	Desc        string `json:"desc,omitempty"`
	DgraphType  string `json:"dgraph.type,omitempty"`
}

func (g *GraphTopic) ToBase() *Topic {
	return &Topic{
		ID:          g.ID,
		TextExplain: g.TextExplain,
    Desc: g.Desc,
	}
}

func (g *GraphTopic) ToGraph() *GraphTopic {
  return &GraphTopic{
    UID: ,
  }
}
