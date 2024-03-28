package model

var (
	ConfigStatusNormal = 1
	ConfigStatusHidden = 2
)

type SortExp []*Experiment

func (a SortExp) Len() int {
	return len(a)
}

func (a SortExp) len() int {
	return len(a)
}

func (a SortExp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SortExp) Less(i, j int) bool {
	iWeight := 0
	jWeight := 0
	if a[i].Platform != "" {
		iWeight += 1
	}
	if a[i].Channel != "" {
		iWeight += 3
	}
	if a[i].Version != "" {
		iWeight += 5
	}

	if a[j].Platform != "" {
		jWeight += 1
	}
	if a[j].Channel != "" {
		jWeight += 3
	}
	if a[j].Version != "" {
		jWeight += 5
	}
	return iWeight > jWeight
}

type Experiment struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Params      string `json:"params"`
	Desc        string `json:"desc"`
	Platform    string `json:"platform"`
	Channel     string `json:"channel"`
	Version     string `json:"version"`
	Layer       int    `json:"layer"`
	Status      int    `json:"status"`
	StartBucket int    `json:"startBucket"`
	EndBucket   int    `json:"endBucket"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
