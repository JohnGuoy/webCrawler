package model

type Profile struct {
	Name              string
	Gender            string
	Age               int
	Height            int
	Weight            int
	Income            string // 工资下限和上限
	Marriage          string
	Education         string
	Occupation        string
	Residence         string // 居住地（工作地）
	Hukou             string // 籍贯
	Xinzuo            string
	House             string
	Car               string
	Introduction      string
	IndividualInfoUrl string
}
