package types

type File struct {
	//ID        string `json:"id"`
	Sha1Hash  string `json:"sha1_hash"`
	Mime      string `json:"mime"`
	RefCount  uint64 `json:"ref_count"`
	Size      uint64 `json:"size"`
	Status    int    `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

type Sutra struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	PlayedCount uint64 `json:"played_count"`
	ItemTotal   uint32 `json:"item_tokal"`
	CreatedAt   int64  `json:"created_at"`
}

type SutraItem struct {
	ID          string `json:"id"`
	SutraID     string `json:"sutra_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Original    string `json:"original"`
	Explanation string `json:"explanation"`
	AudioID     string `json:"audio_id"`
	LyricID     string `json:"lyric_id"`
	Lesson      uint32 `json:"lesson"` // 集数，根据此排序
	PlayedCount uint64 `json:"played_count"`
	Duration    uint32 `json:"duration"`
	CreatedAt   int64  `json:"created_at"`
}
