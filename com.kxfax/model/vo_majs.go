package model 

/**
  vo log for majs
'{"remote_addr":"$remote_addr"'
                                  ',"remote_user":"$remote_user"'
                                  ',"time_local":"$time_local"'
                                  ',"request":"$request"'
                                  ',"status":"$status"'
                                  ',"body_bytes_sent":"$body_bytes_sent"'
                                  ',"http_referer":"$http_referer"'
                                  ',"http_user_agent":"$http_user_agent"'
                                  ',"http_x_forwarded_for":"$http_x_forwarded_for"'
                                  ',"u_domain":"$u_domain"'
                                  ',"u_url":"$u_url"'
                                  ',"u_title":"$u_title"'
                                  ',"u_referrer":"$u_referrer"'
                                  ',"u_sh":"$u_sh"'
                                  ',"u_sw":"$u_sw"'
                                  ',"u_cd":"$u_cd"'
                                  ',"u_lang":"$u_lang"'
                                  ',"u_utrace":"$u_utrace"'
                                  ',"u_account":"$u_account"'
                                  ',"u_loanid":"$u_loanid"'
                                  ',"u_extend":"$u_extend"}';

**/
type Majs struct {
	Remote_addr          string `json:"remote_addr"`
	Remote_user          string `json:"remote_user"`
	Time_local           string `json:"time_local"`
	Request              string `json:"request"`
	Status               string `json:"status"`
	Body_bytes_sent      string `json:"body_bytes_sent"`
	Http_referer         string `json:"http_referer"`
	Http_user_agent      string `json:"http_user_agent"`
	Http_x_forwarded_for string `json:"http_x_forwarded_for"`
	U_domain             string `json:"u_domain"`
	U_url                string `json:"u_url"`
	U_title              string `json:"u_title"`
	U_referrer           string `json:"u_referrer"`
	U_sh                 string `json:"u_sh"`
	U_sw                 string `json:"u_sw"`
	U_cd                 string `json:"u_cd"`
	U_lang               string `json:"u_lang"`
	U_utrace             string `json:"u_utrace"`
	U_account            string `json:"u_account"`
	U_loanid             string `json:"u_loanid"`
	U_extend             string `json:"u_extend"`
}



