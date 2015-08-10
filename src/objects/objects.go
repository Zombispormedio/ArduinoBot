package objects


const Url_pikabot="https://pikabot.herokuapp.com/switch/"

type Messages struct {
	State bool		`json:state`
	Data []string `json:data`
}
