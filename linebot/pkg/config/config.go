package config

// Config - LINE botの設定
type Config struct {
	Host                      string // LINEのホスト(スキームも含む)
	TalkServicePath           string // TalkServiceのパス
	TalkServicePathForPolling string // TalkServiceのパス(Polling用)
	QrPath                    string // QRログインのパス
	PermitNoticePath          string // クライアントへログイン処理させるためのパス
	UserAgent                 string // User-Agentヘッダ
	LINEApp                   string // X-Line-Applicationヘッダ
}

// NewConfig - Configのコンストラクタ
func NewConfig() Config {
	return Config{
		Host:                      "https://gwz.line.naver.jp",
		TalkServicePath:           "/S4",
		TalkServicePathForPolling: "/P4",
		QrPath:                    "/acct/lgn/sq/v1",
		PermitNoticePath:          "/acct/lp/lgn/sq/v1",
		UserAgent:                 "Line/2.5.5",
		LINEApp:                   "CHROMEOS\t2.5.5\tChrome OS\t1;SECONDARY",
	}
}
