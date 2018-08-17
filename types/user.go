package types

type User struct {
	Address       string
	Email         []string `json:"email,omitempty"`
	Phone         []string `json:"phone,omitempty"`
	Name          string   `json:"name,omitempty"`
	NickName      string   `json:"nick_name,omitempty"`
	IDCard        string   `json:"id_card,omitempty"`
	CreditScore   int64    `json:"credit_score,omitempty"`
	Game          []string `json:"game,omitempty"`
	Social        []string `json:"social,omitempty"`
	Job           []string `json:"job,omitempty"`
	Office        []string `json:"office,omitempty"`
	WeChat        []string `json:"wechat,omitempty"`
	Interest      []string `json:"interest,omitempty"`
	Profession    []string
	character     []string
	constellation []string
	birthday      []string
	Mantra        []string
	Like          []string
	hate          []string
	KeyWord       []string
	Events        []string
	BankCard      []string
	CreditCard    []string
	ProvidentFund []string
	education     []string
	Extra         map[string]string
}
