package minimax

const (
	Abab5       = "abab5-chat"
	Abab5Dot5   = "abab5.5-chat"
	Abab5Dot5s  = "abab5.5s-chat"
	Abab6       = "abab6-chat"
	Embo01      = "embo-01"
	Speech01    = "speech-01"
	Speech02    = "speech-02"
	Speech01Pro = "speech-01-pro" // alias speech-01=speech-01-pro

	ModelBot            = "MM智能助理"
	ChatMessageRoleUser = "USER"
	ChatMessageRoleBot  = "BOT"
	EmbeddingsDbType    = "db"
	EmbeddingsQueryType = "query"

	ToolCodeInterpreter = "code_interpreter"
	ToolRetrieval       = "retrieval"
	ToolFunction        = "function"
	ToolWebSearch       = "web_search"
)

// voice_id：音色编号支，持系统音色(id)以及复刻音色（id）两种类型，其中系统音色（ID）如下：
const (
	VoiceMaleQnQingSe            = "male-qn-qingse"             // 青涩青年音色
	VoiceMaleQnJingYing          = "male-qn-jingying"           // 精英青年音色
	VoiceMaleQnBaDao             = "male-qn-badao"              // 霸道青年音色
	VoiceMaleQnDaXueSheng        = "male-qn-daxuesheng"         // 青年大学生音色
	VoiceFemaleShaoNv            = "female-shaonv"              // 少女音色
	VoiceFemaleYuJie             = "female-yujie"               // 御姐音色
	VoiceFemaleChengShu          = "female-chengshu"            // 成熟女性音色
	VoiceFemaleTianMei           = "female-tianmei"             // 甜美女性音色
	VoicePresenterMale           = "presenter_male"             // 男性主持人
	VoicePresenterFemale         = "presenter_female"           // 女性主持人
	VoiceAudiobookMale1          = "audiobook_male_1"           // 男性有声书1
	VoiceAudiobookMale2          = "audiobook_male_2"           // 男性有声书2
	VoiceAudiobookFemale1        = "audiobook_female_1"         // 女性有声书1
	VoiceAudiobookFemale2        = "audiobook_female_2"         // 女性有声书2
	VoiceMaleQnQingSeJingPin     = "male-qn-qingse-jingpin"     // 青涩青年音色-beta
	VoiceMaleQnJingYingJingPin   = "male-qn-jingying-jingpin"   // 精英青年音色-beta
	VoiceMaleQnBaDaoJingPin      = "male-qn-badao-jingpin"      // 霸道青年音色-beta
	VoiceMaleQnDaXueShengJingPin = "male-qn-daxuesheng-jingpin" // 青年大学生音色-beta
	VoiceFemaleShaoNvJingPin     = "female-shaonv-jingpin"      // 少女音色-beta
	VoiceFemaleYuJieJingPin      = "female-yujie-jingpin"       // 御姐音色-beta
	VoiceFemaleChengShuJingPin   = "female-chengshu-jingpin"    // 成熟女性音色-beta
	VoiceFemaleTianMeiJingPin    = "female-tianmei-jingpin"     // 甜美女性音色-beta
)
