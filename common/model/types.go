package model

type File struct {
	//ID        string `json:"id"`
	Sha1Hash  string `json:"sha1_hash"`
	Mime      string `json:"mime"`
	RefCount  uint64 `json:"ref_count"`
	Size      uint64 `json:"size"`
	Status    int    `json:"status"`
	CreatedAt int64  `json:"created_at"`
}
