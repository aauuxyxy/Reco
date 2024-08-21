package model

// フォロー
// フォロワー数計算時はFollowIDに自身のIDが登録されている数＋Stateが2or3の数
type Follow struct {
	UsersID uint
	FollowID uint
	State int //0：両者フォロー解除、1：フォロー側フォロー、2：フォロワー側フォロー、3：相互フォロー
	BlockTarget int // 0：ブロック無し、1：フォロワー側ブロック、2：フォロー側ブロック
}